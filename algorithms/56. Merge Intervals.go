package algorithms

import (
	"sort"
)

//Given a collection of intervals, merge all overlapping intervals.
//
//For example,
//Given [1,3],[2,6],[8,10],[15,18],
//return [1,6],[8,10],[15,18].

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (in Intervals) Len() int {
	return len(in)
}

func (in Intervals) Less(i, j int) bool {
	x, y := in[i], in[j]
	return x.Start < y.Start
}

func (in Intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Sort(Intervals(intervals))
	var r []Interval
	start, end := intervals[0].Start, intervals[0].End
	for _, v := range intervals {
		if v.Start <= end {
			if v.End > end {
				end = v.End
			}
		} else {
			r = append(r, Interval{start, end})
			start = v.Start
			end = v.End
		}
	}
	r = append(r, Interval{start, end})
	return r
}
