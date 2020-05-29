package calendar

// See the LICENSE file.

import "fmt"

func ExamplePersianDate() {
	persianDate := PersianDate{
		Year:  1399,
		Month: 3,
		Day:   9,
	}

	fmt.Printf("%#v\n", persianDate.ToGregorian())

	// Output:
	// calendar.GregorianDate{Year:2020, Month:5, Day:29}
}

func ExampleGregorianDate() {
	gregorianDate := GregorianDate{
		Year:  2020,
		Month: 5,
		Day:   29,
	}

	fmt.Printf("%#v\n", gregorianDate.ToPersian())

	// Output:
	// calendar.PersianDate{Year:1399, Month:3, Day:9}
}
