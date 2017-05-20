package gok

// import (
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/lib/pq"
// 	"time"
// )

// type Futures struct {
// 	Chart
// }

// func NewFutures(id, name string, period Period, beginDate string) *Futures {
// 	ff := &Futures{}
// 	ff.id = id
// 	ff.name = name
// 	ff.period = period

// 	dbname := "bourgogne"
// 	periodString := "day"
// 	if period == M5 {
// 		periodString = "m5"
// 	}

// 	query := fmt.Sprintf("SELECT date,open,close,high,low,volume from %s.%s WHERE date > '%s' ", periodString, id, beginDate)

// 	// fmt.Println(query)
// 	host := "192.168.0.170"
// 	port := "5432"
// 	password := "tangbei"

// 	ff.sticks = make([]Candlestick, 0)

// 	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, "postgres", password, dbname)

// 	// fmt.Println(connectString)
// 	db, err := sql.Open("postgres", connectString)
// 	defer db.Close()
// 	if err != nil {
// 		fmt.Println("postgres conncet err")
// 	}
// 	// fmt.Println("111111111111111111111111111")
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		fmt.Println("postgres query err")
// 		fmt.Println(err)
// 	}

// 	for rows.Next() {
// 		stick := Candlestick{}
// 		var date string
// 		err := rows.Scan(&date, &stick.Open, &stick.Close, &stick.High, &stick.Low, &stick.Volume)
// 		// fmt.Println(date)

// 		tm, _ := time.Parse(GetTimeFormat(period), date)
// 		// fmt.Println(tm.Format("2006-01-02"))
// 		if err != nil {
// 			fmt.Println("postgres Next err")
// 			fmt.Println(err)
// 		}
// 		stick.CloseTime = tm
// 		stick.OpenTime = tm
// 		if period == M5 {
// 			stick.OpenTime = tm.Add(time.Minute * -5)
// 		}
// 		// fmt.Println(stick.CloseTime.Format("2006-01-02"))
// 		ff.sticks = append(ff.sticks, stick)
// 	}
// 	for i := 1; i < len(ff.sticks); i++ {
// 		ff.sticks[i].Prev = &ff.sticks[i-1]
// 	}
// 	// fmt.Println(ff.sticks[len(ff.sticks)-2].CloseTime.Format("2016-01-02"))
// 	return ff
// }
