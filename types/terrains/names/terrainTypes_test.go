package terrainnames

import "testing"

func TestTerrainMatch(t *testing.T) {
	terrainName := TerrainName("Green Earth HQ")
	if !HQ.Match(terrainName) {
		t.Errorf("Given terrain name '%s' did not match 'HQ'", terrainName)
	}
}
