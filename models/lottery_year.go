package models

import "fmt"

// LotteryYear 年份配置
type LotteryYear struct {
	ID     int       `json:"id" yaml:"id"`
	Year   int       `json:"year" yaml:"year"`
	SEList []*SETime `json:"selist" yaml:"selist"`
}

var (
	yearPath = "./yaml/lottery_year.yaml"
)

// GetLotteryYear 获取某一年的配置信息
func (m *LotteryYear) GetLotteryYear(year int) (*LotteryYear, error) {
	lotters := LoadLotteryYearConfig(yearPath)
	if lotters != nil && len(lotters.Years) > 0 {
		for _, item := range lotters.Years {
			if item.Year == year {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("未找到: %d 的年份配置信息", year)
}

func (m *LotteryYear) GetLotteryYearList() ([]*LotteryYear, error) {
	lotters := LoadLotteryYearConfig(yearPath)
	return lotters.Years, nil
}
