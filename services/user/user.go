package user

import (
	"learn_orm/dto"
	"learn_orm/middlewares"
	"learn_orm/models"
	"learn_orm/repository"
)

type UserService interface {
	GetAll() ([]dto.DTOUserRes, error)
	GetById(id uint) (dto.DTOUserRes, error)
	Create(payload dto.DTOUserReq) error
	Update(id uint, payload dto.DTOUserReq) error
	Delete(id uint) error
	Login(payload dto.DTOUserReq) (dto.DTOLoginRes, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(rep repository.UserRepository) *userService {
	return &userService{rep}
}

func (srv *userService) GetAll() ([]dto.DTOUserRes, error) {
	res := []dto.DTOUserRes{}

	users, err := srv.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		dtoUser := dto.DTOUserRes{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
		}

		res = append(res, dtoUser)
	}

	return res, err
}

func (srv *userService) GetById(id uint) (dto.DTOUserRes, error) {

	user, err := srv.userRepository.GetById(id)
	if err != nil {
		return dto.DTOUserRes{}, err
	}

	res := dto.DTOUserRes{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}

	return res, err
}

func (srv *userService) Create(payload dto.DTOUserReq) error {
	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	return srv.userRepository.Create(user)
}

func (srv *userService) Update(id uint, payload dto.DTOUserReq) error {

	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	return srv.userRepository.Update(id, user)
}

func (srv *userService) Delete(id uint) error {
	return srv.userRepository.Delete(id)
}

func (srv *userService) Login(payload dto.DTOUserReq) (dto.DTOLoginRes, error) {

	user, err := srv.userRepository.GetByEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return dto.DTOLoginRes{}, err
	}

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err != nil {
		return dto.DTOLoginRes{}, err
	}

	res := dto.DTOLoginRes{
		user.ID,
		user.Name,
		user.Email,
		token,
	}

	return res, nil
}
