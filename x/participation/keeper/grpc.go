package keeper

import (
	"github.com/spellshape/network/x/participation/types"
)

var _ types.QueryServer = Keeper{}
