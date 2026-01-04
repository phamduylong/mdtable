package mdtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* STRING FUNCTION */
const STRINGS_SHOULD_BE_THE_SAME = "The two strings should be the same"

func TestPadStart(t *testing.T) {
	originalString := "start"
	expected := "     start"
	res, err := padStart(originalString, 10, ' ')

	assert.Nil(t, err, "padStart should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadEnd(t *testing.T) {
	originalString := "end"
	expected := "end       "
	res, err := padEnd(originalString, 10, ' ')

	assert.Nil(t, err, "padEnd should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadCenterEven(t *testing.T) {
	originalString := "eleven"
	expected := "  eleven  "
	res, err := padCenter(originalString, 10, ' ')

	assert.Nil(t, err, "padCenter should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestPadCenterOdd(t *testing.T) {
	originalString := "eight"
	expected := "  eight   "
	res, err := padCenter(originalString, 10, ' ')

	assert.Nil(t, err, "padCenter should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

/* Conversion */
var dataString = [][]string{
	{"First name", "Last name", "Email", "Phone"},
	{"Jane", "Smith", "jane.smith@email.com", "555-555-1212"},
	{"John", "Doe", "john.doe@email.com", "555-555-3434"},
	{"Alice", "Wonder", "alice@wonderland.com", "555-555-5656"},
}

var dataStringWithNarrowColumn = [][]string{
	{"#", "first name", "last name", "email", "gender"},
	{"1", "Herman", "Gribbin", "hgribbin0@deliciousdays.com", "Male"},
	{"2", "Bing", "Langthorne", "blangthorne1@a8.net", "Male"},
	{"3", "Keith", "Hansford", "khansford2@reference.com", "Male"},
}

var dataStringWithPipeCharacters = [][]string{
	{"ID", "Expression", "Description"},
	{"1", "A || B", "Logical OR using pipe"},
	{"2", "foo | bar | baz", "Chained pipe values"},
	{"3", "cmd1 | cmd2", "Unix-style pipe between commands"},
	{"4", "x | y == z", "Comparison involving a pipe operator"},
}

func TestConvertGeneric(t *testing.T) {
	var cfg Config

	expected := `| First name | Last name |        Email         |    Phone     |
| :--------: | :-------: | :------------------: | :----------: |
|    Jane    |   Smith   | jane.smith@email.com | 555-555-1212 |
|    John    |    Doe    |  john.doe@email.com  | 555-555-3434 |
|   Alice    |  Wonder   | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)

}

/* ALIGN */
func TestConvertWithNarrowColumnCenterAlign(t *testing.T) {
	var cfg Config
	cfg.Align = Center

	expected := `|  #  | first name | last name  |            email            | gender |
| :-: | :--------: | :--------: | :-------------------------: | :----: |
|  1  |   Herman   |  Gribbin   | hgribbin0@deliciousdays.com |  Male  |
|  2  |    Bing    | Langthorne |     blangthorne1@a8.net     |  Male  |
|  3  |   Keith    |  Hansford  |  khansford2@reference.com   |  Male  |`

	res, err := Convert(dataStringWithNarrowColumn, cfg)

	assert.Nil(t, err, "Convert with narrow column align center should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertWithNarrowColumnLeftAlign(t *testing.T) {
	var cfg Config
	cfg.Align = Left

	expected := `| #  | first name | last name  | email                       | gender |
| :- | :--------- | :--------- | :-------------------------- | :----- |
| 1  | Herman     | Gribbin    | hgribbin0@deliciousdays.com | Male   |
| 2  | Bing       | Langthorne | blangthorne1@a8.net         | Male   |
| 3  | Keith      | Hansford   | khansford2@reference.com    | Male   |`

	res, err := Convert(dataStringWithNarrowColumn, cfg)

	assert.Nil(t, err, "Convert with narrow column align left should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertWithNarrowColumnRightAlign(t *testing.T) {
	var cfg Config
	cfg.Align = Right

	expected := `|  # | first name |  last name |                       email | gender |
| -: | ---------: | ---------: | --------------------------: | -----: |
|  1 |     Herman |    Gribbin | hgribbin0@deliciousdays.com |   Male |
|  2 |       Bing | Langthorne |         blangthorne1@a8.net |   Male |
|  3 |      Keith |   Hansford |    khansford2@reference.com |   Male |`

	res, err := Convert(dataStringWithNarrowColumn, cfg)

	assert.Nil(t, err, "Convert with narrow column align right should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestLeftAlign(t *testing.T) {
	var cfg Config
	cfg.Align = Left

	expected := `| First name | Last name | Email                | Phone        |
| :--------- | :-------- | :------------------- | :----------- |
| Jane       | Smith     | jane.smith@email.com | 555-555-1212 |
| John       | Doe       | john.doe@email.com   | 555-555-3434 |
| Alice      | Wonder    | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with left align should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestRightAlign(t *testing.T) {
	var cfg Config
	cfg.Align = Right

	expected := `| First name | Last name |                Email |        Phone |
| ---------: | --------: | -------------------: | -----------: |
|       Jane |     Smith | jane.smith@email.com | 555-555-1212 |
|       John |       Doe |   john.doe@email.com | 555-555-3434 |
|      Alice |    Wonder | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with right align should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

/* CAPTION */
func TestWithCaption(t *testing.T) {
	var cfg Config
	cfg.Caption = "Table 2: Customers who are United fans"
	expected := `<!-- Table 2: Customers who are United fans -->
| First name | Last name |        Email         |    Phone     |
| :--------: | :-------: | :------------------: | :----------: |
|    Jane    |   Smith   | jane.smith@email.com | 555-555-1212 |
|    John    |    Doe    |  john.doe@email.com  | 555-555-3434 |
|   Alice    |  Wonder   | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with caption setting should not return an error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

/* COMPACT */
func TestCompactConvertGeneric(t *testing.T) {
	var cfg Config
	cfg.Compact = true

	expected := `|First name|Last name|Email|Phone|
|:-:|:-:|:-:|:-:|
|Jane|Smith|jane.smith@email.com|555-555-1212|
|John|Doe|john.doe@email.com|555-555-3434|
|Alice|Wonder|alice@wonderland.com|555-555-5656|`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert compact should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestCompactConvertGenericLeftAlign(t *testing.T) {
	var cfg Config
	cfg.Compact = true
	cfg.Align = Left

	expected := `|First name|Last name|Email|Phone|
|:-|:-|:-|:-|
|Jane|Smith|jane.smith@email.com|555-555-1212|
|John|Doe|john.doe@email.com|555-555-3434|
|Alice|Wonder|alice@wonderland.com|555-555-5656|`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert compact left align should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestCompactConvertGenericRightAlign(t *testing.T) {
	var cfg Config
	cfg.Compact = true
	cfg.Align = Right

	expected := `|First name|Last name|Email|Phone|
|-:|-:|-:|-:|
|Jane|Smith|jane.smith@email.com|555-555-1212|
|John|Doe|john.doe@email.com|555-555-3434|
|Alice|Wonder|alice@wonderland.com|555-555-5656|`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert compact right align should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

/* COLUMNS EXCLUSION AND SORTING */
func TestConvertExcludeAllColumnsButOne(t *testing.T) {
	var cfg Config
	cfg.ExcludedColumns = []string{"Email", "First name", "Phone"}

	expected := `| Last name |
| :-------: |
|   Smith   |
|    Doe    |
|  Wonder   |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert while excluded all columns but one should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertExcludeSomeColumns(t *testing.T) {
	var cfg Config
	cfg.ExcludedColumns = []string{"Email", "First name"}

	expected := `| Last name |    Phone     |
| :-------: | :----------: |
|   Smith   | 555-555-1212 |
|    Doe    | 555-555-3434 |
|  Wonder   | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with excluded columns should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertExcludeAllColumns(t *testing.T) {
	var cfg Config
	cfg.ExcludedColumns = []string{"Email", "Last name", "First name", "Phone"}

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with all excluded columns should not return a non-nil error")

	assert.Empty(t, res, "String should be empty")
}

func TestConvertExcludeNoColumn(t *testing.T) {
	var cfg Config
	cfg.ExcludedColumns = []string{}

	expected := `| First name | Last name |        Email         |    Phone     |
| :--------: | :-------: | :------------------: | :----------: |
|    Jane    |   Smith   | jane.smith@email.com | 555-555-1212 |
|    John    |    Doe    |  john.doe@email.com  | 555-555-3434 |
|   Alice    |  Wonder   | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with empty list of excluded columns should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertSortColumnsNone(t *testing.T) {
	var cfg Config
	cfg.SortColumns = None

	expected := `| First name | Last name |        Email         |    Phone     |
| :--------: | :-------: | :------------------: | :----------: |
|    Jane    |   Smith   | jane.smith@email.com | 555-555-1212 |
|    John    |    Doe    |  john.doe@email.com  | 555-555-3434 |
|   Alice    |  Wonder   | alice@wonderland.com | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with sorted columns none should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertSortColumnsAscending(t *testing.T) {
	var cfg Config
	cfg.SortColumns = Ascending

	expected := `|        Email         | First name | Last name |    Phone     |
| :------------------: | :--------: | :-------: | :----------: |
| jane.smith@email.com |    Jane    |   Smith   | 555-555-1212 |
|  john.doe@email.com  |    John    |    Doe    | 555-555-3434 |
| alice@wonderland.com |   Alice    |  Wonder   | 555-555-5656 |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with sorted columns ascending should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertSortColumnsDescending(t *testing.T) {
	var cfg Config
	cfg.SortColumns = Descending

	expected := `|    Phone     | Last name | First name |        Email         |
| :----------: | :-------: | :--------: | :------------------: |
| 555-555-1212 |   Smith   |    Jane    | jane.smith@email.com |
| 555-555-3434 |    Doe    |    John    |  john.doe@email.com  |
| 555-555-5656 |  Wonder   |   Alice    | alice@wonderland.com |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with sorted columns descending should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

func TestConvertSortColumnsCustom(t *testing.T) {
	var cfg Config
	cfg.SortColumns = Custom
	cfg.SortFunction = func(a, b string) int {
		return len(a) - len(b)
	}

	expected := `|        Email         |    Phone     | Last name | First name |
| :------------------: | :----------: | :-------: | :--------: |
| jane.smith@email.com | 555-555-1212 |   Smith   |    Jane    |
|  john.doe@email.com  | 555-555-3434 |    Doe    |    John    |
| alice@wonderland.com | 555-555-5656 |  Wonder   |   Alice    |`

	res, err := Convert(dataString, cfg)

	assert.Nil(t, err, "Convert with sorted columns custom should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}

/* EDGE CASES */
func TestEscapePipeCharacter(t *testing.T) {
	var cfg Config
	cfg.Align = Left

	expected := `| ID | Expression        | Description                          |
| :- | :---------------- | :----------------------------------- |
| 1  | A \|\| B          | Logical OR using pipe                |
| 2  | foo \| bar \| baz | Chained pipe values                  |
| 3  | cmd1 \| cmd2      | Unix-style pipe between commands     |
| 4  | x \| y == z       | Comparison involving a pipe operator |`

	res, err := Convert(dataStringWithPipeCharacters, cfg)

	assert.Nil(t, err, "Convert with pipe characters should not return a non-nil error")

	assert.Equal(t, expected, res, STRINGS_SHOULD_BE_THE_SAME)
}
