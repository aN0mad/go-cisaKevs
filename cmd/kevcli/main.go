package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aN0mad/go-cisaKevs/internal/cisa"
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

	// Load CISA KEVs
	kevs, err := cisa.LoadCISAKEVs(dataDir, force, MaxAge)
	if err != nil {
		log.Fatal("Error loading KEVs:", err)
	}

	fmt.Printf("✅ Loaded %d Known Exploited Vulnerabilities\n", len(kevs))
}
