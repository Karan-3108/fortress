package keeper

import (
	"context"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FortressAll(c context.Context, req *types.QueryAllFortressRequest) (*types.QueryAllFortressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fortresss []types.Fortress
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	fortressStore := prefix.NewStore(store, types.KeyPrefix(types.FortressKey))

	pageRes, err := query.Paginate(fortressStore, req.Pagination, func(key []byte, value []byte) error {
		var fortress types.Fortress
		if err := k.cdc.Unmarshal(value, &fortress); err != nil {
			return err
		}

		fortresss = append(fortresss, fortress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFortressResponse{Fortress: fortresss, Pagination: pageRes}, nil
}

func (k Keeper) Fortress(c context.Context, req *types.QueryGetFortressRequest) (*types.QueryGetFortressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	fortress, found := k.GetFortress(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFortressResponse{Fortress: fortress}, nil
}
