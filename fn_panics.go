//go:build !gen

package fn

//generated file should not edit

// Panics01 panic with func 0 in 2 out (last out is an error), returns a Runnable.
func Panics01[A any](fn func() (A, error)) func() {
	return func() {
		var err error
		_, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics02 panic with func 0 in 3 out (last out is an error), returns a Runnable.
func Panics02[A, B any](fn func() (A, B, error)) func() {
	return func() {
		var err error
		_, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics03 panic with func 0 in 4 out (last out is an error), returns a Runnable.
func Panics03[A, B, C any](fn func() (A, B, C, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics04 panic with func 0 in 5 out (last out is an error), returns a Runnable.
func Panics04[A, B, C, D any](fn func() (A, B, C, D, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics05 panic with func 0 in 6 out (last out is an error), returns a Runnable.
func Panics05[A, B, C, D, E any](fn func() (A, B, C, D, E, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics06 panic with func 0 in 7 out (last out is an error), returns a Runnable.
func Panics06[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics07 panic with func 0 in 8 out (last out is an error), returns a Runnable.
func Panics07[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics08 panic with func 0 in 9 out (last out is an error), returns a Runnable.
func Panics08[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics09 panic with func 0 in 10 out (last out is an error), returns a Runnable.
func Panics09[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn()
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics10 panic with func 1 in 1 out (last out is an error), returns a Runnable.
func Panics10[A any](a A, fn func(A) error) func() {
	return func() {
		var err error
		err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics11 panic with func 1 in 2 out (last out is an error), returns a Runnable.
func Panics11[A, B any](a A, fn func(A) (B, error)) func() {
	return func() {
		var err error
		_, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics12 panic with func 1 in 3 out (last out is an error), returns a Runnable.
func Panics12[A, B, C any](a A, fn func(A) (B, C, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics13 panic with func 1 in 4 out (last out is an error), returns a Runnable.
func Panics13[A, B, C, D any](a A, fn func(A) (B, C, D, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics14 panic with func 1 in 5 out (last out is an error), returns a Runnable.
func Panics14[A, B, C, D, E any](a A, fn func(A) (B, C, D, E, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics15 panic with func 1 in 6 out (last out is an error), returns a Runnable.
func Panics15[A, B, C, D, E, F any](a A, fn func(A) (B, C, D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics16 panic with func 1 in 7 out (last out is an error), returns a Runnable.
func Panics16[A, B, C, D, E, F, G any](a A, fn func(A) (B, C, D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics17 panic with func 1 in 8 out (last out is an error), returns a Runnable.
func Panics17[A, B, C, D, E, F, G, H any](a A, fn func(A) (B, C, D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics18 panic with func 1 in 9 out (last out is an error), returns a Runnable.
func Panics18[A, B, C, D, E, F, G, H, I any](a A, fn func(A) (B, C, D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics19 panic with func 1 in 10 out (last out is an error), returns a Runnable.
func Panics19[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A) (B, C, D, E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics20 panic with func 2 in 1 out (last out is an error), returns a Runnable.
func Panics20[A, B any](a A, b B, fn func(A, B) error) func() {
	return func() {
		var err error
		err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics21 panic with func 2 in 2 out (last out is an error), returns a Runnable.
func Panics21[A, B, C any](a A, b B, fn func(A, B) (C, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics22 panic with func 2 in 3 out (last out is an error), returns a Runnable.
func Panics22[A, B, C, D any](a A, b B, fn func(A, B) (C, D, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics23 panic with func 2 in 4 out (last out is an error), returns a Runnable.
func Panics23[A, B, C, D, E any](a A, b B, fn func(A, B) (C, D, E, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics24 panic with func 2 in 5 out (last out is an error), returns a Runnable.
func Panics24[A, B, C, D, E, F any](a A, b B, fn func(A, B) (C, D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics25 panic with func 2 in 6 out (last out is an error), returns a Runnable.
func Panics25[A, B, C, D, E, F, G any](a A, b B, fn func(A, B) (C, D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics26 panic with func 2 in 7 out (last out is an error), returns a Runnable.
func Panics26[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B) (C, D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics27 panic with func 2 in 8 out (last out is an error), returns a Runnable.
func Panics27[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics28 panic with func 2 in 9 out (last out is an error), returns a Runnable.
func Panics28[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics29 panic with func 2 in 10 out (last out is an error), returns a Runnable.
func Panics29[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics30 panic with func 3 in 1 out (last out is an error), returns a Runnable.
func Panics30[A, B, C any](a A, b B, c C, fn func(A, B, C) error) func() {
	return func() {
		var err error
		err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics31 panic with func 3 in 2 out (last out is an error), returns a Runnable.
func Panics31[A, B, C, D any](a A, b B, c C, fn func(A, B, C) (D, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics32 panic with func 3 in 3 out (last out is an error), returns a Runnable.
func Panics32[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C) (D, E, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics33 panic with func 3 in 4 out (last out is an error), returns a Runnable.
func Panics33[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C) (D, E, F, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics34 panic with func 3 in 5 out (last out is an error), returns a Runnable.
func Panics34[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics35 panic with func 3 in 6 out (last out is an error), returns a Runnable.
func Panics35[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics36 panic with func 3 in 7 out (last out is an error), returns a Runnable.
func Panics36[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics37 panic with func 3 in 8 out (last out is an error), returns a Runnable.
func Panics37[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics38 panic with func 3 in 9 out (last out is an error), returns a Runnable.
func Panics38[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics39 panic with func 3 in 10 out (last out is an error), returns a Runnable.
func Panics39[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics40 panic with func 4 in 1 out (last out is an error), returns a Runnable.
func Panics40[A, B, C, D any](a A, b B, c C, d D, fn func(A, B, C, D) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics41 panic with func 4 in 2 out (last out is an error), returns a Runnable.
func Panics41[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D) (E, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics42 panic with func 4 in 3 out (last out is an error), returns a Runnable.
func Panics42[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics43 panic with func 4 in 4 out (last out is an error), returns a Runnable.
func Panics43[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics44 panic with func 4 in 5 out (last out is an error), returns a Runnable.
func Panics44[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics45 panic with func 4 in 6 out (last out is an error), returns a Runnable.
func Panics45[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics46 panic with func 4 in 7 out (last out is an error), returns a Runnable.
func Panics46[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics47 panic with func 4 in 8 out (last out is an error), returns a Runnable.
func Panics47[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics48 panic with func 4 in 9 out (last out is an error), returns a Runnable.
func Panics48[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics49 panic with func 4 in 10 out (last out is an error), returns a Runnable.
func Panics49[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics50 panic with func 5 in 1 out (last out is an error), returns a Runnable.
func Panics50[A, B, C, D, E any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics51 panic with func 5 in 2 out (last out is an error), returns a Runnable.
func Panics51[A, B, C, D, E, F any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics52 panic with func 5 in 3 out (last out is an error), returns a Runnable.
func Panics52[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics53 panic with func 5 in 4 out (last out is an error), returns a Runnable.
func Panics53[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics54 panic with func 5 in 5 out (last out is an error), returns a Runnable.
func Panics54[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics55 panic with func 5 in 6 out (last out is an error), returns a Runnable.
func Panics55[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics56 panic with func 5 in 7 out (last out is an error), returns a Runnable.
func Panics56[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics57 panic with func 5 in 8 out (last out is an error), returns a Runnable.
func Panics57[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics58 panic with func 5 in 9 out (last out is an error), returns a Runnable.
func Panics58[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics59 panic with func 5 in 10 out (last out is an error), returns a Runnable.
func Panics59[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics60 panic with func 6 in 1 out (last out is an error), returns a Runnable.
func Panics60[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics61 panic with func 6 in 2 out (last out is an error), returns a Runnable.
func Panics61[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics62 panic with func 6 in 3 out (last out is an error), returns a Runnable.
func Panics62[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics63 panic with func 6 in 4 out (last out is an error), returns a Runnable.
func Panics63[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics64 panic with func 6 in 5 out (last out is an error), returns a Runnable.
func Panics64[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics65 panic with func 6 in 6 out (last out is an error), returns a Runnable.
func Panics65[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics66 panic with func 6 in 7 out (last out is an error), returns a Runnable.
func Panics66[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics67 panic with func 6 in 8 out (last out is an error), returns a Runnable.
func Panics67[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics68 panic with func 6 in 9 out (last out is an error), returns a Runnable.
func Panics68[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics69 panic with func 6 in 10 out (last out is an error), returns a Runnable.
func Panics69[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics70 panic with func 7 in 1 out (last out is an error), returns a Runnable.
func Panics70[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics71 panic with func 7 in 2 out (last out is an error), returns a Runnable.
func Panics71[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics72 panic with func 7 in 3 out (last out is an error), returns a Runnable.
func Panics72[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics73 panic with func 7 in 4 out (last out is an error), returns a Runnable.
func Panics73[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics74 panic with func 7 in 5 out (last out is an error), returns a Runnable.
func Panics74[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics75 panic with func 7 in 6 out (last out is an error), returns a Runnable.
func Panics75[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics76 panic with func 7 in 7 out (last out is an error), returns a Runnable.
func Panics76[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics77 panic with func 7 in 8 out (last out is an error), returns a Runnable.
func Panics77[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics78 panic with func 7 in 9 out (last out is an error), returns a Runnable.
func Panics78[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics79 panic with func 7 in 10 out (last out is an error), returns a Runnable.
func Panics79[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics80 panic with func 8 in 1 out (last out is an error), returns a Runnable.
func Panics80[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics81 panic with func 8 in 2 out (last out is an error), returns a Runnable.
func Panics81[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics82 panic with func 8 in 3 out (last out is an error), returns a Runnable.
func Panics82[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics83 panic with func 8 in 4 out (last out is an error), returns a Runnable.
func Panics83[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics84 panic with func 8 in 5 out (last out is an error), returns a Runnable.
func Panics84[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics85 panic with func 8 in 6 out (last out is an error), returns a Runnable.
func Panics85[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics86 panic with func 8 in 7 out (last out is an error), returns a Runnable.
func Panics86[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics87 panic with func 8 in 8 out (last out is an error), returns a Runnable.
func Panics87[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics88 panic with func 8 in 9 out (last out is an error), returns a Runnable.
func Panics88[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics89 panic with func 8 in 10 out (last out is an error), returns a Runnable.
func Panics89[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics90 panic with func 9 in 1 out (last out is an error), returns a Runnable.
func Panics90[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) error) func() {
	return func() {
		var err error
		err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics91 panic with func 9 in 2 out (last out is an error), returns a Runnable.
func Panics91[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, error)) func() {
	return func() {
		var err error
		_, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics92 panic with func 9 in 3 out (last out is an error), returns a Runnable.
func Panics92[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, error)) func() {
	return func() {
		var err error
		_, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics93 panic with func 9 in 4 out (last out is an error), returns a Runnable.
func Panics93[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, error)) func() {
	return func() {
		var err error
		_, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics94 panic with func 9 in 5 out (last out is an error), returns a Runnable.
func Panics94[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, error)) func() {
	return func() {
		var err error
		_, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics95 panic with func 9 in 6 out (last out is an error), returns a Runnable.
func Panics95[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics96 panic with func 9 in 7 out (last out is an error), returns a Runnable.
func Panics96[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics97 panic with func 9 in 8 out (last out is an error), returns a Runnable.
func Panics97[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics98 panic with func 9 in 9 out (last out is an error), returns a Runnable.
func Panics98[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}

// Panics99 panic with func 9 in 10 out (last out is an error), returns a Runnable.
func Panics99[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R, error)) func() {
	return func() {
		var err error
		_, _, _, _, _, _, _, _, _, err = fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			panic(err)
		}
		return
	}
}
