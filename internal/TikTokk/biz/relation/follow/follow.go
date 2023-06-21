package follow

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/tools"
	"context"
	"gorm.io/gorm"
)

type FollowRelationBiz interface {
	Action(ctx context.Context, uName string, toUserID, actionType uint) error
	FollowList(ctx context.Context, userID uint) (*api.FollowListRsp, error)
	FollowerList(ctx context.Context, userID uint) (*api.FollowerListRsp, error)
	FriendList(ctx context.Context, userID uint) (*api.FriendListRsp, error)
}

type BFollowRelation struct {
	ds store.DataStore
}

var _ FollowRelationBiz = (*BFollowRelation)(nil)

func New(db store.DataStore) *BFollowRelation {
	return &BFollowRelation{ds: db}
}

func (b *BFollowRelation) Action(ctx context.Context, uName string, toUserID, actionType uint) error {
	//判断点赞关系中点赞状态和操作类型代表的点赞状态是否相同，不相同则修改，相同则不修改
	//获取点赞类型
	var isFollow bool
	var op int
	if actionType == 1 {
		isFollow = true
		op = 1
	} else {
		isFollow = false
		op = -1
	}
	//获取用户信息
	u, err := b.ds.Users().GetByName(ctx, uName)
	if err != nil {
		return err
	}
	uTo, err := b.ds.Users().GetByID(ctx, toUserID)
	if err != nil {
		return err
	}
	//关注关系
	rel, err := b.ds.UserFollowRelation().FirstOrCreate(ctx, u.UserId, uTo.UserId, u.Name, uTo.Name)
	if err != nil {
		return err
	}
	//如果跟原来的关注状态不同,则进行修改
	if !rel.IsFollow && isFollow || rel.IsFollow && !isFollow {
		//事务
		f := func(tx *gorm.DB) error {
			//修改关注关系
			if err := tx.Model(&model.UserFollowed{}).
				Where("user_id=? AND user_name=? AND to_user_id =? AND to_user_name=?",
					u.UserId, u.Name, uTo.UserId, uTo.Name).Update("is_follow", isFollow).Error; err != nil {
				return err
			}
			//用户关注数&粉丝数+1
			if err := tx.Model(&u).Update("follow_count", int(u.FollowCount)+op).Error; err != nil {
				return err
			}
			if err := tx.Model(&uTo).Update("follower_count", int(uTo.FollowerCount)+op).Error; err != nil {
				return err
			}
			return nil
		}
		err := b.ds.UserFollowRelation().Transaction(ctx, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BFollowRelation) FollowList(ctx context.Context, userID uint) (*api.FollowListRsp, error) {
	var rsp api.FollowListRsp
	//根据用户id查询关注列表
	//查询用户关注数
	u, err := b.ds.Users().GetByID(ctx, userID)
	if err != nil {
		return &rsp, err
	}
	//得到所有关注的用户
	l, err := b.ds.UserFollowRelation().FollowList(ctx, int(u.FollowCount), u.UserId)
	if err != nil {
		return &rsp, err
	}
	//将用户转化
	rspList := make([]api.UserDetailRespond, u.FollowCount)
	for i, v := range l {
		rspList[i] = *tools.UserToRsp(&v)
		rspList[i].IsFollow = true
	}
	rsp.UserList = rspList
	return &rsp, nil
}

func (b *BFollowRelation) FollowerList(ctx context.Context, userID uint) (*api.FollowerListRsp, error) {

	var rsp api.FollowerListRsp
	//根据用户id查询粉丝列表,得到所有粉丝
	l, err := b.ds.UserFollowRelation().FollowerList(ctx, userID)
	if err != nil {
		return &rsp, err
	}
	//将用户转化
	rspList := make([]api.UserDetailRespond, len(l))
	for i, v := range l {
		rspList[i] = *tools.UserToRsp(&v)
		rel, err := b.ds.UserFollowRelation().Get(ctx, &model.UserFollowed{UserID: userID, ToUserID: v.UserId})
		if err != nil {
			return &rsp, err
		}
		rspList[i].IsFollow = rel.IsFollow
	}
	rsp.UserList = rspList
	return &rsp, nil
}

func (b *BFollowRelation) FriendList(ctx context.Context, userID uint) (*api.FriendListRsp, error) {
	//根据用户id查询关注列表
	//得到所有关注的用户
	l, err := b.ds.UserFollowRelation().FriendList(ctx, userID)
	if err != nil {
		return &api.FriendListRsp{}, err
	}
	//将用户转化
	rspList := make([]api.UserDetailRespond, len(l))
	for i, v := range l {
		rspList[i] = *tools.UserToRsp(&v)
		rspList[i].IsFollow = true
	}
	return &api.FriendListRsp{UserList: rspList}, nil
}
