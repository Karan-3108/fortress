package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fortress module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
var (
	ErrWrongFortressState = sdkerrors.Register(ModuleName, 1, "wrong fortress state")
	ErrDeadline           = sdkerrors.Register(ModuleName, 2, "deadline")
)
