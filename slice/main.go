package main

import (
	"fmt"
)

// Сравнение слайсов
func main() {
	sl1 := []int{1,2,3,4,5}
	sl2 := []int{1,2,3,4,5}
	
	fmt.Println(SlicesEqual(sl1, sl2))
}

func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// // Удаление элемента по индексу
// func main() {
// 	sl := []int{0,1,2,3,4,5,6,7,8,9} // len 10 / cap 10
// 	index := 5

// 	slIn := RemoveAtIndex(sl, index) // len 9 / cap 10
// 	fmt.Println(slIn)
// }

// func RemoveAtIndex(sl []int, index int) []int {
// 	sl1 := make([]int, len(sl))
// 	copy(sl1, sl)
	
// 	sl1 = append(sl1[:index], sl1[index+1:]...)

// 	return sl1
// }

// // Поиск пересечения 2 слайсов
// func main() {
// 	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 7}
// 	sl2 := []int{1, 9, 9, 9, 9, 1, 3, 5}

// 	dupl := Duplicate(sl1, sl2)

// 	fmt.Println(dupl)
// }

// func Duplicate(sl1, sl2 []int) []int {
// 	m := make(map[int]int)

// 	for _, v := range sl1 {
// 		if _, exists := m[v]; !exists {
// 			m[v]++
// 		}
// 	}

// 	var result []int
// 	m2 := make(map[int]int)

// 	for _, v := range sl2 {
// 		if num, exists := m[v]; num > 0 && exists {
// 			m2[v]++
// 		}
// 	} 
	
// 	for k := range m2 {
// 		result = append(result, k)
// 	}

// 	return result
// }

// // Фильтрация четных чисел
// func main() {
// 	sliceInts := []int{1,2,3,4,5,6,7,8,9,10}

// 	filter := FilterEven(sliceInts)

// 	fmt.Println(filter)
// }

// func FilterEven(slice []int) []int {
// 	var filter []int

// 	for _, v := range slice {
// 		if v % 2 == 0 {
// 			filter = append(filter, v)
// 		}
// 	}

// 	return filter
// }

// // Реверс слайса
// func main() {
// 	sliceInts := []int{1, 2, 3, 4, 5}
// 	sliceStr := []string{"a", "b", "c"}

// 	rev := reverseGeneric(sliceInts)
// 	reverseInPlaceGeneric(sliceStr)

// 	fmt.Println(rev)
// 	fmt.Println(sliceStr)

// 	fmt.Println("str before", sliceStr)
// 	str := reverseGeneric(sliceStr)
// 	fmt.Println("str after", str)
// }

// // Reverse with generic
// func reverseGeneric[T any](sl []T) []T {
// 	reversed := make([]T, len(sl))

// 	for k := range sl {
// 		reversed[k] = sl[len(sl)-1-k]
// 	}

// 	return reversed
// }

// // Generic with reverse in place
// func reverseInPlaceGeneric[T any](sl []T) {
// 	for i, j := 0, len(sl)-1; i <= j; i, j = i+1, j-1 {
// 		sl[i], sl[j] = sl[j], sl[i]
// 	}
// }
