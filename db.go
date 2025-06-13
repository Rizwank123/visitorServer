package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	//dsn := "postgresql://postgres:root@localhost:5433/visitor_db?sslmode=disable"
	dsn := "postgresql://visitor_db_bc1r_user:jdFrkJCkR21kPDh4YzBRqmxbcuDHkYMu@dpg-d15s8andiees73een8u0-a/visitor_db_bc1r"
	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}

	if err := DB.Ping(ctx); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	schema := `CREATE TABLE IF NOT EXISTS visitors (
		id SERIAL PRIMARY KEY,
		ip TEXT,
		network TEXT,
		version TEXT,
		city TEXT,
		region TEXT,
		region_code TEXT,
		country TEXT,
		country_name TEXT,
		country_code TEXT,
		country_code_iso3 TEXT,
		country_capital TEXT,
		country_tld TEXT,
		continent_code TEXT,
		in_eu BOOLEAN,
		postal TEXT,
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		timezone TEXT,
		utc_offset TEXT,
		country_calling_code TEXT,
		currency TEXT,
		currency_name TEXT,
		languages TEXT,
		country_area INTEGER,
		country_population BIGINT,
		asn TEXT,
		org TEXT,
		visited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(ctx, schema)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}
