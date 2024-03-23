package TikTokk

import (
	"TikTokk/internal/TikTokk/controller"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/middleware"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	//创建路由
	e := route()
	//创建并运行server实例
	httpserver := &http.Server{Addr: "192.168.31.29:8080", Handler: e}
	go func() {
		if err := httpserver.ListenAndServe(); err != nil {
			return
		}
	}()
	//优雅关闭
	//监听信号,并使用管道阻塞
	//kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	//kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	cn := make(chan os.Signal, 1)
	signal.Notify(cn, syscall.SIGINT, syscall.SIGTERM)
	<-cn
	//收到信号后,进行清理工作
	log.Println("application will be shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//关闭http服务
	if err := httpserver.Shutdown(ctx); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Println(err.Error())
	}
}

func route() *gin.Engine {
	//gin.SetMode("release")
	e := gin.New()
	e.Use(gin.Recovery(), middleware.GinLogger())
	//e := gin.Default()
	//创建controller实例
	uc := controller.NewCUser(store.S)
	vc := controller.NewCVideo(store.S)
	commC := controller.NewCComment(store.S)
	relFavoriteC := controller.NewCRelFavorite(store.S)
	relFollowC := controller.NewCRelFollow(store.S)
	messageC := controller.NewCMessage(store.S)
	//路由
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
			favoriteG.GET("/list/", middleware.AuthnByQuery(), relFavoriteC.List)
			favoriteG.POST("/action/", middleware.AuthnByQuery(), relFavoriteC.Action)
		}
		//comment
		commentG := g.Group("/comment")
		{
			commentG.POST("/action/", middleware.AuthnByQuery(), commC.Action)
			commentG.GET("/list/", middleware.AuthnByQuery(), commC.List)
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
