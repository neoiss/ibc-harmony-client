package types

import (
	"github.com/cosmos/ibc-go/modules/core/exported"
)

func (cs ConsensusState) ClientType() string {
	return MAPO
}

// GetRoot returns the commitment root of the consensus state,
// which is used for key-value pair verification.
func (cs ConsensusState) GetRoot() exported.Root {
	return root{value: cs.CommitmentRoot}
}

// GetTimestamp returns the timestamp (in nanoseconds) of the consensus state
func (cs ConsensusState) GetTimestamp() uint64 {
	return uint64(cs.Timestamp.UnixNano())
}

func (cs ConsensusState) ValidateBasic() error {
	return nil
}

type root struct {
	value []byte
}

var _ exported.Root = (*root)(nil)

func (r root) GetHash() []byte {
	return r.value
}

func (r root) Empty() bool {
	return len(r.value) == 0
}
