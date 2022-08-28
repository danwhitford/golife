package steppers

func ConwayStep(neighbours int, alive bool) bool {
	if alive && (neighbours == 2 || neighbours == 3) {
		return true
	} else if !alive && neighbours == 3 {
		return true
	} else {
		return false
	}
}

func HighLifeStep(neighbours int, alive bool) bool {
	if alive && (neighbours == 2 || neighbours == 3) {
		return true
	} else if !alive && (neighbours == 3 || neighbours == 6) {
		return true
	} else {
		return false
	}
}

func LiveFreeOrDieStep(neighbours int, alive bool) bool {
	if alive && neighbours == 0 {
		return true
	} else if !alive && neighbours == 2 {
		return false
	} else {
		return false
	}
}
