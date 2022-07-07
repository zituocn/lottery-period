package router

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/lottery-period/handler"
)

// APIRouter api router func
func APIRouter(r *gow.Engine) {
	r.NoRoute(handler.ErrorHandler)
	v1 := r.Group("/api")
	{
		v1.GET("/time", handler.GetOneByTime)     //根据时间取配置
		v1.GET("/period", handler.GetOneByPeriod) //根据期数取配置
		v1.GET("/list", handler.GetList)          //包括当前期数的小列表，即最近N期
		v1.GET("/today", handler.Today)           //今天开奖信息
		v1.GET("/game", handler.Game)             //彩种配置输出
		v1.GET("/year", handler.Year)             //年份配置输出
		v1.GET("/game/yaml", handler.GameYaml)    //彩种配置yaml输出
		v1.GET("/year/yaml", handler.YearYaml)    //年份配置yaml输出
	}
}
