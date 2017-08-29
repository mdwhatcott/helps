package helps

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestJSONFixture(t *testing.T) {
	gunit.Run(new(JSONFixture), t)
}

type JSONFixture struct {
	*gunit.Fixture
}

func (this *JSONFixture) TestFormatJSON() {
	formatted, err := FormatJSON([]byte(`[{"message": "hi"}]`))
	this.So(err, should.BeNil)
	this.So(formatted, should.Equal, `[
  {
    "message": "hi"
  }
]`)
}

func (this *JSONFixture) TestDumpJSON() {
	dumped := DumpJSON(map[string]string{"Hello": "World"})
	this.So(dumped, should.Equal, `{
  "Hello": "World"
}`)
}
