package relation

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"gorm.io/gorm"
)

type IUserFollowRelation interface {
	Update(ctx context.Context, old *model.UserFollowed, IsFallow bool) error
	Get(ctx context.Context, old *model.UserFollowed) (rel *model.UserFollowed, err error)
	Create(ctx context.Context, old *model.UserFollowed) error
	FirstOrCreate(ctx context.Context, fromUserID, ToUserID uint, fromUserName, toUserName string) (*model.UserFollowed, error)
	Transaction(ctx context.Context, f func(tx *gorm.DB) error) error
	FollowList(ctx context.Context, l int, userID uint) ([]model.User, error)
	FollowerListLen(ctx context.Context, l int, userID uint) ([]model.User, error)
	FollowerList(ctx context.Context, userID uint) ([]model.User, error)
	FriendList(ctx context.Context, uID uint) ([]model.User, error)
}

type SUserFollowRelation struct {
	db *gorm.DB
}

var _ IUserFollowRelation = (*SUserFollowRelation)(nil)

func NewSUserFollowRelation(db *gorm.DB) *SUserFollowRelation {
	return &SUserFollowRelation{db: db}
}

func (s *SUserFollowRelation) Update(ctx context.Context, old *model.UserFollowed, IsFallow bool) error {
	return s.db.Model(old).Update("IsFollow", IsFallow).Error
}
func (s *SUserFollowRelation) Get(ctx context.Context, old *model.UserFollowed) (rel *model.UserFollowed, err error) {
	rel = &model.UserFollowed{}
	err = s.db.Model(old).First(rel).Error
	return rel, err
}

func (s *SUserFollowRelation) Create(ctx context.Context, new *model.UserFollowed) error {
	return s.db.Create(new).Error
}

func (s *SUserFollowRelation) FirstOrCreate(ctx context.Context, fromUserID, ToUserID uint, fromUserName, toUserName string) (*model.UserFollowed, error) {
	var r model.UserFollowed
	err := s.db.FirstOrCreate(&r, model.UserFollowed{UserID: fromUserID, UserName: fromUserName, ToUserID: ToUserID, ToUserName: toUserName}).Error
	return &r, err
}

func (s *SUserFollowRelation) Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	return s.db.Transaction(f)
}

func (s *SUserFollowRelation) FollowList(ctx context.Context, l int, userID uint) ([]model.User, error) {
	r := make([]model.User, l)
	err := s.db.Table("users").Where("user_id IN (?)", s.db.Table("user_followeds").
		Select("to_user_id").Where("user_id=? AND is_follow=1", userID)).Find(&r).Error
	return r, err
}

func (s *SUserFollowRelation) FollowerListLen(ctx context.Context, l int, userID uint) ([]model.User, error) {
	r := make([]model.User, l)
	err := s.db.Table("users").Where("user_id IN (?)", s.db.Table("user_followeds").
		Select("user_id").Where("to_user_id=? AND is_follow=1", userID)).Find(&r).Error
	return r, err
}

func (s *SUserFollowRelation) FollowerList(ctx context.Context, userID uint) ([]model.User, error) {
	var r []model.User
	err := s.db.Table("users").Where("user_id IN (?)", s.db.Table("user_followeds").
		Select("user_id").Where("to_user_id=? AND is_follow=1", userID)).Find(&r).Error
	return r, err
}

func (s *SUserFollowRelation) FriendList(ctx context.Context, uID uint) ([]model.User, error) {
	var r []model.User
	//选取uName的关注用户list
	sub := s.db.Table("user_followeds").Select("to_user_id").Where("is_follow=1 AND user_id=?", uID)
	//再选取list中关注uID的用户
	err := s.db.Table("users").Where("user_id IN (?)",
		s.db.Table("user_followeds").Select("user_id").Where("is_follow=1 AND user_id IN (?)", sub)).
		Find(&r).Error
	return r, err
}
