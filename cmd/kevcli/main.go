package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aN0mad/go-cisaKevs/cisakev"
)

var (
	MaxAge = 7 * 24 * time.Hour // 7 days
)

func main() {
	var force bool
	flag.BoolVar(&force, "refresh", false, "Force refresh the CISA KEV file")
	flag.Parse()

	// Verify data dir exists
	dataDir := "data"
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			log.Fatal("Failed to create data directory:", err)
		}
	}

	// Create CISA KEVs instance with default logger
	cisaKevs := cisakev.NewCISAKEVs(nil) // Use default logger
	if cisaKevs == nil {
		log.Fatal("Failed to create CISA KEVs instance")
	}

	// Load KEVs from local file or download if necessary
	err := cisaKevs.LoadCISAKEVs(dataDir, force, MaxAge)
	if err != nil {
		log.Fatal("Error loading KEVs:", err)
	}

	// Get loaded KEVs
	kevs := cisaKevs.GetKEVs()
	fmt.Printf("✅ Loaded %d Known Exploited Vulnerabilities\n", len(kevs))
}
