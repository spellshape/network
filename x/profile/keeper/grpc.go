package keeper

import (
	"github.com/spellshape/network/x/profile/types"
)

var _ types.QueryServer = Keeper{}
