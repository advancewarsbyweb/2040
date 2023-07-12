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
