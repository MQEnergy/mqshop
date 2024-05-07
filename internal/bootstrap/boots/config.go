package boots

import (
	"log/slog"

	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/config"
)

// InitConfig ...
func InitConfig() error {
	var err error
	vars.Config, err = config.New(config.NewViper(), config.Options{
		BasePath: vars.BasePath,
		FileName: "config." + config.ConfEnv,
	})
	if err == nil {
		slog.Info("Server.mode: " + vars.Config.GetString("server.mode"))
		slog.Info("Loading Configuration successfully")
	}
	return err
}
