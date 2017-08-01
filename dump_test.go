package helps

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestFixture(t *testing.T) {
	gunit.Run(new(Fixture), t)
}

type Fixture struct {
	*gunit.Fixture
}

func (this *Fixture) TestFormatJSON() {
	formatted, err := FormatJSON([]byte(`[{"message": "hi"}]`))
	this.So(err, should.BeNil)
	this.So(formatted, should.Equal, `[
  {
    "message": "hi"
  }
]`)
}

func (this *Fixture) TestDumpJSON() {
	dumped := DumpJSON(map[string]string{"Hello": "World"})
	this.So(dumped, should.Equal, `{
  "Hello": "World"
}`)
}

type Thing struct {
	Message string
}

func (this *Fixture) TestDumpXML() {
	dumped := DumpXML(Thing{Message: "Hello, World"})
	this.So(dumped, should.Equal, `<Thing>
  <Message>Hello, World</Message>
</Thing>`)
}

func (this *Fixture) TestFormatXML_ByteOrderMarkIgnored() {
	input := []byte(`<?xml version="1.0" encoding="utf-8"?><candidates><candidate><input_index>0</input_index><candidate_index>0</candidate_index><delivery_line_1>3214 N University Ave</delivery_line_1><last_line>Provo UT 84604-4405</last_line><delivery_point_barcode>846044405140</delivery_point_barcode><components><primary_number>3214</primary_number><street_predirection>N</street_predirection><street_name>University</street_name><street_suffix>Ave</street_suffix><city_name>Provo</city_name><state_abbreviation>UT</state_abbreviation><zipcode>84604</zipcode><plus4_code>4405</plus4_code><delivery_point>14</delivery_point><delivery_point_check_digit>0</delivery_point_check_digit></components><metadata><record_type>S</record_type><zip_type>Standard</zip_type><county_fips>49049</county_fips><county_name>Utah</county_name><carrier_route>C016</carrier_route><congressional_district>03</congressional_district><rdi>Commercial</rdi><elot_sequence>0016</elot_sequence><elot_sort>A</elot_sort><latitude>40.27658</latitude><longitude>-111.65759</longitude><precision>Zip9</precision><time_zone>Mountain</time_zone><utc_offset>-7</utc_offset><dst>true</dst></metadata><analysis><dpv_match_code>Y</dpv_match_code><dpv_footnotes>AABBR1</dpv_footnotes><dpv_cmra>Y</dpv_cmra><dpv_vacant>N</dpv_vacant><active>Y</active><suitelink_match>false</suitelink_match></analysis></candidate></candidates>`)

	output, err := FormatXML(append(utf8ByteOrderMark, input...))

	this.So(err, should.BeNil)
	this.So(output, should.Equal, `<?xml version="1.0" encoding="utf-8"?><candidates>
  <candidate>
    <input_index>0</input_index>
    <candidate_index>0</candidate_index>
    <delivery_line_1>3214 N University Ave</delivery_line_1>
    <last_line>Provo UT 84604-4405</last_line>
    <delivery_point_barcode>846044405140</delivery_point_barcode>
    <components>
      <primary_number>3214</primary_number>
      <street_predirection>N</street_predirection>
      <street_name>University</street_name>
      <street_suffix>Ave</street_suffix>
      <city_name>Provo</city_name>
      <state_abbreviation>UT</state_abbreviation>
      <zipcode>84604</zipcode>
      <plus4_code>4405</plus4_code>
      <delivery_point>14</delivery_point>
      <delivery_point_check_digit>0</delivery_point_check_digit>
    </components>
    <metadata>
      <record_type>S</record_type>
      <zip_type>Standard</zip_type>
      <county_fips>49049</county_fips>
      <county_name>Utah</county_name>
      <carrier_route>C016</carrier_route>
      <congressional_district>03</congressional_district>
      <rdi>Commercial</rdi>
      <elot_sequence>0016</elot_sequence>
      <elot_sort>A</elot_sort>
      <latitude>40.27658</latitude>
      <longitude>-111.65759</longitude>
      <precision>Zip9</precision>
      <time_zone>Mountain</time_zone>
      <utc_offset>-7</utc_offset>
      <dst>true</dst>
    </metadata>
    <analysis>
      <dpv_match_code>Y</dpv_match_code>
      <dpv_footnotes>AABBR1</dpv_footnotes>
      <dpv_cmra>Y</dpv_cmra>
      <dpv_vacant>N</dpv_vacant>
      <active>Y</active>
      <suitelink_match>false</suitelink_match>
    </analysis>
  </candidate>
</candidates>`)
}
