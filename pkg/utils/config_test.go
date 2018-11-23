package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadignore(t *testing.T) {
	Convey("Load Ignore Config", t, func() {
		ign := Loadignore()
		So(len(ign), ShouldEqual, 11)
	})
}
