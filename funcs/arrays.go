package funcs

import "fmt"

// InStringArray 判断 val 是否在 string 数组中
func InStringArray(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// StringArrayUniqueMerge 求 string 集合的并集, 会去重
func StringArrayUniqueMerge(arr ...[]string) []string {
	var rs []string
	m := make(map[string]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// StringArrayUnique string 数组去重
// 返回没有重复值的新数组
func StringArrayUnique(arr []string) []string {
	m := make(map[string]struct{})
	var rs []string
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// StringArrayIntersect 求 string 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func StringArrayIntersect(arr []string, arrs ...[]string) []string {
	// arr 转map
	m := make(map[string]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[string]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]string, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// StringArrayDiff 求 string 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func StringArrayDiff(arr []string, arrs ...[]string) []string {
	m := make(map[string]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[string]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]string, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// StringArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func StringArrayChunk(arr []string, size int) [][]string {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]string, 0), make([]string, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]string, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// StringArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func StringArrayCombine(keys, values []string) (map[string]string, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[string]string)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func StringArrayCountValues(arr []string) map[string]int {
	m := make(map[string]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// StringArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func StringArrayFill(start, num int, value string) ([]string, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]string, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// StringArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func StringArraySearch(needle string, haystack []string) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InInt8Array 判断 val 是否在 int8 数组中
func InInt8Array(arr []int8, val int8) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Int8ArrayUniqueMerge 求 int8 集合的并集, 会去重
func Int8ArrayUniqueMerge(arr ...[]int8) []int8 {
	var rs []int8
	m := make(map[int8]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Int8ArrayUnique int8 数组去重
// 返回没有重复值的新数组
func Int8ArrayUnique(arr []int8) []int8 {
	m := make(map[int8]struct{})
	var rs []int8
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Int8ArrayIntersect 求 int8 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Int8ArrayIntersect(arr []int8, arrs ...[]int8) []int8 {
	// arr 转map
	m := make(map[int8]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int8]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int8, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int8ArrayDiff 求 int8 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Int8ArrayDiff(arr []int8, arrs ...[]int8) []int8 {
	m := make(map[int8]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int8]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int8, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int8ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Int8ArrayChunk(arr []int8, size int) [][]int8 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]int8, 0), make([]int8, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]int8, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Int8ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Int8ArrayCombine(keys, values []int8) (map[int8]int8, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[int8]int8)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Int8ArrayCountValues(arr []int8) map[int8]int {
	m := make(map[int8]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Int8ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Int8ArrayFill(start, num int, value int8) ([]int8, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]int8, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Int8ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Int8ArraySearch(needle int8, haystack []int8) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InInt16Array 判断 val 是否在 int16 数组中
func InInt16Array(arr []int16, val int16) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Int16ArrayUniqueMerge 求 int16 集合的并集, 会去重
func Int16ArrayUniqueMerge(arr ...[]int16) []int16 {
	var rs []int16
	m := make(map[int16]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Int16ArrayUnique int16 数组去重
// 返回没有重复值的新数组
func Int16ArrayUnique(arr []int16) []int16 {
	m := make(map[int16]struct{})
	var rs []int16
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Int16ArrayIntersect 求 int16 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Int16ArrayIntersect(arr []int16, arrs ...[]int16) []int16 {
	// arr 转map
	m := make(map[int16]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int16]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int16, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int16ArrayDiff 求 int16 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Int16ArrayDiff(arr []int16, arrs ...[]int16) []int16 {
	m := make(map[int16]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int16]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int16, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int16ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Int16ArrayChunk(arr []int16, size int) [][]int16 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]int16, 0), make([]int16, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]int16, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Int16ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Int16ArrayCombine(keys, values []int16) (map[int16]int16, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[int16]int16)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Int16ArrayCountValues(arr []int16) map[int16]int {
	m := make(map[int16]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Int16ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Int16ArrayFill(start, num int, value int16) ([]int16, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]int16, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Int16ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Int16ArraySearch(needle int16, haystack []int16) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InInt32Array 判断 val 是否在 int32 数组中
func InInt32Array(arr []int32, val int32) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Int32ArrayUniqueMerge 求 int32 集合的并集, 会去重
func Int32ArrayUniqueMerge(arr ...[]int32) []int32 {
	var rs []int32
	m := make(map[int32]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Int32ArrayUnique int32 数组去重
// 返回没有重复值的新数组
func Int32ArrayUnique(arr []int32) []int32 {
	m := make(map[int32]struct{})
	var rs []int32
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Int32ArrayIntersect 求 int32 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Int32ArrayIntersect(arr []int32, arrs ...[]int32) []int32 {
	// arr 转map
	m := make(map[int32]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int32]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int32, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int32ArrayDiff 求 int32 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Int32ArrayDiff(arr []int32, arrs ...[]int32) []int32 {
	m := make(map[int32]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int32]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int32, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int32ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Int32ArrayChunk(arr []int32, size int) [][]int32 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]int32, 0), make([]int32, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]int32, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Int32ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Int32ArrayCombine(keys, values []int32) (map[int32]int32, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[int32]int32)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Int32ArrayCountValues(arr []int32) map[int32]int {
	m := make(map[int32]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Int32ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Int32ArrayFill(start, num int, value int32) ([]int32, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]int32, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Int32ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Int32ArraySearch(needle int32, haystack []int32) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InRuneArray 判断 val 是否在 rune 数组中
func InRuneArray(arr []rune, val rune) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// RuneArrayUniqueMerge 求 rune 集合的并集, 会去重
func RuneArrayUniqueMerge(arr ...[]rune) []rune {
	var rs []rune
	m := make(map[rune]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// RuneArrayUnique rune 数组去重
// 返回没有重复值的新数组
func RuneArrayUnique(arr []rune) []rune {
	m := make(map[rune]struct{})
	var rs []rune
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// RuneArrayIntersect 求 rune 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func RuneArrayIntersect(arr []rune, arrs ...[]rune) []rune {
	// arr 转map
	m := make(map[rune]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[rune]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]rune, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// RuneArrayDiff 求 rune 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func RuneArrayDiff(arr []rune, arrs ...[]rune) []rune {
	m := make(map[rune]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[rune]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]rune, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// RuneArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func RuneArrayChunk(arr []rune, size int) [][]rune {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]rune, 0), make([]rune, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]rune, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// RuneArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func RuneArrayCombine(keys, values []rune) (map[rune]rune, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[rune]rune)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func RuneArrayCountValues(arr []rune) map[rune]int {
	m := make(map[rune]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// RuneArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func RuneArrayFill(start, num int, value rune) ([]rune, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]rune, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// RuneArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func RuneArraySearch(needle rune, haystack []rune) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InIntArray 判断 val 是否在 int 数组中
func InIntArray(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// IntArrayUniqueMerge 求 int 集合的并集, 会去重
func IntArrayUniqueMerge(arr ...[]int) []int {
	var rs []int
	m := make(map[int]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// IntArrayUnique int 数组去重
// 返回没有重复值的新数组
func IntArrayUnique(arr []int) []int {
	m := make(map[int]struct{})
	var rs []int
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// IntArrayIntersect 求 int 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func IntArrayIntersect(arr []int, arrs ...[]int) []int {
	// arr 转map
	m := make(map[int]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// IntArrayDiff 求 int 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func IntArrayDiff(arr []int, arrs ...[]int) []int {
	m := make(map[int]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// IntArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func IntArrayChunk(arr []int, size int) [][]int {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]int, 0), make([]int, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]int, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// IntArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func IntArrayCombine(keys, values []int) (map[int]int, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[int]int)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func IntArrayCountValues(arr []int) map[int]int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// IntArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func IntArrayFill(start, num int, value int) ([]int, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]int, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// IntArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func IntArraySearch(needle int, haystack []int) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InInt64Array 判断 val 是否在 int64 数组中
func InInt64Array(arr []int64, val int64) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Int64ArrayUniqueMerge 求 int64 集合的并集, 会去重
func Int64ArrayUniqueMerge(arr ...[]int64) []int64 {
	var rs []int64
	m := make(map[int64]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Int64ArrayUnique int64 数组去重
// 返回没有重复值的新数组
func Int64ArrayUnique(arr []int64) []int64 {
	m := make(map[int64]struct{})
	var rs []int64
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Int64ArrayIntersect 求 int64 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Int64ArrayIntersect(arr []int64, arrs ...[]int64) []int64 {
	// arr 转map
	m := make(map[int64]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int64]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int64, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int64ArrayDiff 求 int64 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Int64ArrayDiff(arr []int64, arrs ...[]int64) []int64 {
	m := make(map[int64]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[int64]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]int64, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Int64ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Int64ArrayChunk(arr []int64, size int) [][]int64 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]int64, 0), make([]int64, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]int64, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Int64ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Int64ArrayCombine(keys, values []int64) (map[int64]int64, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[int64]int64)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Int64ArrayCountValues(arr []int64) map[int64]int {
	m := make(map[int64]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Int64ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Int64ArrayFill(start, num int, value int64) ([]int64, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]int64, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Int64ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Int64ArraySearch(needle int64, haystack []int64) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InUint8Array 判断 val 是否在 uint8 数组中
func InUint8Array(arr []uint8, val uint8) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Uint8ArrayUniqueMerge 求 uint8 集合的并集, 会去重
func Uint8ArrayUniqueMerge(arr ...[]uint8) []uint8 {
	var rs []uint8
	m := make(map[uint8]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Uint8ArrayUnique uint8 数组去重
// 返回没有重复值的新数组
func Uint8ArrayUnique(arr []uint8) []uint8 {
	m := make(map[uint8]struct{})
	var rs []uint8
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Uint8ArrayIntersect 求 uint8 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Uint8ArrayIntersect(arr []uint8, arrs ...[]uint8) []uint8 {
	// arr 转map
	m := make(map[uint8]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint8]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint8, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint8ArrayDiff 求 uint8 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Uint8ArrayDiff(arr []uint8, arrs ...[]uint8) []uint8 {
	m := make(map[uint8]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint8]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint8, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint8ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Uint8ArrayChunk(arr []uint8, size int) [][]uint8 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]uint8, 0), make([]uint8, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]uint8, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Uint8ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Uint8ArrayCombine(keys, values []uint8) (map[uint8]uint8, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[uint8]uint8)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Uint8ArrayCountValues(arr []uint8) map[uint8]int {
	m := make(map[uint8]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Uint8ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Uint8ArrayFill(start, num int, value uint8) ([]uint8, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]uint8, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Uint8ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Uint8ArraySearch(needle uint8, haystack []uint8) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InByteArray 判断 val 是否在 byte 数组中
func InByteArray(arr []byte, val byte) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// ByteArrayUniqueMerge 求 byte 集合的并集, 会去重
func ByteArrayUniqueMerge(arr ...[]byte) []byte {
	var rs []byte
	m := make(map[byte]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// ByteArrayUnique byte 数组去重
// 返回没有重复值的新数组
func ByteArrayUnique(arr []byte) []byte {
	m := make(map[byte]struct{})
	var rs []byte
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// ByteArrayIntersect 求 byte 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func ByteArrayIntersect(arr []byte, arrs ...[]byte) []byte {
	// arr 转map
	m := make(map[byte]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[byte]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]byte, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// ByteArrayDiff 求 byte 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func ByteArrayDiff(arr []byte, arrs ...[]byte) []byte {
	m := make(map[byte]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[byte]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]byte, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// ByteArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func ByteArrayChunk(arr []byte, size int) [][]byte {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]byte, 0), make([]byte, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]byte, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// ByteArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func ByteArrayCombine(keys, values []byte) (map[byte]byte, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[byte]byte)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func ByteArrayCountValues(arr []byte) map[byte]int {
	m := make(map[byte]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// ByteArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func ByteArrayFill(start, num int, value byte) ([]byte, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]byte, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// ByteArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func ByteArraySearch(needle byte, haystack []byte) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InUint16Array 判断 val 是否在 uint16 数组中
func InUint16Array(arr []uint16, val uint16) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Uint16ArrayUniqueMerge 求 uint16 集合的并集, 会去重
func Uint16ArrayUniqueMerge(arr ...[]uint16) []uint16 {
	var rs []uint16
	m := make(map[uint16]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Uint16ArrayUnique uint16 数组去重
// 返回没有重复值的新数组
func Uint16ArrayUnique(arr []uint16) []uint16 {
	m := make(map[uint16]struct{})
	var rs []uint16
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Uint16ArrayIntersect 求 uint16 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Uint16ArrayIntersect(arr []uint16, arrs ...[]uint16) []uint16 {
	// arr 转map
	m := make(map[uint16]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint16]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint16, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint16ArrayDiff 求 uint16 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Uint16ArrayDiff(arr []uint16, arrs ...[]uint16) []uint16 {
	m := make(map[uint16]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint16]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint16, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint16ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Uint16ArrayChunk(arr []uint16, size int) [][]uint16 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]uint16, 0), make([]uint16, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]uint16, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Uint16ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Uint16ArrayCombine(keys, values []uint16) (map[uint16]uint16, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[uint16]uint16)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Uint16ArrayCountValues(arr []uint16) map[uint16]int {
	m := make(map[uint16]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Uint16ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Uint16ArrayFill(start, num int, value uint16) ([]uint16, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]uint16, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Uint16ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Uint16ArraySearch(needle uint16, haystack []uint16) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InUint32Array 判断 val 是否在 uint32 数组中
func InUint32Array(arr []uint32, val uint32) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Uint32ArrayUniqueMerge 求 uint32 集合的并集, 会去重
func Uint32ArrayUniqueMerge(arr ...[]uint32) []uint32 {
	var rs []uint32
	m := make(map[uint32]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Uint32ArrayUnique uint32 数组去重
// 返回没有重复值的新数组
func Uint32ArrayUnique(arr []uint32) []uint32 {
	m := make(map[uint32]struct{})
	var rs []uint32
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Uint32ArrayIntersect 求 uint32 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Uint32ArrayIntersect(arr []uint32, arrs ...[]uint32) []uint32 {
	// arr 转map
	m := make(map[uint32]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint32]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint32, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint32ArrayDiff 求 uint32 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Uint32ArrayDiff(arr []uint32, arrs ...[]uint32) []uint32 {
	m := make(map[uint32]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint32]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint32, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint32ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Uint32ArrayChunk(arr []uint32, size int) [][]uint32 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]uint32, 0), make([]uint32, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]uint32, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Uint32ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Uint32ArrayCombine(keys, values []uint32) (map[uint32]uint32, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[uint32]uint32)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Uint32ArrayCountValues(arr []uint32) map[uint32]int {
	m := make(map[uint32]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Uint32ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Uint32ArrayFill(start, num int, value uint32) ([]uint32, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]uint32, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Uint32ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Uint32ArraySearch(needle uint32, haystack []uint32) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InUintArray 判断 val 是否在 uint 数组中
func InUintArray(arr []uint, val uint) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// UintArrayUniqueMerge 求 uint 集合的并集, 会去重
func UintArrayUniqueMerge(arr ...[]uint) []uint {
	var rs []uint
	m := make(map[uint]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// UintArrayUnique uint 数组去重
// 返回没有重复值的新数组
func UintArrayUnique(arr []uint) []uint {
	m := make(map[uint]struct{})
	var rs []uint
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// UintArrayIntersect 求 uint 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func UintArrayIntersect(arr []uint, arrs ...[]uint) []uint {
	// arr 转map
	m := make(map[uint]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// UintArrayDiff 求 uint 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func UintArrayDiff(arr []uint, arrs ...[]uint) []uint {
	m := make(map[uint]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// UintArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func UintArrayChunk(arr []uint, size int) [][]uint {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]uint, 0), make([]uint, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]uint, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// UintArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func UintArrayCombine(keys, values []uint) (map[uint]uint, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[uint]uint)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func UintArrayCountValues(arr []uint) map[uint]int {
	m := make(map[uint]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// UintArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func UintArrayFill(start, num int, value uint) ([]uint, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]uint, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// UintArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func UintArraySearch(needle uint, haystack []uint) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InUint64Array 判断 val 是否在 uint64 数组中
func InUint64Array(arr []uint64, val uint64) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Uint64ArrayUniqueMerge 求 uint64 集合的并集, 会去重
func Uint64ArrayUniqueMerge(arr ...[]uint64) []uint64 {
	var rs []uint64
	m := make(map[uint64]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Uint64ArrayUnique uint64 数组去重
// 返回没有重复值的新数组
func Uint64ArrayUnique(arr []uint64) []uint64 {
	m := make(map[uint64]struct{})
	var rs []uint64
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Uint64ArrayIntersect 求 uint64 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Uint64ArrayIntersect(arr []uint64, arrs ...[]uint64) []uint64 {
	// arr 转map
	m := make(map[uint64]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint64]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint64, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint64ArrayDiff 求 uint64 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Uint64ArrayDiff(arr []uint64, arrs ...[]uint64) []uint64 {
	m := make(map[uint64]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[uint64]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]uint64, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Uint64ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Uint64ArrayChunk(arr []uint64, size int) [][]uint64 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]uint64, 0), make([]uint64, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]uint64, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Uint64ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Uint64ArrayCombine(keys, values []uint64) (map[uint64]uint64, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[uint64]uint64)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Uint64ArrayCountValues(arr []uint64) map[uint64]int {
	m := make(map[uint64]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Uint64ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Uint64ArrayFill(start, num int, value uint64) ([]uint64, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]uint64, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Uint64ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Uint64ArraySearch(needle uint64, haystack []uint64) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InFloat32Array 判断 val 是否在 float32 数组中
func InFloat32Array(arr []float32, val float32) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Float32ArrayUniqueMerge 求 float32 集合的并集, 会去重
func Float32ArrayUniqueMerge(arr ...[]float32) []float32 {
	var rs []float32
	m := make(map[float32]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Float32ArrayUnique float32 数组去重
// 返回没有重复值的新数组
func Float32ArrayUnique(arr []float32) []float32 {
	m := make(map[float32]struct{})
	var rs []float32
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Float32ArrayIntersect 求 float32 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Float32ArrayIntersect(arr []float32, arrs ...[]float32) []float32 {
	// arr 转map
	m := make(map[float32]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[float32]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]float32, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Float32ArrayDiff 求 float32 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Float32ArrayDiff(arr []float32, arrs ...[]float32) []float32 {
	m := make(map[float32]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[float32]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]float32, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Float32ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Float32ArrayChunk(arr []float32, size int) [][]float32 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]float32, 0), make([]float32, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]float32, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Float32ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Float32ArrayCombine(keys, values []float32) (map[float32]float32, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[float32]float32)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Float32ArrayCountValues(arr []float32) map[float32]int {
	m := make(map[float32]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Float32ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Float32ArrayFill(start, num int, value float32) ([]float32, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]float32, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Float32ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Float32ArraySearch(needle float32, haystack []float32) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}

// InFloat64Array 判断 val 是否在 float64 数组中
func InFloat64Array(arr []float64, val float64) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Float64ArrayUniqueMerge 求 float64 集合的并集, 会去重
func Float64ArrayUniqueMerge(arr ...[]float64) []float64 {
	var rs []float64
	m := make(map[float64]struct{})
	for _, v := range arr {
		for _, vv := range v {
			if _, ok := m[vv]; ok {
				continue
			}
			rs = append(rs, vv)
		}
	}
	return rs
}

// Float64ArrayUnique float64 数组去重
// 返回没有重复值的新数组
func Float64ArrayUnique(arr []float64) []float64 {
	m := make(map[float64]struct{})
	var rs []float64
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		rs = append(rs, v)
	}
	return rs
}

// Float64ArrayIntersect 求 float64 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func Float64ArrayIntersect(arr []float64, arrs ...[]float64) []float64 {
	// arr 转map
	m := make(map[float64]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[float64]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; !ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]float64, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Float64ArrayDiff 求 float64 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func Float64ArrayDiff(arr []float64, arrs ...[]float64) []float64 {
	m := make(map[float64]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arrs {
		// 转map
		tmp := make(map[float64]struct{})
		for _, val := range v {
			tmp[val] = struct{}{}
		}
		for k, _ := range m {
			if _, ok := tmp[k]; ok {
				delete(m, k)
			}
		}
	}
	// 保留原来的顺序
	rs := make([]float64, 0, len(m))
	for _, v := range arr {
		if _, ok := m[v]; ok {
			rs = append(rs, v)
		}
	}
	return rs
}

// Float64ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func Float64ArrayChunk(arr []float64, size int) [][]float64 {
	if len(arr) == 0 || size < 0 {
		return nil
	}
	rs, tmp := make([][]float64, 0), make([]float64, 0, size)
	for _, v := range arr {
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]float64, 0, size)
		} else {
			tmp = append(tmp, v)
		}
	}
	return rs
}

// Float64ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func Float64ArrayCombine(keys, values []float64) (map[float64]float64, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys 和 values 长度不同")
	}
	m := make(map[float64]float64)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// 统计数组中所有的值
func Float64ArrayCountValues(arr []float64) map[float64]int {
	m := make(map[float64]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

// Float64ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func Float64ArrayFill(start, num int, value float64) ([]float64, error) {
	if start < 0 {
		return nil, fmt.Errorf("start Must Gte 0")
	}
	if num < 0 {
		return nil, fmt.Errorf("num Must Gte 0")
	}
	length := start + num
	rs := make([]float64, length)
	for i := 0; i < num; i++ {
		rs[start+i] = value
	}
	return rs, nil
}

// Float64ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func Float64ArraySearch(needle float64, haystack []float64) (int, bool) {
	for k, v := range haystack {
		if v == needle {
			return k, true
		}
	}
	return 0, false
}
