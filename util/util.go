package util

import (
	"regexp"
	"errors"
	"math/big"
	"strconv"
)

var dmsRegexp = regexp.MustCompile(`([NSEW]) (\d+) (\d+)'? (\d+(\.\d+)?)"?`)

func ConvertDmsToDecimal(dms string) (*float64, error) {
	if !dmsRegexp.MatchString(dms) {
		return nil, errors.New("invalid dms text")
	}
	
	matches := dmsRegexp.FindAllStringSubmatch(dms, 4)[0]
	
	deg, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return nil, err
	}
	
	dec := big.NewRat(deg, 1)
	
	min, err := strconv.ParseInt(matches[3], 10, 64)
	
	if err != nil {
		return nil, err
	}
	
	dec = dec.Add(dec, big.NewRat(min, 60))
	
	sec, err := strconv.ParseFloat(matches[4], 64)
	
	if err != nil {
		return nil, err
	}
	
	sr := big.NewRat(0, 1).SetFloat64(sec)
	sr = sr.Mul(sr, big.NewRat(1, 3600))
	
	dec = dec.Add(dec, sr)
	
	if matches[1] == "S" || matches[1] == "W" {
		dec = dec.Neg(dec)
	}
	
	f, _ := dec.Float64()
	return &f, nil
}
