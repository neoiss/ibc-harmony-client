package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

var _ exported.Header = (*Header)(nil)

const revisionNumber = 0

func (h Header) ClientType() string {
	return MAPO
}

// GetHeight returns the target height
func (h Header) GetHeight() exported.Height {
	return clienttypes.NewHeight(0, h.Number)
}

// ValidateBasic calls
func (h Header) ValidateBasic() error {
	if h.SignedHeader == nil {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "map signed header cannot be nil")
	}
	if len(h.CommitmentRoot) == 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidHeader, "map commitment root cannot be empty")
	}
	return nil
}
