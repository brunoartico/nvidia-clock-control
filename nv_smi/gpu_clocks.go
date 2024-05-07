package nv_smi

import (
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
)

type GPUClocks struct {
	memory mapset.Set[int]
	MinimumMemoryClock int

	core mapset.Set[int]
	MinimumCoreClock int
}

func (clocks *GPUClocks) Memory() []int {
	return setToSortedSlice(clocks.memory)
}

func (clocks *GPUClocks) Core() []int {
	return setToSortedSlice(clocks.core)
}

func setToSortedSlice(aSet mapset.Set[int]) []int {
	aSlice := aSet.ToSlice()
	sort.Sort(sort.Reverse(sort.IntSlice(aSlice)))
	return aSlice
}