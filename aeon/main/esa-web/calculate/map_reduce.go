package calculate

import (
	"sort"
	"strconv"
	"strings"

	"aeon/utils"
)

// MRUnit A Simple MapReduce
type MRUnit struct {
	Param  string
	Error  error
	Result []string

	MapRes    []tuple2
	ReduceRes []string
}

type tuple2 struct {
	_1 string
	_2 int64
}

func (mr *MRUnit) Map() (mrResult *MRUnit) {
	dataSlice := strings.Split(mr.Param, ",")

	for _, s := range dataSlice {
		tp := strings.Split(s, ":")
		if len(tp) < 2 {
			continue
		}
		value, err := strconv.ParseInt(tp[len(tp)-1], 10, 64)
		if err != nil {
			mr.Error = err
			return mr
		}
		mr.MapRes = append(mr.MapRes, tuple2{utils.StringSliceBuilder(tp[:len(tp)-1]), value})
	}

	return mr
}

// 实现自定义map

func (mr *MRUnit) Reduce() (mrResult *MRUnit) {
	if mr.Error != nil {
		return mr
	}

	mapMiddle := make(map[string]int64)
	for _, re := range mr.MapRes {
		mapMiddle[re._1] = mapMiddle[re._1] + re._2
	}

	for k, v := range mapMiddle {
		mr.Result = append(mr.Result, k+":"+strconv.FormatInt(v, 10))
	}

	// sort.Strings(mr.Result)
	sort.Slice(mr.Result, func(i, j int) bool {
		iSlice := strings.Split(mr.Result[i], ":")
		jSlice := strings.Split(mr.Result[j], ":")
		iNum, _ := strconv.Atoi(iSlice[len(iSlice)-1])
		jNum, _ := strconv.Atoi(jSlice[len(jSlice)-1])
		return iNum > jNum
	})

	return mr
}
