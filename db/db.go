package db

import (
	"github.com/spf13/viper"
	"database/sql"
	
	_ "github.com/lib/pq"
	"gpstrace/db/migrations"
	"github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
)

func init() {

	viper.SetDefault("postgres.url", "postgresql://gpstrace:gpstrace@127.0.0.1/gpstrace?sslmode=disable")
	viper.BindEnv("postgres.url", "GPSTRACE_PG_URL", "PG_HOST", "PGURL")
	
}

type DB struct {
	conn *sql.DB
}

func GetConnection() (*DB, error) {
	conn, err := sql.Open("postgres", viper.GetString("postgres.url"))
	
	if err != nil {
		return nil, err
	}
	
	d := &DB{conn}
	
	if err := d.Migrate(); err != nil {
	
	}
	
	return d, nil
}


func (d *DB) Migrate() error {

	src := bindata.Resource(migrations.AssetNames(), func(name string) ([]byte, error) {
		return migrations.Asset(name)
	})
	
	dc, err := postgres.WithInstance(d.conn, &postgres.Config{})
	
	if err != nil {
		return err
	}
	
	dr, err := bindata.WithInstance(src)
	
	if err != nil {
		return err
	}
	
	m, err := migrate.NewWithInstance("go-bindata", dr,  "postgres", dc)
	
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}
	
	return nil
}
