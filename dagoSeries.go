package dago

import (
	"fmt"
	"math"
	"time"
)

const (
	iType = iota
	sType
	fType
	tType
)

// Series : struct holding each dataset
type Series struct {
	Name    string
	Idata   []int
	Sdata   []string
	Fdata   []float64
	Tdata   []time.Time
	Nildata []int
	Dstats  Sstats
}

// Sstats : struct that holds all statistical information about the series
type Sstats struct {
	// General Stats for all series types
	Length     int
	Dtype      int
	PrettyType string
	Unique     int

	// Numerical Stats
	Mean   float64
	Stddev float64
}

func makeSeries(data interface{}, nilData []int, sName string) Series {
	S := Series{
		Name:    sName,
		Nildata: nilData,
	}
	switch D := data.(type) {
	case []int:
		S.Idata = D
		S.Dstats.Length = len(D)
		S.Dstats.Dtype = iType
		S.Dstats.PrettyType = fmt.Sprintf("%T", D)
	case []string:
		S.Sdata = D
		S.Dstats.Length = len(D)
		S.Dstats.Dtype = sType
		S.Dstats.PrettyType = fmt.Sprintf("%T", D)
	case []float64:
		S.Fdata = D
		S.Dstats.Length = len(D)
		S.Dstats.Dtype = fType
		S.Dstats.PrettyType = fmt.Sprintf("%T", D)
	case []time.Time:
		S.Tdata = D
		S.Dstats.Length = len(D)
		S.Dstats.Dtype = tType
		S.Dstats.PrettyType = fmt.Sprintf("%T", D)
	}
	// Run stats here? or in New func?
	return S
}

func getDataIndicies(start, count, dLen int) (int, int) {
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

func (S *Series) getSeriesData(start, count int) interface{} {
	first, last := getDataIndicies(start, count, S.Dstats.Length)
	switch S.Dstats.Dtype {
	case iType:
		data := S.Idata[first:last]
		return data
	case sType:
		data := S.Sdata[first:last]
		return data
	case fType:
		data := S.Fdata[first:last]
		return data
	case tType:
		data := S.Tdata[first:last]
		return data
	}
	data := -1
	return data
}

func (S *Series) goDescribe() {
	switch S.Dstats.Dtype {
	case iType:
		S.describeIntSeries()
	case sType:
		S.describeStringSeries()
	case fType:
		S.describeFloatSeries()
	case tType:
		S.describeTimeSeries()
	}
}
