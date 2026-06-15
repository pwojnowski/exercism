package lasagnamaster

const defaultPreparationMinutes = 2

func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = defaultPreparationMinutes
	}
	return len(layers) * minutes
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled = make([]float64, len(quantities))
	for i, q := range quantities {
		scaled[i] = (q / float64(2)) * float64(portions)
	}
	return scaled
}
