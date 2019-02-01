package measures

import (
	"bytes"
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestLengthFixture(t *testing.T) {
	gunit.Run(new(LengthFixture), t)
}

type LengthFixture struct {
	*gunit.Fixture
}

func (this *LengthFixture) Test() {
	this.So(Length(1), should.Equal, -1)
	this.So(Length(struct{}{}), should.Equal, -1)
	this.So(Length(time.Now()), should.Equal, -1)

	var a [1]int
	var b []int
	this.So(Length(a), should.Equal, 1)
	this.So(Length(b), should.Equal, 0)
	this.So(Length(""), should.Equal, 0)
	this.So(Length(make([]int, 0)), should.Equal, 0)
	this.So(Length(make(chan int)), should.Equal, 0)
	this.So(Length(make(map[int]int)), should.Equal, 0)
	this.So(Length(new(bytes.Buffer)), should.Equal, 0)
}
