package calendar_test

// See the LICENSE file.

import (
	"testing"

	calendar "github.com/dc0d/persian-calendar"

	assert "github.com/stretchr/testify/require"
)

func Test_convert_dates_for_2020_5_29(t *testing.T) {
	t.Run("persian to gregorian", func(t *testing.T) {
		var (
			assert = assert.New(t)
		)

		var (
			persianDate calendar.PersianDate

			expectedGregorianDate calendar.GregorianDate
		)

		{
			persianDate = calendar.PersianDate{
				Year:  1399,
				Month: 3,
				Day:   9,
			}

			expectedGregorianDate = calendar.GregorianDate{
				Year:  2020,
				Month: 5,
				Day:   29,
			}
		}

		assert.Equal(expectedGregorianDate, persianDate.ToGregorian())
	})

	t.Run("gregorian to persian", func(t *testing.T) {
		var (
			assert = assert.New(t)
		)

		var (
			gregorianDate calendar.GregorianDate

			expectedPersianDate calendar.PersianDate
		)

		{
			gregorianDate = calendar.GregorianDate{
				Year:  2020,
				Month: 5,
				Day:   29,
			}

			expectedPersianDate = calendar.PersianDate{
				Year:  1399,
				Month: 3,
				Day:   9,
			}
		}

		assert.Equal(expectedPersianDate, gregorianDate.ToPersian())
	})
}

func Test_convert_gregorian_to_persian(t *testing.T) {
	for _, tc := range testCases {
		persianDate := tc.gregorianDate.ToPersian()
		expectedPersianDate := tc.persianDate

		assert.Equal(t, expectedPersianDate, persianDate)
	}
}

func Test_convert_persian_to_gregorian(t *testing.T) {
	for _, tc := range testCases {
		gregorianDate := tc.persianDate.ToGregorian()
		expectedGregorianDate := tc.gregorianDate

		assert.Equal(t, expectedGregorianDate, gregorianDate)
	}
}

func Test_persian_leap_years(t *testing.T) {
	for _, tc := range testCases {
		pd := tc.gregorianDate.ToPersian()
		isLeapYear := pd.IsLeap()
		shouldBeLeapYear := tc.isPersianLeap

		assert.Equal(t, shouldBeLeapYear, isLeapYear, "expected to be leap year")
	}
}

var testCases []testCase

func init() {
	for i, isLeap := range initialPersianLeapYears {
		tc := testCase{
			isPersianLeap: isLeap,
			persianDate:   initialPersianDates[i],
			gregorianDate: initialGregorianDates[i],
		}
		testCases = append(testCases, tc)
	}

	testCases = append(testCases, testCase{
		isPersianLeap: false,
		persianDate:   calendar.PersianDate{Year: 1209, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1830, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1210, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1831, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1214, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1835, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1218, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1839, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1226, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1847, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1230, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1851, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1234, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1855, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1238, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1859, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1243, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1864, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1247, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1868, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1251, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1872, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1255, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1876, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1259, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1880, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1263, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1884, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1267, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1888, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1271, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1892, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1276, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1897, Month: 3, Day: 20},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1280, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1901, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1284, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1905, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1288, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1909, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1292, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1913, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1296, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1917, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1300, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1921, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1304, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1925, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1309, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1930, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1313, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1934, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1317, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1938, Month: 3, Day: 21},
	})

	testCases = append(testCases, testCase{
		isPersianLeap: true,
		persianDate:   calendar.PersianDate{Year: 1321, Month: 1, Day: 1},
		gregorianDate: calendar.GregorianDate{Year: 1942, Month: 3, Day: 21},
	})
}

type testCase struct {
	isPersianLeap bool
	persianDate   calendar.PersianDate
	gregorianDate calendar.GregorianDate
}

var initialPersianLeapYears = []bool{
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
}

