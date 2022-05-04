package keeper

import (
	"encoding/binary"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetFortressCount get the total number of fortress
func (k Keeper) GetFortressCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FortressCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFortressCount set the total number of fortress
func (k Keeper) SetFortressCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.FortressCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFortress appends a fortress in the store with a new id and update the count
func (k Keeper) AppendFortress(
	ctx sdk.Context,
	fortress types.Fortress,
) uint64 {
	// Create the fortress
	count := k.GetFortressCount(ctx)

	// Set the ID of the appended value
	fortress.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortressKey))
	appendedValue := k.cdc.MustMarshal(&fortress)
	store.Set(GetFortressIDBytes(fortress.Id), appendedValue)

	// Update fortress count
	k.SetFortressCount(ctx, count+1)

	return count
}

// SetFortress set a specific fortress in the store
func (k Keeper) SetFortress(ctx sdk.Context, fortress types.Fortress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortressKey))
	b := k.cdc.MustMarshal(&fortress)
	store.Set(GetFortressIDBytes(fortress.Id), b)
}

// GetFortress returns a fortress from its id
func (k Keeper) GetFortress(ctx sdk.Context, id uint64) (val types.Fortress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortressKey))
	b := store.Get(GetFortressIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFortress removes a fortress from the store
func (k Keeper) RemoveFortress(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortressKey))
	store.Delete(GetFortressIDBytes(id))
}

// GetAllFortress returns all fortress
func (k Keeper) GetAllFortress(ctx sdk.Context) (list []types.Fortress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortressKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fortress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFortressIDBytes returns the byte representation of the ID
func GetFortressIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetFortressIDFromBytes returns ID in uint64 format from a byte array
func GetFortressIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
