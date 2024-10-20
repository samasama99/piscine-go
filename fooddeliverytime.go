package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}
	menu := map[string]food{
		"burger":  burger,
		"chips":   chips,
		"nuggets": nuggets,
	}
	if menu[order].preptime == 0 {
		return 404
	}
	return menu[order].preptime
}
