# go-cisaKevs

A lightweight Go module to download and parse the [CISA Known Exploited Vulnerabilities (KEV)](https://www.cisa.gov/known-exploited-vulnerabilities-catalog) CSV file.

## Features

- Automatically downloads the KEV CSV file from CISA.
- Caches the file locally and refreshes it if older than 7 days.
- Provides a structured `KEV` object for programmatic use.
- Simple CLI for triggering downloads and inspecting KEV data.

## Installation

```bash
go get github.com/aN0mad/go-cisaKevs
```

## Usage (CLI)

```bash
go run cmd/kevcli/main.go --refresh
```

## Usage (Library)

```go
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

	// Load CISA KEVs
	cisaKevs := cisakev.NewCISAKEVs(nil) // Use default logger
	if cisaKevs == nil {
		log.Fatal("Failed to create CISA KEVs instance")
	}

	err := cisaKevs.LoadCISAKEVs(dataDir, force, MaxAge)
	if err != nil {
		log.Fatal("Error loading KEVs:", err)
	}

	kevs := cisaKevs.GetKEVs()
	fmt.Printf("✅ Loaded %d Known Exploited Vulnerabilities\n", len(kevs))
}

```

## Testing

```bash
go test ./internal/cisa
```