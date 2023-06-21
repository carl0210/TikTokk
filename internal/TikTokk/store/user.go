package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"gorm.io/gorm"
)

type SUser struct {
	db *gorm.DB
}

type UserStore interface {
	GetByName(ctx context.Context, name string) (u *model.User, err error)
	GetByID(ctx context.Context, ID uint) (u *model.User, err error)
	Get(ctx context.Context, f *model.User) (u *model.User, err error)
	Create(ctx context.Context, u *model.User) error
	Update(ctx context.Context, name string, u *model.User) error
	Delete(ctx context.Context, name string) error
}

var _ UserStore = (*SUser)(nil)

func NewUsers(db *gorm.DB) *SUser {
	r := SUser{db: db}
	return &r
}

func (s *SUser) Get(ctx context.Context, f *model.User) (*model.User, error) {
	var u model.User
	err := s.db.Where(f).First(&u).Error
	return &u, err

}

func (s *SUser) GetByName(ctx context.Context, name string) (*model.User, error) {
	var u model.User
	err := s.db.Where("name=?", name).First(&u).Error
	return &u, err
}

func (s *SUser) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var u model.User
	err := s.db.Table("users").Where("user_id=?", id).First(&u).Error
	return &u, err
}

//func GetUserByName(name string) (u model.user) {
//	tools.DB.Where("name=?", name).First(&u)
//	return u
//}

//func GetUserByID(id string) (u model.user, exist bool) {
//	tools.DB.Where("user_id=?", id).First(&u)
//	if u.Name == "" {
//		return model.user{}, false
//	}
//	return u, true
//}

//func CreateUser(u *model.user) {
//	tools.DB.Create(u)
//}

func (s *SUser) Create(ctx context.Context, u *model.User) error {
	return s.db.Create(u).Error
}

//	Update(ctx context.Context, u *model.user) error
//	Delete(ctx context.Context, username string) error

func (s *SUser) Update(ctx context.Context, name string, u *model.User) error {
	return s.db.Where("name=?", name).Updates(u).Error
}

func (s *SUser) Delete(ctx context.Context, name string) error {
	return s.db.Delete(&model.User{Name: name}).Error
}

//func UpdateUser(old *model.user, new *model.user) {
//	tools.DB.Model(old).Updates(*new)
//}
