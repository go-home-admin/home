package database

import (
	"fmt"
	"math"
	"strconv"
)

// Int32 强制转换
func Int32(v interface{}) int32 {
	switch val := v.(type) {
	case int32:
		return val
	case int:
		return int32(val)
	case int64:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		num, err := strconv.Atoi(val)
		if err != nil {
			// 无法解析为数字，返回 0
			return 0
		}
		if num > math.MaxInt32 || num < math.MinInt32 {
			// 超出 int32 的范围，返回 0
			return 0
		}
		return int32(num)
	default:
		s := fmt.Sprintf("%d", v)
		return Int32(s)
	}
}

// UInt32 强制转换
func UInt32(v interface{}) uint32 {
	switch val := v.(type) {
	case uint32:
		return val
	case uint:
		return uint32(val)
	case uint64:
		return uint32(val)
	case float32:
		return uint32(val)
	case float64:
		return uint32(val)
	case string:
		num, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return uint32(num)
	default:
		s := fmt.Sprintf("%d", v)
		return UInt32(s)
	}
}
