package utils

import (
	"time"
)

// TimeCost 耗时统计函数
func TimeCost() func(operate func(duration time.Duration)) {
	start := time.Now()
	return func(operate func(duration time.Duration)) {
		// duration := time.Since(start)
		// fmt.Printf("经过时间：%v\n", duration)
		operate(time.Since(start))
	}
}
