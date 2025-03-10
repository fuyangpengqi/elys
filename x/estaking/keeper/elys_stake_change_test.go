package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *EstakingKeeperTestSuite) TestElysStakeChange() {
	testCases := []struct {
		name                 string
		prerequisiteFunction func() (addr sdk.AccAddress)
		postValidateFunction func(addr sdk.AccAddress)
	}{
		{
			"set and get elys stake change",
			func() (addr sdk.AccAddress) {
				suite.ResetSuite()

				addr = suite.AddAccounts(1, nil)[0]

				// set elys stake change
				suite.app.EstakingKeeper.SetElysStakeChange(suite.ctx, addr)

				return addr
			},
			func(addr sdk.AccAddress) {
				found := suite.app.EstakingKeeper.GetElysStakeChange(suite.ctx, addr)
				suite.Require().True(found)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			addr := tc.prerequisiteFunction()
			tc.postValidateFunction(addr)
		})
	}
}
