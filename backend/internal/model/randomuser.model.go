package model

import (
	"encoding/json"
	"fmt"
	"math"
)

type RandomUserResponse struct {
	Results []UserResult `json:"results"`
	Info    Info         `json:"info"`
}

type UserResult struct {
	Gender     string   `json:"gender"`
	Name       Name     `json:"name"`
	Location   Location `json:"location"`
	Email      string   `json:"email"`
	Login      Login    `json:"login"`
	Dob        DateAge  `json:"dob"`
	Registered DateAge  `json:"registered"`
	Phone      string   `json:"phone"`
	Cell       string   `json:"cell"`
	ID         UserID   `json:"id"`
	Picture    Picture  `json:"picture"`
	Nat        string   `json:"nat"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    PostalCode  `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

type DateAge struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type UserID struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"` // bisa null atau string
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type Info struct{}

// PostalCode handles postcode values that may be returned as a number or a string
// in the randomuser API. We store it as a string internally but accept both
// JSON number and JSON string during unmarshaling.
type PostalCode string

func (p *PostalCode) UnmarshalJSON(b []byte) error {
	// If value is null
	if string(b) == "null" {
		*p = ""
		return nil
	}

	// Try unmarshalling as string
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		*p = PostalCode(s)
		return nil
	}

	// Try unmarshalling as number (int)
	var i int
	if err := json.Unmarshal(b, &i); err == nil {
		*p = PostalCode(fmt.Sprintf("%d", i))
		return nil
	}

	// try as float (some postcodes may be returned as float-like)
	var f float64
	if err := json.Unmarshal(b, &f); err == nil {
		// remove trailing .0 when possible
		if f == math.Trunc(f) {
			*p = PostalCode(fmt.Sprintf("%d", int64(f)))
		} else {
			*p = PostalCode(fmt.Sprintf("%f", f))
		}
		return nil
	}

	return fmt.Errorf("cannot unmarshal PostalCode from %s", string(b))
}
