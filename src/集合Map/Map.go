package main

import "fmt"

/* 集合Map 无序键值对 */

// 声明集合, 默认 map 是 nil
// var map_varibale map[key_data_type]value_data_type
// 使用 make 函数
// map_varibale := make(map[key_data_type]value_data_type)
// 如果不初始化 map ,那么会创建一个 nil map, nil map不能用来存键值对

func main() {
	var map1 map[string]string
	map1 = make(map[string]string)

	map1["a"] = "aaa"
	map1["b"] = "bbb"
	map1["c"] = "ccc"

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	// 查看元素是否在集合内
	key, value := map1["d"] // value: ture: 存在  false: 不存在
	fmt.Println(key, value)

	// 删除元素
	delete(map1, "a")
	fmt.Println(map1)
}
