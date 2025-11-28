package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"test-praweda-app-eng/internal/model"

	"github.com/gin-gonic/gin"
)

type ManipulasiHandler struct {
}

func NewManipulasiHandler() *ManipulasiHandler {
	return &ManipulasiHandler{}
}

func (h *ManipulasiHandler) ManipulasiString(ctx *gin.Context) {
	results := ctx.Query("results")
	if results == "" {
		results = "10"
	}

	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}

	// fetch external API
	resp, err := http.Get(fmt.Sprintf("https://randomuser.me/api?results=%s&page=%s", results, page))
	if err != nil {
		log.Println("Error fetching external API:", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:      http.StatusInternalServerError,
			Err:       "failed to fetch external API",
			IsSuccess: false,
		})
		return
	}
	defer resp.Body.Close()

	// process data from resp.Body if needed
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:      http.StatusInternalServerError,
			Err:       "failed to read response",
			IsSuccess: false,
		})
		return
	}

	// unmashal JSON response
	var randomUserResponse model.RandomUserResponse
	if err = json.Unmarshal(body, &randomUserResponse); err != nil {
		log.Println("Error unmarshaling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Code:      http.StatusInternalServerError,
			Err:       "failed to unmarshal JSON",
			IsSuccess: false,
		})
		return
	}

	// map to manipulated data
	manipulated := mapToManipulated(randomUserResponse)

	// return response
	ctx.JSON(http.StatusOK, model.ResponseData{
		Response: model.Response{
			Code:      http.StatusOK,
			IsSuccess: true,
		},
		Data: gin.H{
			"results": results,
			"page":    page,
			"data":    manipulated,
		},
	})
}

func mapToManipulated(r model.RandomUserResponse) []model.ManipulatedData {
	var out []model.ManipulatedData
	for _, u := range r.Results {
		// build full name
		fullName := fmt.Sprintf("%s %s %s", u.Name.Title, u.Name.First, u.Name.Last)

		// build location string. Postcode is PostalCode type which is underlying string
		location := fmt.Sprintf("%d %s, %s, %s, %s, %s", u.Location.Street.Number, u.Location.Street.Name, u.Location.City, u.Location.State, u.Location.Country, string(u.Location.Postcode))

		pictures := []string{u.Picture.Large, u.Picture.Medium, u.Picture.Thumbnail}

		out = append(out, model.ManipulatedData{
			Name:     fullName,
			Location: location,
			Email:    u.Email,
			Age:      u.Dob.Age,
			Phone:    u.Phone,
			Cell:     u.Cell,
			Picture:  pictures,
		})
	}
	return out
}
