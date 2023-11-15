package db

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
)

type PressFactory struct {
	Press types.Press
}

var Press PressFactory

func NewPressFactory() PressFactory {
	return PressFactory{}
}

func init() {
	Press = NewPressFactory()
}

func (f *PressFactory) Create() *PressFactory {
	f.Press = types.NewPress(models.Press{
		ID:       1,
		Text:     null.StringFrom(faker.Sentence()),
		PlayerID: -1,
		Date:     time.Now(),
		Subject:  null.StringFrom(faker.Word()),
		Type:     "P",
	})
	return f
}

func (f *PressFactory) CreateRelations() *PressFactory {
	p := Player.CreateRelations().BuildInsert()
	f.Press.PlayerID = p.ID
	return f
}

func (f *PressFactory) SetPlayer(p *types.Player) *PressFactory {
	f.Press.PlayerID = p.ID
	return f
}

func (f *PressFactory) Build() types.Press {
	return f.Press
}

func (f *PressFactory) BuildAndInsert(sendToIds []int) types.Press {
	pID, _ := PressRepo.CreatePress(f.Press, sendToIds)
	f.Press.ID = pID
	return f.Press
}
