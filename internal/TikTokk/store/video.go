package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type VideoStore interface {
	Get(ctx context.Context, video *model.Video) (v *model.Video, err error)
	Create(ctx context.Context, u *model.Video) error
	Update(ctx context.Context, videoId uint, v *model.Video) error
	Delete(ctx context.Context, videoId uint) error
	List(ctx context.Context, lastTime time.Time) (list []model.Video, err error)
	Feed(ctx context.Context, l int, lastTime time.Time) (list []model.Video, err error)
	ListAllVideoByAuthorIDLen(ctx context.Context, authorID uint, l int64) (list []model.Video, err error)
}

type SVideo struct {
	db *gorm.DB
	rc *redis.Client
}

var _ VideoStore = (*SVideo)(nil)

func NewVideos(db *gorm.DB, rc *redis.Client) *SVideo {
	return &SVideo{db: db, rc: rc}
}

func (s *SVideo) Get(ctx context.Context, video *model.Video) (*model.Video, error) {
	var v model.Video
	err := s.db.Where(video).First(&v).Error
	return &v, err
}

func (s *SVideo) Create(ctx context.Context, v *model.Video) error {
	return s.db.Create(v).Error
}

func (s *SVideo) Update(ctx context.Context, videoId uint, v *model.Video) error {
	return s.db.Model(&model.Video{VideoID: videoId}).Save(v).Error
}

func (s *SVideo) Delete(ctx context.Context, videoId uint) error {
	return s.db.Delete(&model.Video{VideoID: videoId}).Error
}

func (s *SVideo) List(ctx context.Context, lastTime time.Time) ([]model.Video, error) {
	var list []model.Video
	err := s.db.Table("videos").Order("updated_at desc").Where("updated_at < ?", lastTime).Find(&list).Error
	return list, err
}

func (s *SVideo) Feed(ctx context.Context, l int, lastTime time.Time) ([]model.Video, error) {
	// //如果不存在主键,则直接通过数据库查询,并将结果同步到redis
	// redisKey := fmt.Sprintf("videos-feed")
	// //存在主键则通过redis查询
	// jStr, err := s.rc.Get(ctx, redisKey).Result()
	// if err != nil {
	// 	if err == redis.Nil {
	// 		return s.FeedPartOfMysqlAndSyncToRedis(ctx, l, lastTime, redisKey)
	// 	}
	// 	return nil, err
	// }
	// //如果redis查询不到则通过数据库查询
	// if jStr == "" {
	// 	return s.FeedPartOfMysqlAndSyncToRedis(ctx, l, lastTime, redisKey)
	// }
	// //如果有记录,则解码得到的json
	// ru := make([]model.VideoRedis, l)
	// if err := json.Unmarshal([]byte(jStr), &ru); err != nil {
	// 	return nil, err
	// }
	// //转化并返回
	// rup := make([]model.Video, len(ru))
	// for i, v := range ru {
	// 	rup[i] = v.ToMysqlStruct()
	// }
	// return rup, nil
	//查询数据库
	list := make([]model.Video, 0, l)
	err := s.db.WithContext(ctx).Table("videos").Order("updated_at desc").Where("updated_at < ?", lastTime).Limit(l).Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil

}

func (s *SVideo) FeedPartOfMysqlAndSyncToRedis(ctx context.Context, l int, lastTime time.Time, redisKey string) ([]model.Video, error) {
	//查询数据库
	list := make([]model.Video, 0, l)
	err := s.db.WithContext(ctx).Table("videos").Order("updated_at desc").Where("updated_at < ?", lastTime).Limit(l).Find(&list).Error
	//将结果同步到redis
	err = SyncToRedis(ctx, s.rc, redisKey, list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func (s *SVideo) ListAllVideoByAuthorIDLen(ctx context.Context, authorID uint, l int64) ([]model.Video, error) {
	// //先从redis查询
	// redisKey := fmt.Sprintf("videos-publishList-%d", authorID)
	// j, err := s.rc.Get(ctx, redisKey).Result()
	// if err != nil {
	// 	return nil, err
	// }
	// //如果查询不到则访问数据库,并同步到redis
	// if j == "" {
	// 	return s.GetPartOfMysqlAndSyncToRedis(ctx, authorID, l, redisKey)
	// }
	// //如果查询到则返回结果
	// listR := make([]model.VideoRedis, l)
	// if err := json.Unmarshal([]byte(j), &listR); err != nil {
	// 	return nil, err
	// }
	// //转化并返回
	// listD := make([]model.Video, len(listR))
	// for i, v := range listR {
	// 	listD[i] = v.ToMysqlStruct()
	// }
	// return listD, nil

	//查询数据库
	list := make([]model.Video, l)
	err := s.db.Where("author_id=?", authorID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *SVideo) GetPartOfMysqlAndSyncToRedis(ctx context.Context, authorID uint, l int64, redisKey string) ([]model.Video, error) {
	//查询数据库
	list := make([]model.Video, l)
	err := s.db.Where("author_id=?", authorID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	//将结果同步到redis
	err = SyncToRedis(ctx, s.rc, redisKey, list)
	if err != nil {
		return list, err
	}
	return list, nil
}

//func (s *SVideo) ListAllVideoByAuthorIDLen(ctx context.Context, authorID uint64, l int64) ([]model.Video, error) {
//	list := make([]model.Video, l)
//	err := s.db.Where("author_id=?", authorID).Find(&list).Error
//	return list, err
//}

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
