package keeper

import (
	"github.com/Finschia/finschia/x/settlement/types"
)

var _ types.QueryServer = Keeper{}
