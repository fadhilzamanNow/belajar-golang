package lib

func AddAllArray[K int](arr [3]K) K {
	var s K
	for _,v := range arr {
		s += v
	}
	return s
}
