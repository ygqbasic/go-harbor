package harbor

import (
	"testing"
)

func TestHarborStatistics(t *testing.T) {
	harborClient := NewClient(nil, "https://reg.testcloud.com", "admin", "Harbor12345")
	statisticMap, _, errs := harborClient.GetStatistics()
	if errs == nil {
		t.Log(statisticMap)
	}
}
