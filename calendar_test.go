package calendar_test

// See the LICENSE file.

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	calendar "github.com/dc0d/persian-calendar"

	assert "github.com/stretchr/testify/require"
)

func Test_persian_leap_years(t *testing.T) {
	fixtures := readTestFixtures()

	for fx := range fixtures {
		pd := fx.GregorianDate.ToPersian()
		isLeapYear := pd.IsLeap()
		shouldBeLeapYear := fx.IsPersianLeap

		assert.Equal(t, shouldBeLeapYear, isLeapYear, "expected to be a leap year")
	}
}

func Test_conversion(t *testing.T) {
	fixtures := readTestFixtures()

	for fx := range fixtures {
		fx := fx
		t.Run("persian to gregorian", func(t *testing.T) {
			var (
				persianDate           = fx.PersianDate
				expectedGregorianDate = fx.GregorianDate
			)

			assert.Equal(t, expectedGregorianDate, persianDate.ToGregorian())
		})

		t.Run("gregorian to persian", func(t *testing.T) {
			var (
				gregorianDate       = fx.GregorianDate
				expectedPersianDate = fx.PersianDate
			)

			assert.Equal(t, expectedPersianDate, gregorianDate.ToPersian())
		})
	}
}

func readTestFixtures() chan fixture {
	f, err := os.Open("./test/fixtures/test_cases")
	check(err)

	lines := readLines(f)
	return fixtures(lines)
}

func fixtures(lines chan []byte) chan fixture {
	result := make(chan fixture, 100)

	go func() {
		defer close(result)

		for l := range lines {
			var fx fixture
			fromJSON(l, &fx)

			result <- fx
		}
	}()

	return result
}

func readLines(f *os.File) chan []byte {
	lines := make(chan []byte, 100)

	go func() {
		defer close(lines)

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		buff := make([]byte, 1024*1024*8)
		scanner.Buffer(buff, len(buff))

		for scanner.Scan() {
			lines <- scanner.Bytes()
		}
	}()

	return lines
}

func fromJSON(data []byte, ptr interface{}) {
	if err := json.Unmarshal(data, ptr); err != nil {
		fmt.Println(err)
		fmt.Println(string(data))
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type fixture struct {
	IsPersianLeap bool                   `json:"is_persian_leap,omitempty"`
	PersianDate   calendar.PersianDate   `json:"persian_date,omitempty"`
	GregorianDate calendar.GregorianDate `json:"gregorian_date,omitempty"`
}
