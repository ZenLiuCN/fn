//go:build !gen

package fn

//generated file should not edit

// Clos101 with func( 1 in)(0 out) closure first 1 argument
func Clos101[A any](a A, fn func(A)) func() {
	return func() {
		fn(a)
		return
	}
}

// Clos111 with func( 1 in)(1 out) closure first 1 argument
func Clos111[A, B any](a A, fn func(A) B) func() B {
	return func() (b B) {
		b = fn(a)
		return
	}
}

// Clos121 with func( 1 in)(2 out) closure first 1 argument
func Clos121[A, B, C any](a A, fn func(A) (B, C)) func() (B, C) {
	return func() (b B, c C) {
		b, c = fn(a)
		return
	}
}

// Clos131 with func( 1 in)(3 out) closure first 1 argument
func Clos131[A, B, C, D any](a A, fn func(A) (B, C, D)) func() (B, C, D) {
	return func() (b B, c C, d D) {
		b, c, d = fn(a)
		return
	}
}

// Clos141 with func( 1 in)(4 out) closure first 1 argument
func Clos141[A, B, C, D, E any](a A, fn func(A) (B, C, D, E)) func() (B, C, D, E) {
	return func() (b B, c C, d D, e E) {
		b, c, d, e = fn(a)
		return
	}
}

// Clos151 with func( 1 in)(5 out) closure first 1 argument
func Clos151[A, B, C, D, E, F any](a A, fn func(A) (B, C, D, E, F)) func() (B, C, D, E, F) {
	return func() (b B, c C, d D, e E, f F) {
		b, c, d, e, f = fn(a)
		return
	}
}

// Clos161 with func( 1 in)(6 out) closure first 1 argument
func Clos161[A, B, C, D, E, F, G any](a A, fn func(A) (B, C, D, E, F, G)) func() (B, C, D, E, F, G) {
	return func() (b B, c C, d D, e E, f F, g G) {
		b, c, d, e, f, g = fn(a)
		return
	}
}

// Clos171 with func( 1 in)(7 out) closure first 1 argument
func Clos171[A, B, C, D, E, F, G, H any](a A, fn func(A) (B, C, D, E, F, G, H)) func() (B, C, D, E, F, G, H) {
	return func() (b B, c C, d D, e E, f F, g G, h H) {
		b, c, d, e, f, g, h = fn(a)
		return
	}
}

// Clos181 with func( 1 in)(8 out) closure first 1 argument
func Clos181[A, B, C, D, E, F, G, H, I any](a A, fn func(A) (B, C, D, E, F, G, H, I)) func() (B, C, D, E, F, G, H, I) {
	return func() (b B, c C, d D, e E, f F, g G, h H, i I) {
		b, c, d, e, f, g, h, i = fn(a)
		return
	}
}

// Clos191 with func( 1 in)(9 out) closure first 1 argument
func Clos191[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A) (B, C, D, E, F, G, H, I, J)) func() (B, C, D, E, F, G, H, I, J) {
	return func() (b B, c C, d D, e E, f F, g G, h H, i I, j J) {
		b, c, d, e, f, g, h, i, j = fn(a)
		return
	}
}

// Clos201 with func( 2 in)(0 out) closure first 1 argument
func Clos201[A, B any](a A, fn func(A, B)) func(B) {
	return func(b B) {
		fn(a, b)
		return
	}
}

// ClosLast201 with func( 2 in)(0 out) fix last 1 argument
func ClosLast201[A, B any](b B, fn func(A, B)) func(A) {
	return func(a A) {
		fn(a, b)
		return
	}
}

// Clos202 with func( 2 in)(0 out) closure first 2 argument
func Clos202[A, B any](a A, b B, fn func(A, B)) func() {
	return func() {
		fn(a, b)
		return
	}
}

// Clos211 with func( 2 in)(1 out) closure first 1 argument
func Clos211[A, B, C any](a A, fn func(A, B) C) func(B) C {
	return func(b B) (c C) {
		c = fn(a, b)
		return
	}
}

// ClosLast211 with func( 2 in)(1 out) fix last 1 argument
func ClosLast211[A, B, C any](b B, fn func(A, B) C) func(A) C {
	return func(a A) (c C) {
		c = fn(a, b)
		return
	}
}

// Clos212 with func( 2 in)(1 out) closure first 2 argument
func Clos212[A, B, C any](a A, b B, fn func(A, B) C) func() C {
	return func() (c C) {
		c = fn(a, b)
		return
	}
}

// Clos221 with func( 2 in)(2 out) closure first 1 argument
func Clos221[A, B, C, D any](a A, fn func(A, B) (C, D)) func(B) (C, D) {
	return func(b B) (c C, d D) {
		c, d = fn(a, b)
		return
	}
}

// ClosLast221 with func( 2 in)(2 out) fix last 1 argument
func ClosLast221[A, B, C, D any](b B, fn func(A, B) (C, D)) func(A) (C, D) {
	return func(a A) (c C, d D) {
		c, d = fn(a, b)
		return
	}
}

// Clos222 with func( 2 in)(2 out) closure first 2 argument
func Clos222[A, B, C, D any](a A, b B, fn func(A, B) (C, D)) func() (C, D) {
	return func() (c C, d D) {
		c, d = fn(a, b)
		return
	}
}

// Clos231 with func( 2 in)(3 out) closure first 1 argument
func Clos231[A, B, C, D, E any](a A, fn func(A, B) (C, D, E)) func(B) (C, D, E) {
	return func(b B) (c C, d D, e E) {
		c, d, e = fn(a, b)
		return
	}
}

// ClosLast231 with func( 2 in)(3 out) fix last 1 argument
func ClosLast231[A, B, C, D, E any](b B, fn func(A, B) (C, D, E)) func(A) (C, D, E) {
	return func(a A) (c C, d D, e E) {
		c, d, e = fn(a, b)
		return
	}
}

// Clos232 with func( 2 in)(3 out) closure first 2 argument
func Clos232[A, B, C, D, E any](a A, b B, fn func(A, B) (C, D, E)) func() (C, D, E) {
	return func() (c C, d D, e E) {
		c, d, e = fn(a, b)
		return
	}
}

// Clos241 with func( 2 in)(4 out) closure first 1 argument
func Clos241[A, B, C, D, E, F any](a A, fn func(A, B) (C, D, E, F)) func(B) (C, D, E, F) {
	return func(b B) (c C, d D, e E, f F) {
		c, d, e, f = fn(a, b)
		return
	}
}

// ClosLast241 with func( 2 in)(4 out) fix last 1 argument
func ClosLast241[A, B, C, D, E, F any](b B, fn func(A, B) (C, D, E, F)) func(A) (C, D, E, F) {
	return func(a A) (c C, d D, e E, f F) {
		c, d, e, f = fn(a, b)
		return
	}
}

// Clos242 with func( 2 in)(4 out) closure first 2 argument
func Clos242[A, B, C, D, E, F any](a A, b B, fn func(A, B) (C, D, E, F)) func() (C, D, E, F) {
	return func() (c C, d D, e E, f F) {
		c, d, e, f = fn(a, b)
		return
	}
}

// Clos251 with func( 2 in)(5 out) closure first 1 argument
func Clos251[A, B, C, D, E, F, G any](a A, fn func(A, B) (C, D, E, F, G)) func(B) (C, D, E, F, G) {
	return func(b B) (c C, d D, e E, f F, g G) {
		c, d, e, f, g = fn(a, b)
		return
	}
}

// ClosLast251 with func( 2 in)(5 out) fix last 1 argument
func ClosLast251[A, B, C, D, E, F, G any](b B, fn func(A, B) (C, D, E, F, G)) func(A) (C, D, E, F, G) {
	return func(a A) (c C, d D, e E, f F, g G) {
		c, d, e, f, g = fn(a, b)
		return
	}
}

// Clos252 with func( 2 in)(5 out) closure first 2 argument
func Clos252[A, B, C, D, E, F, G any](a A, b B, fn func(A, B) (C, D, E, F, G)) func() (C, D, E, F, G) {
	return func() (c C, d D, e E, f F, g G) {
		c, d, e, f, g = fn(a, b)
		return
	}
}

