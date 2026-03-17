package common

import "strconv"

func BirthDateFromString(birthDateStr string) (BirthDate, error) {

	yearInt, err := strconv.ParseUint(birthDateStr[0:4], 10, 64)
	if err != nil {
		return BirthDate{}, err
	}

	monthInt, err := strconv.ParseUint(birthDateStr[5:7], 10, 8)
	if err != nil {
		return BirthDate{}, err
	}
	dayInt, err := strconv.ParseUint(birthDateStr[8:10], 10, 8)
	if err != nil {
		return BirthDate{}, err
	}

	return BirthDate{
		day:   uint8(dayInt),
		month: uint8(monthInt),
		year:  yearInt,
	}, nil
}
