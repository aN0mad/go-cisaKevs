package cisakev

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	CISAURL     = "https://www.cisa.gov/sites/default/files/csv/known_exploited_vulnerabilities.csv"
	CISAKEVFile = "known_exploited_vulnerabilities.csv"
)

type CISAKEVS struct {
	KEVs   []KEV `json:"kevs"`
	logger Logger
}

type KEV struct {
	CVEID                      string
	VendorProject              string
	Product                    string
	VulnerabilityName          string
	DateAdded                  string
	ShortDescription           string
	RequiredAction             string
	DueDate                    string
	KnownRansomwareCampaignUse string
	Notes                      string
	CWEs                       string
}

// NewCISAKEVs creates a new CISAKEVS instance with optional logger
func NewCISAKEVs(logger Logger) *CISAKEVS {
	// Use default logger if none provided
	if logger == nil {
		logger = &defaultLogger{}
	}

	return &CISAKEVS{
		KEVs:   []KEV{},
		logger: logger,
	}
}

// LoadCISAKEVs loads KEVs from local file or downloads from CISA
func (c *CISAKEVS) LoadCISAKEVs(dataDir string, forceRefresh bool, maxage time.Duration) error {
	filePath := filepath.Join(dataDir, CISAKEVFile)
	shouldDownload := forceRefresh

	if stat, err := os.Stat(filePath); err == nil {
		if time.Since(stat.ModTime()) > maxage {
			c.logger.Info("CISA KEV file is older than threshold. Downloading new one.")
			shouldDownload = true
		} else {
			c.logger.Info("CISA KEV file is fresh. Using local copy.")
		}
	} else {
		c.logger.Info("CISA KEV file not found. Downloading.")
		shouldDownload = true
	}

	if int(maxage) == 0 {
		shouldDownload = true
	}

	if shouldDownload {
		if err := c.downloadFile(CISAURL, filePath); err != nil {
			return err
		}
		c.logger.Info("CISA KEV file downloaded and saved to: " + filePath)
	}

	kevs, err := c.readCSV(filePath)
	if err != nil {
		return err
	}

	c.KEVs = kevs
	return nil
}

func (c *CISAKEVS) downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to download CISA KEV file")
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func (c *CISAKEVS) readCSV(path string) ([]KEV, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var kevs []KEV
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 11 {
			c.logger.Warn("Skipping malformed row")
			continue
		}
		kev := KEV{
			CVEID:                      row[0],
			VendorProject:              row[1],
			Product:                    row[2],
			VulnerabilityName:          row[3],
			DateAdded:                  row[4],
			ShortDescription:           row[5],
			RequiredAction:             row[6],
			DueDate:                    row[7],
			KnownRansomwareCampaignUse: row[8],
			Notes:                      row[9],
			CWEs:                       row[10],
		}
		kevs = append(kevs, kev)
	}

	return kevs, nil
}

// GetKEVs returns the loaded KEVs
func (c *CISAKEVS) GetKEVs() []KEV {
	return c.KEVs
}

// // For backward compatibility with the package-level function
// func LoadCISAKEVs(dataDir string, forceRefresh bool, maxage time.Duration) ([]KEV, error) {
// 	c := New(nil)
// 	err := c.LoadCISAKEVs(dataDir, forceRefresh, maxage)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return c.GetKEVs(), nil
// }
