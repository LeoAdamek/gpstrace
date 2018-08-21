package util

import (
	"testing"
	"math"
)

var good = map[string]float64 {
	"N 50 49 47.63": 50.8298983,
	"W 1 22 3.41": -1.3676152,
	"N 39 55 1.55": 39.9170981,
	"E 116 23 25.77": 116.3904904,
}


func TestConvertDmsToDecimal(t *testing.T) {
	e := math.Nextafter(1.0, 2.0)
	
	for dms, expected := range good {
		actual, err := ConvertDmsToDecimal(dms)
		
		if err != nil {
			t.Error(err)
		}
		
		
		if math.Abs(*actual - expected) > e {
			t.Errorf("Got %#+v - Expected %#+v", *actual, expected)
		}
	}
}