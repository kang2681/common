package silence

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type CronSchedule struct {
	Spec     string      // 原始表达式
	Trade    int         // 0 不判断  1 交易日  2 非交易日
	Clocks   []CronClock // 每天的 时分秒，支持多个，必填
	Schedule cron.Schedule
	TradeFn  func() (start, end time.Time, tradeList []time.Time)
}

type CronClock struct {
	StartSeconds int
	EndSeconds   int
}

func (c *CronSchedule) GetSpec() string {
	return c.Spec
}

type NextTime struct {
	StartAt time.Time
	EndAt   time.Time
}

func (n NextTime) IsZero() bool {
	return n.StartAt.IsZero() || n.EndAt.IsZero()
}

func (c *CronSchedule) Next(t time.Time) NextTime {
	rs := c.NextN(t, 1)
	if len(rs) == 0 {
		return NextTime{}
	}
	return rs[0]
}

func (c *CronSchedule) NextN(t time.Time, n int) []NextTime {
	rs := make([]NextTime, 0, n)
	// calculate time
	ct := c.timeZeroClock(t).Add(-1)
	for ct.Year()-t.Year() < 10 {
		// Schedule 的 next 生成的是 传入时间以后的时间
		nextTime := c.Schedule.Next(ct)
		if nextTime.IsZero() {
			break
		}
		// 交易日和非交易日判断
		if c.Trade != 0 {
			switch c.Trade {
			case 1: // 交易日
				flag, err := c.IsTradeday(nextTime)
				if err != nil {
					logrus.Errorf("trade day check error %s", err.Error())
					logrus.Errorf("交易日检查失败， 交易日获取失败，使用 周一到周五 兜底")
					weekday := nextTime.Weekday()
					if weekday == time.Saturday || weekday == time.Sunday {
						ct = nextTime
						continue
					}
				}
				if !flag {
					ct = nextTime
					continue
				}
			case 2: // 非交易日
				flag, err := c.IsTradeday(nextTime)
				if err != nil {
					logrus.Errorf("trade day check error %s", err.Error())
					logrus.Errorf("交易日检查失败， 交易日获取失败，使用 周六、周天 兜底")
					weekday := nextTime.Weekday()
					if weekday != time.Saturday && weekday != time.Sunday {
						ct = nextTime
						continue
					}
				}
				if flag {
					ct = nextTime
					continue
				}
			}
		}
		// logrus.Infof("ct:%+v,nextTime:%+v", ct.Format("2006-01-02"), nextTime)
		dayClocks := make([]CronClock, 0, len(c.Clocks))
		if c.isSame(t, nextTime, "20060102") {
			hour, minute, second := t.Clock()
			seconds := hour*3600 + minute*60 + second
			for _, v := range c.Clocks {
				if seconds < v.EndSeconds {
					dayClocks = append(dayClocks, v)
				}
			}
		} else {
			dayClocks = append(dayClocks, c.Clocks...)
		}
		// 排序
		sort.Slice(dayClocks, func(i, j int) bool { return dayClocks[i].StartSeconds < dayClocks[j].StartSeconds })
		size := n - len(rs)
		length := len(dayClocks)

		logrus.Infof("dayClocks:%+v,size:%+v,length:%+v", dayClocks, size, length)
		for i := 0; i < size && i < length; i++ {
			value := dayClocks[i]
			tmp := NextTime{
				StartAt: nextTime.Add(time.Duration(value.StartSeconds) * time.Second),
				EndAt:   nextTime.Add(time.Duration(value.EndSeconds) * time.Second),
			}
			rs = append(rs, tmp)
		}
		logrus.Infof("%+v", rs)
		if len(rs) >= n {
			break
		}
		ct = nextTime
	}
	return rs
}

func (c *CronSchedule) IsTradeday(t time.Time) (bool, error) {
	startTime, endTime, tradeSlice := c.TradeFn()
	ct := c.timeZeroClock(t)
	//  当前时间在 交易日的范围外
	if ct.Before(startTime) || ct.After(endTime) {
		return false, errors.Errorf("当前时间:%+v 在交易日范围 :%+v ~ %+v 以外", ct.Format("2006-01-02"), startTime.Format("2006-01-02"), endTime.Format("2006-01-02"))
	}
	for _, v := range tradeSlice {
		if c.isSame(ct, v, "20060102") {
			return true, nil
		}
	}
	return false, nil
}

