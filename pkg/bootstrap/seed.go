package bootstrap

import databasemigrations "blog/internal/databaseMigrations"

func Seed() {
	Migrate()
	databasemigrations.Seed()
}
