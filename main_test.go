package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/odjhey/playground--some-go-program-with-db/models"
)

func TestGetMessage200(t *testing.T) {

	uri := "/message?id=123"

	r := getRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Print("Invalid Status Code")
		t.Fail()
	}

	if _, err := io.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	}

	fmt.Print()

}

func TestGetMessage404(t *testing.T) {

	uri := "/message?id=12"

	r := getRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 404 {
		fmt.Print("Invalid Status Code")
		t.Fail()
	}

	if _, err := io.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	}

	fmt.Print()

}

func TestPostMessage200(t *testing.T) {

	uri := "/message"

	r := getRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+uri, strings.NewReader(`{"id": 999, "message": "john cena"}`))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Print("Invalid Status Code ", resp.StatusCode)
		t.Fail()
	}

	p, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(string(p))

}

func TestGetMessageDb200(t *testing.T) {

	uri := "/message?id=123"

	r := getRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Print("Invalid Status Code")
		t.Fail()
	}

	if _, err := io.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	}

	fmt.Print()

}

func TestPostMessageDb200(t *testing.T) {

	uri := "/message-db"

	r := getRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	message := models.SomeMessage{ID: 999, Message: "test message"}
	jsonMessage, _ := json.Marshal(message)

	req, err := http.NewRequest("POST", ts.URL+uri, bytes.NewBuffer(jsonMessage))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Print("Invalid Status Code ", resp.StatusCode)
		t.Fail()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var responseMessage models.SomeMessage
	json.Unmarshal(body, &responseMessage)

	if responseMessage.Message != message.Message {
		fmt.Print("Invalid response message")
		t.Fail()
	}

	fmt.Print(string(body))

}
