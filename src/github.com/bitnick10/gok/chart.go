package gok

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"time"
)

type Period int

const (
	M1 Period = iota
	M5
	Day
)

type Chart struct {
	id      string
	name    string
	period  Period
	sticks  []Candlestick
	mas     []*MA
	kdjs    []*KDJ
	macds   []*MACD
	obv     *OBV
	obvsmas []*OBVSMA
}

func GetTimeFormat(period Period) string {
	if period == Day {
		return "2006-01-02"
	} else if period == M5 || period == M1 {
		return "2006-01-02 15:04"
	} else {
		panic("GetTimeFormat")
	}
}

func ParseTime(value string) time.Time {
	if len(value) == 10 {
		t, _ := time.Parse("2006-01-02", value)
		return t
	}
	if len(value) == 16 {
		t, _ := time.Parse("2006-01-02 15:04", value)
		return t
	}
	fmt.Println(value)
	panic("ParseTime")
}

func (self *Chart) Id() string {
	return self.id
}

func (self *Chart) Name() string {
	return self.name
}

func (self *Chart) Sticks() []Candlestick {
	return self.sticks
}

func (self *Chart) Hn(data_i, n int) float64 {
	max := self.Sticks()[data_i].High
	begin := data_i - n + 1
	if begin < 0 {
		begin = 0
	}
	for i := begin; i <= data_i; i++ {
		if self.Sticks()[i].High > max {
			max = self.Sticks()[i].High
		}
	}
	return max
}

func (self *Chart) Ln(data_i, n int) float64 {
	min := self.Sticks()[data_i].Low
	begin := data_i - n + 1
	if begin < 0 {
		begin = 0
	}
	for i := begin; i <= data_i; i++ {
		if self.Sticks()[i].Low < min {
			min = self.Sticks()[i].Low
		}
	}
	return min
}

func (self *Chart) MA(n int) *MA {
	for i := 0; i < len(self.mas); i++ {
		if self.mas[i].N == n {
			return self.mas[i]
		}
	}
	ma := NewMA(self, n)
	self.mas = append(self.mas, ma)
	return self.mas[len(self.mas)-1]
}

func (self *Chart) MACD(shortN, longN, deaN int) *MACD {
	for i := 0; i < len(self.macds); i++ {
		if self.macds[i].ShortN == shortN && self.macds[i].LongN == longN && self.macds[i].DeaN == deaN {
			return self.macds[i]
		}
	}
	macd := NewMACD(self, shortN, longN, deaN)
	self.macds = append(self.macds, macd)
	return self.macds[len(self.macds)-1]
}

func (self *Chart) KDJ(n, sn1, sn2 int) *KDJ {
	for i := 0; i < len(self.kdjs); i++ {
		if self.kdjs[i].N == n && self.kdjs[i].Sn1 == sn1 && self.kdjs[i].Sn2 == sn2 {
			return self.kdjs[i]
		}
	}
	newOne := NewKDJ(self, n, sn1, sn2)
	self.kdjs = append(self.kdjs, newOne)
	return self.kdjs[len(self.kdjs)-1]
}

func (self *Chart) OBV() *OBV {
	if self.obv == nil {
		self.obv = NewOBV(self)
	}
	return self.obv
}

func (self *Chart) OBVSMA(n int) *OBVSMA {
	for i := 0; i < len(self.obvsmas); i++ {
		if self.obvsmas[i].N == n {
			return self.obvsmas[i]
		}
	}
	newOne := NewOBVSMA(self, n)
	self.obvsmas = append(self.obvsmas, newOne)
	return self.obvsmas[len(self.obvsmas)-1]
}

func (self *Chart) GetIndexByDate(date time.Time) int {
	for i := 0; i < len(self.sticks); i++ {
		// fmt.Println(self.sticks[i].CloseTime.Format("2006-01-02"))
		// fmt.Println(date.Format("2006-01-02"))
		if self.sticks[i].CloseTime.Equal(date) {
			return i
		}
	}
	return -1
}

func (self *Chart) GetIndexByDateString(dateString string) int {
	date, _ := time.Parse(GetTimeFormat(self.period), dateString)
	for i := 0; i < len(self.sticks); i++ {
		if self.sticks[i].CloseTime.Equal(date) {
			return i
		}
	}
	return -1
}

func (self *Chart) RawStochasticValues(n int) []float64 {
	rsvs := make([]float64, len(self.Sticks()))
	for i := 0; i < len(self.Sticks()); i++ {
		hn := self.Hn(i, n)
		ln := self.Ln(i, n)
		C := self.Sticks()[i].Close
		rsv := (C - ln) / (hn - ln) * 100
		if hn-ln == 0.0 {
			rsv = 0.0
		}
		if rsv < 0 {
			// print("rsv < 0 at ${i} ${tradeData[i].tradeTime} high:${hn} low:${ln} C:${C}");
			// assert(0 > 1);
			rsv = 0.0
		}
		rsvs[i] = rsv
	}
	return rsvs
}

