package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"gorm.io/gorm"
	"time"
)

type SVideo struct {
	db *gorm.DB
}

func NewVideos(db *gorm.DB) *SVideo {
	return &SVideo{db: db}
}

type VideoStore interface {
	GetByAuthorID(ctx context.Context, authorID uint) (v *model.Video, err error)
	GetByVideoID(ctx context.Context, videoID uint) (v *model.Video, err error)
	Create(ctx context.Context, u *model.Video) error
	Update(ctx context.Context, videoId uint, v *model.Video) error
	Delete(ctx context.Context, videoId uint) error
	List(ctx context.Context, lastTime time.Time) (list []model.Video, err error)
	Feed(ctx context.Context, l int, lastTime time.Time) (list []model.Video, err error)
	ListAllVideoByAuthorID(ctx context.Context, authorID uint) (list []model.Video, err error)
	ListAllVideoByAuthorIDLen(ctx context.Context, authorID uint, l int) (list []model.Video, err error)
}

var _ VideoStore = (*SVideo)(nil)

func (s *SVideo) GetByAuthorID(ctx context.Context, authorID uint) (*model.Video, error) {
	var v model.Video
	err := s.db.Table("videos").Where("author_id=?", authorID).First(&v).Error
	return &v, err
}

func (s *SVideo) GetByVideoID(ctx context.Context, VideoID uint) (*model.Video, error) {
	var v model.Video
	err := s.db.Table("videos").Where("video_id=?", VideoID).First(&v).Error
	return &v, err
}

func (s *SVideo) Create(ctx context.Context, v *model.Video) error {
	return s.db.Create(v).Error
}

func (s *SVideo) Update(ctx context.Context, videoId uint, v *model.Video) error {
	return s.db.Model(&model.Video{VideoId: videoId}).Save(v).Error
}

func (s *SVideo) Delete(ctx context.Context, videoId uint) error {
	return s.db.Delete(&model.Video{VideoId: videoId}).Error
}

func (s *SVideo) List(ctx context.Context, lastTime time.Time) ([]model.Video, error) {
	var list []model.Video
	err := s.db.Table("videos").Order("updated_at desc").Where("updated_at < ?", lastTime).Find(&list).Error
	return list, err
}

func (s *SVideo) Feed(ctx context.Context, l int, lastTime time.Time) ([]model.Video, error) {
	list := make([]model.Video, 0, l)
	err := s.db.Table("videos").Order("updated_at desc").Where("updated_at < ?", lastTime).Limit(l).Find(&list).Error
	return list, err
}

func (s *SVideo) ListAllVideoByAuthorID(ctx context.Context, authorID uint) ([]model.Video, error) {
	var list []model.Video
	err := s.db.Where("author_id=?", authorID).Find(&list).Error
	return list, err
}

func (s *SVideo) ListAllVideoByAuthorIDLen(ctx context.Context, authorID uint, l int) ([]model.Video, error) {
	list := make([]model.Video, l)
	err := s.db.Where("author_id=?", authorID).Find(&list).Error
	return list, err
}

//func CreateVideo(v *model.Video) {
//	tools.DB.Create(v)
//}
//
//func GetVideoByAuthorID(id string) []model.VideoRsp {
//	a, exist := GetUserByID(id)
//	if !exist {
//
//	}
//	l := make([]model.Video, a.WorkCount)
//	tools.DB.Table("videos").Where("author_id=?", id).Find(&l)
//	r := make([]model.VideoRsp, a.WorkCount)
//	for i, v := range l {
//		r[i] = VideoToResMe(&v, &a)
//	}
//	return r
//}
//
//func GetVideoFeed(u model.user) []model.VideoRsp {
//	variable t model.Video
//	variable r model.VideoRsp
//	tools.DB.Find(&t)
//	//判断观看用户是否登录
//	variable user model.user
//	if u.User_ID == 0 {
//		user = model.user{}
//	}
//	author, _ := GetUserByID(strconv.Itoa(int(t.Author_ID)))
//	Tlog.Println("author=", author)
//	r = *VideoToResOther(&t, &user, &author)
//	return []model.VideoRsp{r}
//}
//
//func UpdateVideo(old *model.Video, new *model.Video) {
//	tools.DB.Model(old).Updates(new)
//}
//
//func GetVideoByID(id uint) model.Video {
//	variable r model.Video
//	tools.DB.Where("video_id=?", id).First(&r)
//	return r
//}
