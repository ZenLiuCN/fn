//go:build !gen

package fn

//generated file should not edit

// Panic1 error or return other 1 value
func Panic1[A any](a A, err error) A {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a
}

// Panic2 error or return other 2 value
func Panic2[A, B any](a A, b B, err error) (A, B) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b
}

// Panic3 error or return other 3 value
func Panic3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c
}

// Panic4 error or return other 4 value
func Panic4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c, d
}

// Panic5 error or return other 5 value
func Panic5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c, d, e
}

// Panic6 error or return other 6 value
func Panic6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F, err error) (A, B, C, D, E, F) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c, d, e, f
}

// Panic7 error or return other 7 value
func Panic7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G, err error) (A, B, C, D, E, F, G) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c, d, e, f, g
}

// Panic8 error or return other 8 value
func Panic8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H, err error) (A, B, C, D, E, F, G, H) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c, d, e, f, g, h
}

// Panic9 error or return other 9 value
func Panic9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, err error) (A, B, C, D, E, F, G, H, I) {
	if err != nil {
		panic(Packer(err, 2))
	}
	return a, b, c, d, e, f, g, h, i
}
