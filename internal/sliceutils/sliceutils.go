package sliceutils

func Transform[X, Y any](in []X, f func(X) Y) []Y {
	out := make([]Y, len(in))
	for i := range in {
		out[i] = f(in[i])
	}

	return out
}