// Clos261 with func( 2 in)(6 out) closure first 1 argument
func Clos261[A, B, C, D, E, F, G, H any](a A, fn func(A, B) (C, D, E, F, G, H)) func(B) (C, D, E, F, G, H) {
	return func(b B) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// ClosLast261 with func( 2 in)(6 out) fix last 1 argument
func ClosLast261[A, B, C, D, E, F, G, H any](b B, fn func(A, B) (C, D, E, F, G, H)) func(A) (C, D, E, F, G, H) {
	return func(a A) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// Clos262 with func( 2 in)(6 out) closure first 2 argument
func Clos262[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B) (C, D, E, F, G, H)) func() (C, D, E, F, G, H) {
	return func() (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h = fn(a, b)
		return
	}
}

// Clos271 with func( 2 in)(7 out) closure first 1 argument
func Clos271[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B) (C, D, E, F, G, H, I)) func(B) (C, D, E, F, G, H, I) {
	return func(b B) (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// ClosLast271 with func( 2 in)(7 out) fix last 1 argument
func ClosLast271[A, B, C, D, E, F, G, H, I any](b B, fn func(A, B) (C, D, E, F, G, H, I)) func(A) (C, D, E, F, G, H, I) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// Clos272 with func( 2 in)(7 out) closure first 2 argument
func Clos272[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I)) func() (C, D, E, F, G, H, I) {
	return func() (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// Clos281 with func( 2 in)(8 out) closure first 1 argument
func Clos281[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B) (C, D, E, F, G, H, I, J)) func(B) (C, D, E, F, G, H, I, J) {
	return func(b B) (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// ClosLast281 with func( 2 in)(8 out) fix last 1 argument
func ClosLast281[A, B, C, D, E, F, G, H, I, J any](b B, fn func(A, B) (C, D, E, F, G, H, I, J)) func(A) (C, D, E, F, G, H, I, J) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// Clos282 with func( 2 in)(8 out) closure first 2 argument
func Clos282[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J)) func() (C, D, E, F, G, H, I, J) {
	return func() (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// Clos291 with func( 2 in)(9 out) closure first 1 argument
func Clos291[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(B) (C, D, E, F, G, H, I, J, K) {
	return func(b B) (c C, d D, e E, f F, g G, h H, i I, j J, k K) {
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// ClosLast291 with func( 2 in)(9 out) fix last 1 argument
func ClosLast291[A, B, C, D, E, F, G, H, I, J, K any](b B, fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(A) (C, D, E, F, G, H, I, J, K) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I, j J, k K) {
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// Clos292 with func( 2 in)(9 out) closure first 2 argument
func Clos292[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B) (C, D, E, F, G, H, I, J, K)) func() (C, D, E, F, G, H, I, J, K) {
	return func() (c C, d D, e E, f F, g G, h H, i I, j J, k K) {
		c, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// Clos301 with func( 3 in)(0 out) closure first 1 argument
func Clos301[A, B, C any](a A, fn func(A, B, C)) func(B, C) {
	return func(b B, c C) {
		fn(a, b, c)
		return
	}
}

// ClosLast301 with func( 3 in)(0 out) fix last 1 argument
func ClosLast301[A, B, C any](c C, fn func(A, B, C)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c)
		return
	}
}

// Clos302 with func( 3 in)(0 out) closure first 2 argument
func Clos302[A, B, C any](a A, b B, fn func(A, B, C)) func(C) {
	return func(c C) {
		fn(a, b, c)
		return
	}
}

// ClosLast302 with func( 3 in)(0 out) fix last 2 argument
func ClosLast302[A, B, C any](b B, c C, fn func(A, B, C)) func(A) {
	return func(a A) {
		fn(a, b, c)
		return
	}
}

// Clos303 with func( 3 in)(0 out) closure first 3 argument
func Clos303[A, B, C any](a A, b B, c C, fn func(A, B, C)) func() {
	return func() {
		fn(a, b, c)
		return
	}
}

// Clos311 with func( 3 in)(1 out) closure first 1 argument
func Clos311[A, B, C, D any](a A, fn func(A, B, C) D) func(B, C) D {
	return func(b B, c C) (d D) {
		d = fn(a, b, c)
		return
	}
}

// ClosLast311 with func( 3 in)(1 out) fix last 1 argument
func ClosLast311[A, B, C, D any](c C, fn func(A, B, C) D) func(A, B) D {
	return func(a A, b B) (d D) {
		d = fn(a, b, c)
		return
	}
}

// Clos312 with func( 3 in)(1 out) closure first 2 argument
func Clos312[A, B, C, D any](a A, b B, fn func(A, B, C) D) func(C) D {
	return func(c C) (d D) {
		d = fn(a, b, c)
		return
	}
}

// ClosLast312 with func( 3 in)(1 out) fix last 2 argument
func ClosLast312[A, B, C, D any](b B, c C, fn func(A, B, C) D) func(A) D {
	return func(a A) (d D) {
		d = fn(a, b, c)
		return
	}
}

// Clos313 with func( 3 in)(1 out) closure first 3 argument
func Clos313[A, B, C, D any](a A, b B, c C, fn func(A, B, C) D) func() D {
	return func() (d D) {
		d = fn(a, b, c)
		return
	}
}

// Clos321 with func( 3 in)(2 out) closure first 1 argument
func Clos321[A, B, C, D, E any](a A, fn func(A, B, C) (D, E)) func(B, C) (D, E) {
	return func(b B, c C) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// ClosLast321 with func( 3 in)(2 out) fix last 1 argument
func ClosLast321[A, B, C, D, E any](c C, fn func(A, B, C) (D, E)) func(A, B) (D, E) {
	return func(a A, b B) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// Clos322 with func( 3 in)(2 out) closure first 2 argument
func Clos322[A, B, C, D, E any](a A, b B, fn func(A, B, C) (D, E)) func(C) (D, E) {
	return func(c C) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// ClosLast322 with func( 3 in)(2 out) fix last 2 argument
func ClosLast322[A, B, C, D, E any](b B, c C, fn func(A, B, C) (D, E)) func(A) (D, E) {
	return func(a A) (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// Clos323 with func( 3 in)(2 out) closure first 3 argument
func Clos323[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C) (D, E)) func() (D, E) {
	return func() (d D, e E) {
		d, e = fn(a, b, c)
		return
	}
}

// Clos331 with func( 3 in)(3 out) closure first 1 argument
func Clos331[A, B, C, D, E, F any](a A, fn func(A, B, C) (D, E, F)) func(B, C) (D, E, F) {
	return func(b B, c C) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// ClosLast331 with func( 3 in)(3 out) fix last 1 argument
func ClosLast331[A, B, C, D, E, F any](c C, fn func(A, B, C) (D, E, F)) func(A, B) (D, E, F) {
	return func(a A, b B) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// Clos332 with func( 3 in)(3 out) closure first 2 argument
func Clos332[A, B, C, D, E, F any](a A, b B, fn func(A, B, C) (D, E, F)) func(C) (D, E, F) {
	return func(c C) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// ClosLast332 with func( 3 in)(3 out) fix last 2 argument
func ClosLast332[A, B, C, D, E, F any](b B, c C, fn func(A, B, C) (D, E, F)) func(A) (D, E, F) {
	return func(a A) (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// Clos333 with func( 3 in)(3 out) closure first 3 argument
func Clos333[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C) (D, E, F)) func() (D, E, F) {
	return func() (d D, e E, f F) {
		d, e, f = fn(a, b, c)
		return
	}
}

// Clos341 with func( 3 in)(4 out) closure first 1 argument
func Clos341[A, B, C, D, E, F, G any](a A, fn func(A, B, C) (D, E, F, G)) func(B, C) (D, E, F, G) {
	return func(b B, c C) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// ClosLast341 with func( 3 in)(4 out) fix last 1 argument
func ClosLast341[A, B, C, D, E, F, G any](c C, fn func(A, B, C) (D, E, F, G)) func(A, B) (D, E, F, G) {
	return func(a A, b B) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Clos342 with func( 3 in)(4 out) closure first 2 argument
func Clos342[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C) (D, E, F, G)) func(C) (D, E, F, G) {
	return func(c C) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// ClosLast342 with func( 3 in)(4 out) fix last 2 argument
func ClosLast342[A, B, C, D, E, F, G any](b B, c C, fn func(A, B, C) (D, E, F, G)) func(A) (D, E, F, G) {
	return func(a A) (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Clos343 with func( 3 in)(4 out) closure first 3 argument
func Clos343[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C) (D, E, F, G)) func() (D, E, F, G) {
	return func() (d D, e E, f F, g G) {
		d, e, f, g = fn(a, b, c)
		return
	}
}

// Clos351 with func( 3 in)(5 out) closure first 1 argument
func Clos351[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C) (D, E, F, G, H)) func(B, C) (D, E, F, G, H) {
	return func(b B, c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// ClosLast351 with func( 3 in)(5 out) fix last 1 argument
func ClosLast351[A, B, C, D, E, F, G, H any](c C, fn func(A, B, C) (D, E, F, G, H)) func(A, B) (D, E, F, G, H) {
	return func(a A, b B) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Clos352 with func( 3 in)(5 out) closure first 2 argument
func Clos352[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C) (D, E, F, G, H)) func(C) (D, E, F, G, H) {
	return func(c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// ClosLast352 with func( 3 in)(5 out) fix last 2 argument
func ClosLast352[A, B, C, D, E, F, G, H any](b B, c C, fn func(A, B, C) (D, E, F, G, H)) func(A) (D, E, F, G, H) {
	return func(a A) (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Clos353 with func( 3 in)(5 out) closure first 3 argument
func Clos353[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H)) func() (D, E, F, G, H) {
	return func() (d D, e E, f F, g G, h H) {
		d, e, f, g, h = fn(a, b, c)
		return
	}
}

// Clos361 with func( 3 in)(6 out) closure first 1 argument
func Clos361[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C) (D, E, F, G, H, I)) func(B, C) (D, E, F, G, H, I) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// ClosLast361 with func( 3 in)(6 out) fix last 1 argument
func ClosLast361[A, B, C, D, E, F, G, H, I any](c C, fn func(A, B, C) (D, E, F, G, H, I)) func(A, B) (D, E, F, G, H, I) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Clos362 with func( 3 in)(6 out) closure first 2 argument
func Clos362[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I)) func(C) (D, E, F, G, H, I) {
	return func(c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// ClosLast362 with func( 3 in)(6 out) fix last 2 argument
func ClosLast362[A, B, C, D, E, F, G, H, I any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I)) func(A) (D, E, F, G, H, I) {
	return func(a A) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Clos363 with func( 3 in)(6 out) closure first 3 argument
func Clos363[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I)) func() (D, E, F, G, H, I) {
	return func() (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// Clos371 with func( 3 in)(7 out) closure first 1 argument
func Clos371[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C) (D, E, F, G, H, I, J)) func(B, C) (D, E, F, G, H, I, J) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// ClosLast371 with func( 3 in)(7 out) fix last 1 argument
func ClosLast371[A, B, C, D, E, F, G, H, I, J any](c C, fn func(A, B, C) (D, E, F, G, H, I, J)) func(A, B) (D, E, F, G, H, I, J) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Clos372 with func( 3 in)(7 out) closure first 2 argument
func Clos372[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I, J)) func(C) (D, E, F, G, H, I, J) {
	return func(c C) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// ClosLast372 with func( 3 in)(7 out) fix last 2 argument
func ClosLast372[A, B, C, D, E, F, G, H, I, J any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J)) func(A) (D, E, F, G, H, I, J) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Clos373 with func( 3 in)(7 out) closure first 3 argument
func Clos373[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J)) func() (D, E, F, G, H, I, J) {
	return func() (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// Clos381 with func( 3 in)(8 out) closure first 1 argument
func Clos381[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(B, C) (D, E, F, G, H, I, J, K) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// ClosLast381 with func( 3 in)(8 out) fix last 1 argument
func ClosLast381[A, B, C, D, E, F, G, H, I, J, K any](c C, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(A, B) (D, E, F, G, H, I, J, K) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Clos382 with func( 3 in)(8 out) closure first 2 argument
func Clos382[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(C) (D, E, F, G, H, I, J, K) {
	return func(c C) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// ClosLast382 with func( 3 in)(8 out) fix last 2 argument
func ClosLast382[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(A) (D, E, F, G, H, I, J, K) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Clos383 with func( 3 in)(8 out) closure first 3 argument
func Clos383[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K)) func() (D, E, F, G, H, I, J, K) {
	return func() (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// Clos391 with func( 3 in)(9 out) closure first 1 argument
func Clos391[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(B, C) (D, E, F, G, H, I, J, K, L) {
	return func(b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// ClosLast391 with func( 3 in)(9 out) fix last 1 argument
func ClosLast391[A, B, C, D, E, F, G, H, I, J, K, L any](c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(A, B) (D, E, F, G, H, I, J, K, L) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Clos392 with func( 3 in)(9 out) closure first 2 argument
func Clos392[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(C) (D, E, F, G, H, I, J, K, L) {
	return func(c C) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// ClosLast392 with func( 3 in)(9 out) fix last 2 argument
func ClosLast392[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(A) (D, E, F, G, H, I, J, K, L) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Clos393 with func( 3 in)(9 out) closure first 3 argument
func Clos393[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func() (D, E, F, G, H, I, J, K, L) {
	return func() (d D, e E, f F, g G, h H, i I, j J, k K, l L) {
		d, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// Clos401 with func( 4 in)(0 out) closure first 1 argument
func Clos401[A, B, C, D any](a A, fn func(A, B, C, D)) func(B, C, D) {
	return func(b B, c C, d D) {
		fn(a, b, c, d)
		return
	}
}

// ClosLast401 with func( 4 in)(0 out) fix last 1 argument
func ClosLast401[A, B, C, D any](d D, fn func(A, B, C, D)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d)
		return
	}
}

// Clos402 with func( 4 in)(0 out) closure first 2 argument
func Clos402[A, B, C, D any](a A, b B, fn func(A, B, C, D)) func(C, D) {
	return func(c C, d D) {
		fn(a, b, c, d)
		return
	}
}

// ClosLast402 with func( 4 in)(0 out) fix last 2 argument
func ClosLast402[A, B, C, D any](c C, d D, fn func(A, B, C, D)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d)
		return
	}
}

// Clos403 with func( 4 in)(0 out) closure first 3 argument
func Clos403[A, B, C, D any](a A, b B, c C, fn func(A, B, C, D)) func(D) {
	return func(d D) {
		fn(a, b, c, d)
		return
	}
}

// ClosLast403 with func( 4 in)(0 out) fix last 3 argument
func ClosLast403[A, B, C, D any](b B, c C, d D, fn func(A, B, C, D)) func(A) {
	return func(a A) {
		fn(a, b, c, d)
		return
	}
}

// Clos404 with func( 4 in)(0 out) closure first 4 argument
func Clos404[A, B, C, D any](a A, b B, c C, d D, fn func(A, B, C, D)) func() {
	return func() {
		fn(a, b, c, d)
		return
	}
}

// Clos411 with func( 4 in)(1 out) closure first 1 argument
func Clos411[A, B, C, D, E any](a A, fn func(A, B, C, D) E) func(B, C, D) E {
	return func(b B, c C, d D) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// ClosLast411 with func( 4 in)(1 out) fix last 1 argument
func ClosLast411[A, B, C, D, E any](d D, fn func(A, B, C, D) E) func(A, B, C) E {
	return func(a A, b B, c C) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Clos412 with func( 4 in)(1 out) closure first 2 argument
func Clos412[A, B, C, D, E any](a A, b B, fn func(A, B, C, D) E) func(C, D) E {
	return func(c C, d D) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// ClosLast412 with func( 4 in)(1 out) fix last 2 argument
func ClosLast412[A, B, C, D, E any](c C, d D, fn func(A, B, C, D) E) func(A, B) E {
	return func(a A, b B) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Clos413 with func( 4 in)(1 out) closure first 3 argument
func Clos413[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C, D) E) func(D) E {
	return func(d D) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// ClosLast413 with func( 4 in)(1 out) fix last 3 argument
func ClosLast413[A, B, C, D, E any](b B, c C, d D, fn func(A, B, C, D) E) func(A) E {
	return func(a A) (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Clos414 with func( 4 in)(1 out) closure first 4 argument
func Clos414[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D) E) func() E {
	return func() (e E) {
		e = fn(a, b, c, d)
		return
	}
}

// Clos421 with func( 4 in)(2 out) closure first 1 argument
func Clos421[A, B, C, D, E, F any](a A, fn func(A, B, C, D) (E, F)) func(B, C, D) (E, F) {
	return func(b B, c C, d D) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// ClosLast421 with func( 4 in)(2 out) fix last 1 argument
func ClosLast421[A, B, C, D, E, F any](d D, fn func(A, B, C, D) (E, F)) func(A, B, C) (E, F) {
	return func(a A, b B, c C) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Clos422 with func( 4 in)(2 out) closure first 2 argument
func Clos422[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D) (E, F)) func(C, D) (E, F) {
	return func(c C, d D) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// ClosLast422 with func( 4 in)(2 out) fix last 2 argument
func ClosLast422[A, B, C, D, E, F any](c C, d D, fn func(A, B, C, D) (E, F)) func(A, B) (E, F) {
	return func(a A, b B) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Clos423 with func( 4 in)(2 out) closure first 3 argument
func Clos423[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D) (E, F)) func(D) (E, F) {
	return func(d D) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// ClosLast423 with func( 4 in)(2 out) fix last 3 argument
func ClosLast423[A, B, C, D, E, F any](b B, c C, d D, fn func(A, B, C, D) (E, F)) func(A) (E, F) {
	return func(a A) (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Clos424 with func( 4 in)(2 out) closure first 4 argument
func Clos424[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F)) func() (E, F) {
	return func() (e E, f F) {
		e, f = fn(a, b, c, d)
		return
	}
}

// Clos431 with func( 4 in)(3 out) closure first 1 argument
func Clos431[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D) (E, F, G)) func(B, C, D) (E, F, G) {
	return func(b B, c C, d D) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// ClosLast431 with func( 4 in)(3 out) fix last 1 argument
func ClosLast431[A, B, C, D, E, F, G any](d D, fn func(A, B, C, D) (E, F, G)) func(A, B, C) (E, F, G) {
	return func(a A, b B, c C) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Clos432 with func( 4 in)(3 out) closure first 2 argument
func Clos432[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D) (E, F, G)) func(C, D) (E, F, G) {
	return func(c C, d D) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// ClosLast432 with func( 4 in)(3 out) fix last 2 argument
func ClosLast432[A, B, C, D, E, F, G any](c C, d D, fn func(A, B, C, D) (E, F, G)) func(A, B) (E, F, G) {
	return func(a A, b B) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Clos433 with func( 4 in)(3 out) closure first 3 argument
func Clos433[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D) (E, F, G)) func(D) (E, F, G) {
	return func(d D) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// ClosLast433 with func( 4 in)(3 out) fix last 3 argument
func ClosLast433[A, B, C, D, E, F, G any](b B, c C, d D, fn func(A, B, C, D) (E, F, G)) func(A) (E, F, G) {
	return func(a A) (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Clos434 with func( 4 in)(3 out) closure first 4 argument
func Clos434[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G)) func() (E, F, G) {
	return func() (e E, f F, g G) {
		e, f, g = fn(a, b, c, d)
		return
	}
}

// Clos441 with func( 4 in)(4 out) closure first 1 argument
func Clos441[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D) (E, F, G, H)) func(B, C, D) (E, F, G, H) {
	return func(b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// ClosLast441 with func( 4 in)(4 out) fix last 1 argument
func ClosLast441[A, B, C, D, E, F, G, H any](d D, fn func(A, B, C, D) (E, F, G, H)) func(A, B, C) (E, F, G, H) {
	return func(a A, b B, c C) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Clos442 with func( 4 in)(4 out) closure first 2 argument
func Clos442[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D) (E, F, G, H)) func(C, D) (E, F, G, H) {
	return func(c C, d D) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// ClosLast442 with func( 4 in)(4 out) fix last 2 argument
func ClosLast442[A, B, C, D, E, F, G, H any](c C, d D, fn func(A, B, C, D) (E, F, G, H)) func(A, B) (E, F, G, H) {
	return func(a A, b B) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Clos443 with func( 4 in)(4 out) closure first 3 argument
func Clos443[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H)) func(D) (E, F, G, H) {
	return func(d D) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// ClosLast443 with func( 4 in)(4 out) fix last 3 argument
func ClosLast443[A, B, C, D, E, F, G, H any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H)) func(A) (E, F, G, H) {
	return func(a A) (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Clos444 with func( 4 in)(4 out) closure first 4 argument
func Clos444[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H)) func() (E, F, G, H) {
	return func() (e E, f F, g G, h H) {
		e, f, g, h = fn(a, b, c, d)
		return
	}
}

// Clos451 with func( 4 in)(5 out) closure first 1 argument
func Clos451[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D) (E, F, G, H, I)) func(B, C, D) (E, F, G, H, I) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// ClosLast451 with func( 4 in)(5 out) fix last 1 argument
func ClosLast451[A, B, C, D, E, F, G, H, I any](d D, fn func(A, B, C, D) (E, F, G, H, I)) func(A, B, C) (E, F, G, H, I) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Clos452 with func( 4 in)(5 out) closure first 2 argument
func Clos452[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I)) func(C, D) (E, F, G, H, I) {
	return func(c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// ClosLast452 with func( 4 in)(5 out) fix last 2 argument
func ClosLast452[A, B, C, D, E, F, G, H, I any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I)) func(A, B) (E, F, G, H, I) {
	return func(a A, b B) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Clos453 with func( 4 in)(5 out) closure first 3 argument
func Clos453[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I)) func(D) (E, F, G, H, I) {
	return func(d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// ClosLast453 with func( 4 in)(5 out) fix last 3 argument
func ClosLast453[A, B, C, D, E, F, G, H, I any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I)) func(A) (E, F, G, H, I) {
	return func(a A) (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Clos454 with func( 4 in)(5 out) closure first 4 argument
func Clos454[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I)) func() (E, F, G, H, I) {
	return func() (e E, f F, g G, h H, i I) {
		e, f, g, h, i = fn(a, b, c, d)
		return
	}
}

// Clos461 with func( 4 in)(6 out) closure first 1 argument
func Clos461[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D) (E, F, G, H, I, J)) func(B, C, D) (E, F, G, H, I, J) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// ClosLast461 with func( 4 in)(6 out) fix last 1 argument
func ClosLast461[A, B, C, D, E, F, G, H, I, J any](d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func(A, B, C) (E, F, G, H, I, J) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Clos462 with func( 4 in)(6 out) closure first 2 argument
func Clos462[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J)) func(C, D) (E, F, G, H, I, J) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// ClosLast462 with func( 4 in)(6 out) fix last 2 argument
func ClosLast462[A, B, C, D, E, F, G, H, I, J any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func(A, B) (E, F, G, H, I, J) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Clos463 with func( 4 in)(6 out) closure first 3 argument
func Clos463[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J)) func(D) (E, F, G, H, I, J) {
	return func(d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// ClosLast463 with func( 4 in)(6 out) fix last 3 argument
func ClosLast463[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func(A) (E, F, G, H, I, J) {
	return func(a A) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Clos464 with func( 4 in)(6 out) closure first 4 argument
func Clos464[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J)) func() (E, F, G, H, I, J) {
	return func() (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// Clos471 with func( 4 in)(7 out) closure first 1 argument
func Clos471[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(B, C, D) (E, F, G, H, I, J, K) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// ClosLast471 with func( 4 in)(7 out) fix last 1 argument
func ClosLast471[A, B, C, D, E, F, G, H, I, J, K any](d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(A, B, C) (E, F, G, H, I, J, K) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Clos472 with func( 4 in)(7 out) closure first 2 argument
func Clos472[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(C, D) (E, F, G, H, I, J, K) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// ClosLast472 with func( 4 in)(7 out) fix last 2 argument
func ClosLast472[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(A, B) (E, F, G, H, I, J, K) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Clos473 with func( 4 in)(7 out) closure first 3 argument
func Clos473[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(D) (E, F, G, H, I, J, K) {
	return func(d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// ClosLast473 with func( 4 in)(7 out) fix last 3 argument
func ClosLast473[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(A) (E, F, G, H, I, J, K) {
	return func(a A) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Clos474 with func( 4 in)(7 out) closure first 4 argument
func Clos474[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K)) func() (E, F, G, H, I, J, K) {
	return func() (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// Clos481 with func( 4 in)(8 out) closure first 1 argument
func Clos481[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(B, C, D) (E, F, G, H, I, J, K, L) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// ClosLast481 with func( 4 in)(8 out) fix last 1 argument
func ClosLast481[A, B, C, D, E, F, G, H, I, J, K, L any](d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(A, B, C) (E, F, G, H, I, J, K, L) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Clos482 with func( 4 in)(8 out) closure first 2 argument
func Clos482[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(C, D) (E, F, G, H, I, J, K, L) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// ClosLast482 with func( 4 in)(8 out) fix last 2 argument
func ClosLast482[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(A, B) (E, F, G, H, I, J, K, L) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Clos483 with func( 4 in)(8 out) closure first 3 argument
func Clos483[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(D) (E, F, G, H, I, J, K, L) {
	return func(d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// ClosLast483 with func( 4 in)(8 out) fix last 3 argument
func ClosLast483[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(A) (E, F, G, H, I, J, K, L) {
	return func(a A) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Clos484 with func( 4 in)(8 out) closure first 4 argument
func Clos484[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func() (E, F, G, H, I, J, K, L) {
	return func() (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// Clos491 with func( 4 in)(9 out) closure first 1 argument
func Clos491[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(B, C, D) (E, F, G, H, I, J, K, L, M) {
	return func(b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// ClosLast491 with func( 4 in)(9 out) fix last 1 argument
func ClosLast491[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(A, B, C) (E, F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Clos492 with func( 4 in)(9 out) closure first 2 argument
func Clos492[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(C, D) (E, F, G, H, I, J, K, L, M) {
	return func(c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// ClosLast492 with func( 4 in)(9 out) fix last 2 argument
func ClosLast492[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(A, B) (E, F, G, H, I, J, K, L, M) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Clos493 with func( 4 in)(9 out) closure first 3 argument
func Clos493[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(D) (E, F, G, H, I, J, K, L, M) {
	return func(d D) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// ClosLast493 with func( 4 in)(9 out) fix last 3 argument
func ClosLast493[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(A) (E, F, G, H, I, J, K, L, M) {
	return func(a A) (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Clos494 with func( 4 in)(9 out) closure first 4 argument
func Clos494[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func() (E, F, G, H, I, J, K, L, M) {
	return func() (e E, f F, g G, h H, i I, j J, k K, l L, m M) {
		e, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// Clos501 with func( 5 in)(0 out) closure first 1 argument
func Clos501[A, B, C, D, E any](a A, fn func(A, B, C, D, E)) func(B, C, D, E) {
	return func(b B, c C, d D, e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosLast501 with func( 5 in)(0 out) fix last 1 argument
func ClosLast501[A, B, C, D, E any](e E, fn func(A, B, C, D, E)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e)
		return
	}
}

// Clos502 with func( 5 in)(0 out) closure first 2 argument
func Clos502[A, B, C, D, E any](a A, b B, fn func(A, B, C, D, E)) func(C, D, E) {
	return func(c C, d D, e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosLast502 with func( 5 in)(0 out) fix last 2 argument
func ClosLast502[A, B, C, D, E any](d D, e E, fn func(A, B, C, D, E)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e)
		return
	}
}

// Clos503 with func( 5 in)(0 out) closure first 3 argument
func Clos503[A, B, C, D, E any](a A, b B, c C, fn func(A, B, C, D, E)) func(D, E) {
	return func(d D, e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosLast503 with func( 5 in)(0 out) fix last 3 argument
func ClosLast503[A, B, C, D, E any](c C, d D, e E, fn func(A, B, C, D, E)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e)
		return
	}
}

// Clos504 with func( 5 in)(0 out) closure first 4 argument
func Clos504[A, B, C, D, E any](a A, b B, c C, d D, fn func(A, B, C, D, E)) func(E) {
	return func(e E) {
		fn(a, b, c, d, e)
		return
	}
}

// ClosLast504 with func( 5 in)(0 out) fix last 4 argument
func ClosLast504[A, B, C, D, E any](b B, c C, d D, e E, fn func(A, B, C, D, E)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e)
		return
	}
}

// Clos505 with func( 5 in)(0 out) closure first 5 argument
func Clos505[A, B, C, D, E any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E)) func() {
	return func() {
		fn(a, b, c, d, e)
		return
	}
}

// Clos511 with func( 5 in)(1 out) closure first 1 argument
func Clos511[A, B, C, D, E, F any](a A, fn func(A, B, C, D, E) F) func(B, C, D, E) F {
	return func(b B, c C, d D, e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosLast511 with func( 5 in)(1 out) fix last 1 argument
func ClosLast511[A, B, C, D, E, F any](e E, fn func(A, B, C, D, E) F) func(A, B, C, D) F {
	return func(a A, b B, c C, d D) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Clos512 with func( 5 in)(1 out) closure first 2 argument
func Clos512[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D, E) F) func(C, D, E) F {
	return func(c C, d D, e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosLast512 with func( 5 in)(1 out) fix last 2 argument
func ClosLast512[A, B, C, D, E, F any](d D, e E, fn func(A, B, C, D, E) F) func(A, B, C) F {
	return func(a A, b B, c C) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Clos513 with func( 5 in)(1 out) closure first 3 argument
func Clos513[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D, E) F) func(D, E) F {
	return func(d D, e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosLast513 with func( 5 in)(1 out) fix last 3 argument
func ClosLast513[A, B, C, D, E, F any](c C, d D, e E, fn func(A, B, C, D, E) F) func(A, B) F {
	return func(a A, b B) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Clos514 with func( 5 in)(1 out) closure first 4 argument
func Clos514[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D, E) F) func(E) F {
	return func(e E) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// ClosLast514 with func( 5 in)(1 out) fix last 4 argument
func ClosLast514[A, B, C, D, E, F any](b B, c C, d D, e E, fn func(A, B, C, D, E) F) func(A) F {
	return func(a A) (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Clos515 with func( 5 in)(1 out) closure first 5 argument
func Clos515[A, B, C, D, E, F any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) F) func() F {
	return func() (f F) {
		f = fn(a, b, c, d, e)
		return
	}
}

// Clos521 with func( 5 in)(2 out) closure first 1 argument
func Clos521[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E) (F, G)) func(B, C, D, E) (F, G) {
	return func(b B, c C, d D, e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosLast521 with func( 5 in)(2 out) fix last 1 argument
func ClosLast521[A, B, C, D, E, F, G any](e E, fn func(A, B, C, D, E) (F, G)) func(A, B, C, D) (F, G) {
	return func(a A, b B, c C, d D) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Clos522 with func( 5 in)(2 out) closure first 2 argument
func Clos522[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E) (F, G)) func(C, D, E) (F, G) {
	return func(c C, d D, e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosLast522 with func( 5 in)(2 out) fix last 2 argument
func ClosLast522[A, B, C, D, E, F, G any](d D, e E, fn func(A, B, C, D, E) (F, G)) func(A, B, C) (F, G) {
	return func(a A, b B, c C) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Clos523 with func( 5 in)(2 out) closure first 3 argument
func Clos523[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E) (F, G)) func(D, E) (F, G) {
	return func(d D, e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosLast523 with func( 5 in)(2 out) fix last 3 argument
func ClosLast523[A, B, C, D, E, F, G any](c C, d D, e E, fn func(A, B, C, D, E) (F, G)) func(A, B) (F, G) {
	return func(a A, b B) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Clos524 with func( 5 in)(2 out) closure first 4 argument
func Clos524[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G)) func(E) (F, G) {
	return func(e E) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// ClosLast524 with func( 5 in)(2 out) fix last 4 argument
func ClosLast524[A, B, C, D, E, F, G any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G)) func(A) (F, G) {
	return func(a A) (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Clos525 with func( 5 in)(2 out) closure first 5 argument
func Clos525[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G)) func() (F, G) {
	return func() (f F, g G) {
		f, g = fn(a, b, c, d, e)
		return
	}
}

// Clos531 with func( 5 in)(3 out) closure first 1 argument
func Clos531[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E) (F, G, H)) func(B, C, D, E) (F, G, H) {
	return func(b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosLast531 with func( 5 in)(3 out) fix last 1 argument
func ClosLast531[A, B, C, D, E, F, G, H any](e E, fn func(A, B, C, D, E) (F, G, H)) func(A, B, C, D) (F, G, H) {
	return func(a A, b B, c C, d D) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Clos532 with func( 5 in)(3 out) closure first 2 argument
func Clos532[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E) (F, G, H)) func(C, D, E) (F, G, H) {
	return func(c C, d D, e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosLast532 with func( 5 in)(3 out) fix last 2 argument
func ClosLast532[A, B, C, D, E, F, G, H any](d D, e E, fn func(A, B, C, D, E) (F, G, H)) func(A, B, C) (F, G, H) {
	return func(a A, b B, c C) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Clos533 with func( 5 in)(3 out) closure first 3 argument
func Clos533[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H)) func(D, E) (F, G, H) {
	return func(d D, e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosLast533 with func( 5 in)(3 out) fix last 3 argument
func ClosLast533[A, B, C, D, E, F, G, H any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H)) func(A, B) (F, G, H) {
	return func(a A, b B) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Clos534 with func( 5 in)(3 out) closure first 4 argument
func Clos534[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H)) func(E) (F, G, H) {
	return func(e E) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// ClosLast534 with func( 5 in)(3 out) fix last 4 argument
func ClosLast534[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H)) func(A) (F, G, H) {
	return func(a A) (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Clos535 with func( 5 in)(3 out) closure first 5 argument
func Clos535[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H)) func() (F, G, H) {
	return func() (f F, g G, h H) {
		f, g, h = fn(a, b, c, d, e)
		return
	}
}

// Clos541 with func( 5 in)(4 out) closure first 1 argument
func Clos541[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E) (F, G, H, I)) func(B, C, D, E) (F, G, H, I) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosLast541 with func( 5 in)(4 out) fix last 1 argument
func ClosLast541[A, B, C, D, E, F, G, H, I any](e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A, B, C, D) (F, G, H, I) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Clos542 with func( 5 in)(4 out) closure first 2 argument
func Clos542[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I)) func(C, D, E) (F, G, H, I) {
	return func(c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosLast542 with func( 5 in)(4 out) fix last 2 argument
func ClosLast542[A, B, C, D, E, F, G, H, I any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A, B, C) (F, G, H, I) {
	return func(a A, b B, c C) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Clos543 with func( 5 in)(4 out) closure first 3 argument
func Clos543[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I)) func(D, E) (F, G, H, I) {
	return func(d D, e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosLast543 with func( 5 in)(4 out) fix last 3 argument
func ClosLast543[A, B, C, D, E, F, G, H, I any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A, B) (F, G, H, I) {
	return func(a A, b B) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Clos544 with func( 5 in)(4 out) closure first 4 argument
func Clos544[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I)) func(E) (F, G, H, I) {
	return func(e E) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// ClosLast544 with func( 5 in)(4 out) fix last 4 argument
func ClosLast544[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func(A) (F, G, H, I) {
	return func(a A) (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Clos545 with func( 5 in)(4 out) closure first 5 argument
func Clos545[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I)) func() (F, G, H, I) {
	return func() (f F, g G, h H, i I) {
		f, g, h, i = fn(a, b, c, d, e)
		return
	}
}

// Clos551 with func( 5 in)(5 out) closure first 1 argument
func Clos551[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E) (F, G, H, I, J)) func(B, C, D, E) (F, G, H, I, J) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosLast551 with func( 5 in)(5 out) fix last 1 argument
func ClosLast551[A, B, C, D, E, F, G, H, I, J any](e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A, B, C, D) (F, G, H, I, J) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Clos552 with func( 5 in)(5 out) closure first 2 argument
func Clos552[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J)) func(C, D, E) (F, G, H, I, J) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosLast552 with func( 5 in)(5 out) fix last 2 argument
func ClosLast552[A, B, C, D, E, F, G, H, I, J any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A, B, C) (F, G, H, I, J) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Clos553 with func( 5 in)(5 out) closure first 3 argument
func Clos553[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J)) func(D, E) (F, G, H, I, J) {
	return func(d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosLast553 with func( 5 in)(5 out) fix last 3 argument
func ClosLast553[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A, B) (F, G, H, I, J) {
	return func(a A, b B) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Clos554 with func( 5 in)(5 out) closure first 4 argument
func Clos554[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J)) func(E) (F, G, H, I, J) {
	return func(e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// ClosLast554 with func( 5 in)(5 out) fix last 4 argument
func ClosLast554[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func(A) (F, G, H, I, J) {
	return func(a A) (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Clos555 with func( 5 in)(5 out) closure first 5 argument
func Clos555[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J)) func() (F, G, H, I, J) {
	return func() (f F, g G, h H, i I, j J) {
		f, g, h, i, j = fn(a, b, c, d, e)
		return
	}
}

// Clos561 with func( 5 in)(6 out) closure first 1 argument
func Clos561[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(B, C, D, E) (F, G, H, I, J, K) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosLast561 with func( 5 in)(6 out) fix last 1 argument
func ClosLast561[A, B, C, D, E, F, G, H, I, J, K any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A, B, C, D) (F, G, H, I, J, K) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Clos562 with func( 5 in)(6 out) closure first 2 argument
func Clos562[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(C, D, E) (F, G, H, I, J, K) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosLast562 with func( 5 in)(6 out) fix last 2 argument
func ClosLast562[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A, B, C) (F, G, H, I, J, K) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Clos563 with func( 5 in)(6 out) closure first 3 argument
func Clos563[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(D, E) (F, G, H, I, J, K) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosLast563 with func( 5 in)(6 out) fix last 3 argument
func ClosLast563[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A, B) (F, G, H, I, J, K) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Clos564 with func( 5 in)(6 out) closure first 4 argument
func Clos564[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(E) (F, G, H, I, J, K) {
	return func(e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// ClosLast564 with func( 5 in)(6 out) fix last 4 argument
func ClosLast564[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(A) (F, G, H, I, J, K) {
	return func(a A) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Clos565 with func( 5 in)(6 out) closure first 5 argument
func Clos565[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K)) func() (F, G, H, I, J, K) {
	return func() (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// Clos571 with func( 5 in)(7 out) closure first 1 argument
func Clos571[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(B, C, D, E) (F, G, H, I, J, K, L) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosLast571 with func( 5 in)(7 out) fix last 1 argument
func ClosLast571[A, B, C, D, E, F, G, H, I, J, K, L any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A, B, C, D) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Clos572 with func( 5 in)(7 out) closure first 2 argument
func Clos572[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(C, D, E) (F, G, H, I, J, K, L) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosLast572 with func( 5 in)(7 out) fix last 2 argument
func ClosLast572[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A, B, C) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Clos573 with func( 5 in)(7 out) closure first 3 argument
func Clos573[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(D, E) (F, G, H, I, J, K, L) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosLast573 with func( 5 in)(7 out) fix last 3 argument
func ClosLast573[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A, B) (F, G, H, I, J, K, L) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Clos574 with func( 5 in)(7 out) closure first 4 argument
func Clos574[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(E) (F, G, H, I, J, K, L) {
	return func(e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// ClosLast574 with func( 5 in)(7 out) fix last 4 argument
func ClosLast574[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(A) (F, G, H, I, J, K, L) {
	return func(a A) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Clos575 with func( 5 in)(7 out) closure first 5 argument
func Clos575[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func() (F, G, H, I, J, K, L) {
	return func() (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// Clos581 with func( 5 in)(8 out) closure first 1 argument
func Clos581[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(B, C, D, E) (F, G, H, I, J, K, L, M) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosLast581 with func( 5 in)(8 out) fix last 1 argument
func ClosLast581[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A, B, C, D) (F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Clos582 with func( 5 in)(8 out) closure first 2 argument
func Clos582[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(C, D, E) (F, G, H, I, J, K, L, M) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosLast582 with func( 5 in)(8 out) fix last 2 argument
func ClosLast582[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A, B, C) (F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Clos583 with func( 5 in)(8 out) closure first 3 argument
func Clos583[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(D, E) (F, G, H, I, J, K, L, M) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosLast583 with func( 5 in)(8 out) fix last 3 argument
func ClosLast583[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A, B) (F, G, H, I, J, K, L, M) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Clos584 with func( 5 in)(8 out) closure first 4 argument
func Clos584[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(E) (F, G, H, I, J, K, L, M) {
	return func(e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// ClosLast584 with func( 5 in)(8 out) fix last 4 argument
func ClosLast584[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(A) (F, G, H, I, J, K, L, M) {
	return func(a A) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Clos585 with func( 5 in)(8 out) closure first 5 argument
func Clos585[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func() (F, G, H, I, J, K, L, M) {
	return func() (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// Clos591 with func( 5 in)(9 out) closure first 1 argument
func Clos591[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(B, C, D, E) (F, G, H, I, J, K, L, M, N) {
	return func(b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosLast591 with func( 5 in)(9 out) fix last 1 argument
func ClosLast591[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A, B, C, D) (F, G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Clos592 with func( 5 in)(9 out) closure first 2 argument
func Clos592[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(C, D, E) (F, G, H, I, J, K, L, M, N) {
	return func(c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosLast592 with func( 5 in)(9 out) fix last 2 argument
func ClosLast592[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A, B, C) (F, G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Clos593 with func( 5 in)(9 out) closure first 3 argument
func Clos593[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(D, E) (F, G, H, I, J, K, L, M, N) {
	return func(d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosLast593 with func( 5 in)(9 out) fix last 3 argument
func ClosLast593[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A, B) (F, G, H, I, J, K, L, M, N) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Clos594 with func( 5 in)(9 out) closure first 4 argument
func Clos594[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(E) (F, G, H, I, J, K, L, M, N) {
	return func(e E) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// ClosLast594 with func( 5 in)(9 out) fix last 4 argument
func ClosLast594[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(A) (F, G, H, I, J, K, L, M, N) {
	return func(a A) (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Clos595 with func( 5 in)(9 out) closure first 5 argument
func Clos595[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func() (F, G, H, I, J, K, L, M, N) {
	return func() (f F, g G, h H, i I, j J, k K, l L, m M, n N) {
		f, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// Clos601 with func( 6 in)(0 out) closure first 1 argument
func Clos601[A, B, C, D, E, F any](a A, fn func(A, B, C, D, E, F)) func(B, C, D, E, F) {
	return func(b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast601 with func( 6 in)(0 out) fix last 1 argument
func ClosLast601[A, B, C, D, E, F any](f F, fn func(A, B, C, D, E, F)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Clos602 with func( 6 in)(0 out) closure first 2 argument
func Clos602[A, B, C, D, E, F any](a A, b B, fn func(A, B, C, D, E, F)) func(C, D, E, F) {
	return func(c C, d D, e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast602 with func( 6 in)(0 out) fix last 2 argument
func ClosLast602[A, B, C, D, E, F any](e E, f F, fn func(A, B, C, D, E, F)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Clos603 with func( 6 in)(0 out) closure first 3 argument
func Clos603[A, B, C, D, E, F any](a A, b B, c C, fn func(A, B, C, D, E, F)) func(D, E, F) {
	return func(d D, e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast603 with func( 6 in)(0 out) fix last 3 argument
func ClosLast603[A, B, C, D, E, F any](d D, e E, f F, fn func(A, B, C, D, E, F)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Clos604 with func( 6 in)(0 out) closure first 4 argument
func Clos604[A, B, C, D, E, F any](a A, b B, c C, d D, fn func(A, B, C, D, E, F)) func(E, F) {
	return func(e E, f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast604 with func( 6 in)(0 out) fix last 4 argument
func ClosLast604[A, B, C, D, E, F any](c C, d D, e E, f F, fn func(A, B, C, D, E, F)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Clos605 with func( 6 in)(0 out) closure first 5 argument
func Clos605[A, B, C, D, E, F any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F)) func(F) {
	return func(f F) {
		fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast605 with func( 6 in)(0 out) fix last 5 argument
func ClosLast605[A, B, C, D, E, F any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f)
		return
	}
}

// Clos606 with func( 6 in)(0 out) closure first 6 argument
func Clos606[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F)) func() {
	return func() {
		fn(a, b, c, d, e, f)
		return
	}
}

// Clos611 with func( 6 in)(1 out) closure first 1 argument
func Clos611[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E, F) G) func(B, C, D, E, F) G {
	return func(b B, c C, d D, e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast611 with func( 6 in)(1 out) fix last 1 argument
func ClosLast611[A, B, C, D, E, F, G any](f F, fn func(A, B, C, D, E, F) G) func(A, B, C, D, E) G {
	return func(a A, b B, c C, d D, e E) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Clos612 with func( 6 in)(1 out) closure first 2 argument
func Clos612[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E, F) G) func(C, D, E, F) G {
	return func(c C, d D, e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast612 with func( 6 in)(1 out) fix last 2 argument
func ClosLast612[A, B, C, D, E, F, G any](e E, f F, fn func(A, B, C, D, E, F) G) func(A, B, C, D) G {
	return func(a A, b B, c C, d D) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Clos613 with func( 6 in)(1 out) closure first 3 argument
func Clos613[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E, F) G) func(D, E, F) G {
	return func(d D, e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast613 with func( 6 in)(1 out) fix last 3 argument
func ClosLast613[A, B, C, D, E, F, G any](d D, e E, f F, fn func(A, B, C, D, E, F) G) func(A, B, C) G {
	return func(a A, b B, c C) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Clos614 with func( 6 in)(1 out) closure first 4 argument
func Clos614[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) G) func(E, F) G {
	return func(e E, f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast614 with func( 6 in)(1 out) fix last 4 argument
func ClosLast614[A, B, C, D, E, F, G any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) G) func(A, B) G {
	return func(a A, b B) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Clos615 with func( 6 in)(1 out) closure first 5 argument
func Clos615[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) G) func(F) G {
	return func(f F) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast615 with func( 6 in)(1 out) fix last 5 argument
func ClosLast615[A, B, C, D, E, F, G any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) G) func(A) G {
	return func(a A) (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Clos616 with func( 6 in)(1 out) closure first 6 argument
func Clos616[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) G) func() G {
	return func() (g G) {
		g = fn(a, b, c, d, e, f)
		return
	}
}

// Clos621 with func( 6 in)(2 out) closure first 1 argument
func Clos621[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E, F) (G, H)) func(B, C, D, E, F) (G, H) {
	return func(b B, c C, d D, e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast621 with func( 6 in)(2 out) fix last 1 argument
func ClosLast621[A, B, C, D, E, F, G, H any](f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B, C, D, E) (G, H) {
	return func(a A, b B, c C, d D, e E) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Clos622 with func( 6 in)(2 out) closure first 2 argument
func Clos622[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E, F) (G, H)) func(C, D, E, F) (G, H) {
	return func(c C, d D, e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast622 with func( 6 in)(2 out) fix last 2 argument
func ClosLast622[A, B, C, D, E, F, G, H any](e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B, C, D) (G, H) {
	return func(a A, b B, c C, d D) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Clos623 with func( 6 in)(2 out) closure first 3 argument
func Clos623[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H)) func(D, E, F) (G, H) {
	return func(d D, e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast623 with func( 6 in)(2 out) fix last 3 argument
func ClosLast623[A, B, C, D, E, F, G, H any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B, C) (G, H) {
	return func(a A, b B, c C) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Clos624 with func( 6 in)(2 out) closure first 4 argument
func Clos624[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H)) func(E, F) (G, H) {
	return func(e E, f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast624 with func( 6 in)(2 out) fix last 4 argument
func ClosLast624[A, B, C, D, E, F, G, H any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A, B) (G, H) {
	return func(a A, b B) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Clos625 with func( 6 in)(2 out) closure first 5 argument
func Clos625[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H)) func(F) (G, H) {
	return func(f F) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast625 with func( 6 in)(2 out) fix last 5 argument
func ClosLast625[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func(A) (G, H) {
	return func(a A) (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Clos626 with func( 6 in)(2 out) closure first 6 argument
func Clos626[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H)) func() (G, H) {
	return func() (g G, h H) {
		g, h = fn(a, b, c, d, e, f)
		return
	}
}

// Clos631 with func( 6 in)(3 out) closure first 1 argument
func Clos631[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F) (G, H, I)) func(B, C, D, E, F) (G, H, I) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast631 with func( 6 in)(3 out) fix last 1 argument
func ClosLast631[A, B, C, D, E, F, G, H, I any](f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B, C, D, E) (G, H, I) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Clos632 with func( 6 in)(3 out) closure first 2 argument
func Clos632[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I)) func(C, D, E, F) (G, H, I) {
	return func(c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast632 with func( 6 in)(3 out) fix last 2 argument
func ClosLast632[A, B, C, D, E, F, G, H, I any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B, C, D) (G, H, I) {
	return func(a A, b B, c C, d D) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Clos633 with func( 6 in)(3 out) closure first 3 argument
func Clos633[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I)) func(D, E, F) (G, H, I) {
	return func(d D, e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast633 with func( 6 in)(3 out) fix last 3 argument
func ClosLast633[A, B, C, D, E, F, G, H, I any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B, C) (G, H, I) {
	return func(a A, b B, c C) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Clos634 with func( 6 in)(3 out) closure first 4 argument
func Clos634[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I)) func(E, F) (G, H, I) {
	return func(e E, f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast634 with func( 6 in)(3 out) fix last 4 argument
func ClosLast634[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A, B) (G, H, I) {
	return func(a A, b B) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Clos635 with func( 6 in)(3 out) closure first 5 argument
func Clos635[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I)) func(F) (G, H, I) {
	return func(f F) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast635 with func( 6 in)(3 out) fix last 5 argument
func ClosLast635[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func(A) (G, H, I) {
	return func(a A) (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Clos636 with func( 6 in)(3 out) closure first 6 argument
func Clos636[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I)) func() (G, H, I) {
	return func() (g G, h H, i I) {
		g, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// Clos641 with func( 6 in)(4 out) closure first 1 argument
func Clos641[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F) (G, H, I, J)) func(B, C, D, E, F) (G, H, I, J) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast641 with func( 6 in)(4 out) fix last 1 argument
func ClosLast641[A, B, C, D, E, F, G, H, I, J any](f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B, C, D, E) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Clos642 with func( 6 in)(4 out) closure first 2 argument
func Clos642[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J)) func(C, D, E, F) (G, H, I, J) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast642 with func( 6 in)(4 out) fix last 2 argument
func ClosLast642[A, B, C, D, E, F, G, H, I, J any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B, C, D) (G, H, I, J) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Clos643 with func( 6 in)(4 out) closure first 3 argument
func Clos643[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J)) func(D, E, F) (G, H, I, J) {
	return func(d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast643 with func( 6 in)(4 out) fix last 3 argument
func ClosLast643[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B, C) (G, H, I, J) {
	return func(a A, b B, c C) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Clos644 with func( 6 in)(4 out) closure first 4 argument
func Clos644[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J)) func(E, F) (G, H, I, J) {
	return func(e E, f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast644 with func( 6 in)(4 out) fix last 4 argument
func ClosLast644[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A, B) (G, H, I, J) {
	return func(a A, b B) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Clos645 with func( 6 in)(4 out) closure first 5 argument
func Clos645[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J)) func(F) (G, H, I, J) {
	return func(f F) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast645 with func( 6 in)(4 out) fix last 5 argument
func ClosLast645[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func(A) (G, H, I, J) {
	return func(a A) (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Clos646 with func( 6 in)(4 out) closure first 6 argument
func Clos646[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J)) func() (G, H, I, J) {
	return func() (g G, h H, i I, j J) {
		g, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// Clos651 with func( 6 in)(5 out) closure first 1 argument
func Clos651[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(B, C, D, E, F) (G, H, I, J, K) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast651 with func( 6 in)(5 out) fix last 1 argument
func ClosLast651[A, B, C, D, E, F, G, H, I, J, K any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B, C, D, E) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Clos652 with func( 6 in)(5 out) closure first 2 argument
func Clos652[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(C, D, E, F) (G, H, I, J, K) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast652 with func( 6 in)(5 out) fix last 2 argument
func ClosLast652[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B, C, D) (G, H, I, J, K) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Clos653 with func( 6 in)(5 out) closure first 3 argument
func Clos653[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(D, E, F) (G, H, I, J, K) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast653 with func( 6 in)(5 out) fix last 3 argument
func ClosLast653[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B, C) (G, H, I, J, K) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Clos654 with func( 6 in)(5 out) closure first 4 argument
func Clos654[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(E, F) (G, H, I, J, K) {
	return func(e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast654 with func( 6 in)(5 out) fix last 4 argument
func ClosLast654[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A, B) (G, H, I, J, K) {
	return func(a A, b B) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Clos655 with func( 6 in)(5 out) closure first 5 argument
func Clos655[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(F) (G, H, I, J, K) {
	return func(f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast655 with func( 6 in)(5 out) fix last 5 argument
func ClosLast655[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(A) (G, H, I, J, K) {
	return func(a A) (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Clos656 with func( 6 in)(5 out) closure first 6 argument
func Clos656[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K)) func() (G, H, I, J, K) {
	return func() (g G, h H, i I, j J, k K) {
		g, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// Clos661 with func( 6 in)(6 out) closure first 1 argument
func Clos661[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(B, C, D, E, F) (G, H, I, J, K, L) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast661 with func( 6 in)(6 out) fix last 1 argument
func ClosLast661[A, B, C, D, E, F, G, H, I, J, K, L any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B, C, D, E) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Clos662 with func( 6 in)(6 out) closure first 2 argument
func Clos662[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(C, D, E, F) (G, H, I, J, K, L) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast662 with func( 6 in)(6 out) fix last 2 argument
func ClosLast662[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B, C, D) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Clos663 with func( 6 in)(6 out) closure first 3 argument
func Clos663[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(D, E, F) (G, H, I, J, K, L) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast663 with func( 6 in)(6 out) fix last 3 argument
func ClosLast663[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B, C) (G, H, I, J, K, L) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Clos664 with func( 6 in)(6 out) closure first 4 argument
func Clos664[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(E, F) (G, H, I, J, K, L) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast664 with func( 6 in)(6 out) fix last 4 argument
func ClosLast664[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A, B) (G, H, I, J, K, L) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Clos665 with func( 6 in)(6 out) closure first 5 argument
func Clos665[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(F) (G, H, I, J, K, L) {
	return func(f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast665 with func( 6 in)(6 out) fix last 5 argument
func ClosLast665[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(A) (G, H, I, J, K, L) {
	return func(a A) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Clos666 with func( 6 in)(6 out) closure first 6 argument
func Clos666[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func() (G, H, I, J, K, L) {
	return func() (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// Clos671 with func( 6 in)(7 out) closure first 1 argument
func Clos671[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(B, C, D, E, F) (G, H, I, J, K, L, M) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast671 with func( 6 in)(7 out) fix last 1 argument
func ClosLast671[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B, C, D, E) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Clos672 with func( 6 in)(7 out) closure first 2 argument
func Clos672[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(C, D, E, F) (G, H, I, J, K, L, M) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast672 with func( 6 in)(7 out) fix last 2 argument
func ClosLast672[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B, C, D) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Clos673 with func( 6 in)(7 out) closure first 3 argument
func Clos673[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(D, E, F) (G, H, I, J, K, L, M) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast673 with func( 6 in)(7 out) fix last 3 argument
func ClosLast673[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B, C) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Clos674 with func( 6 in)(7 out) closure first 4 argument
func Clos674[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(E, F) (G, H, I, J, K, L, M) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast674 with func( 6 in)(7 out) fix last 4 argument
func ClosLast674[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A, B) (G, H, I, J, K, L, M) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Clos675 with func( 6 in)(7 out) closure first 5 argument
func Clos675[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(F) (G, H, I, J, K, L, M) {
	return func(f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast675 with func( 6 in)(7 out) fix last 5 argument
func ClosLast675[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(A) (G, H, I, J, K, L, M) {
	return func(a A) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Clos676 with func( 6 in)(7 out) closure first 6 argument
func Clos676[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func() (G, H, I, J, K, L, M) {
	return func() (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// Clos681 with func( 6 in)(8 out) closure first 1 argument
func Clos681[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(B, C, D, E, F) (G, H, I, J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast681 with func( 6 in)(8 out) fix last 1 argument
func ClosLast681[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B, C, D, E) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Clos682 with func( 6 in)(8 out) closure first 2 argument
func Clos682[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(C, D, E, F) (G, H, I, J, K, L, M, N) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast682 with func( 6 in)(8 out) fix last 2 argument
func ClosLast682[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B, C, D) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Clos683 with func( 6 in)(8 out) closure first 3 argument
func Clos683[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(D, E, F) (G, H, I, J, K, L, M, N) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast683 with func( 6 in)(8 out) fix last 3 argument
func ClosLast683[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B, C) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Clos684 with func( 6 in)(8 out) closure first 4 argument
func Clos684[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(E, F) (G, H, I, J, K, L, M, N) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast684 with func( 6 in)(8 out) fix last 4 argument
func ClosLast684[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A, B) (G, H, I, J, K, L, M, N) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Clos685 with func( 6 in)(8 out) closure first 5 argument
func Clos685[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(F) (G, H, I, J, K, L, M, N) {
	return func(f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast685 with func( 6 in)(8 out) fix last 5 argument
func ClosLast685[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(A) (G, H, I, J, K, L, M, N) {
	return func(a A) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Clos686 with func( 6 in)(8 out) closure first 6 argument
func Clos686[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func() (G, H, I, J, K, L, M, N) {
	return func() (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// Clos691 with func( 6 in)(9 out) closure first 1 argument
func Clos691[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(B, C, D, E, F) (G, H, I, J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast691 with func( 6 in)(9 out) fix last 1 argument
func ClosLast691[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B, C, D, E) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Clos692 with func( 6 in)(9 out) closure first 2 argument
func Clos692[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(C, D, E, F) (G, H, I, J, K, L, M, N, O) {
	return func(c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast692 with func( 6 in)(9 out) fix last 2 argument
func ClosLast692[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B, C, D) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Clos693 with func( 6 in)(9 out) closure first 3 argument
func Clos693[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(D, E, F) (G, H, I, J, K, L, M, N, O) {
	return func(d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast693 with func( 6 in)(9 out) fix last 3 argument
func ClosLast693[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B, C) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Clos694 with func( 6 in)(9 out) closure first 4 argument
func Clos694[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(E, F) (G, H, I, J, K, L, M, N, O) {
	return func(e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast694 with func( 6 in)(9 out) fix last 4 argument
func ClosLast694[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A, B) (G, H, I, J, K, L, M, N, O) {
	return func(a A, b B) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Clos695 with func( 6 in)(9 out) closure first 5 argument
func Clos695[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(F) (G, H, I, J, K, L, M, N, O) {
	return func(f F) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// ClosLast695 with func( 6 in)(9 out) fix last 5 argument
func ClosLast695[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(A) (G, H, I, J, K, L, M, N, O) {
	return func(a A) (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Clos696 with func( 6 in)(9 out) closure first 6 argument
func Clos696[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func() (G, H, I, J, K, L, M, N, O) {
	return func() (g G, h H, i I, j J, k K, l L, m M, n N, o O) {
		g, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// Clos701 with func( 7 in)(0 out) closure first 1 argument
func Clos701[A, B, C, D, E, F, G any](a A, fn func(A, B, C, D, E, F, G)) func(B, C, D, E, F, G) {
	return func(b B, c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast701 with func( 7 in)(0 out) fix last 1 argument
func ClosLast701[A, B, C, D, E, F, G any](g G, fn func(A, B, C, D, E, F, G)) func(A, B, C, D, E, F) {
	return func(a A, b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos702 with func( 7 in)(0 out) closure first 2 argument
func Clos702[A, B, C, D, E, F, G any](a A, b B, fn func(A, B, C, D, E, F, G)) func(C, D, E, F, G) {
	return func(c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast702 with func( 7 in)(0 out) fix last 2 argument
func ClosLast702[A, B, C, D, E, F, G any](f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos703 with func( 7 in)(0 out) closure first 3 argument
func Clos703[A, B, C, D, E, F, G any](a A, b B, c C, fn func(A, B, C, D, E, F, G)) func(D, E, F, G) {
	return func(d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast703 with func( 7 in)(0 out) fix last 3 argument
func ClosLast703[A, B, C, D, E, F, G any](e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos704 with func( 7 in)(0 out) closure first 4 argument
func Clos704[A, B, C, D, E, F, G any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G)) func(E, F, G) {
	return func(e E, f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast704 with func( 7 in)(0 out) fix last 4 argument
func ClosLast704[A, B, C, D, E, F, G any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos705 with func( 7 in)(0 out) closure first 5 argument
func Clos705[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G)) func(F, G) {
	return func(f F, g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast705 with func( 7 in)(0 out) fix last 5 argument
func ClosLast705[A, B, C, D, E, F, G any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos706 with func( 7 in)(0 out) closure first 6 argument
func Clos706[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G)) func(G) {
	return func(g G) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast706 with func( 7 in)(0 out) fix last 6 argument
func ClosLast706[A, B, C, D, E, F, G any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos707 with func( 7 in)(0 out) closure first 7 argument
func Clos707[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G)) func() {
	return func() {
		fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos711 with func( 7 in)(1 out) closure first 1 argument
func Clos711[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E, F, G) H) func(B, C, D, E, F, G) H {
	return func(b B, c C, d D, e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast711 with func( 7 in)(1 out) fix last 1 argument
func ClosLast711[A, B, C, D, E, F, G, H any](g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C, D, E, F) H {
	return func(a A, b B, c C, d D, e E, f F) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos712 with func( 7 in)(1 out) closure first 2 argument
func Clos712[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E, F, G) H) func(C, D, E, F, G) H {
	return func(c C, d D, e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast712 with func( 7 in)(1 out) fix last 2 argument
func ClosLast712[A, B, C, D, E, F, G, H any](f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C, D, E) H {
	return func(a A, b B, c C, d D, e E) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos713 with func( 7 in)(1 out) closure first 3 argument
func Clos713[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E, F, G) H) func(D, E, F, G) H {
	return func(d D, e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast713 with func( 7 in)(1 out) fix last 3 argument
func ClosLast713[A, B, C, D, E, F, G, H any](e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C, D) H {
	return func(a A, b B, c C, d D) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos714 with func( 7 in)(1 out) closure first 4 argument
func Clos714[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) H) func(E, F, G) H {
	return func(e E, f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast714 with func( 7 in)(1 out) fix last 4 argument
func ClosLast714[A, B, C, D, E, F, G, H any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B, C) H {
	return func(a A, b B, c C) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos715 with func( 7 in)(1 out) closure first 5 argument
func Clos715[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) H) func(F, G) H {
	return func(f F, g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast715 with func( 7 in)(1 out) fix last 5 argument
func ClosLast715[A, B, C, D, E, F, G, H any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A, B) H {
	return func(a A, b B) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos716 with func( 7 in)(1 out) closure first 6 argument
func Clos716[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) H) func(G) H {
	return func(g G) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast716 with func( 7 in)(1 out) fix last 6 argument
func ClosLast716[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func(A) H {
	return func(a A) (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos717 with func( 7 in)(1 out) closure first 7 argument
func Clos717[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) H) func() H {
	return func() (h H) {
		h = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos721 with func( 7 in)(2 out) closure first 1 argument
func Clos721[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F, G) (H, I)) func(B, C, D, E, F, G) (H, I) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast721 with func( 7 in)(2 out) fix last 1 argument
func ClosLast721[A, B, C, D, E, F, G, H, I any](g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C, D, E, F) (H, I) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos722 with func( 7 in)(2 out) closure first 2 argument
func Clos722[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I)) func(C, D, E, F, G) (H, I) {
	return func(c C, d D, e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast722 with func( 7 in)(2 out) fix last 2 argument
func ClosLast722[A, B, C, D, E, F, G, H, I any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C, D, E) (H, I) {
	return func(a A, b B, c C, d D, e E) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos723 with func( 7 in)(2 out) closure first 3 argument
func Clos723[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I)) func(D, E, F, G) (H, I) {
	return func(d D, e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast723 with func( 7 in)(2 out) fix last 3 argument
func ClosLast723[A, B, C, D, E, F, G, H, I any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C, D) (H, I) {
	return func(a A, b B, c C, d D) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos724 with func( 7 in)(2 out) closure first 4 argument
func Clos724[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I)) func(E, F, G) (H, I) {
	return func(e E, f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast724 with func( 7 in)(2 out) fix last 4 argument
func ClosLast724[A, B, C, D, E, F, G, H, I any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B, C) (H, I) {
	return func(a A, b B, c C) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos725 with func( 7 in)(2 out) closure first 5 argument
func Clos725[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I)) func(F, G) (H, I) {
	return func(f F, g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast725 with func( 7 in)(2 out) fix last 5 argument
func ClosLast725[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A, B) (H, I) {
	return func(a A, b B) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos726 with func( 7 in)(2 out) closure first 6 argument
func Clos726[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I)) func(G) (H, I) {
	return func(g G) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast726 with func( 7 in)(2 out) fix last 6 argument
func ClosLast726[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func(A) (H, I) {
	return func(a A) (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos727 with func( 7 in)(2 out) closure first 7 argument
func Clos727[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I)) func() (H, I) {
	return func() (h H, i I) {
		h, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos731 with func( 7 in)(3 out) closure first 1 argument
func Clos731[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F, G) (H, I, J)) func(B, C, D, E, F, G) (H, I, J) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast731 with func( 7 in)(3 out) fix last 1 argument
func ClosLast731[A, B, C, D, E, F, G, H, I, J any](g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C, D, E, F) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos732 with func( 7 in)(3 out) closure first 2 argument
func Clos732[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J)) func(C, D, E, F, G) (H, I, J) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast732 with func( 7 in)(3 out) fix last 2 argument
func ClosLast732[A, B, C, D, E, F, G, H, I, J any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C, D, E) (H, I, J) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos733 with func( 7 in)(3 out) closure first 3 argument
func Clos733[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J)) func(D, E, F, G) (H, I, J) {
	return func(d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast733 with func( 7 in)(3 out) fix last 3 argument
func ClosLast733[A, B, C, D, E, F, G, H, I, J any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C, D) (H, I, J) {
	return func(a A, b B, c C, d D) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos734 with func( 7 in)(3 out) closure first 4 argument
func Clos734[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J)) func(E, F, G) (H, I, J) {
	return func(e E, f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast734 with func( 7 in)(3 out) fix last 4 argument
func ClosLast734[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B, C) (H, I, J) {
	return func(a A, b B, c C) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos735 with func( 7 in)(3 out) closure first 5 argument
func Clos735[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J)) func(F, G) (H, I, J) {
	return func(f F, g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast735 with func( 7 in)(3 out) fix last 5 argument
func ClosLast735[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A, B) (H, I, J) {
	return func(a A, b B) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos736 with func( 7 in)(3 out) closure first 6 argument
func Clos736[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J)) func(G) (H, I, J) {
	return func(g G) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast736 with func( 7 in)(3 out) fix last 6 argument
func ClosLast736[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func(A) (H, I, J) {
	return func(a A) (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos737 with func( 7 in)(3 out) closure first 7 argument
func Clos737[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J)) func() (H, I, J) {
	return func() (h H, i I, j J) {
		h, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos741 with func( 7 in)(4 out) closure first 1 argument
func Clos741[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(B, C, D, E, F, G) (H, I, J, K) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast741 with func( 7 in)(4 out) fix last 1 argument
func ClosLast741[A, B, C, D, E, F, G, H, I, J, K any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C, D, E, F) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos742 with func( 7 in)(4 out) closure first 2 argument
func Clos742[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(C, D, E, F, G) (H, I, J, K) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast742 with func( 7 in)(4 out) fix last 2 argument
func ClosLast742[A, B, C, D, E, F, G, H, I, J, K any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C, D, E) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos743 with func( 7 in)(4 out) closure first 3 argument
func Clos743[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(D, E, F, G) (H, I, J, K) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast743 with func( 7 in)(4 out) fix last 3 argument
func ClosLast743[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C, D) (H, I, J, K) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos744 with func( 7 in)(4 out) closure first 4 argument
func Clos744[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(E, F, G) (H, I, J, K) {
	return func(e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast744 with func( 7 in)(4 out) fix last 4 argument
func ClosLast744[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B, C) (H, I, J, K) {
	return func(a A, b B, c C) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos745 with func( 7 in)(4 out) closure first 5 argument
func Clos745[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(F, G) (H, I, J, K) {
	return func(f F, g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast745 with func( 7 in)(4 out) fix last 5 argument
func ClosLast745[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A, B) (H, I, J, K) {
	return func(a A, b B) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos746 with func( 7 in)(4 out) closure first 6 argument
func Clos746[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(G) (H, I, J, K) {
	return func(g G) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast746 with func( 7 in)(4 out) fix last 6 argument
func ClosLast746[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(A) (H, I, J, K) {
	return func(a A) (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos747 with func( 7 in)(4 out) closure first 7 argument
func Clos747[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K)) func() (H, I, J, K) {
	return func() (h H, i I, j J, k K) {
		h, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos751 with func( 7 in)(5 out) closure first 1 argument
func Clos751[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(B, C, D, E, F, G) (H, I, J, K, L) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast751 with func( 7 in)(5 out) fix last 1 argument
func ClosLast751[A, B, C, D, E, F, G, H, I, J, K, L any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C, D, E, F) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos752 with func( 7 in)(5 out) closure first 2 argument
func Clos752[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(C, D, E, F, G) (H, I, J, K, L) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast752 with func( 7 in)(5 out) fix last 2 argument
func ClosLast752[A, B, C, D, E, F, G, H, I, J, K, L any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C, D, E) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos753 with func( 7 in)(5 out) closure first 3 argument
func Clos753[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(D, E, F, G) (H, I, J, K, L) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast753 with func( 7 in)(5 out) fix last 3 argument
func ClosLast753[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C, D) (H, I, J, K, L) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos754 with func( 7 in)(5 out) closure first 4 argument
func Clos754[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(E, F, G) (H, I, J, K, L) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast754 with func( 7 in)(5 out) fix last 4 argument
func ClosLast754[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B, C) (H, I, J, K, L) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos755 with func( 7 in)(5 out) closure first 5 argument
func Clos755[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(F, G) (H, I, J, K, L) {
	return func(f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast755 with func( 7 in)(5 out) fix last 5 argument
func ClosLast755[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A, B) (H, I, J, K, L) {
	return func(a A, b B) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos756 with func( 7 in)(5 out) closure first 6 argument
func Clos756[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(G) (H, I, J, K, L) {
	return func(g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast756 with func( 7 in)(5 out) fix last 6 argument
func ClosLast756[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(A) (H, I, J, K, L) {
	return func(a A) (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos757 with func( 7 in)(5 out) closure first 7 argument
func Clos757[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func() (H, I, J, K, L) {
	return func() (h H, i I, j J, k K, l L) {
		h, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos761 with func( 7 in)(6 out) closure first 1 argument
func Clos761[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(B, C, D, E, F, G) (H, I, J, K, L, M) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast761 with func( 7 in)(6 out) fix last 1 argument
func ClosLast761[A, B, C, D, E, F, G, H, I, J, K, L, M any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C, D, E, F) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos762 with func( 7 in)(6 out) closure first 2 argument
func Clos762[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(C, D, E, F, G) (H, I, J, K, L, M) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast762 with func( 7 in)(6 out) fix last 2 argument
func ClosLast762[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C, D, E) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos763 with func( 7 in)(6 out) closure first 3 argument
func Clos763[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(D, E, F, G) (H, I, J, K, L, M) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast763 with func( 7 in)(6 out) fix last 3 argument
func ClosLast763[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C, D) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos764 with func( 7 in)(6 out) closure first 4 argument
func Clos764[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(E, F, G) (H, I, J, K, L, M) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast764 with func( 7 in)(6 out) fix last 4 argument
func ClosLast764[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B, C) (H, I, J, K, L, M) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos765 with func( 7 in)(6 out) closure first 5 argument
func Clos765[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(F, G) (H, I, J, K, L, M) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast765 with func( 7 in)(6 out) fix last 5 argument
func ClosLast765[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A, B) (H, I, J, K, L, M) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos766 with func( 7 in)(6 out) closure first 6 argument
func Clos766[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(G) (H, I, J, K, L, M) {
	return func(g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast766 with func( 7 in)(6 out) fix last 6 argument
func ClosLast766[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(A) (H, I, J, K, L, M) {
	return func(a A) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos767 with func( 7 in)(6 out) closure first 7 argument
func Clos767[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func() (H, I, J, K, L, M) {
	return func() (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos771 with func( 7 in)(7 out) closure first 1 argument
func Clos771[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(B, C, D, E, F, G) (H, I, J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast771 with func( 7 in)(7 out) fix last 1 argument
func ClosLast771[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C, D, E, F) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos772 with func( 7 in)(7 out) closure first 2 argument
func Clos772[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(C, D, E, F, G) (H, I, J, K, L, M, N) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast772 with func( 7 in)(7 out) fix last 2 argument
func ClosLast772[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C, D, E) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos773 with func( 7 in)(7 out) closure first 3 argument
func Clos773[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(D, E, F, G) (H, I, J, K, L, M, N) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast773 with func( 7 in)(7 out) fix last 3 argument
func ClosLast773[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C, D) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos774 with func( 7 in)(7 out) closure first 4 argument
func Clos774[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(E, F, G) (H, I, J, K, L, M, N) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast774 with func( 7 in)(7 out) fix last 4 argument
func ClosLast774[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B, C) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos775 with func( 7 in)(7 out) closure first 5 argument
func Clos775[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(F, G) (H, I, J, K, L, M, N) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast775 with func( 7 in)(7 out) fix last 5 argument
func ClosLast775[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A, B) (H, I, J, K, L, M, N) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos776 with func( 7 in)(7 out) closure first 6 argument
func Clos776[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(G) (H, I, J, K, L, M, N) {
	return func(g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast776 with func( 7 in)(7 out) fix last 6 argument
func ClosLast776[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(A) (H, I, J, K, L, M, N) {
	return func(a A) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos777 with func( 7 in)(7 out) closure first 7 argument
func Clos777[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func() (H, I, J, K, L, M, N) {
	return func() (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos781 with func( 7 in)(8 out) closure first 1 argument
func Clos781[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(B, C, D, E, F, G) (H, I, J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast781 with func( 7 in)(8 out) fix last 1 argument
func ClosLast781[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C, D, E, F) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos782 with func( 7 in)(8 out) closure first 2 argument
func Clos782[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(C, D, E, F, G) (H, I, J, K, L, M, N, O) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast782 with func( 7 in)(8 out) fix last 2 argument
func ClosLast782[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C, D, E) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos783 with func( 7 in)(8 out) closure first 3 argument
func Clos783[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(D, E, F, G) (H, I, J, K, L, M, N, O) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast783 with func( 7 in)(8 out) fix last 3 argument
func ClosLast783[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C, D) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos784 with func( 7 in)(8 out) closure first 4 argument
func Clos784[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(E, F, G) (H, I, J, K, L, M, N, O) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast784 with func( 7 in)(8 out) fix last 4 argument
func ClosLast784[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B, C) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos785 with func( 7 in)(8 out) closure first 5 argument
func Clos785[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(F, G) (H, I, J, K, L, M, N, O) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast785 with func( 7 in)(8 out) fix last 5 argument
func ClosLast785[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A, B) (H, I, J, K, L, M, N, O) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos786 with func( 7 in)(8 out) closure first 6 argument
func Clos786[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(G) (H, I, J, K, L, M, N, O) {
	return func(g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast786 with func( 7 in)(8 out) fix last 6 argument
func ClosLast786[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(A) (H, I, J, K, L, M, N, O) {
	return func(a A) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos787 with func( 7 in)(8 out) closure first 7 argument
func Clos787[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func() (H, I, J, K, L, M, N, O) {
	return func() (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos791 with func( 7 in)(9 out) closure first 1 argument
func Clos791[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast791 with func( 7 in)(9 out) fix last 1 argument
func ClosLast791[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C, D, E, F) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos792 with func( 7 in)(9 out) closure first 2 argument
func Clos792[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(C, D, E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast792 with func( 7 in)(9 out) fix last 2 argument
func ClosLast792[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C, D, E) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos793 with func( 7 in)(9 out) closure first 3 argument
func Clos793[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(D, E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast793 with func( 7 in)(9 out) fix last 3 argument
func ClosLast793[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C, D) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos794 with func( 7 in)(9 out) closure first 4 argument
func Clos794[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(E, F, G) (H, I, J, K, L, M, N, O, P) {
	return func(e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast794 with func( 7 in)(9 out) fix last 4 argument
func ClosLast794[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B, C) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos795 with func( 7 in)(9 out) closure first 5 argument
func Clos795[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(F, G) (H, I, J, K, L, M, N, O, P) {
	return func(f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast795 with func( 7 in)(9 out) fix last 5 argument
func ClosLast795[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A, B) (H, I, J, K, L, M, N, O, P) {
	return func(a A, b B) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos796 with func( 7 in)(9 out) closure first 6 argument
func Clos796[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(G) (H, I, J, K, L, M, N, O, P) {
	return func(g G) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// ClosLast796 with func( 7 in)(9 out) fix last 6 argument
func ClosLast796[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(A) (H, I, J, K, L, M, N, O, P) {
	return func(a A) (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos797 with func( 7 in)(9 out) closure first 7 argument
func Clos797[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func() (H, I, J, K, L, M, N, O, P) {
	return func() (h H, i I, j J, k K, l L, m M, n N, o O, p P) {
		h, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// Clos801 with func( 8 in)(0 out) closure first 1 argument
func Clos801[A, B, C, D, E, F, G, H any](a A, fn func(A, B, C, D, E, F, G, H)) func(B, C, D, E, F, G, H) {
	return func(b B, c C, d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast801 with func( 8 in)(0 out) fix last 1 argument
func ClosLast801[A, B, C, D, E, F, G, H any](h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D, E, F, G) {
	return func(a A, b B, c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos802 with func( 8 in)(0 out) closure first 2 argument
func Clos802[A, B, C, D, E, F, G, H any](a A, b B, fn func(A, B, C, D, E, F, G, H)) func(C, D, E, F, G, H) {
	return func(c C, d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast802 with func( 8 in)(0 out) fix last 2 argument
func ClosLast802[A, B, C, D, E, F, G, H any](g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D, E, F) {
	return func(a A, b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos803 with func( 8 in)(0 out) closure first 3 argument
func Clos803[A, B, C, D, E, F, G, H any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H)) func(D, E, F, G, H) {
	return func(d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast803 with func( 8 in)(0 out) fix last 3 argument
func ClosLast803[A, B, C, D, E, F, G, H any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos804 with func( 8 in)(0 out) closure first 4 argument
func Clos804[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H)) func(E, F, G, H) {
	return func(e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast804 with func( 8 in)(0 out) fix last 4 argument
func ClosLast804[A, B, C, D, E, F, G, H any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos805 with func( 8 in)(0 out) closure first 5 argument
func Clos805[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H)) func(F, G, H) {
	return func(f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast805 with func( 8 in)(0 out) fix last 5 argument
func ClosLast805[A, B, C, D, E, F, G, H any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos806 with func( 8 in)(0 out) closure first 6 argument
func Clos806[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H)) func(G, H) {
	return func(g G, h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast806 with func( 8 in)(0 out) fix last 6 argument
func ClosLast806[A, B, C, D, E, F, G, H any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos807 with func( 8 in)(0 out) closure first 7 argument
func Clos807[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H)) func(H) {
	return func(h H) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast807 with func( 8 in)(0 out) fix last 7 argument
func ClosLast807[A, B, C, D, E, F, G, H any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos808 with func( 8 in)(0 out) closure first 8 argument
func Clos808[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H)) func() {
	return func() {
		fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos811 with func( 8 in)(1 out) closure first 1 argument
func Clos811[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F, G, H) I) func(B, C, D, E, F, G, H) I {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast811 with func( 8 in)(1 out) fix last 1 argument
func ClosLast811[A, B, C, D, E, F, G, H, I any](h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D, E, F, G) I {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos812 with func( 8 in)(1 out) closure first 2 argument
func Clos812[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F, G, H) I) func(C, D, E, F, G, H) I {
	return func(c C, d D, e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast812 with func( 8 in)(1 out) fix last 2 argument
func ClosLast812[A, B, C, D, E, F, G, H, I any](g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D, E, F) I {
	return func(a A, b B, c C, d D, e E, f F) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos813 with func( 8 in)(1 out) closure first 3 argument
func Clos813[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) I) func(D, E, F, G, H) I {
	return func(d D, e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast813 with func( 8 in)(1 out) fix last 3 argument
func ClosLast813[A, B, C, D, E, F, G, H, I any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D, E) I {
	return func(a A, b B, c C, d D, e E) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos814 with func( 8 in)(1 out) closure first 4 argument
func Clos814[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) I) func(E, F, G, H) I {
	return func(e E, f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast814 with func( 8 in)(1 out) fix last 4 argument
func ClosLast814[A, B, C, D, E, F, G, H, I any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C, D) I {
	return func(a A, b B, c C, d D) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos815 with func( 8 in)(1 out) closure first 5 argument
func Clos815[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) I) func(F, G, H) I {
	return func(f F, g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast815 with func( 8 in)(1 out) fix last 5 argument
func ClosLast815[A, B, C, D, E, F, G, H, I any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B, C) I {
	return func(a A, b B, c C) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos816 with func( 8 in)(1 out) closure first 6 argument
func Clos816[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) I) func(G, H) I {
	return func(g G, h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast816 with func( 8 in)(1 out) fix last 6 argument
func ClosLast816[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A, B) I {
	return func(a A, b B) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos817 with func( 8 in)(1 out) closure first 7 argument
func Clos817[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) I) func(H) I {
	return func(h H) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast817 with func( 8 in)(1 out) fix last 7 argument
func ClosLast817[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func(A) I {
	return func(a A) (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos818 with func( 8 in)(1 out) closure first 8 argument
func Clos818[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) I) func() I {
	return func() (i I) {
		i = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos821 with func( 8 in)(2 out) closure first 1 argument
func Clos821[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F, G, H) (I, J)) func(B, C, D, E, F, G, H) (I, J) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast821 with func( 8 in)(2 out) fix last 1 argument
func ClosLast821[A, B, C, D, E, F, G, H, I, J any](h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D, E, F, G) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos822 with func( 8 in)(2 out) closure first 2 argument
func Clos822[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J)) func(C, D, E, F, G, H) (I, J) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast822 with func( 8 in)(2 out) fix last 2 argument
func ClosLast822[A, B, C, D, E, F, G, H, I, J any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D, E, F) (I, J) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos823 with func( 8 in)(2 out) closure first 3 argument
func Clos823[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J)) func(D, E, F, G, H) (I, J) {
	return func(d D, e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast823 with func( 8 in)(2 out) fix last 3 argument
func ClosLast823[A, B, C, D, E, F, G, H, I, J any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D, E) (I, J) {
	return func(a A, b B, c C, d D, e E) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos824 with func( 8 in)(2 out) closure first 4 argument
func Clos824[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J)) func(E, F, G, H) (I, J) {
	return func(e E, f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast824 with func( 8 in)(2 out) fix last 4 argument
func ClosLast824[A, B, C, D, E, F, G, H, I, J any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C, D) (I, J) {
	return func(a A, b B, c C, d D) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos825 with func( 8 in)(2 out) closure first 5 argument
func Clos825[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J)) func(F, G, H) (I, J) {
	return func(f F, g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast825 with func( 8 in)(2 out) fix last 5 argument
func ClosLast825[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B, C) (I, J) {
	return func(a A, b B, c C) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos826 with func( 8 in)(2 out) closure first 6 argument
func Clos826[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J)) func(G, H) (I, J) {
	return func(g G, h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast826 with func( 8 in)(2 out) fix last 6 argument
func ClosLast826[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A, B) (I, J) {
	return func(a A, b B) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos827 with func( 8 in)(2 out) closure first 7 argument
func Clos827[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J)) func(H) (I, J) {
	return func(h H) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast827 with func( 8 in)(2 out) fix last 7 argument
func ClosLast827[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func(A) (I, J) {
	return func(a A) (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos828 with func( 8 in)(2 out) closure first 8 argument
func Clos828[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J)) func() (I, J) {
	return func() (i I, j J) {
		i, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos831 with func( 8 in)(3 out) closure first 1 argument
func Clos831[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(B, C, D, E, F, G, H) (I, J, K) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast831 with func( 8 in)(3 out) fix last 1 argument
func ClosLast831[A, B, C, D, E, F, G, H, I, J, K any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D, E, F, G) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos832 with func( 8 in)(3 out) closure first 2 argument
func Clos832[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(C, D, E, F, G, H) (I, J, K) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast832 with func( 8 in)(3 out) fix last 2 argument
func ClosLast832[A, B, C, D, E, F, G, H, I, J, K any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D, E, F) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos833 with func( 8 in)(3 out) closure first 3 argument
func Clos833[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(D, E, F, G, H) (I, J, K) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast833 with func( 8 in)(3 out) fix last 3 argument
func ClosLast833[A, B, C, D, E, F, G, H, I, J, K any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D, E) (I, J, K) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos834 with func( 8 in)(3 out) closure first 4 argument
func Clos834[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(E, F, G, H) (I, J, K) {
	return func(e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast834 with func( 8 in)(3 out) fix last 4 argument
func ClosLast834[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C, D) (I, J, K) {
	return func(a A, b B, c C, d D) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos835 with func( 8 in)(3 out) closure first 5 argument
func Clos835[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(F, G, H) (I, J, K) {
	return func(f F, g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast835 with func( 8 in)(3 out) fix last 5 argument
func ClosLast835[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B, C) (I, J, K) {
	return func(a A, b B, c C) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos836 with func( 8 in)(3 out) closure first 6 argument
func Clos836[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(G, H) (I, J, K) {
	return func(g G, h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast836 with func( 8 in)(3 out) fix last 6 argument
func ClosLast836[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A, B) (I, J, K) {
	return func(a A, b B) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos837 with func( 8 in)(3 out) closure first 7 argument
func Clos837[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(H) (I, J, K) {
	return func(h H) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast837 with func( 8 in)(3 out) fix last 7 argument
func ClosLast837[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(A) (I, J, K) {
	return func(a A) (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos838 with func( 8 in)(3 out) closure first 8 argument
func Clos838[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K)) func() (I, J, K) {
	return func() (i I, j J, k K) {
		i, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos841 with func( 8 in)(4 out) closure first 1 argument
func Clos841[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(B, C, D, E, F, G, H) (I, J, K, L) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast841 with func( 8 in)(4 out) fix last 1 argument
func ClosLast841[A, B, C, D, E, F, G, H, I, J, K, L any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D, E, F, G) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos842 with func( 8 in)(4 out) closure first 2 argument
func Clos842[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(C, D, E, F, G, H) (I, J, K, L) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast842 with func( 8 in)(4 out) fix last 2 argument
func ClosLast842[A, B, C, D, E, F, G, H, I, J, K, L any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D, E, F) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos843 with func( 8 in)(4 out) closure first 3 argument
func Clos843[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(D, E, F, G, H) (I, J, K, L) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast843 with func( 8 in)(4 out) fix last 3 argument
func ClosLast843[A, B, C, D, E, F, G, H, I, J, K, L any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D, E) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos844 with func( 8 in)(4 out) closure first 4 argument
func Clos844[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(E, F, G, H) (I, J, K, L) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast844 with func( 8 in)(4 out) fix last 4 argument
func ClosLast844[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C, D) (I, J, K, L) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos845 with func( 8 in)(4 out) closure first 5 argument
func Clos845[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(F, G, H) (I, J, K, L) {
	return func(f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast845 with func( 8 in)(4 out) fix last 5 argument
func ClosLast845[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B, C) (I, J, K, L) {
	return func(a A, b B, c C) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos846 with func( 8 in)(4 out) closure first 6 argument
func Clos846[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(G, H) (I, J, K, L) {
	return func(g G, h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast846 with func( 8 in)(4 out) fix last 6 argument
func ClosLast846[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A, B) (I, J, K, L) {
	return func(a A, b B) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos847 with func( 8 in)(4 out) closure first 7 argument
func Clos847[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(H) (I, J, K, L) {
	return func(h H) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast847 with func( 8 in)(4 out) fix last 7 argument
func ClosLast847[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(A) (I, J, K, L) {
	return func(a A) (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos848 with func( 8 in)(4 out) closure first 8 argument
func Clos848[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func() (I, J, K, L) {
	return func() (i I, j J, k K, l L) {
		i, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos851 with func( 8 in)(5 out) closure first 1 argument
func Clos851[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(B, C, D, E, F, G, H) (I, J, K, L, M) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast851 with func( 8 in)(5 out) fix last 1 argument
func ClosLast851[A, B, C, D, E, F, G, H, I, J, K, L, M any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D, E, F, G) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos852 with func( 8 in)(5 out) closure first 2 argument
func Clos852[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(C, D, E, F, G, H) (I, J, K, L, M) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast852 with func( 8 in)(5 out) fix last 2 argument
func ClosLast852[A, B, C, D, E, F, G, H, I, J, K, L, M any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D, E, F) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos853 with func( 8 in)(5 out) closure first 3 argument
func Clos853[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(D, E, F, G, H) (I, J, K, L, M) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast853 with func( 8 in)(5 out) fix last 3 argument
func ClosLast853[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D, E) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos854 with func( 8 in)(5 out) closure first 4 argument
func Clos854[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(E, F, G, H) (I, J, K, L, M) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast854 with func( 8 in)(5 out) fix last 4 argument
func ClosLast854[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C, D) (I, J, K, L, M) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos855 with func( 8 in)(5 out) closure first 5 argument
func Clos855[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(F, G, H) (I, J, K, L, M) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast855 with func( 8 in)(5 out) fix last 5 argument
func ClosLast855[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B, C) (I, J, K, L, M) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos856 with func( 8 in)(5 out) closure first 6 argument
func Clos856[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(G, H) (I, J, K, L, M) {
	return func(g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast856 with func( 8 in)(5 out) fix last 6 argument
func ClosLast856[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A, B) (I, J, K, L, M) {
	return func(a A, b B) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos857 with func( 8 in)(5 out) closure first 7 argument
func Clos857[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(H) (I, J, K, L, M) {
	return func(h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast857 with func( 8 in)(5 out) fix last 7 argument
func ClosLast857[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(A) (I, J, K, L, M) {
	return func(a A) (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos858 with func( 8 in)(5 out) closure first 8 argument
func Clos858[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func() (I, J, K, L, M) {
	return func() (i I, j J, k K, l L, m M) {
		i, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos861 with func( 8 in)(6 out) closure first 1 argument
func Clos861[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast861 with func( 8 in)(6 out) fix last 1 argument
func ClosLast861[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos862 with func( 8 in)(6 out) closure first 2 argument
func Clos862[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(C, D, E, F, G, H) (I, J, K, L, M, N) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast862 with func( 8 in)(6 out) fix last 2 argument
func ClosLast862[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D, E, F) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos863 with func( 8 in)(6 out) closure first 3 argument
func Clos863[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(D, E, F, G, H) (I, J, K, L, M, N) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast863 with func( 8 in)(6 out) fix last 3 argument
func ClosLast863[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D, E) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos864 with func( 8 in)(6 out) closure first 4 argument
func Clos864[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(E, F, G, H) (I, J, K, L, M, N) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast864 with func( 8 in)(6 out) fix last 4 argument
func ClosLast864[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C, D) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos865 with func( 8 in)(6 out) closure first 5 argument
func Clos865[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(F, G, H) (I, J, K, L, M, N) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast865 with func( 8 in)(6 out) fix last 5 argument
func ClosLast865[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B, C) (I, J, K, L, M, N) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos866 with func( 8 in)(6 out) closure first 6 argument
func Clos866[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(G, H) (I, J, K, L, M, N) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast866 with func( 8 in)(6 out) fix last 6 argument
func ClosLast866[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A, B) (I, J, K, L, M, N) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos867 with func( 8 in)(6 out) closure first 7 argument
func Clos867[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(H) (I, J, K, L, M, N) {
	return func(h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast867 with func( 8 in)(6 out) fix last 7 argument
func ClosLast867[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(A) (I, J, K, L, M, N) {
	return func(a A) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos868 with func( 8 in)(6 out) closure first 8 argument
func Clos868[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func() (I, J, K, L, M, N) {
	return func() (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos871 with func( 8 in)(7 out) closure first 1 argument
func Clos871[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast871 with func( 8 in)(7 out) fix last 1 argument
func ClosLast871[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos872 with func( 8 in)(7 out) closure first 2 argument
func Clos872[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(C, D, E, F, G, H) (I, J, K, L, M, N, O) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast872 with func( 8 in)(7 out) fix last 2 argument
func ClosLast872[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D, E, F) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos873 with func( 8 in)(7 out) closure first 3 argument
func Clos873[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(D, E, F, G, H) (I, J, K, L, M, N, O) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast873 with func( 8 in)(7 out) fix last 3 argument
func ClosLast873[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D, E) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos874 with func( 8 in)(7 out) closure first 4 argument
func Clos874[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(E, F, G, H) (I, J, K, L, M, N, O) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast874 with func( 8 in)(7 out) fix last 4 argument
func ClosLast874[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C, D) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos875 with func( 8 in)(7 out) closure first 5 argument
func Clos875[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(F, G, H) (I, J, K, L, M, N, O) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast875 with func( 8 in)(7 out) fix last 5 argument
func ClosLast875[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B, C) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos876 with func( 8 in)(7 out) closure first 6 argument
func Clos876[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(G, H) (I, J, K, L, M, N, O) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast876 with func( 8 in)(7 out) fix last 6 argument
func ClosLast876[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A, B) (I, J, K, L, M, N, O) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos877 with func( 8 in)(7 out) closure first 7 argument
func Clos877[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(H) (I, J, K, L, M, N, O) {
	return func(h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast877 with func( 8 in)(7 out) fix last 7 argument
func ClosLast877[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(A) (I, J, K, L, M, N, O) {
	return func(a A) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos878 with func( 8 in)(7 out) closure first 8 argument
func Clos878[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func() (I, J, K, L, M, N, O) {
	return func() (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos881 with func( 8 in)(8 out) closure first 1 argument
func Clos881[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast881 with func( 8 in)(8 out) fix last 1 argument
func ClosLast881[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos882 with func( 8 in)(8 out) closure first 2 argument
func Clos882[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(C, D, E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast882 with func( 8 in)(8 out) fix last 2 argument
func ClosLast882[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D, E, F) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos883 with func( 8 in)(8 out) closure first 3 argument
func Clos883[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(D, E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast883 with func( 8 in)(8 out) fix last 3 argument
func ClosLast883[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D, E) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos884 with func( 8 in)(8 out) closure first 4 argument
func Clos884[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(E, F, G, H) (I, J, K, L, M, N, O, P) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast884 with func( 8 in)(8 out) fix last 4 argument
func ClosLast884[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C, D) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos885 with func( 8 in)(8 out) closure first 5 argument
func Clos885[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(F, G, H) (I, J, K, L, M, N, O, P) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast885 with func( 8 in)(8 out) fix last 5 argument
func ClosLast885[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B, C) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos886 with func( 8 in)(8 out) closure first 6 argument
func Clos886[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(G, H) (I, J, K, L, M, N, O, P) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast886 with func( 8 in)(8 out) fix last 6 argument
func ClosLast886[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A, B) (I, J, K, L, M, N, O, P) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos887 with func( 8 in)(8 out) closure first 7 argument
func Clos887[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(H) (I, J, K, L, M, N, O, P) {
	return func(h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast887 with func( 8 in)(8 out) fix last 7 argument
func ClosLast887[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(A) (I, J, K, L, M, N, O, P) {
	return func(a A) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos888 with func( 8 in)(8 out) closure first 8 argument
func Clos888[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func() (I, J, K, L, M, N, O, P) {
	return func() (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos891 with func( 8 in)(9 out) closure first 1 argument
func Clos891[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast891 with func( 8 in)(9 out) fix last 1 argument
func ClosLast891[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F, G) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos892 with func( 8 in)(9 out) closure first 2 argument
func Clos892[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast892 with func( 8 in)(9 out) fix last 2 argument
func ClosLast892[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos893 with func( 8 in)(9 out) closure first 3 argument
func Clos893[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(D, E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast893 with func( 8 in)(9 out) fix last 3 argument
func ClosLast893[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D, E) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos894 with func( 8 in)(9 out) closure first 4 argument
func Clos894[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(E, F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast894 with func( 8 in)(9 out) fix last 4 argument
func ClosLast894[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C, D) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos895 with func( 8 in)(9 out) closure first 5 argument
func Clos895[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(F, G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast895 with func( 8 in)(9 out) fix last 5 argument
func ClosLast895[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B, C) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos896 with func( 8 in)(9 out) closure first 6 argument
func Clos896[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(G, H) (I, J, K, L, M, N, O, P, Q) {
	return func(g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast896 with func( 8 in)(9 out) fix last 6 argument
func ClosLast896[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A, B) (I, J, K, L, M, N, O, P, Q) {
	return func(a A, b B) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos897 with func( 8 in)(9 out) closure first 7 argument
func Clos897[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(H) (I, J, K, L, M, N, O, P, Q) {
	return func(h H) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// ClosLast897 with func( 8 in)(9 out) fix last 7 argument
func ClosLast897[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(A) (I, J, K, L, M, N, O, P, Q) {
	return func(a A) (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos898 with func( 8 in)(9 out) closure first 8 argument
func Clos898[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func() (I, J, K, L, M, N, O, P, Q) {
	return func() (i I, j J, k K, l L, m M, n N, o O, p P, q Q) {
		i, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Clos901 with func( 9 in)(0 out) closure first 1 argument
func Clos901[A, B, C, D, E, F, G, H, I any](a A, fn func(A, B, C, D, E, F, G, H, I)) func(B, C, D, E, F, G, H, I) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast901 with func( 9 in)(0 out) fix last 1 argument
func ClosLast901[A, B, C, D, E, F, G, H, I any](i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E, F, G, H) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos902 with func( 9 in)(0 out) closure first 2 argument
func Clos902[A, B, C, D, E, F, G, H, I any](a A, b B, fn func(A, B, C, D, E, F, G, H, I)) func(C, D, E, F, G, H, I) {
	return func(c C, d D, e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast902 with func( 9 in)(0 out) fix last 2 argument
func ClosLast902[A, B, C, D, E, F, G, H, I any](h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E, F, G) {
	return func(a A, b B, c C, d D, e E, f F, g G) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos903 with func( 9 in)(0 out) closure first 3 argument
func Clos903[A, B, C, D, E, F, G, H, I any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I)) func(D, E, F, G, H, I) {
	return func(d D, e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast903 with func( 9 in)(0 out) fix last 3 argument
func ClosLast903[A, B, C, D, E, F, G, H, I any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E, F) {
	return func(a A, b B, c C, d D, e E, f F) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos904 with func( 9 in)(0 out) closure first 4 argument
func Clos904[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I)) func(E, F, G, H, I) {
	return func(e E, f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast904 with func( 9 in)(0 out) fix last 4 argument
func ClosLast904[A, B, C, D, E, F, G, H, I any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D, E) {
	return func(a A, b B, c C, d D, e E) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos905 with func( 9 in)(0 out) closure first 5 argument
func Clos905[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I)) func(F, G, H, I) {
	return func(f F, g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast905 with func( 9 in)(0 out) fix last 5 argument
func ClosLast905[A, B, C, D, E, F, G, H, I any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C, D) {
	return func(a A, b B, c C, d D) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos906 with func( 9 in)(0 out) closure first 6 argument
func Clos906[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I)) func(G, H, I) {
	return func(g G, h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast906 with func( 9 in)(0 out) fix last 6 argument
func ClosLast906[A, B, C, D, E, F, G, H, I any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B, C) {
	return func(a A, b B, c C) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos907 with func( 9 in)(0 out) closure first 7 argument
func Clos907[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I)) func(H, I) {
	return func(h H, i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast907 with func( 9 in)(0 out) fix last 7 argument
func ClosLast907[A, B, C, D, E, F, G, H, I any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A, B) {
	return func(a A, b B) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos908 with func( 9 in)(0 out) closure first 8 argument
func Clos908[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I)) func(I) {
	return func(i I) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast908 with func( 9 in)(0 out) fix last 8 argument
func ClosLast908[A, B, C, D, E, F, G, H, I any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func(A) {
	return func(a A) {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos909 with func( 9 in)(0 out) closure first 9 argument
func Clos909[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I)) func() {
	return func() {
		fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos911 with func( 9 in)(1 out) closure first 1 argument
func Clos911[A, B, C, D, E, F, G, H, I, J any](a A, fn func(A, B, C, D, E, F, G, H, I) J) func(B, C, D, E, F, G, H, I) J {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast911 with func( 9 in)(1 out) fix last 1 argument
func ClosLast911[A, B, C, D, E, F, G, H, I, J any](i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E, F, G, H) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos912 with func( 9 in)(1 out) closure first 2 argument
func Clos912[A, B, C, D, E, F, G, H, I, J any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) J) func(C, D, E, F, G, H, I) J {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast912 with func( 9 in)(1 out) fix last 2 argument
func ClosLast912[A, B, C, D, E, F, G, H, I, J any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E, F, G) J {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos913 with func( 9 in)(1 out) closure first 3 argument
func Clos913[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) J) func(D, E, F, G, H, I) J {
	return func(d D, e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast913 with func( 9 in)(1 out) fix last 3 argument
func ClosLast913[A, B, C, D, E, F, G, H, I, J any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E, F) J {
	return func(a A, b B, c C, d D, e E, f F) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos914 with func( 9 in)(1 out) closure first 4 argument
func Clos914[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) J) func(E, F, G, H, I) J {
	return func(e E, f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast914 with func( 9 in)(1 out) fix last 4 argument
func ClosLast914[A, B, C, D, E, F, G, H, I, J any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D, E) J {
	return func(a A, b B, c C, d D, e E) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos915 with func( 9 in)(1 out) closure first 5 argument
func Clos915[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) J) func(F, G, H, I) J {
	return func(f F, g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast915 with func( 9 in)(1 out) fix last 5 argument
func ClosLast915[A, B, C, D, E, F, G, H, I, J any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C, D) J {
	return func(a A, b B, c C, d D) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos916 with func( 9 in)(1 out) closure first 6 argument
func Clos916[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) J) func(G, H, I) J {
	return func(g G, h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast916 with func( 9 in)(1 out) fix last 6 argument
func ClosLast916[A, B, C, D, E, F, G, H, I, J any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B, C) J {
	return func(a A, b B, c C) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos917 with func( 9 in)(1 out) closure first 7 argument
func Clos917[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) J) func(H, I) J {
	return func(h H, i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast917 with func( 9 in)(1 out) fix last 7 argument
func ClosLast917[A, B, C, D, E, F, G, H, I, J any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A, B) J {
	return func(a A, b B) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos918 with func( 9 in)(1 out) closure first 8 argument
func Clos918[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) J) func(I) J {
	return func(i I) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast918 with func( 9 in)(1 out) fix last 8 argument
func ClosLast918[A, B, C, D, E, F, G, H, I, J any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func(A) J {
	return func(a A) (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos919 with func( 9 in)(1 out) closure first 9 argument
func Clos919[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) J) func() J {
	return func() (j J) {
		j = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos921 with func( 9 in)(2 out) closure first 1 argument
func Clos921[A, B, C, D, E, F, G, H, I, J, K any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(B, C, D, E, F, G, H, I) (J, K) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast921 with func( 9 in)(2 out) fix last 1 argument
func ClosLast921[A, B, C, D, E, F, G, H, I, J, K any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E, F, G, H) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos922 with func( 9 in)(2 out) closure first 2 argument
func Clos922[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(C, D, E, F, G, H, I) (J, K) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast922 with func( 9 in)(2 out) fix last 2 argument
func ClosLast922[A, B, C, D, E, F, G, H, I, J, K any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E, F, G) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos923 with func( 9 in)(2 out) closure first 3 argument
func Clos923[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(D, E, F, G, H, I) (J, K) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast923 with func( 9 in)(2 out) fix last 3 argument
func ClosLast923[A, B, C, D, E, F, G, H, I, J, K any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E, F) (J, K) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos924 with func( 9 in)(2 out) closure first 4 argument
func Clos924[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(E, F, G, H, I) (J, K) {
	return func(e E, f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast924 with func( 9 in)(2 out) fix last 4 argument
func ClosLast924[A, B, C, D, E, F, G, H, I, J, K any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D, E) (J, K) {
	return func(a A, b B, c C, d D, e E) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos925 with func( 9 in)(2 out) closure first 5 argument
func Clos925[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(F, G, H, I) (J, K) {
	return func(f F, g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast925 with func( 9 in)(2 out) fix last 5 argument
func ClosLast925[A, B, C, D, E, F, G, H, I, J, K any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C, D) (J, K) {
	return func(a A, b B, c C, d D) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos926 with func( 9 in)(2 out) closure first 6 argument
func Clos926[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(G, H, I) (J, K) {
	return func(g G, h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast926 with func( 9 in)(2 out) fix last 6 argument
func ClosLast926[A, B, C, D, E, F, G, H, I, J, K any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B, C) (J, K) {
	return func(a A, b B, c C) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos927 with func( 9 in)(2 out) closure first 7 argument
func Clos927[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(H, I) (J, K) {
	return func(h H, i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast927 with func( 9 in)(2 out) fix last 7 argument
func ClosLast927[A, B, C, D, E, F, G, H, I, J, K any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A, B) (J, K) {
	return func(a A, b B) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos928 with func( 9 in)(2 out) closure first 8 argument
func Clos928[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(I) (J, K) {
	return func(i I) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast928 with func( 9 in)(2 out) fix last 8 argument
func ClosLast928[A, B, C, D, E, F, G, H, I, J, K any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(A) (J, K) {
	return func(a A) (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos929 with func( 9 in)(2 out) closure first 9 argument
func Clos929[A, B, C, D, E, F, G, H, I, J, K any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K)) func() (J, K) {
	return func() (j J, k K) {
		j, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos931 with func( 9 in)(3 out) closure first 1 argument
func Clos931[A, B, C, D, E, F, G, H, I, J, K, L any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(B, C, D, E, F, G, H, I) (J, K, L) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast931 with func( 9 in)(3 out) fix last 1 argument
func ClosLast931[A, B, C, D, E, F, G, H, I, J, K, L any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E, F, G, H) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos932 with func( 9 in)(3 out) closure first 2 argument
func Clos932[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(C, D, E, F, G, H, I) (J, K, L) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast932 with func( 9 in)(3 out) fix last 2 argument
func ClosLast932[A, B, C, D, E, F, G, H, I, J, K, L any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E, F, G) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos933 with func( 9 in)(3 out) closure first 3 argument
func Clos933[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(D, E, F, G, H, I) (J, K, L) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast933 with func( 9 in)(3 out) fix last 3 argument
func ClosLast933[A, B, C, D, E, F, G, H, I, J, K, L any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E, F) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos934 with func( 9 in)(3 out) closure first 4 argument
func Clos934[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(E, F, G, H, I) (J, K, L) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast934 with func( 9 in)(3 out) fix last 4 argument
func ClosLast934[A, B, C, D, E, F, G, H, I, J, K, L any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D, E) (J, K, L) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos935 with func( 9 in)(3 out) closure first 5 argument
func Clos935[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(F, G, H, I) (J, K, L) {
	return func(f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast935 with func( 9 in)(3 out) fix last 5 argument
func ClosLast935[A, B, C, D, E, F, G, H, I, J, K, L any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C, D) (J, K, L) {
	return func(a A, b B, c C, d D) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos936 with func( 9 in)(3 out) closure first 6 argument
func Clos936[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(G, H, I) (J, K, L) {
	return func(g G, h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast936 with func( 9 in)(3 out) fix last 6 argument
func ClosLast936[A, B, C, D, E, F, G, H, I, J, K, L any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B, C) (J, K, L) {
	return func(a A, b B, c C) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos937 with func( 9 in)(3 out) closure first 7 argument
func Clos937[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(H, I) (J, K, L) {
	return func(h H, i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast937 with func( 9 in)(3 out) fix last 7 argument
func ClosLast937[A, B, C, D, E, F, G, H, I, J, K, L any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A, B) (J, K, L) {
	return func(a A, b B) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos938 with func( 9 in)(3 out) closure first 8 argument
func Clos938[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(I) (J, K, L) {
	return func(i I) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast938 with func( 9 in)(3 out) fix last 8 argument
func ClosLast938[A, B, C, D, E, F, G, H, I, J, K, L any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(A) (J, K, L) {
	return func(a A) (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos939 with func( 9 in)(3 out) closure first 9 argument
func Clos939[A, B, C, D, E, F, G, H, I, J, K, L any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func() (J, K, L) {
	return func() (j J, k K, l L) {
		j, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos941 with func( 9 in)(4 out) closure first 1 argument
func Clos941[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(B, C, D, E, F, G, H, I) (J, K, L, M) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast941 with func( 9 in)(4 out) fix last 1 argument
func ClosLast941[A, B, C, D, E, F, G, H, I, J, K, L, M any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E, F, G, H) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos942 with func( 9 in)(4 out) closure first 2 argument
func Clos942[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(C, D, E, F, G, H, I) (J, K, L, M) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast942 with func( 9 in)(4 out) fix last 2 argument
func ClosLast942[A, B, C, D, E, F, G, H, I, J, K, L, M any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E, F, G) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos943 with func( 9 in)(4 out) closure first 3 argument
func Clos943[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(D, E, F, G, H, I) (J, K, L, M) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast943 with func( 9 in)(4 out) fix last 3 argument
func ClosLast943[A, B, C, D, E, F, G, H, I, J, K, L, M any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E, F) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos944 with func( 9 in)(4 out) closure first 4 argument
func Clos944[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(E, F, G, H, I) (J, K, L, M) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast944 with func( 9 in)(4 out) fix last 4 argument
func ClosLast944[A, B, C, D, E, F, G, H, I, J, K, L, M any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D, E) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos945 with func( 9 in)(4 out) closure first 5 argument
func Clos945[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(F, G, H, I) (J, K, L, M) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast945 with func( 9 in)(4 out) fix last 5 argument
func ClosLast945[A, B, C, D, E, F, G, H, I, J, K, L, M any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C, D) (J, K, L, M) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos946 with func( 9 in)(4 out) closure first 6 argument
func Clos946[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(G, H, I) (J, K, L, M) {
	return func(g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast946 with func( 9 in)(4 out) fix last 6 argument
func ClosLast946[A, B, C, D, E, F, G, H, I, J, K, L, M any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B, C) (J, K, L, M) {
	return func(a A, b B, c C) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos947 with func( 9 in)(4 out) closure first 7 argument
func Clos947[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(H, I) (J, K, L, M) {
	return func(h H, i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast947 with func( 9 in)(4 out) fix last 7 argument
func ClosLast947[A, B, C, D, E, F, G, H, I, J, K, L, M any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A, B) (J, K, L, M) {
	return func(a A, b B) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos948 with func( 9 in)(4 out) closure first 8 argument
func Clos948[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(I) (J, K, L, M) {
	return func(i I) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast948 with func( 9 in)(4 out) fix last 8 argument
func ClosLast948[A, B, C, D, E, F, G, H, I, J, K, L, M any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(A) (J, K, L, M) {
	return func(a A) (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos949 with func( 9 in)(4 out) closure first 9 argument
func Clos949[A, B, C, D, E, F, G, H, I, J, K, L, M any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func() (J, K, L, M) {
	return func() (j J, k K, l L, m M) {
		j, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos951 with func( 9 in)(5 out) closure first 1 argument
func Clos951[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast951 with func( 9 in)(5 out) fix last 1 argument
func ClosLast951[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos952 with func( 9 in)(5 out) closure first 2 argument
func Clos952[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(C, D, E, F, G, H, I) (J, K, L, M, N) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast952 with func( 9 in)(5 out) fix last 2 argument
func ClosLast952[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E, F, G) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos953 with func( 9 in)(5 out) closure first 3 argument
func Clos953[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(D, E, F, G, H, I) (J, K, L, M, N) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast953 with func( 9 in)(5 out) fix last 3 argument
func ClosLast953[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E, F) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos954 with func( 9 in)(5 out) closure first 4 argument
func Clos954[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(E, F, G, H, I) (J, K, L, M, N) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast954 with func( 9 in)(5 out) fix last 4 argument
func ClosLast954[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D, E) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos955 with func( 9 in)(5 out) closure first 5 argument
func Clos955[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(F, G, H, I) (J, K, L, M, N) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast955 with func( 9 in)(5 out) fix last 5 argument
func ClosLast955[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C, D) (J, K, L, M, N) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos956 with func( 9 in)(5 out) closure first 6 argument
func Clos956[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(G, H, I) (J, K, L, M, N) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast956 with func( 9 in)(5 out) fix last 6 argument
func ClosLast956[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B, C) (J, K, L, M, N) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos957 with func( 9 in)(5 out) closure first 7 argument
func Clos957[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(H, I) (J, K, L, M, N) {
	return func(h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast957 with func( 9 in)(5 out) fix last 7 argument
func ClosLast957[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A, B) (J, K, L, M, N) {
	return func(a A, b B) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos958 with func( 9 in)(5 out) closure first 8 argument
func Clos958[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(I) (J, K, L, M, N) {
	return func(i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast958 with func( 9 in)(5 out) fix last 8 argument
func ClosLast958[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(A) (J, K, L, M, N) {
	return func(a A) (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos959 with func( 9 in)(5 out) closure first 9 argument
func Clos959[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func() (J, K, L, M, N) {
	return func() (j J, k K, l L, m M, n N) {
		j, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos961 with func( 9 in)(6 out) closure first 1 argument
func Clos961[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast961 with func( 9 in)(6 out) fix last 1 argument
func ClosLast961[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos962 with func( 9 in)(6 out) closure first 2 argument
func Clos962[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast962 with func( 9 in)(6 out) fix last 2 argument
func ClosLast962[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos963 with func( 9 in)(6 out) closure first 3 argument
func Clos963[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(D, E, F, G, H, I) (J, K, L, M, N, O) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast963 with func( 9 in)(6 out) fix last 3 argument
func ClosLast963[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E, F) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos964 with func( 9 in)(6 out) closure first 4 argument
func Clos964[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(E, F, G, H, I) (J, K, L, M, N, O) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast964 with func( 9 in)(6 out) fix last 4 argument
func ClosLast964[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D, E) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos965 with func( 9 in)(6 out) closure first 5 argument
func Clos965[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(F, G, H, I) (J, K, L, M, N, O) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast965 with func( 9 in)(6 out) fix last 5 argument
func ClosLast965[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C, D) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos966 with func( 9 in)(6 out) closure first 6 argument
func Clos966[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(G, H, I) (J, K, L, M, N, O) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast966 with func( 9 in)(6 out) fix last 6 argument
func ClosLast966[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B, C) (J, K, L, M, N, O) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos967 with func( 9 in)(6 out) closure first 7 argument
func Clos967[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(H, I) (J, K, L, M, N, O) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast967 with func( 9 in)(6 out) fix last 7 argument
func ClosLast967[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A, B) (J, K, L, M, N, O) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos968 with func( 9 in)(6 out) closure first 8 argument
func Clos968[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(I) (J, K, L, M, N, O) {
	return func(i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast968 with func( 9 in)(6 out) fix last 8 argument
func ClosLast968[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(A) (J, K, L, M, N, O) {
	return func(a A) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos969 with func( 9 in)(6 out) closure first 9 argument
func Clos969[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func() (J, K, L, M, N, O) {
	return func() (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos971 with func( 9 in)(7 out) closure first 1 argument
func Clos971[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast971 with func( 9 in)(7 out) fix last 1 argument
func ClosLast971[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos972 with func( 9 in)(7 out) closure first 2 argument
func Clos972[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast972 with func( 9 in)(7 out) fix last 2 argument
func ClosLast972[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos973 with func( 9 in)(7 out) closure first 3 argument
func Clos973[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(D, E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast973 with func( 9 in)(7 out) fix last 3 argument
func ClosLast973[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E, F) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos974 with func( 9 in)(7 out) closure first 4 argument
func Clos974[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(E, F, G, H, I) (J, K, L, M, N, O, P) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast974 with func( 9 in)(7 out) fix last 4 argument
func ClosLast974[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D, E) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos975 with func( 9 in)(7 out) closure first 5 argument
func Clos975[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(F, G, H, I) (J, K, L, M, N, O, P) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast975 with func( 9 in)(7 out) fix last 5 argument
func ClosLast975[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C, D) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos976 with func( 9 in)(7 out) closure first 6 argument
func Clos976[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(G, H, I) (J, K, L, M, N, O, P) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast976 with func( 9 in)(7 out) fix last 6 argument
func ClosLast976[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B, C) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos977 with func( 9 in)(7 out) closure first 7 argument
func Clos977[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(H, I) (J, K, L, M, N, O, P) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast977 with func( 9 in)(7 out) fix last 7 argument
func ClosLast977[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A, B) (J, K, L, M, N, O, P) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos978 with func( 9 in)(7 out) closure first 8 argument
func Clos978[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(I) (J, K, L, M, N, O, P) {
	return func(i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast978 with func( 9 in)(7 out) fix last 8 argument
func ClosLast978[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(A) (J, K, L, M, N, O, P) {
	return func(a A) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos979 with func( 9 in)(7 out) closure first 9 argument
func Clos979[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func() (J, K, L, M, N, O, P) {
	return func() (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos981 with func( 9 in)(8 out) closure first 1 argument
func Clos981[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast981 with func( 9 in)(8 out) fix last 1 argument
func ClosLast981[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos982 with func( 9 in)(8 out) closure first 2 argument
func Clos982[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast982 with func( 9 in)(8 out) fix last 2 argument
func ClosLast982[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos983 with func( 9 in)(8 out) closure first 3 argument
func Clos983[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(D, E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast983 with func( 9 in)(8 out) fix last 3 argument
func ClosLast983[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E, F) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos984 with func( 9 in)(8 out) closure first 4 argument
func Clos984[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(E, F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast984 with func( 9 in)(8 out) fix last 4 argument
func ClosLast984[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D, E) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos985 with func( 9 in)(8 out) closure first 5 argument
func Clos985[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(F, G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast985 with func( 9 in)(8 out) fix last 5 argument
func ClosLast985[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C, D) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos986 with func( 9 in)(8 out) closure first 6 argument
func Clos986[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(G, H, I) (J, K, L, M, N, O, P, Q) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast986 with func( 9 in)(8 out) fix last 6 argument
func ClosLast986[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B, C) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos987 with func( 9 in)(8 out) closure first 7 argument
func Clos987[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(H, I) (J, K, L, M, N, O, P, Q) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast987 with func( 9 in)(8 out) fix last 7 argument
func ClosLast987[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A, B) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos988 with func( 9 in)(8 out) closure first 8 argument
func Clos988[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(I) (J, K, L, M, N, O, P, Q) {
	return func(i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast988 with func( 9 in)(8 out) fix last 8 argument
func ClosLast988[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(A) (J, K, L, M, N, O, P, Q) {
	return func(a A) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos989 with func( 9 in)(8 out) closure first 9 argument
func Clos989[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func() (J, K, L, M, N, O, P, Q) {
	return func() (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos991 with func( 9 in)(9 out) closure first 1 argument
func Clos991[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast991 with func( 9 in)(9 out) fix last 1 argument
func ClosLast991[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E, F, G, H) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos992 with func( 9 in)(9 out) closure first 2 argument
func Clos992[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast992 with func( 9 in)(9 out) fix last 2 argument
func ClosLast992[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E, F, G) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos993 with func( 9 in)(9 out) closure first 3 argument
func Clos993[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast993 with func( 9 in)(9 out) fix last 3 argument
func ClosLast993[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E, F) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos994 with func( 9 in)(9 out) closure first 4 argument
func Clos994[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(E, F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast994 with func( 9 in)(9 out) fix last 4 argument
func ClosLast994[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D, E) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos995 with func( 9 in)(9 out) closure first 5 argument
func Clos995[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(F, G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast995 with func( 9 in)(9 out) fix last 5 argument
func ClosLast995[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C, D) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos996 with func( 9 in)(9 out) closure first 6 argument
func Clos996[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(G, H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast996 with func( 9 in)(9 out) fix last 6 argument
func ClosLast996[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B, C) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos997 with func( 9 in)(9 out) closure first 7 argument
func Clos997[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(H, I) (J, K, L, M, N, O, P, Q, R) {
	return func(h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast997 with func( 9 in)(9 out) fix last 7 argument
func ClosLast997[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A, B) (J, K, L, M, N, O, P, Q, R) {
	return func(a A, b B) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos998 with func( 9 in)(9 out) closure first 8 argument
func Clos998[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, h H, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(I) (J, K, L, M, N, O, P, Q, R) {
	return func(i I) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// ClosLast998 with func( 9 in)(9 out) fix last 8 argument
func ClosLast998[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(A) (J, K, L, M, N, O, P, Q, R) {
	return func(a A) (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Clos999 with func( 9 in)(9 out) closure first 9 argument
func Clos999[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](a A, b B, c C, d D, e E, f F, g G, h H, i I, fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func() (J, K, L, M, N, O, P, Q, R) {
	return func() (j J, k K, l L, m M, n N, o O, p P, q Q, r R) {
		j, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}
