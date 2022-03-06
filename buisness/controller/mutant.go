package controller

import (
	"math"
	"net/http"
	"tests/entities/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	MutantControllerI interface {
		IsMutant(ctx echo.Context) error
		Stats(ctx echo.Context) error
	}

	mutantController struct {
		dnaRepository repository.DnaRepositoryI
	}

	//directions in x and y  axis
	direction struct {
		I int
		J int
	}

	DNADto struct {
		Sequence []string `json:"dna"`
	}

	StatsDto struct {
		CountHumanDna int     `json:"count_human_dna"`
		CountMutanDna int     `json:"count_mutan_dna"`
		Ratio         float64 `json:"ratio"`
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

func NewMutantController(dnaRepository repository.DnaRepositoryI) *mutantController {
	return &mutantController{dnaRepository}
}

//return stauts 200 if dna body represent a mutant 403 in other case
func (mc *mutantController) IsMutant(ctx echo.Context) error {
	dna := &DNADto{}
	if err := ctx.Bind(dna); err != nil {
		return echo.ErrBadRequest
	}
	Dna := repository.DNA{}
	Dna.Sequence = dna.Sequence
	var statusCode int
	Dna.IsMutant = isMutant(dna.Sequence)
	if Dna.IsMutant {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusForbidden
	}
	Dna.ID = uuid.New().String()
	if err := mc.dnaRepository.Save(Dna); err != nil {
		log.Error(err)
		return err
	}
	return ctx.NoContent(statusCode)
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

func (mc *mutantController) Stats(ctx echo.Context) error {

	humanCount, err := mc.dnaRepository.GetHumansCount()
	if err != nil {
		log.Error(err)
		return err
	}
	mutantCount, err := mc.dnaRepository.GetMutantsCount()
	if err != nil {
		log.Error(err)
		return err
	}
	var stats StatsDto
	totalCount := (float64(mutantCount) + float64(humanCount))
	if totalCount != 0.0 {
		ratio := float64(mutantCount) / totalCount
		ratio = math.Round(ratio*100.0) / 100
		stats = StatsDto{
			CountHumanDna: humanCount,
			CountMutanDna: mutantCount,
			Ratio:         ratio,
		}
	}

	return ctx.JSON(http.StatusOK, stats)
}
