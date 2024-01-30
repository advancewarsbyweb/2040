package models

import (
	"fmt"
	"testing"

	terraintypes "github.com/awbw/2040/types/terrains"
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

	assert.Equal(terrainMap[0][0].Name, terraintypes.Plain, fmt.Sprintf("(0, 0) must be %s", terraintypes.Plain))
	assert.Equal(terrainMap[0][1].Name, terraintypes.Mountain, fmt.Sprintf("(0, 1) must be %s", terraintypes.Mountain))

	HBridge := terrainMap[1][0]
	assert.Equal(HBridge.Name, terraintypes.Names[HBridge.ID], fmt.Sprintf("(1, 0) must be %s", terraintypes.Names[HBridge.ID]))

	GreenEarthLab := terrainMap[2][2]
	assert.Equal(GreenEarthLab.Name, terraintypes.Names[GreenEarthLab.ID], fmt.Sprintf("(2, 2) must be %s", terraintypes.Names[GreenEarthLab.ID]))
}
