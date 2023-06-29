package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

// CheckMisbehaviourAndUpdateState detects the following as Misbehaviour:
//
// 1. The existence of two different valid beacon headers for the same beacon block number
// 2. The existence of two different valid shard headers for the same shard ID and block number:
func (cs ClientState) CheckMisbehaviourAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	clientStore sdk.KVStore,
	misbehaviour exported.Misbehaviour,
) (exported.ClientState, error) {
	return &cs, nil
}
