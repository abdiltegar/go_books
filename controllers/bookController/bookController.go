package bookController

import (
	"github.com/labstack/echo/v4"
	"learn_orm/dto"
	"learn_orm/services/book"
	"net/http"
	"strconv"
)

type bookController struct {
	bookService book.BookService
}

func NewBookController(srv book.BookService) *bookController {
	return &bookController{
		srv,
	}
}

// get all books
func (ctrl *bookController) GetBooksController(c echo.Context) error {
	books, err := ctrl.bookService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func (ctrl *bookController) GetBookController(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	book, err := ctrl.bookService.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// create book
func (ctrl *bookController) CreateBookController(c echo.Context) error {
	book := dto.DTOBookReq{}
	c.Bind(&book)

	err := ctrl.bookService.Create(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create book",
	})
}

// delete book by id
func (ctrl *bookController) DeleteBookController(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	err := ctrl.bookService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func (ctrl *bookController) UpdateBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 16, 32)

	book := dto.DTOBookReq{}
	c.Bind(&book)

	err := ctrl.bookService.Update(uint(id), book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
	})
}
