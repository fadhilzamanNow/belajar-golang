package withgenerics

func SumIntOrFloat[K string, V int64 | float64](m map[K]V) V{
	var s V
	for _,v := range m {
		s += v
	}
	return s
}
