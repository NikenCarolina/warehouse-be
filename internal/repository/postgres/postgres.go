package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/NikenCarolina/warehouse-be/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Init(config *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
