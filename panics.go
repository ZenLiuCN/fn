package fn

//Panic error
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// Panic1 error or return other 1 value
func Panic1[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

// Panic2 error or return other 2 value
func Panic2[A, B any](a A, b B, err error) (A, B) {
	if err != nil {
		panic(err)
	}
	return a, b
}

// Panic3 error or return other 3 value
func Panic3[A, B, C any](a A, b B, c C, err error) (A, B, C) {
	if err != nil {
		panic(err)
	}
	return a, b, c
}

// Panic4 error or return other 4 value
func Panic4[A, B, C, D any](a A, b B, c C, d D, err error) (A, B, C, D) {
	if err != nil {
		panic(err)
	}
	return a, b, c, d
}

// Panic5 error or return other 5 value
func Panic5[A, B, C, D, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E) {
	if err != nil {
		panic(err)
	}
	return a, b, c, d, e
}
