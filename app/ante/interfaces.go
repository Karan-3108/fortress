package ante

import (
	"math/big"

	evmtypes "github.com/Karan-3108/ethermint/x/evm/types"
	"github.com/Karan-3108/fortress/v4/x/fees/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/params"
)

// FeesKeeper defines the expected keeper interface used on the AnteHandler
type FeesKeeper interface {
	GetParams(ctx sdk.Context) (params types.Params)
}

// EvmKeeper defines the expected keeper interface used on the AnteHandler
type EvmKeeper interface {
	GetParams(ctx sdk.Context) (params evmtypes.Params)
	ChainID() *big.Int
	GetBaseFee(ctx sdk.Context, ethCfg *params.ChainConfig) *big.Int
}
