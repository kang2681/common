package convert

import (
	"fmt"
	"math"
	"strconv"
)

func Int64ToInt32(n int64) (int32, error) {
	if n <= math.MaxInt32 {
		return int32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int32 range", n)
}

func Int64ToInt16(n int64) (int16, error) {
	if n <= math.MaxInt16 {
		return int16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int16 range", n)
}

func Int64ToInt8(n int64) (int8, error) {
	if n <= math.MaxInt8 {
		return int8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int8 range", n)
}

func Int64ToUint64(n int64) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int64ToUint32(n int64) (uint32, error) {
	if 0 <= n && n <= math.MaxUint32 {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int64ToUint16(n int64) (uint16, error) {
	if 0 <= n && n <= math.MaxUint16 {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int64ToUint8(n int64) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func Int32ToInt16(n int32) (int16, error) {
	if n <= math.MaxInt16 {
		return int16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int16 range", n)
}

func Int32ToInt8(n int32) (int8, error) {
	if n <= math.MaxInt8 {
		return int8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int8 range", n)
}

func Int32ToUint64(n int32) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int32ToUint32(n int32) (uint32, error) {
	if 0 <= n {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int32ToUint16(n int32) (uint16, error) {
	if 0 <= n && n <= math.MaxUint16 {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int32ToUint8(n int32) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func Int16ToInt8(n int16) (int8, error) {
	if n <= math.MaxInt8 {
		return int8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the int8 range", n)
}

func Int16ToUint64(n int16) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int16ToUint32(n int16) (uint32, error) {
	if 0 <= n {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int16ToUint16(n int16) (uint16, error) {
	if 0 <= n {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int16ToUint8(n int16) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func Int8ToUint64(n int8) (uint64, error) {
	if 0 <= n {
		return uint64(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint64 range", n)
}

func Int8ToUint32(n int8) (uint32, error) {
	if 0 <= n {
		return uint32(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint32 range", n)
}

func Int8ToUint16(n int8) (uint16, error) {
	if 0 <= n {
		return uint16(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint16 range", n)
}

func Int8ToUint8(n int8) (uint8, error) {
	if 0 <= n {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func StringToInt8(str string) (int8, error) {
	n, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(n), nil
}

func StringToInt16(str string) (int16, error) {
	n, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(n), nil
}

func StringToInt(str string) (int, error) {
	n, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func StringToInt32(str string) (int32, error) {
	n, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

func StringToInt64(str string) (int64, error) {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}

func StringToUint8(str string) (uint8, error) {
	n, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(n), nil
}

func StringToUint16(str string) (uint16, error) {
	n, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(n), nil
}

func StringToUint(str string) (uint, error) {
	n, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}

func StringToUint32(str string) (uint32, error) {
	n, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(n), nil
}

func StringToUint64(str string) (uint64, error) {
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(n), nil
}

func StringToFloat32(str string) (float32, error) {
	n, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(n), nil
}

func StringToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func StringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}
