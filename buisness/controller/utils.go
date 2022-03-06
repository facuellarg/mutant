package controller

//search the mutant sequence in the direction gived
func search(dna []string, i, j int, d direction, acum int) bool {
	if acum == DEPTH {
		return true
	}
	nextI, nextJ := i+d.I, j+d.J

	if dna[i][j] == dna[nextI][nextJ] {
		return search(dna, nextI, nextJ, d, acum+1)
	}

	return false
}

//Determine if the value is inside the range passed
func IsInRange(val, limitDown, limitUp int) bool {
	return val >= limitDown && val <= limitUp
}
