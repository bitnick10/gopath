package gok

// import (
// 	. "github.com/smartystreets/goconvey/convey"
// 	"testing"
// 	// "time"
// 	"fmt"
// )

// func TestSpec(t *testing.T) {

// 	// Only pass t into top-level Convey calls
// 	Convey("Given some integer with a starting value", t, func() {
// 		x := 1

// 		Convey("When the integer is incremented", func() {
// 			x++

// 			Convey("The value should be greater by one", func() {
// 				So(x, ShouldEqual, 2)
// 			})
// 		})
// 	})
// }

// func TestFutures1(t *testing.T) {
// 	index := NewFutures("index399001", "深证成指", M5, "2015-12-27")
// 	Convey("Given some integer with a starting value", t, func() {
// 		//So(index.OBV().Ovb[index.GetIndexByDateString("2017-05-05 14:55")], ShouldEqual, 16)
// 		// So(index.OBV().Ovb[index.GetIndexByDateString("2017-05-05 15:00")], ShouldEqual, 5)
// 		So(index.OBV().Ovb[index.GetIndexByDateString("2017-05-04 15:00")], ShouldEqual, 5)
// 		//So(index.OBV().Ovb[index.GetIndexByDateString("2017-05-05 14:55")], ShouldEqual, 16)
// 		fmt.Println(index.Sticks()[index.GetIndexByDateString("2017-05-05 15:00")].Volume)
// 	})
// }

// // func TestFutures(t *testing.T) {
// // 	ff := NewFutures("rb", "rebar", Day, "2017-01-01")
// // 	// Only pass t into top-level Convey calls
// // 	Convey("Given some integer with a starting value", t, func() {
// // 		So(ff.id, ShouldEqual, "rb")
// // 		tm, _ := time.Parse("2006-01-02", "2017-04-11")
// // 		ff.GetIndexByDate(tm)
// // 		So(ff.Sticks()[ff.GetIndexByDate(tm)].Close, ShouldEqual, 2982)
// // 		So(ff.Sticks()[ff.GetIndexByDateString("2017-04-12")].Amplitude(), ShouldAlmostEqual, 3.49/100, 0.01/100)

// // 		So(ff.KDJ(9, 3, 3).K[ff.GetIndexByDateString("2017-04-17")], ShouldAlmostEqual, 11.26, 0.01)
// // 		So(ff.KDJ(9, 3, 3).D[ff.GetIndexByDateString("2017-04-17")], ShouldAlmostEqual, 14.34, 0.01)

// // 		So(ff.MA(5).Ma[ff.GetIndexByDateString("2017-04-11")], ShouldAlmostEqual, 3081, 0.5)
// // 		// So(ff.sticks[len(ff.sticks)-1].tradeTime.Format("2006-01-02"), ShouldEqual, "2017-014-17")

// // 		Convey("5 minutes test", func() {
// // 			f5 := NewFutures("rb", "rebar", M5, "2017-01-01")
// // 			// So(f5.Sticks()[len(f5.Sticks())-1].Volume, ShouldEqual, 104168)
// // 			// f5.Sticks()[len(f5.Sticks())-1].Volume
// // 			volume := f5.GetVolumeBetween("2017-04-16 21:05", "2017-04-16 23:00")
// // 			So(volume/1000/1000, ShouldAlmostEqual, 1.83, 0.01)
// // 			Convey("The value should be greater by one", func() {
// // 				// So(x, ShouldEqual, 2)
// // 			})
// // 		})
// // 	})
// // }
