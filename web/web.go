package web

import (
	"github.com/gin-gonic/gin"
	"gpstrace/util"
	"net/http"
	"errors"
	"gpstrace/db"
	"time"
	"github.com/spf13/viper"
	"log"
	"encoding/json"
)

func init() {
	viper.SetDefault("http.listen", ":8080")
}

// New creates a new app HTTP handler
func New(d *db.DB) http.Handler {
	h := gin.New()
	
	h.Use(func (c *gin.Context) {
		c.Set("db", d)
		c.Next()
	})
	
	h.POST("/data", dataIngestHandler)
	h.GET("/data", allDataHandler)
	
	return h
}

func dataIngestHandler(c *gin.Context) {
	fLat := c.PostForm("lat")
	fLng := c.PostForm("lng")
	fEnt := c.PostForm("ent")
	
	lat, err := util.ConvertDmsToDecimal(fLat)
	
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	lng, err := util.ConvertDmsToDecimal(fLng)
	
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	if fEnt == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("entity is required"))
	}
	
	ts := time.Now()
	
	entry := db.Entry{
		Latitude:   *lat,
		Longitude:  *lng,
		Entity:     fEnt,
		IngestTime: ts,
		Timestamp:  ts,
	}
	
	 d := c.MustGet("db").(*db.DB)
	
	if err := d.Add(entry); err != nil {
		log.Println("Error inserting entry:", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	log.Printf("lat=%f lng=%f ent=%s ts=%s", entry.Latitude, entry.Longitude, entry.Entity, entry.IngestTime)
	
	c.Status(http.StatusCreated)
	return
}


func allDataHandler(c *gin.Context) {
	d := c.MustGet("db").(*db.DB)
	
	c.Header("Content-Type", "application/json")
	
	entries, err := d.AllEntries()
	
	if err != nil {
		log.Println("Error getting entries:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to get entries"})
		return
	}

	c.Status(http.StatusOK)
	e := json.NewEncoder(c.Writer)
	e.SetIndent("", "    ")
	e.Encode(entries)
}
