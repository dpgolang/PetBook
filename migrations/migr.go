package migrations

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	
)

func GetMigrations() *migrate.FileMigrationSource {
	return &migrate.FileMigrationSource{
		Dir: "migrations",
	}
}

func MigrateUp(db *sql.DB,mig *migrate.FileMigrationSource){
n, err := migrate.Exec(db, "postgres", mig, migrate.Up)
if err != nil {
log.Printf("can't work migrations:%v", err)
}
fmt.Printf("Applied %d migrations!\n", n)
err = db.Ping()
if err != nil {
log.Printf("can't ping:%v", err)
}
}