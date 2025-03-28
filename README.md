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
kevs, err := cisa.LoadCISAKEVs("data", false)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Loaded %d KEVs\n", len(kevs))
```

## Testing

```bash
go test ./internal/cisa
```