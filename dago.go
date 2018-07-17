package dago

import (
	"fmt"
	"time"
)

// DataFrame : data type that holds all data from dago package
type DataFrame struct {
	Sets  []Series
	level bool
}

// Options : Struct to hold user parameters during certain operations
type Options struct {
}

/*
type series struct {
	Name    string
	Data    dset
	NilData []int // stores the indices of any nil values
	// Dtype   string
}

type dset struct {
	iData []int
	sData []string
	fData []float64
	tData []time.Time
}
*/

// New : Returns a new DataFrame object
func New(data ...interface{}) DataFrame {
	DF := DataFrame{}
	for _, item := range data {
		switch d := item.(type) {
		case []int:
			DF.Sets = append(DF.Sets, makeIntSeries(d, []int{}, ""))
		case []string:
			DF.Sets = append(DF.Sets, makeStringSeries(d, []int{}, ""))
		case []float64:
			DF.Sets = append(DF.Sets, makeFloatSeries(d, []int{}, ""))
		case []time.Time:
			DF.Sets = append(DF.Sets, makeTimeSeries(d, []int{}, ""))
		case [][]int:
			for _, v := range d {
				DF.Sets = append(DF.Sets, makeIntSeries(v, []int{}, ""))
			}
		case map[string][]int:
			for k, v := range d {
				DF.Sets = append(DF.Sets, makeIntSeries(v, []int{}, k))
			}
		case map[string][]string:
			for k, v := range d {
				DF.Sets = append(DF.Sets, makeStringSeries(v, []int{}, k))
			}
		default:
			// do something
		}
	}

	return DF
}

// Add : Add new data to a pre-existing DataFrame
func (DF *DataFrame) Add(data interface{}) {
	newDf := New(data)
	if DF.level {
		if newDf.level {
			if newDf.Sets[0].getLength() == DF.Sets[0].getLength() {
				for _, v := range newDf.Sets {
					DF.Sets = append(DF.Sets, v)
				}
			}
		}

	} else {
		for _, v := range newDf.Sets {
			DF.Sets = append(DF.Sets, v)
		}
	}

}

// GetInts : Returns the slice of ints in that series, or an empty slice if series has no ints
func (DF *DataFrame) GetInts(index int) []int {
	SS := DF.Sets[index]
	rawData, _ := SS.getSeriesData(0, -1)
	data, ok := rawData.([]int)
	if ok {
		return data
	}
	return []int{}
}

func (DF *DataFrame) getIndicies(sets ...interface{}) []int {
	selectedIndicies := map[int]bool{}
	for _, v := range sets {
		switch name := v.(type) {
		case string:
			for i, ss := range DF.Sets {
				setName := ss.getName()
				if setName == name {
					selectedIndicies[i] = true
				}
			}
		case int:
			selectedIndicies[name] = true
		}
	}
	outIndicies := []int{}
	for k := range selectedIndicies {
		outIndicies = append(outIndicies, k)
	}
	return outIndicies
}

// Select : Make a selection from the current DataFrame
func (DF *DataFrame) Select(sets ...interface{}) DataFrame {
	selectedIndicies := DF.getIndicies(sets)
	newDF := DataFrame{}
	for _, v := range selectedIndicies {
		if v >= 0 && v < len(DF.Sets) {
			newDF.Sets = append(newDF.Sets, DF.Sets[v])
		}
	}
	return newDF
}

/*
func (DF *DataFrame) deleteIndex(index int) {
	if index < len(DF.Sets) && index >= 0 {
		if len(DF.Sets) == 1 {
			delete(DF.Sets, index)
		} else {
			for i := 0; i < len(DF.Sets); i++ {
				if i > index {
					keepss := DF.Sets[i]
					DF.Sets[i-1] = keepss
				}
			}
			delete(DF.Sets, len(DF.Sets)-1)
		}
	}
}

// Delete : removes series from the DataFrame
func (DF *DataFrame) Delete(index int, name string) {
	if index < 0 {
		indicies := DF.getIndicies(name)
		for _, v := range indicies {
			DF.deleteIndex(v)
		}
	} else {
		DF.deleteIndex(index)
	}
}


*/

// Rename : Allows user to edit the name of a series
func (DF *DataFrame) Rename(index int, name string) {
	if index < len(DF.Sets) && index >= 0 {
		ss := DF.Sets[index]
		ss = ss.changeName(name)
		DF.Sets[index] = ss
	}
}

// Levelize : makes all series in the DataFrame the same length. Many methods
// will not work unless the DataFrame is level
func (DF *DataFrame) Levelize() {

}

// Head : Prints the first few rows in the dataframe
func (DF *DataFrame) Head(rows int) {
	for i, v := range DF.Sets {
		data, dtype := v.getSeriesData(0, rows)
		DType := string(dtype[0]) + fmt.Sprintf("%v", DF.Sets[i].getLength()) + string(dtype[1:])
		fmt.Printf("%5v %15v %20v\t%v\n", i, DF.Sets[i].getName(), DType, data)
	}
}

// Tail : Prints the last few rows in the dataframe
func (DF *DataFrame) Tail(rows int) {
	for i, v := range DF.Sets {
		data, dtype := v.getSeriesData(-rows, rows)
		DType := string(dtype[0]) + fmt.Sprintf("%v", DF.Sets[i].getLength()) + string(dtype[1:])
		fmt.Printf("%5v %15v %20v\t%v\n", i, DF.Sets[i].getName(), DType, data)
	}
}

/*

// Reorder : Reorders the DataFrame to the order specified in the input
func (DF *DataFrame) Reorder(ser ...string) {
	indicies := DF.getIndicies(ser)

}

// Convert : converts a series of one type to a series of another type
func (DF *DataFrame) Convert(index int, dtype string) {

}

// Transpose : tranposes the dataset... probably not very useful unless you load in the data manually
func (DF *DataFrame) Transpose() {

}

// Restructure : Restructures the dataset with specified column
func (DF *DataFrame) Restructure(by string) {

}

// Describe : Returns another DataFrame with a full set of descriptive statistics of the dataset
func (DF *DataFrame) Describe() DataFrame {
  return DataFrame{}
}

type Group struct {
	Groups []DataFrame
}

// GroupBy : Returns a slice of DataFrames each one grouped by the given parameter
func (DF *DataFrame) GroupBy(groupBy string) []DataFrame {
  return []DataFrame{}
}



*/
