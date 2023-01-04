package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Expense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func TestGetAllExpense(t *testing.T) {
	seedExpense(t)
	var custs []Expense

	res := request(http.MethodGet, uri("expense"), nil)
	err := res.Decode(&custs)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(custs), 0)
}

func TestCreateExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		
		"title"		: "orange juice",
		"amount"	: 1.0,
		"note"		: "no discount",
		"tags"		: [
			"beverage"
		  ]
	}`)
	var cust Expense

	res := request(http.MethodPost, uri("expense"), body)
	err := res.Decode(&cust)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, cust.ID)
	assert.Equal(t, "orange juice", cust.Title)
	assert.Equal(t, 1.0, cust.Amount)
	assert.Equal(t, "no discount", cust.Note)
	assert.Equal(t, []string{"beverage"}, cust.Tags)
}

func TestGetExpenseByID(t *testing.T) {
	c := seedExpense(t)

	var lastCust Expense
	res := request(http.MethodGet, uri("expense", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&lastCust)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, lastCust.ID)
	assert.NotEmpty(t, lastCust.Title)
	assert.NotEmpty(t, lastCust.Amount)
	assert.NotEmpty(t, lastCust.Note)
	assert.NotEmpty(t, lastCust.Tags)
}

func TestUpdateExpense(t *testing.T) {
	id := seedExpense(t).ID
	c := Expense{
		ID		:     id,
		Title	:  "orange juice",
		Amount	: 1,
		Note	:   "no discount",
		Tags	:   []string{"beverage"},
	}
	payload, _ := json.Marshal(c)
	res := request(http.MethodPut, uri("expense", strconv.Itoa(id)), bytes.NewBuffer(payload))
	var info Expense
	err := res.Decode(&info)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.Title, info.Title)
	assert.Equal(t, c.Amount, info.Amount)
	assert.Equal(t, c.Note, info.Note)
	assert.Equal(t, c.Tags, info.Tags)
}

func seedExpense(t *testing.T) Expense {
	var c Expense
	body := bytes.NewBufferString(`{
		
		"title": "orange juice",
		"amount": 1.0,
		"note": "no discount",
		"tags": [
			"beverage"
		  ]
	}`)
	err := request(http.MethodPost, uri("expense"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create expense:", err)
	}
	return c
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	token := os.Getenv("AUTH_TOKEN")
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
