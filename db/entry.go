package db

import (
	"time"
)

type Entry struct {
	ID uint64 `json:"id"`
	Entity string `json:"entity"`
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lng"`

	Timestamp time.Time `json:"ts"`
	IngestTime time.Time `json:"its"`
}

// Add adds a new Entry to the database
//
// e Entry - Entry to add
func (c *DB) Add(e Entry) error {
	tx, err := c.conn.Begin()
	
	if err != nil {
		return err
	}
	
	query, err := tx.Prepare(`
     INSERT INTO "locations" ("entity", "latitude", "longitude", "timestamp", "ingest_timestamp")
     VALUES ($1, $2, $3, $4, $5)`)
	
	if err != nil {
		return err
	}
	
	if _, err := query.Exec(e.Entity, e.Latitude, e.Longitude, e.Timestamp, e.IngestTime); err != nil {
		return err
	}
	
	if err := query.Close(); err != nil {
		return err
	}
	
	err = tx.Commit()
	
	return err
}

func (c *DB) AllEntries() ([]Entry, error) {
	
	rows, err := c.conn.Query(`SELECT "id", "entity", "timestamp", "ingest_timestamp", "latitude", "longitude" FROM "locations" ORDER BY "ingest_timestamp" DESC`)
	
	if err != nil {
		return nil, err
	}
	
	var entries []Entry
	
	for rows.Next() {
		var e Entry
		
		err := rows.Scan(&e.ID, &e.Entity, &e.Timestamp, &e.IngestTime, &e.Latitude, &e.Longitude)
		
		if err != nil {
			return nil, err
		}
		
		entries = append(entries, e)
	}
	
	return entries, err
}