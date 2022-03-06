package main

import (
	migrate "web-server/database/migrate"
)

// var reset bool

// migrateCmd represents the migrate command
func main() {

	// if reset {
	// 	migrate.Reset()
	// }

	migrate.Migrate(nil)

}
