package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RedelegatePool lets a user redelegate from one staker to another staker
func (k msgServer) RedelegatePool(goCtx context.Context, msg *types.MsgRedelegatePool) (*types.MsgRedelegatePoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO implement cooldown.

	// Perform undelegation
	if err := k.Undelegate(ctx, msg.FromStaker, msg.FromPoolId, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// Perform undelegation
	if err := k.Delegate(ctx, msg.ToStaker, msg.ToPoolId, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// Emit a delegation event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventRedelegatePool{
		Address:  msg.Creator,
		FromPool: msg.FromPoolId,
		FromNode: msg.FromStaker,
		ToPool:   msg.ToPoolId,
		ToNode:   msg.ToStaker,
		Amount:   msg.Amount,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgRedelegatePoolResponse{}, nil
}
