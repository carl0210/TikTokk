package route

import (
	"TikTokk/controller"
	"TikTokk/midware"
	"TikTokk/store"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {

	e := gin.Default()
	uc := controller.NewCUser(store.S)
	vc := controller.NewCVideo(store.S)
	commC := controller.NewCComment(store.S)
	relFavoriteC := controller.NewCRelFavorite(store.S)
	relFollowC := controller.NewCRelation(store.S)
	messageC := controller.NewCMessage(store.S)
	e.Static("/publicSrc", "./publicSrc/video")
	e.POST("/uploads/", controller.NewFile(store.S).Uploads)
	g := e.Group("/douyin")
	{
		// /user
		userG := g.Group("/user/")
		{
			userG.POST("/register/", uc.Register)
			userG.POST("/login/", uc.Login)
			userG.GET("/", midware.AuthnByQuery(), uc.GetDetail)
		}
		//feed
		g.GET("/feed/", vc.Feed)
		//publish
		publishG := g.Group("/publish")
		{
			publishG.POST("/action/", midware.AuthnByBody(), vc.PublishAction)
			publishG.GET("/list/", midware.AuthnByQuery(), vc.PublishList)
		}
		//favorite
		favoriteG := g.Group("/favorite")
		{
			favoriteG.GET("/list/", midware.AuthnByQuery(), relFavoriteC.FollowList)
			favoriteG.POST("/action/", midware.AuthnByQuery(), relFavoriteC.FollowAction)
		}
		//comment
		commentG := g.Group("/comment")
		{
			commentG.POST("/action/", midware.AuthnByQuery(), commC.FollowAction)
			commentG.GET("/list/", midware.AuthnByQuery(), commC.FollowList)
		}
		//relation
		relationG := g.Group("/relation")
		{
			relationG.POST("/action/", midware.AuthnByQuery(), relFollowC.FollowAction)
			relationG.GET("/follow/list/", midware.AuthnByQuery(), relFollowC.FollowList)
			relationG.GET("/follower/list/", midware.AuthnByQuery(), relFollowC.FollowerList)
			relationG.GET("/friend/list/", midware.AuthnByQuery(), relFollowC.FriendListList)
		}
		messageG := g.Group("/message")
		{
			messageG.POST("/action/", midware.AuthnByQuery(), messageC.Action)
			messageG.GET("/chat/", midware.AuthnByQuery(), messageC.Chat)
		}
	}

	return e
}
