package fn

//region Drop

// Drop021 with func( 0 in)(2 out) drop first 1 result
func Drop021[A, B any](fn func() (A, B)) func() B {
	return func() (b B) {
		_, b = fn()
		return
	}
}

// DropLast021 with func( 0 in)(2 out) drop last 1 result
func DropLast021[A, B any](fn func() (A, B)) func() A {
	return func() (a A) {
		a, _ = fn()
		return
	}
}

// Drop031 with func( 0 in)(3 out) drop first 1 result
func Drop031[A, B, C any](fn func() (A, B, C)) func() (B, C) {
	return func() (b B, c C) {
		_, b, c = fn()
		return
	}
}

// DropLast031 with func( 0 in)(3 out) drop last 1 result
func DropLast031[A, B, C any](fn func() (A, B, C)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _ = fn()
		return
	}
}

// Drop032 with func( 0 in)(3 out) drop first 2 result
func Drop032[A, B, C any](fn func() (A, B, C)) func() C {
	return func() (c C) {
		_, _, c = fn()
		return
	}
}

// DropLast032 with func( 0 in)(3 out) drop last 2 result
func DropLast032[A, B, C any](fn func() (A, B, C)) func() A {
	return func() (a A) {
		a, _, _ = fn()
		return
	}
}

// Drop041 with func( 0 in)(4 out) drop first 1 result
func Drop041[A, B, C, D any](fn func() (A, B, C, D)) func() (B, C, D) {
	return func() (b B, c C, d D) {
		_, b, c, d = fn()
		return
	}
}

// DropLast041 with func( 0 in)(4 out) drop last 1 result
func DropLast041[A, B, C, D any](fn func() (A, B, C, D)) func() (A, B, C) {
	return func() (a A, b B, c C) {
		a, b, c, _ = fn()
		return
	}
}

// Drop042 with func( 0 in)(4 out) drop first 2 result
func Drop042[A, B, C, D any](fn func() (A, B, C, D)) func() (C, D) {
	return func() (c C, d D) {
		_, _, c, d = fn()
		return
	}
}

// DropLast042 with func( 0 in)(4 out) drop last 2 result
func DropLast042[A, B, C, D any](fn func() (A, B, C, D)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _, _ = fn()
		return
	}
}

// Drop043 with func( 0 in)(4 out) drop first 3 result
func Drop043[A, B, C, D any](fn func() (A, B, C, D)) func() D {
	return func() (d D) {
		_, _, _, d = fn()
		return
	}
}

// DropLast043 with func( 0 in)(4 out) drop last 3 result
func DropLast043[A, B, C, D any](fn func() (A, B, C, D)) func() A {
	return func() (a A) {
		a, _, _, _ = fn()
		return
	}
}

