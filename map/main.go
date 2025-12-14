package main

import (
	"fmt"
	
	
	"sort"
)

var Database = map[string]string {
	"Серёга": "Омск",
    "Соня":   "Москва",
    "Дима":   "Челябинск",
    "Алина":  "Красноярск",
    "Егор":   "Пермь",
    "Коля":   "Екатеринбург",
}

func main() {
	fmt.Println("Привет, я Алиса!")
    askAlice("Сколько у меня друзей?")
    askAlice("Кто все мои друзья?")
    askAlice("Где все мои друзья?")
}

func askAlice(query string) {
	if query == "Сколько у меня друзей?" {
		count := len(Database)
		fmt.Println("У тебя", count, "друзей")
		return
	}

	if query == "Кто все мои друзья?" {
		friendString := ""

		sortedFriends := getSortedMap(Database)

		for _, v := range sortedFriends{
			friendString += v + " "
		}

		fmt.Println("Твои друзья:", friendString)
		return
	}

	if query == "Где все мои друзья?" {
		ci := []string{}
		city := ""

		for _, v := range Database {
			ci = append(ci, v)
		}

		m := make(map[string]int)
		sl := []string{}

		for _, v := range ci {
			m[v]++
		}
		
		for k := range m {
			sl = append(sl, k)
		}
		
		for _, v := range sl {
			city += v + " "
		}

		fmt.Printf("Твои друзья в городах: %s", city)
		return
	}
}

func getSortedMap(m map[string]string) []string {
	keys := make([]string, 0)

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

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

