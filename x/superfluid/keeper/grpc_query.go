package keeper

import (
	"github.com/KYVENetwork/chain/x/superfluid/types"
)

var _ types.QueryServer = Keeper{}
