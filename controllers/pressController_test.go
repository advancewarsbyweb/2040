package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	db.NewTestDatabase()
}

func TestGetAll(t *testing.T) {
	assert := assert.New(t)

	g := db.GameFactory.Create().BuildInsert()

	p1 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()
	p2 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()
	p3 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()

	_ = db.PressFactory.Create().SetPlayer(&p1).BuildInsert([]int{p2.ID, p3.ID})
	_ = db.PressFactory.Create().SetPlayer(&p2).BuildInsert([]int{p1.ID, p3.ID})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "GET",
		Header: http.Header{},
	}
	c.Params = []gin.Param{
		{
			Key:   "playerId",
			Value: strconv.Itoa(p1.ID),
		},
	}
	c.Set("PlayerUser", p1)

	Press.GetAll(c)

	if w.Result().StatusCode != http.StatusOK {
		var response interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		t.Fatalf("Wrong status code returned. Got (%d), want (%d). %v", w.Result().StatusCode, http.StatusOK, response)
	}

	var response []models.Press
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(len(response), 2, "Incorrect amount of Press returned")
}

func TestCreatePress(t *testing.T) {
	assert := assert.New(t)

	g := db.GameFactory.Create().BuildInsert()

	p1 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()
	p2 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()
	p3 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()

	press := db.PressFactory.Create().SetPlayer(&p3).Build()

	data, _ := json.Marshal(map[string]interface{}{
		"playerId":   press.PlayerID,
		"subject":    press.GetSubject(),
		"text":       press.GetText(),
		"recipients": []int{p1.ID, p2.ID},
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "POST",
		Body:   io.NopCloser(bytes.NewBuffer(data)),
		Header: http.Header{},
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("PlayerUser", p3)

	Press.Create(c)

	if w.Result().StatusCode != http.StatusOK {
		var response interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		t.Fatalf("Wrong status code returned. Got (%d), want (%d). %v", w.Result().StatusCode, http.StatusOK, response)
	}

	var res models.Press
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(press.GetSubject(), res.GetSubject(), "Press subject does not match given one. Got (%s), want (%s)")
}
