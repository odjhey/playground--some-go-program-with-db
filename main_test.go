package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
