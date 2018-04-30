package aco

import (
	"math"
	"math/rand"

	"github.com/vahriin/BigGraph/lib/model"
)

type Ant struct {
	Taboo       map[uint64]struct{}
	Path        model.Path
	CurrentCity uint64
}

func MakeAnts(pm *PathMatrix, startCity uint64, antAmount int) []Ant {
	ants := make([]Ant, 0, antAmount)
	for antAmount > 0 {
		ant := Ant{
			Taboo: make(map[uint64]struct{}),
			Path: model.Path{
				Points: make([]uint64, 0, len(pm.cities)+1),
				Len:    0.0},
			CurrentCity: startCity,
		}
		ant.Taboo[ant.CurrentCity] = struct{}{}
		ant.Path.Points = append(ant.Path.Points, startCity)
		ants = append(ants, ant)
		antAmount--
	}

	return ants
}

func (ant *Ant) MakePath(pathCh chan<- model.Path, pm *PathMatrix) {
	for len(ant.Taboo) != len(pm.cities) {
		ant.nextCity(pm)
	}

	ant.Path.Len += pm.Distance(ant.CurrentCity, ant.Path.Points[0])
	ant.Path.Points = append(ant.Path.Points, ant.Path.Points[0])
	ant.CurrentCity = ant.Path.Points[0]

	pathCh <- ant.Path.Copy()

	// update pheromones
	for i := 1; i < len(ant.Path.Points); i++ {
		pm.UpdatePheromone(ant.Path.Points[i-1], ant.Path.Points[i], phMult/ant.Path.Len)
	}

	// set path to clear
	ant.Path = model.Path{
		Points: make([]uint64, 0, len(ant.Path.Points)),
		Len:    0,
	}
	ant.Path.Points = append(ant.Path.Points, ant.CurrentCity)

	// clear Taboo
	for city := range ant.Taboo {
		if city != ant.CurrentCity {
			delete(ant.Taboo, city)
		}
	}
}

func (ant *Ant) nextCity(pm *PathMatrix) {
	var denom float64
	for city := range pm.cities {
		if _, ok := ant.Taboo[city]; ok {
			continue
		}

		denom += math.Pow(pm.Pheromone(ant.CurrentCity, city), alpha) *
			math.Pow(pm.Distance(ant.CurrentCity, city), beta)

	}

	citiesOrder := make([]uint64, len(pm.cities)-len(ant.Taboo))
	probabilities := make([]float64, len(pm.cities)-len(ant.Taboo))
	orderIndex := 0
	var probabilitiesSum float64
	for city := range pm.cities {
		if _, ok := ant.Taboo[city]; ok {
			continue
		}

		probabilitiesSum += math.Pow(pm.Pheromone(ant.CurrentCity, city), alpha) *
			math.Pow(pm.Distance(ant.CurrentCity, city), beta) / denom

		citiesOrder[orderIndex] = city
		probabilities[orderIndex] = probabilitiesSum
		orderIndex++
	}

	random := rand.Float64()
	for i, prob := range probabilities {
		if prob > random {
			ant.Path.Points = append(ant.Path.Points, citiesOrder[i])
			dist := pm.Distance(ant.CurrentCity, citiesOrder[i])
			ant.Path.Len += dist
			ant.CurrentCity = citiesOrder[i]
			ant.Taboo[ant.CurrentCity] = struct{}{}
			break
		}
	}
}
