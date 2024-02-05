package cos

import "github.com/awbw/2040/models"

type andy struct {
	models.Co
}

func NewAndy() models.ICo {
	return &andy{}
}
