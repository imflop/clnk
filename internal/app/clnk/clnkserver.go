package clnk

import (
	"net/http"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// Run server ...
func Run(config *Config) error {
	db, err := connectDB(config.Database.URL)
	if err != nil {
		return err
	}

	defer db.Close()
	srv := NewServer()
	return http.ListenAndServe(config.Server.Host+":"+config.Server.Port, srv)
}

func connectDB(databaseURL string) (*sqlx.DB, error) {
	d := &stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			PreferSimpleProtocol: true,
			RuntimeParams: map[string]string{
				"standard_conforming_strings": "on",
			},
		},
	}
	stdlib.RegisterDriverConfig(d)

	db, err := sqlx.Connect("pgx", d.ConnectionString(databaseURL))
	if err != nil {
		return nil, err
	}

	return db, nil
}
