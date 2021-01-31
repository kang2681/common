package convert

import (
	"fmt"
	"math"
	"strconv"
)

func Int64ConvertInt32(n int64) (int32, error) {
	if n <= math.MaxInt32 {
		return int32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int32 range", n)
}

func Int64ConvertInt16(n int64) (int16, error) {
	if n <= math.MaxInt16 {
		return int16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int16 range", n)
}

func Int64ConvertInt8(n int64) (int8, error) {
	if n <= math.MaxInt8 {
		return int8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int8 range", n)
}

func Int64ConvertUint64(n int64) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int64ConvertUint32(n int64) (uint32, error) {
	if 0 <= n && n <= math.MaxUint32 {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int64ConvertUint16(n int64) (uint16, error) {
	if 0 <= n && n <= math.MaxUint16 {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int64ConvertUint8(n int64) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func Int32ConvertInt16(n int32) (int16, error) {
	if n <= math.MaxInt16 {
		return int16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int16 range", n)
}

func Int32ConvertInt8(n int32) (int8, error) {
	if n <= math.MaxInt8 {
		return int8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int8 range", n)
}

func Int32ConvertUint64(n int32) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int32ConvertUint32(n int32) (uint32, error) {
	if 0 <= n {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int32ConvertUint16(n int32) (uint16, error) {
	if 0 <= n && n <= math.MaxUint16 {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int32ConvertUint8(n int32) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func Int16ConvertInt8(n int16) (int8, error) {
	if n <= math.MaxInt8 {
		return int8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int8 range", n)
}

func Int16ConvertUint64(n int16) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int16ConvertUint32(n int16) (uint32, error) {
	if 0 <= n {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int16ConvertUint16(n int16) (uint16, error) {
	if 0 <= n {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int16ConvertUint8(n int16) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func Int8ConvertUint64(n int8) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int8ConvertUint32(n int8) (uint32, error) {
	if 0 <= n {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int8ConvertUint16(n int8) (uint16, error) {
	if 0 <= n {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int8ConvertUint8(n int8) (uint8, error) {
	if 0 <= n {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func StringConvertInt8(str string) (int8, error) {
	n, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(n), nil
}

func StringConvertInt16(str string) (int16, error) {
	n, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(n), nil
}

func StringConvertInt(str string) (int, error) {
	n, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func StringConvertInt32(str string) (int32, error) {
	n, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

func StringConvertInt64(str string) (int64, error) {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}

func StringConvertUint8(str string) (uint8, error) {
	n, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(n), nil
}

func StringConvertUint16(str string) (uint16, error) {
	n, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(n), nil
}

func StringConvertUint(str string) (uint, error) {
	n, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}

func StringConvertUint32(str string) (uint32, error) {
	n, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(n), nil
}

func StringConvertUint64(str string) (uint64, error) {
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(n), nil
}

func StringConvertFloat32(str string) (float32, error) {
	n, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(n), nil
}

func StringConvertFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func StringConvertBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}
