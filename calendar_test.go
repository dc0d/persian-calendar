package calendar_test

import (
	"fmt"
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

func Test_convert_gregorian_to_persian_for_samples(t *testing.T) {
	for i, gregorianDate := range sampleGregorianDates {
		persianDate := gregorianDate.ToPersian()
		expectedPersianDate := samplePersianDates[i]

		t.Run(fmt.Sprintf("year %d", persianDate.Year), func(t *testing.T) {
			assert.Equal(t, expectedPersianDate, persianDate)
		})
	}
}

var samplePersianDates = []calendar.PersianDate{
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

var sampleGregorianDates = []calendar.GregorianDate{
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
