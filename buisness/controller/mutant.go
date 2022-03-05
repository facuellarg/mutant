package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	MutantControllerI interface {
		IsMutant(ctx echo.Context) error
	}

	mutantController struct{}

	//directions in x and y  axis
	direction struct {
		I int
		J int
	}

	DNA struct {
		Sequence []string `json:"dna"`
	}
	//enum for directions
	Direction int
)

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

func NewMutantController() *mutantController {
	return &mutantController{}
}

//return stauts 200 if dna body represent a mutant 403 in other case
func (mc *mutantController) IsMutant(ctx echo.Context) error {
	dna := &DNA{}
	if err := ctx.Bind(dna); err != nil {
		return echo.ErrBadRequest
	}
	fmt.Printf("dna: %v\n", dna)
	if !isMutant(dna.Sequence) {
		return ctx.NoContent(http.StatusForbidden)
	}
	return ctx.NoContent(http.StatusOK)
}

//determine if dna sequence belongs to a mutant or not
func isMutant(dna []string) bool {
	for i := range dna {
		for j := range dna[i] {
			//abajo
			for _, d := range directions {

				if IsInRange(i+d.I*(DEPTH-1), 0, len(dna)-1) &&
					IsInRange(j+d.J*(DEPTH-1), 0, len(dna)-1) &&
					search(dna, i, j, d, 1) {
					return true
				}
			}

		}
	}
	return false
}

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
