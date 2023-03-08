package database

// StrPointer 字符串转指针
func StrPointer(str interface{}) *string {
	switch str.(type) {
	case string:
		s := str.(string)
		return &s
	case *string:
		if str == nil {
			return nil
		}
		s := str.(*string)
		return s
	}
	return nil
}

// IntPointer 转指针
func IntPointer(i interface{}) *int {
	switch i.(type) {
	case int:
		s := i.(int)
		return &s
	case *int:
		if i == nil {
			return nil
		}
		s := i.(*int)
		return s
	}
	return nil
}

func Int64Pointer(i interface{}) *int64 {
	switch i.(type) {
	case int64:
		s := i.(int64)
		return &s
	case *int64:
		if i == nil {
			return nil
		}
		s := i.(*int64)
		return s
	}
	return nil
}

func UInt64Pointer(i interface{}) *uint64 {
	switch i.(type) {
	case uint64:
		s := i.(uint64)
		return &s
	case *uint64:
		if i == nil {
			return nil
		}
		s := i.(*uint64)
		return s
	}
	return nil
}

func UInt32Pointer(i interface{}) *uint32 {
	switch i.(type) {
	case uint32:
		s := i.(uint32)
		return &s
	case *uint32:
		if i == nil {
			return nil
		}
		s := i.(*uint32)
		return s
	}
	return nil
}

func Int32Pointer(i interface{}) *int32 {
	switch i.(type) {
	case int32:
		s := i.(int32)
		return &s
	case *int32:
		if i == nil {
			return nil
		}
		s := i.(*int32)
		return s
	}
	return nil
}
