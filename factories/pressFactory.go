package factories

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
)

type PressFactory struct{}

var Press PressFactory

func NewPressFactory() PressFactory {
	return PressFactory{}
}

func init() {
	Press = NewPressFactory()
}

func (f *PressFactory) Create() types.Press {
	return types.NewPress(models.Press{
		ID:       1,
		Text:     null.StringFrom(faker.Sentence()),
		PlayerID: -1,
		Date:     time.Now(),
		Subject:  null.StringFrom(faker.Word()),
		Type:     "P",
	})
}
