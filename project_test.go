package harbor

import (
	"testing"
)

func TestHarborProject(t *testing.T) {
	harborClient := NewClient(nil, "https://reg.testcloud.com", "admin", "Harbor12345")
	opt := &ListProjectsOptions{Name: "test"}

	projects, _, err := harborClient.Projects.ListProject(opt)
	if err == nil {
		t.Log(projects)
	}
}
