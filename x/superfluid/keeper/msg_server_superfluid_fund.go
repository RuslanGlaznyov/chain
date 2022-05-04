package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/superfluid/types"
)

// SuperfluidFund handles the logic of an SDK message that allows funders to superfluid fund a specified pool.
func (k msgServer) SuperfluidFund(
	goCtx context.Context, msg *types.MsgSuperfluidFund,
) (*types.MsgSuperfluidFundResponse, error) {
	return &types.MsgSuperfluidFundResponse{}, nil
}
