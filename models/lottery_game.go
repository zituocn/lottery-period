package models

import "fmt"

// LotteryGame 彩种配置
type LotteryGame struct {
	Gamecode int       `json:"gamecode" yaml:"gamecode"` //彩种ID
	Oid      int       `json:"oid" yaml:"oid"`           //1福彩/2体彩
	Typeid   int       `json:"typeid" yaml:"typeid"`     //类型
	Name     string    `json:"name" ymal:"name"`         //中文名
	Sname    string    `json:"sname" yaml:"sname"`       //中文名
	Ename    string    `json:"ename" yaml:"ename"`       //英文名或拼音
	Opendate string    `json:"opendate" yaml:"opendate"` //开奖日期
	Opentime string    `json:"opentime" ymal:"opentime"` //开奖时间
	State    int       `json:"state" ymal:"state"`       //状态
	Islocal  int       `json:"islocal" ymal:"islocal"`   //地方彩种
	Province string    `json:"province" ymal:"province"` //地方彩的所在地
	Formatqi string    `json:"formatqi" ymal:"formatqi"` //期数格式
	SEList   []*SETime `json:"selist" yaml:"selist"`     //彩种的时间配置
}

//SETime 时间控制
type SETime struct {
	STime string `json:"stime" yaml:"stime"` //开始时间
	ETime string `json:"etime" yaml:"etime"` //结束时间
}

var (
	gamePath = "./yaml/lottery_game.yaml"
)

// GetLotteryGame 获取某个game的配置信息
func (m *LotteryGame) GetLotteryGame(gamecode int) (*LotteryGame, error) {
	lotters := LoadLotteryGameConfig(gamePath)
	if lotters != nil && len(lotters.LotteryGame) > 0 {
		for _, item := range lotters.LotteryGame {
			if item.Gamecode == gamecode {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("未找到gamecode: %d 的配置信息", gamecode)
}

// GetLotteryGameList yaml转json配置
func (m *LotteryGame) GetLotteryGameList() ([]*LotteryGame, error) {
	data := LoadLotteryGameConfig(gamePath)
	return data.LotteryGame, nil
}
