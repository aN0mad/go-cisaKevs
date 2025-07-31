package cisakev

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadCISAKEVs(t *testing.T) {
	var maxage = 0
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, CISAKEVFile)

	// Remove if already exists
	_ = os.Remove(filePath)

	c := NewCISAKEVs(nil)
	err := c.LoadCISAKEVs(tmpDir, true, time.Duration(maxage))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	kevs := c.GetKEVs()

	if len(kevs) == 0 {
		t.Fatalf("expected some KEVs, got 0")
	}

	// Check for required fields
	for _, k := range kevs {
		if k.CVEID == "" {
			t.Errorf("missing CVEID in KEV: %+v", k)
		}
	}
}
