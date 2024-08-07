package entity

import (
	"github.com/MQEnergy/mqshop/internal/app/entity/admin"
	"github.com/MQEnergy/mqshop/pkg/helper"
)

type MethodMaps map[string][]any

// methodMaps 需手动添加配置
var methodMaps = MethodMaps{
	"ms_c_admin": {
		func(Querier) {},
		func(admin.Querier) {},
	},
}

// Load ...
func Load(models []string) (tableMethods MethodMaps) {
	tableMethods = make(MethodMaps)
	if len(models) > 0 {
		for table, methods := range methodMaps {
			if helper.InAnySlice(models, table) {
				tableMethods[table] = methods
			} else {
				tableMethods[table] = []any{}
			}
		}
	} else {
		tableMethods = methodMaps
	}
	return
}
