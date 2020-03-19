package space

// Planet is a string containing a planet's name
type Planet string

// Age is exported
func Age(seconds float64, planet Planet) float64 {
	earthYrs := seconds / 31557600
	switch planet {
	case "Mercury":
		return earthYrs / 0.2408467
	case "Venus":
		return earthYrs / 0.61519726
	case "Earth":
		return earthYrs
	case "Mars":
		return earthYrs / 1.8808158
	case "Jupiter":
		return earthYrs / 11.862615
	case "Saturn":
		return earthYrs / 29.447498
	case "Uranus":
		return earthYrs / 84.016846
	case "Neptune":
		return earthYrs / 164.79132
	default:
		return -1
	}
}
