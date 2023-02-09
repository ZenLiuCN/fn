//go:build !gen

package fn

//generated file should not edit

// Recover01 panic with func 0 in 1 out
func Recover01[A any](fn func() A) func() (A, error) {
	return func() (a A, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a = fn()
		return
	}
}

// Recover02 panic with func 0 in 2 out
func Recover02[A, B any](fn func() (A, B)) func() (A, B, error) {
	return func() (a A, b B, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b = fn()
		return
	}
}

// Recover03 panic with func 0 in 3 out
func Recover03[A, B, C any](fn func() (A, B, C)) func() (A, B, C, error) {
	return func() (a A, b B, c C, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c = fn()
		return
	}
}

// Recover04 panic with func 0 in 4 out
func Recover04[A, B, C, D any](fn func() (A, B, C, D)) func() (A, B, C, D, error) {
	return func() (a A, b B, c C, d D, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c, d = fn()
		return
	}
}

// Recover05 panic with func 0 in 5 out
func Recover05[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (A, B, C, D, E, error) {
	return func() (a A, b B, c C, d D, e E, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c, d, e = fn()
		return
	}
}

// Recover06 panic with func 0 in 6 out
func Recover06[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (A, B, C, D, E, F, error) {
	return func() (a A, b B, c C, d D, e E, f F, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c, d, e, f = fn()
		return
	}
}

// Recover07 panic with func 0 in 7 out
func Recover07[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (A, B, C, D, E, F, G, error) {
	return func() (a A, b B, c C, d D, e E, f F, g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c, d, e, f, g = fn()
		return
	}
}

// Recover08 panic with func 0 in 8 out
func Recover08[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B, C, D, E, F, G, H, error) {
	return func() (a A, b B, c C, d D, e E, f F, g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c, d, e, f, g, h = fn()
		return
	}
}

// Recover09 panic with func 0 in 9 out
func Recover09[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C, D, E, F, G, H, I, error) {
	return func() (a A, b B, c C, d D, e E, f F, g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		a, b, c, d, e, f, g, h, i = fn()
		return
	}
}

// Recover10 panic with func 1 in 0 out
func Recover10[A any](fn func(A)) func(a A) error {
	return func(a A) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a)
		return
	}
}

// Recover11 panic with func 1 in 1 out
func Recover11[A, B any](fn func(A) B) func(a A) (B, error) {
	return func(a A) (b B, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b = fn(a)
		return
	}
}

// Recover12 panic with func 1 in 2 out
func Recover12[A, B, C any](fn func(A) (B, C)) func(a A) (B, C, error) {
	return func(a A) (b B, c C, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c = fn(a)
		return
	}
}

// Recover13 panic with func 1 in 3 out
func Recover13[A, B, C, D any](fn func(A) (B, C, D)) func(a A) (B, C, D, error) {
	return func(a A) (b B, c C, d D, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d = fn(a)
		return
	}
}

// Recover14 panic with func 1 in 4 out
func Recover14[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) (B, C, D, E, error) {
	return func(a A) (b B, c C, d D, e E, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d, e = fn(a)
		return
	}
}

// Recover15 panic with func 1 in 5 out
func Recover15[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (B, C, D, E, F, error) {
	return func(a A) (b B, c C, d D, e E, f F, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d, e, f = fn(a)
		return
	}
}

// Recover16 panic with func 1 in 6 out
func Recover16[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (B, C, D, E, F, G, error) {
	return func(a A) (b B, c C, d D, e E, f F, g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d, e, f, g = fn(a)
		return
	}
}

// Recover17 panic with func 1 in 7 out
func Recover17[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (B, C, D, E, F, G, H, error) {
	return func(a A) (b B, c C, d D, e E, f F, g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d, e, f, g, h = fn(a)
		return
	}
}

// Recover18 panic with func 1 in 8 out
func Recover18[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C, D, E, F, G, H, I, error) {
	return func(a A) (b B, c C, d D, e E, f F, g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d, e, f, g, h, i = fn(a)
		return
	}
}

// Recover19 panic with func 1 in 9 out
func Recover19[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D, E, F, G, H, I, J, error) {
	return func(a A) (b B, c C, d D, e E, f F, g G, h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		b, c, d, e, f, g, h, i, j = fn(a)
		return
	}
}

// Recover20 panic with func 2 in 0 out
func Recover20[A, B any](fn func(A, B)) func(a A, b B) error {
	return func(a A, b B) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b)
		return
	}
}

// Recover21 panic with func 2 in 1 out
func Recover21[A, B, C any](fn func(A, B) C) func(a A, b B) (C, error) {
	return func(a A, b B) (c C, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c = fn(a, b)
		return
	}
}

// Recover22 panic with func 2 in 2 out
func Recover22[A, B, C, D any](fn func(A, B) (C, D)) func(a A, b B) (C, D, error) {
	return func(a A, b B) (c C, d D, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d = fn(a, b)
		return
	}
}

// Recover23 panic with func 2 in 3 out
func Recover23[A, B, C, D, E any](fn func(A, B) (C, D, E)) func(a A, b B) (C, D, E, error) {
	return func(a A, b B) (c C, d D, e E, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e = fn(a, b)
		return
	}
}

// Recover24 panic with func 2 in 4 out
func Recover24[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) (C, D, E, F, error) {
	return func(a A, b B) (c C, d D, e E, f F, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e, f = fn(a, b)
		return
	}
}

// Recover25 panic with func 2 in 5 out
func Recover25[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (C, D, E, F, G, error) {
	return func(a A, b B) (c C, d D, e E, f F, g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e, f, g = fn(a, b)
		return
	}
}

// Recover26 panic with func 2 in 6 out
func Recover26[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (C, D, E, F, G, H, error) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// Recover27 panic with func 2 in 7 out
func Recover27[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (C, D, E, F, G, H, I, error) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// Recover28 panic with func 2 in 8 out
func Recover28[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D, E, F, G, H, I, J, error) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// Recover29 panic with func 2 in 9 out
func Recover29[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E, F, G, H, I, J, K, error) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// Recover30 panic with func 3 in 0 out
func Recover30[A, B, C any](fn func(A, B, C)) func(a A, b B, c C) error {
	return func(a A, b B, c C) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c)
		return
	}
}

// Recover31 panic with func 3 in 1 out
func Recover31[A, B, C, D any](fn func(A, B, C) D) func(a A, b B, c C) (D, error) {
	return func(a A, b B, c C) (d D, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d = fn(a, b, c)
		return
	}
}

// Recover32 panic with func 3 in 2 out
func Recover32[A, B, C, D, E any](fn func(A, B, C) (D, E)) func(a A, b B, c C) (D, E, error) {
	return func(a A, b B, c C) (d D, e E, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e = fn(a, b, c)
		return
	}
}

// Recover33 panic with func 3 in 3 out
func Recover33[A, B, C, D, E, F any](fn func(A, B, C) (D, E, F)) func(a A, b B, c C) (D, E, F, error) {
	return func(a A, b B, c C) (d D, e E, f F, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f = fn(a, b, c)
		return
	}
}

// Recover34 panic with func 3 in 4 out
func Recover34[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) (D, E, F, G, error) {
	return func(a A, b B, c C) (d D, e E, f F, g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Recover35 panic with func 3 in 5 out
func Recover35[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (D, E, F, G, H, error) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Recover36 panic with func 3 in 6 out
func Recover36[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (D, E, F, G, H, I, error) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Recover37 panic with func 3 in 7 out
func Recover37[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (D, E, F, G, H, I, J, error) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Recover38 panic with func 3 in 8 out
func Recover38[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E, F, G, H, I, J, K, error) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Recover39 panic with func 3 in 9 out
func Recover39[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F, G, H, I, J, K, L, error) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Recover40 panic with func 4 in 0 out
func Recover40[A, B, C, D any](fn func(A, B, C, D)) func(a A, b B, c C, d D) error {
	return func(a A, b B, c C, d D) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c, d)
		return
	}
}

// Recover41 panic with func 4 in 1 out
func Recover41[A, B, C, D, E any](fn func(A, B, C, D) E) func(a A, b B, c C, d D) (E, error) {
	return func(a A, b B, c C, d D) (e E, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e = fn(a, b, c, d)
		return
	}
}

// Recover42 panic with func 4 in 2 out
func Recover42[A, B, C, D, E, F any](fn func(A, B, C, D) (E, F)) func(a A, b B, c C, d D) (E, F, error) {
	return func(a A, b B, c C, d D) (e E, f F, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f = fn(a, b, c, d)
		return
	}
}

// Recover43 panic with func 4 in 3 out
func Recover43[A, B, C, D, E, F, G any](fn func(A, B, C, D) (E, F, G)) func(a A, b B, c C, d D) (E, F, G, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Recover44 panic with func 4 in 4 out
func Recover44[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) (E, F, G, H, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Recover45 panic with func 4 in 5 out
func Recover45[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (E, F, G, H, I, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Recover46 panic with func 4 in 6 out
func Recover46[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (E, F, G, H, I, J, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Recover47 panic with func 4 in 7 out
func Recover47[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (E, F, G, H, I, J, K, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Recover48 panic with func 4 in 8 out
func Recover48[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F, G, H, I, J, K, L, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Recover49 panic with func 4 in 9 out
func Recover49[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G, H, I, J, K, L, M, error) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Recover50 panic with func 5 in 0 out
func Recover50[A, B, C, D, E any](fn func(A, B, C, D, E)) func(a A, b B, c C, d D, e E) error {
	return func(a A, b B, c C, d D, e E) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c, d, e)
		return
	}
}

// Recover51 panic with func 5 in 1 out
func Recover51[A, B, C, D, E, F any](fn func(A, B, C, D, E) F) func(a A, b B, c C, d D, e E) (F, error) {
	return func(a A, b B, c C, d D, e E) (f F, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f = fn(a, b, c, d, e)
		return
	}
}

// Recover52 panic with func 5 in 2 out
func Recover52[A, B, C, D, E, F, G any](fn func(A, B, C, D, E) (F, G)) func(a A, b B, c C, d D, e E) (F, G, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Recover53 panic with func 5 in 3 out
func Recover53[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E) (F, G, H)) func(a A, b B, c C, d D, e E) (F, G, H, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Recover54 panic with func 5 in 4 out
func Recover54[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) (F, G, H, I, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Recover55 panic with func 5 in 5 out
func Recover55[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Recover56 panic with func 5 in 6 out
func Recover56[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Recover57 panic with func 5 in 7 out
func Recover57[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, L, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Recover58 panic with func 5 in 8 out
func Recover58[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, L, M, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Recover59 panic with func 5 in 9 out
func Recover59[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, L, M, N, error) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Recover60 panic with func 6 in 0 out
func Recover60[A, B, C, D, E, F any](fn func(A, B, C, D, E, F)) func(a A, b B, c C, d D, e E, f F) error {
	return func(a A, b B, c C, d D, e E, f F) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c, d, e, f)
		return
	}
}

// Recover61 panic with func 6 in 1 out
func Recover61[A, B, C, D, E, F, G any](fn func(A, B, C, D, E, F) G) func(a A, b B, c C, d D, e E, f F) (G, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Recover62 panic with func 6 in 2 out
func Recover62[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F) (G, H)) func(a A, b B, c C, d D, e E, f F) (G, H, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Recover63 panic with func 6 in 3 out
func Recover63[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F) (G, H, I)) func(a A, b B, c C, d D, e E, f F) (G, H, I, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Recover64 panic with func 6 in 4 out
func Recover64[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Recover65 panic with func 6 in 5 out
func Recover65[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Recover66 panic with func 6 in 6 out
func Recover66[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Recover67 panic with func 6 in 7 out
func Recover67[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, M, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Recover68 panic with func 6 in 8 out
func Recover68[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, M, N, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Recover69 panic with func 6 in 9 out
func Recover69[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, M, N, O, error) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Recover70 panic with func 7 in 0 out
func Recover70[A, B, C, D, E, F, G any](fn func(A, B, C, D, E, F, G)) func(a A, b B, c C, d D, e E, f F, g G) error {
	return func(a A, b B, c C, d D, e E, f F, g G) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover71 panic with func 7 in 1 out
func Recover71[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F, G) H) func(a A, b B, c C, d D, e E, f F, g G) (H, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover72 panic with func 7 in 2 out
func Recover72[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G) (H, I)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover73 panic with func 7 in 3 out
func Recover73[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G) (H, I, J)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover74 panic with func 7 in 4 out
func Recover74[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover75 panic with func 7 in 5 out
func Recover75[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover76 panic with func 7 in 6 out
func Recover76[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover77 panic with func 7 in 7 out
func Recover77[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, N, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover78 panic with func 7 in 8 out
func Recover78[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, N, O, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover79 panic with func 7 in 9 out
func Recover79[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, N, O, P, error) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Recover80 panic with func 8 in 0 out
func Recover80[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F, G, H)) func(a A, b B, c C, d D, e E, f F, g G, h H) error {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover81 panic with func 8 in 1 out
func Recover81[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G, H) I) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover82 panic with func 8 in 2 out
func Recover82[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H) (I, J)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover83 panic with func 8 in 3 out
func Recover83[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover84 panic with func 8 in 4 out
func Recover84[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover85 panic with func 8 in 5 out
func Recover85[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover86 panic with func 8 in 6 out
func Recover86[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover87 panic with func 8 in 7 out
func Recover87[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, O, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover88 panic with func 8 in 8 out
func Recover88[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, O, P, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover89 panic with func 8 in 9 out
func Recover89[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, O, P, Q, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Recover90 panic with func 9 in 0 out
func Recover90[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G, H, I)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) error {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover91 panic with func 9 in 1 out
func Recover91[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H, I) J) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover92 panic with func 9 in 2 out
func Recover92[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover93 panic with func 9 in 3 out
func Recover93[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover94 panic with func 9 in 4 out
func Recover94[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover95 panic with func 9 in 5 out
func Recover95[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover96 panic with func 9 in 6 out
func Recover96[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover97 panic with func 9 in 7 out
func Recover97[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, P, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover98 panic with func 9 in 8 out
func Recover98[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, P, Q, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Recover99 panic with func 9 in 9 out
func Recover99[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, P, Q, R, error) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R, err error) {
		defer func() {
			if z := recover(); z != nil {
				err = Packer(z, 3)
			}
		}()
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}
