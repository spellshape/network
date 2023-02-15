package keeper

import (
	"github.com/spellshape/network/x/campaign/types"
)

var _ types.QueryServer = Keeper{}
