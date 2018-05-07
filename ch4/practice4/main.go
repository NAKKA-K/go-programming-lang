// 渡された配列を1つずらして、ローテーションさせる関数
package main

import "fmt"

func rotate(slice []int, isRightRotate bool) []int {
	var out []int

	if isRightRotate {
		out = append(out, slice[len(slice)-1])
		out = appendIntSlice(out, slice[:len(slice)-1])
		return out
	}

	out = appendIntSlice(out, slice[1:])
	out = append(out, slice[0])
	return out
}

func appendIntSlice(slice []int, appendSlice []int) []int {
	out := slice
	for _, v := range appendSlice {
		out = append(out, v)
	}
	return out
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(rotate(slice, true))
	fmt.Println(rotate(slice, false))
}
