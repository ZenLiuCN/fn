package fn

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenPanics(t *testing.T) {
	for i := 1; i < 6; i++ {
		printf("// Panic%d error or return other %d value\n", i, i)
		printf(`func Panic%d[`, i)
		n(i, a)
		print(` any](`)
		n(i, aA)
		print(", err error)(")
		n(i, A)
		print("){\n\tif err!=nil{\n\t\tpanic(err)\n\t}\n\treturn ")
		n(i, a)
		print("\n}\n\n")
	}
}
func TestGenRecovers(t *testing.T) {
	for i := 0; i < 6; i++ {
	no:
		for o := 0; o < 6; o++ {
			if i == o && o == 0 {
				continue no
			}
			printf("// Recover%d%d panic with func %d in %d out\n", i, o, i, o)
			printf(`func Recover%d%d[`, i, o)
			n(i+o, A)
			print(` any](fn func(`)
			n(i, A)
			print(`)(`)
			r(i, i+o, A)
			print(`) )func(`)
			n(i, aA)
			print(")(")
			r(i, i+o, A)
			if o != 0 {
				print(",")
			}
			print(`error){
	return func(`)
			n(i, aA)
			print(")(")
			r(i, i+o, aA)
			if o != 0 {
				print(",")
			}
			print("err error){\n")
			print(`		defer func() {
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
		`)
			r(i, i+o, a)
			if o != 0 {
				print("= ")
			}
			print("fn(")
			n(i, a)
			print(")\n\t\treturn\n\t}\n}\n\n")
		}

	}
}
func TestGenDrop(t *testing.T) {
	for i := 0; i < 6; i++ {
	no:
		for o := 0; o < 6; o++ {
			for d := 1; d < o; d++ {
				if o == 0 {
					continue no
				}
				printf("// Drop%d%d%d with func( %d in)(%d out) drop first %d result \n", i, o, d, i, o, d)
				printf(`func Drop%d%d%d[`, i, o, d)
				n(i+o, A)
				print(` any](fn func(`)
				//in param
				n(i, A)
				print(`)(`)
				//out param
				r(i, i+o, A)
				print(`) )func(`)
				//in param
				n(i, aA)
				print(")(")
				//out after drop
				r(i+d, i+o, A)
				print(`){
	return func(`)
				n(i, aA)
				print(")(")
				r(i+d, i+o, aA)
				print("){\n\t\t")
				f(d, "_")
				print(",")
				r(i+d, i+o, a)
				print(" = ")
				print("fn(")
				n(i, a)
				print(")\n\t\treturn\n\t}\n}\n\n")

				printf("// DropLast%d%d%d with func( %d in)(%d out) drop last %d result \n", i, o, d, i, o, d)
				printf(`func DropLast%d%d%d[`, i, o, d)
				n(i+o, A)
				print(` any](fn func(`)
				//in param
				n(i, A)
				print(`)(`)
				//out param
				r(i, i+o, A)
				print(`) )func(`)
				//in param
				n(i, aA)
				print(")(")
				//out after drop
				r(i, i+o-d, A)
				print(`){
	return func(`)
				n(i, aA)
				print(")(")
				r(i, i+o-d, aA)
				print("){\n\t\t")
				r(i, i+o-d, a)
				print(",")
				f(d, "_")
				print(" = ")
				print("fn(")
				n(i, a)
				print(")\n\t\treturn\n\t}\n}\n\n")
			}
		}

	}
}
func TestGenFix(t *testing.T) {
ni:
	for i := 0; i < 6; i++ {
		for o := 0; o < 6; o++ {
			for d := 1; d < i; d++ {
				if i == 0 {
					continue ni
				}
				printf("// Closure%d%d%d with func( %d in)(%d out) fix first %d argument \n", i, o, d, i, o, d)
				printf(`func Closure%d%d%d[`, i, o, d)
				n(i+o, A)
				print(` any](`)
				r(0, i-d, aA)
				print(`,fn func(`)
				//in param
				n(i, A)
				print(`)(`)
				//out param
				r(i, i+o, A)
				print(`)`)
				print(` )func(`)
				//in param
				r(i-d, i, A)
				print(")(")
				//outs
				r(i, i+o, A)
				print(`){
	return func(`)
				r(i-d, i, aA)
				print(")(")
				r(i, i+o, aA)
				print("){\n\t\t")
				r(i, i+o, a)
				if o != 0 {
					print(" = ")
				}
				print("fn(")
				n(i, a)
				print(")\n\t\treturn\n\t}\n}\n\n")

				printf("// ClosureLast%d%d%d with func( %d in)(%d out) fix last %d argument \n", i, o, d, i, o, d)
				printf(`func ClosureLast%d%d%d[`, i, o, d)
				n(i+o, A)
				print(` any](`)
				r(i-d, i, aA)
				print(`,fn func(`)
				//in param
				n(i, A)
				print(`)(`)
				//out param
				r(i, i+o, A)
				print(`)`)
				print(` )func(`)
				//in param
				r(0, i-d, A)
				print(")(")
				//outs
				r(i, i+o, A)
				print(`){
	return func(`)
				r(0, i-d, aA)
				print(")(")
				r(i, i+o, aA)
				print("){\n\t\t")
				r(i, i+o, a)
				if o != 0 {
					print(" = ")
				}
				print("fn(")
				n(i, a)
				print(")\n\t\treturn\n\t}\n}\n\n")
			}
		}

	}
}

func printf(p string, a ...any) {
	fmt.Printf(p, a...)
}
func id(i int, v string) {
	print(strings.Repeat("\t", i), v)
}
func a(i int) {
	print(string(rune(i + 'a')))
}
func aA(i int) {
	print(string(rune(i + 'a')))
	print(" ")
	print(string(rune(i + 'A')))
}
func A(i int) {
	print(string(rune(i + 'A')))
}
func n(i int, v func(j int)) {
	for j := 0; j < i; j++ {
		if j != 0 {
			print(",")
		}
		v(j)
	}
}
func f(i int, v string) {
	for j := 0; j < i; j++ {
		if j != 0 {
			print(",")
		}
		print(v)
	}
}
func r(i int, u int, v func(j int)) {
	for j := i; j < u; j++ {
		if j != i {
			print(",")
		}
		v(j)
	}
}
