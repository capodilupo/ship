package config

import (
	"context"

	"github.com/replicatedcom/ship/pkg/api"
	"github.com/replicatedcom/ship/pkg/fs"
	"github.com/replicatedcom/ship/pkg/logger"
	"github.com/replicatedcom/ship/pkg/ui"
	"github.com/spf13/viper"
)

// Resolver is a thing that can resolve configuration options
type Resolver interface {
	ResolveConfig(context.Context, *api.Release, map[string]interface{}) (map[string]interface{}, error)
}

func ResolverFromViper(v *viper.Viper) Resolver {
	if v.GetBool("headless") {
		return &CLIResolver{
			Logger: logger.FromViper(v),
			UI:     ui.FromViper(v),
			Viper:  v,
		}
	}

	daemon := DaemonFromViper(v)

	return &DaemonResolver{
		Logger:             logger.FromViper(v),
		MaybeRunningDaemon: daemon,
	}

}
func DaemonFromViper(v *viper.Viper) *Daemon {
	return &Daemon{
		Logger:           logger.FromViper(v),
		Fs:               fs.FromViper(v),
		UI:               ui.FromViper(v),
		Viper:            v,
		ConfigSaved:      make(chan interface{}),
		MessageConfirmed: make(chan string, 1),
	}
}