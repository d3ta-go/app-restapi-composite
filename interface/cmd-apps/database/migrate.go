package database

import (
	"fmt"

	migDBAccount "github.com/d3ta-go/ddd-mod-account/modules/account/infrastructure/migration"
	migDBCovid19 "github.com/d3ta-go/ddd-mod-covid19/modules/covid19/infrastructure/migration"
	migDBEmail "github.com/d3ta-go/ddd-mod-email/modules/email/infrastructure/migration"
	migDBGeoLoc "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/infrastructure/migration"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/fsnotify/fsnotify"
)

func initConfig(h *handler.Handler) (*config.Config, error) {
	//init config
	cfg, viper, err := config.NewConfig("./")
	if err != nil {
		panic(err)
	}
	h.SetDefaultConfig(cfg)
	h.SetViper("config", viper)

	viper.OnConfigChange(func(e fsnotify.Event) {
		// fmt.Println("config file changed:", e.Name)
		c := new(config.Config)
		if err := viper.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}

		h.SetDefaultConfig(c)
		initializeSystems(h)
	})

	return cfg, nil
}

func initializeSystems(h *handler.Handler) error {

	// initialize database
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		panic(err)
	}

	// initialize cacher
	if err := initialize.OpenAllCacheConnection(h); err != nil {
		panic(err)
	}

	return nil
}

// RunDBMigrate run DB Migration
func RunDBMigrate() error {
	// Add your custom code
	// init super handler
	superHandler := new(handler.Handler)

	// init configuration
	_, err := initConfig(superHandler)
	if err != nil {
		return err
	}

	// initialize Systems
	err = initializeSystems(superHandler)
	if err != nil {
		return err
	}
	defer initialize.CloseDBConnections(superHandler)

	// account
	if err := _runAccountDBMigrate(superHandler); err != nil {
		return err
	}

	// email
	if err := _runEmailDBMigrate(superHandler); err != nil {
		return err
	}

	// geo-location
	if err := _runGeoLocDBMigrate(superHandler); err != nil {
		return err
	}

	// covid19
	if err := _runCovid19DBMigrate(superHandler); err != nil {
		return err
	}

	return nil
}

func _runAccountDBMigrate(h *handler.Handler) error {
	mig, err := migDBAccount.NewRDBMSMigration(h)
	if err != nil {
		return err
	}

	if err := mig.Run(); err != nil {
		if err := mig.RollBack(); err != nil {
			return err
		}
	}
	return nil
}

func _runEmailDBMigrate(h *handler.Handler) error {
	mig, err := migDBEmail.NewRDBMSMigration(h)
	if err != nil {
		return err
	}

	if err := mig.Run(); err != nil {
		if err := mig.RollBack(); err != nil {
			return err
		}
	}
	return nil
}

func _runGeoLocDBMigrate(h *handler.Handler) error {
	mig, err := migDBGeoLoc.NewRDBMSMigration(h)
	if err != nil {
		return err
	}

	if err := mig.Run(); err != nil {
		if err := mig.RollBack(); err != nil {
			return err
		}
	}
	return nil
}

func _runCovid19DBMigrate(h *handler.Handler) error {
	mig, err := migDBCovid19.NewRDBMSMigration(h)
	if err != nil {
		return err
	}

	if err := mig.Run(); err != nil {
		if err := mig.RollBack(); err != nil {
			return err
		}
	}
	return nil
}