func (c *CronSchedule) timeZeroClock(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func (c *CronSchedule) isSame(t1, t2 time.Time, format string) bool {
	return t1.Format(format) == t2.Format(format)
}

// NewCronSchedule
// expr crontab定时 "0 00:00:00-00:08:00 * * *"
// fn 交易日获取方法
func NewCronSchedule(expr string, tradeFn func() (time.Time, time.Time, []time.Time)) (*CronSchedule, error) {
	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return nil, errors.Errorf("expr error")
	}
	cp := cronParam{
		Trade: fields[0],
		Hms:   strings.Split(fields[1], ","),
		Day:   fields[2],
		Month: fields[3],
		Week:  fields[4],
	}
	cs := CronSchedule{}
	//  trade 参数验证
	switch cp.Trade {
	case "0":
		cs.Trade = 0
	case "1":
		cs.Trade = 1
		if tradeFn == nil {
			return nil, errors.Errorf(" trade day check open,  trade function empty")
		}
		cs.TradeFn = tradeFn
	case "2":
		cs.Trade = 2
		if tradeFn == nil {
			return nil, errors.Errorf(" trade day check open,  trade function empty")
		}
		cs.TradeFn = tradeFn
	default:
		return nil, errors.Errorf("trade not support value %d", cp.Trade)
	}
	clocks, err := parse_clock(cp.Hms)
	if err != nil {
		return nil, errors.Wrap(err, "parse_clock")
	}
	cs.Clocks = clocks
	p := cron.NewParser(cron.Dom | cron.Month | cron.Dow)
	s, err := p.Parse(fmt.Sprintf("%s %s %s", cp.Day, cp.Month, cp.Week))
	if err != nil {
		return nil, errors.Wrapf(err, "cron parse error")
	}
	cs.Schedule = s
	return &cs, nil
}

func parse_clock(exprs []string) ([]CronClock, error) {
	if len(exprs) == 0 {
		return nil, errors.Errorf("HMS empty")
	}
	clocks := make([]CronClock, 0, len(exprs))
	for _, expr := range exprs {
		// hh:mm:ss-hh:mm:ss
		arr := strings.Split(expr, "-")
		if len(arr) != 2 {
			return nil, errors.Errorf("HMS %s format error", expr)
		}
		startTime, err := time.Parse("15:04:05", arr[0])
		if err != nil {
			return nil, errors.Wrapf(err, "HMS %s parse error", arr[0])
		}
		endTime, err := time.Parse("15:04:05", arr[1])
		if err != nil {
			return nil, errors.Wrapf(err, "HMS %s parse error", arr[1])
		}
		startHour, startMinue, startSecond := startTime.Clock()
		endHour, endMinue, endSecond := endTime.Clock()

		startSeconds := startHour*3600 + startMinue*60 + startSecond
		endSeconds := endHour*3600 + endMinue*60 + endSecond
		// 开始时间和结束时间判断
		if startSeconds < 0 {
			return nil, errors.Wrapf(err, "start HMS error:%s  parse seconds %d", arr[0], startSeconds)
		}
		if endSeconds >= 86400 {
			return nil, errors.Wrapf(err, "end HMS error:%s  parse seconds %d", arr[1], endSeconds)
		}
		if startSeconds >= endSeconds {
			return nil, errors.Wrapf(err, "start HMS must less than end  end HMS, start HMS： end HMS:%s", arr[0], arr[1])
		}

		c := CronClock{
			StartSeconds: startSeconds,
			EndSeconds:   endSeconds,
		}
		clocks = append(clocks, c)
	}
	sort.Slice(clocks, func(i, j int) bool { return clocks[i].StartSeconds < clocks[j].StartSeconds })
	length := len(clocks) - 1
	for i := 0; i < length; i++ {
		if clocks[i].EndSeconds > clocks[i+1].StartSeconds {
			return nil, errors.Errorf("HMS has overlapping")
		}
	}
	return clocks, nil
}

type cronParam struct {
	// 0 不判断
	// 1 交易日
	// 2 非交易日
	Trade string
	// hh:mm:ss-hh:mm:ss
	Hms []string
	// ? 未指定
	// * 每一天
	// 1-6 1~6号
	// 3/2 从3号开始,间隔2天
	// 3-6/2 从3号到6号,间隔2天
	// 1,2,3 指定时间 1,2,3 号
	Day string
	// ? 未指定
	// * 每个月
	// 1-2 1月到2月
	// 4/2 从4月开始,间隔2月
	// 2-4/2 从2月到4月，间隔2月
	// 1,2,3 指定时间 1,2,3月
	Month string // month
	// ? 不指定
	// * 每周
	// 1-2 周一到周二
	// 4/2 周四开始,间隔2天
	// 2-6/2 从2号到6号,间隔2天
	// 1,2,3 指定时间 周1,2,3
	Week string // week 周天是0
	// ? 不指定
	// * 每周
	// 2021-2022 2021~2022年
	// 2021/2 2021年，间隔2年
	// 2021-2026/2 2021~2026年，间隔2年
	// 1,2,3 指定时间 周1,2,3
	// Year string `json:"year"`
}
