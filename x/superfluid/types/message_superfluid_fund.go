package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSuperfluidFund = "superfluid_fund"

var _ sdk.Msg = &MsgSuperfluidFund{}

func NewMsgSuperfluidFund(creator string, poolId uint64, valAddr string, amount uint64) *MsgSuperfluidFund {
	return &MsgSuperfluidFund{
		Creator: creator,
		PoolId:  poolId,
		ValAddr: valAddr,
		Amount:  amount,
	}
}

func (msg *MsgSuperfluidFund) Route() string {
	return RouterKey
}

func (msg *MsgSuperfluidFund) Type() string {
	return TypeMsgSuperfluidFund
}

func (msg *MsgSuperfluidFund) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSuperfluidFund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSuperfluidFund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