// Drop051 with func( 0 in)(5 out) drop first 1 result
func Drop051[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (B, C, D, E) {
	return func() (b B, c C, d D, e E) {
		_, b, c, d, e = fn()
		return
	}
}

// DropLast051 with func( 0 in)(5 out) drop last 1 result
func DropLast051[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (A, B, C, D) {
	return func() (a A, b B, c C, d D) {
		a, b, c, d, _ = fn()
		return
	}
}

// Drop052 with func( 0 in)(5 out) drop first 2 result
func Drop052[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (C, D, E) {
	return func() (c C, d D, e E) {
		_, _, c, d, e = fn()
		return
	}
}

// DropLast052 with func( 0 in)(5 out) drop last 2 result
func DropLast052[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (A, B, C) {
	return func() (a A, b B, c C) {
		a, b, c, _, _ = fn()
		return
	}
}

// Drop053 with func( 0 in)(5 out) drop first 3 result
func Drop053[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (D, E) {
	return func() (d D, e E) {
		_, _, _, d, e = fn()
		return
	}
}

// DropLast053 with func( 0 in)(5 out) drop last 3 result
func DropLast053[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _, _, _ = fn()
		return
	}
}

// Drop054 with func( 0 in)(5 out) drop first 4 result
func Drop054[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() E {
	return func() (e E) {
		_, _, _, _, e = fn()
		return
	}
}

// DropLast054 with func( 0 in)(5 out) drop last 4 result
func DropLast054[A, B, C, D, E any](fn func() (A, B, C, D, E)) func() A {
	return func() (a A) {
		a, _, _, _, _ = fn()
		return
	}
}

// Drop121 with func( 1 in)(2 out) drop first 1 result
func Drop121[A, B, C any](fn func(A) (B, C)) func(a A) C {
	return func(a A) (c C) {
		_, c = fn(a)
		return
	}
}

// DropLast121 with func( 1 in)(2 out) drop last 1 result
func DropLast121[A, B, C any](fn func(A) (B, C)) func(a A) B {
	return func(a A) (b B) {
		b, _ = fn(a)
		return
	}
}

// Drop131 with func( 1 in)(3 out) drop first 1 result
func Drop131[A, B, C, D any](fn func(A) (B, C, D)) func(a A) (C, D) {
	return func(a A) (c C, d D) {
		_, c, d = fn(a)
		return
	}
}

// DropLast131 with func( 1 in)(3 out) drop last 1 result
func DropLast131[A, B, C, D any](fn func(A) (B, C, D)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _ = fn(a)
		return
	}
}

// Drop132 with func( 1 in)(3 out) drop first 2 result
func Drop132[A, B, C, D any](fn func(A) (B, C, D)) func(a A) D {
	return func(a A) (d D) {
		_, _, d = fn(a)
		return
	}
}

// DropLast132 with func( 1 in)(3 out) drop last 2 result
func DropLast132[A, B, C, D any](fn func(A) (B, C, D)) func(a A) B {
	return func(a A) (b B) {
		b, _, _ = fn(a)
		return
	}
}

// Drop141 with func( 1 in)(4 out) drop first 1 result
func Drop141[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) (C, D, E) {
	return func(a A) (c C, d D, e E) {
		_, c, d, e = fn(a)
		return
	}
}

// DropLast141 with func( 1 in)(4 out) drop last 1 result
func DropLast141[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) (B, C, D) {
	return func(a A) (b B, c C, d D) {
		b, c, d, _ = fn(a)
		return
	}
}

// Drop142 with func( 1 in)(4 out) drop first 2 result
func Drop142[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) (D, E) {
	return func(a A) (d D, e E) {
		_, _, d, e = fn(a)
		return
	}
}

// DropLast142 with func( 1 in)(4 out) drop last 2 result
func DropLast142[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _, _ = fn(a)
		return
	}
}

// Drop143 with func( 1 in)(4 out) drop first 3 result
func Drop143[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) E {
	return func(a A) (e E) {
		_, _, _, e = fn(a)
		return
	}
}

// DropLast143 with func( 1 in)(4 out) drop last 3 result
func DropLast143[A, B, C, D, E any](fn func(A) (B, C, D, E)) func(a A) B {
	return func(a A) (b B) {
		b, _, _, _ = fn(a)
		return
	}
}

// Drop151 with func( 1 in)(5 out) drop first 1 result
func Drop151[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (C, D, E, F) {
	return func(a A) (c C, d D, e E, f F) {
		_, c, d, e, f = fn(a)
		return
	}
}

// DropLast151 with func( 1 in)(5 out) drop last 1 result
func DropLast151[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (B, C, D, E) {
	return func(a A) (b B, c C, d D, e E) {
		b, c, d, e, _ = fn(a)
		return
	}
}

// Drop152 with func( 1 in)(5 out) drop first 2 result
func Drop152[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (D, E, F) {
	return func(a A) (d D, e E, f F) {
		_, _, d, e, f = fn(a)
		return
	}
}

// DropLast152 with func( 1 in)(5 out) drop last 2 result
func DropLast152[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (B, C, D) {
	return func(a A) (b B, c C, d D) {
		b, c, d, _, _ = fn(a)
		return
	}
}

// Drop153 with func( 1 in)(5 out) drop first 3 result
func Drop153[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (E, F) {
	return func(a A) (e E, f F) {
		_, _, _, e, f = fn(a)
		return
	}
}

// DropLast153 with func( 1 in)(5 out) drop last 3 result
func DropLast153[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _, _, _ = fn(a)
		return
	}
}

// Drop154 with func( 1 in)(5 out) drop first 4 result
func Drop154[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) F {
	return func(a A) (f F) {
		_, _, _, _, f = fn(a)
		return
	}
}

// DropLast154 with func( 1 in)(5 out) drop last 4 result
func DropLast154[A, B, C, D, E, F any](fn func(A) (B, C, D, E, F)) func(a A) B {
	return func(a A) (b B) {
		b, _, _, _, _ = fn(a)
		return
	}
}

// Drop221 with func( 2 in)(2 out) drop first 1 result
func Drop221[A, B, C, D any](fn func(A, B) (C, D)) func(a A, b B) D {
	return func(a A, b B) (d D) {
		_, d = fn(a, b)
		return
	}
}

// DropLast221 with func( 2 in)(2 out) drop last 1 result
func DropLast221[A, B, C, D any](fn func(A, B) (C, D)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _ = fn(a, b)
		return
	}
}

// Drop231 with func( 2 in)(3 out) drop first 1 result
func Drop231[A, B, C, D, E any](fn func(A, B) (C, D, E)) func(a A, b B) (D, E) {
	return func(a A, b B) (d D, e E) {
		_, d, e = fn(a, b)
		return
	}
}

// DropLast231 with func( 2 in)(3 out) drop last 1 result
func DropLast231[A, B, C, D, E any](fn func(A, B) (C, D, E)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _ = fn(a, b)
		return
	}
}

// Drop232 with func( 2 in)(3 out) drop first 2 result
func Drop232[A, B, C, D, E any](fn func(A, B) (C, D, E)) func(a A, b B) E {

	return func(a A, b B) (e E) {
		_, _, e = fn(a, b)
		return
	}
}

// DropLast232 with func( 2 in)(3 out) drop last 2 result
func DropLast232[A, B, C, D, E any](fn func(A, B) (C, D, E)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _ = fn(a, b)
		return
	}
}

// Drop241 with func( 2 in)(4 out) drop first 1 result
func Drop241[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) (D, E, F) {
	return func(a A, b B) (d D, e E, f F) {
		_, d, e, f = fn(a, b)
		return
	}
}

// DropLast241 with func( 2 in)(4 out) drop last 1 result
func DropLast241[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) (C, D, E) {
	return func(a A, b B) (c C, d D, e E) {
		c, d, e, _ = fn(a, b)
		return
	}
}

// Drop242 with func( 2 in)(4 out) drop first 2 result
func Drop242[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) (E, F) {
	return func(a A, b B) (e E, f F) {
		_, _, e, f = fn(a, b)
		return
	}
}

// DropLast242 with func( 2 in)(4 out) drop last 2 result
func DropLast242[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _, _ = fn(a, b)
		return
	}
}

// Drop243 with func( 2 in)(4 out) drop first 3 result
func Drop243[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) F {
	return func(a A, b B) (f F) {
		_, _, _, f = fn(a, b)
		return
	}
}

// DropLast243 with func( 2 in)(4 out) drop last 3 result
func DropLast243[A, B, C, D, E, F any](fn func(A, B) (C, D, E, F)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _, _ = fn(a, b)
		return
	}
}

// Drop251 with func( 2 in)(5 out) drop first 1 result
func Drop251[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (D, E, F, G) {
	return func(a A, b B) (d D, e E, f F, g G) {
		_, d, e, f, g = fn(a, b)
		return
	}
}

// DropLast251 with func( 2 in)(5 out) drop last 1 result
func DropLast251[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (C, D, E, F) {
	return func(a A, b B) (c C, d D, e E, f F) {
		c, d, e, f, _ = fn(a, b)
		return
	}
}

// Drop252 with func( 2 in)(5 out) drop first 2 result
func Drop252[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (E, F, G) {
	return func(a A, b B) (e E, f F, g G) {
		_, _, e, f, g = fn(a, b)
		return
	}
}

// DropLast252 with func( 2 in)(5 out) drop last 2 result
func DropLast252[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (C, D, E) {
	return func(a A, b B) (c C, d D, e E) {
		c, d, e, _, _ = fn(a, b)
		return
	}
}

// Drop253 with func( 2 in)(5 out) drop first 3 result
func Drop253[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (F, G) {
	return func(a A, b B) (f F, g G) {
		_, _, _, f, g = fn(a, b)
		return
	}
}

// DropLast253 with func( 2 in)(5 out) drop last 3 result
func DropLast253[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _, _, _ = fn(a, b)
		return
	}
}

// Drop254 with func( 2 in)(5 out) drop first 4 result
func Drop254[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) G {
	return func(a A, b B) (g G) {
		_, _, _, _, g = fn(a, b)
		return
	}
}

// DropLast254 with func( 2 in)(5 out) drop last 4 result
func DropLast254[A, B, C, D, E, F, G any](fn func(A, B) (C, D, E, F, G)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop321 with func( 3 in)(2 out) drop first 1 result
func Drop321[A, B, C, D, E any](fn func(A, B, C) (D, E)) func(a A, b B, c C) E {
	return func(a A, b B, c C) (e E) {
		_, e = fn(a, b, c)
		return
	}
}

// DropLast321 with func( 3 in)(2 out) drop last 1 result
func DropLast321[A, B, C, D, E any](fn func(A, B, C) (D, E)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _ = fn(a, b, c)
		return
	}
}

// Drop331 with func( 3 in)(3 out) drop first 1 result
func Drop331[A, B, C, D, E, F any](fn func(A, B, C) (D, E, F)) func(a A, b B, c C) (E, F) {
	return func(a A, b B, c C) (e E, f F) {
		_, e, f = fn(a, b, c)
		return
	}
}

// DropLast331 with func( 3 in)(3 out) drop last 1 result
func DropLast331[A, B, C, D, E, F any](fn func(A, B, C) (D, E, F)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _ = fn(a, b, c)
		return
	}
}

// Drop332 with func( 3 in)(3 out) drop first 2 result
func Drop332[A, B, C, D, E, F any](fn func(A, B, C) (D, E, F)) func(a A, b B, c C) F {
	return func(a A, b B, c C) (f F) {
		_, _, f = fn(a, b, c)
		return
	}
}

// DropLast332 with func( 3 in)(3 out) drop last 2 result
func DropLast332[A, B, C, D, E, F any](fn func(A, B, C) (D, E, F)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _ = fn(a, b, c)
		return
	}
}

// Drop341 with func( 3 in)(4 out) drop first 1 result
func Drop341[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) (E, F, G) {
	return func(a A, b B, c C) (e E, f F, g G) {
		_, e, f, g = fn(a, b, c)
		return
	}
}

// DropLast341 with func( 3 in)(4 out) drop last 1 result
func DropLast341[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) (D, E, F) {
	return func(a A, b B, c C) (d D, e E, f F) {
		d, e, f, _ = fn(a, b, c)
		return
	}
}

// Drop342 with func( 3 in)(4 out) drop first 2 result
func Drop342[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) (F, G) {
	return func(a A, b B, c C) (f F, g G) {
		_, _, f, g = fn(a, b, c)
		return
	}
}

// DropLast342 with func( 3 in)(4 out) drop last 2 result
func DropLast342[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _, _ = fn(a, b, c)
		return
	}
}

// Drop343 with func( 3 in)(4 out) drop first 3 result
func Drop343[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) G {
	return func(a A, b B, c C) (g G) {
		_, _, _, g = fn(a, b, c)
		return
	}
}

// DropLast343 with func( 3 in)(4 out) drop last 3 result
func DropLast343[A, B, C, D, E, F, G any](fn func(A, B, C) (D, E, F, G)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop351 with func( 3 in)(5 out) drop first 1 result
func Drop351[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (E, F, G, H) {
	return func(a A, b B, c C) (e E, f F, g G, h H) {
		_, e, f, g, h = fn(a, b, c)
		return
	}
}

// DropLast351 with func( 3 in)(5 out) drop last 1 result
func DropLast351[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (D, E, F, G) {
	return func(a A, b B, c C) (d D, e E, f F, g G) {
		d, e, f, g, _ = fn(a, b, c)
		return
	}
}

// Drop352 with func( 3 in)(5 out) drop first 2 result
func Drop352[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (F, G, H) {
	return func(a A, b B, c C) (f F, g G, h H) {
		_, _, f, g, h = fn(a, b, c)
		return
	}
}

// DropLast352 with func( 3 in)(5 out) drop last 2 result
func DropLast352[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (D, E, F) {
	return func(a A, b B, c C) (d D, e E, f F) {
		d, e, f, _, _ = fn(a, b, c)
		return
	}
}

// Drop353 with func( 3 in)(5 out) drop first 3 result
func Drop353[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (G, H) {
	return func(a A, b B, c C) (g G, h H) {
		_, _, _, g, h = fn(a, b, c)
		return
	}
}

// DropLast353 with func( 3 in)(5 out) drop last 3 result
func DropLast353[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop354 with func( 3 in)(5 out) drop first 4 result
func Drop354[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) H {
	return func(a A, b B, c C) (h H) {
		_, _, _, _, h = fn(a, b, c)
		return
	}
}

// DropLast354 with func( 3 in)(5 out) drop last 4 result
func DropLast354[A, B, C, D, E, F, G, H any](fn func(A, B, C) (D, E, F, G, H)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop421 with func( 4 in)(2 out) drop first 1 result
func Drop421[A, B, C, D, E, F any](fn func(A, B, C, D) (E, F)) func(a A, b B, c C, d D) F {
	return func(a A, b B, c C, d D) (f F) {
		_, f = fn(a, b, c, d)
		return
	}
}

// DropLast421 with func( 4 in)(2 out) drop last 1 result
func DropLast421[A, B, C, D, E, F any](fn func(A, B, C, D) (E, F)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _ = fn(a, b, c, d)
		return
	}
}

// Drop431 with func( 4 in)(3 out) drop first 1 result
func Drop431[A, B, C, D, E, F, G any](fn func(A, B, C, D) (E, F, G)) func(a A, b B, c C, d D) (F, G) {
	return func(a A, b B, c C, d D) (f F, g G) {
		_, f, g = fn(a, b, c, d)
		return
	}
}

// DropLast431 with func( 4 in)(3 out) drop last 1 result
func DropLast431[A, B, C, D, E, F, G any](fn func(A, B, C, D) (E, F, G)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _ = fn(a, b, c, d)
		return
	}
}

// Drop432 with func( 4 in)(3 out) drop first 2 result
func Drop432[A, B, C, D, E, F, G any](fn func(A, B, C, D) (E, F, G)) func(a A, b B, c C, d D) G {
	return func(a A, b B, c C, d D) (g G) {
		_, _, g = fn(a, b, c, d)
		return
	}
}

// DropLast432 with func( 4 in)(3 out) drop last 2 result
func DropLast432[A, B, C, D, E, F, G any](fn func(A, B, C, D) (E, F, G)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop441 with func( 4 in)(4 out) drop first 1 result
func Drop441[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) (F, G, H) {
	return func(a A, b B, c C, d D) (f F, g G, h H) {
		_, f, g, h = fn(a, b, c, d)
		return
	}
}

// DropLast441 with func( 4 in)(4 out) drop last 1 result
func DropLast441[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) (E, F, G) {
	return func(a A, b B, c C, d D) (e E, f F, g G) {
		e, f, g, _ = fn(a, b, c, d)
		return
	}
}

// Drop442 with func( 4 in)(4 out) drop first 2 result
func Drop442[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) (G, H) {
	return func(a A, b B, c C, d D) (g G, h H) {
		_, _, g, h = fn(a, b, c, d)
		return
	}
}

// DropLast442 with func( 4 in)(4 out) drop last 2 result
func DropLast442[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop443 with func( 4 in)(4 out) drop first 3 result
func Drop443[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) H {
	return func(a A, b B, c C, d D) (h H) {
		_, _, _, h = fn(a, b, c, d)
		return
	}
}

// DropLast443 with func( 4 in)(4 out) drop last 3 result
func DropLast443[A, B, C, D, E, F, G, H any](fn func(A, B, C, D) (E, F, G, H)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop451 with func( 4 in)(5 out) drop first 1 result
func Drop451[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (F, G, H, I) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I) {
		_, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// DropLast451 with func( 4 in)(5 out) drop last 1 result
func DropLast451[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (E, F, G, H) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h, _ = fn(a, b, c, d)
		return
	}
}

// Drop452 with func( 4 in)(5 out) drop first 2 result
func Drop452[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (G, H, I) {
	return func(a A, b B, c C, d D) (g G, h H, i I) {
		_, _, g, h, i = fn(a, b, c, d)
		return
	}
}

// DropLast452 with func( 4 in)(5 out) drop last 2 result
func DropLast452[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (E, F, G) {
	return func(a A, b B, c C, d D) (e E, f F, g G) {
		e, f, g, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop453 with func( 4 in)(5 out) drop first 3 result
func Drop453[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (H, I) {
	return func(a A, b B, c C, d D) (h H, i I) {
		_, _, _, h, i = fn(a, b, c, d)
		return
	}
}

// DropLast453 with func( 4 in)(5 out) drop last 3 result
func DropLast453[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop454 with func( 4 in)(5 out) drop first 4 result
func Drop454[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) I {
	return func(a A, b B, c C, d D) (i I) {
		_, _, _, _, i = fn(a, b, c, d)
		return
	}
}

// DropLast454 with func( 4 in)(5 out) drop last 4 result
func DropLast454[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D) (E, F, G, H, I)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop521 with func( 5 in)(2 out) drop first 1 result
func Drop521[A, B, C, D, E, F, G any](fn func(A, B, C, D, E) (F, G)) func(a A, b B, c C, d D, e E) G {
	return func(a A, b B, c C, d D, e E) (g G) {
		_, g = fn(a, b, c, d, e)
		return
	}
}

// DropLast521 with func( 5 in)(2 out) drop last 1 result
func DropLast521[A, B, C, D, E, F, G any](fn func(A, B, C, D, E) (F, G)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop531 with func( 5 in)(3 out) drop first 1 result
func Drop531[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E) (F, G, H)) func(a A, b B, c C, d D, e E) (G, H) {
	return func(a A, b B, c C, d D, e E) (g G, h H) {
		_, g, h = fn(a, b, c, d, e)
		return
	}
}

// DropLast531 with func( 5 in)(3 out) drop last 1 result
func DropLast531[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E) (F, G, H)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop532 with func( 5 in)(3 out) drop first 2 result
func Drop532[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E) (F, G, H)) func(a A, b B, c C, d D, e E) H {
	return func(a A, b B, c C, d D, e E) (h H) {
		_, _, h = fn(a, b, c, d, e)
		return
	}
}

// DropLast532 with func( 5 in)(3 out) drop last 2 result
func DropLast532[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E) (F, G, H)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop541 with func( 5 in)(4 out) drop first 1 result
func Drop541[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) (G, H, I) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I) {
		_, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// DropLast541 with func( 5 in)(4 out) drop last 1 result
func DropLast541[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) (F, G, H) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop542 with func( 5 in)(4 out) drop first 2 result
func Drop542[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) (H, I) {
	return func(a A, b B, c C, d D, e E) (h H, i I) {
		_, _, h, i = fn(a, b, c, d, e)
		return
	}
}

// DropLast542 with func( 5 in)(4 out) drop last 2 result
func DropLast542[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop543 with func( 5 in)(4 out) drop first 3 result
func Drop543[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) I {
	return func(a A, b B, c C, d D, e E) (i I) {
		_, _, _, i = fn(a, b, c, d, e)
		return
	}
}

// DropLast543 with func( 5 in)(4 out) drop last 3 result
func DropLast543[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E) (F, G, H, I)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop551 with func( 5 in)(5 out) drop first 1 result
func Drop551[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J) {
		_, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// DropLast551 with func( 5 in)(5 out) drop last 1 result
func DropLast551[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (F, G, H, I) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop552 with func( 5 in)(5 out) drop first 2 result
func Drop552[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (H, I, J) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J) {
		_, _, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// DropLast552 with func( 5 in)(5 out) drop last 2 result
func DropLast552[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (F, G, H) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop553 with func( 5 in)(5 out) drop first 3 result
func Drop553[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (I, J) {
	return func(a A, b B, c C, d D, e E) (i I, j J) {
		_, _, _, i, j = fn(a, b, c, d, e)
		return
	}
}

// DropLast553 with func( 5 in)(5 out) drop last 3 result
func DropLast553[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop554 with func( 5 in)(5 out) drop first 4 result
func Drop554[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) J {
	return func(a A, b B, c C, d D, e E) (j J) {
		_, _, _, _, j = fn(a, b, c, d, e)
		return
	}
}

// DropLast554 with func( 5 in)(5 out) drop last 4 result
func DropLast554[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E) (F, G, H, I, J)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

//endregion

//region Closure

// Closure201 with func( 2 in)(0 out) fix first 1 argument
func Closure201[A, B any](a A, fn func(A, B)) func(B) {
	return func(b B) {
		fn(a, b)
		return
	}
}

// ClosureLast201 with func( 2 in)(0 out) fix last 1 argument
func ClosureLast201[A, B any](b B, fn func(A, B)) func(A) {
	return func(a A) {
		fn(a, b)
		return
	}
}

// Closure211 with func( 2 in)(1 out) fix first 1 argument
func Closure211[A, B, C any](a A, fn func(A, B) C) func(B) C {
	return func(b B) (c C) {
		c = fn(a, b)
		return
	}
}

// ClosureLast211 with func( 2 in)(1 out) fix last 1 argument
func ClosureLast211[A, B, C any](b B, fn func(A, B) C) func(A) C {
	return func(a A) (c C) {
		c = fn(a, b)
		return
	}
}

// Closure221 with func( 2 in)(2 out) fix first 1 argument
func Closure221[A, B, C, D any](a A, fn func(A, B) (C, D)) func(B) (C, D) {
	return func(b B) (c C, d D) {
		c, d = fn(a, b)
		return
	}
}

// ClosureLast221 with func( 2 in)(2 out) fix last 1 argument
func ClosureLast221[A, B, C, D any](b B, fn func(A, B) (C, D)) func(A) (C, D) {
	return func(a A) (c C, d D) {
		c, d = fn(a, b)
		return
	}
}

// Closure231 with func( 2 in)(3 out) fix first 1 argument
func Closure231[A, B, C, D, E any](a A, fn func(A, B) (C, D, E)) func(B) (C, D, E) {
	return func(b B) (c C, d D, e E) {
		c, d, e = fn(a, b)
		return
	}
}

// ClosureLast231 with func( 2 in)(3 out) fix last 1 argument
func ClosureLast231[A, B, C, D, E any](b B, fn func(A, B) (C, D, E)) func(A) (C, D, E) {
	return func(a A) (c C, d D, e E) {
		c, d, e = fn(a, b)
		return
	}
}

// Closure241 with func( 2 in)(4 out) fix first 1 argument
func Closure241[A, B, C, D, E, F any](a A, fn func(A, B) (C, D, E, F)) func(B) (C, D, E, F) {
	return func(b B) (c C, d D, e E, f F) {
		c, d, e, f = fn(a, b)
		return
	}
}

// ClosureLast241 with func( 2 in)(4 out) fix last 1 argument
func ClosureLast241[A, B, C, D, E, F any](b B, fn func(A, B) (C, D, E, F)) func(A) (C, D, E, F) {
	return func(a A) (c C, d D, e E, f F) {
		c, d, e, f = fn(a, b)
		return
	}
}

// Closure251 with func( 2 in)(5 out) fix first 1 argument
func Closure251[A, B, C, D, E, F, G any](a A, fn func(A, B) (C, D, E, F, G)) func(B) (C, D, E, F, G) {
	return func(b B) (c C, d D, e E, f F, g G) {
		c, d, e, f, g = fn(a, b)
		return
	}
}

// ClosureLast251 with func( 2 in)(5 out) fix last 1 argument
func ClosureLast251[A, B, C, D, E, F, G any](b B, fn func(A, B) (C, D, E, F, G)) func(A) (C, D, E, F, G) {
	return func(a A) (c C, d D, e E, f F, g G) {
		c, d, e, f, g = fn(a, b)
		return
	}
}

// Closure301 with func( 3 in)(0 out) fix first 1 argument
func Closure301[A, B, C any](a A, b B, fn func(A, B, C)) func(C) {
	return func(c C) {
		fn(a, b, c)
		return
	}
}

// ClosureLast301 with func( 3 in)(0 out) fix last 1 argument
func ClosureLast301[A, B, C any](c C, fn func(A, B, C)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c)
		return
	}
}

// Closure302 with func( 3 in)(0 out) fix first 2 argument
func Closure302[A, B, C any](a A, fn func(A, B, C)) func(B, C) {
	return func(b B, c C) {
		fn(a, b, c)
		return
	}
}

// ClosureLast302 with func( 3 in)(0 out) fix last 2 argument
func ClosureLast302[A, B, C any](b B, c C, fn func(A, B, C)) func(A) {
	return func(a A) {
		fn(a, b, c)
		return
	}
}

// Closure311 with func( 3 in)(1 out) fix first 1 argument
func Closure311[A, B, C, D any](a A, b B, fn func(A, B, C) D) func(C) D {
	return func(c C) (d D) {
		d = fn(a, b, c)
		return
	}
}

// ClosureLast311 with func( 3 in)(1 out) fix last 1 argument
func ClosureLast311[A, B, C, D any](c C, fn func(A, B, C) D) func(A, B) D {
	return func(a A, b B) (d D) {
		d = fn(a, b, c)
		return
	}
}

// Closure312 with func( 3 in)(1 out) fix first 2 argument
func Closure312[A, B, C, D any](a A, fn func(A, B, C) D) func(B, C) D {
	return func(b B, c C) (d D) {
		d = fn(a, b, c)
		return
	}
}

// ClosureLast312 with func( 3 in)(1 out) fix last 2 argument
func ClosureLast312[A, B, C, D any](b B, c C, fn func(A, B, C) D) func(A) D {
	return func(a A) (d D) {
		d = fn(a, b, c)
		return
	}
}

// Closure321 with func( 3 in)(2 out) fix first 1 argument
func Closure321[A, B, C, D, E any](a A, b B, fn func(A, B, C) (D, E)) func(C) (D, E) {
	return func(c C) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// ClosureLast321 with func( 3 in)(2 out) fix last 1 argument
func ClosureLast321[A, B, C, D, E any](c C, fn func(A, B, C) (D, E)) func(A, B) (D, E) {
	return func(a A, b B) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// Closure322 with func( 3 in)(2 out) fix first 2 argument
func Closure322[A, B, C, D, E any](a A, fn func(A, B, C) (D, E)) func(B, C) (D, E) {
	return func(b B, c C) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// ClosureLast322 with func( 3 in)(2 out) fix last 2 argument
func ClosureLast322[A, B, C, D, E any](b B, c C, fn func(A, B, C) (D, E)) func(A) (D, E) {
	return func(a A) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// Closure331 with func( 3 in)(3 out) fix first 1 argument
func Closure331[A, B, C, D, E, F any](a A, b B, fn func(A, B, C) (D, E, F)) func(C) (D, E, F) {
	return func(c C) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// ClosureLast331 with func( 3 in)(3 out) fix last 1 argument
func ClosureLast331[A, B, C, D, E, F any](c C, fn func(A, B, C) (D, E, F)) func(A, B) (D, E, F) {
	return func(a A, b B) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// Closure332 with func( 3 in)(3 out) fix first 2 argument
func Closure332[A, B, C, D, E, F any](a A, fn func(A, B, C) (D, E, F)) func(B, C) (D, E, F) {
	return func(b B, c C) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// ClosureLast332 with func( 3 in)(3 out) fix last 2 argument
func ClosureLast332[A, B, C, D, E, F any](b B, c C, fn func(A, B, C) (D, E, F)) func(A) (D, E, F) {
	return func(a A) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// Closure341 with func( 3 in)(4 out) fix first 1 argument
func Closure341[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C) (D, E, F, G)) func(C) (D, E, F, G) {
	return func(c C) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// ClosureLast341 with func( 3 in)(4 out) fix last 1 argument
func ClosureLast341[A, B, C, D, E, F, G any](c C, fn func(A, B, C) (D, E, F, G)) func(A, B) (D, E, F, G) {
	return func(a A, b B) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Closure342 with func( 3 in)(4 out) fix first 2 argument
func Closure342[A, B, C, D, E, F, G any](a A, fn func(A, B, C) (D, E, F, G)) func(B, C) (D, E, F, G) {
	return func(b B, c C) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// ClosureLast342 with func( 3 in)(4 out) fix last 2 argument
func ClosureLast342[A, B, C, D, E, F, G any](b B, c C, fn func(A, B, C) (D, E, F, G)) func(A) (D, E, F, G) {
	return func(a A) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Closure351 with func( 3 in)(5 out) fix first 1 argument
func Closure351[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C) (D, E, F, G, H)) func(C) (D, E, F, G, H) {
	return func(c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// ClosureLast351 with func( 3 in)(5 out) fix last 1 argument
func ClosureLast351[A, B, C, D, E, F, G, H any](c C, fn func(A, B, C) (D, E, F, G, H)) func(A, B) (D, E, F, G, H) {
	return func(a A, b B) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Closure352 with func( 3 in)(5 out) fix first 2 argument
func Closure352[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C) (D, E, F, G, H)) func(B, C) (D, E, F, G, H) {
	return func(b B, c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// ClosureLast352 with func( 3 in)(5 out) fix last 2 argument
func ClosureLast352[A, B, C, D, E, F, G, H any](b B, c C, fn func(A, B, C) (D, E, F, G, H)) func(A) (D, E, F, G, H) {
	return func(a A) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Closure401 with func( 4 in)(0 out) fix first 1 argument
func Closure401[A, B, C, D any](a A, b B, c C, fn func(A, B, C, D)) func(D) {
	return func(d D) {
		fn(a, b, c, d)
		return
	}
}

// ClosureLast401 with func( 4 in)(0 out) fix last 1 argument
func ClosureLast401[A, B, C, D any](d D, fn func(A, B, C, D)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d)
		return
	}
}

// Closure402 with func( 4 in)(0 out) fix first 2 argument
func Closure402[A, B, C, D any](a A, b B, fn func(A, B, C, D)) func(C, D) {
	return func(c C, d D) {
		fn(a, b, c, d)
		return
	}
}

// ClosureLast402 with func( 4 in)(0 out) fix last 2 argument
func ClosureLast402[A, B, C, D any](c C, d D, fn func(A, B, C, D)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d)
		return
	}
}

// Closure403 with func( 4 in)(0 out) fix first 3 argument
func Closure403[A, B, C, D any](a A, fn func(A, B, C, D)) func(B, C, D) {
	return func(b B, c C, d D) {
		fn(a, b, c, d)
		return
	}
}

// ClosureLast403 with func( 4 in)(0 out) fix last 3 argument
func ClosureLast403[A, B, C, D any](b B, c C, d D, fn func(A, B, C, D)) func(A) {
	return func(a A) {
		fn(a, b, c, d)
		return
	}
}

// Closure411 with func( 4 in)(1 out) fix first 1 argument
func Closure411[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C, D) E) func(D) E {
	return func(d D) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// ClosureLast411 with func( 4 in)(1 out) fix last 1 argument
func ClosureLast411[A, B, C, D, E any](d D, fn func(A, B, C, D) E) func(A, B, C) E {
	return func(a A, b B, c C) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Closure412 with func( 4 in)(1 out) fix first 2 argument
func Closure412[A, B, C, D, E any](a A, b B, fn func(A, B, C, D) E) func(C, D) E {
	return func(c C, d D) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// ClosureLast412 with func( 4 in)(1 out) fix last 2 argument
func ClosureLast412[A, B, C, D, E any](c C, d D, fn func(A, B, C, D) E) func(A, B) E {
	return func(a A, b B) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Closure413 with func( 4 in)(1 out) fix first 3 argument
func Closure413[A, B, C, D, E any](a A, fn func(A, B, C, D) E) func(B, C, D) E {
	return func(b B, c C, d D) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// ClosureLast413 with func( 4 in)(1 out) fix last 3 argument
func ClosureLast413[A, B, C, D, E any](b B, c C, d D, fn func(A, B, C, D) E) func(A) E {
	return func(a A) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Closure421 with func( 4 in)(2 out) fix first 1 argument
func Closure421[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D) (E, F)) func(D) (E, F) {
	return func(d D) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// ClosureLast421 with func( 4 in)(2 out) fix last 1 argument
func ClosureLast421[A, B, C, D, E, F any](d D, fn func(A, B, C, D) (E, F)) func(A, B, C) (E, F) {
	return func(a A, b B, c C) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Closure422 with func( 4 in)(2 out) fix first 2 argument
func Closure422[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D) (E, F)) func(C, D) (E, F) {
	return func(c C, d D) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// ClosureLast422 with func( 4 in)(2 out) fix last 2 argument
func ClosureLast422[A, B, C, D, E, F any](c C, d D, fn func(A, B, C, D) (E, F)) func(A, B) (E, F) {
	return func(a A, b B) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Closure423 with func( 4 in)(2 out) fix first 3 argument
func Closure423[A, B, C, D, E, F any](a A, fn func(A, B, C, D) (E, F)) func(B, C, D) (E, F) {
	return func(b B, c C, d D) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// ClosureLast423 with func( 4 in)(2 out) fix last 3 argument
func ClosureLast423[A, B, C, D, E, F any](b B, c C, d D, fn func(A, B, C, D) (E, F)) func(A) (E, F) {
	return func(a A) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Closure431 with func( 4 in)(3 out) fix first 1 argument
func Closure431[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D) (E, F, G)) func(D) (E, F, G) {
	return func(d D) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// ClosureLast431 with func( 4 in)(3 out) fix last 1 argument
func ClosureLast431[A, B, C, D, E, F, G any](d D, fn func(A, B, C, D) (E, F, G)) func(A, B, C) (E, F, G) {
	return func(a A, b B, c C) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Closure432 with func( 4 in)(3 out) fix first 2 argument
func Closure432[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D) (E, F, G)) func(C, D) (E, F, G) {
	return func(c C, d D) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// ClosureLast432 with func( 4 in)(3 out) fix last 2 argument
func ClosureLast432[A, B, C, D, E, F, G any](c C, d D, fn func(A, B, C, D) (E, F, G)) func(A, B) (E, F, G) {
	return func(a A, b B) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Closure433 with func( 4 in)(3 out) fix first 3 argument
func Closure433[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D) (E, F, G)) func(B, C, D) (E, F, G) {
	return func(b B, c C, d D) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// ClosureLast433 with func( 4 in)(3 out) fix last 3 argument
func ClosureLast433[A, B, C, D, E, F, G any](b B, c C, d D, fn func(A, B, C, D) (E, F, G)) func(A) (E, F, G) {
	return func(a A) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Closure441 with func( 4 in)(4 out) fix first 1 argument
func Closure441[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H)) func(D) (E, F, G, H) {
	return func(d D) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// ClosureLast441 with func( 4 in)(4 out) fix last 1 argument
func ClosureLast441[A, B, C, D, E, F, G, H any](d D, fn func(A, B, C, D) (E, F, G, H)) func(A, B, C) (E, F, G, H) {
	return func(a A, b B, c C) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Closure442 with func( 4 in)(4 out) fix first 2 argument
func Closure442[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D) (E, F, G, H)) func(C, D) (E, F, G, H) {
	return func(c C, d D) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// ClosureLast442 with func( 4 in)(4 out) fix last 2 argument
func ClosureLast442[A, B, C, D, E, F, G, H any](c C, d D, fn func(A, B, C, D) (E, F, G, H)) func(A, B) (E, F, G, H) {
	return func(a A, b B) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Closure443 with func( 4 in)(4 out) fix first 3 argument
func Closure443[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D) (E, F, G, H)) func(B, C, D) (E, F, G, H) {
	return func(b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// ClosureLast443 with func( 4 in)(4 out) fix last 3 argument
func ClosureLast443[A, B, C, D, E, F, G, H any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H)) func(A) (E, F, G, H) {
	return func(a A) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Closure451 with func( 4 in)(5 out) fix first 1 argument
func Closure451[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I)) func(D) (E, F, G, H, I) {
	return func(d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// ClosureLast451 with func( 4 in)(5 out) fix last 1 argument
func ClosureLast451[A, B, C, D, E, F, G, H, I any](d D, fn func(A, B, C, D) (E, F, G, H, I)) func(A, B, C) (E, F, G, H, I) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Closure452 with func( 4 in)(5 out) fix first 2 argument
func Closure452[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I)) func(C, D) (E, F, G, H, I) {
	return func(c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// ClosureLast452 with func( 4 in)(5 out) fix last 2 argument
func ClosureLast452[A, B, C, D, E, F, G, H, I any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I)) func(A, B) (E, F, G, H, I) {
	return func(a A, b B) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Closure453 with func( 4 in)(5 out) fix first 3 argument
func Closure453[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D) (E, F, G, H, I)) func(B, C, D) (E, F, G, H, I) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// ClosureLast453 with func( 4 in)(5 out) fix last 3 argument
func ClosureLast453[A, B, C, D, E, F, G, H, I any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I)) func(A) (E, F, G, H, I) {
	return func(a A) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Closure501 with func( 5 in)(0 out) fix first 1 argument
func Closure501[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D, E)) func(E) {
	return func(e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosureLast501 with func( 5 in)(0 out) fix last 1 argument
func ClosureLast501[A, B, C, D, E any](e E, fn func(A, B, C, D, E)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e)
		return
	}
}

// Closure502 with func( 5 in)(0 out) fix first 2 argument
func Closure502[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C, D, E)) func(D, E) {
	return func(d D, e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosureLast502 with func( 5 in)(0 out) fix last 2 argument
func ClosureLast502[A, B, C, D, E any](d D, e E, fn func(A, B, C, D, E)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e)
		return
	}
}

// Closure503 with func( 5 in)(0 out) fix first 3 argument
func Closure503[A, B, C, D, E any](a A, b B, fn func(A, B, C, D, E)) func(C, D, E) {
	return func(c C, d D, e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosureLast503 with func( 5 in)(0 out) fix last 3 argument
func ClosureLast503[A, B, C, D, E any](c C, d D, e E, fn func(A, B, C, D, E)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e)
		return
	}
}

// Closure504 with func( 5 in)(0 out) fix first 4 argument
func Closure504[A, B, C, D, E any](a A, fn func(A, B, C, D, E)) func(B, C, D, E) {
	return func(b B, c C, d D, e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosureLast504 with func( 5 in)(0 out) fix last 4 argument
func ClosureLast504[A, B, C, D, E any](b B, c C, d D, e E, fn func(A, B, C, D, E)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e)
		return
	}
}

// Closure511 with func( 5 in)(1 out) fix first 1 argument
func Closure511[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D, E) F) func(E) F {
	return func(e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast511 with func( 5 in)(1 out) fix last 1 argument
func ClosureLast511[A, B, C, D, E, F any](e E, fn func(A, B, C, D, E) F) func(A, B, C, D) F {
	return func(a A, b B, c C, d D) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Closure512 with func( 5 in)(1 out) fix first 2 argument
func Closure512[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D, E) F) func(D, E) F {
	return func(d D, e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast512 with func( 5 in)(1 out) fix last 2 argument
func ClosureLast512[A, B, C, D, E, F any](d D, e E, fn func(A, B, C, D, E) F) func(A, B, C) F {
	return func(a A, b B, c C) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Closure513 with func( 5 in)(1 out) fix first 3 argument
func Closure513[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D, E) F) func(C, D, E) F {
	return func(c C, d D, e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast513 with func( 5 in)(1 out) fix last 3 argument
func ClosureLast513[A, B, C, D, E, F any](c C, d D, e E, fn func(A, B, C, D, E) F) func(A, B) F {
	return func(a A, b B) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Closure514 with func( 5 in)(1 out) fix first 4 argument
func Closure514[A, B, C, D, E, F any](a A, fn func(A, B, C, D, E) F) func(B, C, D, E) F {
	return func(b B, c C, d D, e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast514 with func( 5 in)(1 out) fix last 4 argument
func ClosureLast514[A, B, C, D, E, F any](b B, c C, d D, e E, fn func(A, B, C, D, E) F) func(A) F {
	return func(a A) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Closure521 with func( 5 in)(2 out) fix first 1 argument
func Closure521[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G)) func(E) (F, G) {
	return func(e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast521 with func( 5 in)(2 out) fix last 1 argument
func ClosureLast521[A, B, C, D, E, F, G any](e E, fn func(A, B, C, D, E) (F, G)) func(A, B, C, D) (F, G) {
	return func(a A, b B, c C, d D) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Closure522 with func( 5 in)(2 out) fix first 2 argument
func Closure522[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E) (F, G)) func(D, E) (F, G) {
	return func(d D, e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast522 with func( 5 in)(2 out) fix last 2 argument
func ClosureLast522[A, B, C, D, E, F, G any](d D, e E, fn func(A, B, C, D, E) (F, G)) func(A, B, C) (F, G) {
	return func(a A, b B, c C) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Closure523 with func( 5 in)(2 out) fix first 3 argument
func Closure523[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E) (F, G)) func(C, D, E) (F, G) {
	return func(c C, d D, e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast523 with func( 5 in)(2 out) fix last 3 argument
func ClosureLast523[A, B, C, D, E, F, G any](c C, d D, e E, fn func(A, B, C, D, E) (F, G)) func(A, B) (F, G) {
	return func(a A, b B) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Closure524 with func( 5 in)(2 out) fix first 4 argument
func Closure524[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E) (F, G)) func(B, C, D, E) (F, G) {
	return func(b B, c C, d D, e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast524 with func( 5 in)(2 out) fix last 4 argument
func ClosureLast524[A, B, C, D, E, F, G any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G)) func(A) (F, G) {
	return func(a A) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Closure531 with func( 5 in)(3 out) fix first 1 argument
func Closure531[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H)) func(E) (F, G, H) {
	return func(e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast531 with func( 5 in)(3 out) fix last 1 argument
func ClosureLast531[A, B, C, D, E, F, G, H any](e E, fn func(A, B, C, D, E) (F, G, H)) func(A, B, C, D) (F, G, H) {
	return func(a A, b B, c C, d D) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Closure532 with func( 5 in)(3 out) fix first 2 argument
func Closure532[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H)) func(D, E) (F, G, H) {
	return func(d D, e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast532 with func( 5 in)(3 out) fix last 2 argument
func ClosureLast532[A, B, C, D, E, F, G, H any](d D, e E, fn func(A, B, C, D, E) (F, G, H)) func(A, B, C) (F, G, H) {
	return func(a A, b B, c C) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Closure533 with func( 5 in)(3 out) fix first 3 argument
func Closure533[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E) (F, G, H)) func(C, D, E) (F, G, H) {
	return func(c C, d D, e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast533 with func( 5 in)(3 out) fix last 3 argument
func ClosureLast533[A, B, C, D, E, F, G, H any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H)) func(A, B) (F, G, H) {
	return func(a A, b B) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Closure534 with func( 5 in)(3 out) fix first 4 argument
func Closure534[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E) (F, G, H)) func(B, C, D, E) (F, G, H) {
	return func(b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast534 with func( 5 in)(3 out) fix last 4 argument
func ClosureLast534[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H)) func(A) (F, G, H) {
	return func(a A) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Closure541 with func( 5 in)(4 out) fix first 1 argument
func Closure541[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I)) func(E) (F, G, H, I) {
	return func(e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast541 with func( 5 in)(4 out) fix last 1 argument
func ClosureLast541[A, B, C, D, E, F, G, H, I any](e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A, B, C, D) (F, G, H, I) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Closure542 with func( 5 in)(4 out) fix first 2 argument
func Closure542[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I)) func(D, E) (F, G, H, I) {
	return func(d D, e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast542 with func( 5 in)(4 out) fix last 2 argument
func ClosureLast542[A, B, C, D, E, F, G, H, I any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A, B, C) (F, G, H, I) {
	return func(a A, b B, c C) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Closure543 with func( 5 in)(4 out) fix first 3 argument
func Closure543[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I)) func(C, D, E) (F, G, H, I) {
	return func(c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast543 with func( 5 in)(4 out) fix last 3 argument
func ClosureLast543[A, B, C, D, E, F, G, H, I any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A, B) (F, G, H, I) {
	return func(a A, b B) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Closure544 with func( 5 in)(4 out) fix first 4 argument
func Closure544[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E) (F, G, H, I)) func(B, C, D, E) (F, G, H, I) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast544 with func( 5 in)(4 out) fix last 4 argument
func ClosureLast544[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A) (F, G, H, I) {
	return func(a A) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Closure551 with func( 5 in)(5 out) fix first 1 argument
func Closure551[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J)) func(E) (F, G, H, I, J) {
	return func(e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast551 with func( 5 in)(5 out) fix last 1 argument
func ClosureLast551[A, B, C, D, E, F, G, H, I, J any](e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A, B, C, D) (F, G, H, I, J) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Closure552 with func( 5 in)(5 out) fix first 2 argument
func Closure552[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J)) func(D, E) (F, G, H, I, J) {
	return func(d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast552 with func( 5 in)(5 out) fix last 2 argument
func ClosureLast552[A, B, C, D, E, F, G, H, I, J any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A, B, C) (F, G, H, I, J) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Closure553 with func( 5 in)(5 out) fix first 3 argument
func Closure553[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J)) func(C, D, E) (F, G, H, I, J) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast553 with func( 5 in)(5 out) fix last 3 argument
func ClosureLast553[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A, B) (F, G, H, I, J) {
	return func(a A, b B) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Closure554 with func( 5 in)(5 out) fix first 4 argument
func Closure554[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E) (F, G, H, I, J)) func(B, C, D, E) (F, G, H, I, J) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast554 with func( 5 in)(5 out) fix last 4 argument
func ClosureLast554[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A) (F, G, H, I, J) {
	return func(a A) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

//endregion
