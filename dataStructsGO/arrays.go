package main

import "fmt"

func main() {

	//The workaround using an array
	var age [5]int = [5]int{1, 3, 5, 8}
	fmt.Println(age)
	names := [3]string{"Mumbi", "Wamaitha", "Mahinda"}
	fmt.Println(names)
	height := [...]int{153, 156, 172, 189, 201} //any length array
	fmt.Println(len(height))
	fmt.Println(height[3]) //Accessing elements in an array
	//Modifying the array
	numbers := [5]int{1, 2, 4, 5, 7}
	numbers[2] = 3
	numbers[4] = 8
	fmt.Println(numbers)

	//Slices
	//1. From an array
	arr := [7]int{1, 2, 3, 5, 8, 4, 6}
	slice := arr[3:6]
	fmt.Println(slice)

	//2. Using make() func
	slice1 := make([]int, 4)
	slice2 := make([]string, 3)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("Cap: ", cap(slice2))

	//3. Using slice literals - Just like an array but Don't specify the size
	slice3 := []int{10, 20, 30, 50}
	fmt.Println(slice3)

	//Slicing slices
	slice4 := []int{10, 20, 30, 50, 40, 60, 90}
	fmt.Println(slice4[1:6])
	fmt.Println(slice4[5:])
	fmt.Println(slice4[:6])
	fmt.Println(slice4[:])

	//Appending to slice
	slice5 := []int{11, 12, 13, 15}
	slice5 = append(slice5, 14, 16, 18, 19)
	fmt.Println(slice5)
	slice5 = append(slice5, slice3...)
	fmt.Println(slice5)
	fmt.Println("Cap: ", cap(slice5))

	//Copying a slice
	src := []int{5, 0, 8, 5, 0, 7}
	dst := []int{2, 2, 0, 9}
	num := copy(src, dst)
	fmt.Println("Copied: ", num)
	fmt.Println(dst)
	fmt.Println(src)
}
