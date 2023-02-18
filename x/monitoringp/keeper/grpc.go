package keeper

import (
	"github.com/spellshape/network/x/monitoringp/types"
)

var _ types.QueryServer = Keeper{}
