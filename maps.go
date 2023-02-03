package fn

import "sort"

//Map the object style map
type Map[K key, V any] map[K]V

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Values() []V {
	return MapValues(m)
}
func (m Map[K, V]) Keys() []K {
	return MapKeys(m)
}

func (m Map[K, V]) HasValue(cmp func(V) bool) bool {
	return MapHasValueBy(m, cmp)
}
func (m Map[K, V]) HasKey(k K) bool {
	return MapHasKey(m, k)
}

func (m Map[K, V]) KeyBy(pred func(V) bool) K {
	return MapKeyBy(m, pred)
}
func (m Map[K, V]) KeysBy(pred func(V) bool) []K {
	return MapKeysBy(m, pred)
}

func (m Map[K, V]) ValuesBy(pred func(K) bool) []V {
	return MapValuesBy(m, pred)
}

func (m Map[K, V]) Entries() []*Entry[K, V] {
	return MapEntries(m)
}

func (m Map[K, V]) Filter(pred func(K, V) bool) []*Entry[K, V] {
	return MapFilter(m, pred)
}

//FilterSitu filter and remove
func (m Map[K, V]) FilterSitu(pred func(K, V) bool) {
	MapFilterSitu(m, pred)
}

//Pick picker returns: negative for discard, otherwise for put map N
func (m Map[K, V]) Pick(picker func(K, V) int) []map[K]V {
	return MapPick(m, picker)
}

func (m Map[K, V]) Each(c func(K, V)) {
	MapEach(m, c)
}

func (m Map[K, V]) Sort(lesser func(V, V) bool) []K {
	return MapSortByValue(m, lesser)
}

func (m Map[K, V]) SortStable(lesser func(V, V) bool) []K {
	return MapSortByValueStable(m, lesser)
}

//MapEach consume each key and value
func MapEach[K key, V any, M ~map[K]V](m M, c func(K, V)) {
	for k, v := range m {
		c(k, v)
	}
}

//MapEntries map to Entry slice
func MapEntries[K key, V any, M ~map[K]V](m M) (r []*Entry[K, V]) {
	for k, v := range m {
		r = append(r, &Entry[K, V]{k, v})
	}
	return
}

//MapFilter pick subset of entry
func MapFilter[K key, V any, M ~map[K]V](m M, cmp func(K, V) bool) (r []*Entry[K, V]) {
	for k, v := range m {
		if cmp(k, v) {
			r = append(r, &Entry[K, V]{k, v})
		}
	}
	return
}

//MapKeys key slice of map
func MapKeys[K key, V any, M ~map[K]V](m M) (r []K) {
	for k := range m {
		r = append(r, k)
	}
	return
}

//MapHasKey check key is exists
func MapHasKey[K key, V any, M ~map[K]V](m M, k K) (ok bool) {
	_, ok = m[k]
	return
}

//MapKeyOf found key of a value otherwise the K's default value returned
func MapKeyOf[K key, V key, M ~map[K]V](m M, v V) (k K) {
	for k1, val := range m {
		if val == v {
			return k1
		}
	}
	return
}

//MapKeysOf found keys of a value
func MapKeysOf[K key, V key, M ~map[K]V](m M, v V) (k []K) {
	for k1, val := range m {
		if val == v {
			k = append(k, k1)
		}
	}
	return
}

//MapKeyBy found key by predicate
func MapKeyBy[K key, V any, M ~map[K]V](m M, p func(V) bool) (k K) {
	for k1, val := range m {
		if p(val) {
			return k1
		}
	}
	return
}

//MapKeysBy pick keys by value predicate
func MapKeysBy[K key, V any, M ~map[K]V](m M, p func(V) bool) (k []K) {
	for k1, val := range m {
		if p(val) {
			k = append(k, k1)
		}
	}
	return
}

//MapKeyMapping mapping key
func MapKeyMapping[K, K1 key, V any, M ~map[K]V](m M, fn func(K) K1) (o map[K1]V) {
	if m == nil || len(m) == 0 {
		return
	}
	o = make(map[K1]V, len(m))
	for k, val := range m {
		o[fn(k)] = val
	}
	return
}

//MapValues value slice of map
func MapValues[K key, V any, M ~map[K]V](m M) (r []V) {
	for _, v := range m {
		r = append(r, v)
	}
	return
}

//MapValueMapping mapping value
func MapValueMapping[K key, V, V1 any, M ~map[K]V](m M, fn func(V) V1) (o map[K]V1) {
	if m == nil || len(m) == 0 {
		return
	}
	o = make(map[K]V1, len(m))
	for k, val := range m {
		o[k] = fn(val)
	}
	return
}

//MapValueTo make map value to v
func MapValueTo[K key, V, V1 any, M ~map[K]V](m M, v V1) (o map[K]V1) {
	if m == nil || len(m) == 0 {
		return
	}
	o = make(map[K]V1, len(m))
	for k := range m {
		o[k] = v
	}
	return
}

//MapValuesBy picks values
func MapValuesBy[K key, V any, M ~map[K]V](m M, fn func(K) bool) (o []V) {
	if m == nil || len(m) == 0 {
		return
	}
	for k, v := range m {
		if fn(k) {
			o = append(o, v)
		}
	}
	return
}

//MapHasValue check value exists
func MapHasValue[K key, V key, M ~map[K]V](m M, v V) (ok bool) {
	for _, val := range m {
		if val == v {
			return true
		}
	}
	return
}

//MapHasValueBy check value exists by cmp
func MapHasValueBy[K key, V any, M ~map[K]V](m M, cmp func(V) bool) (ok bool) {
	for _, val := range m {
		if cmp(val) {
			return true
		}
	}
	return
}

//MapFilterSitu filter and remove entries (not create new ones)
func MapFilterSitu[K key, V any, M ~map[K]V](m M, pred func(K, V) bool) {
	for k, v := range m {
		if !pred(k, v) {
			delete(m, k)
		}
	}
	return
}

//MapPick in to new maps, pick returns: negative for discard, otherwise for put map N
func MapPick[K key, V any, M ~map[K]V](m M, pick func(K, V) int) (r []map[K]V) {
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

//MapSortByValue sort by value and returns the sorted key slice, see sort.Slice
func MapSortByValue[K key, V any, M ~map[K]V](m M, cmp func(V, V) bool) (r []K) {
	r = append(r, MapKeys(m)...)
	sort.Slice(r, func(i, j int) bool {
		return cmp(m[r[i]], m[r[j]])
	})
	return
}

//MapSortByValueStable sort by value and returns the sorted key slice, see sort.SliceStable
func MapSortByValueStable[K key, V any, M ~map[K]V](m M, cmp func(V, V) bool) (r []K) {
	r = append(r, MapKeys(m)...)
	sort.SliceStable(r, func(i, j int) bool {
		return cmp(m[r[i]], m[r[j]])
	})
	return
}

type Entry[K key, V any] struct {
	Key K
	Val V
}
