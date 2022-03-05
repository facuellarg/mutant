package main

import (
	"fmt"
	"log"
	"tests/buisness/controller"
	"tests/buisness/router"

	"github.com/labstack/echo/v4"
)

type direction struct {
	I int
	J int
}

type Direction int

const (
	ABAJO Direction = iota
	ADELANTE
	DIAGONAL_ABAJO
	DIAGONAL_ARRIBA
)

const (
	DEPTH = 4
)

var directions = map[Direction]direction{
	ABAJO:           {1, 0},
	ADELANTE:        {0, 1},
	DIAGONAL_ABAJO:  {1, 1},
	DIAGONAL_ARRIBA: {-1, 1},
}

func main() {

	server := echo.New()
	mutantController := controller.NewMutantController()
	router := router.NewRouter(*server, mutantController)
	log.Fatal(router.Start())

}

func IsMutant(dna []string) (bool, error) {
	if isMutant(dna) {
		return true, nil
	}
	return false, fmt.Errorf("no Mutant")
}

func isMutant(dna []string) bool {
	for i := range dna {
		for j := range dna[i] {
			//abajo
			for _, d := range directions {

				if IsInRange(i+(d.I*(DEPTH-1)), 0, len(dna)-1) &&
					IsInRange(j+(d.J*(DEPTH-1)), 0, len(dna[i])-1) &&
					search(dna, i, j, d, 1) {
					return true
				}
			}

		}
	}
	return false
}

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

func IsInRange(val, limitDown, limitUp int) bool {
	return val >= limitDown && val <= limitUp
}
