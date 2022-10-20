package bookController

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"learn_orm/dto"
	"learn_orm/services/book/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type BookTestSuite struct {
	suite.Suite
	bookService *mock.BookMock
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}

func (suite *BookTestSuite) TestGetAll() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
	}{
		{
			name:          "berhasil",
			path:          "/books",
			expectMessage: "success get all books",
			expectStatus:  http.StatusOK,
			method:        http.MethodGet,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
		},
	}

	for _, testCase := range testCases {

		e := echo.New()
		req := httptest.NewRequest(testCase.method, "/", nil)
		rec := httptest.NewRecorder()
		req.Header = testCase.header
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		//Assertions
		controller := NewBookController(suite.bookService)
		if suite.NoError(controller.GetBooksController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *BookTestSuite) TestGetById() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		id            uint
	}{
		{
			name:          "berhasil",
			path:          "/books/:id",
			expectMessage: "success get book",
			expectStatus:  http.StatusOK,
			method:        http.MethodGet,
			id:            1,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
		},
		{
			name:          "gagal",
			path:          "/books/:id",
			expectMessage: "data not found",
			expectStatus:  http.StatusBadRequest,
			method:        http.MethodGet,
			id:            3,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
		},
	}

	for _, testCase := range testCases {

		e := echo.New()
		req := httptest.NewRequest(testCase.method, "/", nil)
		rec := httptest.NewRecorder()
		req.Header = testCase.header
		c := e.NewContext(req, rec)

		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(int(testCase.id)))
		c.SetPath(testCase.path)

		//Assertions
		controller := NewBookController(suite.bookService)
		if suite.NoError(controller.GetBookController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *BookTestSuite) TestCreate() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		bodyParams    dto.DTOBookReq
	}{
		{
			name:          "berhasil",
			path:          "/books",
			expectMessage: "success create book",
			expectStatus:  http.StatusOK,
			method:        http.MethodPost,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			bodyParams: dto.DTOBookReq{
				Title:       "Learn Javascript",
				Description: "Lorem Ipsum",
				Author:      "Armadilo",
				Publisher:   "Gramedia",
			},
		},
	}

	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.bodyParams)

		e := echo.New()
		req := httptest.NewRequest(testCase.method, "/", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		req.Header = testCase.header
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		//Assertions
		controller := NewBookController(suite.bookService)
		if suite.NoError(controller.CreateBookController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *BookTestSuite) TestUpdate() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		id            uint
		bodyParams    dto.DTOBookReq
	}{
		{
			name:          "berhasil",
			path:          "/books/:id",
			expectMessage: "success update book",
			expectStatus:  http.StatusOK,
			method:        http.MethodPut,
			id:            1,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			bodyParams: dto.DTOBookReq{
				Title:       "Learn Javascript",
				Description: "Lorem Ipsum",
				Author:      "Armadilo",
				Publisher:   "Gramedia",
			},
		},
	}

	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.bodyParams)

		e := echo.New()
		req := httptest.NewRequest(testCase.method, "/", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		req.Header = testCase.header
		c := e.NewContext(req, rec)

		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(int(testCase.id)))
		c.SetPath(testCase.path)

		//Assertions
		controller := NewBookController(suite.bookService)
		if suite.NoError(controller.UpdateBookController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *BookTestSuite) TestDelete() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		id            int
	}{
		{
			name:          "berhasil",
			path:          "/books/:id",
			expectMessage: "success delete book",
			expectStatus:  http.StatusOK,
			method:        http.MethodDelete,
			id:            1,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
		},
	}

	for _, testCase := range testCases {

		e := echo.New()
		req := httptest.NewRequest(testCase.method, "/", nil)
		rec := httptest.NewRecorder()
		req.Header = testCase.header
		c := e.NewContext(req, rec)

		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))
		c.SetPath(testCase.path)

		//Assertions
		controller := NewBookController(suite.bookService)
		if suite.NoError(controller.DeleteBookController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}
