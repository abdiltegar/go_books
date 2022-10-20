package userController

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"learn_orm/dto"
	"learn_orm/services/user/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	userService *mock.UserMock
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) TestGetAll() {
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
			path:          "/users",
			expectMessage: "success get all users",
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
		controller := NewUserController(suite.userService)
		if suite.NoError(controller.GetUsersController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *UserTestSuite) TestGetById() {
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
			path:          "/users/:id",
			expectMessage: "success get user",
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
			path:          "/users/:id",
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
		controller := NewUserController(suite.userService)
		if suite.NoError(controller.GetUserController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *UserTestSuite) TestCreate() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		bodyParams    dto.DTOUserReq
	}{
		{
			name:          "berhasil",
			path:          "/users",
			expectMessage: "success create user",
			expectStatus:  http.StatusOK,
			method:        http.MethodPost,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			bodyParams: dto.DTOUserReq{
				Name:     "Admin",
				Email:    "admin@mail.com",
				Password: "admin123",
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
		controller := NewUserController(suite.userService)
		if suite.NoError(controller.CreateUserController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *UserTestSuite) TestUpdate() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		id            uint
		bodyParams    dto.DTOUserReq
	}{
		{
			name:          "berhasil",
			path:          "/users/:id",
			expectMessage: "success update user",
			expectStatus:  http.StatusOK,
			method:        http.MethodPut,
			id:            1,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			bodyParams: dto.DTOUserReq{
				Name:     "Admin",
				Email:    "admin@mail.com",
				Password: "admin123",
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
		controller := NewUserController(suite.userService)
		if suite.NoError(controller.UpdateUserController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *UserTestSuite) TestDelete() {
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
			path:          "/users/:id",
			expectMessage: "success delete user",
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
		controller := NewUserController(suite.userService)
		if suite.NoError(controller.DeleteUserController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}

func (suite *UserTestSuite) TestLogin() {
	var testCases = []struct {
		name          string
		path          string
		expectStatus  int
		expectMessage string
		header        map[string][]string
		method        string
		bodyParams    dto.DTOUserReq
	}{
		{
			name:          "berhasil",
			path:          "/login",
			expectMessage: "login success",
			expectStatus:  http.StatusOK,
			method:        http.MethodPut,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			bodyParams: dto.DTOUserReq{
				Name:     "Admin",
				Email:    "admin@mail.com",
				Password: "admin123",
			},
		},
		{
			name:          "berhasil",
			path:          "/login",
			expectMessage: "data not found",
			expectStatus:  http.StatusBadRequest,
			method:        http.MethodPut,
			header: map[string][]string{
				"Content-Type":    {"application/json"},
				"Accept":          {"*/*"},
				"Accept-Encoding": {"gzip", "deflate", "br"},
			},
			bodyParams: dto.DTOUserReq{
				Name:     "Admin",
				Email:    "admin2@mail.com",
				Password: "admin123",
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
		controller := NewUserController(suite.userService)
		if suite.NoError(controller.LoginController(c)) {
			suite.Equal(testCase.expectStatus, rec.Code)

			var resp map[string]interface{}
			err := json.NewDecoder(rec.Result().Body).Decode(&resp)
			suite.NoError(err)
			suite.Equal(testCase.expectMessage, resp["message"])
		}
	}
}
