package config

import (
	"log"

	migrate "github.com/rubenv/sql-migrate"
)


func RunMigration() {
	
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations", 
	}

	
	db := DB 
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Printf("Applied %d migrations!\n", n)
}


func RollbackMigration() {
	
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	
	db := DB 

	
	n, err := migrate.ExecMax(db, "postgres", migrations, migrate.Down, 1)
	if err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}

	log.Printf("Rolled back %d migrations!\n", n)
}
