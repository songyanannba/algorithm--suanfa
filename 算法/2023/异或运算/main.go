package main

import "fmt"

// 异或是 二进制 无进位相加
// 满足交换律

func main() {
	// 1
	a := 6
	b := -101
	fmt.Println("异或前", a, b)
	a = a ^ b // a =  6 ^ -101 ,b = -101
	b = a ^ b // a = 6 ^ -101 , b = 6 ^ -101 ^ -101 = 6
	a = a ^ b // a = 6 ^ -101 ^ 6 = -101
	fmt.Println("异或后", a, b)

	// 2
	arr := []int{1, 2, 3}
	fmt.Printf("%v arr \n", arr)
	Swap(arr, 0, 1)
	fmt.Printf("%v arr \n", arr)

	Swap(arr, 0, 0) //异或如果内存地址指向同一个位置 同一个位置异或会成为 0

	// 3 一组数据 只有一个数字是奇数个 求这个数
	slice := []int{1, 1, 2, 2, 3, 3, 3, 3, 9, 4, 4, 788, 788}
	eor := 0
	for i := 0; i < len(slice); i++ {
		eor ^= slice[i]
	}

	fmt.Println("eor == ", eor)
}

func Swap(arr []int, a, b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
}
