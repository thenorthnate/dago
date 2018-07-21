package dago

import (
	"fmt"
	"math"
	"time"
)

// Series : interface for series objects
type Series interface {
	getSeriesData(start int, count int) (interface{}, string)
	getName() string
	changeName(name string) Series
	getLength() int
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

// INT SERIES
type intSeries struct {
	Name    string
	Data    []int
	NilData []int
}

func makeIntSeries(data []int, nilData []int, name string) intSeries {
	ss := intSeries{
		Name:    name,
		Data:    data,
		NilData: nilData,
	}
	return ss
}

func (ss intSeries) getSeriesData(start int, count int) (interface{}, string) {
	first, last := getDataIndicies(start, count, len(ss.Data))
	return ss.Data[first:last], fmt.Sprintf("%T", ss.Data)
}

func (ss intSeries) changeName(name string) Series {
	ss.Name = name
	return ss
}

func (ss intSeries) getName() string {
	return ss.Name
}

func (ss intSeries) getLength() int {
	return len(ss.Data)
}

func (ss intSeries) convertTo(dtype string) Series {
	switch dtype {
	case "int":
		return ss
	case "string":
		ns := make([]string, len(ss.Data))
		for i, val := range ss.Data {
			ns[i] = fmt.Sprintf("%v", val)
		}
		return makeStringSeries(ns, ss.NilData, ss.Name)
	case "float64":
		nf := make([]float64, len(ss.Data))
		for i, val := range ss.Data {
			nf[i] = float64(val)
		}
		return makeFloatSeries(nf, ss.NilData, ss.Name)
	case "time.Time":
		nt := make([]time.Time, len(ss.Data))
		for i, val := range ss.Data {
			nt[i] = time.Unix(int64(val), 0)
		}
		return makeTimeSeries(nt, ss.NilData, ss.Name)
	}
	return ss
}

// STRING SERIES
type stringSeries struct {
	Name    string
	Data    []string
	NilData []int
}

func makeStringSeries(data []string, nilData []int, name string) stringSeries {
	ss := stringSeries{
		Name:    name,
		Data:    data,
		NilData: nilData,
	}
	return ss
}

func (ss stringSeries) getSeriesData(start int, count int) (interface{}, string) {
	first, last := getDataIndicies(start, count, len(ss.Data))
	return ss.Data[first:last], fmt.Sprintf("%T", ss.Data)
}

func (ss stringSeries) changeName(name string) Series {
	ss.Name = name
	return ss
}

func (ss stringSeries) getName() string {
	return ss.Name
}

func (ss stringSeries) getLength() int {
	return len(ss.Data)
}

// FLOAT SERIES
type float64Series struct {
	Name    string
	Data    []float64
	NilData []int
}

func makeFloatSeries(data []float64, nilData []int, name string) float64Series {
	ss := float64Series{
		Name:    name,
		Data:    data,
		NilData: nilData,
	}
	return ss
}

func (ss float64Series) getSeriesData(start int, count int) (interface{}, string) {
	first, last := getDataIndicies(start, count, len(ss.Data))
	return ss.Data[first:last], fmt.Sprintf("%T", ss.Data)
}

func (ss float64Series) changeName(name string) Series {
	ss.Name = name
	return ss
}

func (ss float64Series) getName() string {
	return ss.Name
}

func (ss float64Series) getLength() int {
	return len(ss.Data)
}

// TIME SERIES
type timeSeries struct {
	Name    string
	Data    []time.Time
	NilData []int
}

func makeTimeSeries(data []time.Time, nilData []int, name string) timeSeries {
	ss := timeSeries{
		Name:    name,
		Data:    data,
		NilData: nilData,
	}
	return ss
}

func (ss timeSeries) getSeriesData(start int, count int) (interface{}, string) {
	first, last := getDataIndicies(start, count, len(ss.Data))
	return ss.Data[first:last], fmt.Sprintf("%T", ss.Data)
}

func (ss timeSeries) changeName(name string) Series {
	ss.Name = name
	return ss
}

func (ss timeSeries) getName() string {
	return ss.Name
}

func (ss timeSeries) getLength() int {
	return len(ss.Data)
}
