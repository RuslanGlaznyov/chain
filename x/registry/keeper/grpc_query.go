package keeper

import (
	"kyve/x/registry/types"
)

var _ types.QueryServer = Keeper{}
