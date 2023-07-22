package fn

func Map[K interface{}, V interface{}](a []K, function func(str K) V) []V {
	out := []V{}

	for _, ele := range a {
		out = append(out, function(ele))
	}

	return out
}

func Filter[K interface{}](arr []K, function func(b K) bool) []K {
	out := []K{}

	for _, ele := range arr {
		if function(ele) {
			out = append(out, ele)
		}
	}
	return out
}
