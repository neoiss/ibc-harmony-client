package types

import (
	"strings"
	time "time"

	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

func NewClientState() *ClientState {
	return &ClientState{}
}

func (cs ClientState) ClientType() string {
	return MAPO
}

func (cs ClientState) GetLatestHeight() exported.Height {
	return clienttypes.NewHeight(revisionNumber, cs.LatestHeight)
}

func (cs ClientState) Validate() error {
	if strings.TrimSpace(cs.ClientIdentifier) == "" {
		return sdkerrors.Wrap(ErrInvalidClientIdentifier, "client identifier cannot be empty string")
	}
	if cs.LatestEpoch == 0 {
		return sdkerrors.Wrap(ErrInvalidLatestEpoch, "latest epoch must be greater than zero")
	}
	if cs.EpochSize == 0 {
		return sdkerrors.Wrap(ErrInvalidEpochSize, "epoch size must be greater than zero")
	}
	if cs.LatestHeight == 0 {
		return sdkerrors.Wrap(ErrInvalidLatestHeight, "latest height must be greater than zero")
	}
	return nil
}

func (cs ClientState) GetProofSpecs() []*ics23.ProofSpec {
	return nil
}

// Initialization function
// Clients must validate the initial consensus state, and may store any client-specific metadata
// necessary for correct light client operation
func (cs ClientState) Initialize(ctx sdk.Context, cdc codec.BinaryCodec, store sdk.KVStore, consState exported.ConsensusState) error {
	if _, ok := consState.(*ConsensusState); !ok {
		return sdkerrors.Wrapf(clienttypes.ErrInvalidConsensus, "invalid initial consensus state. expected type: %T, got: %T",
			&ConsensusState{}, consState)
	}
	// TODO set metadata for initial consensus state
	return nil
}

// Status function
// Clients must return their status. Only Active clients are allowed to process packets.
func (cs ClientState) Status(ctx sdk.Context, clientStore sdk.KVStore, cdc codec.BinaryCodec) exported.Status {
	if cs.Frozen {
		return exported.Frozen
	}

	return exported.Active
}

// Genesis function
func (cs ClientState) ExportMetadata(_ sdk.KVStore) []exported.GenesisMetadata {
	return nil
}

// CheckHeaderAndUpdateState verifies that:
// - the beacon header with the associated committee signature and bitmap. It also verifies the quorum using the committee of the target epoch.
// - (for shard 1+) the shard header with the cross-link of the beacon header.
// If the target header's epoch is older than the epoch of ClientState,
// `header` must have "epoch header(s)", which is/are the last beacon header(s) of each epoch for updating epoch of ClientState.
func (cs ClientState) CheckHeaderAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, clientStore sdk.KVStore,
	header exported.Header,
) (exported.ClientState, exported.ConsensusState, error) {
	h, ok := header.(*Header)
	if !ok {
		return nil, nil, sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader, "expected type %T, got %T", &Header{}, header,
		)
	}

	consensusState := &ConsensusState{
		Epoch: cs.LatestEpoch,
		Validators: &ValidatorSet{
			PairKeys: []*PairKey{
				{
					Address:  "0xff",
					G2PubKey: []byte{1},
				},
			},
			Weights: []uint64{1},
		},
		CommitmentRoot: h.CommitmentRoot,
		Timestamp:      time.Now(),
	}
	return &cs, consensusState, nil
}

// Future work is needed
func (cs ClientState) CheckSubstituteAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore sdk.KVStore, substituteClientStore sdk.KVStore, substituteClient exported.ClientState) (exported.ClientState, error) {
	return nil, sdkerrors.Wrapf(
		clienttypes.ErrUpdateClientFailed,
		"mapo client is not allowed to updated with a proposal",
	)
}

// Future work is needed
func (cs ClientState) VerifyUpgradeAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, store sdk.KVStore, newClient exported.ClientState, newConsState exported.ConsensusState, proofUpgradeClient []byte, proofUpgradeConsState []byte) (exported.ClientState, exported.ConsensusState, error) {
	return nil, nil, sdkerrors.Wrap(clienttypes.ErrInvalidUpgradeClient, "cannot upgrade mapo client")
}

// Utility function that zeroes out any client customizable fields in client state
// Ledger enforced fields are maintained while all custom fields are zero values
// Used to verify upgrades
func (cs ClientState) ZeroCustomFields() exported.ClientState {
	return &ClientState{
		Frozen:       cs.Frozen,
		LatestEpoch:  cs.LatestEpoch,
		EpochSize:    cs.EpochSize,
		LatestHeight: cs.LatestHeight,
	}
}

// State verification functions
func (cs ClientState) VerifyClientState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, counterpartyClientIdentifier string, proof []byte, clientState exported.ClientState) error {
	return nil
}

func (cs ClientState) VerifyClientConsensusState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, counterpartyClientIdentifier string, consensusHeight exported.Height, prefix exported.Prefix, proof []byte, consensusState exported.ConsensusState) error {
	return nil
}

func (cs ClientState) VerifyConnectionState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, connectionID string, connectionEnd exported.ConnectionI) error {
	return nil
}

func (cs ClientState) VerifyChannelState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, portID string, channelID string, channel exported.ChannelI) error {
	return nil
}

func (cs ClientState) VerifyPacketCommitment(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64, commitmentBytes []byte) error {
	return nil
}

func (cs ClientState) VerifyPacketAcknowledgement(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64, acknowledgement []byte) error {
	return nil
}

func (cs ClientState) VerifyPacketReceiptAbsence(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64) error {
	return nil
}

func (cs ClientState) VerifyNextSequenceRecv(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, nextSequenceRecv uint64) error {
	return nil
}