func (self *Chart) GetVolumeBetween(beginDateString, endDateString string) float64 {
	volume := 0.0
	// beginDate, _ := time.Parse("2006-01-02 15:04", beginDateString)
	// endDate, _ := time.Parse("2006-01-02 15:04", endDateString)
	// for _, v := range self.Sticks() {
	// fmt.Println(v.CloseTime.Format("2006-01-02 15:04"))
	// }
	begin_i := self.GetIndexByDateString(beginDateString)
	end_i := self.GetIndexByDateString(endDateString)
	if begin_i == -1 {
		panic("fuck begin_i == -1")
		fmt.Println("...")
	}
	if end_i == -1 {
		panic("fuck end_i == -1")
		fmt.Println("...")
	}
	// fmt.Println(begin_i)
	// fmt.Println(end_i)
	for i := begin_i; i <= end_i; i++ {
		// fmt.Println(self.Sticks()[i].Volume)
		volume += self.Sticks()[i].Volume
	}
	return volume
}

func NewChart(id, name string, period Period, beginTime, endTime time.Time) *Chart {
	cc := &Chart{}
	cc.id = id
	cc.name = name
	cc.period = period

	dbname := "bourgogne"
	if len(id) == 7 {
		// fmt.Println(id)
		dbname = "stock"
	}

	periodString := "day"
	if period == M5 {
		periodString = "m5"
	}
	if period == M1 {
		periodString = "m1"
	}

	query := fmt.Sprintf("SELECT date,open,close,high,low,volume from %s.%s WHERE '%s' <= date AND date<= '%s'", periodString, id,
		beginTime.Format(GetTimeFormat(period)), endTime.Format(GetTimeFormat(period)))

	// fmt.Println(query)
	host := "192.168.0.170"
	// host := "127.0.0.1"
	port := "5432"
	password := "tangbei"

	cc.sticks = make([]Candlestick, 0)

	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, "postgres", password, dbname)

	// fmt.Println(connectString)
	// fmt.Println(connectString)
	// fmt.Println(time.Now().String())
	db, err := sql.Open("postgres", connectString)
	// fmt.Println(time.Now().String())
	defer db.Close()
	if err != nil {
		fmt.Println("postgres conncet err")
	}
	// fmt.Println("111111111111111111111111111")
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("postgres query err")
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		stick := Candlestick{}
		var date string
		err := rows.Scan(&date, &stick.Open, &stick.Close, &stick.High, &stick.Low, &stick.Volume)
		// fmt.Println(date)

		tm, _ := time.Parse(GetTimeFormat(period), date)
		// fmt.Println(tm.Format("2006-01-02"))
		if err != nil {
			fmt.Println("postgres Next err")
			fmt.Println(err)
		}
		stick.CloseTime = tm
		stick.OpenTime = tm
		if period == M5 {
			stick.OpenTime = tm.Add(time.Minute * -5)
		}
		// fmt.Println(stick.CloseTime.Format("2006-01-02"))
		cc.sticks = append(cc.sticks, stick)
	}

	for i := 1; i < len(cc.sticks); i++ {
		cc.sticks[i].Prev = &cc.sticks[i-1]
	}
	// fmt.Println(cc.sticks[len(cc.sticks)-2].CloseTime.Format("2016-01-02"))
	return cc
}

type JSONCandlestick struct {
	Date                   string
	Open, Close, High, Low float64
	Volume                 float64
}

func NewChartFromRedis(id, name string, period Period, beginTime, endTime time.Time) *Chart {
	sticks := make([]JSONCandlestick, 0)
	ret := &Chart{id: id, name: name, period: period}
	// ret.sticks = make([]Candlestick, 0)
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	prefix := "day"
	if period == M5 {
		prefix = "m5"
	}
	key := fmt.Sprintf("%s.%s", prefix, id)
	// fmt.Println("key:", key)
	bb, _ := redis.Bytes(conn.Do("GET", key)) // redis.Bytes(reply, err) conn.Do("GET", "m5.000002")
	// str, _ := redis.Sidtring(conn.Do("GET", key))
	// fmt.Println(str)
	// fmt.Println(bb)
	err = json.Unmarshal(bb, &sticks)
	// fmt.Println("len(sticks)", len(sticks))
	if err != nil {
		fmt.Println("json.Unmarshal err", id)
		fmt.Println(err)
	}
	ret.sticks = make([]Candlestick, len(sticks))
	for i := 0; i < len(sticks); i++ {
		// c := Candlestick{}
		ret.sticks[i].CloseTime = ParseTime(sticks[i].Date)
		ret.sticks[i].Open = sticks[i].Open
		ret.sticks[i].Close = sticks[i].Close
		ret.sticks[i].High = sticks[i].High
		ret.sticks[i].Low = sticks[i].Low
		ret.sticks[i].Volume = sticks[i].Volume
	}
	//fmt.Println(ret.sticks[0].Close)
	// fmt.Printf("%#v", ret.sticks)
	return ret
}
