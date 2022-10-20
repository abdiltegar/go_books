package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"learn_orm/constants"
	"learn_orm/controllers/bookController"
	"learn_orm/controllers/userController"
	mdlwr "learn_orm/middlewares"
	"learn_orm/repository"
	"learn_orm/services/book"
	"learn_orm/services/user"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	// Repositories
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Services
	bookService := book.NewBookService(bookRepo)
	userService := user.NewUserService(userRepo)

	// Controllers
	bookCtrl := bookController.NewBookController(bookService)
	usrCtrl := userController.NewUserController(userService)

	e.POST("/login", usrCtrl.LoginController)

	jwt := middleware.JWT([]byte(constants.SECRET_JWT))

	// Route For Users
	usrRoute := e.Group("/users")
	usrRoute.GET("", usrCtrl.GetUsersController, jwt)
	usrRoute.GET("/:id", usrCtrl.GetUserController, jwt)
	usrRoute.POST("", usrCtrl.CreateUserController)
	usrRoute.DELETE("/:id", usrCtrl.DeleteUserController, jwt)
	usrRoute.PUT("/:id", usrCtrl.UpdateUserController, jwt)

	// Route For Books
	bookRoute := e.Group("/books")
	bookRoute.GET("", bookCtrl.GetBooksController)
	bookRoute.GET("/:id", bookCtrl.GetBookController)
	bookRoute.POST("", bookCtrl.CreateBookController, jwt)
	bookRoute.DELETE("/:id", bookCtrl.DeleteBookController, jwt)
	bookRoute.PUT("/:id", bookCtrl.UpdateBookController, jwt)

	mdlwr.LogMiddleware(e)

	return e
}
