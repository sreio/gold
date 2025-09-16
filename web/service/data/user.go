package data

import (
	"errors"
	"github.com/sreio/gold/web/dto"
	"github.com/sreio/gold/web/model"
	"github.com/sreio/gold/web/repository"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(r *repository.UserRepo) *UserService { return &UserService{repo: r} }

func (s *UserService) List(q dto.QueryUser) ([]model.User, int64, error) {
	return s.repo.List(q)
}

func (s *UserService) Create(dto dto.CreateUserDTO) (*model.User, error) {
	u := &model.User{
		Name: dto.Name, Cron: dto.Cron, SaveDay: dto.SaveDay,
	}
	if err := s.repo.CreateWithConf(u, dto.UserConf); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) Update(id uint, dto dto.UpdateUserDTO) error {
	return s.repo.UpdateAndSyncConf(id, dto)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) Exits(name string) (bool, error) {
	u, err := s.repo.GetByName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // 没找到=不存在，不算错误
	}

	if err != nil {
		return false, err
	}

	if u.ID > 0 {
		return true, nil
	}

	return false, nil
}
