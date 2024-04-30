package pkg

func CheckConditions(edges []Edge) (bool, string) {
	// Условие 1: Проверяем, что каждая дуга ориентирована
	for _, edge := range edges {
		if edge.Start == edge.End {
			return false, "Условие 1 не выполнено"
		}
	}

	// Условие 2: Проверяем, что узел V0 является истоком (инвалентность равна нулю)
	hasV0 := false
	for _, edge := range edges {
		if edge.Start == "V0" {
			hasV0 = true
			break
		}
	}
	if !hasV0 {
		return false, "Условие 2 не выполнено"
	}

	// Условие 3: Проверяем, что узел Vn является стоком (аутвалентность равна нулю)
	hasVn := false
	for _, edge := range edges {
		if edge.End == "Vn" {
			hasVn = true
			break
		}
	}
	if !hasVn {
		return false, "Условие 3 не выполнено"
	}

	// Условие 4: Проверяем, что пропускные способности всех дуг больше нуля
	for _, edge := range edges {
		if edge.Weight <= 0 {
			return false, "Условие 4 не выполнено"
		}
	}

	// Все условия выполнены
	return true, "Условия выполнены"
}
