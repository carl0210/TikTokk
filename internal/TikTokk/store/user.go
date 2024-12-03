package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, u *model.User) error
	Update(ctx context.Context, name string, u *model.User) error
	Delete(ctx context.Context, name string) error

	//Get
	Get(ctx context.Context, f *model.User) (u *model.User, err error)
}

type SUser struct {
	db *gorm.DB
	rc *redis.Client
}

var _ UserStore = (*SUser)(nil)

func NewUsers(db *gorm.DB, rc *redis.Client) *SUser {
	r := SUser{db: db, rc: rc}
	return &r
}

func (s *SUser) Get(ctx context.Context, f *model.User) (*model.User, error) {
	var u model.User
	if err := s.db.Where(f).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil

}

func (s *SUser) Create(ctx context.Context, u *model.User) error {
	return s.db.Create(u).Error
}

func (s *SUser) Update(ctx context.Context, name string, u *model.User) error {
	return s.db.Where("name=?", name).Updates(u).Error
}

func (s *SUser) Delete(ctx context.Context, name string) error {
	return s.db.Delete(&model.User{Name: name}).Error
}
