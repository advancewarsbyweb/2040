package types

import "github.com/awbw/2040/models"

type Press struct {
	models.Press
	Text       string
	Subject    string
	Recipients string
}

func NewPress(p models.Press) Press {
	return Press{
		Press:      p,
		Text:       p.Text.String,
		Subject:    p.Subject.String,
		Recipients: p.Recipients.String,
	}
}
