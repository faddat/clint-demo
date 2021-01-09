package keeper

import (
	"github.com/faddat/clint-demo/x/clintdemo/types"
)

var _ types.QueryServer = Keeper{}
