package helps

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestJSONFixture(t *testing.T) {
	gunit.Run(new(JSONFixture), t)
}

type JSONFixture struct {
	*gunit.Fixture
}

func (this *JSONFixture) TestFormatJSON() {
	formatted := FormatJSON([]byte(`[{"message": "hi"}]`))
	this.So(formatted, should.Equal, `[
  {
    "message": "hi"
  }
]`)
}

func (this *JSONFixture) TestFormatJSON_BadJSON() {
	var formatted string
	callback := func() { formatted = FormatJSON([]byte(`I can't haz json'`)) }
	this.So(callback, should.Panic)
	this.So(formatted, should.BeEmpty)
}

func (this *JSONFixture) TestFormatJSONSafe() {
	formatted, err := FormatJSONSafe([]byte(`[{"message": "hi"}]`))
	this.So(err, should.BeNil)
	this.So(formatted, should.Equal, `[
  {
    "message": "hi"
  }
]`)
}

func (this *JSONFixture) TestFormatJSONSafe_BadJSON() {
	formatted, err := FormatJSONSafe([]byte(`I can't haz json'`))
	this.So(err, should.NotBeNil)
	this.So(formatted, should.BeEmpty)
}

func (this *JSONFixture) TestDumpJSON() {
	dumped := DumpJSON(map[string]string{"Hello": "World"})
	this.So(dumped, should.Equal, `{
  "Hello": "World"
}`)
}

func (this *JSONFixture) TestDumpJSON_NotMarshalable() {
	var dumped string
	callback := func() { dumped = DumpJSON(make(chan int)) }
	this.So(callback, should.Panic)
	this.So(dumped, should.BeEmpty)

}

func (this *JSONFixture) TestDumpJSONSafe() {
	dumped, err := DumpJSONSafe(map[string]string{"Hello": "World"})
	this.So(err, should.BeNil)
	this.So(dumped, should.Equal, `{
  "Hello": "World"
}`)
}

func (this *JSONFixture) TestDumpJSONSafe_NotMarshalable() {
	dumped, err := DumpJSONSafe(make(chan int))
	this.So(err, should.NotBeNil)
	this.So(dumped, should.BeBlank)
}
