package migration

import (
	"quiz3/config"
)

func RunMigration() {
   
    config.ConnectDatabase()

    config.RunMigration()
}
