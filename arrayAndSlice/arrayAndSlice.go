package arrayandslice

import "fmt"

func ArraySlice() {
	// array -> fix size ต้องจองพื้นที่ก่อน **ตอนเอาเข้าฟังชั่นก์เช่น array[i] *= 2 มันจะ copy ทำให้ค่าใน array ไม่เปลี่ยน
	// เพราะไม่ใช่ตัวแปรตัวเดิมที่ address เดียวกัน
	// ถ้าไม่กำหนดค่าเริ่มต้นจะเท่ากับ [0,0,0]
	var arrayList [3]int
	// 3 is specific index
	// ... ใช้เวลาเราต้องพิม arrays หลาย ๆ ตัวโดยที่ไม่รู้ว่าทั้งหมดมันมีขนาดเท่าไหร่ ก็ไม่ต้องไปไล่นับ ใช้ ... ได้เลย
	arrayList2 := [...]string{"apple", "mango", 3: "orange"}
	arrayList2[2] = "banana"
	fmt.Println(arrayList)
	fmt.Println(arrayList2)
	for _, k := range arrayList2 {
		fmt.Println(k)
	}
	// size
	fmt.Println(len(arrayList2))

	// slice -> dynamic size **ตอนเอาเข้าฟังชั่นก์เช่น slice[i] *= 2 มันจะ copy ทำให้ค่าใน slice เปลี่ยน
	// เพราะเป็นตัวแปรตัวเดิมที่ address เดียวกัน -> pass by reference อันตรายนะครับผม
	//ถ้าไม่กำหนดค่าเริ่มต้นจะเท่ากับ []
	var slice []int
	slice2 := []string{"cat", "dog", "human"}
	slice2 = append(slice2, "bird")
	fmt.Println(slice)
	fmt.Println(slice2)
	for _, k := range slice2 {
		fmt.Println(k)
	}
	// all element
	fmt.Println("All element =", slice2[0:])
	// last element
	fmt.Println("Last element =", slice2[len(slice2)-1])
	// last two element
	fmt.Println("Last two element are =", slice2[len(slice2)-2:])
	// ประกาษ slice เปล่าๆ
	var sliceEmpty []int
	sliceEmpty2 := []int{}
	// make คือการจองพื้นที่ให้กับ slice, map ไว้ก่อนแต่ถ้าเป็น int และขนาดไม่ใช่่ 0 มันจะเหมือน array ตรงที่ให้ [0,0,0] เป็น default
	sliceEmpty3 := make([]int, 3)
	sliceEmpty4 := make([]string, 3)
	fmt.Println(sliceEmpty)
	fmt.Println(sliceEmpty2)
	fmt.Println(sliceEmpty3)
	fmt.Println(sliceEmpty4)

	// ระวังเรื่อง slice pass by reference ตัวให้ค่าเปลี่ยน แก้โดยการ trade of เพิ่มการใช้ memory นิดหน่อย เพื่อสร้าง ตัวมารับ return
	cautionSlice := []int{1, 2, 3, 4}
	fmt.Println("Before pass in function =", cautionSlice)
	cautionSliceMultiplyBy2 := double(cautionSlice)
	fmt.Println("After pass in function =", cautionSlice)
	fmt.Println("New fixer slice =", cautionSliceMultiplyBy2)
}

func double(nums []int) []int {
	// This is wrong, pass by reference -> the orginal caution slice will also change
	// newSlice := nums
	// fmt.Println(&newSlice[0]) same address
	// fmt.Println(&nums[0]) same address

	// แก้แล้วแต่ยัง runtime error
	// newSlice := []int{} -> อันนี้ก็ไม่ได้เพราะไม่ได้จองพื้นที่ไว้ก่อน แก้โดยใช้ make ทำให้ slice จองขนาดพื้นที่คล้าย ๆ array หรือเปลี่ยนเป็น append
	// for i := range nums {
	// 	newSlice[i] = nums[i] * 2
	// }

	// นี่ก็ถูกนะ แต่ชอบแบบล่างมากกว่า
	// newSlice := []int{}
	// for i := range nums {
	// 	newSlice = append(newSlice, nums[i]*2)
	// }

	// ถูกและดี
	newSlice := make([]int, len(nums))
	for i := range nums {
		newSlice[i] = nums[i] * 2
	}
	return newSlice
}
