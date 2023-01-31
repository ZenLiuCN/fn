package fn

//generated file should not edit

// PanicsFn01 panic with func 0 in 2 out (last out is an error), returns a Runnable.
func PanicsFn01[A any](fn func() (A, error)) func() {
	return func() {
		var err error
		_, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn02 panic with func 0 in 3 out (last out is an error), returns a Runnable.
func PanicsFn02[A, B any](fn func() (A, B, error)) func() {
	return func() {
		var err error
		_, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn03 panic with func 0 in 4 out (last out is an error), returns a Runnable.
func PanicsFn03[A, B, C any](fn func() (A, B, C, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn04 panic with func 0 in 5 out (last out is an error), returns a Runnable.
func PanicsFn04[A, B, C, D any](fn func() (A, B, C, D, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn05 panic with func 0 in 6 out (last out is an error), returns a Runnable.
func PanicsFn05[A, B, C, D, E any](fn func() (A, B, C, D, E, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn06 panic with func 0 in 7 out (last out is an error), returns a Runnable.
func PanicsFn06[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn07 panic with func 0 in 8 out (last out is an error), returns a Runnable.
func PanicsFn07[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn08 panic with func 0 in 9 out (last out is an error), returns a Runnable.
func PanicsFn08[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn09 panic with func 0 in 10 out (last out is an error), returns a Runnable.
func PanicsFn09[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn10 panic with func 1 in 1 out (last out is an error), returns a Runnable.
func PanicsFn10[A any](a A, fn func(A) error) func() {
	return func() {
		var err error
		err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn11 panic with func 1 in 2 out (last out is an error), returns a Runnable.
func PanicsFn11[A, B any](a A, fn func(A) (B, error)) func() {
	return func() {
		var err error
		_, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn12 panic with func 1 in 3 out (last out is an error), returns a Runnable.
func PanicsFn12[A, B, C any](a A, fn func(A) (B, C, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn13 panic with func 1 in 4 out (last out is an error), returns a Runnable.
func PanicsFn13[A, B, C, D any](a A, fn func(A) (B, C, D, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn14 panic with func 1 in 5 out (last out is an error), returns a Runnable.
func PanicsFn14[A, B, C, D, E any](a A, fn func(A) (B, C, D, E, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn15 panic with func 1 in 6 out (last out is an error), returns a Runnable.
func PanicsFn15[A, B, C, D, E, F any](a A, fn func(A) (B, C, D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn16 panic with func 1 in 7 out (last out is an error), returns a Runnable.
func PanicsFn16[A, B, C, D, E, F, G any](a A, fn func(A) (B, C, D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn17 panic with func 1 in 8 out (last out is an error), returns a Runnable.
func PanicsFn17[A, B, C, D, E, F, G, H any](a A, fn func(A) (B, C, D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn18 panic with func 1 in 9 out (last out is an error), returns a Runnable.
func PanicsFn18[A, B, C, D, E, F, G, H, I any](a A, fn func(A) (B, C, D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn19 panic with func 1 in 10 out (last out is an error), returns a Runnable.
func PanicsFn19[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A) (B, C, D, E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn20 panic with func 2 in 1 out (last out is an error), returns a Runnable.
func PanicsFn20[A, B any](a A, b B, fn func(A, B) error) func() {
	return func() {
		var err error
		err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn21 panic with func 2 in 2 out (last out is an error), returns a Runnable.
func PanicsFn21[A, B, C any](a A, b B, fn func(A, B) (C, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn22 panic with func 2 in 3 out (last out is an error), returns a Runnable.
func PanicsFn22[A, B, C, D any](a A, b B, fn func(A, B) (C, D, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn23 panic with func 2 in 4 out (last out is an error), returns a Runnable.
func PanicsFn23[A, B, C, D, E any](a A, b B, fn func(A, B) (C, D, E, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn24 panic with func 2 in 5 out (last out is an error), returns a Runnable.
func PanicsFn24[A, B, C, D, E, F any](a A, b B, fn func(A, B) (C, D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn25 panic with func 2 in 6 out (last out is an error), returns a Runnable.
func PanicsFn25[A, B, C, D, E, F, G any](a A, b B, fn func(A, B) (C, D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn26 panic with func 2 in 7 out (last out is an error), returns a Runnable.
func PanicsFn26[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B) (C, D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn27 panic with func 2 in 8 out (last out is an error), returns a Runnable.
func PanicsFn27[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn28 panic with func 2 in 9 out (last out is an error), returns a Runnable.
func PanicsFn28[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn29 panic with func 2 in 10 out (last out is an error), returns a Runnable.
func PanicsFn29[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn30 panic with func 3 in 1 out (last out is an error), returns a Runnable.
func PanicsFn30[A, B, C any](a A, b B, c C, fn func(A, B, C) error) func() {
	return func() {
		var err error
		err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn31 panic with func 3 in 2 out (last out is an error), returns a Runnable.
func PanicsFn31[A, B, C, D any](a A, b B, c C, fn func(A, B, C) (D, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn32 panic with func 3 in 3 out (last out is an error), returns a Runnable.
func PanicsFn32[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C) (D, E, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn33 panic with func 3 in 4 out (last out is an error), returns a Runnable.
func PanicsFn33[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C) (D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn34 panic with func 3 in 5 out (last out is an error), returns a Runnable.
func PanicsFn34[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn35 panic with func 3 in 6 out (last out is an error), returns a Runnable.
func PanicsFn35[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn36 panic with func 3 in 7 out (last out is an error), returns a Runnable.
func PanicsFn36[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn37 panic with func 3 in 8 out (last out is an error), returns a Runnable.
func PanicsFn37[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn38 panic with func 3 in 9 out (last out is an error), returns a Runnable.
func PanicsFn38[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn39 panic with func 3 in 10 out (last out is an error), returns a Runnable.
func PanicsFn39[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn40 panic with func 4 in 1 out (last out is an error), returns a Runnable.
func PanicsFn40[A, B, C, D any](a A, b B, c C, d D, fn func(A, B, C, D) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn41 panic with func 4 in 2 out (last out is an error), returns a Runnable.
func PanicsFn41[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D) (E, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn42 panic with func 4 in 3 out (last out is an error), returns a Runnable.
func PanicsFn42[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn43 panic with func 4 in 4 out (last out is an error), returns a Runnable.
func PanicsFn43[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn44 panic with func 4 in 5 out (last out is an error), returns a Runnable.
func PanicsFn44[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn45 panic with func 4 in 6 out (last out is an error), returns a Runnable.
func PanicsFn45[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn46 panic with func 4 in 7 out (last out is an error), returns a Runnable.
func PanicsFn46[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn47 panic with func 4 in 8 out (last out is an error), returns a Runnable.
func PanicsFn47[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn48 panic with func 4 in 9 out (last out is an error), returns a Runnable.
func PanicsFn48[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn49 panic with func 4 in 10 out (last out is an error), returns a Runnable.
func PanicsFn49[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn50 panic with func 5 in 1 out (last out is an error), returns a Runnable.
func PanicsFn50[A, B, C, D, E any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn51 panic with func 5 in 2 out (last out is an error), returns a Runnable.
func PanicsFn51[A, B, C, D, E, F any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn52 panic with func 5 in 3 out (last out is an error), returns a Runnable.
func PanicsFn52[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn53 panic with func 5 in 4 out (last out is an error), returns a Runnable.
func PanicsFn53[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn54 panic with func 5 in 5 out (last out is an error), returns a Runnable.
func PanicsFn54[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn55 panic with func 5 in 6 out (last out is an error), returns a Runnable.
func PanicsFn55[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn56 panic with func 5 in 7 out (last out is an error), returns a Runnable.
func PanicsFn56[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn57 panic with func 5 in 8 out (last out is an error), returns a Runnable.
func PanicsFn57[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn58 panic with func 5 in 9 out (last out is an error), returns a Runnable.
func PanicsFn58[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn59 panic with func 5 in 10 out (last out is an error), returns a Runnable.
func PanicsFn59[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn60 panic with func 6 in 1 out (last out is an error), returns a Runnable.
func PanicsFn60[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn61 panic with func 6 in 2 out (last out is an error), returns a Runnable.
func PanicsFn61[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn62 panic with func 6 in 3 out (last out is an error), returns a Runnable.
func PanicsFn62[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn63 panic with func 6 in 4 out (last out is an error), returns a Runnable.
func PanicsFn63[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn64 panic with func 6 in 5 out (last out is an error), returns a Runnable.
func PanicsFn64[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn65 panic with func 6 in 6 out (last out is an error), returns a Runnable.
func PanicsFn65[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn66 panic with func 6 in 7 out (last out is an error), returns a Runnable.
func PanicsFn66[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn67 panic with func 6 in 8 out (last out is an error), returns a Runnable.
func PanicsFn67[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn68 panic with func 6 in 9 out (last out is an error), returns a Runnable.
func PanicsFn68[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn69 panic with func 6 in 10 out (last out is an error), returns a Runnable.
func PanicsFn69[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn70 panic with func 7 in 1 out (last out is an error), returns a Runnable.
func PanicsFn70[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn71 panic with func 7 in 2 out (last out is an error), returns a Runnable.
func PanicsFn71[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn72 panic with func 7 in 3 out (last out is an error), returns a Runnable.
func PanicsFn72[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn73 panic with func 7 in 4 out (last out is an error), returns a Runnable.
func PanicsFn73[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn74 panic with func 7 in 5 out (last out is an error), returns a Runnable.
func PanicsFn74[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn75 panic with func 7 in 6 out (last out is an error), returns a Runnable.
func PanicsFn75[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn76 panic with func 7 in 7 out (last out is an error), returns a Runnable.
func PanicsFn76[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn77 panic with func 7 in 8 out (last out is an error), returns a Runnable.
func PanicsFn77[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn78 panic with func 7 in 9 out (last out is an error), returns a Runnable.
func PanicsFn78[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn79 panic with func 7 in 10 out (last out is an error), returns a Runnable.
func PanicsFn79[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn80 panic with func 8 in 1 out (last out is an error), returns a Runnable.
func PanicsFn80[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn81 panic with func 8 in 2 out (last out is an error), returns a Runnable.
func PanicsFn81[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn82 panic with func 8 in 3 out (last out is an error), returns a Runnable.
func PanicsFn82[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn83 panic with func 8 in 4 out (last out is an error), returns a Runnable.
func PanicsFn83[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn84 panic with func 8 in 5 out (last out is an error), returns a Runnable.
func PanicsFn84[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn85 panic with func 8 in 6 out (last out is an error), returns a Runnable.
func PanicsFn85[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn86 panic with func 8 in 7 out (last out is an error), returns a Runnable.
func PanicsFn86[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn87 panic with func 8 in 8 out (last out is an error), returns a Runnable.
func PanicsFn87[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn88 panic with func 8 in 9 out (last out is an error), returns a Runnable.
func PanicsFn88[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn89 panic with func 8 in 10 out (last out is an error), returns a Runnable.
func PanicsFn89[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn90 panic with func 9 in 1 out (last out is an error), returns a Runnable.
func PanicsFn90[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn91 panic with func 9 in 2 out (last out is an error), returns a Runnable.
func PanicsFn91[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn92 panic with func 9 in 3 out (last out is an error), returns a Runnable.
func PanicsFn92[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn93 panic with func 9 in 4 out (last out is an error), returns a Runnable.
func PanicsFn93[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn94 panic with func 9 in 5 out (last out is an error), returns a Runnable.
func PanicsFn94[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn95 panic with func 9 in 6 out (last out is an error), returns a Runnable.
func PanicsFn95[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn96 panic with func 9 in 7 out (last out is an error), returns a Runnable.
func PanicsFn96[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn97 panic with func 9 in 8 out (last out is an error), returns a Runnable.
func PanicsFn97[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn98 panic with func 9 in 9 out (last out is an error), returns a Runnable.
func PanicsFn98[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// PanicsFn99 panic with func 9 in 10 out (last out is an error), returns a Runnable.
func PanicsFn99[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}
