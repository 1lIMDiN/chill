package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	K1 = 0.035
	K2 = 0.029
)

var (
	Format     = "20060102 15:04:05" // формат даты и времени
	StepLength = 0.65                // длина шага в метрах
	Weight     = 75.0                // вес кг
	Height     = 1.75                // рост м
	Speed      = 1.39                // скорость м/с
)

// parsePackage разбирает входящий пакет в параметре data.
// Возвращаемые значения:
// t — дата и время, указанные в пакете
// steps — количество шагов
// ok — true, если время и шаги указаны корректно, и false — в противном случае
func parsePackage(data string) (t time.Time, steps int, ok bool) {
	ds := strings.Split(data, ",")
	if len(ds) != 2 {
		return
	}
	var err error

	t, err = time.Parse(Format, ds[0])
	if err != nil {
		return
	}

	steps, err = strconv.Atoi(ds[1])
	if err != nil || steps < 0 {
		return
	}

	ok = true
	return
}

// stepsDay перебирает все записи слайса, подсчитывает и возвращает
// общее количество шагов
func stepsDay(storage []string) int {
	// тема оптимизации не затрагивается, поэтому можно
	// использовать parsePackage для каждого элемента списка
	count := 0
	for _, v := range storage {
		_, steps, _ := parsePackage(v)
		count += steps
	}

	return count
}

// calories возвращает количество килокалорий, которые потрачены на
// прохождение указанной дистанции (в метрах) со скоростью 5 км/ч
func calories(distance float64) float64 {
	result := K1*Weight + (Speed*Speed/Height)*K2*Weight
	timeInGo := distance * 1000 / Speed / 60
	return result * timeInGo
}

// achievement возвращает мотивирующее сообщение в зависимости от
// пройденного расстояния в километрах
func achievement(distance float64) string {
	if distance >= 6.5 {
		return "Отличный результат! Цель достигнута."
	} else if distance >= 3.9 {
		return "Неплохо! День был продуктивный."
	} else if distance >= 2 {
		return "Завтра наверстаем!"
	} else {
		return "Лежать тоже полезно. Главное — участие, а не победа!"
	}
}

func showMessage(s string) {
	fmt.Print(s, "\n\n")
}

// AcceptPackage обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func AcceptPackage(data string, storage []string) []string {
	//    Используйте parsePackage для разбора пакета
	//    выведите сообщение в случае ошибки
	//    также проверьте количество шагов на равенство нулю
	t, steps, ok := parsePackage(data)
	if !ok || steps == 0 {
		showMessage(`ошибочный формат пакета`)
		return storage
	}
	//	  Получите текущее UTC-время и сравните дни
	//    выведите сообщение, если день в пакете t.Day() не совпадает
	//    с текущим днём
	now := time.Now().UTC()
	if t.Day() != now.Day() {
		showMessage(`неверный день`)
		return storage
	}

	// выводим ошибку, если время в пакете больше текущего времени
	if t.After(now) {
		showMessage(`некорректное значение времени`)
		return storage
	}
	// проверки для непустого storage
	if len(storage) > 0 {
		num := len(Format)
		if data[:num] <= storage[len(storage)-1][:num] {
			showMessage(`некорректное значение времени`)
			return storage
		}

		// смотрим, наступили ли новые сутки: YYYYMMDD — 8 символов
		if data[:8] != storage[len(storage)-1][:8] {
			// если наступили,
			// то обнуляем слайс с накопленными данными
			storage = storage[:0]
		}
	}
	// Добавить пакет в storage
	storage = append(storage, data)

	// Получить общее количество шагов
	allSteps := stepsDay(storage)

	// Вычислить общее расстояние (в метрах)
	distance := float64(allSteps) * StepLength / 1000

	// Получить потраченные килокалории
	energy := calories(distance)

	// Получить мотивирующий текст
	achiev := achievement(distance)

	// Сформировать и вывести полный текст сообщения
	msg := fmt.Sprintf(`Время: %s.
Количество шагов за сегодня: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
%s`, t.Format("15:04:05"), allSteps, distance, energy, achiev)
	showMessage(msg)

	return storage
}

func main() {
	now := time.Now().UTC()
	today := now.Format("20060102")

	input := []string{
		"01:41:03,-100",
		",3456",
		"12:40:00, 3456 ",
		"something is wrong",
		"02:11:34,678",
		"02:11:34,792",
		"17:01:30,1078",
		"03:25:59,7830",
		"04:00:46,5325",
		"04:45:21,3123",
	}

	var storage []string
	storage = AcceptPackage("20230720 00:11:33,100", storage)
	for _, v := range input {
		storage = AcceptPackage(today+" "+v, storage)
	}
}
