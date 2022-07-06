package models

import (
	"io/ioutil"

	"github.com/zituocn/gow/lib/logy"
	"gopkg.in/yaml.v3"
)

type LotteryGameConfig struct {
	Des         string         `json:"desc" yaml:"desc"`
	LotteryGame []*LotteryGame `json:"lotterys" yaml:"lotterys"`
}

type LotteryYearConfig struct {
	Des   string         `json:"desc" yaml:"desc"`
	Years []*LotteryYear `json:"years" yaml:"years"`
}

// LoadLotteryGameConfig 读取yaml配置文件
func LoadLotteryGameConfig(filepath string) *LotteryGameConfig {
	b, err := loadFile(filepath)
	if err != nil {
		logy.Panicf("读取配置文件错误: %s", err.Error())
	}
	lotteryGameConfig := new(LotteryGameConfig)
	err = yaml.Unmarshal(b, &lotteryGameConfig)
	if err != nil {
		logy.Panicf("解析配置文件 错误: %s", err.Error())
	}

	return lotteryGameConfig
}

// LoadLotteryYearConfig 读取yaml本文件
func LoadLotteryYearConfig(filepath string) *LotteryYearConfig {
	b, err := loadFile(filepath)
	if err != nil {
		logy.Panicf("读取配置文件错误: %s", err.Error())
	}
	lotteryYearConfig := new(LotteryYearConfig)
	err = yaml.Unmarshal(b, &lotteryYearConfig)
	if err != nil {
		logy.Panicf("解析配置文件错误: %s", err.Error())
	}

	return lotteryYearConfig
}

func loadFile(filepath string) (b []byte, err error) {
	b, err = ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	return
}
