package clnkserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/imflop/clnk/internal/app/store/sqlstore"
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

	store := sqlstore.New(db)
	srv := NewServer(store)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{config.OriginURL.Frontend})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions})

	fmt.Printf("server running at %s:%s\n", config.Server.Host, config.Server.Port)
	return http.ListenAndServe(
		config.Server.Host+":"+config.Server.Port,
		handlers.CORS(allowedMethods, allowedOrigins, allowedHeaders)(srv))
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
