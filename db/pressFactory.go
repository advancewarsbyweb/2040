package db

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
)

type pressFactory struct {
	Press models.Press
}

var PressFactory pressFactory

func NewPressFactory() pressFactory {
	return pressFactory{}
}

func init() {
	PressFactory = NewPressFactory()
}

func (f *pressFactory) Create() *pressFactory {
	f.Press = models.Press{
		ID:       1,
		Text:     null.StringFrom(faker.Sentence()),
		PlayerID: -1,
		Date:     time.Now(),
		Subject:  null.StringFrom(faker.Word()),
		Type:     "P",
	}
	return f
}

func (f *pressFactory) CreateRelations() *pressFactory {
	p := PlayerFactory.CreateRelations().BuildInsert()
	f.Press.PlayerID = p.ID
	return f
}

func (f *pressFactory) SetPlayer(p *models.Player) *pressFactory {
	f.Press.PlayerID = p.ID
	f.Press.Player = p
	return f
}

func (f *pressFactory) Build() models.Press {
	return f.Press
}

func (f *pressFactory) BuildInsert(sendToIds []int) models.Press {
	pID, _ := PressRepo.CreatePress(f.Press, sendToIds)
	f.Press.ID = pID
	return f.Press
}
