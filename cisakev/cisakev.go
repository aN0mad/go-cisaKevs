package cisakev

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	CISAURL     = "https://www.cisa.gov/sites/default/files/csv/known_exploited_vulnerabilities.csv"
	CISAKEVFile = "known_exploited_vulnerabilities.csv"
)

func LoadCISAKEVs(dataDir string, forceRefresh bool, maxage time.Duration) ([]KEV, error) {
	filePath := filepath.Join(dataDir, CISAKEVFile)
	shouldDownload := forceRefresh

	if stat, err := os.Stat(filePath); err == nil {
		if time.Since(stat.ModTime()) > maxage {
			log.Println("CISA KEV file is older than 7 days. Downloading new one.")
			shouldDownload = true
		} else {
			log.Println("CISA KEV file is fresh. Using local copy.")
		}
	} else {
		log.Println("CISA KEV file not found. Downloading.")
		shouldDownload = true
	}

	if int(maxage) == 0 {
		shouldDownload = true
	}

	if shouldDownload {
		if err := downloadFile(CISAURL, filePath); err != nil {
			return nil, err
		}
		log.Println("CISA KEV file downloaded and saved to:", filePath)
	}

	return readCSV(filePath)
}

func downloadFile(url, filePath string) error {
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

func readCSV(path string) ([]KEV, error) {
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
			log.Println("Skipping malformed row:", row)
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
