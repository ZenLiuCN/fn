package fn

import "fmt"

// Recover a runnable
func Recover(fn func()) (err error) {
	defer func() {
		z := recover()
		switch z.(type) {
		case error:
			err = z.(error)
			return
		case nil:
			return
		default:
			err = fmt.Errorf("%#v", z)
		}
	}()
	fn()
	return
}

// Recover01 panic with func 0 in 1 out
func Recover01[A any](fn func() A) func() (A, error) {
	return func() (a A, err error) {
		defer func() {
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
			}
		}()
		a, b, c, d, e = fn()
		return
	}
}

// Recover10 panic with func 1 in 0 out
func Recover10[A any](fn func(A)) func(a A) error {
	return func(a A) (err error) {
		defer func() {
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
			}
		}()
		b, c, d, e, f = fn(a)
		return
	}
}

// Recover20 panic with func 2 in 0 out
func Recover20[A, B any](fn func(A, B)) func(a A, b B) error {
	return func(a A, b B) (err error) {
		defer func() {
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
			}
		}()
		c, d, e, f, g = fn(a, b)
		return
	}
}

// Recover30 panic with func 3 in 0 out
func Recover30[A, B, C any](fn func(A, B, C)) func(a A, b B, c C) error {
	return func(a A, b B, c C) (err error) {
		defer func() {
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
			}
		}()
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Recover40 panic with func 4 in 0 out
func Recover40[A, B, C, D any](fn func(A, B, C, D)) func(a A, b B, c C, d D) error {
	return func(a A, b B, c C, d D) (err error) {
		defer func() {
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
			}
		}()
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Recover50 panic with func 5 in 0 out
func Recover50[A, B, C, D, E any](fn func(A, B, C, D, E)) func(a A, b B, c C, d D, e E) error {
	return func(a A, b B, c C, d D, e E) (err error) {
		defer func() {
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
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
			z := recover()
			switch z.(type) {
			case error:
				err = z.(error)
				return
			case nil:
				return
			default:
				err = fmt.Errorf("%#v", z)
			}
		}()
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}
