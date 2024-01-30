package terraintypes

import (
	"fmt"
	"testing"

	countrynames "github.com/awbw/2040/types/countries/names"
)

func TestFormatName(t *testing.T) {
	name := FormatBuilding(countrynames.OrangeStar, Base)
	want := fmt.Sprintf("%s %s", countrynames.OrangeStar, Base)
	if name != want {
		t.Fatalf("Failed to format name. Got %s, want %s", name, want)
	}
}

func TestCreateTerrainNames(t *testing.T) {
	terrains := NewTerrainNames()
	want := string(Plain)

	if terrains[1] != want {
		t.Fatalf(fmt.Sprintf("Terrain names initialization error. Got: %s. Want: %s", terrains[1], want))
	}

	roaddirs := Roaddirs()
	withdirs := TerrainNames{
		4:  string(River),
		15: string(Road),
	}

	for terrainId, terrainName := range withdirs {
		want = terrainName
		for i, dir := range roaddirs {
			if terrains[terrainId+i] != fmt.Sprintf("%s%s", dir, want) {
				t.Logf(fmt.Sprintf("River & Road names initialization error. Got: %s (id %d). Want: %s", terrains[terrainId+i], terrainId+i, want))
			}
		}
	}
}
