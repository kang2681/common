package funcs

import "fmt"

{{range .data}}
// In{{.FuncName}}Array 判断 val 是否在 {{.Name}} 数组中
func In{{.FuncName}}Array(arr []{{.Name}}, val {{.Name}}) bool {
    for _, v := range arr {
        if v == val {
            return true
        }
    }
    return false
}

// {{.FuncName}}ArrayUniqueMerge 求 {{.Name}} 集合的并集, 会去重
func {{.FuncName}}ArrayUniqueMerge(arr ...[]{{.Name}}) []{{.Name}} {
    var rs []{{.Name}}
    m := make(map[{{.Name}}]struct{})
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

// {{.FuncName}}ArrayUnique {{.Name}} 数组去重
// 返回没有重复值的新数组
func {{.FuncName}}ArrayUnique(arr []{{.Name}}) []{{.Name}} {
    m := make(map[{{.Name}}]struct{})
    var rs []{{.Name}}
    for _, v := range arr {
        if _, ok := m[v]; ok {
            continue
        }
        rs = append(rs, v)
    }
    return rs
}

// {{.FuncName}}ArrayIntersect 求 {{.Name}} 集合的交集
// arr 要检查的数组，作为主值。
// arrs 要对比的数组列表。
// 返回一个数组，该数组包含了所有在 arr 中也同时出现在所有其它参数数组中的值。
func {{.FuncName}}ArrayIntersect(arr []{{.Name}}, arrs ...[]{{.Name}}) []{{.Name}} {
    // arr 转map
    m := make(map[{{.Name}}]struct{})
    for _, v := range arr {
        m[v] = struct{}{}
    }
    for _, v := range arrs {
        // 转map
        tmp := make(map[{{.Name}}]struct{})
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
    rs := make([]{{.Name}}, 0, len(m))
    for _, v := range arr {
        if _, ok := m[v]; ok {
            rs = append(rs, v)
        }
    }
    return rs
}

// {{.FuncName}}ArrayDiff 求 {{.Name}} 数组差集
// arr1 要被对比的数组
// arrs 和这个数组进行比较
// 对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func {{.FuncName}}ArrayDiff(arr []{{.Name}}, arrs ...[]{{.Name}}) []{{.Name}} {
    m := make(map[{{.Name}}]struct{})
    for _, v := range arr {
        m[v] = struct{}{}
    }
    for _, v := range arrs {
        // 转map
        tmp := make(map[{{.Name}}]struct{})
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
    rs := make([]{{.Name}}, 0, len(m))
    for _, v := range arr {
        if _, ok := m[v]; ok {
            rs = append(rs, v)
        }
    }
    return rs
}

// {{.FuncName}}ArrayChunk 将一个数组分割成多个
// arr 需要操作的数组
// size 每个数组的单元数目
//将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个。
func {{.FuncName}}ArrayChunk(arr []{{.Name}}, size int) [][]{{.Name}} {
    if len(arr) == 0 || size < 0 {
        return nil
    }
    rs, tmp := make([][]{{.Name}}, 0), make([]{{.Name}}, 0, size)
    for _, v := range arr {
        if len(tmp) == size {
            rs = append(rs, tmp)
            tmp = make([]{{.Name}}, 0, size)
        } else {
            tmp = append(tmp, v)
        }
    }
    return rs
}

// {{.FuncName}}ArrayCombine 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
// keys 将被作为map的键
// values 将被作为map的值
func {{.FuncName}}ArrayCombine(keys, values []{{.Name}}) (map[{{.Name}}]{{.Name}}, error) {
    if len(keys) != len(values) {
        return nil, fmt.Errorf("keys 和 values 长度不同")
    }
    m := make(map[{{.Name}}]{{.Name}})
    for k, v := range keys {
        m[v] = values[k]
    }
    return m, nil
}

// 统计数组中所有的值
func {{.FuncName}}ArrayCountValues(arr []{{.Name}}) map[{{.Name}}]int {
    m := make(map[{{.Name}}]int)
    for _, v := range arr {
        m[v]++
    }
    return m
}

// {{.FuncName}}ArrayFill 用给定的值填充数组
// start 数组的开始索引值，必须大于等于0
// num 插入元素的数量 必须大于等于0
// 如果start 大于0开始，0~ start 将使用空字符串填充
func {{.FuncName}}ArrayFill(start, num int, value {{.Name}}) ([]{{.Name}}, error) {
    if start < 0 {
        return nil, fmt.Errorf("start Must Gte 0")
    }
    if num < 0 {
        return nil, fmt.Errorf("num Must Gte 0")
    }
    length := start + num
    rs := make([]{{.Name}}, length)
    for i := 0; i < num; i++ {
        rs[start+i] = value
    }
    return rs, nil
}

// {{.FuncName}}ArraySearch 在数组haystack中搜索给定的值，如果成功则返回首个相应的下标
func {{.FuncName}}ArraySearch(needle {{.Name}}, haystack []{{.Name}}) (int, bool) {
    for k, v := range haystack {
        if v == needle {
            return k, true
        }
    }
    return 0, false
}
{{end}}