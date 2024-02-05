package terraintypes

import (
	"fmt"
	"testing"

	terrainnames "github.com/awbw/2040/types/terrains/names"
	"github.com/stretchr/testify/assert"
)

func TestCreateTerrainMap(t *testing.T) {
	assert := assert.New(t)

	tileIDs := [][]int{
		{1, 2, 3},
		{26, 1, 1},
		{1, 1, 142},
	}

	mapTiles := TestMaps.FromTiles(tileIDs)
	terrainMap := CreateTerrainMap(mapTiles)

	assert.Equal(terrainMap[0][0].Name, terrainnames.Plain, fmt.Sprintf("(0, 0) must be %s", terrainnames.Plain))
	assert.Equal(terrainMap[0][1].Name, terrainnames.Mountain, fmt.Sprintf("(0, 1) must be %s", terrainnames.Mountain))

	HBridge := terrainMap[1][0]
	assert.Equal(HBridge.Name, terrainnames.Names[HBridge.ID], fmt.Sprintf("(1, 0) must be %s", terrainnames.Names[HBridge.ID]))

	GreenEarthLab := terrainMap[2][2]
	assert.Equal(GreenEarthLab.Name, terrainnames.Names[GreenEarthLab.ID], fmt.Sprintf("(2, 2) must be %s", terrainnames.Names[GreenEarthLab.ID]))
}
