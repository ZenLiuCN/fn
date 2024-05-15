package fn

import "bytes"

var (
	bytesPool = NewByteBufferPoolInitializedLimited(256, 1024*4) //4k
)

func GetBuffer() *bytes.Buffer {
	return bytesPool.Get()
}
func GetPooledBuffer() *Pooled[*bytes.Buffer] {
	return bytesPool.GetPooled()
}
func PutBuffer(b *bytes.Buffer) {
	bytesPool.Put(b)
}
