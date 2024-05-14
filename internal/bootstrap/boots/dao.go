// Code generated by go-skeleton.
// Code generated by go-skeleton.
// Code generated by go-skeleton.

package boots

import (
	"sync"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/vars"
)

var once sync.Once

// InitDao ...
func InitDao() {
	once.Do(func() {
		// dao default set
		if vars.MDB["default"] != nil {
			dao.SetDefault(vars.MDB["default"])
		}
	})
}
