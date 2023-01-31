package fn

//generated file should not edit

// Closure101 with func( 1 in)(0 out) closure first 1 argument
func Closure101[A any](a A, fn func(A)) func() {
	return func() {
		fn(a)
		return
	}
}

// Closure111 with func( 1 in)(1 out) closure first 1 argument
func Closure111[A, B any](a A, fn func(A) B) func() B {
	return func() (b B) {
		b = fn(a)
		return
	}
}

// Closure121 with func( 1 in)(2 out) closure first 1 argument
func Closure121[A, B, C any](a A, fn func(A) (B, C)) func() (B, C) {
	return func() (b B, c C) {
		b, c = fn(a)
		return
	}
}

// Closure131 with func( 1 in)(3 out) closure first 1 argument
func Closure131[A, B, C, D any](a A, fn func(A) (B, C, D)) func() (B, C, D) {
	return func() (b B, c C, d D) {
		b, c, d = fn(a)
		return
	}
}

// Closure141 with func( 1 in)(4 out) closure first 1 argument
func Closure141[A, B, C, D, E any](a A, fn func(A) (B, C, D, E)) func() (B, C, D, E) {
	return func() (b B, c C, d D, e E) {
		b, c, d, e = fn(a)
		return
	}
}

// Closure151 with func( 1 in)(5 out) closure first 1 argument
func Closure151[A, B, C, D, E, F any](a A, fn func(A) (B, C, D, E, F)) func() (B, C, D, E, F) {
	return func() (b B, c C, d D, e E, f F) {
		b, c, d, e, f = fn(a)
		return
	}
}

// Closure161 with func( 1 in)(6 out) closure first 1 argument
func Closure161[A, B, C, D, E, F, G any](a A, fn func(A) (B, C, D, E, F, G)) func() (B, C, D, E, F, G) {
	return func() (b B, c C, d D, e E, f F, g G) {
		b, c, d, e, f, g = fn(a)
		return
	}
}

// Closure171 with func( 1 in)(7 out) closure first 1 argument
func Closure171[A, B, C, D, E, F, G, H any](a A, fn func(A) (B, C, D, E, F, G, H)) func() (B, C, D, E, F, G, H) {
	return func() (b B, c C, d D, e E, f F, g G, h H) {
		b, c, d, e, f, g, h = fn(a)
		return
	}
}

// Closure181 with func( 1 in)(8 out) closure first 1 argument
func Closure181[A, B, C, D, E, F, G, H, I any](a A, fn func(A) (B, C, D, E, F, G, H, I)) func() (B, C, D, E, F, G, H, I) {
	return func() (b B, c C, d D, e E, f F, g G, h H, i I) {
		b, c, d, e, f, g, h, i = fn(a)
		return
	}
}

// Closure191 with func( 1 in)(9 out) closure first 1 argument
func Closure191[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A) (B, C, D, E, F, G, H, I, J)) func() (B, C, D, E, F, G, H, I, J) {
	return func() (b B, c C, d D, e E, f F, g G, h H, i I, j J) {
		b, c, d, e, f, g, h, i, j = fn(a)
		return
	}
}

// Closure201 with func( 2 in)(0 out) closure first 1 argument
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

// Closure202 with func( 2 in)(0 out) closure first 2 argument
func Closure202[A, B any](a A, b B, fn func(A, B)) func() {
	return func() {
		fn(a, b)
		return
	}
}

// Closure211 with func( 2 in)(1 out) closure first 1 argument
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

// Closure212 with func( 2 in)(1 out) closure first 2 argument
func Closure212[A, B, C any](a A, b B, fn func(A, B) C) func() C {
	return func() (c C) {
		c = fn(a, b)
		return
	}
}

// Closure221 with func( 2 in)(2 out) closure first 1 argument
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

// Closure222 with func( 2 in)(2 out) closure first 2 argument
func Closure222[A, B, C, D any](a A, b B, fn func(A, B) (C, D)) func() (C, D) {
	return func() (c C, d D) {
		c, d = fn(a, b)
		return
	}
}

// Closure231 with func( 2 in)(3 out) closure first 1 argument
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

// Closure232 with func( 2 in)(3 out) closure first 2 argument
func Closure232[A, B, C, D, E any](a A, b B, fn func(A, B) (C, D, E)) func() (C, D, E) {
	return func() (c C, d D, e E) {
		c, d, e = fn(a, b)
		return
	}
}

// Closure241 with func( 2 in)(4 out) closure first 1 argument
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

// Closure242 with func( 2 in)(4 out) closure first 2 argument
func Closure242[A, B, C, D, E, F any](a A, b B, fn func(A, B) (C, D, E, F)) func() (C, D, E, F) {
	return func() (c C, d D, e E, f F) {
		c, d, e, f = fn(a, b)
		return
	}
}

// Closure251 with func( 2 in)(5 out) closure first 1 argument
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

// Closure252 with func( 2 in)(5 out) closure first 2 argument
func Closure252[A, B, C, D, E, F, G any](a A, b B, fn func(A, B) (C, D, E, F, G)) func() (C, D, E, F, G) {
	return func() (c C, d D, e E, f F, g G) {
		c, d, e, f, g = fn(a, b)
		return
	}
}

