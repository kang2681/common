package arrays

import (
	"fmt"
	"math"
)

// InStringArray 判断 val 是否在 arr 数组中
func InStringArray(val string, arr []string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// StringArrayUniqueMerge 求 arr 集合的并集, 会去重, 会保留传入的顺序
func StringArrayUniqueMerge(arrs ...[]string) []string {
	var rs []string
	m := make(map[string]struct{})
	for _, arr := range arrs {
		for _, v := range arr {
			if _, ok := m[v]; ok {
				continue
			}
			m[v] = struct{}{}
			rs = append(rs, v)
		}
	}
	return rs
}

// StringArrayUnique arr 数组去重
// 返回没有重复值的新数组，会保留原来的顺序
func StringArrayUnique(arr []string) []string {
	m := make(map[string]struct{})
	var rs []string
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		rs = append(rs, v)
	}
	return rs
}

// StringArrayIntersect 求 arr 集合的交集
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

// StringArrayDiff 求 arr 数组与 arrs 差集
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
// size 每个数组的单元数目 必须大于 0
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func StringArrayChunk(arr []string, size int) [][]string {
	if size < 1 {
		panic("size can not less than 1 ")
	}
	capSize := int(math.Ceil(float64(len(arr)) / float64(size)))
	rs, tmp := make([][]string, 0, capSize), make([]string, 0, size)
	for _, v := range arr {
		tmp = append(tmp, v)
		if len(tmp) == size {
			rs = append(rs, tmp)
			tmp = make([]string, 0, size)
		}
	}
	if len(tmp) > 0 {
		rs = append(rs, tmp)
	}
	return rs
}

// StringArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func StringArrayCombine(keys, values []string) (map[string]string, error) {
	if len(keys) != len(values) {
		return nil, fmt.Errorf("keys and values length Not Equal")
	}
	m := make(map[string]string)
	for k, v := range keys {
		m[v] = values[k]
	}
	return m, nil
}

// StringArrayCountValues 统计数组中所有单元值出现的次数
func StringArrayCountValues(arr []string) map[string]int {
	m := make(map[string]int)
	for _, v := range arr {
		m[v]++
	}
	return m
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
