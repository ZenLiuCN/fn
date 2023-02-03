//go:build gen

package fn

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	generate("fn_panic.go", genPanics)
	generate("fn_recovers.go", genRecovers)
	generate("fn_drops.go", genDrops)
	generate("fn_clos.go", genClosures)
	generate("fn_panics.go", genPanicsFn)
	/*	buf = bufio.NewWriter(os.Stdout)
		P(next(10, A))
		Panic(buf.Flush())*/
}
