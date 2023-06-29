package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MAPO client sentinel errors
var (
	ErrInvalidClientIdentifier = sdkerrors.Register(ModuleName, 1, "invalid client identifier")
	ErrInvalidLatestEpoch      = sdkerrors.Register(ModuleName, 2, "invalid latest epoch")
	ErrInvalidEpochSize        = sdkerrors.Register(ModuleName, 3, "invalid epoch size")
	ErrInvalidLatestHeight     = sdkerrors.Register(ModuleName, 4, "invalid latest height")
)
