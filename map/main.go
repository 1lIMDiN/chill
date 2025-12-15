package main

import (
	"fmt"
	//"strings"
	//"sort"
)

// ---------------------------------------------------------------------
// Инвертированная map

func main() {
	m := map[int]string {
		1: "a",
		2: "b",
		3: "a",
		4: "d",
	}

	swapped := swapMap(m)
	fmt.Println(swapped)
}

func swapMap(swap map[int]string) map[string][]int {
	swapped := make(map[string][]int)

	for key, value := range swap {
		swapped[value] = append(swapped[value], key)
	}

	return swapped
}
  
// ---------------------------------------------------------------------
//
// //	Группировка по категории

// type Product struct {
// 	ID       int
// 	Name     string
// 	Category string
// 	Price    float64
// }

// func main() {
// 	products := []Product{
// 		{ID: 1, Name: "Ноутбук", Category: "Электроника", Price: 50000},
// 		{ID: 2, Name: "Мышь", Category: "Электроника", Price: 1500},
// 		{ID: 3, Name: "Стул", Category: "Мебель", Price: 8000},
// 		{ID: 4, Name: "Стол", Category: "Мебель", Price: 12000},
// 		{ID: 5, Name: "Наушники", Category: "Электроника", Price: 3000},
// 		{ID: 6, Name: "Книга", Category: "Книги", Price: 500},
// 		{ID: 7, Name: "Коврик", Category: "Мебель", Price: 2000},
// 	}

// 	grouped := groupByCategory(products)

// 	print(grouped)
// }

// func groupByCategory(products []Product) map[string][]Product {
// 	group := make(map[string][]Product) 

// 	for _, product := range products {
// 		group[product.Category] = append(group[product.Category], product)
// 	}

// 	return group
// }

// func print(group map[string][]Product) {
// 	for category, products := range group {
// 		fmt.Printf("Категория: %s\n", category)
// 		fmt.Printf("Количество товара: %d\n", len(products))
// 		fmt.Println("Товары: ")
// 		for _, product := range products {
// 			fmt.Printf(" -%s (ID: %d, Price: %.2f)\n",
// 		product.Name, product.ID, product.Price)
// 		}
// 		fmt.Println("---")
// 	}
// }

// ---------------------------------------------------------------------
//
// //	Фильтр по значению
// func main() {
// 	m := map[string]int{
// 		"A": 2,
// 		"B": 2,
// 		"C": 3,
// 		"D": 5,
// 		"E": 1,
// 		"F": 2,
// 	}
// 	value := 5

// 	result := mapFilter(m, value)

// 	for k, v := range result {
// 		fmt.Printf("%s: %d\n", k, v)
// 	}
// }

// func mapFilter(m map[string]int, value int) map[string]int {
// 	count := 0
// 	for _, v := range m {
// 		if v >= value {
// 			count++
// 		}
// 	}

// 	result := make(map[string]int, count)

// 	for k, v := range m {
// 		if v >= value {
// 			result[k] = v
// 		}
// 	}

// 	return result
// }

//---------------------------------------------------------------------
// // 							Объединение map
// func main() {
// 	words1 := map[string]int{
// 		"Пират": 1,
// 		"Кошка": 1,
// 		"Пакет": 3,
// 		"Окунь": 5,
// 		"Ухо":   2,
// 	}
// 	words2 := map[string]int{
// 		"Пират": 2,
// 		"Пакет": 3,
// 		"Гошка": 4,
// 		"Кошка": 1,
// 	}

// 	result := mapCombining(words1, words2)

// 	for word, num := range result {
// 		fmt.Printf("%s: %d\n", word, num)
// 	}
// }

// func mapCombining(map1, map2 map[string]int) map[string]int {
// 	wordCount := make(map[string]int)

// 	for key, val := range map1 {
// 		wordCount[key] = val
// 	}

// 	for k, v := range map2 {
// 		wordCount[k] += v
// 	}

// 	return wordCount
// }
//---------------------------------------------------------------------
// // 							Счетчик слов
// func main() {
// 	txt := "Привет мир, привет Go! Go - это здорово."

// 	result := wordCounter(txt)

// 	for word, num := range result {
// 		fmt.Printf("%s: %d\n", word, num)
// 	}
// }

// func wordCounter(text string) map[string]int {
// 	wordCount := make(map[string]int)

// 	words := strings.Fields(text)

// 	for _, word := range words {
// 		word = strings.ToLower(strings.Trim(word, "-.,!?;:\"'()[]{}"))

// 		if word != "" {
// 			wordCount[word]++
// 		}
// 	}

// 	return wordCount
// }
//---------------------------------------------------------------------
// var Database = map[string]string {
// 	"Серёга": "Омск",
//     "Соня":   "Москва",
//     "Дима":   "Челябинск",
//     "Алина":  "Красноярск",
//     "Егор":   "Пермь",
//     "Коля":   "Екатеринбург",
// }

// func main() {
// 	fmt.Println("Привет, я Алиса!")
//     askAlice("Сколько у меня друзей?")
//     askAlice("Кто все мои друзья?")
//     askAlice("Где все мои друзья?")
// }

// func askAlice(query string) {
// 	if query == "Сколько у меня друзей?" {
// 		count := len(Database)
// 		fmt.Println("У тебя", count, "друзей")
// 		return
// 	}

// 	if query == "Кто все мои друзья?" {
// 		friendString := ""

// 		sortedFriends := getSortedMap(Database)

// 		for _, v := range sortedFriends{
// 			friendString += v + " "
// 		}

// 		fmt.Println("Твои друзья:", friendString)
// 		return
// 	}

// 	if query == "Где все мои друзья?" {
// 		ci := []string{}
// 		city := ""

// 		for _, v := range Database {
// 			ci = append(ci, v)
// 		}

// 		m := make(map[string]int)
// 		sl := []string{}

// 		for _, v := range ci {
// 			m[v]++
// 		}

// 		for k := range m {
// 			sl = append(sl, k)
// 		}

// 		for _, v := range sl {
// 			city += v + " "
// 		}

// 		fmt.Printf("Твои друзья в городах: %s", city)
// 		return
// 	}
// }

// func getSortedMap(m map[string]string) []string {
// 	keys := make([]string, 0)

// 	for k := range m {
// 		keys = append(keys, k)
// 	}

// 	sort.Strings(keys)

// 	return keys
// }
//---------------------------------------------------------------------
// // Task 5
// func main() {
// 	favoriteSongs := map[string][]string{
// 		"Серёга": {"Unforgiven", "Holiday", "Highway to Hell"},
// 		"Соня":   {"Shake It Out", "The Show Must Go On", "Наше лето"},
// 		"Дима":   {"Владимирский централ", "Мурка", "Третье сентября"},
// 	}

// 	count := favoriteSongs["Серёга"]
// 	fmt.Println(len(count))

// 	sonya := favoriteSongs["Соня"]
// 	for _, v := range sonya {
// 		fmt.Println(v)
// 	}
// }
