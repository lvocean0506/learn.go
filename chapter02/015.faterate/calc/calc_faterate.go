package calc

func ClacFateRate(bmi float64, age int, sex string) (fateRate float64) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fateRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageweight float64) {
	ageweight = 0.23
	if age >= 30 && age <= 40 {
		ageweight = 0.22
	}
	return
}
