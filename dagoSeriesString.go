package dago

import (
	"fmt"
	"strconv"
)

func (S *Series) describeStringSeries() {
	S.Dstats.Dtype = sType
	S.Dstats.Length = len(S.Sdata)
	S.Dstats.PrettyType = fmt.Sprintf("%T", S.Sdata)
}

func (S *Series) convertStringToInt() {
	S.Idata = []int{}
	tmpNildata := []int{}
	for i, v := range S.Sdata {
		value, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			tmpNildata = append(tmpNildata, i)
		}
		S.Idata = append(S.Idata, int(value))
	}
	S.Nildata = mergeIntSlices(S.Nildata, tmpNildata)
	S.describeIntSeries()
}
