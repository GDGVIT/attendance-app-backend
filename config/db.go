package config

import (
	"github.com/spf13/viper"
)

func DbConfiguration() string {
	// check viper bool DEBUG, if true then PLANETSCALE_DB_DSN_DEV else PLANETSCALE_DB_DSN
	planetscaleDBDSN := viper.GetString("PLANETSCALE_DB_DSN")
	if viper.GetBool("DEBUG") {
		planetscaleDBDSN = viper.GetString("PLANETSCALE_DB_DSN_DEV")
	}

	return planetscaleDBDSN
}
