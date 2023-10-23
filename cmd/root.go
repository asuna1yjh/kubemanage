/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"gin_demo/cmd/option"
	"gin_demo/logic"
	"gin_demo/pkg/snowflake"
	"gin_demo/pkg/validator"
	"gin_demo/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gin_demo",
	Short: "my_app",
	Long:  `my_app is a demo for web application`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 加载配置
		Option := option.NewOption()
		//settings.Init()
		// 2. 初始化日志
		err := Option.Complete()
		if err != nil {
			return
		}
		// 初始化雪花算法
		if err := snowflake.Init(Option.Config.StartTime, Option.Config.MachineID); err != nil {
			panic(err)
		}
		// 初始化logic
		useCase := logic.NewUserUseCase(Option.Factory)
		r := routes.Setup(Option, useCase)
		// 5.1 修改gin框架里面的Validator引擎属性，实现自定制
		err = validator.InitTrans("zh")
		if err != nil {
			return
		}

		// 6. 优雅的关机
		Run(r)
	},
}

// 优雅启动服务
func Run(r *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
