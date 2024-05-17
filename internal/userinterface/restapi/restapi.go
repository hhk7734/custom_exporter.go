package restapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	PORT_KEY = "port"
)

type RestAPIConfig struct {
	Port string
}

func RestAPIPFlags() *pflag.FlagSet {
	f := pflag.NewFlagSet("restapi", pflag.ContinueOnError)
	f.String(PORT_KEY, "9107", "restapi port")
	return f
}

func RestAPIConfigFromViper() RestAPIConfig {
	return RestAPIConfig{
		Port: viper.GetString(PORT_KEY),
	}
}

type RestAPI struct {
	server *http.Server
}

func NewRestAPI(cfg RestAPIConfig) *RestAPI {
	zap.L().Info("restapi config", zap.Dict("config",
		zap.String(PORT_KEY, cfg.Port),
	))
	return &RestAPI{
		server: &http.Server{
			Addr: fmt.Sprintf(":%s", cfg.Port),
		},
	}
}

func (r *RestAPI) Run() error {
	return r.server.ListenAndServe()
}

func (r *RestAPI) Shutdown() error {
	return r.server.Shutdown(context.Background())
}
