package gok

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	// "time"
	// "fmt"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		cc := NewChart("rb", "rb", M5, ParseTime("2016-01-01 00:00"), ParseTime("2020-01-01 00:00"))
		ss := NewChart("s000002", "a", Day, ParseTime("2016-01-01 00:00"), ParseTime("2020-01-01 00:00"))
		So(len(cc.Sticks()), ShouldBeGreaterThan, 0)
		So(ss.MACD(12, 26, 9).Diff[ss.GetIndexByDateString("2017-05-18")], ShouldAlmostEqual, -0.379, 0.001)

		Convey("When the integer is incremented", func() {
			So(2, ShouldEqual, 2)

			Convey("The value should be greater by one", func() {
				So(2, ShouldEqual, 2)
			})
		})
	})
}
