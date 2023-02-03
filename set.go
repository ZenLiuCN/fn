package fn

//Set current base on map that's why it inside package maps
type Set[K comparable] interface {
	Clear()
	Values() []K
	Each(func(K) bool)
	EachIndex(func(int, K) bool)
	Add(K) bool
	Remove(K) bool
	Put(K)
	Delete(K)
	Len() int
	Exists(K) bool
}

//HashSet Set impl with map[K]struct{}
type HashSet[K comparable] map[K]struct{}

func NewHashSet[K comparable]() HashSet[K] {
	return make(HashSet[K])
}
func NewHashSetInit[K comparable](size int) HashSet[K] {
	return make(HashSet[K], max(0, size))
}
func max(n0, n1 int) int {
	if n0 > n1 {
		return n0
	}
	return n1
}

func (m HashSet[K]) Exists(k K) bool {
	return m[k] != Nothing
}
func (m HashSet[K]) Clear() {
	for k := range m {
		delete(m, k)
	}
}
func (m HashSet[K]) Values() []K {
	return MapKeys(m)
}
func (m HashSet[K]) Each(fn func(K) bool) {
	for k := range m {
		if fn(k) {
			break
		}
	}
}
func (m HashSet[K]) EachIndex(fn func(int, K) bool) {
	i := 0
	for k := range m {
		if fn(i, k) {
			break
		}
		i++
	}
}
func (m HashSet[K]) Add(k K) bool {
	if m.Exists(k) {
		return false
	}
	m[k] = Nothing
	return true
}
func (m HashSet[K]) Remove(k K) bool {
	if m.Exists(k) {
		delete(m, k)
		return true
	}
	return false
}
func (m HashSet[K]) Put(k K) {
	m[k] = Nothing
}
func (m HashSet[K]) Delete(k K) {
	if m.Exists(k) {
		delete(m, k)
	}
}
func (m HashSet[K]) Len() int {
	return len(m)
}
