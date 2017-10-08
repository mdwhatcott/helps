package measures

import (
	"bytes"
	"testing"

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
	this.So(func() { Length(1) }, should.Panic)

	var a [1]int
	this.So(Length(a), should.Equal, 1)
	this.So(Length(""), should.Equal, 0)
	this.So(Length(make([]int, 0)), should.Equal, 0)
	this.So(Length(make(chan int)), should.Equal, 0)
	this.So(Length(make(map[int]int)), should.Equal, 0)
	this.So(Length(new(bytes.Buffer)), should.Equal, 0)
}
