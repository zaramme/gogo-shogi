package bench

func myFuncA(a, b int) int {
	return a * b * 1
}

func myFuncPointer(a, b *int) int {
	return *a * *b * 1
}

func callMyFuncA(a, b int) int {
	return myFuncA(a, b)
}

func myFuncB(a, b int) int {
	return a * b * 2
}

func myFuncC(a, b int) int {
	return a * b * 3
}

func myFuncD(a, b int) int {
	return a * b * 4
}

func myFuncE(a, b int) int {
	return a * b * 5
}

func myFuncF(a, b int) int {
	return a * b * 6
}

func myFuncG(a, b int) int {
	return a * b * 7
}

func myFuncH(a, b int) int {
	return a * b * 8
}

func myFuncI(a, b int) int {
	return a * b * 9
}
func myFuncSwitchB(num, a, b int) int {
	switch num {
	case 1:
		return a * b * 1
	case 2:
		return a * b * 2
	case 3:
		return a * b * 3
	case 4:
		return a * b * 4
	case 5:
		return a * b * 5
	case 6:
		return a * b * 6
	case 7:
		return a * b * 7
	case 8:
		return a * b * 8
	case 9:
		return a * b * 9
	}
	return 0
}

func myFuncIfB(num, a, b int) int {
	if num == 1 {
		return a * b * 1
	}
	if num == 2 {
		return a * b * 2
	}
	if num == 3 {
		return a * b * 3
	}
	if num == 4 {
		return a * b * 4
	}
	if num == 5 {
		return a * b * 5
	}
	if num == 6 {
		return a * b * 6
	}
	if num == 7 {
		return a * b * 7
	}
	if num == 8 {
		return a * b * 8
	}
	return a * b * 9
}

func myFuncSwitch(num, a, b int) int {
	switch num {
	case 1:
		return myFuncA(a, b)
	case 2:
		return myFuncB(a, b)
	case 3:
		return myFuncC(a, b)
	case 4:
		return myFuncD(a, b)
	case 5:
		return myFuncE(a, b)
	case 6:
		return myFuncF(a, b)
	case 7:
		return myFuncG(a, b)
	case 8:
		return myFuncH(a, b)
	case 9:
		return myFuncI(a, b)
	}
	return 0
}
