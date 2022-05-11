package migrations

import (
	"embed"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/merdan2000/internal/settings"
	"github.com/pressly/goose"
)

var embedMigrations embed.FS

func MigrationUp(settings *settings.Settings) error {
	fmt.Println("Migration Started")
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		settings.Host, settings.Port, settings.DbName, settings.Password,
		settings.User, settings.SSLmode)

	db, err := goose.OpenDBWithDriver("pgx", sqlInfo)
	if err != nil {
		return err
	}

	defer db.Close()

	dir := "./migrations/"

	if err := goose.Up(db, dir); err != nil {
		return err
	}

	fmt.Println("Migration End")
	return nil
}
