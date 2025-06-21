//go:build !solution

package hotelbusiness

import (
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	if len(guests) == 0 {
		return []Load{}
	}

	// Записываем изменение гостей по дням
	delta := make(map[int]int)
	for _, guest := range guests {
		delta[guest.CheckInDate]  += 1
		delta[guest.CheckOutDate] -= 1
	}

	// Создаём сортированный список дней
	dates := make([]int, 0, len(delta))
	for day := range delta {
		dates = append(dates, day)
	}
	sort.Ints(dates)

	cur_guests := 0
	result := make([]Load, 0)
	for _, day := range dates {
		cur_guests += delta[day]
		if (len(result) == 0) || (cur_guests != result[len(result) - 1].GuestCount) {
			result = append(result, Load{day, cur_guests})
		}
	}

	return result
}
