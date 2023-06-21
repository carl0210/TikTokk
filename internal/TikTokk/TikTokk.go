package TikTokk

import (
	"TikTokk/internal/TikTokk/controller"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := route()
	engine.Run(":6666")
}

func route() *gin.Engine {

	e := gin.Default()
	uc := controller.NewCUser(store.S)
	vc := controller.NewCVideo(store.S)
	commC := controller.NewCComment(store.S)
	relFavoriteC := controller.NewCRelFavorite(store.S)
	relFollowC := controller.NewCRelation(store.S)
	messageC := controller.NewCMessage(store.S)
	e.Static("/asset", "./asset/video")
	e.POST("/uploads/", controller.NewFile(store.S).Uploads)
	g := e.Group("/douyin")
	{
		// /user
		userG := g.Group("/user/")
		{
			userG.POST("/register/", uc.Register)
			userG.POST("/login/", uc.Login)
			userG.GET("/", middleware.AuthnByQuery(), uc.GetDetail)
		}
		//feed
		g.GET("/feed/", vc.Feed)
		//publish
		publishG := g.Group("/publish")
		{
			publishG.POST("/action/", middleware.AuthnByBody(), vc.PublishAction)
			publishG.GET("/list/", middleware.AuthnByQuery(), vc.PublishList)
		}
		//favorite
		favoriteG := g.Group("/favorite")
		{
			favoriteG.GET("/list/", middleware.AuthnByQuery(), relFavoriteC.FollowList)
			favoriteG.POST("/action/", middleware.AuthnByQuery(), relFavoriteC.FollowAction)
		}
		//comment
		commentG := g.Group("/comment")
		{
			commentG.POST("/action/", middleware.AuthnByQuery(), commC.FollowAction)
			commentG.GET("/list/", middleware.AuthnByQuery(), commC.FollowList)
		}
		//relation
		relationG := g.Group("/relation")
		{
			relationG.POST("/action/", middleware.AuthnByQuery(), relFollowC.FollowAction)
			relationG.GET("/follow/list/", middleware.AuthnByQuery(), relFollowC.FollowList)
			relationG.GET("/follower/list/", middleware.AuthnByQuery(), relFollowC.FollowerList)
			relationG.GET("/friend/list/", middleware.AuthnByQuery(), relFollowC.FriendListList)
		}
		messageG := g.Group("/message")
		{
			messageG.POST("/action/", middleware.AuthnByQuery(), messageC.Action)
			messageG.GET("/chat/", middleware.AuthnByQuery(), messageC.Chat)
		}
	}

	return e
}
