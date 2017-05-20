package gok

import (
	"database/sql"
	"fmt"
	// "github.com/jeffail/tunny"
	_ "github.com/lib/pq"
	// "runtime"
	"sync"
	// "time"
)

type StockMarket struct {
	stocks []*Chart
}

func (self *StockMarket) Stocks() []*Chart {
	return self.stocks
}

func (self *StockMarket) GetStockById(id string) *Chart {
	for _, s := range self.stocks {
		if s.Id() == id {
			return s
		}
	}
	return nil
}

func NewStockMarket(period Period) *StockMarket {
	ret := &StockMarket{}
	ret.stocks = make([]*Chart, 0)
	dbname := "stock"

	query := fmt.Sprintf("SELECT code,name from basics order by code")

	// fmt.Println(query)
	host := "192.168.0.170"
	// host := "127.0.0.1"
	port := "5432"
	password := "tangbei"

	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, "postgres", password, dbname)

	// fmt.Println(connectString)
	// fmt.Println(connectString)
	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		fmt.Println("postgres conncet err")
	}
	// fmt.Println("111111111111111111111111111")
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("postgres query err")
		fmt.Println(err)
	}

	var mutex = &sync.Mutex{}

	// exampleChannel := make(chan int)
	// runtime.GOMAXPROCS(10 + 1)
	// pool, _ := tunny.CreatePoolGeneric(16).Open()
	// defer pool.Close()
	// err = pool.SendWork(func() {
	// 	 Do your hard work here, usual rules of closures apply here,
	// 	 * so you can return values like so:

	// 	// exampleChannel <- 10
	// })

	if err != nil {
		// You done goofed
	}

	for rows.Next() {
		var code string
		var name string
		err := rows.Scan(&code, &name)

		if err != nil {
			fmt.Println("postgres Next err")
			fmt.Println(err)
		}
		code = "" + code
		if code[0] == '3' {
			continue
		}
		// fmt.Println(code)
		var chart = NewChartFromRedis(code, name, period, ParseTime("2017-01-01"), ParseTime("2090-01-01"))
		mutex.Lock()
		ret.stocks = append(ret.stocks, chart)
		mutex.Unlock()
		// fmt.Println(code)
		// go func() {
		// 	_, err = pool.SendWork(func() {
		// 		// fmt.Println(code)
		// 		// fmt.Println(time.Now().String())
		// 		var chart = NewChartFromRedis(code, name, period, ParseTime("2016-01-01"), ParseTime("2090-01-01"))
		// 		mutex.Lock()
		// 		ret.stocks = append(ret.stocks, chart)
		// 		mutex.Unlock()
		// 		// fmt.Println(time.Now().String())
		// 	})
		// }()

		if err != nil {
			fmt.Println("eeeeeeeeeeeeee")
			// You done goofed
		}
		// go func() {
		// 	var chart = NewChart(code, name, period, ParseTime("2017-01-01"), ParseTime("2090-01-01"))
		// 	mutex.Lock()
		// 	ret.stocks = append(ret.stocks, chart)
		// 	mutex.Unlock()
		// }()
		// ret.stocks = append(ret.stocks, NewChart(code, name, period, ParseTime("2017-01-01"), ParseTime("2090-01-01")))
	}
	return ret
}
