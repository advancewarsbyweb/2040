package terraintypes

import "testing"

func TestTerrainMatch(t *testing.T) {
	terrainName := "Green Earth HQ"
	if !HQ.Match(terrainName) {
		t.Errorf("Given terrain name '%s' did not match 'HQ'", terrainName)
	}
}
