package maps

import "sort"
import . "github.com/ZenLiuCN/fn"

//EachEntries consume each key and value
func EachEntries[K comparable, V any, M ~map[K]V](m M, c func(K, V)) {
	for k, v := range m {
		c(k, v)
	}
}

//Entries map to Entry slice
func Entries[K comparable, V any, M ~map[K]V](m M) (r []*Entry[K, V]) {
	for k, v := range m {
		r = append(r, &Entry[K, V]{k, v})
	}
	return
}

//EntriesFilter pick subset of entry
func EntriesFilter[K comparable, V any, M ~map[K]V](m M, cmp func(K, V) bool) (r []*Entry[K, V]) {
	for k, v := range m {
		if cmp(k, v) {
			r = append(r, &Entry[K, V]{k, v})
		}
	}
	return
}

//Keys key slice of map
func Keys[K comparable, V any, M ~map[K]V](m M) (r []K) {
	for k := range m {
		r = append(r, k)
	}
	return
}

//HasKey check key is exists
func HasKey[K comparable, V any, M ~map[K]V](m M, k K) (ok bool) {
	_, ok = m[k]
	return
}

//KeyOf found key of a value otherwise the K's default value returned
func KeyOf[K comparable, V comparable, M ~map[K]V](m M, v V) (k K) {
	for k1, val := range m {
		if val == v {
			return k1
		}
	}
	return
}

//MapKey mapping key
func MapKey[K, K1 comparable, V any, M ~map[K]V](m M, fn func(K) K1) (o map[K1]V) {
	if m == nil || len(m) == 0 {
		return
	}
	for k, val := range m {
		o[fn(k)] = val
	}
	return
}

//Values value slice of map
func Values[K comparable, V any, M ~map[K]V](m M) (r []V) {
	for _, v := range m {
		r = append(r, v)
	}
	return
}

//MapValue mapping value
func MapValue[K comparable, V, V1 any, M ~map[K]V](m M, fn func(V) V1) (o map[K]V1) {
	if m == nil || len(m) == 0 {
		return
	}
	for k, val := range m {
		o[k] = fn(val)
	}
	return
}

//MapValueTo make map value to v
func MapValueTo[K comparable, V, V1 any, M ~map[K]V](m M, v V1) (o map[K]V1) {
	if m == nil || len(m) == 0 {
		return
	}
	for k := range m {
		o[k] = v
	}
	return
}

//HasValue check value exists
func HasValue[K comparable, V comparable, M ~map[K]V](m M, v V) (ok bool) {
	for _, val := range m {
		if val == v {
			return true
		}
	}
	return
}

//Filter filter entries
func Filter[K comparable, V, V1 any, M ~map[K]V](m M, pred func(K, V) bool) (o map[K]V) {
	for k, v := range m {
		if pred(k, v) {
			o[k] = v
		}
	}
	return
}

//Pick entries in to new map
func Pick[K comparable, V any, M ~map[K]V](m M, pick func(K, V) int) (r []map[K]V) {
	p := 0
	for k, v := range m {
		i := pick(k, v)
		if i < 0 {
			continue
		}
		if i > p {
			for x := 0; x < p-i; x++ {
				r = append(r, make(map[K]V))
			}
		}
		r[i][k] = v
	}
	return
}

//SortByValue sort by value and returns the key slice, see sort.Slice
func SortByValue[K comparable, V any, M ~map[K]V](m M, cmp func(V, V) bool) (r []K) {
	r = append(r, Keys(m)...)
	sort.Slice(r, func(i, j int) bool {
		return cmp(m[r[i]], m[r[j]])
	})
	return
}

//SortByValueStable sort by value and returns the key slice, see sort.SliceStable
func SortByValueStable[K comparable, V any, M ~map[K]V](m M, cmp func(V, V) bool) (r []K) {
	r = append(r, Keys(m)...)
	sort.SliceStable(r, func(i, j int) bool {
		return cmp(m[r[i]], m[r[j]])
	})
	return
}

type HashSet[K comparable] interface {
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

//Set HashSet impl with map[K]struct{}
type Set[K comparable] map[K]struct{}

func MewMapSet[K comparable]() Set[K] {
	return make(Set[K])
}
func MewMapSetInit[K comparable](initCap int) Set[K] {
	return make(Set[K], initCap)
}

func (m Set[K]) Exists(k K) bool {
	return m[k] != Nothing
}
func (m Set[K]) Clear() {
	for k := range m {
		delete(m, k)
	}
}
func (m Set[K]) Values() []K {
	return Keys(m)
}
func (m Set[K]) Each(fn func(K) bool) {
	for k := range m {
		if fn(k) {
			break
		}
	}
}
func (m Set[K]) EachIndex(fn func(int, K) bool) {
	i := 0
	for k := range m {
		if fn(i, k) {
			break
		}
		i++
	}
}
func (m Set[K]) Add(k K) bool {
	if m.Exists(k) {
		return false
	}
	m[k] = Nothing
	return true
}
func (m Set[K]) Remove(k K) bool {
	if m.Exists(k) {
		delete(m, k)
		return true
	}
	return false
}
func (m Set[K]) Put(k K) {
	m[k] = Nothing
}
func (m Set[K]) Delete(k K) {
	if m.Exists(k) {
		delete(m, k)
	}
}
func (m Set[K]) Len() int {
	return len(m)
}

type Entry[K comparable, V any] struct {
	Key K
	Val V
}
