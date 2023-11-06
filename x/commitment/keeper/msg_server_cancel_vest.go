package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/elys-network/elys/x/commitment/types"
)

func (k msgServer) CancelVest(goCtx context.Context, msg *types.MsgCancelVest) (*types.MsgCancelVestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	vestingInfo, _ := k.GetVestingInfo(ctx, msg.Denom)

	if vestingInfo == nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom: %s", msg.Denom)
	}

	// Get the Commitments for the creator
	commitments, found := k.GetCommitments(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCommitmentsNotFound, "creator: %s", msg.Creator)
	}

	remainingToCancel := msg.Amount

	newVestingTokens := make([]*types.VestingTokens, 0, len(commitments.VestingTokens))

	for _, vesting := range commitments.VestingTokens {
		cancelAmount := sdk.MinInt(remainingToCancel, vesting.UnvestedAmount)
		vesting.TotalAmount = vesting.TotalAmount.Sub(cancelAmount)
		vesting.UnvestedAmount = vesting.UnvestedAmount.Sub(cancelAmount)

		if !vesting.TotalAmount.IsZero() {
			newVestingTokens = append(newVestingTokens, vesting)
		}

		remainingToCancel = remainingToCancel.Sub(cancelAmount)
		if remainingToCancel.IsZero() {
			break
		}
	}

	commitments.VestingTokens = newVestingTokens

	if !remainingToCancel.IsZero() {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientVestingTokens, "denom: %s, amount: %s", msg.Denom, msg.Amount)
	}

	// Update the uncommitted tokens amount
	uncommittedToken, found := commitments.GetUncommittedTokensForDenom(msg.Denom)

	if found {
		uncommittedToken.Amount = uncommittedToken.Amount.Add(msg.Amount)
	} else {
		uncommittedTokens := commitments.GetUncommittedTokens()
		uncommittedTokens = append(uncommittedTokens, &types.UncommittedTokens{
			Denom:  msg.Denom,
			Amount: msg.Amount,
		})
		commitments.UncommittedTokens = uncommittedTokens
	}

	k.SetCommitments(ctx, commitments)

	// Emit Hook commitment changed
	k.AfterCommitmentChange(ctx, msg.Creator, sdk.NewCoin(msg.Denom, msg.Amount))

	// Emit blockchain event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCommitmentChanged,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeDenom, msg.Denom),
		),
	)

	return &types.MsgCancelVestResponse{}, nil
}
