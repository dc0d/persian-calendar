package calendar

type PersianDate struct{ Year, Month, Day int }

func (pd PersianDate) ToGregorian() (gd GregorianDate) {
	gd.Year, gd.Month, gd.Day = persianToGregorian(pd.Year, pd.Month, pd.Day)
	return
}

func (pd PersianDate) IsLeap() bool {
	nextPD := PersianDate{Year: pd.Year + 1, Month: 1, Day: 1}
	gd := nextPD.ToGregorian()
	gd.Day--
	lastDay := gd.ToPersian().Day
	return lastDay == 30
}

type GregorianDate struct{ Year, Month, Day int }

func (gd GregorianDate) ToPersian() (pd PersianDate) {
	pd.Year, pd.Month, pd.Day = gregorianToPersian(gd.Year, gd.Month, gd.Day)
	return
}

func persianToGregorian(pyear, pmonth, pday int) (gyear, gmonth, gday int) {
	jd := jal2jd(int32(pyear), int32(pmonth), int32(pday))
	L, M, N := jd2jg(jd, 0)
	gyear, gmonth, gday = int(L), int(M), int(N)
	return
}

func gregorianToPersian(gyear, gmonth, gday int) (pyear, pmonth, pday int) {
	jd := jg2jd(int32(gyear), int32(gmonth), int32(gday), 0)
	L, M, N := jd2jal(jd)
	pyear, pmonth, pday = int(L), int(M), int(N)
	return
}

func jal2jd(Jy, Jm, Jd int32) int32 {
	//Converts a date of the Jalaali calendar to the Julian Day Number
	//Input:  Jy - Jalaali year (1 to 3100)
	//        Jm - month (1 to 12)
	//        Jd - day (1 to 29/31)
	//Output: Jal2JD - the Julian Day Number
	_, iGy, March := jalCal(Jy)
	return jg2jd(iGy, 3, March, 0) + (Jm-1)*31 - Jm/7*(Jm-7) + Jd - 1
}

func jd2jal(JDN int32) (Jy, Jm, Jd int32) {
	//Converts the Julian Day number to a date in the Jalaali calendar
	//Input: JDN - the Julian Day number
	//Output: Jy - Jalaali year (1 to 3100)
	//        Jm - month (1 to 12)
	//        Jd - day (1 to 29/31)

	//Calculate Gregorian year (L)
	L, _, _ := jd2jg(JDN, 0)
	Jy = L - 621
	leap, _, March := jalCal(Jy)
	JDN1F := jg2jd(L, 3, March, 0)
	//Find number of days that passed since 1 Farvardin
	k := JDN - JDN1F
	if k >= 0 {
		if k <= 185 {
			//The first 6 months
			Jm = 1 + k/31
			Jd = mod(k, 31) + 1
			return
		}

		//The remaining months
		k = k - 186
	} else {
		//previous Jalaali year
		Jy = Jy - 1
		k = k + 179
		if leap == 1 {
			k = k + 1
		}
	}
	Jm = 7 + k/30
	Jd = mod(k, 30) + 1

	return
}

