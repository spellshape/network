package keeper

import (
	"github.com/spellshape/network/x/monitoringc/types"
)

var _ types.QueryServer = Keeper{}