// Closure261 with func( 2 in)(6 out) closure first 1 argument
func Closure261[A, B, C, D, E, F, G, H any](a A, fn func(A, B) (C, D, E, F, G, H)) func(B) (C, D, E, F, G, H) {
	return func(b B) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// ClosureLast261 with func( 2 in)(6 out) fix last 1 argument
func ClosureLast261[A, B, C, D, E, F, G, H any](b B, fn func(A, B) (C, D, E, F, G, H)) func(A) (C, D, E, F, G, H) {
	return func(a A) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// Closure262 with func( 2 in)(6 out) closure first 2 argument
func Closure262[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B) (C, D, E, F, G, H)) func() (C, D, E, F, G, H) {
	return func() (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// Closure271 with func( 2 in)(7 out) closure first 1 argument
func Closure271[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B) (C, D, E, F, G, H, I)) func(B) (C, D, E, F, G, H, I) {
	return func(b B) (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// ClosureLast271 with func( 2 in)(7 out) fix last 1 argument
func ClosureLast271[A, B, C, D, E, F, G, H, I any](b B, fn func(A, B) (C, D, E, F, G, H, I)) func(A) (C, D, E, F, G, H, I) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// Closure272 with func( 2 in)(7 out) closure first 2 argument
func Closure272[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I)) func() (C, D, E, F, G, H, I) {
	return func() (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// Closure281 with func( 2 in)(8 out) closure first 1 argument
func Closure281[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B) (C, D, E, F, G, H, I, J)) func(B) (C, D, E, F, G, H, I, J) {
	return func(b B) (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// ClosureLast281 with func( 2 in)(8 out) fix last 1 argument
func ClosureLast281[A, B, C, D, E, F, G, H, I, J any](b B, fn func(A, B) (C, D, E, F, G, H, I, J)) func(A) (C, D, E, F, G, H, I, J) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// Closure282 with func( 2 in)(8 out) closure first 2 argument
func Closure282[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J)) func() (C, D, E, F, G, H, I, J) {
	return func() (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// Closure291 with func( 2 in)(9 out) closure first 1 argument
func Closure291[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(B) (C, D, E, F, G, H, I, J, K) {
	return func(b B) (c C, d D, e E, f F, g G, h H, i I, j J, k K) {
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// ClosureLast291 with func( 2 in)(9 out) fix last 1 argument
func ClosureLast291[A, B, C, D, E, F, G, H, I, J, K any](b B, fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(A) (C, D, E, F, G, H, I, J, K) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I, j J, k K) {
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// Closure292 with func( 2 in)(9 out) closure first 2 argument
func Closure292[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J, K)) func() (C, D, E, F, G, H, I, J, K) {
	return func() (c C, d D, e E, f F, g G, h H, i I, j J, k K) {
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// Closure301 with func( 3 in)(0 out) closure first 1 argument
func Closure301[A, B, C any](a A, fn func(A, B, C)) func(B, C) {
	return func(b B, c C) {
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

// Closure302 with func( 3 in)(0 out) closure first 2 argument
func Closure302[A, B, C any](a A, b B, fn func(A, B, C)) func(C) {
	return func(c C) {
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

// Closure303 with func( 3 in)(0 out) closure first 3 argument
func Closure303[A, B, C any](a A, b B, c C, fn func(A, B, C)) func() {
	return func() {
		fn(a, b, c)
		return
	}
}

// Closure311 with func( 3 in)(1 out) closure first 1 argument
func Closure311[A, B, C, D any](a A, fn func(A, B, C) D) func(B, C) D {
	return func(b B, c C) (d D) {
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

// Closure312 with func( 3 in)(1 out) closure first 2 argument
func Closure312[A, B, C, D any](a A, b B, fn func(A, B, C) D) func(C) D {
	return func(c C) (d D) {
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

// Closure313 with func( 3 in)(1 out) closure first 3 argument
func Closure313[A, B, C, D any](a A, b B, c C, fn func(A, B, C) D) func() D {
	return func() (d D) {
		d = fn(a, b, c)
		return
	}
}

// Closure321 with func( 3 in)(2 out) closure first 1 argument
func Closure321[A, B, C, D, E any](a A, fn func(A, B, C) (D, E)) func(B, C) (D, E) {
	return func(b B, c C) (d D, e E) {
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

// Closure322 with func( 3 in)(2 out) closure first 2 argument
func Closure322[A, B, C, D, E any](a A, b B, fn func(A, B, C) (D, E)) func(C) (D, E) {
	return func(c C) (d D, e E) {
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

// Closure323 with func( 3 in)(2 out) closure first 3 argument
func Closure323[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C) (D, E)) func() (D, E) {
	return func() (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// Closure331 with func( 3 in)(3 out) closure first 1 argument
func Closure331[A, B, C, D, E, F any](a A, fn func(A, B, C) (D, E, F)) func(B, C) (D, E, F) {
	return func(b B, c C) (d D, e E, f F) {
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

// Closure332 with func( 3 in)(3 out) closure first 2 argument
func Closure332[A, B, C, D, E, F any](a A, b B, fn func(A, B, C) (D, E, F)) func(C) (D, E, F) {
	return func(c C) (d D, e E, f F) {
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

// Closure333 with func( 3 in)(3 out) closure first 3 argument
func Closure333[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C) (D, E, F)) func() (D, E, F) {
	return func() (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// Closure341 with func( 3 in)(4 out) closure first 1 argument
func Closure341[A, B, C, D, E, F, G any](a A, fn func(A, B, C) (D, E, F, G)) func(B, C) (D, E, F, G) {
	return func(b B, c C) (d D, e E, f F, g G) {
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

// Closure342 with func( 3 in)(4 out) closure first 2 argument
func Closure342[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C) (D, E, F, G)) func(C) (D, E, F, G) {
	return func(c C) (d D, e E, f F, g G) {
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

// Closure343 with func( 3 in)(4 out) closure first 3 argument
func Closure343[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C) (D, E, F, G)) func() (D, E, F, G) {
	return func() (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Closure351 with func( 3 in)(5 out) closure first 1 argument
func Closure351[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C) (D, E, F, G, H)) func(B, C) (D, E, F, G, H) {
	return func(b B, c C) (d D, e E, f F, g G, h H) {
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

// Closure352 with func( 3 in)(5 out) closure first 2 argument
func Closure352[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C) (D, E, F, G, H)) func(C) (D, E, F, G, H) {
	return func(c C) (d D, e E, f F, g G, h H) {
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

// Closure353 with func( 3 in)(5 out) closure first 3 argument
func Closure353[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H)) func() (D, E, F, G, H) {
	return func() (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Closure361 with func( 3 in)(6 out) closure first 1 argument
func Closure361[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C) (D, E, F, G, H, I)) func(B, C) (D, E, F, G, H, I) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// ClosureLast361 with func( 3 in)(6 out) fix last 1 argument
func ClosureLast361[A, B, C, D, E, F, G, H, I any](c C, fn func(A, B, C) (D, E, F, G, H, I)) func(A, B) (D, E, F, G, H, I) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Closure362 with func( 3 in)(6 out) closure first 2 argument
func Closure362[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I)) func(C) (D, E, F, G, H, I) {
	return func(c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// ClosureLast362 with func( 3 in)(6 out) fix last 2 argument
func ClosureLast362[A, B, C, D, E, F, G, H, I any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I)) func(A) (D, E, F, G, H, I) {
	return func(a A) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Closure363 with func( 3 in)(6 out) closure first 3 argument
func Closure363[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I)) func() (D, E, F, G, H, I) {
	return func() (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Closure371 with func( 3 in)(7 out) closure first 1 argument
func Closure371[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C) (D, E, F, G, H, I, J)) func(B, C) (D, E, F, G, H, I, J) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// ClosureLast371 with func( 3 in)(7 out) fix last 1 argument
func ClosureLast371[A, B, C, D, E, F, G, H, I, J any](c C, fn func(A, B, C) (D, E, F, G, H, I, J)) func(A, B) (D, E, F, G, H, I, J) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Closure372 with func( 3 in)(7 out) closure first 2 argument
func Closure372[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I, J)) func(C) (D, E, F, G, H, I, J) {
	return func(c C) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// ClosureLast372 with func( 3 in)(7 out) fix last 2 argument
func ClosureLast372[A, B, C, D, E, F, G, H, I, J any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J)) func(A) (D, E, F, G, H, I, J) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Closure373 with func( 3 in)(7 out) closure first 3 argument
func Closure373[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J)) func() (D, E, F, G, H, I, J) {
	return func() (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Closure381 with func( 3 in)(8 out) closure first 1 argument
func Closure381[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(B, C) (D, E, F, G, H, I, J, K) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// ClosureLast381 with func( 3 in)(8 out) fix last 1 argument
func ClosureLast381[A, B, C, D, E, F, G, H, I, J, K any](c C, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(A, B) (D, E, F, G, H, I, J, K) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Closure382 with func( 3 in)(8 out) closure first 2 argument
func Closure382[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(C) (D, E, F, G, H, I, J, K) {
	return func(c C) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// ClosureLast382 with func( 3 in)(8 out) fix last 2 argument
func ClosureLast382[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(A) (D, E, F, G, H, I, J, K) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Closure383 with func( 3 in)(8 out) closure first 3 argument
func Closure383[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func() (D, E, F, G, H, I, J, K) {
	return func() (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Closure391 with func( 3 in)(9 out) closure first 1 argument
func Closure391[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(B, C) (D, E, F, G, H, I, J, K, L) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// ClosureLast391 with func( 3 in)(9 out) fix last 1 argument
func ClosureLast391[A, B, C, D, E, F, G, H, I, J, K, L any](c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(A, B) (D, E, F, G, H, I, J, K, L) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Closure392 with func( 3 in)(9 out) closure first 2 argument
func Closure392[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(C) (D, E, F, G, H, I, J, K, L) {
	return func(c C) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// ClosureLast392 with func( 3 in)(9 out) fix last 2 argument
func ClosureLast392[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(A) (D, E, F, G, H, I, J, K, L) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Closure393 with func( 3 in)(9 out) closure first 3 argument
func Closure393[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func() (D, E, F, G, H, I, J, K, L) {
	return func() (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Closure401 with func( 4 in)(0 out) closure first 1 argument
func Closure401[A, B, C, D any](a A, fn func(A, B, C, D)) func(B, C, D) {
	return func(b B, c C, d D) {
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

// Closure402 with func( 4 in)(0 out) closure first 2 argument
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

// Closure403 with func( 4 in)(0 out) closure first 3 argument
func Closure403[A, B, C, D any](a A, b B, c C, fn func(A, B, C, D)) func(D) {
	return func(d D) {
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

// Closure404 with func( 4 in)(0 out) closure first 4 argument
func Closure404[A, B, C, D any](a A, b B, c C, d D, fn func(A, B, C, D)) func() {
	return func() {
		fn(a, b, c, d)
		return
	}
}

// Closure411 with func( 4 in)(1 out) closure first 1 argument
func Closure411[A, B, C, D, E any](a A, fn func(A, B, C, D) E) func(B, C, D) E {
	return func(b B, c C, d D) (e E) {
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

// Closure412 with func( 4 in)(1 out) closure first 2 argument
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

// Closure413 with func( 4 in)(1 out) closure first 3 argument
func Closure413[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C, D) E) func(D) E {
	return func(d D) (e E) {
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

// Closure414 with func( 4 in)(1 out) closure first 4 argument
func Closure414[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D) E) func() E {
	return func() (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Closure421 with func( 4 in)(2 out) closure first 1 argument
func Closure421[A, B, C, D, E, F any](a A, fn func(A, B, C, D) (E, F)) func(B, C, D) (E, F) {
	return func(b B, c C, d D) (e E, f F) {
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

// Closure422 with func( 4 in)(2 out) closure first 2 argument
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

// Closure423 with func( 4 in)(2 out) closure first 3 argument
func Closure423[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D) (E, F)) func(D) (E, F) {
	return func(d D) (e E, f F) {
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

// Closure424 with func( 4 in)(2 out) closure first 4 argument
func Closure424[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F)) func() (E, F) {
	return func() (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Closure431 with func( 4 in)(3 out) closure first 1 argument
func Closure431[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D) (E, F, G)) func(B, C, D) (E, F, G) {
	return func(b B, c C, d D) (e E, f F, g G) {
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

// Closure432 with func( 4 in)(3 out) closure first 2 argument
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

// Closure433 with func( 4 in)(3 out) closure first 3 argument
func Closure433[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D) (E, F, G)) func(D) (E, F, G) {
	return func(d D) (e E, f F, g G) {
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

// Closure434 with func( 4 in)(3 out) closure first 4 argument
func Closure434[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G)) func() (E, F, G) {
	return func() (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Closure441 with func( 4 in)(4 out) closure first 1 argument
func Closure441[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D) (E, F, G, H)) func(B, C, D) (E, F, G, H) {
	return func(b B, c C, d D) (e E, f F, g G, h H) {
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

// Closure442 with func( 4 in)(4 out) closure first 2 argument
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

// Closure443 with func( 4 in)(4 out) closure first 3 argument
func Closure443[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H)) func(D) (E, F, G, H) {
	return func(d D) (e E, f F, g G, h H) {
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

// Closure444 with func( 4 in)(4 out) closure first 4 argument
func Closure444[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H)) func() (E, F, G, H) {
	return func() (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Closure451 with func( 4 in)(5 out) closure first 1 argument
func Closure451[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D) (E, F, G, H, I)) func(B, C, D) (E, F, G, H, I) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I) {
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

// Closure452 with func( 4 in)(5 out) closure first 2 argument
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

// Closure453 with func( 4 in)(5 out) closure first 3 argument
func Closure453[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I)) func(D) (E, F, G, H, I) {
	return func(d D) (e E, f F, g G, h H, i I) {
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

// Closure454 with func( 4 in)(5 out) closure first 4 argument
func Closure454[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I)) func() (E, F, G, H, I) {
	return func() (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Closure461 with func( 4 in)(6 out) closure first 1 argument
func Closure461[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D) (E, F, G, H, I, J)) func(B, C, D) (E, F, G, H, I, J) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// ClosureLast461 with func( 4 in)(6 out) fix last 1 argument
func ClosureLast461[A, B, C, D, E, F, G, H, I, J any](d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func(A, B, C) (E, F, G, H, I, J) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Closure462 with func( 4 in)(6 out) closure first 2 argument
func Closure462[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J)) func(C, D) (E, F, G, H, I, J) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// ClosureLast462 with func( 4 in)(6 out) fix last 2 argument
func ClosureLast462[A, B, C, D, E, F, G, H, I, J any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func(A, B) (E, F, G, H, I, J) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Closure463 with func( 4 in)(6 out) closure first 3 argument
func Closure463[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J)) func(D) (E, F, G, H, I, J) {
	return func(d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// ClosureLast463 with func( 4 in)(6 out) fix last 3 argument
func ClosureLast463[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func(A) (E, F, G, H, I, J) {
	return func(a A) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Closure464 with func( 4 in)(6 out) closure first 4 argument
func Closure464[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func() (E, F, G, H, I, J) {
	return func() (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Closure471 with func( 4 in)(7 out) closure first 1 argument
func Closure471[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(B, C, D) (E, F, G, H, I, J, K) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// ClosureLast471 with func( 4 in)(7 out) fix last 1 argument
func ClosureLast471[A, B, C, D, E, F, G, H, I, J, K any](d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(A, B, C) (E, F, G, H, I, J, K) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Closure472 with func( 4 in)(7 out) closure first 2 argument
func Closure472[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(C, D) (E, F, G, H, I, J, K) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// ClosureLast472 with func( 4 in)(7 out) fix last 2 argument
func ClosureLast472[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(A, B) (E, F, G, H, I, J, K) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Closure473 with func( 4 in)(7 out) closure first 3 argument
func Closure473[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(D) (E, F, G, H, I, J, K) {
	return func(d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// ClosureLast473 with func( 4 in)(7 out) fix last 3 argument
func ClosureLast473[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(A) (E, F, G, H, I, J, K) {
	return func(a A) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Closure474 with func( 4 in)(7 out) closure first 4 argument
func Closure474[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func() (E, F, G, H, I, J, K) {
	return func() (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Closure481 with func( 4 in)(8 out) closure first 1 argument
func Closure481[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(B, C, D) (E, F, G, H, I, J, K, L) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// ClosureLast481 with func( 4 in)(8 out) fix last 1 argument
func ClosureLast481[A, B, C, D, E, F, G, H, I, J, K, L any](d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(A, B, C) (E, F, G, H, I, J, K, L) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Closure482 with func( 4 in)(8 out) closure first 2 argument
func Closure482[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(C, D) (E, F, G, H, I, J, K, L) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// ClosureLast482 with func( 4 in)(8 out) fix last 2 argument
func ClosureLast482[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(A, B) (E, F, G, H, I, J, K, L) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Closure483 with func( 4 in)(8 out) closure first 3 argument
func Closure483[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(D) (E, F, G, H, I, J, K, L) {
	return func(d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// ClosureLast483 with func( 4 in)(8 out) fix last 3 argument
func ClosureLast483[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(A) (E, F, G, H, I, J, K, L) {
	return func(a A) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Closure484 with func( 4 in)(8 out) closure first 4 argument
func Closure484[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func() (E, F, G, H, I, J, K, L) {
	return func() (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Closure491 with func( 4 in)(9 out) closure first 1 argument
func Closure491[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(B, C, D) (E, F, G, H, I, J, K, L, M) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// ClosureLast491 with func( 4 in)(9 out) fix last 1 argument
func ClosureLast491[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(A, B, C) (E, F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Closure492 with func( 4 in)(9 out) closure first 2 argument
func Closure492[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(C, D) (E, F, G, H, I, J, K, L, M) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// ClosureLast492 with func( 4 in)(9 out) fix last 2 argument
func ClosureLast492[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(A, B) (E, F, G, H, I, J, K, L, M) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Closure493 with func( 4 in)(9 out) closure first 3 argument
func Closure493[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(D) (E, F, G, H, I, J, K, L, M) {
	return func(d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// ClosureLast493 with func( 4 in)(9 out) fix last 3 argument
func ClosureLast493[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(A) (E, F, G, H, I, J, K, L, M) {
	return func(a A) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Closure494 with func( 4 in)(9 out) closure first 4 argument
func Closure494[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func() (E, F, G, H, I, J, K, L, M) {
	return func() (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Closure501 with func( 5 in)(0 out) closure first 1 argument
func Closure501[A, B, C, D, E any](a A, fn func(A, B, C, D, E)) func(B, C, D, E) {
	return func(b B, c C, d D, e E) {
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

// Closure502 with func( 5 in)(0 out) closure first 2 argument
func Closure502[A, B, C, D, E any](a A, b B, fn func(A, B, C, D, E)) func(C, D, E) {
	return func(c C, d D, e E) {
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

// Closure503 with func( 5 in)(0 out) closure first 3 argument
func Closure503[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C, D, E)) func(D, E) {
	return func(d D, e E) {
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

// Closure504 with func( 5 in)(0 out) closure first 4 argument
func Closure504[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D, E)) func(E) {
	return func(e E) {
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

// Closure505 with func( 5 in)(0 out) closure first 5 argument
func Closure505[A, B, C, D, E any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E)) func() {
	return func() {
		fn(a, b, c, d, e)
		return
	}
}

// Closure511 with func( 5 in)(1 out) closure first 1 argument
func Closure511[A, B, C, D, E, F any](a A, fn func(A, B, C, D, E) F) func(B, C, D, E) F {
	return func(b B, c C, d D, e E) (f F) {
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

// Closure512 with func( 5 in)(1 out) closure first 2 argument
func Closure512[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D, E) F) func(C, D, E) F {
	return func(c C, d D, e E) (f F) {
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

// Closure513 with func( 5 in)(1 out) closure first 3 argument
func Closure513[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D, E) F) func(D, E) F {
	return func(d D, e E) (f F) {
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

// Closure514 with func( 5 in)(1 out) closure first 4 argument
func Closure514[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D, E) F) func(E) F {
	return func(e E) (f F) {
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

// Closure515 with func( 5 in)(1 out) closure first 5 argument
func Closure515[A, B, C, D, E, F any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) F) func() F {
	return func() (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Closure521 with func( 5 in)(2 out) closure first 1 argument
func Closure521[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E) (F, G)) func(B, C, D, E) (F, G) {
	return func(b B, c C, d D, e E) (f F, g G) {
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

// Closure522 with func( 5 in)(2 out) closure first 2 argument
func Closure522[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E) (F, G)) func(C, D, E) (F, G) {
	return func(c C, d D, e E) (f F, g G) {
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

// Closure523 with func( 5 in)(2 out) closure first 3 argument
func Closure523[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E) (F, G)) func(D, E) (F, G) {
	return func(d D, e E) (f F, g G) {
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

// Closure524 with func( 5 in)(2 out) closure first 4 argument
func Closure524[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G)) func(E) (F, G) {
	return func(e E) (f F, g G) {
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

// Closure525 with func( 5 in)(2 out) closure first 5 argument
func Closure525[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G)) func() (F, G) {
	return func() (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Closure531 with func( 5 in)(3 out) closure first 1 argument
func Closure531[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E) (F, G, H)) func(B, C, D, E) (F, G, H) {
	return func(b B, c C, d D, e E) (f F, g G, h H) {
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

// Closure532 with func( 5 in)(3 out) closure first 2 argument
func Closure532[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E) (F, G, H)) func(C, D, E) (F, G, H) {
	return func(c C, d D, e E) (f F, g G, h H) {
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

// Closure533 with func( 5 in)(3 out) closure first 3 argument
func Closure533[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H)) func(D, E) (F, G, H) {
	return func(d D, e E) (f F, g G, h H) {
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

// Closure534 with func( 5 in)(3 out) closure first 4 argument
func Closure534[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H)) func(E) (F, G, H) {
	return func(e E) (f F, g G, h H) {
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

// Closure535 with func( 5 in)(3 out) closure first 5 argument
func Closure535[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H)) func() (F, G, H) {
	return func() (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Closure541 with func( 5 in)(4 out) closure first 1 argument
func Closure541[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E) (F, G, H, I)) func(B, C, D, E) (F, G, H, I) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I) {
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

// Closure542 with func( 5 in)(4 out) closure first 2 argument
func Closure542[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I)) func(C, D, E) (F, G, H, I) {
	return func(c C, d D, e E) (f F, g G, h H, i I) {
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

// Closure543 with func( 5 in)(4 out) closure first 3 argument
func Closure543[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I)) func(D, E) (F, G, H, I) {
	return func(d D, e E) (f F, g G, h H, i I) {
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

// Closure544 with func( 5 in)(4 out) closure first 4 argument
func Closure544[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I)) func(E) (F, G, H, I) {
	return func(e E) (f F, g G, h H, i I) {
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

// Closure545 with func( 5 in)(4 out) closure first 5 argument
func Closure545[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func() (F, G, H, I) {
	return func() (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Closure551 with func( 5 in)(5 out) closure first 1 argument
func Closure551[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E) (F, G, H, I, J)) func(B, C, D, E) (F, G, H, I, J) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
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

// Closure552 with func( 5 in)(5 out) closure first 2 argument
func Closure552[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J)) func(C, D, E) (F, G, H, I, J) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J) {
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

// Closure553 with func( 5 in)(5 out) closure first 3 argument
func Closure553[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J)) func(D, E) (F, G, H, I, J) {
	return func(d D, e E) (f F, g G, h H, i I, j J) {
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

// Closure554 with func( 5 in)(5 out) closure first 4 argument
func Closure554[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J)) func(E) (F, G, H, I, J) {
	return func(e E) (f F, g G, h H, i I, j J) {
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

// Closure555 with func( 5 in)(5 out) closure first 5 argument
func Closure555[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func() (F, G, H, I, J) {
	return func() (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Closure561 with func( 5 in)(6 out) closure first 1 argument
func Closure561[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(B, C, D, E) (F, G, H, I, J, K) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast561 with func( 5 in)(6 out) fix last 1 argument
func ClosureLast561[A, B, C, D, E, F, G, H, I, J, K any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A, B, C, D) (F, G, H, I, J, K) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Closure562 with func( 5 in)(6 out) closure first 2 argument
func Closure562[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(C, D, E) (F, G, H, I, J, K) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast562 with func( 5 in)(6 out) fix last 2 argument
func ClosureLast562[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A, B, C) (F, G, H, I, J, K) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Closure563 with func( 5 in)(6 out) closure first 3 argument
func Closure563[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(D, E) (F, G, H, I, J, K) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast563 with func( 5 in)(6 out) fix last 3 argument
func ClosureLast563[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A, B) (F, G, H, I, J, K) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Closure564 with func( 5 in)(6 out) closure first 4 argument
func Closure564[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(E) (F, G, H, I, J, K) {
	return func(e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast564 with func( 5 in)(6 out) fix last 4 argument
func ClosureLast564[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A) (F, G, H, I, J, K) {
	return func(a A) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Closure565 with func( 5 in)(6 out) closure first 5 argument
func Closure565[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func() (F, G, H, I, J, K) {
	return func() (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Closure571 with func( 5 in)(7 out) closure first 1 argument
func Closure571[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(B, C, D, E) (F, G, H, I, J, K, L) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast571 with func( 5 in)(7 out) fix last 1 argument
func ClosureLast571[A, B, C, D, E, F, G, H, I, J, K, L any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A, B, C, D) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Closure572 with func( 5 in)(7 out) closure first 2 argument
func Closure572[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(C, D, E) (F, G, H, I, J, K, L) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast572 with func( 5 in)(7 out) fix last 2 argument
func ClosureLast572[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A, B, C) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Closure573 with func( 5 in)(7 out) closure first 3 argument
func Closure573[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(D, E) (F, G, H, I, J, K, L) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast573 with func( 5 in)(7 out) fix last 3 argument
func ClosureLast573[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A, B) (F, G, H, I, J, K, L) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Closure574 with func( 5 in)(7 out) closure first 4 argument
func Closure574[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(E) (F, G, H, I, J, K, L) {
	return func(e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast574 with func( 5 in)(7 out) fix last 4 argument
func ClosureLast574[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A) (F, G, H, I, J, K, L) {
	return func(a A) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Closure575 with func( 5 in)(7 out) closure first 5 argument
func Closure575[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func() (F, G, H, I, J, K, L) {
	return func() (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Closure581 with func( 5 in)(8 out) closure first 1 argument
func Closure581[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(B, C, D, E) (F, G, H, I, J, K, L, M) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast581 with func( 5 in)(8 out) fix last 1 argument
func ClosureLast581[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A, B, C, D) (F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Closure582 with func( 5 in)(8 out) closure first 2 argument
func Closure582[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(C, D, E) (F, G, H, I, J, K, L, M) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast582 with func( 5 in)(8 out) fix last 2 argument
func ClosureLast582[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A, B, C) (F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Closure583 with func( 5 in)(8 out) closure first 3 argument
func Closure583[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(D, E) (F, G, H, I, J, K, L, M) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast583 with func( 5 in)(8 out) fix last 3 argument
func ClosureLast583[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A, B) (F, G, H, I, J, K, L, M) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Closure584 with func( 5 in)(8 out) closure first 4 argument
func Closure584[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(E) (F, G, H, I, J, K, L, M) {
	return func(e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast584 with func( 5 in)(8 out) fix last 4 argument
func ClosureLast584[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A) (F, G, H, I, J, K, L, M) {
	return func(a A) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Closure585 with func( 5 in)(8 out) closure first 5 argument
func Closure585[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func() (F, G, H, I, J, K, L, M) {
	return func() (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Closure591 with func( 5 in)(9 out) closure first 1 argument
func Closure591[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(B, C, D, E) (F, G, H, I, J, K, L, M, N) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast591 with func( 5 in)(9 out) fix last 1 argument
func ClosureLast591[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A, B, C, D) (F, G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Closure592 with func( 5 in)(9 out) closure first 2 argument
func Closure592[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(C, D, E) (F, G, H, I, J, K, L, M, N) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast592 with func( 5 in)(9 out) fix last 2 argument
func ClosureLast592[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A, B, C) (F, G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Closure593 with func( 5 in)(9 out) closure first 3 argument
func Closure593[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(D, E) (F, G, H, I, J, K, L, M, N) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast593 with func( 5 in)(9 out) fix last 3 argument
func ClosureLast593[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A, B) (F, G, H, I, J, K, L, M, N) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Closure594 with func( 5 in)(9 out) closure first 4 argument
func Closure594[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(E) (F, G, H, I, J, K, L, M, N) {
	return func(e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosureLast594 with func( 5 in)(9 out) fix last 4 argument
func ClosureLast594[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A) (F, G, H, I, J, K, L, M, N) {
	return func(a A) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Closure595 with func( 5 in)(9 out) closure first 5 argument
func Closure595[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func() (F, G, H, I, J, K, L, M, N) {
	return func() (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Closure601 with func( 6 in)(0 out) closure first 1 argument
func Closure601[A, B, C, D, E, F any](a A, fn func(A, B, C, D, E, F)) func(B, C, D, E, F) {
	return func(b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast601 with func( 6 in)(0 out) fix last 1 argument
func ClosureLast601[A, B, C, D, E, F any](f F, fn func(A, B, C, D, E, F)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Closure602 with func( 6 in)(0 out) closure first 2 argument
func Closure602[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D, E, F)) func(C, D, E, F) {
	return func(c C, d D, e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast602 with func( 6 in)(0 out) fix last 2 argument
func ClosureLast602[A, B, C, D, E, F any](e E, f F, fn func(A, B, C, D, E, F)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Closure603 with func( 6 in)(0 out) closure first 3 argument
func Closure603[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D, E, F)) func(D, E, F) {
	return func(d D, e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast603 with func( 6 in)(0 out) fix last 3 argument
func ClosureLast603[A, B, C, D, E, F any](d D, e E, f F, fn func(A, B, C, D, E, F)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Closure604 with func( 6 in)(0 out) closure first 4 argument
func Closure604[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D, E, F)) func(E, F) {
	return func(e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast604 with func( 6 in)(0 out) fix last 4 argument
func ClosureLast604[A, B, C, D, E, F any](c C, d D, e E, f F, fn func(A, B, C, D, E, F)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Closure605 with func( 6 in)(0 out) closure first 5 argument
func Closure605[A, B, C, D, E, F any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F)) func(F) {
	return func(f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast605 with func( 6 in)(0 out) fix last 5 argument
func ClosureLast605[A, B, C, D, E, F any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Closure606 with func( 6 in)(0 out) closure first 6 argument
func Closure606[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F)) func() {
	return func() {
		fn(a, b, c, d, e, f)
		return
	}
}

// Closure611 with func( 6 in)(1 out) closure first 1 argument
func Closure611[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E, F) G) func(B, C, D, E, F) G {
	return func(b B, c C, d D, e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast611 with func( 6 in)(1 out) fix last 1 argument
func ClosureLast611[A, B, C, D, E, F, G any](f F, fn func(A, B, C, D, E, F) G) func(A, B, C, D, E) G {
	return func(a A, b B, c C, d D, e E) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Closure612 with func( 6 in)(1 out) closure first 2 argument
func Closure612[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E, F) G) func(C, D, E, F) G {
	return func(c C, d D, e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast612 with func( 6 in)(1 out) fix last 2 argument
func ClosureLast612[A, B, C, D, E, F, G any](e E, f F, fn func(A, B, C, D, E, F) G) func(A, B, C, D) G {
	return func(a A, b B, c C, d D) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Closure613 with func( 6 in)(1 out) closure first 3 argument
func Closure613[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E, F) G) func(D, E, F) G {
	return func(d D, e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast613 with func( 6 in)(1 out) fix last 3 argument
func ClosureLast613[A, B, C, D, E, F, G any](d D, e E, f F, fn func(A, B, C, D, E, F) G) func(A, B, C) G {
	return func(a A, b B, c C) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Closure614 with func( 6 in)(1 out) closure first 4 argument
func Closure614[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) G) func(E, F) G {
	return func(e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast614 with func( 6 in)(1 out) fix last 4 argument
func ClosureLast614[A, B, C, D, E, F, G any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) G) func(A, B) G {
	return func(a A, b B) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Closure615 with func( 6 in)(1 out) closure first 5 argument
func Closure615[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) G) func(F) G {
	return func(f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast615 with func( 6 in)(1 out) fix last 5 argument
func ClosureLast615[A, B, C, D, E, F, G any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) G) func(A) G {
	return func(a A) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Closure616 with func( 6 in)(1 out) closure first 6 argument
func Closure616[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) G) func() G {
	return func() (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Closure621 with func( 6 in)(2 out) closure first 1 argument
func Closure621[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E, F) (G, H)) func(B, C, D, E, F) (G, H) {
	return func(b B, c C, d D, e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast621 with func( 6 in)(2 out) fix last 1 argument
func ClosureLast621[A, B, C, D, E, F, G, H any](f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B, C, D, E) (G, H) {
	return func(a A, b B, c C, d D, e E) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Closure622 with func( 6 in)(2 out) closure first 2 argument
func Closure622[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E, F) (G, H)) func(C, D, E, F) (G, H) {
	return func(c C, d D, e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast622 with func( 6 in)(2 out) fix last 2 argument
func ClosureLast622[A, B, C, D, E, F, G, H any](e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B, C, D) (G, H) {
	return func(a A, b B, c C, d D) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Closure623 with func( 6 in)(2 out) closure first 3 argument
func Closure623[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H)) func(D, E, F) (G, H) {
	return func(d D, e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast623 with func( 6 in)(2 out) fix last 3 argument
func ClosureLast623[A, B, C, D, E, F, G, H any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B, C) (G, H) {
	return func(a A, b B, c C) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Closure624 with func( 6 in)(2 out) closure first 4 argument
func Closure624[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H)) func(E, F) (G, H) {
	return func(e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast624 with func( 6 in)(2 out) fix last 4 argument
func ClosureLast624[A, B, C, D, E, F, G, H any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B) (G, H) {
	return func(a A, b B) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Closure625 with func( 6 in)(2 out) closure first 5 argument
func Closure625[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H)) func(F) (G, H) {
	return func(f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast625 with func( 6 in)(2 out) fix last 5 argument
func ClosureLast625[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A) (G, H) {
	return func(a A) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Closure626 with func( 6 in)(2 out) closure first 6 argument
func Closure626[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func() (G, H) {
	return func() (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Closure631 with func( 6 in)(3 out) closure first 1 argument
func Closure631[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F) (G, H, I)) func(B, C, D, E, F) (G, H, I) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast631 with func( 6 in)(3 out) fix last 1 argument
func ClosureLast631[A, B, C, D, E, F, G, H, I any](f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B, C, D, E) (G, H, I) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Closure632 with func( 6 in)(3 out) closure first 2 argument
func Closure632[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I)) func(C, D, E, F) (G, H, I) {
	return func(c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast632 with func( 6 in)(3 out) fix last 2 argument
func ClosureLast632[A, B, C, D, E, F, G, H, I any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B, C, D) (G, H, I) {
	return func(a A, b B, c C, d D) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Closure633 with func( 6 in)(3 out) closure first 3 argument
func Closure633[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I)) func(D, E, F) (G, H, I) {
	return func(d D, e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast633 with func( 6 in)(3 out) fix last 3 argument
func ClosureLast633[A, B, C, D, E, F, G, H, I any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B, C) (G, H, I) {
	return func(a A, b B, c C) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Closure634 with func( 6 in)(3 out) closure first 4 argument
func Closure634[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I)) func(E, F) (G, H, I) {
	return func(e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast634 with func( 6 in)(3 out) fix last 4 argument
func ClosureLast634[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B) (G, H, I) {
	return func(a A, b B) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Closure635 with func( 6 in)(3 out) closure first 5 argument
func Closure635[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I)) func(F) (G, H, I) {
	return func(f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast635 with func( 6 in)(3 out) fix last 5 argument
func ClosureLast635[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A) (G, H, I) {
	return func(a A) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Closure636 with func( 6 in)(3 out) closure first 6 argument
func Closure636[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func() (G, H, I) {
	return func() (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Closure641 with func( 6 in)(4 out) closure first 1 argument
func Closure641[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F) (G, H, I, J)) func(B, C, D, E, F) (G, H, I, J) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast641 with func( 6 in)(4 out) fix last 1 argument
func ClosureLast641[A, B, C, D, E, F, G, H, I, J any](f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B, C, D, E) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Closure642 with func( 6 in)(4 out) closure first 2 argument
func Closure642[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J)) func(C, D, E, F) (G, H, I, J) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast642 with func( 6 in)(4 out) fix last 2 argument
func ClosureLast642[A, B, C, D, E, F, G, H, I, J any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B, C, D) (G, H, I, J) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Closure643 with func( 6 in)(4 out) closure first 3 argument
func Closure643[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J)) func(D, E, F) (G, H, I, J) {
	return func(d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast643 with func( 6 in)(4 out) fix last 3 argument
func ClosureLast643[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B, C) (G, H, I, J) {
	return func(a A, b B, c C) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Closure644 with func( 6 in)(4 out) closure first 4 argument
func Closure644[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J)) func(E, F) (G, H, I, J) {
	return func(e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast644 with func( 6 in)(4 out) fix last 4 argument
func ClosureLast644[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B) (G, H, I, J) {
	return func(a A, b B) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Closure645 with func( 6 in)(4 out) closure first 5 argument
func Closure645[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J)) func(F) (G, H, I, J) {
	return func(f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast645 with func( 6 in)(4 out) fix last 5 argument
func ClosureLast645[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A) (G, H, I, J) {
	return func(a A) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Closure646 with func( 6 in)(4 out) closure first 6 argument
func Closure646[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func() (G, H, I, J) {
	return func() (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Closure651 with func( 6 in)(5 out) closure first 1 argument
func Closure651[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(B, C, D, E, F) (G, H, I, J, K) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast651 with func( 6 in)(5 out) fix last 1 argument
func ClosureLast651[A, B, C, D, E, F, G, H, I, J, K any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B, C, D, E) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Closure652 with func( 6 in)(5 out) closure first 2 argument
func Closure652[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(C, D, E, F) (G, H, I, J, K) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast652 with func( 6 in)(5 out) fix last 2 argument
func ClosureLast652[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B, C, D) (G, H, I, J, K) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Closure653 with func( 6 in)(5 out) closure first 3 argument
func Closure653[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(D, E, F) (G, H, I, J, K) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast653 with func( 6 in)(5 out) fix last 3 argument
func ClosureLast653[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B, C) (G, H, I, J, K) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Closure654 with func( 6 in)(5 out) closure first 4 argument
func Closure654[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(E, F) (G, H, I, J, K) {
	return func(e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast654 with func( 6 in)(5 out) fix last 4 argument
func ClosureLast654[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B) (G, H, I, J, K) {
	return func(a A, b B) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Closure655 with func( 6 in)(5 out) closure first 5 argument
func Closure655[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(F) (G, H, I, J, K) {
	return func(f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast655 with func( 6 in)(5 out) fix last 5 argument
func ClosureLast655[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A) (G, H, I, J, K) {
	return func(a A) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Closure656 with func( 6 in)(5 out) closure first 6 argument
func Closure656[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func() (G, H, I, J, K) {
	return func() (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Closure661 with func( 6 in)(6 out) closure first 1 argument
func Closure661[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(B, C, D, E, F) (G, H, I, J, K, L) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast661 with func( 6 in)(6 out) fix last 1 argument
func ClosureLast661[A, B, C, D, E, F, G, H, I, J, K, L any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B, C, D, E) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Closure662 with func( 6 in)(6 out) closure first 2 argument
func Closure662[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(C, D, E, F) (G, H, I, J, K, L) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast662 with func( 6 in)(6 out) fix last 2 argument
func ClosureLast662[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B, C, D) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Closure663 with func( 6 in)(6 out) closure first 3 argument
func Closure663[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(D, E, F) (G, H, I, J, K, L) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast663 with func( 6 in)(6 out) fix last 3 argument
func ClosureLast663[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B, C) (G, H, I, J, K, L) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Closure664 with func( 6 in)(6 out) closure first 4 argument
func Closure664[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(E, F) (G, H, I, J, K, L) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast664 with func( 6 in)(6 out) fix last 4 argument
func ClosureLast664[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B) (G, H, I, J, K, L) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Closure665 with func( 6 in)(6 out) closure first 5 argument
func Closure665[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(F) (G, H, I, J, K, L) {
	return func(f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast665 with func( 6 in)(6 out) fix last 5 argument
func ClosureLast665[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A) (G, H, I, J, K, L) {
	return func(a A) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Closure666 with func( 6 in)(6 out) closure first 6 argument
func Closure666[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func() (G, H, I, J, K, L) {
	return func() (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Closure671 with func( 6 in)(7 out) closure first 1 argument
func Closure671[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(B, C, D, E, F) (G, H, I, J, K, L, M) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast671 with func( 6 in)(7 out) fix last 1 argument
func ClosureLast671[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B, C, D, E) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Closure672 with func( 6 in)(7 out) closure first 2 argument
func Closure672[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(C, D, E, F) (G, H, I, J, K, L, M) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast672 with func( 6 in)(7 out) fix last 2 argument
func ClosureLast672[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B, C, D) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Closure673 with func( 6 in)(7 out) closure first 3 argument
func Closure673[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(D, E, F) (G, H, I, J, K, L, M) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast673 with func( 6 in)(7 out) fix last 3 argument
func ClosureLast673[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B, C) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Closure674 with func( 6 in)(7 out) closure first 4 argument
func Closure674[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(E, F) (G, H, I, J, K, L, M) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast674 with func( 6 in)(7 out) fix last 4 argument
func ClosureLast674[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B) (G, H, I, J, K, L, M) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Closure675 with func( 6 in)(7 out) closure first 5 argument
func Closure675[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(F) (G, H, I, J, K, L, M) {
	return func(f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast675 with func( 6 in)(7 out) fix last 5 argument
func ClosureLast675[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A) (G, H, I, J, K, L, M) {
	return func(a A) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Closure676 with func( 6 in)(7 out) closure first 6 argument
func Closure676[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func() (G, H, I, J, K, L, M) {
	return func() (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Closure681 with func( 6 in)(8 out) closure first 1 argument
func Closure681[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(B, C, D, E, F) (G, H, I, J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast681 with func( 6 in)(8 out) fix last 1 argument
func ClosureLast681[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B, C, D, E) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Closure682 with func( 6 in)(8 out) closure first 2 argument
func Closure682[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(C, D, E, F) (G, H, I, J, K, L, M, N) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast682 with func( 6 in)(8 out) fix last 2 argument
func ClosureLast682[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B, C, D) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Closure683 with func( 6 in)(8 out) closure first 3 argument
func Closure683[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(D, E, F) (G, H, I, J, K, L, M, N) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast683 with func( 6 in)(8 out) fix last 3 argument
func ClosureLast683[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B, C) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Closure684 with func( 6 in)(8 out) closure first 4 argument
func Closure684[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(E, F) (G, H, I, J, K, L, M, N) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast684 with func( 6 in)(8 out) fix last 4 argument
func ClosureLast684[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B) (G, H, I, J, K, L, M, N) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Closure685 with func( 6 in)(8 out) closure first 5 argument
func Closure685[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(F) (G, H, I, J, K, L, M, N) {
	return func(f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast685 with func( 6 in)(8 out) fix last 5 argument
func ClosureLast685[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A) (G, H, I, J, K, L, M, N) {
	return func(a A) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Closure686 with func( 6 in)(8 out) closure first 6 argument
func Closure686[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func() (G, H, I, J, K, L, M, N) {
	return func() (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Closure691 with func( 6 in)(9 out) closure first 1 argument
func Closure691[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(B, C, D, E, F) (G, H, I, J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast691 with func( 6 in)(9 out) fix last 1 argument
func ClosureLast691[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B, C, D, E) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Closure692 with func( 6 in)(9 out) closure first 2 argument
func Closure692[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(C, D, E, F) (G, H, I, J, K, L, M, N, O) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast692 with func( 6 in)(9 out) fix last 2 argument
func ClosureLast692[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B, C, D) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Closure693 with func( 6 in)(9 out) closure first 3 argument
func Closure693[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(D, E, F) (G, H, I, J, K, L, M, N, O) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast693 with func( 6 in)(9 out) fix last 3 argument
func ClosureLast693[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B, C) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Closure694 with func( 6 in)(9 out) closure first 4 argument
func Closure694[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(E, F) (G, H, I, J, K, L, M, N, O) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast694 with func( 6 in)(9 out) fix last 4 argument
func ClosureLast694[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Closure695 with func( 6 in)(9 out) closure first 5 argument
func Closure695[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(F) (G, H, I, J, K, L, M, N, O) {
	return func(f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosureLast695 with func( 6 in)(9 out) fix last 5 argument
func ClosureLast695[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A) (G, H, I, J, K, L, M, N, O) {
	return func(a A) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Closure696 with func( 6 in)(9 out) closure first 6 argument
func Closure696[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func() (G, H, I, J, K, L, M, N, O) {
	return func() (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Closure701 with func( 7 in)(0 out) closure first 1 argument
func Closure701[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E, F, G)) func(B, C, D, E, F, G) {
	return func(b B, c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast701 with func( 7 in)(0 out) fix last 1 argument
func ClosureLast701[A, B, C, D, E, F, G any](g G, fn func(A, B, C, D, E, F, G)) func(A, B, C, D, E, F) {
	return func(a A, b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure702 with func( 7 in)(0 out) closure first 2 argument
func Closure702[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E, F, G)) func(C, D, E, F, G) {
	return func(c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast702 with func( 7 in)(0 out) fix last 2 argument
func ClosureLast702[A, B, C, D, E, F, G any](f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure703 with func( 7 in)(0 out) closure first 3 argument
func Closure703[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E, F, G)) func(D, E, F, G) {
	return func(d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast703 with func( 7 in)(0 out) fix last 3 argument
func ClosureLast703[A, B, C, D, E, F, G any](e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure704 with func( 7 in)(0 out) closure first 4 argument
func Closure704[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G)) func(E, F, G) {
	return func(e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast704 with func( 7 in)(0 out) fix last 4 argument
func ClosureLast704[A, B, C, D, E, F, G any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure705 with func( 7 in)(0 out) closure first 5 argument
func Closure705[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G)) func(F, G) {
	return func(f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast705 with func( 7 in)(0 out) fix last 5 argument
func ClosureLast705[A, B, C, D, E, F, G any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure706 with func( 7 in)(0 out) closure first 6 argument
func Closure706[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G)) func(G) {
	return func(g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast706 with func( 7 in)(0 out) fix last 6 argument
func ClosureLast706[A, B, C, D, E, F, G any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure707 with func( 7 in)(0 out) closure first 7 argument
func Closure707[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func() {
	return func() {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure711 with func( 7 in)(1 out) closure first 1 argument
func Closure711[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E, F, G) H) func(B, C, D, E, F, G) H {
	return func(b B, c C, d D, e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast711 with func( 7 in)(1 out) fix last 1 argument
func ClosureLast711[A, B, C, D, E, F, G, H any](g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C, D, E, F) H {
	return func(a A, b B, c C, d D, e E, f F) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure712 with func( 7 in)(1 out) closure first 2 argument
func Closure712[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E, F, G) H) func(C, D, E, F, G) H {
	return func(c C, d D, e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast712 with func( 7 in)(1 out) fix last 2 argument
func ClosureLast712[A, B, C, D, E, F, G, H any](f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C, D, E) H {
	return func(a A, b B, c C, d D, e E) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure713 with func( 7 in)(1 out) closure first 3 argument
func Closure713[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E, F, G) H) func(D, E, F, G) H {
	return func(d D, e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast713 with func( 7 in)(1 out) fix last 3 argument
func ClosureLast713[A, B, C, D, E, F, G, H any](e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C, D) H {
	return func(a A, b B, c C, d D) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure714 with func( 7 in)(1 out) closure first 4 argument
func Closure714[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) H) func(E, F, G) H {
	return func(e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast714 with func( 7 in)(1 out) fix last 4 argument
func ClosureLast714[A, B, C, D, E, F, G, H any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C) H {
	return func(a A, b B, c C) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure715 with func( 7 in)(1 out) closure first 5 argument
func Closure715[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) H) func(F, G) H {
	return func(f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast715 with func( 7 in)(1 out) fix last 5 argument
func ClosureLast715[A, B, C, D, E, F, G, H any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B) H {
	return func(a A, b B) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure716 with func( 7 in)(1 out) closure first 6 argument
func Closure716[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) H) func(G) H {
	return func(g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast716 with func( 7 in)(1 out) fix last 6 argument
func ClosureLast716[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A) H {
	return func(a A) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure717 with func( 7 in)(1 out) closure first 7 argument
func Closure717[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func() H {
	return func() (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure721 with func( 7 in)(2 out) closure first 1 argument
func Closure721[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F, G) (H, I)) func(B, C, D, E, F, G) (H, I) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast721 with func( 7 in)(2 out) fix last 1 argument
func ClosureLast721[A, B, C, D, E, F, G, H, I any](g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C, D, E, F) (H, I) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure722 with func( 7 in)(2 out) closure first 2 argument
func Closure722[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I)) func(C, D, E, F, G) (H, I) {
	return func(c C, d D, e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast722 with func( 7 in)(2 out) fix last 2 argument
func ClosureLast722[A, B, C, D, E, F, G, H, I any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C, D, E) (H, I) {
	return func(a A, b B, c C, d D, e E) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure723 with func( 7 in)(2 out) closure first 3 argument
func Closure723[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I)) func(D, E, F, G) (H, I) {
	return func(d D, e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast723 with func( 7 in)(2 out) fix last 3 argument
func ClosureLast723[A, B, C, D, E, F, G, H, I any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C, D) (H, I) {
	return func(a A, b B, c C, d D) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure724 with func( 7 in)(2 out) closure first 4 argument
func Closure724[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I)) func(E, F, G) (H, I) {
	return func(e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast724 with func( 7 in)(2 out) fix last 4 argument
func ClosureLast724[A, B, C, D, E, F, G, H, I any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C) (H, I) {
	return func(a A, b B, c C) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure725 with func( 7 in)(2 out) closure first 5 argument
func Closure725[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I)) func(F, G) (H, I) {
	return func(f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast725 with func( 7 in)(2 out) fix last 5 argument
func ClosureLast725[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B) (H, I) {
	return func(a A, b B) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure726 with func( 7 in)(2 out) closure first 6 argument
func Closure726[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I)) func(G) (H, I) {
	return func(g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast726 with func( 7 in)(2 out) fix last 6 argument
func ClosureLast726[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A) (H, I) {
	return func(a A) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure727 with func( 7 in)(2 out) closure first 7 argument
func Closure727[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func() (H, I) {
	return func() (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure731 with func( 7 in)(3 out) closure first 1 argument
func Closure731[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F, G) (H, I, J)) func(B, C, D, E, F, G) (H, I, J) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast731 with func( 7 in)(3 out) fix last 1 argument
func ClosureLast731[A, B, C, D, E, F, G, H, I, J any](g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C, D, E, F) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure732 with func( 7 in)(3 out) closure first 2 argument
func Closure732[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J)) func(C, D, E, F, G) (H, I, J) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast732 with func( 7 in)(3 out) fix last 2 argument
func ClosureLast732[A, B, C, D, E, F, G, H, I, J any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C, D, E) (H, I, J) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure733 with func( 7 in)(3 out) closure first 3 argument
func Closure733[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J)) func(D, E, F, G) (H, I, J) {
	return func(d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast733 with func( 7 in)(3 out) fix last 3 argument
func ClosureLast733[A, B, C, D, E, F, G, H, I, J any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C, D) (H, I, J) {
	return func(a A, b B, c C, d D) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure734 with func( 7 in)(3 out) closure first 4 argument
func Closure734[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J)) func(E, F, G) (H, I, J) {
	return func(e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast734 with func( 7 in)(3 out) fix last 4 argument
func ClosureLast734[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C) (H, I, J) {
	return func(a A, b B, c C) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure735 with func( 7 in)(3 out) closure first 5 argument
func Closure735[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J)) func(F, G) (H, I, J) {
	return func(f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast735 with func( 7 in)(3 out) fix last 5 argument
func ClosureLast735[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B) (H, I, J) {
	return func(a A, b B) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure736 with func( 7 in)(3 out) closure first 6 argument
func Closure736[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J)) func(G) (H, I, J) {
	return func(g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast736 with func( 7 in)(3 out) fix last 6 argument
func ClosureLast736[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A) (H, I, J) {
	return func(a A) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure737 with func( 7 in)(3 out) closure first 7 argument
func Closure737[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func() (H, I, J) {
	return func() (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure741 with func( 7 in)(4 out) closure first 1 argument
func Closure741[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(B, C, D, E, F, G) (H, I, J, K) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast741 with func( 7 in)(4 out) fix last 1 argument
func ClosureLast741[A, B, C, D, E, F, G, H, I, J, K any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C, D, E, F) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure742 with func( 7 in)(4 out) closure first 2 argument
func Closure742[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(C, D, E, F, G) (H, I, J, K) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast742 with func( 7 in)(4 out) fix last 2 argument
func ClosureLast742[A, B, C, D, E, F, G, H, I, J, K any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C, D, E) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure743 with func( 7 in)(4 out) closure first 3 argument
func Closure743[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(D, E, F, G) (H, I, J, K) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast743 with func( 7 in)(4 out) fix last 3 argument
func ClosureLast743[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C, D) (H, I, J, K) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure744 with func( 7 in)(4 out) closure first 4 argument
func Closure744[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(E, F, G) (H, I, J, K) {
	return func(e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast744 with func( 7 in)(4 out) fix last 4 argument
func ClosureLast744[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C) (H, I, J, K) {
	return func(a A, b B, c C) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure745 with func( 7 in)(4 out) closure first 5 argument
func Closure745[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(F, G) (H, I, J, K) {
	return func(f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast745 with func( 7 in)(4 out) fix last 5 argument
func ClosureLast745[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B) (H, I, J, K) {
	return func(a A, b B) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure746 with func( 7 in)(4 out) closure first 6 argument
func Closure746[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(G) (H, I, J, K) {
	return func(g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast746 with func( 7 in)(4 out) fix last 6 argument
func ClosureLast746[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A) (H, I, J, K) {
	return func(a A) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure747 with func( 7 in)(4 out) closure first 7 argument
func Closure747[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func() (H, I, J, K) {
	return func() (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure751 with func( 7 in)(5 out) closure first 1 argument
func Closure751[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(B, C, D, E, F, G) (H, I, J, K, L) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast751 with func( 7 in)(5 out) fix last 1 argument
func ClosureLast751[A, B, C, D, E, F, G, H, I, J, K, L any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C, D, E, F) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure752 with func( 7 in)(5 out) closure first 2 argument
func Closure752[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(C, D, E, F, G) (H, I, J, K, L) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast752 with func( 7 in)(5 out) fix last 2 argument
func ClosureLast752[A, B, C, D, E, F, G, H, I, J, K, L any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C, D, E) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure753 with func( 7 in)(5 out) closure first 3 argument
func Closure753[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(D, E, F, G) (H, I, J, K, L) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast753 with func( 7 in)(5 out) fix last 3 argument
func ClosureLast753[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C, D) (H, I, J, K, L) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure754 with func( 7 in)(5 out) closure first 4 argument
func Closure754[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(E, F, G) (H, I, J, K, L) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast754 with func( 7 in)(5 out) fix last 4 argument
func ClosureLast754[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C) (H, I, J, K, L) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure755 with func( 7 in)(5 out) closure first 5 argument
func Closure755[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(F, G) (H, I, J, K, L) {
	return func(f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast755 with func( 7 in)(5 out) fix last 5 argument
func ClosureLast755[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B) (H, I, J, K, L) {
	return func(a A, b B) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure756 with func( 7 in)(5 out) closure first 6 argument
func Closure756[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(G) (H, I, J, K, L) {
	return func(g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast756 with func( 7 in)(5 out) fix last 6 argument
func ClosureLast756[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A) (H, I, J, K, L) {
	return func(a A) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure757 with func( 7 in)(5 out) closure first 7 argument
func Closure757[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func() (H, I, J, K, L) {
	return func() (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure761 with func( 7 in)(6 out) closure first 1 argument
func Closure761[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(B, C, D, E, F, G) (H, I, J, K, L, M) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast761 with func( 7 in)(6 out) fix last 1 argument
func ClosureLast761[A, B, C, D, E, F, G, H, I, J, K, L, M any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C, D, E, F) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure762 with func( 7 in)(6 out) closure first 2 argument
func Closure762[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(C, D, E, F, G) (H, I, J, K, L, M) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast762 with func( 7 in)(6 out) fix last 2 argument
func ClosureLast762[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C, D, E) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure763 with func( 7 in)(6 out) closure first 3 argument
func Closure763[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(D, E, F, G) (H, I, J, K, L, M) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast763 with func( 7 in)(6 out) fix last 3 argument
func ClosureLast763[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C, D) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure764 with func( 7 in)(6 out) closure first 4 argument
func Closure764[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(E, F, G) (H, I, J, K, L, M) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast764 with func( 7 in)(6 out) fix last 4 argument
func ClosureLast764[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C) (H, I, J, K, L, M) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure765 with func( 7 in)(6 out) closure first 5 argument
func Closure765[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(F, G) (H, I, J, K, L, M) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast765 with func( 7 in)(6 out) fix last 5 argument
func ClosureLast765[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B) (H, I, J, K, L, M) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure766 with func( 7 in)(6 out) closure first 6 argument
func Closure766[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(G) (H, I, J, K, L, M) {
	return func(g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast766 with func( 7 in)(6 out) fix last 6 argument
func ClosureLast766[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A) (H, I, J, K, L, M) {
	return func(a A) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure767 with func( 7 in)(6 out) closure first 7 argument
func Closure767[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func() (H, I, J, K, L, M) {
	return func() (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure771 with func( 7 in)(7 out) closure first 1 argument
func Closure771[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(B, C, D, E, F, G) (H, I, J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast771 with func( 7 in)(7 out) fix last 1 argument
func ClosureLast771[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C, D, E, F) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure772 with func( 7 in)(7 out) closure first 2 argument
func Closure772[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(C, D, E, F, G) (H, I, J, K, L, M, N) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast772 with func( 7 in)(7 out) fix last 2 argument
func ClosureLast772[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C, D, E) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure773 with func( 7 in)(7 out) closure first 3 argument
func Closure773[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(D, E, F, G) (H, I, J, K, L, M, N) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast773 with func( 7 in)(7 out) fix last 3 argument
func ClosureLast773[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C, D) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure774 with func( 7 in)(7 out) closure first 4 argument
func Closure774[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(E, F, G) (H, I, J, K, L, M, N) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast774 with func( 7 in)(7 out) fix last 4 argument
func ClosureLast774[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure775 with func( 7 in)(7 out) closure first 5 argument
func Closure775[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(F, G) (H, I, J, K, L, M, N) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast775 with func( 7 in)(7 out) fix last 5 argument
func ClosureLast775[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B) (H, I, J, K, L, M, N) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure776 with func( 7 in)(7 out) closure first 6 argument
func Closure776[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(G) (H, I, J, K, L, M, N) {
	return func(g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast776 with func( 7 in)(7 out) fix last 6 argument
func ClosureLast776[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A) (H, I, J, K, L, M, N) {
	return func(a A) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure777 with func( 7 in)(7 out) closure first 7 argument
func Closure777[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func() (H, I, J, K, L, M, N) {
	return func() (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure781 with func( 7 in)(8 out) closure first 1 argument
func Closure781[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(B, C, D, E, F, G) (H, I, J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast781 with func( 7 in)(8 out) fix last 1 argument
func ClosureLast781[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C, D, E, F) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure782 with func( 7 in)(8 out) closure first 2 argument
func Closure782[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(C, D, E, F, G) (H, I, J, K, L, M, N, O) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast782 with func( 7 in)(8 out) fix last 2 argument
func ClosureLast782[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C, D, E) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure783 with func( 7 in)(8 out) closure first 3 argument
func Closure783[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(D, E, F, G) (H, I, J, K, L, M, N, O) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast783 with func( 7 in)(8 out) fix last 3 argument
func ClosureLast783[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C, D) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure784 with func( 7 in)(8 out) closure first 4 argument
func Closure784[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(E, F, G) (H, I, J, K, L, M, N, O) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast784 with func( 7 in)(8 out) fix last 4 argument
func ClosureLast784[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure785 with func( 7 in)(8 out) closure first 5 argument
func Closure785[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(F, G) (H, I, J, K, L, M, N, O) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast785 with func( 7 in)(8 out) fix last 5 argument
func ClosureLast785[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B) (H, I, J, K, L, M, N, O) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure786 with func( 7 in)(8 out) closure first 6 argument
func Closure786[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(G) (H, I, J, K, L, M, N, O) {
	return func(g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast786 with func( 7 in)(8 out) fix last 6 argument
func ClosureLast786[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A) (H, I, J, K, L, M, N, O) {
	return func(a A) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure787 with func( 7 in)(8 out) closure first 7 argument
func Closure787[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func() (H, I, J, K, L, M, N, O) {
	return func() (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure791 with func( 7 in)(9 out) closure first 1 argument
func Closure791[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast791 with func( 7 in)(9 out) fix last 1 argument
func ClosureLast791[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C, D, E, F) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure792 with func( 7 in)(9 out) closure first 2 argument
func Closure792[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(C, D, E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast792 with func( 7 in)(9 out) fix last 2 argument
func ClosureLast792[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C, D, E) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure793 with func( 7 in)(9 out) closure first 3 argument
func Closure793[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(D, E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast793 with func( 7 in)(9 out) fix last 3 argument
func ClosureLast793[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C, D) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure794 with func( 7 in)(9 out) closure first 4 argument
func Closure794[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast794 with func( 7 in)(9 out) fix last 4 argument
func ClosureLast794[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure795 with func( 7 in)(9 out) closure first 5 argument
func Closure795[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(F, G) (H, I, J, K, L, M, N, O, P) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast795 with func( 7 in)(9 out) fix last 5 argument
func ClosureLast795[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure796 with func( 7 in)(9 out) closure first 6 argument
func Closure796[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(G) (H, I, J, K, L, M, N, O, P) {
	return func(g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosureLast796 with func( 7 in)(9 out) fix last 6 argument
func ClosureLast796[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A) (H, I, J, K, L, M, N, O, P) {
	return func(a A) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure797 with func( 7 in)(9 out) closure first 7 argument
func Closure797[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func() (H, I, J, K, L, M, N, O, P) {
	return func() (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Closure801 with func( 8 in)(0 out) closure first 1 argument
func Closure801[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E, F, G, H)) func(B, C, D, E, F, G, H) {
	return func(b B, c C, d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast801 with func( 8 in)(0 out) fix last 1 argument
func ClosureLast801[A, B, C, D, E, F, G, H any](h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D, E, F, G) {
	return func(a A, b B, c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure802 with func( 8 in)(0 out) closure first 2 argument
func Closure802[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E, F, G, H)) func(C, D, E, F, G, H) {
	return func(c C, d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast802 with func( 8 in)(0 out) fix last 2 argument
func ClosureLast802[A, B, C, D, E, F, G, H any](g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D, E, F) {
	return func(a A, b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure803 with func( 8 in)(0 out) closure first 3 argument
func Closure803[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H)) func(D, E, F, G, H) {
	return func(d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast803 with func( 8 in)(0 out) fix last 3 argument
func ClosureLast803[A, B, C, D, E, F, G, H any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure804 with func( 8 in)(0 out) closure first 4 argument
func Closure804[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H)) func(E, F, G, H) {
	return func(e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast804 with func( 8 in)(0 out) fix last 4 argument
func ClosureLast804[A, B, C, D, E, F, G, H any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure805 with func( 8 in)(0 out) closure first 5 argument
func Closure805[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H)) func(F, G, H) {
	return func(f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast805 with func( 8 in)(0 out) fix last 5 argument
func ClosureLast805[A, B, C, D, E, F, G, H any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure806 with func( 8 in)(0 out) closure first 6 argument
func Closure806[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H)) func(G, H) {
	return func(g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast806 with func( 8 in)(0 out) fix last 6 argument
func ClosureLast806[A, B, C, D, E, F, G, H any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure807 with func( 8 in)(0 out) closure first 7 argument
func Closure807[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H)) func(H) {
	return func(h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast807 with func( 8 in)(0 out) fix last 7 argument
func ClosureLast807[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure808 with func( 8 in)(0 out) closure first 8 argument
func Closure808[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func() {
	return func() {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure811 with func( 8 in)(1 out) closure first 1 argument
func Closure811[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F, G, H) I) func(B, C, D, E, F, G, H) I {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast811 with func( 8 in)(1 out) fix last 1 argument
func ClosureLast811[A, B, C, D, E, F, G, H, I any](h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D, E, F, G) I {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure812 with func( 8 in)(1 out) closure first 2 argument
func Closure812[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F, G, H) I) func(C, D, E, F, G, H) I {
	return func(c C, d D, e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast812 with func( 8 in)(1 out) fix last 2 argument
func ClosureLast812[A, B, C, D, E, F, G, H, I any](g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D, E, F) I {
	return func(a A, b B, c C, d D, e E, f F) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure813 with func( 8 in)(1 out) closure first 3 argument
func Closure813[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) I) func(D, E, F, G, H) I {
	return func(d D, e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast813 with func( 8 in)(1 out) fix last 3 argument
func ClosureLast813[A, B, C, D, E, F, G, H, I any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D, E) I {
	return func(a A, b B, c C, d D, e E) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure814 with func( 8 in)(1 out) closure first 4 argument
func Closure814[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) I) func(E, F, G, H) I {
	return func(e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast814 with func( 8 in)(1 out) fix last 4 argument
func ClosureLast814[A, B, C, D, E, F, G, H, I any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D) I {
	return func(a A, b B, c C, d D) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure815 with func( 8 in)(1 out) closure first 5 argument
func Closure815[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) I) func(F, G, H) I {
	return func(f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast815 with func( 8 in)(1 out) fix last 5 argument
func ClosureLast815[A, B, C, D, E, F, G, H, I any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C) I {
	return func(a A, b B, c C) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure816 with func( 8 in)(1 out) closure first 6 argument
func Closure816[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) I) func(G, H) I {
	return func(g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast816 with func( 8 in)(1 out) fix last 6 argument
func ClosureLast816[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B) I {
	return func(a A, b B) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure817 with func( 8 in)(1 out) closure first 7 argument
func Closure817[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) I) func(H) I {
	return func(h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast817 with func( 8 in)(1 out) fix last 7 argument
func ClosureLast817[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A) I {
	return func(a A) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure818 with func( 8 in)(1 out) closure first 8 argument
func Closure818[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func() I {
	return func() (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure821 with func( 8 in)(2 out) closure first 1 argument
func Closure821[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F, G, H) (I, J)) func(B, C, D, E, F, G, H) (I, J) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast821 with func( 8 in)(2 out) fix last 1 argument
func ClosureLast821[A, B, C, D, E, F, G, H, I, J any](h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D, E, F, G) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure822 with func( 8 in)(2 out) closure first 2 argument
func Closure822[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J)) func(C, D, E, F, G, H) (I, J) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast822 with func( 8 in)(2 out) fix last 2 argument
func ClosureLast822[A, B, C, D, E, F, G, H, I, J any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D, E, F) (I, J) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure823 with func( 8 in)(2 out) closure first 3 argument
func Closure823[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J)) func(D, E, F, G, H) (I, J) {
	return func(d D, e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast823 with func( 8 in)(2 out) fix last 3 argument
func ClosureLast823[A, B, C, D, E, F, G, H, I, J any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D, E) (I, J) {
	return func(a A, b B, c C, d D, e E) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure824 with func( 8 in)(2 out) closure first 4 argument
func Closure824[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J)) func(E, F, G, H) (I, J) {
	return func(e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast824 with func( 8 in)(2 out) fix last 4 argument
func ClosureLast824[A, B, C, D, E, F, G, H, I, J any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D) (I, J) {
	return func(a A, b B, c C, d D) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure825 with func( 8 in)(2 out) closure first 5 argument
func Closure825[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J)) func(F, G, H) (I, J) {
	return func(f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast825 with func( 8 in)(2 out) fix last 5 argument
func ClosureLast825[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C) (I, J) {
	return func(a A, b B, c C) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure826 with func( 8 in)(2 out) closure first 6 argument
func Closure826[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J)) func(G, H) (I, J) {
	return func(g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast826 with func( 8 in)(2 out) fix last 6 argument
func ClosureLast826[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B) (I, J) {
	return func(a A, b B) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure827 with func( 8 in)(2 out) closure first 7 argument
func Closure827[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J)) func(H) (I, J) {
	return func(h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast827 with func( 8 in)(2 out) fix last 7 argument
func ClosureLast827[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A) (I, J) {
	return func(a A) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure828 with func( 8 in)(2 out) closure first 8 argument
func Closure828[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func() (I, J) {
	return func() (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure831 with func( 8 in)(3 out) closure first 1 argument
func Closure831[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(B, C, D, E, F, G, H) (I, J, K) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast831 with func( 8 in)(3 out) fix last 1 argument
func ClosureLast831[A, B, C, D, E, F, G, H, I, J, K any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D, E, F, G) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure832 with func( 8 in)(3 out) closure first 2 argument
func Closure832[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(C, D, E, F, G, H) (I, J, K) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast832 with func( 8 in)(3 out) fix last 2 argument
func ClosureLast832[A, B, C, D, E, F, G, H, I, J, K any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D, E, F) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure833 with func( 8 in)(3 out) closure first 3 argument
func Closure833[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(D, E, F, G, H) (I, J, K) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast833 with func( 8 in)(3 out) fix last 3 argument
func ClosureLast833[A, B, C, D, E, F, G, H, I, J, K any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D, E) (I, J, K) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure834 with func( 8 in)(3 out) closure first 4 argument
func Closure834[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(E, F, G, H) (I, J, K) {
	return func(e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast834 with func( 8 in)(3 out) fix last 4 argument
func ClosureLast834[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D) (I, J, K) {
	return func(a A, b B, c C, d D) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure835 with func( 8 in)(3 out) closure first 5 argument
func Closure835[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(F, G, H) (I, J, K) {
	return func(f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast835 with func( 8 in)(3 out) fix last 5 argument
func ClosureLast835[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C) (I, J, K) {
	return func(a A, b B, c C) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure836 with func( 8 in)(3 out) closure first 6 argument
func Closure836[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(G, H) (I, J, K) {
	return func(g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast836 with func( 8 in)(3 out) fix last 6 argument
func ClosureLast836[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B) (I, J, K) {
	return func(a A, b B) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure837 with func( 8 in)(3 out) closure first 7 argument
func Closure837[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(H) (I, J, K) {
	return func(h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast837 with func( 8 in)(3 out) fix last 7 argument
func ClosureLast837[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A) (I, J, K) {
	return func(a A) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure838 with func( 8 in)(3 out) closure first 8 argument
func Closure838[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func() (I, J, K) {
	return func() (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure841 with func( 8 in)(4 out) closure first 1 argument
func Closure841[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(B, C, D, E, F, G, H) (I, J, K, L) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast841 with func( 8 in)(4 out) fix last 1 argument
func ClosureLast841[A, B, C, D, E, F, G, H, I, J, K, L any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D, E, F, G) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure842 with func( 8 in)(4 out) closure first 2 argument
func Closure842[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(C, D, E, F, G, H) (I, J, K, L) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast842 with func( 8 in)(4 out) fix last 2 argument
func ClosureLast842[A, B, C, D, E, F, G, H, I, J, K, L any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D, E, F) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure843 with func( 8 in)(4 out) closure first 3 argument
func Closure843[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(D, E, F, G, H) (I, J, K, L) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast843 with func( 8 in)(4 out) fix last 3 argument
func ClosureLast843[A, B, C, D, E, F, G, H, I, J, K, L any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D, E) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure844 with func( 8 in)(4 out) closure first 4 argument
func Closure844[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(E, F, G, H) (I, J, K, L) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast844 with func( 8 in)(4 out) fix last 4 argument
func ClosureLast844[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D) (I, J, K, L) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure845 with func( 8 in)(4 out) closure first 5 argument
func Closure845[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(F, G, H) (I, J, K, L) {
	return func(f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast845 with func( 8 in)(4 out) fix last 5 argument
func ClosureLast845[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C) (I, J, K, L) {
	return func(a A, b B, c C) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure846 with func( 8 in)(4 out) closure first 6 argument
func Closure846[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(G, H) (I, J, K, L) {
	return func(g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast846 with func( 8 in)(4 out) fix last 6 argument
func ClosureLast846[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B) (I, J, K, L) {
	return func(a A, b B) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure847 with func( 8 in)(4 out) closure first 7 argument
func Closure847[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(H) (I, J, K, L) {
	return func(h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast847 with func( 8 in)(4 out) fix last 7 argument
func ClosureLast847[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A) (I, J, K, L) {
	return func(a A) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure848 with func( 8 in)(4 out) closure first 8 argument
func Closure848[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func() (I, J, K, L) {
	return func() (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure851 with func( 8 in)(5 out) closure first 1 argument
func Closure851[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(B, C, D, E, F, G, H) (I, J, K, L, M) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast851 with func( 8 in)(5 out) fix last 1 argument
func ClosureLast851[A, B, C, D, E, F, G, H, I, J, K, L, M any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D, E, F, G) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure852 with func( 8 in)(5 out) closure first 2 argument
func Closure852[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(C, D, E, F, G, H) (I, J, K, L, M) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast852 with func( 8 in)(5 out) fix last 2 argument
func ClosureLast852[A, B, C, D, E, F, G, H, I, J, K, L, M any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D, E, F) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure853 with func( 8 in)(5 out) closure first 3 argument
func Closure853[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(D, E, F, G, H) (I, J, K, L, M) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast853 with func( 8 in)(5 out) fix last 3 argument
func ClosureLast853[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D, E) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure854 with func( 8 in)(5 out) closure first 4 argument
func Closure854[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(E, F, G, H) (I, J, K, L, M) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast854 with func( 8 in)(5 out) fix last 4 argument
func ClosureLast854[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D) (I, J, K, L, M) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure855 with func( 8 in)(5 out) closure first 5 argument
func Closure855[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(F, G, H) (I, J, K, L, M) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast855 with func( 8 in)(5 out) fix last 5 argument
func ClosureLast855[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C) (I, J, K, L, M) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure856 with func( 8 in)(5 out) closure first 6 argument
func Closure856[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(G, H) (I, J, K, L, M) {
	return func(g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast856 with func( 8 in)(5 out) fix last 6 argument
func ClosureLast856[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B) (I, J, K, L, M) {
	return func(a A, b B) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure857 with func( 8 in)(5 out) closure first 7 argument
func Closure857[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(H) (I, J, K, L, M) {
	return func(h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast857 with func( 8 in)(5 out) fix last 7 argument
func ClosureLast857[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A) (I, J, K, L, M) {
	return func(a A) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure858 with func( 8 in)(5 out) closure first 8 argument
func Closure858[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func() (I, J, K, L, M) {
	return func() (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure861 with func( 8 in)(6 out) closure first 1 argument
func Closure861[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast861 with func( 8 in)(6 out) fix last 1 argument
func ClosureLast861[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure862 with func( 8 in)(6 out) closure first 2 argument
func Closure862[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(C, D, E, F, G, H) (I, J, K, L, M, N) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast862 with func( 8 in)(6 out) fix last 2 argument
func ClosureLast862[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D, E, F) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure863 with func( 8 in)(6 out) closure first 3 argument
func Closure863[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(D, E, F, G, H) (I, J, K, L, M, N) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast863 with func( 8 in)(6 out) fix last 3 argument
func ClosureLast863[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D, E) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure864 with func( 8 in)(6 out) closure first 4 argument
func Closure864[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(E, F, G, H) (I, J, K, L, M, N) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast864 with func( 8 in)(6 out) fix last 4 argument
func ClosureLast864[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure865 with func( 8 in)(6 out) closure first 5 argument
func Closure865[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(F, G, H) (I, J, K, L, M, N) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast865 with func( 8 in)(6 out) fix last 5 argument
func ClosureLast865[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C) (I, J, K, L, M, N) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure866 with func( 8 in)(6 out) closure first 6 argument
func Closure866[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(G, H) (I, J, K, L, M, N) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast866 with func( 8 in)(6 out) fix last 6 argument
func ClosureLast866[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B) (I, J, K, L, M, N) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure867 with func( 8 in)(6 out) closure first 7 argument
func Closure867[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(H) (I, J, K, L, M, N) {
	return func(h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast867 with func( 8 in)(6 out) fix last 7 argument
func ClosureLast867[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A) (I, J, K, L, M, N) {
	return func(a A) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure868 with func( 8 in)(6 out) closure first 8 argument
func Closure868[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func() (I, J, K, L, M, N) {
	return func() (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure871 with func( 8 in)(7 out) closure first 1 argument
func Closure871[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast871 with func( 8 in)(7 out) fix last 1 argument
func ClosureLast871[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure872 with func( 8 in)(7 out) closure first 2 argument
func Closure872[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(C, D, E, F, G, H) (I, J, K, L, M, N, O) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast872 with func( 8 in)(7 out) fix last 2 argument
func ClosureLast872[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D, E, F) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure873 with func( 8 in)(7 out) closure first 3 argument
func Closure873[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(D, E, F, G, H) (I, J, K, L, M, N, O) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast873 with func( 8 in)(7 out) fix last 3 argument
func ClosureLast873[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D, E) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure874 with func( 8 in)(7 out) closure first 4 argument
func Closure874[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(E, F, G, H) (I, J, K, L, M, N, O) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast874 with func( 8 in)(7 out) fix last 4 argument
func ClosureLast874[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure875 with func( 8 in)(7 out) closure first 5 argument
func Closure875[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(F, G, H) (I, J, K, L, M, N, O) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast875 with func( 8 in)(7 out) fix last 5 argument
func ClosureLast875[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure876 with func( 8 in)(7 out) closure first 6 argument
func Closure876[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(G, H) (I, J, K, L, M, N, O) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast876 with func( 8 in)(7 out) fix last 6 argument
func ClosureLast876[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B) (I, J, K, L, M, N, O) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure877 with func( 8 in)(7 out) closure first 7 argument
func Closure877[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(H) (I, J, K, L, M, N, O) {
	return func(h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast877 with func( 8 in)(7 out) fix last 7 argument
func ClosureLast877[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A) (I, J, K, L, M, N, O) {
	return func(a A) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure878 with func( 8 in)(7 out) closure first 8 argument
func Closure878[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func() (I, J, K, L, M, N, O) {
	return func() (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure881 with func( 8 in)(8 out) closure first 1 argument
func Closure881[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast881 with func( 8 in)(8 out) fix last 1 argument
func ClosureLast881[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure882 with func( 8 in)(8 out) closure first 2 argument
func Closure882[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(C, D, E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast882 with func( 8 in)(8 out) fix last 2 argument
func ClosureLast882[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D, E, F) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure883 with func( 8 in)(8 out) closure first 3 argument
func Closure883[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(D, E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast883 with func( 8 in)(8 out) fix last 3 argument
func ClosureLast883[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D, E) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure884 with func( 8 in)(8 out) closure first 4 argument
func Closure884[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast884 with func( 8 in)(8 out) fix last 4 argument
func ClosureLast884[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure885 with func( 8 in)(8 out) closure first 5 argument
func Closure885[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(F, G, H) (I, J, K, L, M, N, O, P) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast885 with func( 8 in)(8 out) fix last 5 argument
func ClosureLast885[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure886 with func( 8 in)(8 out) closure first 6 argument
func Closure886[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(G, H) (I, J, K, L, M, N, O, P) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast886 with func( 8 in)(8 out) fix last 6 argument
func ClosureLast886[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B) (I, J, K, L, M, N, O, P) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure887 with func( 8 in)(8 out) closure first 7 argument
func Closure887[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(H) (I, J, K, L, M, N, O, P) {
	return func(h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast887 with func( 8 in)(8 out) fix last 7 argument
func ClosureLast887[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A) (I, J, K, L, M, N, O, P) {
	return func(a A) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure888 with func( 8 in)(8 out) closure first 8 argument
func Closure888[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func() (I, J, K, L, M, N, O, P) {
	return func() (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure891 with func( 8 in)(9 out) closure first 1 argument
func Closure891[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast891 with func( 8 in)(9 out) fix last 1 argument
func ClosureLast891[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure892 with func( 8 in)(9 out) closure first 2 argument
func Closure892[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast892 with func( 8 in)(9 out) fix last 2 argument
func ClosureLast892[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure893 with func( 8 in)(9 out) closure first 3 argument
func Closure893[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(D, E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast893 with func( 8 in)(9 out) fix last 3 argument
func ClosureLast893[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D, E) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure894 with func( 8 in)(9 out) closure first 4 argument
func Closure894[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast894 with func( 8 in)(9 out) fix last 4 argument
func ClosureLast894[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure895 with func( 8 in)(9 out) closure first 5 argument
func Closure895[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast895 with func( 8 in)(9 out) fix last 5 argument
func ClosureLast895[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure896 with func( 8 in)(9 out) closure first 6 argument
func Closure896[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast896 with func( 8 in)(9 out) fix last 6 argument
func ClosureLast896[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure897 with func( 8 in)(9 out) closure first 7 argument
func Closure897[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(H) (I, J, K, L, M, N, O, P, Q) {
	return func(h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosureLast897 with func( 8 in)(9 out) fix last 7 argument
func ClosureLast897[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A) (I, J, K, L, M, N, O, P, Q) {
	return func(a A) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure898 with func( 8 in)(9 out) closure first 8 argument
func Closure898[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func() (I, J, K, L, M, N, O, P, Q) {
	return func() (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Closure901 with func( 9 in)(0 out) closure first 1 argument
func Closure901[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F, G, H, I)) func(B, C, D, E, F, G, H, I) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast901 with func( 9 in)(0 out) fix last 1 argument
func ClosureLast901[A, B, C, D, E, F, G, H, I any](i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E, F, G, H) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure902 with func( 9 in)(0 out) closure first 2 argument
func Closure902[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F, G, H, I)) func(C, D, E, F, G, H, I) {
	return func(c C, d D, e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast902 with func( 9 in)(0 out) fix last 2 argument
func ClosureLast902[A, B, C, D, E, F, G, H, I any](h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E, F, G) {
	return func(a A, b B, c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure903 with func( 9 in)(0 out) closure first 3 argument
func Closure903[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I)) func(D, E, F, G, H, I) {
	return func(d D, e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast903 with func( 9 in)(0 out) fix last 3 argument
func ClosureLast903[A, B, C, D, E, F, G, H, I any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E, F) {
	return func(a A, b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure904 with func( 9 in)(0 out) closure first 4 argument
func Closure904[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I)) func(E, F, G, H, I) {
	return func(e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast904 with func( 9 in)(0 out) fix last 4 argument
func ClosureLast904[A, B, C, D, E, F, G, H, I any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure905 with func( 9 in)(0 out) closure first 5 argument
func Closure905[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I)) func(F, G, H, I) {
	return func(f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast905 with func( 9 in)(0 out) fix last 5 argument
func ClosureLast905[A, B, C, D, E, F, G, H, I any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure906 with func( 9 in)(0 out) closure first 6 argument
func Closure906[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I)) func(G, H, I) {
	return func(g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast906 with func( 9 in)(0 out) fix last 6 argument
func ClosureLast906[A, B, C, D, E, F, G, H, I any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure907 with func( 9 in)(0 out) closure first 7 argument
func Closure907[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I)) func(H, I) {
	return func(h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast907 with func( 9 in)(0 out) fix last 7 argument
func ClosureLast907[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure908 with func( 9 in)(0 out) closure first 8 argument
func Closure908[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I)) func(I) {
	return func(i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast908 with func( 9 in)(0 out) fix last 8 argument
func ClosureLast908[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure909 with func( 9 in)(0 out) closure first 9 argument
func Closure909[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func() {
	return func() {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure911 with func( 9 in)(1 out) closure first 1 argument
func Closure911[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F, G, H, I) J) func(B, C, D, E, F, G, H, I) J {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast911 with func( 9 in)(1 out) fix last 1 argument
func ClosureLast911[A, B, C, D, E, F, G, H, I, J any](i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E, F, G, H) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure912 with func( 9 in)(1 out) closure first 2 argument
func Closure912[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) J) func(C, D, E, F, G, H, I) J {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast912 with func( 9 in)(1 out) fix last 2 argument
func ClosureLast912[A, B, C, D, E, F, G, H, I, J any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E, F, G) J {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure913 with func( 9 in)(1 out) closure first 3 argument
func Closure913[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) J) func(D, E, F, G, H, I) J {
	return func(d D, e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast913 with func( 9 in)(1 out) fix last 3 argument
func ClosureLast913[A, B, C, D, E, F, G, H, I, J any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E, F) J {
	return func(a A, b B, c C, d D, e E, f F) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure914 with func( 9 in)(1 out) closure first 4 argument
func Closure914[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) J) func(E, F, G, H, I) J {
	return func(e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast914 with func( 9 in)(1 out) fix last 4 argument
func ClosureLast914[A, B, C, D, E, F, G, H, I, J any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E) J {
	return func(a A, b B, c C, d D, e E) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure915 with func( 9 in)(1 out) closure first 5 argument
func Closure915[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) J) func(F, G, H, I) J {
	return func(f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast915 with func( 9 in)(1 out) fix last 5 argument
func ClosureLast915[A, B, C, D, E, F, G, H, I, J any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D) J {
	return func(a A, b B, c C, d D) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure916 with func( 9 in)(1 out) closure first 6 argument
func Closure916[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) J) func(G, H, I) J {
	return func(g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast916 with func( 9 in)(1 out) fix last 6 argument
func ClosureLast916[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C) J {
	return func(a A, b B, c C) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure917 with func( 9 in)(1 out) closure first 7 argument
func Closure917[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) J) func(H, I) J {
	return func(h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast917 with func( 9 in)(1 out) fix last 7 argument
func ClosureLast917[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B) J {
	return func(a A, b B) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure918 with func( 9 in)(1 out) closure first 8 argument
func Closure918[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) J) func(I) J {
	return func(i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast918 with func( 9 in)(1 out) fix last 8 argument
func ClosureLast918[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A) J {
	return func(a A) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure919 with func( 9 in)(1 out) closure first 9 argument
func Closure919[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func() J {
	return func() (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure921 with func( 9 in)(2 out) closure first 1 argument
func Closure921[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(B, C, D, E, F, G, H, I) (J, K) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast921 with func( 9 in)(2 out) fix last 1 argument
func ClosureLast921[A, B, C, D, E, F, G, H, I, J, K any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E, F, G, H) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure922 with func( 9 in)(2 out) closure first 2 argument
func Closure922[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(C, D, E, F, G, H, I) (J, K) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast922 with func( 9 in)(2 out) fix last 2 argument
func ClosureLast922[A, B, C, D, E, F, G, H, I, J, K any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E, F, G) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure923 with func( 9 in)(2 out) closure first 3 argument
func Closure923[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(D, E, F, G, H, I) (J, K) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast923 with func( 9 in)(2 out) fix last 3 argument
func ClosureLast923[A, B, C, D, E, F, G, H, I, J, K any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E, F) (J, K) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure924 with func( 9 in)(2 out) closure first 4 argument
func Closure924[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(E, F, G, H, I) (J, K) {
	return func(e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast924 with func( 9 in)(2 out) fix last 4 argument
func ClosureLast924[A, B, C, D, E, F, G, H, I, J, K any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E) (J, K) {
	return func(a A, b B, c C, d D, e E) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure925 with func( 9 in)(2 out) closure first 5 argument
func Closure925[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(F, G, H, I) (J, K) {
	return func(f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast925 with func( 9 in)(2 out) fix last 5 argument
func ClosureLast925[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D) (J, K) {
	return func(a A, b B, c C, d D) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure926 with func( 9 in)(2 out) closure first 6 argument
func Closure926[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(G, H, I) (J, K) {
	return func(g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast926 with func( 9 in)(2 out) fix last 6 argument
func ClosureLast926[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C) (J, K) {
	return func(a A, b B, c C) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure927 with func( 9 in)(2 out) closure first 7 argument
func Closure927[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(H, I) (J, K) {
	return func(h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast927 with func( 9 in)(2 out) fix last 7 argument
func ClosureLast927[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B) (J, K) {
	return func(a A, b B) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure928 with func( 9 in)(2 out) closure first 8 argument
func Closure928[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(I) (J, K) {
	return func(i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast928 with func( 9 in)(2 out) fix last 8 argument
func ClosureLast928[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A) (J, K) {
	return func(a A) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure929 with func( 9 in)(2 out) closure first 9 argument
func Closure929[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func() (J, K) {
	return func() (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure931 with func( 9 in)(3 out) closure first 1 argument
func Closure931[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(B, C, D, E, F, G, H, I) (J, K, L) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast931 with func( 9 in)(3 out) fix last 1 argument
func ClosureLast931[A, B, C, D, E, F, G, H, I, J, K, L any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E, F, G, H) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure932 with func( 9 in)(3 out) closure first 2 argument
func Closure932[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(C, D, E, F, G, H, I) (J, K, L) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast932 with func( 9 in)(3 out) fix last 2 argument
func ClosureLast932[A, B, C, D, E, F, G, H, I, J, K, L any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E, F, G) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure933 with func( 9 in)(3 out) closure first 3 argument
func Closure933[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(D, E, F, G, H, I) (J, K, L) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast933 with func( 9 in)(3 out) fix last 3 argument
func ClosureLast933[A, B, C, D, E, F, G, H, I, J, K, L any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E, F) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure934 with func( 9 in)(3 out) closure first 4 argument
func Closure934[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(E, F, G, H, I) (J, K, L) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast934 with func( 9 in)(3 out) fix last 4 argument
func ClosureLast934[A, B, C, D, E, F, G, H, I, J, K, L any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E) (J, K, L) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure935 with func( 9 in)(3 out) closure first 5 argument
func Closure935[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(F, G, H, I) (J, K, L) {
	return func(f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast935 with func( 9 in)(3 out) fix last 5 argument
func ClosureLast935[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D) (J, K, L) {
	return func(a A, b B, c C, d D) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure936 with func( 9 in)(3 out) closure first 6 argument
func Closure936[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(G, H, I) (J, K, L) {
	return func(g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast936 with func( 9 in)(3 out) fix last 6 argument
func ClosureLast936[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C) (J, K, L) {
	return func(a A, b B, c C) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure937 with func( 9 in)(3 out) closure first 7 argument
func Closure937[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(H, I) (J, K, L) {
	return func(h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast937 with func( 9 in)(3 out) fix last 7 argument
func ClosureLast937[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B) (J, K, L) {
	return func(a A, b B) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure938 with func( 9 in)(3 out) closure first 8 argument
func Closure938[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(I) (J, K, L) {
	return func(i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast938 with func( 9 in)(3 out) fix last 8 argument
func ClosureLast938[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A) (J, K, L) {
	return func(a A) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure939 with func( 9 in)(3 out) closure first 9 argument
func Closure939[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func() (J, K, L) {
	return func() (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure941 with func( 9 in)(4 out) closure first 1 argument
func Closure941[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(B, C, D, E, F, G, H, I) (J, K, L, M) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast941 with func( 9 in)(4 out) fix last 1 argument
func ClosureLast941[A, B, C, D, E, F, G, H, I, J, K, L, M any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E, F, G, H) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure942 with func( 9 in)(4 out) closure first 2 argument
func Closure942[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(C, D, E, F, G, H, I) (J, K, L, M) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast942 with func( 9 in)(4 out) fix last 2 argument
func ClosureLast942[A, B, C, D, E, F, G, H, I, J, K, L, M any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E, F, G) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure943 with func( 9 in)(4 out) closure first 3 argument
func Closure943[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(D, E, F, G, H, I) (J, K, L, M) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast943 with func( 9 in)(4 out) fix last 3 argument
func ClosureLast943[A, B, C, D, E, F, G, H, I, J, K, L, M any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E, F) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure944 with func( 9 in)(4 out) closure first 4 argument
func Closure944[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(E, F, G, H, I) (J, K, L, M) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast944 with func( 9 in)(4 out) fix last 4 argument
func ClosureLast944[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure945 with func( 9 in)(4 out) closure first 5 argument
func Closure945[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(F, G, H, I) (J, K, L, M) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast945 with func( 9 in)(4 out) fix last 5 argument
func ClosureLast945[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D) (J, K, L, M) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure946 with func( 9 in)(4 out) closure first 6 argument
func Closure946[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(G, H, I) (J, K, L, M) {
	return func(g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast946 with func( 9 in)(4 out) fix last 6 argument
func ClosureLast946[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C) (J, K, L, M) {
	return func(a A, b B, c C) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure947 with func( 9 in)(4 out) closure first 7 argument
func Closure947[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(H, I) (J, K, L, M) {
	return func(h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast947 with func( 9 in)(4 out) fix last 7 argument
func ClosureLast947[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B) (J, K, L, M) {
	return func(a A, b B) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure948 with func( 9 in)(4 out) closure first 8 argument
func Closure948[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(I) (J, K, L, M) {
	return func(i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast948 with func( 9 in)(4 out) fix last 8 argument
func ClosureLast948[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A) (J, K, L, M) {
	return func(a A) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure949 with func( 9 in)(4 out) closure first 9 argument
func Closure949[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func() (J, K, L, M) {
	return func() (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure951 with func( 9 in)(5 out) closure first 1 argument
func Closure951[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast951 with func( 9 in)(5 out) fix last 1 argument
func ClosureLast951[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure952 with func( 9 in)(5 out) closure first 2 argument
func Closure952[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(C, D, E, F, G, H, I) (J, K, L, M, N) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast952 with func( 9 in)(5 out) fix last 2 argument
func ClosureLast952[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E, F, G) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure953 with func( 9 in)(5 out) closure first 3 argument
func Closure953[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(D, E, F, G, H, I) (J, K, L, M, N) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast953 with func( 9 in)(5 out) fix last 3 argument
func ClosureLast953[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E, F) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure954 with func( 9 in)(5 out) closure first 4 argument
func Closure954[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(E, F, G, H, I) (J, K, L, M, N) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast954 with func( 9 in)(5 out) fix last 4 argument
func ClosureLast954[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure955 with func( 9 in)(5 out) closure first 5 argument
func Closure955[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(F, G, H, I) (J, K, L, M, N) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast955 with func( 9 in)(5 out) fix last 5 argument
func ClosureLast955[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D) (J, K, L, M, N) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure956 with func( 9 in)(5 out) closure first 6 argument
func Closure956[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(G, H, I) (J, K, L, M, N) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast956 with func( 9 in)(5 out) fix last 6 argument
func ClosureLast956[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C) (J, K, L, M, N) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure957 with func( 9 in)(5 out) closure first 7 argument
func Closure957[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(H, I) (J, K, L, M, N) {
	return func(h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast957 with func( 9 in)(5 out) fix last 7 argument
func ClosureLast957[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B) (J, K, L, M, N) {
	return func(a A, b B) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure958 with func( 9 in)(5 out) closure first 8 argument
func Closure958[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(I) (J, K, L, M, N) {
	return func(i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast958 with func( 9 in)(5 out) fix last 8 argument
func ClosureLast958[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A) (J, K, L, M, N) {
	return func(a A) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure959 with func( 9 in)(5 out) closure first 9 argument
func Closure959[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func() (J, K, L, M, N) {
	return func() (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure961 with func( 9 in)(6 out) closure first 1 argument
func Closure961[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast961 with func( 9 in)(6 out) fix last 1 argument
func ClosureLast961[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure962 with func( 9 in)(6 out) closure first 2 argument
func Closure962[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast962 with func( 9 in)(6 out) fix last 2 argument
func ClosureLast962[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure963 with func( 9 in)(6 out) closure first 3 argument
func Closure963[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(D, E, F, G, H, I) (J, K, L, M, N, O) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast963 with func( 9 in)(6 out) fix last 3 argument
func ClosureLast963[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E, F) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure964 with func( 9 in)(6 out) closure first 4 argument
func Closure964[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(E, F, G, H, I) (J, K, L, M, N, O) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast964 with func( 9 in)(6 out) fix last 4 argument
func ClosureLast964[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure965 with func( 9 in)(6 out) closure first 5 argument
func Closure965[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(F, G, H, I) (J, K, L, M, N, O) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast965 with func( 9 in)(6 out) fix last 5 argument
func ClosureLast965[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure966 with func( 9 in)(6 out) closure first 6 argument
func Closure966[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(G, H, I) (J, K, L, M, N, O) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast966 with func( 9 in)(6 out) fix last 6 argument
func ClosureLast966[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C) (J, K, L, M, N, O) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure967 with func( 9 in)(6 out) closure first 7 argument
func Closure967[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(H, I) (J, K, L, M, N, O) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast967 with func( 9 in)(6 out) fix last 7 argument
func ClosureLast967[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B) (J, K, L, M, N, O) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure968 with func( 9 in)(6 out) closure first 8 argument
func Closure968[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(I) (J, K, L, M, N, O) {
	return func(i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast968 with func( 9 in)(6 out) fix last 8 argument
func ClosureLast968[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A) (J, K, L, M, N, O) {
	return func(a A) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure969 with func( 9 in)(6 out) closure first 9 argument
func Closure969[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func() (J, K, L, M, N, O) {
	return func() (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure971 with func( 9 in)(7 out) closure first 1 argument
func Closure971[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast971 with func( 9 in)(7 out) fix last 1 argument
func ClosureLast971[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure972 with func( 9 in)(7 out) closure first 2 argument
func Closure972[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast972 with func( 9 in)(7 out) fix last 2 argument
func ClosureLast972[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure973 with func( 9 in)(7 out) closure first 3 argument
func Closure973[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(D, E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast973 with func( 9 in)(7 out) fix last 3 argument
func ClosureLast973[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E, F) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure974 with func( 9 in)(7 out) closure first 4 argument
func Closure974[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast974 with func( 9 in)(7 out) fix last 4 argument
func ClosureLast974[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure975 with func( 9 in)(7 out) closure first 5 argument
func Closure975[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(F, G, H, I) (J, K, L, M, N, O, P) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast975 with func( 9 in)(7 out) fix last 5 argument
func ClosureLast975[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure976 with func( 9 in)(7 out) closure first 6 argument
func Closure976[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(G, H, I) (J, K, L, M, N, O, P) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast976 with func( 9 in)(7 out) fix last 6 argument
func ClosureLast976[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure977 with func( 9 in)(7 out) closure first 7 argument
func Closure977[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(H, I) (J, K, L, M, N, O, P) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast977 with func( 9 in)(7 out) fix last 7 argument
func ClosureLast977[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B) (J, K, L, M, N, O, P) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure978 with func( 9 in)(7 out) closure first 8 argument
func Closure978[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(I) (J, K, L, M, N, O, P) {
	return func(i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast978 with func( 9 in)(7 out) fix last 8 argument
func ClosureLast978[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A) (J, K, L, M, N, O, P) {
	return func(a A) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure979 with func( 9 in)(7 out) closure first 9 argument
func Closure979[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func() (J, K, L, M, N, O, P) {
	return func() (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure981 with func( 9 in)(8 out) closure first 1 argument
func Closure981[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast981 with func( 9 in)(8 out) fix last 1 argument
func ClosureLast981[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure982 with func( 9 in)(8 out) closure first 2 argument
func Closure982[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast982 with func( 9 in)(8 out) fix last 2 argument
func ClosureLast982[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure983 with func( 9 in)(8 out) closure first 3 argument
func Closure983[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(D, E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast983 with func( 9 in)(8 out) fix last 3 argument
func ClosureLast983[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure984 with func( 9 in)(8 out) closure first 4 argument
func Closure984[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast984 with func( 9 in)(8 out) fix last 4 argument
func ClosureLast984[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure985 with func( 9 in)(8 out) closure first 5 argument
func Closure985[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast985 with func( 9 in)(8 out) fix last 5 argument
func ClosureLast985[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure986 with func( 9 in)(8 out) closure first 6 argument
func Closure986[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast986 with func( 9 in)(8 out) fix last 6 argument
func ClosureLast986[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure987 with func( 9 in)(8 out) closure first 7 argument
func Closure987[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(H, I) (J, K, L, M, N, O, P, Q) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast987 with func( 9 in)(8 out) fix last 7 argument
func ClosureLast987[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure988 with func( 9 in)(8 out) closure first 8 argument
func Closure988[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(I) (J, K, L, M, N, O, P, Q) {
	return func(i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast988 with func( 9 in)(8 out) fix last 8 argument
func ClosureLast988[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A) (J, K, L, M, N, O, P, Q) {
	return func(a A) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure989 with func( 9 in)(8 out) closure first 9 argument
func Closure989[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func() (J, K, L, M, N, O, P, Q) {
	return func() (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure991 with func( 9 in)(9 out) closure first 1 argument
func Closure991[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast991 with func( 9 in)(9 out) fix last 1 argument
func ClosureLast991[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure992 with func( 9 in)(9 out) closure first 2 argument
func Closure992[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast992 with func( 9 in)(9 out) fix last 2 argument
func ClosureLast992[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure993 with func( 9 in)(9 out) closure first 3 argument
func Closure993[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast993 with func( 9 in)(9 out) fix last 3 argument
func ClosureLast993[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E, F) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure994 with func( 9 in)(9 out) closure first 4 argument
func Closure994[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast994 with func( 9 in)(9 out) fix last 4 argument
func ClosureLast994[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure995 with func( 9 in)(9 out) closure first 5 argument
func Closure995[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast995 with func( 9 in)(9 out) fix last 5 argument
func ClosureLast995[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure996 with func( 9 in)(9 out) closure first 6 argument
func Closure996[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast996 with func( 9 in)(9 out) fix last 6 argument
func ClosureLast996[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure997 with func( 9 in)(9 out) closure first 7 argument
func Closure997[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast997 with func( 9 in)(9 out) fix last 7 argument
func ClosureLast997[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure998 with func( 9 in)(9 out) closure first 8 argument
func Closure998[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(I) (J, K, L, M, N, O, P, Q, R) {
	return func(i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosureLast998 with func( 9 in)(9 out) fix last 8 argument
func ClosureLast998[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A) (J, K, L, M, N, O, P, Q, R) {
	return func(a A) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Closure999 with func( 9 in)(9 out) closure first 9 argument
func Closure999[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func() (J, K, L, M, N, O, P, Q, R) {
	return func() (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}
