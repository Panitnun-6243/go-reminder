package mapobject

import "fmt"

func MapObject() {
	// ประกาศแบบให้ค่า
	mapA := map[int]int{
		0: 1,
		1: 3,
		2: 5,
	}
	// ประกาศไว้เฉย ๆ
	mapB := make(map[string]float32)
	mapB["pi"] = 3.1416
	fmt.Println(mapA)
	fmt.Println(mapA[2])
	fmt.Println(mapB["pi"])
	// update map
	mapB["pi"] = 3.14
	fmt.Println(mapB["pi"])
	// delete
	delete(mapB, "pi")
	fmt.Println(mapB["pi"])
	// check if the key exist or not
	if _, ok := mapA[4]; ok {
		fmt.Println("The key is existed")
	} else {
		fmt.Println("The value is not exist")
	}

	// ประยุกต์ใช้กับ graph
	var useGraph = map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
	}
	depthFirstSearch(useGraph, 1, make(map[int]bool))
}

// vertex จุด (node) ที่มัน visit ไปเรื่อย ๆ เริ่ม
func depthFirstSearch(graph map[int][]int, vertex int, visited map[int]bool) {
	visited[vertex] = true
	for _, v := range graph[vertex] {
		fmt.Println(v)
		depthFirstSearch(graph, v, visited)
	}
}
