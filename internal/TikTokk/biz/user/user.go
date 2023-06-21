package user

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/TikTokk/store"
	"context"
	"fmt"
	"strconv"
)

type UserBiz interface {
	Login(c context.Context, req *api.LoginUserRequest) (rsp *api.LoginUserRespond, err error)
	Register(c context.Context, req *api.RegisterUserRequest) (rsp *api.RegisterUserRespond, err error)
	GetDetail(c context.Context, req *api.GetDetailUserRequest) (rsp *api.GetDetailUserRespond, err error)
}

type BUser struct {
	ds store.DataStore
}

var _ UserBiz = (*BUser)(nil)

func New(b store.DataStore) *BUser {
	return &BUser{ds: b}
}

func (b *BUser) Login(ctx context.Context, req *api.LoginUserRequest) (rsp *api.LoginUserRespond, err error) {
	//得到要登录的用户名和密码
	username := req.Username
	password := req.Password
	//得到要要登录的账号结构体
	u, err := b.ds.Users().GetByName(ctx, username)
	//如果不存在账号则返回
	if err != nil {
		rsp = &api.LoginUserRespond{UserID: -1}
		return rsp, err
	}
	//存在则对比密码是否一致,一致则返回id
	if u.Password == password {
		rsp = &api.LoginUserRespond{UserID: int64(u.UserId)}
		return rsp, nil
	}
	//不一致
	rsp = &api.LoginUserRespond{UserID: -1}
	return rsp, fmt.Errorf("密码错误")
}

func (b *BUser) Register(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserRespond, error) {
	var rsp api.RegisterUserRespond
	//获取用户名和密码
	username := req.Username
	password := req.Password
	//查询用户名是否存在,用户名存在则返回错误
	if _, err := b.ds.Users().GetByName(ctx, username); err == nil {
		rsp = api.RegisterUserRespond{UserID: -1}
		return &rsp, fmt.Errorf("用户名存在")
	}
	//用户名不存在则创建用户
	user := model.User{
		Name:            username,
		Password:        password,
		Avatar:          "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		BackgroundImage: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		TotalFavorited:  "0",
		Signature:       "hello,world",
	}
	if err := b.ds.Users().Create(ctx, &user); err != nil {
		//创建失败
		rsp = api.RegisterUserRespond{UserID: -1}
		return &rsp, err
	}
	//创建成功则获取用户id
	u, err := b.ds.Users().GetByName(ctx, username)
	//获取失败
	if err != nil {
		rsp = api.RegisterUserRespond{UserID: -1}
		return &rsp, err
	}
	//获取成功则返回
	rsp = api.RegisterUserRespond{UserID: int64(u.UserId)}
	return &rsp, nil
}
func (b *BUser) GetDetail(ctx context.Context, req *api.GetDetailUserRequest) (*api.GetDetailUserRespond, error) {
	var rsp api.GetDetailUserRespond
	//获取被查询用户具体信息
	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		rsp.User.ID = -1
		return &rsp, err
	}
	u, err := b.ds.Users().GetByID(ctx, uint(userID))
	if err != nil {
		//不存在该用户则返回错误
		rsp.User.ID = -1
		return &rsp, err
	}
	//存在则返回详情
	rsp = api.GetDetailUserRespond{User: api.UserDetailRespond{
		Name:            u.Name,
		BackgroundImage: u.BackgroundImage,
		FavoriteCount:   int64(u.FavoriteCount),
		FollowCount:     int64(u.FollowCount),
		FollowerCount:   int64(u.FollowerCount),
		ID:              int64(u.UserId),
		Avatar:          u.Avatar,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       int64(u.WorkCount),
	}}
	//得到关注状态
	//得到发起查询的人的信息
	username := ctx.Value("username").(string)
	user, err := b.ds.Users().GetByName(ctx, username)
	if err != nil {
		rsp.User.IsFollow = false
		return &rsp, err
	}
	//查询记录
	f := model.UserFollowed{UserID: user.UserId, UserName: user.Name, ToUserID: u.UserId, ToUserName: u.Name}
	rel, err := b.ds.UserFollowRelation().Get(ctx, &f)
	//如果记录不存在则创建记录并返回该记录
	if err != nil {
		f.IsFollow = false
		err := b.ds.UserFollowRelation().Create(ctx, &f)
		//如果创建失败
		if err != nil {
			return nil, err
		}
		//创建成功则返回
		rsp.User.IsFollow = f.IsFollow
		return &rsp, nil
	}
	//存在则写入并返回
	rsp.User.IsFollow = rel.IsFollow
	return &rsp, nil

}
