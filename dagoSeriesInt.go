package dago

import (
	"fmt"
)

func (S *Series) describeIntSeries() {
	S.Dstats.Dtype = iType
	S.Dstats.Length = len(S.Idata)
	S.Dstats.PrettyType = fmt.Sprintf("%T", S.Idata)
}

func (S *Series) convertIntToString() {

}

// merges two slices with the assumption that each slice is ordered
func mergeIntSlices(slice1, slice2 []int) []int {
	outSlice := []int{}
	for {
		if len(slice1) == 0 && len(slice2) == 0 {
			return outSlice
		}
		if len(slice1) > 0 {
			if len(slice2) > 0 {
				if slice1[0] < slice2[0] {
					outSlice = append(outSlice, slice1[0])
					slice1 = slice1[1:]
				} else {
					outSlice = append(outSlice, slice2[0])
					slice2 = slice2[1:]
				}
			} else {
				outSlice = append(outSlice, slice1[0])
				slice1 = slice1[1:]
			}
		} else {
			outSlice = append(outSlice, slice2[0])
			slice2 = slice2[1:]
		}
	}
}