var initialPersianDates = []calendar.PersianDate{
	{Year: 1354, Month: 1, Day: 1},
	{Year: 1355, Month: 1, Day: 1},
	{Year: 1356, Month: 1, Day: 1},
	{Year: 1357, Month: 1, Day: 1},
	{Year: 1358, Month: 1, Day: 1},
	{Year: 1359, Month: 1, Day: 1},
	{Year: 1360, Month: 1, Day: 1},
	{Year: 1361, Month: 1, Day: 1},
	{Year: 1362, Month: 1, Day: 1},
	{Year: 1363, Month: 1, Day: 1},
	{Year: 1364, Month: 1, Day: 1},
	{Year: 1365, Month: 1, Day: 1},
	{Year: 1366, Month: 1, Day: 1},
	{Year: 1367, Month: 1, Day: 1},
	{Year: 1368, Month: 1, Day: 1},
	{Year: 1369, Month: 1, Day: 1},
	{Year: 1370, Month: 1, Day: 1},
	{Year: 1371, Month: 1, Day: 1},
	{Year: 1372, Month: 1, Day: 1},
	{Year: 1373, Month: 1, Day: 1},
	{Year: 1374, Month: 1, Day: 1},
	{Year: 1375, Month: 1, Day: 1},
	{Year: 1376, Month: 1, Day: 1},
	{Year: 1377, Month: 1, Day: 1},
	{Year: 1378, Month: 1, Day: 1},
	{Year: 1379, Month: 1, Day: 1},
	{Year: 1380, Month: 1, Day: 1},
	{Year: 1381, Month: 1, Day: 1},
	{Year: 1382, Month: 1, Day: 1},
	{Year: 1383, Month: 1, Day: 1},
	{Year: 1384, Month: 1, Day: 1},
	{Year: 1385, Month: 1, Day: 1},
	{Year: 1386, Month: 1, Day: 1},
	{Year: 1387, Month: 1, Day: 1},
	{Year: 1388, Month: 1, Day: 1},
	{Year: 1389, Month: 1, Day: 1},
	{Year: 1390, Month: 1, Day: 1},
	{Year: 1391, Month: 1, Day: 1},
	{Year: 1392, Month: 1, Day: 1},
	{Year: 1393, Month: 1, Day: 1},
	{Year: 1394, Month: 1, Day: 1},
	{Year: 1395, Month: 1, Day: 1},
	{Year: 1396, Month: 1, Day: 1},
	{Year: 1397, Month: 1, Day: 1},
	{Year: 1398, Month: 1, Day: 1},
	{Year: 1399, Month: 1, Day: 1},
	{Year: 1400, Month: 1, Day: 1},
	{Year: 1401, Month: 1, Day: 1},
	{Year: 1402, Month: 1, Day: 1},
	{Year: 1403, Month: 1, Day: 1},
	{Year: 1404, Month: 1, Day: 1},
	{Year: 1405, Month: 1, Day: 1},
	{Year: 1406, Month: 1, Day: 1},
	{Year: 1407, Month: 1, Day: 1},
	{Year: 1408, Month: 1, Day: 1},
	{Year: 1409, Month: 1, Day: 1},
	{Year: 1410, Month: 1, Day: 1},
	{Year: 1411, Month: 1, Day: 1},
	{Year: 1412, Month: 1, Day: 1},
	{Year: 1413, Month: 1, Day: 1},
	{Year: 1414, Month: 1, Day: 1},
	{Year: 1415, Month: 1, Day: 1},
	{Year: 1416, Month: 1, Day: 1},
	{Year: 1417, Month: 1, Day: 1},
	{Year: 1418, Month: 1, Day: 1},
	{Year: 1419, Month: 1, Day: 1},
}

var initialGregorianDates = []calendar.GregorianDate{
	{Year: 1975, Month: 3, Day: 21},
	{Year: 1976, Month: 3, Day: 21},
	{Year: 1977, Month: 3, Day: 21},
	{Year: 1978, Month: 3, Day: 21},
	{Year: 1979, Month: 3, Day: 21},
	{Year: 1980, Month: 3, Day: 21},
	{Year: 1981, Month: 3, Day: 21},
	{Year: 1982, Month: 3, Day: 21},
	{Year: 1983, Month: 3, Day: 21},
	{Year: 1984, Month: 3, Day: 21},
	{Year: 1985, Month: 3, Day: 21},
	{Year: 1986, Month: 3, Day: 21},
	{Year: 1987, Month: 3, Day: 21},
	{Year: 1988, Month: 3, Day: 21},
	{Year: 1989, Month: 3, Day: 21},
	{Year: 1990, Month: 3, Day: 21},
	{Year: 1991, Month: 3, Day: 21},
	{Year: 1992, Month: 3, Day: 21},
	{Year: 1993, Month: 3, Day: 21},
	{Year: 1994, Month: 3, Day: 21},
	{Year: 1995, Month: 3, Day: 21},
	{Year: 1996, Month: 3, Day: 20},
	{Year: 1997, Month: 3, Day: 21},
	{Year: 1998, Month: 3, Day: 21},
	{Year: 1999, Month: 3, Day: 21},
	{Year: 2000, Month: 3, Day: 20},
	{Year: 2001, Month: 3, Day: 21},
	{Year: 2002, Month: 3, Day: 21},
	{Year: 2003, Month: 3, Day: 21},
	{Year: 2004, Month: 3, Day: 20},
	{Year: 2005, Month: 3, Day: 21},
	{Year: 2006, Month: 3, Day: 21},
	{Year: 2007, Month: 3, Day: 21},
	{Year: 2008, Month: 3, Day: 20},
	{Year: 2009, Month: 3, Day: 21},
	{Year: 2010, Month: 3, Day: 21},
	{Year: 2011, Month: 3, Day: 21},
	{Year: 2012, Month: 3, Day: 20},
	{Year: 2013, Month: 3, Day: 21},
	{Year: 2014, Month: 3, Day: 21},
	{Year: 2015, Month: 3, Day: 21},
	{Year: 2016, Month: 3, Day: 20},
	{Year: 2017, Month: 3, Day: 21},
	{Year: 2018, Month: 3, Day: 21},
	{Year: 2019, Month: 3, Day: 21},
	{Year: 2020, Month: 3, Day: 20},
	{Year: 2021, Month: 3, Day: 21},
	{Year: 2022, Month: 3, Day: 21},
	{Year: 2023, Month: 3, Day: 21},
	{Year: 2024, Month: 3, Day: 20},
	{Year: 2025, Month: 3, Day: 21},
	{Year: 2026, Month: 3, Day: 21},
	{Year: 2027, Month: 3, Day: 21},
	{Year: 2028, Month: 3, Day: 20},
	{Year: 2029, Month: 3, Day: 20},
	{Year: 2030, Month: 3, Day: 21},
	{Year: 2031, Month: 3, Day: 21},
	{Year: 2032, Month: 3, Day: 20},
	{Year: 2033, Month: 3, Day: 20},
	{Year: 2034, Month: 3, Day: 21},
	{Year: 2035, Month: 3, Day: 21},
	{Year: 2036, Month: 3, Day: 20},
	{Year: 2037, Month: 3, Day: 20},
	{Year: 2038, Month: 3, Day: 21},
	{Year: 2039, Month: 3, Day: 21},
	{Year: 2040, Month: 3, Day: 20},
}