//This procedure determines if the Jalaali (Persian) year is
//leap (366-day long) or is the common year (365 days), and
//finds the day in March (Gregorian calendar) of the first
//day of the Jalaali year (Jy)
//Input:  Jy - Jalaali calendar year (-61 to 3177)
//Output:
//  leap  - number of years since the last leap year (0 to 4)
//  Gy    - Gregorian year of the beginning of Jalaali year
//  March - the March day of Farvardin the 1st (1st day of Jy)
//Jalaali years starting the 33-year rule
func jalCal(Jy int32) (leap, Gy, March int32) {
	leap = 0
	Gy = 0
	March = 0
	Gy = Jy + 621
	leapJ := int32(-14)
	jp := breaks[0]
	if Jy < jp || Jy > breaks[19] {
		//      print'(10x,a,i5,a,i5,a)',
		//*' Invalid Jalaali year number:',Jy,' (=',Gy,' Gregorian)'
	}
	//Find the limiting years for the Jalaali year Jy
	jump := int32(0)
	for j := int32(1); j <= 19; j++ {
		jm := breaks[j]
		jump = jm - jp
		if Jy < jm {
			goto Label2
		}
		//Q:should these 2 lines be in the for loop?
		leapJ = leapJ + jump/33*8 + mod(jump, 33)/4
		//Label1:
		jp = jm
	}
Label2:
	N := int32(Jy - jp)
	//Find the number of leap years from AD 621 to the beginning
	//of the current Jalaali year in the Persian calendar
	leapJ = leapJ + N/33*8 + (mod(N, 33)+3)/4
	if mod(jump, 33) == 4 && (jump-N) == 4 {
		leapJ = leapJ + 1
	}
	//and the same in the Gregorian calendar (until the year Gy)
	leapG := int32(Gy/4 - (Gy/100+1)*3/4 - 150)
	//Determine the Gregorian date of Farvardin the 1st
	March = 20 + leapJ - leapG
	//Find how many years have passed since the last leap year
	if (jump - N) < 6 {
		N = N - jump + (jump+4)/33*33
	}
	leap = mod(mod(N+1, 33)-1, 4)
	if leap == -1 {
		leap = 4
	}

	return
}

var (
	breaks = []int32{
		-61, 9, 38, 199, 426,
		686, 756, 818, 1111, 1181,
		1210, 1635, 2060, 2097, 2192,
		2262, 2324, 2394, 2456, 3178}
)

//Input:  JD   - Julian Day number
//        J1G0 - to be set to 1 for Julian and to 0 for Gregorian calendar
//Output: L - calendar year (years BC numbered 0, -1, -2, ...)
//        M - calendar month (for January M=1, February M=2, ... M=12)
//        N - calendar day of the month M (1 to 28/29/30/31)
// Calculates Gregorian and Julian calendar dates from the Julian Day number
// (JD) for the period since JD=-34839655 (i.e. the year -100100 of both
// the calendars) to some millions (10**6) years ahead of the present.
// The algorithm is based on D.A. Hatcher, Q.Jl.R.Astron.Soc. 25(1984), 53-55
// slightly modified by me (K.M. Borkowski, Post.Astron. 25(1987), 275-279).
func jd2jg(JD, J1G0 int32) (L, M, N int32) {
	var I, J int32
	J = 4*JD + 139361631
	if J1G0 <= 0 {
		J = J + (4*JD+183187720)/146097*3/4*4 - 3908
	}
	I = mod(J, 1461)/4*5 + 308
	N = mod(I, 153)/5 + 1
	M = mod(I/153, 12) + 1
	L = J/1461 - 100100 + (8-M)/6

	return
}

//Input:  L - calendar year (years BC numbered 0, -1, -2, ...)
//        M - calendar month (for January M=1, February M=2, ..., M=12)
//        N - calendar day of the month M (1 to 28/29/30/31)
//     J1G0 - to be set to 1 for Julian and to 0 for Gregorian calendar
//Output: jg2jd - Julian Day number
// Calculates the Julian Day number (jg2jd) from Gregorian or Julian
// calendar dates. This integer number corresponds to the noon of
// the date (i.e. 12 hours of Universal Time).
// The procedure was tested to be good since 1 March, -100100 (of both
// the calendars) up to a few millions (10**6) years into the future.
// The algorithm is based on D.A. Hatcher, Q.Jl.R.Astron.Soc. 25(1984), 53-55
// slightly modified by me (K.M. Borkowski, Post.Astron. 25(1987), 275-279).
func jg2jd(L, M, N, J1G0 int32) (res int32) {
	res = (L+(M-8)/6+100100)*1461/4 + (153*mod(M+9, 12)+2)/5 + N - 34840408
	if J1G0 <= 0 {
		res = res - (L+100100+(M-8)/6)/100*3/4 + 752
	}
	//MJD=jg2jd-2400000.5   ! this formula gives Modified Julian Day number
	return
}

func mod(a, b int32) int32 {
	return a % b
}
