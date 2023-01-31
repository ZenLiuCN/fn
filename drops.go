package fn

//generated file should not edit

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

// Drop061 with func( 0 in)(6 out) drop first 1 result
func Drop061[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (B, C, D, E, F) {
	return func() (b B, c C, d D, e E, f F) {
		_, b, c, d, e, f = fn()
		return
	}
}

// DropLast061 with func( 0 in)(6 out) drop last 1 result
func DropLast061[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (A, B, C, D, E) {
	return func() (a A, b B, c C, d D, e E) {
		a, b, c, d, e, _ = fn()
		return
	}
}

// Drop062 with func( 0 in)(6 out) drop first 2 result
func Drop062[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (C, D, E, F) {
	return func() (c C, d D, e E, f F) {
		_, _, c, d, e, f = fn()
		return
	}
}

// DropLast062 with func( 0 in)(6 out) drop last 2 result
func DropLast062[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (A, B, C, D) {
	return func() (a A, b B, c C, d D) {
		a, b, c, d, _, _ = fn()
		return
	}
}

// Drop063 with func( 0 in)(6 out) drop first 3 result
func Drop063[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (D, E, F) {
	return func() (d D, e E, f F) {
		_, _, _, d, e, f = fn()
		return
	}
}

// DropLast063 with func( 0 in)(6 out) drop last 3 result
func DropLast063[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (A, B, C) {
	return func() (a A, b B, c C) {
		a, b, c, _, _, _ = fn()
		return
	}
}

// Drop064 with func( 0 in)(6 out) drop first 4 result
func Drop064[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (E, F) {
	return func() (e E, f F) {
		_, _, _, _, e, f = fn()
		return
	}
}

// DropLast064 with func( 0 in)(6 out) drop last 4 result
func DropLast064[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _, _, _, _ = fn()
		return
	}
}

// Drop065 with func( 0 in)(6 out) drop first 5 result
func Drop065[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() F {
	return func() (f F) {
		_, _, _, _, _, f = fn()
		return
	}
}

// DropLast065 with func( 0 in)(6 out) drop last 5 result
func DropLast065[A, B, C, D, E, F any](fn func() (A, B, C, D, E, F)) func() A {
	return func() (a A) {
		a, _, _, _, _, _ = fn()
		return
	}
}

// Drop071 with func( 0 in)(7 out) drop first 1 result
func Drop071[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (B, C, D, E, F, G) {
	return func() (b B, c C, d D, e E, f F, g G) {
		_, b, c, d, e, f, g = fn()
		return
	}
}

// DropLast071 with func( 0 in)(7 out) drop last 1 result
func DropLast071[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (A, B, C, D, E, F) {
	return func() (a A, b B, c C, d D, e E, f F) {
		a, b, c, d, e, f, _ = fn()
		return
	}
}

// Drop072 with func( 0 in)(7 out) drop first 2 result
func Drop072[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (C, D, E, F, G) {
	return func() (c C, d D, e E, f F, g G) {
		_, _, c, d, e, f, g = fn()
		return
	}
}

// DropLast072 with func( 0 in)(7 out) drop last 2 result
func DropLast072[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (A, B, C, D, E) {
	return func() (a A, b B, c C, d D, e E) {
		a, b, c, d, e, _, _ = fn()
		return
	}
}

// Drop073 with func( 0 in)(7 out) drop first 3 result
func Drop073[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (D, E, F, G) {
	return func() (d D, e E, f F, g G) {
		_, _, _, d, e, f, g = fn()
		return
	}
}

// DropLast073 with func( 0 in)(7 out) drop last 3 result
func DropLast073[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (A, B, C, D) {
	return func() (a A, b B, c C, d D) {
		a, b, c, d, _, _, _ = fn()
		return
	}
}

// Drop074 with func( 0 in)(7 out) drop first 4 result
func Drop074[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (E, F, G) {
	return func() (e E, f F, g G) {
		_, _, _, _, e, f, g = fn()
		return
	}
}

// DropLast074 with func( 0 in)(7 out) drop last 4 result
func DropLast074[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (A, B, C) {
	return func() (a A, b B, c C) {
		a, b, c, _, _, _, _ = fn()
		return
	}
}

// Drop075 with func( 0 in)(7 out) drop first 5 result
func Drop075[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (F, G) {
	return func() (f F, g G) {
		_, _, _, _, _, f, g = fn()
		return
	}
}

// DropLast075 with func( 0 in)(7 out) drop last 5 result
func DropLast075[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _, _, _, _, _ = fn()
		return
	}
}

// Drop076 with func( 0 in)(7 out) drop first 6 result
func Drop076[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() G {
	return func() (g G) {
		_, _, _, _, _, _, g = fn()
		return
	}
}

// DropLast076 with func( 0 in)(7 out) drop last 6 result
func DropLast076[A, B, C, D, E, F, G any](fn func() (A, B, C, D, E, F, G)) func() A {
	return func() (a A) {
		a, _, _, _, _, _, _ = fn()
		return
	}
}

// Drop081 with func( 0 in)(8 out) drop first 1 result
func Drop081[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (B, C, D, E, F, G, H) {
	return func() (b B, c C, d D, e E, f F, g G, h H) {
		_, b, c, d, e, f, g, h = fn()
		return
	}
}

// DropLast081 with func( 0 in)(8 out) drop last 1 result
func DropLast081[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B, C, D, E, F, G) {
	return func() (a A, b B, c C, d D, e E, f F, g G) {
		a, b, c, d, e, f, g, _ = fn()
		return
	}
}

// Drop082 with func( 0 in)(8 out) drop first 2 result
func Drop082[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (C, D, E, F, G, H) {
	return func() (c C, d D, e E, f F, g G, h H) {
		_, _, c, d, e, f, g, h = fn()
		return
	}
}

// DropLast082 with func( 0 in)(8 out) drop last 2 result
func DropLast082[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B, C, D, E, F) {
	return func() (a A, b B, c C, d D, e E, f F) {
		a, b, c, d, e, f, _, _ = fn()
		return
	}
}

// Drop083 with func( 0 in)(8 out) drop first 3 result
func Drop083[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (D, E, F, G, H) {
	return func() (d D, e E, f F, g G, h H) {
		_, _, _, d, e, f, g, h = fn()
		return
	}
}

// DropLast083 with func( 0 in)(8 out) drop last 3 result
func DropLast083[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B, C, D, E) {
	return func() (a A, b B, c C, d D, e E) {
		a, b, c, d, e, _, _, _ = fn()
		return
	}
}

// Drop084 with func( 0 in)(8 out) drop first 4 result
func Drop084[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (E, F, G, H) {
	return func() (e E, f F, g G, h H) {
		_, _, _, _, e, f, g, h = fn()
		return
	}
}

// DropLast084 with func( 0 in)(8 out) drop last 4 result
func DropLast084[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B, C, D) {
	return func() (a A, b B, c C, d D) {
		a, b, c, d, _, _, _, _ = fn()
		return
	}
}

// Drop085 with func( 0 in)(8 out) drop first 5 result
func Drop085[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (F, G, H) {
	return func() (f F, g G, h H) {
		_, _, _, _, _, f, g, h = fn()
		return
	}
}

// DropLast085 with func( 0 in)(8 out) drop last 5 result
func DropLast085[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B, C) {
	return func() (a A, b B, c C) {
		a, b, c, _, _, _, _, _ = fn()
		return
	}
}

// Drop086 with func( 0 in)(8 out) drop first 6 result
func Drop086[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (G, H) {
	return func() (g G, h H) {
		_, _, _, _, _, _, g, h = fn()
		return
	}
}

// DropLast086 with func( 0 in)(8 out) drop last 6 result
func DropLast086[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _, _, _, _, _, _ = fn()
		return
	}
}

// Drop087 with func( 0 in)(8 out) drop first 7 result
func Drop087[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() H {
	return func() (h H) {
		_, _, _, _, _, _, _, h = fn()
		return
	}
}

// DropLast087 with func( 0 in)(8 out) drop last 7 result
func DropLast087[A, B, C, D, E, F, G, H any](fn func() (A, B, C, D, E, F, G, H)) func() A {
	return func() (a A) {
		a, _, _, _, _, _, _, _ = fn()
		return
	}
}

// Drop091 with func( 0 in)(9 out) drop first 1 result
func Drop091[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (B, C, D, E, F, G, H, I) {
	return func() (b B, c C, d D, e E, f F, g G, h H, i I) {
		_, b, c, d, e, f, g, h, i = fn()
		return
	}
}

// DropLast091 with func( 0 in)(9 out) drop last 1 result
func DropLast091[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C, D, E, F, G, H) {
	return func() (a A, b B, c C, d D, e E, f F, g G, h H) {
		a, b, c, d, e, f, g, h, _ = fn()
		return
	}
}

// Drop092 with func( 0 in)(9 out) drop first 2 result
func Drop092[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (C, D, E, F, G, H, I) {
	return func() (c C, d D, e E, f F, g G, h H, i I) {
		_, _, c, d, e, f, g, h, i = fn()
		return
	}
}

// DropLast092 with func( 0 in)(9 out) drop last 2 result
func DropLast092[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C, D, E, F, G) {
	return func() (a A, b B, c C, d D, e E, f F, g G) {
		a, b, c, d, e, f, g, _, _ = fn()
		return
	}
}

// Drop093 with func( 0 in)(9 out) drop first 3 result
func Drop093[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (D, E, F, G, H, I) {
	return func() (d D, e E, f F, g G, h H, i I) {
		_, _, _, d, e, f, g, h, i = fn()
		return
	}
}

// DropLast093 with func( 0 in)(9 out) drop last 3 result
func DropLast093[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C, D, E, F) {
	return func() (a A, b B, c C, d D, e E, f F) {
		a, b, c, d, e, f, _, _, _ = fn()
		return
	}
}

// Drop094 with func( 0 in)(9 out) drop first 4 result
func Drop094[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (E, F, G, H, I) {
	return func() (e E, f F, g G, h H, i I) {
		_, _, _, _, e, f, g, h, i = fn()
		return
	}
}

// DropLast094 with func( 0 in)(9 out) drop last 4 result
func DropLast094[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C, D, E) {
	return func() (a A, b B, c C, d D, e E) {
		a, b, c, d, e, _, _, _, _ = fn()
		return
	}
}

// Drop095 with func( 0 in)(9 out) drop first 5 result
func Drop095[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (F, G, H, I) {
	return func() (f F, g G, h H, i I) {
		_, _, _, _, _, f, g, h, i = fn()
		return
	}
}

// DropLast095 with func( 0 in)(9 out) drop last 5 result
func DropLast095[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C, D) {
	return func() (a A, b B, c C, d D) {
		a, b, c, d, _, _, _, _, _ = fn()
		return
	}
}

// Drop096 with func( 0 in)(9 out) drop first 6 result
func Drop096[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (G, H, I) {
	return func() (g G, h H, i I) {
		_, _, _, _, _, _, g, h, i = fn()
		return
	}
}

// DropLast096 with func( 0 in)(9 out) drop last 6 result
func DropLast096[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B, C) {
	return func() (a A, b B, c C) {
		a, b, c, _, _, _, _, _, _ = fn()
		return
	}
}

// Drop097 with func( 0 in)(9 out) drop first 7 result
func Drop097[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (H, I) {
	return func() (h H, i I) {
		_, _, _, _, _, _, _, h, i = fn()
		return
	}
}

// DropLast097 with func( 0 in)(9 out) drop last 7 result
func DropLast097[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() (A, B) {
	return func() (a A, b B) {
		a, b, _, _, _, _, _, _, _ = fn()
		return
	}
}

// Drop098 with func( 0 in)(9 out) drop first 8 result
func Drop098[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() I {
	return func() (i I) {
		_, _, _, _, _, _, _, _, i = fn()
		return
	}
}

// DropLast098 with func( 0 in)(9 out) drop last 8 result
func DropLast098[A, B, C, D, E, F, G, H, I any](fn func() (A, B, C, D, E, F, G, H, I)) func() A {
	return func() (a A) {
		a, _, _, _, _, _, _, _, _ = fn()
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

// Drop161 with func( 1 in)(6 out) drop first 1 result
func Drop161[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (C, D, E, F, G) {
	return func(a A) (c C, d D, e E, f F, g G) {
		_, c, d, e, f, g = fn(a)
		return
	}
}

// DropLast161 with func( 1 in)(6 out) drop last 1 result
func DropLast161[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (B, C, D, E, F) {
	return func(a A) (b B, c C, d D, e E, f F) {
		b, c, d, e, f, _ = fn(a)
		return
	}
}

// Drop162 with func( 1 in)(6 out) drop first 2 result
func Drop162[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (D, E, F, G) {
	return func(a A) (d D, e E, f F, g G) {
		_, _, d, e, f, g = fn(a)
		return
	}
}

// DropLast162 with func( 1 in)(6 out) drop last 2 result
func DropLast162[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (B, C, D, E) {
	return func(a A) (b B, c C, d D, e E) {
		b, c, d, e, _, _ = fn(a)
		return
	}
}

// Drop163 with func( 1 in)(6 out) drop first 3 result
func Drop163[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (E, F, G) {
	return func(a A) (e E, f F, g G) {
		_, _, _, e, f, g = fn(a)
		return
	}
}

// DropLast163 with func( 1 in)(6 out) drop last 3 result
func DropLast163[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (B, C, D) {
	return func(a A) (b B, c C, d D) {
		b, c, d, _, _, _ = fn(a)
		return
	}
}

// Drop164 with func( 1 in)(6 out) drop first 4 result
func Drop164[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (F, G) {
	return func(a A) (f F, g G) {
		_, _, _, _, f, g = fn(a)
		return
	}
}

// DropLast164 with func( 1 in)(6 out) drop last 4 result
func DropLast164[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _, _, _, _ = fn(a)
		return
	}
}

// Drop165 with func( 1 in)(6 out) drop first 5 result
func Drop165[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) G {
	return func(a A) (g G) {
		_, _, _, _, _, g = fn(a)
		return
	}
}

// DropLast165 with func( 1 in)(6 out) drop last 5 result
func DropLast165[A, B, C, D, E, F, G any](fn func(A) (B, C, D, E, F, G)) func(a A) B {
	return func(a A) (b B) {
		b, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop171 with func( 1 in)(7 out) drop first 1 result
func Drop171[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (C, D, E, F, G, H) {
	return func(a A) (c C, d D, e E, f F, g G, h H) {
		_, c, d, e, f, g, h = fn(a)
		return
	}
}

// DropLast171 with func( 1 in)(7 out) drop last 1 result
func DropLast171[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (B, C, D, E, F, G) {
	return func(a A) (b B, c C, d D, e E, f F, g G) {
		b, c, d, e, f, g, _ = fn(a)
		return
	}
}

// Drop172 with func( 1 in)(7 out) drop first 2 result
func Drop172[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (D, E, F, G, H) {
	return func(a A) (d D, e E, f F, g G, h H) {
		_, _, d, e, f, g, h = fn(a)
		return
	}
}

// DropLast172 with func( 1 in)(7 out) drop last 2 result
func DropLast172[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (B, C, D, E, F) {
	return func(a A) (b B, c C, d D, e E, f F) {
		b, c, d, e, f, _, _ = fn(a)
		return
	}
}

// Drop173 with func( 1 in)(7 out) drop first 3 result
func Drop173[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (E, F, G, H) {
	return func(a A) (e E, f F, g G, h H) {
		_, _, _, e, f, g, h = fn(a)
		return
	}
}

// DropLast173 with func( 1 in)(7 out) drop last 3 result
func DropLast173[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (B, C, D, E) {
	return func(a A) (b B, c C, d D, e E) {
		b, c, d, e, _, _, _ = fn(a)
		return
	}
}

// Drop174 with func( 1 in)(7 out) drop first 4 result
func Drop174[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (F, G, H) {
	return func(a A) (f F, g G, h H) {
		_, _, _, _, f, g, h = fn(a)
		return
	}
}

// DropLast174 with func( 1 in)(7 out) drop last 4 result
func DropLast174[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (B, C, D) {
	return func(a A) (b B, c C, d D) {
		b, c, d, _, _, _, _ = fn(a)
		return
	}
}

// Drop175 with func( 1 in)(7 out) drop first 5 result
func Drop175[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (G, H) {
	return func(a A) (g G, h H) {
		_, _, _, _, _, g, h = fn(a)
		return
	}
}

// DropLast175 with func( 1 in)(7 out) drop last 5 result
func DropLast175[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop176 with func( 1 in)(7 out) drop first 6 result
func Drop176[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) H {
	return func(a A) (h H) {
		_, _, _, _, _, _, h = fn(a)
		return
	}
}

// DropLast176 with func( 1 in)(7 out) drop last 6 result
func DropLast176[A, B, C, D, E, F, G, H any](fn func(A) (B, C, D, E, F, G, H)) func(a A) B {
	return func(a A) (b B) {
		b, _, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop181 with func( 1 in)(8 out) drop first 1 result
func Drop181[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (C, D, E, F, G, H, I) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I) {
		_, c, d, e, f, g, h, i = fn(a)
		return
	}
}

// DropLast181 with func( 1 in)(8 out) drop last 1 result
func DropLast181[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C, D, E, F, G, H) {
	return func(a A) (b B, c C, d D, e E, f F, g G, h H) {
		b, c, d, e, f, g, h, _ = fn(a)
		return
	}
}

// Drop182 with func( 1 in)(8 out) drop first 2 result
func Drop182[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (D, E, F, G, H, I) {
	return func(a A) (d D, e E, f F, g G, h H, i I) {
		_, _, d, e, f, g, h, i = fn(a)
		return
	}
}

// DropLast182 with func( 1 in)(8 out) drop last 2 result
func DropLast182[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C, D, E, F, G) {
	return func(a A) (b B, c C, d D, e E, f F, g G) {
		b, c, d, e, f, g, _, _ = fn(a)
		return
	}
}

// Drop183 with func( 1 in)(8 out) drop first 3 result
func Drop183[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (E, F, G, H, I) {
	return func(a A) (e E, f F, g G, h H, i I) {
		_, _, _, e, f, g, h, i = fn(a)
		return
	}
}

// DropLast183 with func( 1 in)(8 out) drop last 3 result
func DropLast183[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C, D, E, F) {
	return func(a A) (b B, c C, d D, e E, f F) {
		b, c, d, e, f, _, _, _ = fn(a)
		return
	}
}

// Drop184 with func( 1 in)(8 out) drop first 4 result
func Drop184[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (F, G, H, I) {
	return func(a A) (f F, g G, h H, i I) {
		_, _, _, _, f, g, h, i = fn(a)
		return
	}
}

// DropLast184 with func( 1 in)(8 out) drop last 4 result
func DropLast184[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C, D, E) {
	return func(a A) (b B, c C, d D, e E) {
		b, c, d, e, _, _, _, _ = fn(a)
		return
	}
}

// Drop185 with func( 1 in)(8 out) drop first 5 result
func Drop185[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (G, H, I) {
	return func(a A) (g G, h H, i I) {
		_, _, _, _, _, g, h, i = fn(a)
		return
	}
}

// DropLast185 with func( 1 in)(8 out) drop last 5 result
func DropLast185[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C, D) {
	return func(a A) (b B, c C, d D) {
		b, c, d, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop186 with func( 1 in)(8 out) drop first 6 result
func Drop186[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (H, I) {
	return func(a A) (h H, i I) {
		_, _, _, _, _, _, h, i = fn(a)
		return
	}
}

// DropLast186 with func( 1 in)(8 out) drop last 6 result
func DropLast186[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop187 with func( 1 in)(8 out) drop first 7 result
func Drop187[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) I {
	return func(a A) (i I) {
		_, _, _, _, _, _, _, i = fn(a)
		return
	}
}

// DropLast187 with func( 1 in)(8 out) drop last 7 result
func DropLast187[A, B, C, D, E, F, G, H, I any](fn func(A) (B, C, D, E, F, G, H, I)) func(a A) B {
	return func(a A) (b B) {
		b, _, _, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop191 with func( 1 in)(9 out) drop first 1 result
func Drop191[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (C, D, E, F, G, H, I, J) {
	return func(a A) (c C, d D, e E, f F, g G, h H, i I, j J) {
		_, c, d, e, f, g, h, i, j = fn(a)
		return
	}
}

// DropLast191 with func( 1 in)(9 out) drop last 1 result
func DropLast191[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D, E, F, G, H, I) {
	return func(a A) (b B, c C, d D, e E, f F, g G, h H, i I) {
		b, c, d, e, f, g, h, i, _ = fn(a)
		return
	}
}

// Drop192 with func( 1 in)(9 out) drop first 2 result
func Drop192[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (D, E, F, G, H, I, J) {
	return func(a A) (d D, e E, f F, g G, h H, i I, j J) {
		_, _, d, e, f, g, h, i, j = fn(a)
		return
	}
}

// DropLast192 with func( 1 in)(9 out) drop last 2 result
func DropLast192[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D, E, F, G, H) {
	return func(a A) (b B, c C, d D, e E, f F, g G, h H) {
		b, c, d, e, f, g, h, _, _ = fn(a)
		return
	}
}

// Drop193 with func( 1 in)(9 out) drop first 3 result
func Drop193[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (E, F, G, H, I, J) {
	return func(a A) (e E, f F, g G, h H, i I, j J) {
		_, _, _, e, f, g, h, i, j = fn(a)
		return
	}
}

// DropLast193 with func( 1 in)(9 out) drop last 3 result
func DropLast193[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D, E, F, G) {
	return func(a A) (b B, c C, d D, e E, f F, g G) {
		b, c, d, e, f, g, _, _, _ = fn(a)
		return
	}
}

// Drop194 with func( 1 in)(9 out) drop first 4 result
func Drop194[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (F, G, H, I, J) {
	return func(a A) (f F, g G, h H, i I, j J) {
		_, _, _, _, f, g, h, i, j = fn(a)
		return
	}
}

// DropLast194 with func( 1 in)(9 out) drop last 4 result
func DropLast194[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D, E, F) {
	return func(a A) (b B, c C, d D, e E, f F) {
		b, c, d, e, f, _, _, _, _ = fn(a)
		return
	}
}

// Drop195 with func( 1 in)(9 out) drop first 5 result
func Drop195[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (G, H, I, J) {
	return func(a A) (g G, h H, i I, j J) {
		_, _, _, _, _, g, h, i, j = fn(a)
		return
	}
}

// DropLast195 with func( 1 in)(9 out) drop last 5 result
func DropLast195[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D, E) {
	return func(a A) (b B, c C, d D, e E) {
		b, c, d, e, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop196 with func( 1 in)(9 out) drop first 6 result
func Drop196[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (H, I, J) {
	return func(a A) (h H, i I, j J) {
		_, _, _, _, _, _, h, i, j = fn(a)
		return
	}
}

// DropLast196 with func( 1 in)(9 out) drop last 6 result
func DropLast196[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C, D) {
	return func(a A) (b B, c C, d D) {
		b, c, d, _, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop197 with func( 1 in)(9 out) drop first 7 result
func Drop197[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (I, J) {
	return func(a A) (i I, j J) {
		_, _, _, _, _, _, _, i, j = fn(a)
		return
	}
}

// DropLast197 with func( 1 in)(9 out) drop last 7 result
func DropLast197[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) (B, C) {
	return func(a A) (b B, c C) {
		b, c, _, _, _, _, _, _, _ = fn(a)
		return
	}
}

// Drop198 with func( 1 in)(9 out) drop first 8 result
func Drop198[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) J {
	return func(a A) (j J) {
		_, _, _, _, _, _, _, _, j = fn(a)
		return
	}
}

// DropLast198 with func( 1 in)(9 out) drop last 8 result
func DropLast198[A, B, C, D, E, F, G, H, I, J any](fn func(A) (B, C, D, E, F, G, H, I, J)) func(a A) B {
	return func(a A) (b B) {
		b, _, _, _, _, _, _, _, _ = fn(a)
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

// Drop261 with func( 2 in)(6 out) drop first 1 result
func Drop261[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (D, E, F, G, H) {
	return func(a A, b B) (d D, e E, f F, g G, h H) {
		_, d, e, f, g, h = fn(a, b)
		return
	}
}

// DropLast261 with func( 2 in)(6 out) drop last 1 result
func DropLast261[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (C, D, E, F, G) {
	return func(a A, b B) (c C, d D, e E, f F, g G) {
		c, d, e, f, g, _ = fn(a, b)
		return
	}
}

// Drop262 with func( 2 in)(6 out) drop first 2 result
func Drop262[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (E, F, G, H) {
	return func(a A, b B) (e E, f F, g G, h H) {
		_, _, e, f, g, h = fn(a, b)
		return
	}
}

// DropLast262 with func( 2 in)(6 out) drop last 2 result
func DropLast262[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (C, D, E, F) {
	return func(a A, b B) (c C, d D, e E, f F) {
		c, d, e, f, _, _ = fn(a, b)
		return
	}
}

// Drop263 with func( 2 in)(6 out) drop first 3 result
func Drop263[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (F, G, H) {
	return func(a A, b B) (f F, g G, h H) {
		_, _, _, f, g, h = fn(a, b)
		return
	}
}

// DropLast263 with func( 2 in)(6 out) drop last 3 result
func DropLast263[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (C, D, E) {
	return func(a A, b B) (c C, d D, e E) {
		c, d, e, _, _, _ = fn(a, b)
		return
	}
}

// Drop264 with func( 2 in)(6 out) drop first 4 result
func Drop264[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (G, H) {
	return func(a A, b B) (g G, h H) {
		_, _, _, _, g, h = fn(a, b)
		return
	}
}

// DropLast264 with func( 2 in)(6 out) drop last 4 result
func DropLast264[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop265 with func( 2 in)(6 out) drop first 5 result
func Drop265[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) H {
	return func(a A, b B) (h H) {
		_, _, _, _, _, h = fn(a, b)
		return
	}
}

// DropLast265 with func( 2 in)(6 out) drop last 5 result
func DropLast265[A, B, C, D, E, F, G, H any](fn func(A, B) (C, D, E, F, G, H)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop271 with func( 2 in)(7 out) drop first 1 result
func Drop271[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (D, E, F, G, H, I) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I) {
		_, d, e, f, g, h, i = fn(a, b)
		return
	}
}

// DropLast271 with func( 2 in)(7 out) drop last 1 result
func DropLast271[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (C, D, E, F, G, H) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h, _ = fn(a, b)
		return
	}
}

// Drop272 with func( 2 in)(7 out) drop first 2 result
func Drop272[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (E, F, G, H, I) {
	return func(a A, b B) (e E, f F, g G, h H, i I) {
		_, _, e, f, g, h, i = fn(a, b)
		return
	}
}

// DropLast272 with func( 2 in)(7 out) drop last 2 result
func DropLast272[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (C, D, E, F, G) {
	return func(a A, b B) (c C, d D, e E, f F, g G) {
		c, d, e, f, g, _, _ = fn(a, b)
		return
	}
}

// Drop273 with func( 2 in)(7 out) drop first 3 result
func Drop273[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (F, G, H, I) {
	return func(a A, b B) (f F, g G, h H, i I) {
		_, _, _, f, g, h, i = fn(a, b)
		return
	}
}

// DropLast273 with func( 2 in)(7 out) drop last 3 result
func DropLast273[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (C, D, E, F) {
	return func(a A, b B) (c C, d D, e E, f F) {
		c, d, e, f, _, _, _ = fn(a, b)
		return
	}
}

// Drop274 with func( 2 in)(7 out) drop first 4 result
func Drop274[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (G, H, I) {
	return func(a A, b B) (g G, h H, i I) {
		_, _, _, _, g, h, i = fn(a, b)
		return
	}
}

// DropLast274 with func( 2 in)(7 out) drop last 4 result
func DropLast274[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (C, D, E) {
	return func(a A, b B) (c C, d D, e E) {
		c, d, e, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop275 with func( 2 in)(7 out) drop first 5 result
func Drop275[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (H, I) {
	return func(a A, b B) (h H, i I) {
		_, _, _, _, _, h, i = fn(a, b)
		return
	}
}

// DropLast275 with func( 2 in)(7 out) drop last 5 result
func DropLast275[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop276 with func( 2 in)(7 out) drop first 6 result
func Drop276[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) I {
	return func(a A, b B) (i I) {
		_, _, _, _, _, _, i = fn(a, b)
		return
	}
}

// DropLast276 with func( 2 in)(7 out) drop last 6 result
func DropLast276[A, B, C, D, E, F, G, H, I any](fn func(A, B) (C, D, E, F, G, H, I)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop281 with func( 2 in)(8 out) drop first 1 result
func Drop281[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (D, E, F, G, H, I, J) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J) {
		_, d, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// DropLast281 with func( 2 in)(8 out) drop last 1 result
func DropLast281[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D, E, F, G, H, I) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i, _ = fn(a, b)
		return
	}
}

// Drop282 with func( 2 in)(8 out) drop first 2 result
func Drop282[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (E, F, G, H, I, J) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J) {
		_, _, e, f, g, h, i, j = fn(a, b)
		return
	}
}

// DropLast282 with func( 2 in)(8 out) drop last 2 result
func DropLast282[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D, E, F, G, H) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h, _, _ = fn(a, b)
		return
	}
}

// Drop283 with func( 2 in)(8 out) drop first 3 result
func Drop283[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (F, G, H, I, J) {
	return func(a A, b B) (f F, g G, h H, i I, j J) {
		_, _, _, f, g, h, i, j = fn(a, b)
		return
	}
}

// DropLast283 with func( 2 in)(8 out) drop last 3 result
func DropLast283[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D, E, F, G) {
	return func(a A, b B) (c C, d D, e E, f F, g G) {
		c, d, e, f, g, _, _, _ = fn(a, b)
		return
	}
}

// Drop284 with func( 2 in)(8 out) drop first 4 result
func Drop284[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (G, H, I, J) {
	return func(a A, b B) (g G, h H, i I, j J) {
		_, _, _, _, g, h, i, j = fn(a, b)
		return
	}
}

// DropLast284 with func( 2 in)(8 out) drop last 4 result
func DropLast284[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D, E, F) {
	return func(a A, b B) (c C, d D, e E, f F) {
		c, d, e, f, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop285 with func( 2 in)(8 out) drop first 5 result
func Drop285[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (H, I, J) {
	return func(a A, b B) (h H, i I, j J) {
		_, _, _, _, _, h, i, j = fn(a, b)
		return
	}
}

// DropLast285 with func( 2 in)(8 out) drop last 5 result
func DropLast285[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D, E) {
	return func(a A, b B) (c C, d D, e E) {
		c, d, e, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop286 with func( 2 in)(8 out) drop first 6 result
func Drop286[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (I, J) {
	return func(a A, b B) (i I, j J) {
		_, _, _, _, _, _, i, j = fn(a, b)
		return
	}
}

// DropLast286 with func( 2 in)(8 out) drop last 6 result
func DropLast286[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop287 with func( 2 in)(8 out) drop first 7 result
func Drop287[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) J {
	return func(a A, b B) (j J) {
		_, _, _, _, _, _, _, j = fn(a, b)
		return
	}
}

// DropLast287 with func( 2 in)(8 out) drop last 7 result
func DropLast287[A, B, C, D, E, F, G, H, I, J any](fn func(A, B) (C, D, E, F, G, H, I, J)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop291 with func( 2 in)(9 out) drop first 1 result
func Drop291[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (D, E, F, G, H, I, J, K) {
	return func(a A, b B) (d D, e E, f F, g G, h H, i I, j J, k K) {
		_, d, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// DropLast291 with func( 2 in)(9 out) drop last 1 result
func DropLast291[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E, F, G, H, I, J) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, i I, j J) {
		c, d, e, f, g, h, i, j, _ = fn(a, b)
		return
	}
}

// Drop292 with func( 2 in)(9 out) drop first 2 result
func Drop292[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (E, F, G, H, I, J, K) {
	return func(a A, b B) (e E, f F, g G, h H, i I, j J, k K) {
		_, _, e, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// DropLast292 with func( 2 in)(9 out) drop last 2 result
func DropLast292[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E, F, G, H, I) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H, i I) {
		c, d, e, f, g, h, i, _, _ = fn(a, b)
		return
	}
}

// Drop293 with func( 2 in)(9 out) drop first 3 result
func Drop293[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (F, G, H, I, J, K) {
	return func(a A, b B) (f F, g G, h H, i I, j J, k K) {
		_, _, _, f, g, h, i, j, k = fn(a, b)
		return
	}
}

// DropLast293 with func( 2 in)(9 out) drop last 3 result
func DropLast293[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E, F, G, H) {
	return func(a A, b B) (c C, d D, e E, f F, g G, h H) {
		c, d, e, f, g, h, _, _, _ = fn(a, b)
		return
	}
}

// Drop294 with func( 2 in)(9 out) drop first 4 result
func Drop294[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (G, H, I, J, K) {
	return func(a A, b B) (g G, h H, i I, j J, k K) {
		_, _, _, _, g, h, i, j, k = fn(a, b)
		return
	}
}

// DropLast294 with func( 2 in)(9 out) drop last 4 result
func DropLast294[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E, F, G) {
	return func(a A, b B) (c C, d D, e E, f F, g G) {
		c, d, e, f, g, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop295 with func( 2 in)(9 out) drop first 5 result
func Drop295[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (H, I, J, K) {
	return func(a A, b B) (h H, i I, j J, k K) {
		_, _, _, _, _, h, i, j, k = fn(a, b)
		return
	}
}

// DropLast295 with func( 2 in)(9 out) drop last 5 result
func DropLast295[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E, F) {
	return func(a A, b B) (c C, d D, e E, f F) {
		c, d, e, f, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop296 with func( 2 in)(9 out) drop first 6 result
func Drop296[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (I, J, K) {
	return func(a A, b B) (i I, j J, k K) {
		_, _, _, _, _, _, i, j, k = fn(a, b)
		return
	}
}

// DropLast296 with func( 2 in)(9 out) drop last 6 result
func DropLast296[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D, E) {
	return func(a A, b B) (c C, d D, e E) {
		c, d, e, _, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop297 with func( 2 in)(9 out) drop first 7 result
func Drop297[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (J, K) {
	return func(a A, b B) (j J, k K) {
		_, _, _, _, _, _, _, j, k = fn(a, b)
		return
	}
}

// DropLast297 with func( 2 in)(9 out) drop last 7 result
func DropLast297[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) (C, D) {
	return func(a A, b B) (c C, d D) {
		c, d, _, _, _, _, _, _, _ = fn(a, b)
		return
	}
}

// Drop298 with func( 2 in)(9 out) drop first 8 result
func Drop298[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) K {
	return func(a A, b B) (k K) {
		_, _, _, _, _, _, _, _, k = fn(a, b)
		return
	}
}

// DropLast298 with func( 2 in)(9 out) drop last 8 result
func DropLast298[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B) (C, D, E, F, G, H, I, J, K)) func(a A, b B) C {
	return func(a A, b B) (c C) {
		c, _, _, _, _, _, _, _, _ = fn(a, b)
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

// Drop361 with func( 3 in)(6 out) drop first 1 result
func Drop361[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (E, F, G, H, I) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I) {
		_, e, f, g, h, i = fn(a, b, c)
		return
	}
}

// DropLast361 with func( 3 in)(6 out) drop last 1 result
func DropLast361[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (D, E, F, G, H) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h, _ = fn(a, b, c)
		return
	}
}

// Drop362 with func( 3 in)(6 out) drop first 2 result
func Drop362[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (F, G, H, I) {
	return func(a A, b B, c C) (f F, g G, h H, i I) {
		_, _, f, g, h, i = fn(a, b, c)
		return
	}
}

// DropLast362 with func( 3 in)(6 out) drop last 2 result
func DropLast362[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (D, E, F, G) {
	return func(a A, b B, c C) (d D, e E, f F, g G) {
		d, e, f, g, _, _ = fn(a, b, c)
		return
	}
}

// Drop363 with func( 3 in)(6 out) drop first 3 result
func Drop363[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (G, H, I) {
	return func(a A, b B, c C) (g G, h H, i I) {
		_, _, _, g, h, i = fn(a, b, c)
		return
	}
}

// DropLast363 with func( 3 in)(6 out) drop last 3 result
func DropLast363[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (D, E, F) {
	return func(a A, b B, c C) (d D, e E, f F) {
		d, e, f, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop364 with func( 3 in)(6 out) drop first 4 result
func Drop364[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (H, I) {
	return func(a A, b B, c C) (h H, i I) {
		_, _, _, _, h, i = fn(a, b, c)
		return
	}
}

// DropLast364 with func( 3 in)(6 out) drop last 4 result
func DropLast364[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop365 with func( 3 in)(6 out) drop first 5 result
func Drop365[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) I {
	return func(a A, b B, c C) (i I) {
		_, _, _, _, _, i = fn(a, b, c)
		return
	}
}

// DropLast365 with func( 3 in)(6 out) drop last 5 result
func DropLast365[A, B, C, D, E, F, G, H, I any](fn func(A, B, C) (D, E, F, G, H, I)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop371 with func( 3 in)(7 out) drop first 1 result
func Drop371[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (E, F, G, H, I, J) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J) {
		_, e, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// DropLast371 with func( 3 in)(7 out) drop last 1 result
func DropLast371[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (D, E, F, G, H, I) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i, _ = fn(a, b, c)
		return
	}
}

// Drop372 with func( 3 in)(7 out) drop first 2 result
func Drop372[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (F, G, H, I, J) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J) {
		_, _, f, g, h, i, j = fn(a, b, c)
		return
	}
}

// DropLast372 with func( 3 in)(7 out) drop last 2 result
func DropLast372[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (D, E, F, G, H) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h, _, _ = fn(a, b, c)
		return
	}
}

// Drop373 with func( 3 in)(7 out) drop first 3 result
func Drop373[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (G, H, I, J) {
	return func(a A, b B, c C) (g G, h H, i I, j J) {
		_, _, _, g, h, i, j = fn(a, b, c)
		return
	}
}

// DropLast373 with func( 3 in)(7 out) drop last 3 result
func DropLast373[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (D, E, F, G) {
	return func(a A, b B, c C) (d D, e E, f F, g G) {
		d, e, f, g, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop374 with func( 3 in)(7 out) drop first 4 result
func Drop374[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (H, I, J) {
	return func(a A, b B, c C) (h H, i I, j J) {
		_, _, _, _, h, i, j = fn(a, b, c)
		return
	}
}

// DropLast374 with func( 3 in)(7 out) drop last 4 result
func DropLast374[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (D, E, F) {
	return func(a A, b B, c C) (d D, e E, f F) {
		d, e, f, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop375 with func( 3 in)(7 out) drop first 5 result
func Drop375[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (I, J) {
	return func(a A, b B, c C) (i I, j J) {
		_, _, _, _, _, i, j = fn(a, b, c)
		return
	}
}

// DropLast375 with func( 3 in)(7 out) drop last 5 result
func DropLast375[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop376 with func( 3 in)(7 out) drop first 6 result
func Drop376[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) J {
	return func(a A, b B, c C) (j J) {
		_, _, _, _, _, _, j = fn(a, b, c)
		return
	}
}

// DropLast376 with func( 3 in)(7 out) drop last 6 result
func DropLast376[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C) (D, E, F, G, H, I, J)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop381 with func( 3 in)(8 out) drop first 1 result
func Drop381[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (E, F, G, H, I, J, K) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K) {
		_, e, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// DropLast381 with func( 3 in)(8 out) drop last 1 result
func DropLast381[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E, F, G, H, I, J) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j, _ = fn(a, b, c)
		return
	}
}

// Drop382 with func( 3 in)(8 out) drop first 2 result
func Drop382[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (F, G, H, I, J, K) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K) {
		_, _, f, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// DropLast382 with func( 3 in)(8 out) drop last 2 result
func DropLast382[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E, F, G, H, I) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i, _, _ = fn(a, b, c)
		return
	}
}

// Drop383 with func( 3 in)(8 out) drop first 3 result
func Drop383[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (G, H, I, J, K) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K) {
		_, _, _, g, h, i, j, k = fn(a, b, c)
		return
	}
}

// DropLast383 with func( 3 in)(8 out) drop last 3 result
func DropLast383[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E, F, G, H) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop384 with func( 3 in)(8 out) drop first 4 result
func Drop384[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (H, I, J, K) {
	return func(a A, b B, c C) (h H, i I, j J, k K) {
		_, _, _, _, h, i, j, k = fn(a, b, c)
		return
	}
}

// DropLast384 with func( 3 in)(8 out) drop last 4 result
func DropLast384[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E, F, G) {
	return func(a A, b B, c C) (d D, e E, f F, g G) {
		d, e, f, g, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop385 with func( 3 in)(8 out) drop first 5 result
func Drop385[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (I, J, K) {
	return func(a A, b B, c C) (i I, j J, k K) {
		_, _, _, _, _, i, j, k = fn(a, b, c)
		return
	}
}

// DropLast385 with func( 3 in)(8 out) drop last 5 result
func DropLast385[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E, F) {
	return func(a A, b B, c C) (d D, e E, f F) {
		d, e, f, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop386 with func( 3 in)(8 out) drop first 6 result
func Drop386[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (J, K) {
	return func(a A, b B, c C) (j J, k K) {
		_, _, _, _, _, _, j, k = fn(a, b, c)
		return
	}
}

// DropLast386 with func( 3 in)(8 out) drop last 6 result
func DropLast386[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop387 with func( 3 in)(8 out) drop first 7 result
func Drop387[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) K {
	return func(a A, b B, c C) (k K) {
		_, _, _, _, _, _, _, k = fn(a, b, c)
		return
	}
}

// DropLast387 with func( 3 in)(8 out) drop last 7 result
func DropLast387[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C) (D, E, F, G, H, I, J, K)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop391 with func( 3 in)(9 out) drop first 1 result
func Drop391[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (E, F, G, H, I, J, K, L) {
	return func(a A, b B, c C) (e E, f F, g G, h H, i I, j J, k K, l L) {
		_, e, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// DropLast391 with func( 3 in)(9 out) drop last 1 result
func DropLast391[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F, G, H, I, J, K) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, j J, k K) {
		d, e, f, g, h, i, j, k, _ = fn(a, b, c)
		return
	}
}

// Drop392 with func( 3 in)(9 out) drop first 2 result
func Drop392[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C) (f F, g G, h H, i I, j J, k K, l L) {
		_, _, f, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// DropLast392 with func( 3 in)(9 out) drop last 2 result
func DropLast392[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F, G, H, I, J) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I, j J) {
		d, e, f, g, h, i, j, _, _ = fn(a, b, c)
		return
	}
}

// Drop393 with func( 3 in)(9 out) drop first 3 result
func Drop393[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (G, H, I, J, K, L) {
	return func(a A, b B, c C) (g G, h H, i I, j J, k K, l L) {
		_, _, _, g, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// DropLast393 with func( 3 in)(9 out) drop last 3 result
func DropLast393[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F, G, H, I) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H, i I) {
		d, e, f, g, h, i, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop394 with func( 3 in)(9 out) drop first 4 result
func Drop394[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (H, I, J, K, L) {
	return func(a A, b B, c C) (h H, i I, j J, k K, l L) {
		_, _, _, _, h, i, j, k, l = fn(a, b, c)
		return
	}
}

// DropLast394 with func( 3 in)(9 out) drop last 4 result
func DropLast394[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F, G, H) {
	return func(a A, b B, c C) (d D, e E, f F, g G, h H) {
		d, e, f, g, h, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop395 with func( 3 in)(9 out) drop first 5 result
func Drop395[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (I, J, K, L) {
	return func(a A, b B, c C) (i I, j J, k K, l L) {
		_, _, _, _, _, i, j, k, l = fn(a, b, c)
		return
	}
}

// DropLast395 with func( 3 in)(9 out) drop last 5 result
func DropLast395[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F, G) {
	return func(a A, b B, c C) (d D, e E, f F, g G) {
		d, e, f, g, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop396 with func( 3 in)(9 out) drop first 6 result
func Drop396[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (J, K, L) {
	return func(a A, b B, c C) (j J, k K, l L) {
		_, _, _, _, _, _, j, k, l = fn(a, b, c)
		return
	}
}

// DropLast396 with func( 3 in)(9 out) drop last 6 result
func DropLast396[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E, F) {
	return func(a A, b B, c C) (d D, e E, f F) {
		d, e, f, _, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop397 with func( 3 in)(9 out) drop first 7 result
func Drop397[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (K, L) {
	return func(a A, b B, c C) (k K, l L) {
		_, _, _, _, _, _, _, k, l = fn(a, b, c)
		return
	}
}

// DropLast397 with func( 3 in)(9 out) drop last 7 result
func DropLast397[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) (D, E) {
	return func(a A, b B, c C) (d D, e E) {
		d, e, _, _, _, _, _, _, _ = fn(a, b, c)
		return
	}
}

// Drop398 with func( 3 in)(9 out) drop first 8 result
func Drop398[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) L {
	return func(a A, b B, c C) (l L) {
		_, _, _, _, _, _, _, _, l = fn(a, b, c)
		return
	}
}

// DropLast398 with func( 3 in)(9 out) drop last 8 result
func DropLast398[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C) (D, E, F, G, H, I, J, K, L)) func(a A, b B, c C) D {
	return func(a A, b B, c C) (d D) {
		d, _, _, _, _, _, _, _, _ = fn(a, b, c)
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

// Drop461 with func( 4 in)(6 out) drop first 1 result
func Drop461[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (F, G, H, I, J) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J) {
		_, f, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// DropLast461 with func( 4 in)(6 out) drop last 1 result
func DropLast461[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (E, F, G, H, I) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i, _ = fn(a, b, c, d)
		return
	}
}

// Drop462 with func( 4 in)(6 out) drop first 2 result
func Drop462[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (G, H, I, J) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J) {
		_, _, g, h, i, j = fn(a, b, c, d)
		return
	}
}

// DropLast462 with func( 4 in)(6 out) drop last 2 result
func DropLast462[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (E, F, G, H) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop463 with func( 4 in)(6 out) drop first 3 result
func Drop463[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (H, I, J) {
	return func(a A, b B, c C, d D) (h H, i I, j J) {
		_, _, _, h, i, j = fn(a, b, c, d)
		return
	}
}

// DropLast463 with func( 4 in)(6 out) drop last 3 result
func DropLast463[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (E, F, G) {
	return func(a A, b B, c C, d D) (e E, f F, g G) {
		e, f, g, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop464 with func( 4 in)(6 out) drop first 4 result
func Drop464[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (I, J) {
	return func(a A, b B, c C, d D) (i I, j J) {
		_, _, _, _, i, j = fn(a, b, c, d)
		return
	}
}

// DropLast464 with func( 4 in)(6 out) drop last 4 result
func DropLast464[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop465 with func( 4 in)(6 out) drop first 5 result
func Drop465[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) J {
	return func(a A, b B, c C, d D) (j J) {
		_, _, _, _, _, j = fn(a, b, c, d)
		return
	}
}

// DropLast465 with func( 4 in)(6 out) drop last 5 result
func DropLast465[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D) (E, F, G, H, I, J)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop471 with func( 4 in)(7 out) drop first 1 result
func Drop471[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (F, G, H, I, J, K) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K) {
		_, f, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// DropLast471 with func( 4 in)(7 out) drop last 1 result
func DropLast471[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (E, F, G, H, I, J) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j, _ = fn(a, b, c, d)
		return
	}
}

// Drop472 with func( 4 in)(7 out) drop first 2 result
func Drop472[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (G, H, I, J, K) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K) {
		_, _, g, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// DropLast472 with func( 4 in)(7 out) drop last 2 result
func DropLast472[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (E, F, G, H, I) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop473 with func( 4 in)(7 out) drop first 3 result
func Drop473[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (H, I, J, K) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K) {
		_, _, _, h, i, j, k = fn(a, b, c, d)
		return
	}
}

// DropLast473 with func( 4 in)(7 out) drop last 3 result
func DropLast473[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (E, F, G, H) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop474 with func( 4 in)(7 out) drop first 4 result
func Drop474[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (I, J, K) {
	return func(a A, b B, c C, d D) (i I, j J, k K) {
		_, _, _, _, i, j, k = fn(a, b, c, d)
		return
	}
}

// DropLast474 with func( 4 in)(7 out) drop last 4 result
func DropLast474[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (E, F, G) {
	return func(a A, b B, c C, d D) (e E, f F, g G) {
		e, f, g, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop475 with func( 4 in)(7 out) drop first 5 result
func Drop475[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (J, K) {
	return func(a A, b B, c C, d D) (j J, k K) {
		_, _, _, _, _, j, k = fn(a, b, c, d)
		return
	}
}

// DropLast475 with func( 4 in)(7 out) drop last 5 result
func DropLast475[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop476 with func( 4 in)(7 out) drop first 6 result
func Drop476[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) K {
	return func(a A, b B, c C, d D) (k K) {
		_, _, _, _, _, _, k = fn(a, b, c, d)
		return
	}
}

// DropLast476 with func( 4 in)(7 out) drop last 6 result
func DropLast476[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D) (E, F, G, H, I, J, K)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop481 with func( 4 in)(8 out) drop first 1 result
func Drop481[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L) {
		_, f, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// DropLast481 with func( 4 in)(8 out) drop last 1 result
func DropLast481[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F, G, H, I, J, K) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k, _ = fn(a, b, c, d)
		return
	}
}

// Drop482 with func( 4 in)(8 out) drop first 2 result
func Drop482[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L) {
		_, _, g, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// DropLast482 with func( 4 in)(8 out) drop last 2 result
func DropLast482[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F, G, H, I, J) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop483 with func( 4 in)(8 out) drop first 3 result
func Drop483[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (H, I, J, K, L) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L) {
		_, _, _, h, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// DropLast483 with func( 4 in)(8 out) drop last 3 result
func DropLast483[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F, G, H, I) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop484 with func( 4 in)(8 out) drop first 4 result
func Drop484[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (I, J, K, L) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L) {
		_, _, _, _, i, j, k, l = fn(a, b, c, d)
		return
	}
}

// DropLast484 with func( 4 in)(8 out) drop last 4 result
func DropLast484[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F, G, H) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop485 with func( 4 in)(8 out) drop first 5 result
func Drop485[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (J, K, L) {
	return func(a A, b B, c C, d D) (j J, k K, l L) {
		_, _, _, _, _, j, k, l = fn(a, b, c, d)
		return
	}
}

// DropLast485 with func( 4 in)(8 out) drop last 5 result
func DropLast485[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F, G) {
	return func(a A, b B, c C, d D) (e E, f F, g G) {
		e, f, g, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop486 with func( 4 in)(8 out) drop first 6 result
func Drop486[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (K, L) {
	return func(a A, b B, c C, d D) (k K, l L) {
		_, _, _, _, _, _, k, l = fn(a, b, c, d)
		return
	}
}

// DropLast486 with func( 4 in)(8 out) drop last 6 result
func DropLast486[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop487 with func( 4 in)(8 out) drop first 7 result
func Drop487[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) L {
	return func(a A, b B, c C, d D) (l L) {
		_, _, _, _, _, _, _, l = fn(a, b, c, d)
		return
	}
}

// DropLast487 with func( 4 in)(8 out) drop last 7 result
func DropLast487[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop491 with func( 4 in)(9 out) drop first 1 result
func Drop491[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (f F, g G, h H, i I, j J, k K, l L, m M) {
		_, f, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast491 with func( 4 in)(9 out) drop last 1 result
func DropLast491[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G, H, I, J, K, L) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K, l L) {
		e, f, g, h, i, j, k, l, _ = fn(a, b, c, d)
		return
	}
}

// Drop492 with func( 4 in)(9 out) drop first 2 result
func Drop492[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (g G, h H, i I, j J, k K, l L, m M) {
		_, _, g, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast492 with func( 4 in)(9 out) drop last 2 result
func DropLast492[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G, H, I, J, K) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J, k K) {
		e, f, g, h, i, j, k, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop493 with func( 4 in)(9 out) drop first 3 result
func Drop493[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D) (h H, i I, j J, k K, l L, m M) {
		_, _, _, h, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast493 with func( 4 in)(9 out) drop last 3 result
func DropLast493[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G, H, I, J) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I, j J) {
		e, f, g, h, i, j, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop494 with func( 4 in)(9 out) drop first 4 result
func Drop494[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (I, J, K, L, M) {
	return func(a A, b B, c C, d D) (i I, j J, k K, l L, m M) {
		_, _, _, _, i, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast494 with func( 4 in)(9 out) drop last 4 result
func DropLast494[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G, H, I) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H, i I) {
		e, f, g, h, i, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop495 with func( 4 in)(9 out) drop first 5 result
func Drop495[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (J, K, L, M) {
	return func(a A, b B, c C, d D) (j J, k K, l L, m M) {
		_, _, _, _, _, j, k, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast495 with func( 4 in)(9 out) drop last 5 result
func DropLast495[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G, H) {
	return func(a A, b B, c C, d D) (e E, f F, g G, h H) {
		e, f, g, h, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop496 with func( 4 in)(9 out) drop first 6 result
func Drop496[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (K, L, M) {
	return func(a A, b B, c C, d D) (k K, l L, m M) {
		_, _, _, _, _, _, k, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast496 with func( 4 in)(9 out) drop last 6 result
func DropLast496[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F, G) {
	return func(a A, b B, c C, d D) (e E, f F, g G) {
		e, f, g, _, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop497 with func( 4 in)(9 out) drop first 7 result
func Drop497[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (L, M) {
	return func(a A, b B, c C, d D) (l L, m M) {
		_, _, _, _, _, _, _, l, m = fn(a, b, c, d)
		return
	}
}

// DropLast497 with func( 4 in)(9 out) drop last 7 result
func DropLast497[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) (E, F) {
	return func(a A, b B, c C, d D) (e E, f F) {
		e, f, _, _, _, _, _, _, _ = fn(a, b, c, d)
		return
	}
}

// Drop498 with func( 4 in)(9 out) drop first 8 result
func Drop498[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) M {
	return func(a A, b B, c C, d D) (m M) {
		_, _, _, _, _, _, _, _, m = fn(a, b, c, d)
		return
	}
}

// DropLast498 with func( 4 in)(9 out) drop last 8 result
func DropLast498[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D) (E, F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D) E {
	return func(a A, b B, c C, d D) (e E) {
		e, _, _, _, _, _, _, _, _ = fn(a, b, c, d)
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

// Drop561 with func( 5 in)(6 out) drop first 1 result
func Drop561[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K) {
		_, g, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// DropLast561 with func( 5 in)(6 out) drop last 1 result
func DropLast561[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (F, G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop562 with func( 5 in)(6 out) drop first 2 result
func Drop562[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K) {
		_, _, h, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// DropLast562 with func( 5 in)(6 out) drop last 2 result
func DropLast562[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (F, G, H, I) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop563 with func( 5 in)(6 out) drop first 3 result
func Drop563[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (I, J, K) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K) {
		_, _, _, i, j, k = fn(a, b, c, d, e)
		return
	}
}

// DropLast563 with func( 5 in)(6 out) drop last 3 result
func DropLast563[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (F, G, H) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop564 with func( 5 in)(6 out) drop first 4 result
func Drop564[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (J, K) {
	return func(a A, b B, c C, d D, e E) (j J, k K) {
		_, _, _, _, j, k = fn(a, b, c, d, e)
		return
	}
}

// DropLast564 with func( 5 in)(6 out) drop last 4 result
func DropLast564[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop565 with func( 5 in)(6 out) drop first 5 result
func Drop565[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) K {
	return func(a A, b B, c C, d D, e E) (k K) {
		_, _, _, _, _, k = fn(a, b, c, d, e)
		return
	}
}

// DropLast565 with func( 5 in)(6 out) drop last 5 result
func DropLast565[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E) (F, G, H, I, J, K)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop571 with func( 5 in)(7 out) drop first 1 result
func Drop571[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L) {
		_, g, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// DropLast571 with func( 5 in)(7 out) drop last 1 result
func DropLast571[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop572 with func( 5 in)(7 out) drop first 2 result
func Drop572[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L) {
		_, _, h, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// DropLast572 with func( 5 in)(7 out) drop last 2 result
func DropLast572[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (F, G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop573 with func( 5 in)(7 out) drop first 3 result
func Drop573[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L) {
		_, _, _, i, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// DropLast573 with func( 5 in)(7 out) drop last 3 result
func DropLast573[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (F, G, H, I) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop574 with func( 5 in)(7 out) drop first 4 result
func Drop574[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (J, K, L) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L) {
		_, _, _, _, j, k, l = fn(a, b, c, d, e)
		return
	}
}

// DropLast574 with func( 5 in)(7 out) drop last 4 result
func DropLast574[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (F, G, H) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop575 with func( 5 in)(7 out) drop first 5 result
func Drop575[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (K, L) {
	return func(a A, b B, c C, d D, e E) (k K, l L) {
		_, _, _, _, _, k, l = fn(a, b, c, d, e)
		return
	}
}

// DropLast575 with func( 5 in)(7 out) drop last 5 result
func DropLast575[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop576 with func( 5 in)(7 out) drop first 6 result
func Drop576[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) L {
	return func(a A, b B, c C, d D, e E) (l L) {
		_, _, _, _, _, _, l = fn(a, b, c, d, e)
		return
	}
}

// DropLast576 with func( 5 in)(7 out) drop last 6 result
func DropLast576[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop581 with func( 5 in)(8 out) drop first 1 result
func Drop581[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M) {
		_, g, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast581 with func( 5 in)(8 out) drop last 1 result
func DropLast581[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop582 with func( 5 in)(8 out) drop first 2 result
func Drop582[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M) {
		_, _, h, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast582 with func( 5 in)(8 out) drop last 2 result
func DropLast582[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop583 with func( 5 in)(8 out) drop first 3 result
func Drop583[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M) {
		_, _, _, i, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast583 with func( 5 in)(8 out) drop last 3 result
func DropLast583[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop584 with func( 5 in)(8 out) drop first 4 result
func Drop584[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M) {
		_, _, _, _, j, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast584 with func( 5 in)(8 out) drop last 4 result
func DropLast584[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G, H, I) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop585 with func( 5 in)(8 out) drop first 5 result
func Drop585[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (K, L, M) {
	return func(a A, b B, c C, d D, e E) (k K, l L, m M) {
		_, _, _, _, _, k, l, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast585 with func( 5 in)(8 out) drop last 5 result
func DropLast585[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G, H) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop586 with func( 5 in)(8 out) drop first 6 result
func Drop586[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (L, M) {
	return func(a A, b B, c C, d D, e E) (l L, m M) {
		_, _, _, _, _, _, l, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast586 with func( 5 in)(8 out) drop last 6 result
func DropLast586[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop587 with func( 5 in)(8 out) drop first 7 result
func Drop587[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) M {
	return func(a A, b B, c C, d D, e E) (m M) {
		_, _, _, _, _, _, _, m = fn(a, b, c, d, e)
		return
	}
}

// DropLast587 with func( 5 in)(8 out) drop last 7 result
func DropLast587[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop591 with func( 5 in)(9 out) drop first 1 result
func Drop591[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (g G, h H, i I, j J, k K, l L, m M, n N) {
		_, g, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast591 with func( 5 in)(9 out) drop last 1 result
func DropLast591[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L, m M) {
		f, g, h, i, j, k, l, m, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop592 with func( 5 in)(9 out) drop first 2 result
func Drop592[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (h H, i I, j J, k K, l L, m M, n N) {
		_, _, h, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast592 with func( 5 in)(9 out) drop last 2 result
func DropLast592[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K, l L) {
		f, g, h, i, j, k, l, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop593 with func( 5 in)(9 out) drop first 3 result
func Drop593[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (i I, j J, k K, l L, m M, n N) {
		_, _, _, i, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast593 with func( 5 in)(9 out) drop last 3 result
func DropLast593[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J, k K) {
		f, g, h, i, j, k, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop594 with func( 5 in)(9 out) drop first 4 result
func Drop594[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (j J, k K, l L, m M, n N) {
		_, _, _, _, j, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast594 with func( 5 in)(9 out) drop last 4 result
func DropLast594[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H, I, J) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I, j J) {
		f, g, h, i, j, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop595 with func( 5 in)(9 out) drop first 5 result
func Drop595[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (K, L, M, N) {
	return func(a A, b B, c C, d D, e E) (k K, l L, m M, n N) {
		_, _, _, _, _, k, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast595 with func( 5 in)(9 out) drop last 5 result
func DropLast595[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H, I) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H, i I) {
		f, g, h, i, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop596 with func( 5 in)(9 out) drop first 6 result
func Drop596[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (L, M, N) {
	return func(a A, b B, c C, d D, e E) (l L, m M, n N) {
		_, _, _, _, _, _, l, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast596 with func( 5 in)(9 out) drop last 6 result
func DropLast596[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G, H) {
	return func(a A, b B, c C, d D, e E) (f F, g G, h H) {
		f, g, h, _, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop597 with func( 5 in)(9 out) drop first 7 result
func Drop597[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (M, N) {
	return func(a A, b B, c C, d D, e E) (m M, n N) {
		_, _, _, _, _, _, _, m, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast597 with func( 5 in)(9 out) drop last 7 result
func DropLast597[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) (F, G) {
	return func(a A, b B, c C, d D, e E) (f F, g G) {
		f, g, _, _, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop598 with func( 5 in)(9 out) drop first 8 result
func Drop598[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) N {
	return func(a A, b B, c C, d D, e E) (n N) {
		_, _, _, _, _, _, _, _, n = fn(a, b, c, d, e)
		return
	}
}

// DropLast598 with func( 5 in)(9 out) drop last 8 result
func DropLast598[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E) (F, G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E) F {
	return func(a A, b B, c C, d D, e E) (f F) {
		f, _, _, _, _, _, _, _, _ = fn(a, b, c, d, e)
		return
	}
}

// Drop621 with func( 6 in)(2 out) drop first 1 result
func Drop621[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F) (G, H)) func(a A, b B, c C, d D, e E, f F) H {
	return func(a A, b B, c C, d D, e E, f F) (h H) {
		_, h = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast621 with func( 6 in)(2 out) drop last 1 result
func DropLast621[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F) (G, H)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop631 with func( 6 in)(3 out) drop first 1 result
func Drop631[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F) (G, H, I)) func(a A, b B, c C, d D, e E, f F) (H, I) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I) {
		_, h, i = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast631 with func( 6 in)(3 out) drop last 1 result
func DropLast631[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F) (G, H, I)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop632 with func( 6 in)(3 out) drop first 2 result
func Drop632[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F) (G, H, I)) func(a A, b B, c C, d D, e E, f F) I {
	return func(a A, b B, c C, d D, e E, f F) (i I) {
		_, _, i = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast632 with func( 6 in)(3 out) drop last 2 result
func DropLast632[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F) (G, H, I)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop641 with func( 6 in)(4 out) drop first 1 result
func Drop641[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J) {
		_, h, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast641 with func( 6 in)(4 out) drop last 1 result
func DropLast641[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) (G, H, I) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop642 with func( 6 in)(4 out) drop first 2 result
func Drop642[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) (I, J) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J) {
		_, _, i, j = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast642 with func( 6 in)(4 out) drop last 2 result
func DropLast642[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop643 with func( 6 in)(4 out) drop first 3 result
func Drop643[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) J {
	return func(a A, b B, c C, d D, e E, f F) (j J) {
		_, _, _, j = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast643 with func( 6 in)(4 out) drop last 3 result
func DropLast643[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F) (G, H, I, J)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop651 with func( 6 in)(5 out) drop first 1 result
func Drop651[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K) {
		_, h, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast651 with func( 6 in)(5 out) drop last 1 result
func DropLast651[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop652 with func( 6 in)(5 out) drop first 2 result
func Drop652[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K) {
		_, _, i, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast652 with func( 6 in)(5 out) drop last 2 result
func DropLast652[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (G, H, I) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop653 with func( 6 in)(5 out) drop first 3 result
func Drop653[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (J, K) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K) {
		_, _, _, j, k = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast653 with func( 6 in)(5 out) drop last 3 result
func DropLast653[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop654 with func( 6 in)(5 out) drop first 4 result
func Drop654[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) K {
	return func(a A, b B, c C, d D, e E, f F) (k K) {
		_, _, _, _, k = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast654 with func( 6 in)(5 out) drop last 4 result
func DropLast654[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F) (G, H, I, J, K)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop661 with func( 6 in)(6 out) drop first 1 result
func Drop661[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L) {
		_, h, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast661 with func( 6 in)(6 out) drop last 1 result
func DropLast661[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop662 with func( 6 in)(6 out) drop first 2 result
func Drop662[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L) {
		_, _, i, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast662 with func( 6 in)(6 out) drop last 2 result
func DropLast662[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop663 with func( 6 in)(6 out) drop first 3 result
func Drop663[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L) {
		_, _, _, j, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast663 with func( 6 in)(6 out) drop last 3 result
func DropLast663[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (G, H, I) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop664 with func( 6 in)(6 out) drop first 4 result
func Drop664[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (K, L) {
	return func(a A, b B, c C, d D, e E, f F) (k K, l L) {
		_, _, _, _, k, l = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast664 with func( 6 in)(6 out) drop last 4 result
func DropLast664[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop665 with func( 6 in)(6 out) drop first 5 result
func Drop665[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) L {
	return func(a A, b B, c C, d D, e E, f F) (l L) {
		_, _, _, _, _, l = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast665 with func( 6 in)(6 out) drop last 5 result
func DropLast665[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop671 with func( 6 in)(7 out) drop first 1 result
func Drop671[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M) {
		_, h, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast671 with func( 6 in)(7 out) drop last 1 result
func DropLast671[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop672 with func( 6 in)(7 out) drop first 2 result
func Drop672[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M) {
		_, _, i, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast672 with func( 6 in)(7 out) drop last 2 result
func DropLast672[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop673 with func( 6 in)(7 out) drop first 3 result
func Drop673[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M) {
		_, _, _, j, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast673 with func( 6 in)(7 out) drop last 3 result
func DropLast673[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop674 with func( 6 in)(7 out) drop first 4 result
func Drop674[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (k K, l L, m M) {
		_, _, _, _, k, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast674 with func( 6 in)(7 out) drop last 4 result
func DropLast674[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (G, H, I) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop675 with func( 6 in)(7 out) drop first 5 result
func Drop675[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (L, M) {
	return func(a A, b B, c C, d D, e E, f F) (l L, m M) {
		_, _, _, _, _, l, m = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast675 with func( 6 in)(7 out) drop last 5 result
func DropLast675[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop676 with func( 6 in)(7 out) drop first 6 result
func Drop676[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) M {
	return func(a A, b B, c C, d D, e E, f F) (m M) {
		_, _, _, _, _, _, m = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast676 with func( 6 in)(7 out) drop last 6 result
func DropLast676[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop681 with func( 6 in)(8 out) drop first 1 result
func Drop681[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N) {
		_, h, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast681 with func( 6 in)(8 out) drop last 1 result
func DropLast681[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop682 with func( 6 in)(8 out) drop first 2 result
func Drop682[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N) {
		_, _, i, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast682 with func( 6 in)(8 out) drop last 2 result
func DropLast682[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop683 with func( 6 in)(8 out) drop first 3 result
func Drop683[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N) {
		_, _, _, j, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast683 with func( 6 in)(8 out) drop last 3 result
func DropLast683[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop684 with func( 6 in)(8 out) drop first 4 result
func Drop684[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (k K, l L, m M, n N) {
		_, _, _, _, k, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast684 with func( 6 in)(8 out) drop last 4 result
func DropLast684[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop685 with func( 6 in)(8 out) drop first 5 result
func Drop685[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (l L, m M, n N) {
		_, _, _, _, _, l, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast685 with func( 6 in)(8 out) drop last 5 result
func DropLast685[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H, I) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop686 with func( 6 in)(8 out) drop first 6 result
func Drop686[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (M, N) {
	return func(a A, b B, c C, d D, e E, f F) (m M, n N) {
		_, _, _, _, _, _, m, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast686 with func( 6 in)(8 out) drop last 6 result
func DropLast686[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop687 with func( 6 in)(8 out) drop first 7 result
func Drop687[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) N {
	return func(a A, b B, c C, d D, e E, f F) (n N) {
		_, _, _, _, _, _, _, n = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast687 with func( 6 in)(8 out) drop last 7 result
func DropLast687[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop691 with func( 6 in)(9 out) drop first 1 result
func Drop691[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (h H, i I, j J, k K, l L, m M, n N, o O) {
		_, h, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast691 with func( 6 in)(9 out) drop last 1 result
func DropLast691[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M, n N) {
		g, h, i, j, k, l, m, n, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop692 with func( 6 in)(9 out) drop first 2 result
func Drop692[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (i I, j J, k K, l L, m M, n N, o O) {
		_, _, i, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast692 with func( 6 in)(9 out) drop last 2 result
func DropLast692[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L, m M) {
		g, h, i, j, k, l, m, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop693 with func( 6 in)(9 out) drop first 3 result
func Drop693[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (j J, k K, l L, m M, n N, o O) {
		_, _, _, j, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast693 with func( 6 in)(9 out) drop last 3 result
func DropLast693[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K, l L) {
		g, h, i, j, k, l, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop694 with func( 6 in)(9 out) drop first 4 result
func Drop694[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (k K, l L, m M, n N, o O) {
		_, _, _, _, k, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast694 with func( 6 in)(9 out) drop last 4 result
func DropLast694[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J, k K) {
		g, h, i, j, k, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop695 with func( 6 in)(9 out) drop first 5 result
func Drop695[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (l L, m M, n N, o O) {
		_, _, _, _, _, l, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast695 with func( 6 in)(9 out) drop last 5 result
func DropLast695[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I, J) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I, j J) {
		g, h, i, j, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop696 with func( 6 in)(9 out) drop first 6 result
func Drop696[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (M, N, O) {
	return func(a A, b B, c C, d D, e E, f F) (m M, n N, o O) {
		_, _, _, _, _, _, m, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast696 with func( 6 in)(9 out) drop last 6 result
func DropLast696[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H, I) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H, i I) {
		g, h, i, _, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop697 with func( 6 in)(9 out) drop first 7 result
func Drop697[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (N, O) {
	return func(a A, b B, c C, d D, e E, f F) (n N, o O) {
		_, _, _, _, _, _, _, n, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast697 with func( 6 in)(9 out) drop last 7 result
func DropLast697[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) (G, H) {
	return func(a A, b B, c C, d D, e E, f F) (g G, h H) {
		g, h, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop698 with func( 6 in)(9 out) drop first 8 result
func Drop698[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) O {
	return func(a A, b B, c C, d D, e E, f F) (o O) {
		_, _, _, _, _, _, _, _, o = fn(a, b, c, d, e, f)
		return
	}
}

// DropLast698 with func( 6 in)(9 out) drop last 8 result
func DropLast698[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F) (G, H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F) G {
	return func(a A, b B, c C, d D, e E, f F) (g G) {
		g, _, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f)
		return
	}
}

// Drop721 with func( 7 in)(2 out) drop first 1 result
func Drop721[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G) (H, I)) func(a A, b B, c C, d D, e E, f F, g G) I {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I) {
		_, i = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast721 with func( 7 in)(2 out) drop last 1 result
func DropLast721[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G) (H, I)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop731 with func( 7 in)(3 out) drop first 1 result
func Drop731[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G) (H, I, J)) func(a A, b B, c C, d D, e E, f F, g G) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J) {
		_, i, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast731 with func( 7 in)(3 out) drop last 1 result
func DropLast731[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G) (H, I, J)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop732 with func( 7 in)(3 out) drop first 2 result
func Drop732[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G) (H, I, J)) func(a A, b B, c C, d D, e E, f F, g G) J {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J) {
		_, _, j = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast732 with func( 7 in)(3 out) drop last 2 result
func DropLast732[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G) (H, I, J)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop741 with func( 7 in)(4 out) drop first 1 result
func Drop741[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K) {
		_, i, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast741 with func( 7 in)(4 out) drop last 1 result
func DropLast741[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop742 with func( 7 in)(4 out) drop first 2 result
func Drop742[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K) {
		_, _, j, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast742 with func( 7 in)(4 out) drop last 2 result
func DropLast742[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop743 with func( 7 in)(4 out) drop first 3 result
func Drop743[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) K {
	return func(a A, b B, c C, d D, e E, f F, g G) (k K) {
		_, _, _, k = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast743 with func( 7 in)(4 out) drop last 3 result
func DropLast743[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G) (H, I, J, K)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop751 with func( 7 in)(5 out) drop first 1 result
func Drop751[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L) {
		_, i, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast751 with func( 7 in)(5 out) drop last 1 result
func DropLast751[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop752 with func( 7 in)(5 out) drop first 2 result
func Drop752[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L) {
		_, _, j, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast752 with func( 7 in)(5 out) drop last 2 result
func DropLast752[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop753 with func( 7 in)(5 out) drop first 3 result
func Drop753[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (k K, l L) {
		_, _, _, k, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast753 with func( 7 in)(5 out) drop last 3 result
func DropLast753[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop754 with func( 7 in)(5 out) drop first 4 result
func Drop754[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) L {
	return func(a A, b B, c C, d D, e E, f F, g G) (l L) {
		_, _, _, _, l = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast754 with func( 7 in)(5 out) drop last 4 result
func DropLast754[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop761 with func( 7 in)(6 out) drop first 1 result
func Drop761[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M) {
		_, i, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast761 with func( 7 in)(6 out) drop last 1 result
func DropLast761[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop762 with func( 7 in)(6 out) drop first 2 result
func Drop762[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M) {
		_, _, j, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast762 with func( 7 in)(6 out) drop last 2 result
func DropLast762[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop763 with func( 7 in)(6 out) drop first 3 result
func Drop763[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (k K, l L, m M) {
		_, _, _, k, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast763 with func( 7 in)(6 out) drop last 3 result
func DropLast763[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop764 with func( 7 in)(6 out) drop first 4 result
func Drop764[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (l L, m M) {
		_, _, _, _, l, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast764 with func( 7 in)(6 out) drop last 4 result
func DropLast764[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop765 with func( 7 in)(6 out) drop first 5 result
func Drop765[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) M {
	return func(a A, b B, c C, d D, e E, f F, g G) (m M) {
		_, _, _, _, _, m = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast765 with func( 7 in)(6 out) drop last 5 result
func DropLast765[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop771 with func( 7 in)(7 out) drop first 1 result
func Drop771[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N) {
		_, i, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast771 with func( 7 in)(7 out) drop last 1 result
func DropLast771[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop772 with func( 7 in)(7 out) drop first 2 result
func Drop772[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N) {
		_, _, j, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast772 with func( 7 in)(7 out) drop last 2 result
func DropLast772[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop773 with func( 7 in)(7 out) drop first 3 result
func Drop773[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (k K, l L, m M, n N) {
		_, _, _, k, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast773 with func( 7 in)(7 out) drop last 3 result
func DropLast773[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop774 with func( 7 in)(7 out) drop first 4 result
func Drop774[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (l L, m M, n N) {
		_, _, _, _, l, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast774 with func( 7 in)(7 out) drop last 4 result
func DropLast774[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop775 with func( 7 in)(7 out) drop first 5 result
func Drop775[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (m M, n N) {
		_, _, _, _, _, m, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast775 with func( 7 in)(7 out) drop last 5 result
func DropLast775[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop776 with func( 7 in)(7 out) drop first 6 result
func Drop776[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) N {
	return func(a A, b B, c C, d D, e E, f F, g G) (n N) {
		_, _, _, _, _, _, n = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast776 with func( 7 in)(7 out) drop last 6 result
func DropLast776[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop781 with func( 7 in)(8 out) drop first 1 result
func Drop781[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O) {
		_, i, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast781 with func( 7 in)(8 out) drop last 1 result
func DropLast781[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop782 with func( 7 in)(8 out) drop first 2 result
func Drop782[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O) {
		_, _, j, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast782 with func( 7 in)(8 out) drop last 2 result
func DropLast782[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop783 with func( 7 in)(8 out) drop first 3 result
func Drop783[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (k K, l L, m M, n N, o O) {
		_, _, _, k, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast783 with func( 7 in)(8 out) drop last 3 result
func DropLast783[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop784 with func( 7 in)(8 out) drop first 4 result
func Drop784[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (l L, m M, n N, o O) {
		_, _, _, _, l, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast784 with func( 7 in)(8 out) drop last 4 result
func DropLast784[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop785 with func( 7 in)(8 out) drop first 5 result
func Drop785[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (m M, n N, o O) {
		_, _, _, _, _, m, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast785 with func( 7 in)(8 out) drop last 5 result
func DropLast785[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop786 with func( 7 in)(8 out) drop first 6 result
func Drop786[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (n N, o O) {
		_, _, _, _, _, _, n, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast786 with func( 7 in)(8 out) drop last 6 result
func DropLast786[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop787 with func( 7 in)(8 out) drop first 7 result
func Drop787[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) O {
	return func(a A, b B, c C, d D, e E, f F, g G) (o O) {
		_, _, _, _, _, _, _, o = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast787 with func( 7 in)(8 out) drop last 7 result
func DropLast787[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop791 with func( 7 in)(9 out) drop first 1 result
func Drop791[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (i I, j J, k K, l L, m M, n N, o O, p P) {
		_, i, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast791 with func( 7 in)(9 out) drop last 1 result
func DropLast791[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N, o O) {
		h, i, j, k, l, m, n, o, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop792 with func( 7 in)(9 out) drop first 2 result
func Drop792[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (j J, k K, l L, m M, n N, o O, p P) {
		_, _, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast792 with func( 7 in)(9 out) drop last 2 result
func DropLast792[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M, n N) {
		h, i, j, k, l, m, n, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop793 with func( 7 in)(9 out) drop first 3 result
func Drop793[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (k K, l L, m M, n N, o O, p P) {
		_, _, _, k, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast793 with func( 7 in)(9 out) drop last 3 result
func DropLast793[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L, m M) {
		h, i, j, k, l, m, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop794 with func( 7 in)(9 out) drop first 4 result
func Drop794[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (l L, m M, n N, o O, p P) {
		_, _, _, _, l, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast794 with func( 7 in)(9 out) drop last 4 result
func DropLast794[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K, l L) {
		h, i, j, k, l, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop795 with func( 7 in)(9 out) drop first 5 result
func Drop795[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (m M, n N, o O, p P) {
		_, _, _, _, _, m, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast795 with func( 7 in)(9 out) drop last 5 result
func DropLast795[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J, k K) {
		h, i, j, k, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop796 with func( 7 in)(9 out) drop first 6 result
func Drop796[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (n N, o O, p P) {
		_, _, _, _, _, _, n, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast796 with func( 7 in)(9 out) drop last 6 result
func DropLast796[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I, j J) {
		h, i, j, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop797 with func( 7 in)(9 out) drop first 7 result
func Drop797[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G) (o O, p P) {
		_, _, _, _, _, _, _, o, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast797 with func( 7 in)(9 out) drop last 7 result
func DropLast797[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) (H, I) {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H, i I) {
		h, i, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop798 with func( 7 in)(9 out) drop first 8 result
func Drop798[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) P {
	return func(a A, b B, c C, d D, e E, f F, g G) (p P) {
		_, _, _, _, _, _, _, _, p = fn(a, b, c, d, e, f, g)
		return
	}
}

// DropLast798 with func( 7 in)(9 out) drop last 8 result
func DropLast798[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G) (H, I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) (h H) {
		h, _, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g)
		return
	}
}

// Drop821 with func( 8 in)(2 out) drop first 1 result
func Drop821[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H) (I, J)) func(a A, b B, c C, d D, e E, f F, g G, h H) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J) {
		_, j = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast821 with func( 8 in)(2 out) drop last 1 result
func DropLast821[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H) (I, J)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop831 with func( 8 in)(3 out) drop first 1 result
func Drop831[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K) {
		_, j, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast831 with func( 8 in)(3 out) drop last 1 result
func DropLast831[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop832 with func( 8 in)(3 out) drop first 2 result
func Drop832[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H) K {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K) {
		_, _, k = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast832 with func( 8 in)(3 out) drop last 2 result
func DropLast832[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H) (I, J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop841 with func( 8 in)(4 out) drop first 1 result
func Drop841[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L) {
		_, j, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast841 with func( 8 in)(4 out) drop last 1 result
func DropLast841[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop842 with func( 8 in)(4 out) drop first 2 result
func Drop842[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) (K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K, l L) {
		_, _, k, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast842 with func( 8 in)(4 out) drop last 2 result
func DropLast842[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop843 with func( 8 in)(4 out) drop first 3 result
func Drop843[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) L {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (l L) {
		_, _, _, l = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast843 with func( 8 in)(4 out) drop last 3 result
func DropLast843[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop851 with func( 8 in)(5 out) drop first 1 result
func Drop851[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M) {
		_, j, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast851 with func( 8 in)(5 out) drop last 1 result
func DropLast851[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop852 with func( 8 in)(5 out) drop first 2 result
func Drop852[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K, l L, m M) {
		_, _, k, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast852 with func( 8 in)(5 out) drop last 2 result
func DropLast852[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop853 with func( 8 in)(5 out) drop first 3 result
func Drop853[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (l L, m M) {
		_, _, _, l, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast853 with func( 8 in)(5 out) drop last 3 result
func DropLast853[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop854 with func( 8 in)(5 out) drop first 4 result
func Drop854[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) M {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (m M) {
		_, _, _, _, m = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast854 with func( 8 in)(5 out) drop last 4 result
func DropLast854[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop861 with func( 8 in)(6 out) drop first 1 result
func Drop861[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N) {
		_, j, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast861 with func( 8 in)(6 out) drop last 1 result
func DropLast861[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop862 with func( 8 in)(6 out) drop first 2 result
func Drop862[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K, l L, m M, n N) {
		_, _, k, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast862 with func( 8 in)(6 out) drop last 2 result
func DropLast862[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop863 with func( 8 in)(6 out) drop first 3 result
func Drop863[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (l L, m M, n N) {
		_, _, _, l, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast863 with func( 8 in)(6 out) drop last 3 result
func DropLast863[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop864 with func( 8 in)(6 out) drop first 4 result
func Drop864[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (m M, n N) {
		_, _, _, _, m, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast864 with func( 8 in)(6 out) drop last 4 result
func DropLast864[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop865 with func( 8 in)(6 out) drop first 5 result
func Drop865[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) N {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (n N) {
		_, _, _, _, _, n = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast865 with func( 8 in)(6 out) drop last 5 result
func DropLast865[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop871 with func( 8 in)(7 out) drop first 1 result
func Drop871[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O) {
		_, j, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast871 with func( 8 in)(7 out) drop last 1 result
func DropLast871[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop872 with func( 8 in)(7 out) drop first 2 result
func Drop872[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K, l L, m M, n N, o O) {
		_, _, k, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast872 with func( 8 in)(7 out) drop last 2 result
func DropLast872[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop873 with func( 8 in)(7 out) drop first 3 result
func Drop873[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (l L, m M, n N, o O) {
		_, _, _, l, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast873 with func( 8 in)(7 out) drop last 3 result
func DropLast873[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop874 with func( 8 in)(7 out) drop first 4 result
func Drop874[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (m M, n N, o O) {
		_, _, _, _, m, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast874 with func( 8 in)(7 out) drop last 4 result
func DropLast874[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop875 with func( 8 in)(7 out) drop first 5 result
func Drop875[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (n N, o O) {
		_, _, _, _, _, n, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast875 with func( 8 in)(7 out) drop last 5 result
func DropLast875[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop876 with func( 8 in)(7 out) drop first 6 result
func Drop876[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) O {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (o O) {
		_, _, _, _, _, _, o = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast876 with func( 8 in)(7 out) drop last 6 result
func DropLast876[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop881 with func( 8 in)(8 out) drop first 1 result
func Drop881[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P) {
		_, j, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast881 with func( 8 in)(8 out) drop last 1 result
func DropLast881[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop882 with func( 8 in)(8 out) drop first 2 result
func Drop882[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K, l L, m M, n N, o O, p P) {
		_, _, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast882 with func( 8 in)(8 out) drop last 2 result
func DropLast882[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop883 with func( 8 in)(8 out) drop first 3 result
func Drop883[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (l L, m M, n N, o O, p P) {
		_, _, _, l, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast883 with func( 8 in)(8 out) drop last 3 result
func DropLast883[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop884 with func( 8 in)(8 out) drop first 4 result
func Drop884[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (m M, n N, o O, p P) {
		_, _, _, _, m, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast884 with func( 8 in)(8 out) drop last 4 result
func DropLast884[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop885 with func( 8 in)(8 out) drop first 5 result
func Drop885[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (n N, o O, p P) {
		_, _, _, _, _, n, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast885 with func( 8 in)(8 out) drop last 5 result
func DropLast885[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop886 with func( 8 in)(8 out) drop first 6 result
func Drop886[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (o O, p P) {
		_, _, _, _, _, _, o, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast886 with func( 8 in)(8 out) drop last 6 result
func DropLast886[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop887 with func( 8 in)(8 out) drop first 7 result
func Drop887[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) P {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (p P) {
		_, _, _, _, _, _, _, p = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast887 with func( 8 in)(8 out) drop last 7 result
func DropLast887[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop891 with func( 8 in)(9 out) drop first 1 result
func Drop891[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		_, j, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast891 with func( 8 in)(9 out) drop last 1 result
func DropLast891[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O, p P) {
		i, j, k, l, m, n, o, p, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop892 with func( 8 in)(9 out) drop first 2 result
func Drop892[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (k K, l L, m M, n N, o O, p P, q Q) {
		_, _, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast892 with func( 8 in)(9 out) drop last 2 result
func DropLast892[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N, o O) {
		i, j, k, l, m, n, o, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop893 with func( 8 in)(9 out) drop first 3 result
func Drop893[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (l L, m M, n N, o O, p P, q Q) {
		_, _, _, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast893 with func( 8 in)(9 out) drop last 3 result
func DropLast893[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M, n N) {
		i, j, k, l, m, n, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop894 with func( 8 in)(9 out) drop first 4 result
func Drop894[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (m M, n N, o O, p P, q Q) {
		_, _, _, _, m, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast894 with func( 8 in)(9 out) drop last 4 result
func DropLast894[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L, m M) {
		i, j, k, l, m, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop895 with func( 8 in)(9 out) drop first 5 result
func Drop895[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (n N, o O, p P, q Q) {
		_, _, _, _, _, n, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast895 with func( 8 in)(9 out) drop last 5 result
func DropLast895[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K, l L) {
		i, j, k, l, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop896 with func( 8 in)(9 out) drop first 6 result
func Drop896[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (o O, p P, q Q) {
		_, _, _, _, _, _, o, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast896 with func( 8 in)(9 out) drop last 6 result
func DropLast896[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J, k K) {
		i, j, k, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop897 with func( 8 in)(9 out) drop first 7 result
func Drop897[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (p P, q Q) {
		_, _, _, _, _, _, _, p, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast897 with func( 8 in)(9 out) drop last 7 result
func DropLast897[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) (I, J) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I, j J) {
		i, j, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop898 with func( 8 in)(9 out) drop first 8 result
func Drop898[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) Q {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (q Q) {
		_, _, _, _, _, _, _, _, q = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// DropLast898 with func( 8 in)(9 out) drop last 8 result
func DropLast898[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H) (I, J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) (i I) {
		i, _, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h)
		return
	}
}

// Drop921 with func( 9 in)(2 out) drop first 1 result
func Drop921[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) K {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K) {
		_, k = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast921 with func( 9 in)(2 out) drop last 1 result
func DropLast921[A, B, C, D, E, F, G, H, I, J, K any](fn func(A, B, C, D, E, F, G, H, I) (J, K)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop931 with func( 9 in)(3 out) drop first 1 result
func Drop931[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L) {
		_, k, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast931 with func( 9 in)(3 out) drop last 1 result
func DropLast931[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop932 with func( 9 in)(3 out) drop first 2 result
func Drop932[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) L {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L) {
		_, _, l = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast932 with func( 9 in)(3 out) drop last 2 result
func DropLast932[A, B, C, D, E, F, G, H, I, J, K, L any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop941 with func( 9 in)(4 out) drop first 1 result
func Drop941[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L, m M) {
		_, k, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast941 with func( 9 in)(4 out) drop last 1 result
func DropLast941[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop942 with func( 9 in)(4 out) drop first 2 result
func Drop942[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L, m M) {
		_, _, l, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast942 with func( 9 in)(4 out) drop last 2 result
func DropLast942[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop943 with func( 9 in)(4 out) drop first 3 result
func Drop943[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) M {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (m M) {
		_, _, _, m = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast943 with func( 9 in)(4 out) drop last 3 result
func DropLast943[A, B, C, D, E, F, G, H, I, J, K, L, M any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop951 with func( 9 in)(5 out) drop first 1 result
func Drop951[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L, m M, n N) {
		_, k, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast951 with func( 9 in)(5 out) drop last 1 result
func DropLast951[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop952 with func( 9 in)(5 out) drop first 2 result
func Drop952[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L, m M, n N) {
		_, _, l, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast952 with func( 9 in)(5 out) drop last 2 result
func DropLast952[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop953 with func( 9 in)(5 out) drop first 3 result
func Drop953[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (m M, n N) {
		_, _, _, m, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast953 with func( 9 in)(5 out) drop last 3 result
func DropLast953[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop954 with func( 9 in)(5 out) drop first 4 result
func Drop954[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) N {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (n N) {
		_, _, _, _, n = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast954 with func( 9 in)(5 out) drop last 4 result
func DropLast954[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop961 with func( 9 in)(6 out) drop first 1 result
func Drop961[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L, m M, n N, o O) {
		_, k, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast961 with func( 9 in)(6 out) drop last 1 result
func DropLast961[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop962 with func( 9 in)(6 out) drop first 2 result
func Drop962[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L, m M, n N, o O) {
		_, _, l, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast962 with func( 9 in)(6 out) drop last 2 result
func DropLast962[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop963 with func( 9 in)(6 out) drop first 3 result
func Drop963[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (m M, n N, o O) {
		_, _, _, m, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast963 with func( 9 in)(6 out) drop last 3 result
func DropLast963[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop964 with func( 9 in)(6 out) drop first 4 result
func Drop964[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (n N, o O) {
		_, _, _, _, n, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast964 with func( 9 in)(6 out) drop last 4 result
func DropLast964[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop965 with func( 9 in)(6 out) drop first 5 result
func Drop965[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) O {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (o O) {
		_, _, _, _, _, o = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast965 with func( 9 in)(6 out) drop last 5 result
func DropLast965[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop971 with func( 9 in)(7 out) drop first 1 result
func Drop971[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L, m M, n N, o O, p P) {
		_, k, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast971 with func( 9 in)(7 out) drop last 1 result
func DropLast971[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop972 with func( 9 in)(7 out) drop first 2 result
func Drop972[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L, m M, n N, o O, p P) {
		_, _, l, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast972 with func( 9 in)(7 out) drop last 2 result
func DropLast972[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop973 with func( 9 in)(7 out) drop first 3 result
func Drop973[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (m M, n N, o O, p P) {
		_, _, _, m, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast973 with func( 9 in)(7 out) drop last 3 result
func DropLast973[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop974 with func( 9 in)(7 out) drop first 4 result
func Drop974[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (n N, o O, p P) {
		_, _, _, _, n, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast974 with func( 9 in)(7 out) drop last 4 result
func DropLast974[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop975 with func( 9 in)(7 out) drop first 5 result
func Drop975[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (o O, p P) {
		_, _, _, _, _, o, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast975 with func( 9 in)(7 out) drop last 5 result
func DropLast975[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop976 with func( 9 in)(7 out) drop first 6 result
func Drop976[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) P {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (p P) {
		_, _, _, _, _, _, p = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast976 with func( 9 in)(7 out) drop last 6 result
func DropLast976[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop981 with func( 9 in)(8 out) drop first 1 result
func Drop981[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L, m M, n N, o O, p P, q Q) {
		_, k, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast981 with func( 9 in)(8 out) drop last 1 result
func DropLast981[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop982 with func( 9 in)(8 out) drop first 2 result
func Drop982[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L, m M, n N, o O, p P, q Q) {
		_, _, l, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast982 with func( 9 in)(8 out) drop last 2 result
func DropLast982[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop983 with func( 9 in)(8 out) drop first 3 result
func Drop983[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (m M, n N, o O, p P, q Q) {
		_, _, _, m, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast983 with func( 9 in)(8 out) drop last 3 result
func DropLast983[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop984 with func( 9 in)(8 out) drop first 4 result
func Drop984[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (n N, o O, p P, q Q) {
		_, _, _, _, n, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast984 with func( 9 in)(8 out) drop last 4 result
func DropLast984[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop985 with func( 9 in)(8 out) drop first 5 result
func Drop985[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (o O, p P, q Q) {
		_, _, _, _, _, o, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast985 with func( 9 in)(8 out) drop last 5 result
func DropLast985[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop986 with func( 9 in)(8 out) drop first 6 result
func Drop986[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (p P, q Q) {
		_, _, _, _, _, _, p, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast986 with func( 9 in)(8 out) drop last 6 result
func DropLast986[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop987 with func( 9 in)(8 out) drop first 7 result
func Drop987[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Q {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (q Q) {
		_, _, _, _, _, _, _, q = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast987 with func( 9 in)(8 out) drop last 7 result
func DropLast987[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop991 with func( 9 in)(9 out) drop first 1 result
func Drop991[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (K, L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (k K, l L, m M, n N, o O, p P, q Q, r R) {
		_, k, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast991 with func( 9 in)(9 out) drop last 1 result
func DropLast991[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, P, Q) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P, q Q) {
		j, k, l, m, n, o, p, q, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop992 with func( 9 in)(9 out) drop first 2 result
func Drop992[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (L, M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (l L, m M, n N, o O, p P, q Q, r R) {
		_, _, l, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast992 with func( 9 in)(9 out) drop last 2 result
func DropLast992[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O, P) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O, p P) {
		j, k, l, m, n, o, p, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop993 with func( 9 in)(9 out) drop first 3 result
func Drop993[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (M, N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (m M, n N, o O, p P, q Q, r R) {
		_, _, _, m, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast993 with func( 9 in)(9 out) drop last 3 result
func DropLast993[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N, O) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N, o O) {
		j, k, l, m, n, o, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop994 with func( 9 in)(9 out) drop first 4 result
func Drop994[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (N, O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (n N, o O, p P, q Q, r R) {
		_, _, _, _, n, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast994 with func( 9 in)(9 out) drop last 4 result
func DropLast994[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M, N) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M, n N) {
		j, k, l, m, n, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop995 with func( 9 in)(9 out) drop first 5 result
func Drop995[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (O, P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (o O, p P, q Q, r R) {
		_, _, _, _, _, o, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast995 with func( 9 in)(9 out) drop last 5 result
func DropLast995[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L, M) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L, m M) {
		j, k, l, m, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop996 with func( 9 in)(9 out) drop first 6 result
func Drop996[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (P, Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (p P, q Q, r R) {
		_, _, _, _, _, _, p, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast996 with func( 9 in)(9 out) drop last 6 result
func DropLast996[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K, L) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K, l L) {
		j, k, l, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop997 with func( 9 in)(9 out) drop first 7 result
func Drop997[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (Q, R) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (q Q, r R) {
		_, _, _, _, _, _, _, q, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast997 with func( 9 in)(9 out) drop last 7 result
func DropLast997[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (J, K) {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J, k K) {
		j, k, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// Drop998 with func( 9 in)(9 out) drop first 8 result
func Drop998[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) R {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (r R) {
		_, _, _, _, _, _, _, _, r = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}

// DropLast998 with func( 9 in)(9 out) drop last 8 result
func DropLast998[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](fn func(A, B, C, D, E, F, G, H, I) (J, K, L, M, N, O, P, Q, R)) func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) (j J) {
		j, _, _, _, _, _, _, _, _ = fn(a, b, c, d, e, f, g, h, i)
		return
	}
}
