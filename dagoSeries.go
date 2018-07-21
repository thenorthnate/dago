package dago

import (
	"fmt"
	"math"
	"time"
)

// Series : struct holding each dataset
type Series struct {
	Name    string
	Idata   []int
	Sdata   []string
	Fdata   []float64
	Tdata   []time.Time
	Nildata []int
	Dstats  sstats
}

type sstats struct {
	length int
	dtype  string
}

func makeSeries(data interface{}, nilData []int, sName string) Series {
	S := Series{
		Name:    sName,
		Nildata: nilData,
	}
	switch D := data.(type) {
	case []int:
		S.Idata = D
		S.Dstats.length = len(D)
		S.Dstats.dtype = "int"
	case []string:
		S.Sdata = D
		S.Dstats.length = len(D)
		S.Dstats.dtype = "string"
	case []float64:
		S.Fdata = D
		S.Dstats.length = len(D)
		S.Dstats.dtype = "float64"
	case []time.Time:
		S.Tdata = D
		S.Dstats.length = len(D)
		S.Dstats.dtype = "time.Time"
	}
	// Run stats here? or in New func?
	return S
}

func getDataIndicies(start int, count int, dLen int) (int, int) {
	if math.Abs(float64(start)) < float64(dLen) {
		if start >= 0 {
			if count > 0 {
				if start+count <= dLen {
					return start, start + count
				}
				return start, dLen
			}
			return 0, dLen
		}
		realStart := dLen + start
		if count > 0 {
			if realStart+count <= dLen {
				return realStart, realStart + count
			}
			return realStart, dLen
		}
		return 0, dLen
	}
	return dLen, dLen
}

func (S *Series) getSeriesData(start int, count int) (interface{}, string) {
	first, last := getDataIndicies(start, count, S.Dstats.length)
	switch S.Dstats.dtype {
	case "int":
		data := S.Idata[first:last]
		return data, fmt.Sprintf("%T", data)
	case "string":
		data := S.Sdata[first:last]
		return data, fmt.Sprintf("%T", data)
	case "float64":
		data := S.Fdata[first:last]
		return data, fmt.Sprintf("%T", data)
	case "time.Time":
		data := S.Tdata[first:last]
		return data, fmt.Sprintf("%T", data)
	}
	data := -1
	return data, fmt.Sprintf("%T", data)
}

func (S *Series) goDescribe() {

}
