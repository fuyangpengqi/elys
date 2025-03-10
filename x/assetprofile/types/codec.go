package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgUpdateEntry{}, "assetprofile/MsgUpdateEntry")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteEntry{}, "assetprofile/MsgDeleteEntry")
	legacy.RegisterAminoMsg(cdc, &MsgAddEntry{}, "assetprofile/MsgAddEntry")
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddEntry{},
		&MsgUpdateEntry{},
		&MsgDeleteEntry{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
