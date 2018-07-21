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
	Dstats  map[string]interface{}
}

func makeSeries(data interface{}, nilData []int, sName string) Series {
	S := Series{
		Name:    sName,
		Nildata: nilData,
		Dstats:  make(map[string]interface{}),
	}
	switch D := data.(type) {
	case []int:
		S.Idata = D
		S.Dstats["len"] = len(D)
		S.Dstats["type"] = "int"
	case []string:
		S.Sdata = D
		S.Dstats["len"] = len(D)
		S.Dstats["type"] = "string"
	case []float64:
		S.Fdata = D
		S.Dstats["len"] = len(D)
		S.Dstats["type"] = "float64"
	case []time.Time:
		S.Tdata = D
		S.Dstats["len"] = len(D)
		S.Dstats["type"] = "time.Time"
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
	first, last := getDataIndicies(start, count, S.getLength())
	switch S.getType() {
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

func (S *Series) getLength() int {
	dLen, _ := S.Dstats["len"].(int)
	return dLen
}

func (S *Series) getType() string {
	dType, _ := S.Dstats["type"].(string)
	return dType
}

func (S *Series) goDescribe() {

}
