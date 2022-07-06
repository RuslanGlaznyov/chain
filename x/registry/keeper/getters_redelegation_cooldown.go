package keeper

import (
	"encoding/binary"
	"github.com/KYVENetwork/chain/x/registry/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRedelegationCooldown ...
func (k Keeper) SetRedelegationCooldown(ctx sdk.Context, redelegationCooldown types.RedelegationCooldown) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedelegationCooldownPrefix)
	store.Set(types.RedelegationCooldownKey(
		redelegationCooldown.Address,
		redelegationCooldown.CreatedBlock,
	), []byte{1})
}

// GetRedelegationCooldownEntries ...
func (k Keeper) GetRedelegationCooldownEntries(ctx sdk.Context, delegatorAddress string) (blocks []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedelegationCooldownPrefix)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixBuilder{}.AString(delegatorAddress).Key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		blocks = append(blocks, binary.BigEndian.Uint64(iterator.Key()))
	}
	return
}

// RemoveRedelegationCooldown ...
func (k Keeper) RemoveRedelegationCooldown(ctx sdk.Context, delegatorAddress string, block uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedelegationCooldownPrefix)
	store.Delete(types.RedelegationCooldownKey(delegatorAddress, block))
}

// GetAllRedelegationCooldownEntries ...
func (k Keeper) GetAllRedelegationCooldownEntries(ctx sdk.Context) (list []types.RedelegationCooldown) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RedelegationCooldownPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := types.RedelegationCooldown{
			Address:      string(iterator.Key()[0:43]),
			CreatedBlock: binary.BigEndian.Uint64(iterator.Key()[44:52]),
		}
		list = append(list, val)
	}

	return
}
