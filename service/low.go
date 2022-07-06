package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zituocn/lottery-period/models"
)

const (
	format = "2006-01-02 15:04:05"
)

type Low struct{}

// GetOneByPeriod 根据彩种和期数获取数据
//
func (m *Low) GetOneByPeriod(gamecode int, period int64) (result *models.Result, err error) {
	var (
		year int
	)

	qiStr := strconv.FormatInt(period, 10)

	lotteryGame, err := new(models.LotteryGame).GetLotteryGame(gamecode)
	if err != nil {
		return
	}

	if strings.Contains(lotteryGame.Formatqi, "YYYY") {
		year, err = strconv.Atoi(qiStr[0:4])
		if err != nil {
			err = fmt.Errorf("期数转年份出错 :%s", lotteryGame.Ename)
			return
		}
	} else {
		year, err = strconv.Atoi(qiStr[0:4])
		if err != nil {
			err = fmt.Errorf("期数转年份出错 :%s", lotteryGame.Ename)
			return
		}
	}
	data, err := m.GetCompleteList(lotteryGame, year)
	if err != nil {
		return
	}
	for _, item := range data {
		if item.Period == period {
			result = item
			return
		}
	}
	return
}

// GetOneByTime 获取某一期数据
//	根据彩种和时间
//	gamecode = 彩种
//	ts = 时间戳
//	typeid -1=上一期 0=当前期 1=下一期
func (m *Low) GetOneByTime(gamecode int, ts int64, typeid int) (result *models.Result, err error) {
	lotteryGame, err := new(models.LotteryGame).GetLotteryGame(gamecode)
	if err != nil {
		return
	}

	dt := time.Unix(ts, 0)
	year := dt.Year()

	data, err := m.GetCompleteList(lotteryGame, year)
	if err != nil {
		return
	}

	length := len(data)
	for i := 0; i < length; i++ {
		item := data[i]

		// 上一期
		if typeid == -1 && i > 0 {
			if item.At-ts >= 0 {
				result = data[i-1]
				return
			}
		}

		//当前期
		if typeid == 0 {
			if item.At-ts >= 0 {
				result = data[i]
				return
			}
		}

		//下一期
		if typeid == 1 && i < length-1 {
			if item.At-ts >= 0 {
				result = data[i+1]
				return
			}
		}
	}
	return
}

// GetCompleteList 完整的列表
//	当前年的上下两年数据
func (m *Low) GetCompleteList(lotteryGame *models.LotteryGame, arg ...int) (data []*models.Result, err error) {
	var (
		year int
	)
	if len(arg) > 0 {
		year = arg[0]
	} else {
		year = time.Now().Year()
	}

	last, _ := m.getBaseList(lotteryGame, year-1)
	current, _ := m.getBaseList(lotteryGame, year)
	next, _ := m.getBaseList(lotteryGame, year+1)

	data = append(data, last...)
	data = append(data, current...)
	data = append(data, next...)

	return
}

// getBaseList 拿到基本的数据列表
//	期数算法核心
func (m *Low) getBaseList(lotteryGame *models.LotteryGame, arg ...int) (data []*models.Result, err error) {
	var (
		year int
		qi   int64
		list = make([]*models.Result, 0)
		now  = time.Now()
	)

	if len(arg) > 0 {
		year = arg[0]
	} else {
		year = now.Year()
	}

	if lotteryGame == nil || lotteryGame.Gamecode == 0 {
		err = errors.New("参数错误,gamecode=0")
		return
	}

	lotteryYear, err := new(models.LotteryYear).GetLotteryYear(year)
	if err != nil {
		return
	}

	formats := strings.Split(strings.TrimSpace(lotteryGame.Formatqi), ",")
	if len(formats) != 2 {
		err = fmt.Errorf("formatqi 格式错误")
		return
	}
	openTime := lotteryGame.Opentime + ":00"
	stime, err := time.ParseInLocation(format, fmt.Sprintf("%d-01-01 %s", year, openTime), time.Local)
	if err != nil {
		err = fmt.Errorf("opentime格式错误 :%s", err.Error())
		return
	}

	etime, err := time.ParseInLocation(format, fmt.Sprintf("%d-12-31 %s", year, openTime), time.Local)
	if err != nil {
		err = fmt.Errorf("opentime 格式错误 :%s", err.Error())
		return
	}

	sumDay := int(etime.Sub(stime).Seconds() / 86400)
	tqi := 0

	for i := 0; i <= sumDay; i++ {
		temp := stime.AddDate(0, 0, i)
		//中间不开奖时间处理
		if !m.IsSETime(lotteryYear, lotteryGame, temp) {
			//星期处理
			if m.IsWeek(lotteryGame, temp) {
				tqi++
				item := new(models.Result)
				qiStr := fmt.Sprintf("%d"+formats[1], temp.Year(), tqi)
				qi, _ = strconv.ParseInt(qiStr, 10, 64)

				item.Period = qi
				item.Bt = temp.Unix()
				item.At = temp.Unix()
				item.Et = temp.Unix()

				list = append(list, item)
			}
		}
	}

	for _, item := range list {
		data = append(data, item)
	}

	return
}

// IsSETime 时间跨度问题
func (m *Low) IsSETime(lotteryYear *models.LotteryYear, lotteryGame *models.LotteryGame, dt time.Time) (flag bool) {
	yList := lotteryYear.SEList
	gList := lotteryGame.SEList
	for _, item := range yList {
		s, _ := time.ParseInLocation("2006-01-02", item.STime, time.Local)
		e, _ := time.ParseInLocation("2006-01-02", item.ETime, time.Local)
		//在这个时间区间内
		if dt.Sub(s).Seconds()/86400 >= 0 && e.Sub(dt).Seconds()/86400 >= 0 {
			return true
		}
	}

	for _, item := range gList {
		s, _ := time.ParseInLocation("2006-01-02", item.STime, time.Local)
		e, _ := time.ParseInLocation("2006-01-02", item.ETime, time.Local)
		//在这个时间区间内
		if dt.Sub(s).Seconds()/86400 >= 0 && e.Sub(dt).Seconds()/86400 >= 0 {
			return true
		}
	}

	return false
}

// IsWeek 星期
func (m *Low) IsWeek(lotteryGame *models.LotteryGame, dt time.Time) bool {
	weeks := strings.Split(lotteryGame.Opendate, ",")
	week := dt.Weekday().String()
	for _, item := range weeks {
		if item == week {
			return true
		}
	}
	return false
}
