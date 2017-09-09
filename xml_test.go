package helps

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestXMLFixture(t *testing.T) {
	gunit.Run(new(XMLFixture), t)
}

type XMLFixture struct {
	*gunit.Fixture
}

func (this *XMLFixture) TestDumpXMLSafe_ByteOrderMarkIgnored() {
	output, err := DumpXMLSafe(Thing{Message: "Hello, World"})
	this.So(err, should.BeNil)
	this.So(output, should.Equal, `<Thing>
  <Message>Hello, World</Message>
</Thing>`)
}

func (this *XMLFixture) TestDumpXMLSafe_BadXML() {
	output, err := DumpXMLSafe(make(chan int))
	this.So(err, should.NotBeNil)
	this.So(output, should.BeBlank)
}

func (this *XMLFixture) TestDumpXML() {
	dumped := DumpXML(Thing{Message: "Hello, World"})
	this.So(dumped, should.Equal, `<Thing>
  <Message>Hello, World</Message>
</Thing>`)
}

func (this *XMLFixture) TestDumpXML_NotMarshalable() {
	var dumped string
	callback := func() { dumped = DumpXML(make(chan int)) }
	this.So(callback, should.Panic)
	this.So(dumped, should.BeEmpty)
}

func (this *XMLFixture) TestFormatXMLSafe_ByteOrderMarkIgnored() {
	output, err := FormatXMLSafe(inputXML)
	this.So(err, should.BeNil)
	this.So(output, should.Equal, expectedFormattedXML)
}

func (this *XMLFixture) TestFormatXMLSafe_BadXML() {
	output, err := FormatXMLSafe([]byte(`<asdf`))
	this.So(err, should.NotBeNil)
	this.So(output, should.BeBlank)
}

func (this *XMLFixture) TestFormatXML() {
	output := FormatXML(inputXML)
	this.So(output, should.Equal, expectedFormattedXML)
}

func (this *XMLFixture) TestFormatXML_BadXML() {
	var output string
	callback := func() { output = FormatXML([]byte(`<asdf`)) }
	this.So(callback, should.Panic)
	this.So(output, should.BeBlank)
}

type Thing struct {
	Message string
}

var inputXML = append(utf8ByteOrderMark,
	[]byte(`<?xml version="1.0" encoding="utf-8"?><candidates><candidate><input_index>0</input_index><candidate_index>0</candidate_index><delivery_line_1>3214 N University Ave</delivery_line_1><last_line>Provo UT 84604-4405</last_line><delivery_point_barcode>846044405140</delivery_point_barcode><components><primary_number>3214</primary_number><street_predirection>N</street_predirection><street_name>University</street_name><street_suffix>Ave</street_suffix><city_name>Provo</city_name><state_abbreviation>UT</state_abbreviation><zipcode>84604</zipcode><plus4_code>4405</plus4_code><delivery_point>14</delivery_point><delivery_point_check_digit>0</delivery_point_check_digit></components><metadata><record_type>S</record_type><zip_type>Standard</zip_type><county_fips>49049</county_fips><county_name>Utah</county_name><carrier_route>C016</carrier_route><congressional_district>03</congressional_district><rdi>Commercial</rdi><elot_sequence>0016</elot_sequence><elot_sort>A</elot_sort><latitude>40.27658</latitude><longitude>-111.65759</longitude><precision>Zip9</precision><time_zone>Mountain</time_zone><utc_offset>-7</utc_offset><dst>true</dst></metadata><analysis><dpv_match_code>Y</dpv_match_code><dpv_footnotes>AABBR1</dpv_footnotes><dpv_cmra>Y</dpv_cmra><dpv_vacant>N</dpv_vacant><active>Y</active><suitelink_match>false</suitelink_match></analysis></candidate></candidates>`)...)

const expectedFormattedXML = `<?xml version="1.0" encoding="utf-8"?><candidates>
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
</candidates>`
