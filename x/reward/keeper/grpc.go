package keeper

import (
	"github.com/spellshape/network/x/reward/types"
)

var _ types.QueryServer = Keeper{}
