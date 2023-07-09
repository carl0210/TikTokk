package user

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/encryption"
	"TikTokk/internal/pkg/token"
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
)

type UserBiz interface {
	Login(c context.Context, req *api.LoginUserRequest) (rsp *api.LoginUserRespond, err error)
	Register(c context.Context, req *api.RegisterUserRequest) (rsp *api.RegisterUserRespond, err error)
	GetDetail(c context.Context, req *api.GetDetailUserRequest, userID uint) (rsp *api.GetDetailUserRespond, err error)
}

type BUser struct {
	ds store.DataStore
}

var _ UserBiz = (*BUser)(nil)

func New(b store.DataStore) *BUser {
	return &BUser{ds: b}
}

// Login 数据库中查找用户,对比密码是否相同。如果不相同或出现错误,返回用户id=-1的rsp,否则返回真实用户id的rsp
func (b *BUser) Login(ctx context.Context, req *api.LoginUserRequest) (rsp *api.LoginUserRespond, err error) {
	//得到要要登录的账号结构体
	u, err := b.ds.Users().Get(ctx, &model.User{Name: req.Username})
	//如果不存在账号则返回
	if err != nil {
		return &api.LoginUserRespond{UserID: 0}, err
	}
	//存在则对比密码是否一致,一致则返回id
	if encryption.CheckPassword(req.Password, u.Password) {
		rsp = &api.LoginUserRespond{UserID: int64(u.UserID), Token: token.Sign(strconv.Itoa(int(u.UserID)))}
		return rsp, nil
	}
	//不一致
	return &api.LoginUserRespond{UserID: 0}, fmt.Errorf("密码错误")
}

// Register
func (b *BUser) Register(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserRespond, error) {
	//查询用户名是否存在,用户名存在则返回错误
	if _, err := b.ds.Users().Get(ctx, &model.User{Name: req.Username}); err == nil {
		return &api.RegisterUserRespond{UserID: -1}, fmt.Errorf("用户名存在")
	}
	//加密密码
	enPw := encryption.Encryption(req.Password)
	pw := hex.EncodeToString(enPw[:])
	//用户名不存在则创建用户
	user := model.User{
		Name:            req.Username,
		Password:        pw,
		Avatar:          "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		BackgroundImage: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		TotalFavorite:   "0",
		Signature:       "hello,world",
	}
	if err := b.ds.Users().Create(ctx, &user); err != nil {
		//创建失败
		return &api.RegisterUserRespond{UserID: -1}, err
	}
	//创建成功则获取用户id
	u, err := b.ds.Users().Get(ctx, &model.User{Name: req.Username})
	//获取失败
	if err != nil {
		return &api.RegisterUserRespond{UserID: -1}, err
	}
	//获取成功则返回
	return &api.RegisterUserRespond{UserID: int64(u.UserID), Token: token.Sign(strconv.Itoa(int(u.UserID)))}, nil
}
func (b *BUser) GetDetail(ctx context.Context, req *api.GetDetailUserRequest, userID uint) (*api.GetDetailUserRespond, error) {
	u, err := b.ds.Users().Get(ctx, &model.User{UserID: uint(req.UserID)})
	if err != nil {
		//不存在该用户则返回错误
		return &api.GetDetailUserRespond{User: api.UserDetailRespond{ID: 0}}, err
	}
	//存在则返回详情
	rsp := api.GetDetailUserRespond{User: api.UserDetailRespond{
		Name:            u.Name,
		BackgroundImage: u.BackgroundImage,
		FavoriteCount:   u.FavoriteCount,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		ID:              int64(u.UserID),
		Avatar:          u.Avatar,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorite,
		WorkCount:       u.WorkCount,
	}}
	//得到关注状态
	//得到发起查询的人的信息
	user, err := b.ds.Users().Get(ctx, &model.User{UserID: userID})
	if err != nil {
		rsp.User.IsFollow = false
		return &rsp, err
	}
	//查询记录,如果记录不存在则创建记录并返回该记录
	rel, err := b.ds.UserFollowRelation().FirstOrCreate(ctx, user.UserID, u.UserID, user.Name, u.Name)
	if err != nil {
		return nil, err
	}
	//写回
	rsp.User.IsFollow = rel.IsFollow
	return &rsp, nil

}
