package boots

import (
	"github.com/MQEnergy/mqshop/internal/vars"
	logger2 "github.com/MQEnergy/mqshop/pkg/logger"
)

// InitLogger ...
func InitLogger() {
	logger2.New(
		vars.Config.GetString("log.fileName"),
		vars.Config,
	)
}
