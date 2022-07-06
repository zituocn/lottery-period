package handler

import (
	"time"

	"github.com/zituocn/gow"
	"github.com/zituocn/lottery-period/models"
	"github.com/zituocn/lottery-period/service"
)

const (
	format = "2006-01-02 15:04:05"
)

// GetOneByTime 根据时间取配置
func GetOneByTime(c *gow.Context) {
	gamecode, _ := c.GetInt("gamecode", 0)
	ts, _ := c.GetInt64("ts", 0)
	typeid, _ := c.GetInt("typeid", 0)

	if gamecode < 1 {
		JSON(c, 1, "[gamecode]不能为0")
		return
	}

	if ts < 1 {
		ts = time.Now().Unix()
	}

	result, err := new(service.Low).GetOneByTime(gamecode, ts, typeid)
	if err != nil {
		JSON(c, 1, err.Error())
		return
	}
	JSON(c, convertItem(result))

}

// GetOneByPeriod 根据期数取配置
func GetOneByPeriod(c *gow.Context) {
	gamecode, _ := c.GetInt("gamecode", 0)
	period, _ := c.GetInt64("period", 0)

	if gamecode < 1 || period < 1 {
		JSON(c, 1, "[gamecode] 和 [period] 不能为0")
		return
	}

	result, err := new(service.Low).GetOneByPeriod(gamecode, period)

	if err != nil {
		JSON(c, 1, err.Error())
		return
	}

	JSON(c, convertItem(result))
}

func convertItem(item *models.Result) *models.ViewResult {
	if item != nil {
		return &models.ViewResult{
			Period: item.Period,
			At:     time.Unix(item.At, 0).Format(format),
			Et:     time.Unix(item.Et, 0).Format(format),
			Bt:     time.Unix(item.Bt, 0).Format(format),
		}
	}

	return nil
}

func Today(c *gow.Context) {

}

func Game(c *gow.Context) {
	data, _ := new(models.LotteryGame).GetLotteryGameList()
	JSON(c, data)
}

func GameYaml(c *gow.Context) {
	data, _ := new(models.LotteryGame).GetLotteryGameList()
	c.YAML(data)
}

func Year(c *gow.Context) {
	data, _ := new(models.LotteryYear).GetLotteryYearList()
	JSON(c, data)
}

func YearYaml(c *gow.Context) {
	data, _ := new(models.LotteryYear).GetLotteryYearList()
	c.YAML(data)
}
