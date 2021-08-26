package space

import "fmt"

type Planet string

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

func Age(age float64, planet Planet) float64 {

	earthyrs := age * (1 / float64(31557600))
	switch planet {
	case Earth:
		return earthyrs
	case Mercury:
		return earthyrs / 0.2408467
	case Venus:
		return earthyrs / 0.61519726
	case Mars:
		return earthyrs / 1.8808158
	case Jupiter:
		return earthyrs / 11.862615
	case Saturn:
		return earthyrs / 29.447498
	case Uranus:
		return earthyrs / 84.016846
	case Neptune:
		return earthyrs / 164.79132
	default:
		fmt.Print("Sorry Can't figure out the planet name ", planet)
		return -1
	}

}
