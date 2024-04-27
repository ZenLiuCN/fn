package fn

import "sort"

// MapEx the object style map support for mappings
type MapEx[K, K1 key, V, V1 any] Map[K, V]

// MapExOfMap same as MapEx[K, K1, V, V1](m)
func MapExOfMap[K, K1 key, V, V1 any](m Map[K, V]) MapEx[K, K1, V, V1] {
	return MapEx[K, K1, V, V1](m)
}

// MapExOf  same as MapEx[K, K1, V, V1](m)
func MapExOf[K, K1 key, V, V1 any](m map[K]V) MapEx[K, K1, V, V1] {
	return m
}

// AsMap same as Map[K, V](m)
func (m MapEx[K, K1, V, V1]) AsMap() Map[K, V] {
	return Map[K, V](m)
}

// AsRaw same as map[K]V(m)
func (m MapEx[K, K1, V, V1]) AsRaw() map[K]V {
	return m
}

// MappingKey see MapKeyMapping
func (m MapEx[K, K1, V, V1]) MappingKey(fn func(K) K1) (r Map[K1, V]) {
	return MapKeyMapping(m, fn)
}

// MappingValue  see MapValueMapping
func (m MapEx[K, K1, V, V1]) MappingValue(fn func(V) V1) (r Map[K, V1]) {
	return MapValueMapping(m, fn)
}

// MappingValueTo see MapValueTo
func (m MapEx[K, K1, V, V1]) MappingValueTo(v V1) (r Map[K, V1]) {
	return MapValueTo(m, v)
}

// Len same as len(m)
func (m MapEx[K, K1, V, V1]) Len() int {
	return len(m)
}

// Values see MapValues
func (m MapEx[K, K1, V, V1]) Values() []V {
	return MapValues(m)
}

// Keys see MapKeys
func (m MapEx[K, K1, V, V1]) Keys() []K {
	return MapKeys(m)
}

// HasValue see MapHasValueBy
func (m MapEx[K, K1, V, V1]) HasValue(cmp func(V) bool) bool {
	return MapHasValueBy(m, cmp)
}

// HasKey see MapHasKey
func (m MapEx[K, K1, V, V1]) HasKey(k K) bool {
	return MapHasKey(m, k)
}

// KeyBy  see MapKeyBy
func (m MapEx[K, K1, V, V1]) KeyBy(pred func(V) bool) K {
	return MapKeyBy(m, pred)
}

// KeysBy  see MapKeysBy
func (m MapEx[K, K1, V, V1]) KeysBy(pred func(V) bool) []K {
	return MapKeysBy(m, pred)
}

// ValuesBy see 	return MapValuesBy(m, pred)
func (m MapEx[K, K1, V, V1]) ValuesBy(pred func(K) bool) []V {
	return MapValuesBy(m, pred)
}

// Entries see MapEntries
func (m MapEx[K, K1, V, V1]) Entries() []*Entry[K, V] {
	return MapEntries(m)
}

// Filter see MapFilter
func (m MapEx[K, K1, V, V1]) Filter(pred func(K, V) bool) []*Entry[K, V] {
	return MapFilter(m, pred)
}

// FilterSitu see MapFilterSitu
func (m MapEx[K, K1, V, V1]) FilterSitu(pred func(K, V) bool) {
	MapFilterSitu(m, pred)
}

// Pick see MapPick
func (m MapEx[K, K1, V, V1]) Pick(picker func(K, V) int) []map[K]V {
	return MapPick(m, picker)
}

// Each see MapEach
func (m MapEx[K, K1, V, V1]) Each(c func(K, V)) {
	MapEach(m, c)
}

// Sort see MapSortByValue
func (m MapEx[K, K1, V, V1]) Sort(lesser func(V, V) bool) []K {
	return MapSortByValue(m, lesser)
}

// SortStable see MapSortByValueStable
func (m MapEx[K, K1, V, V1]) SortStable(lesser func(V, V) bool) []K {
	return MapSortByValueStable(m, lesser)
}

// Map the object style map
type Map[K key, V any] map[K]V

// Len same as len(m)
func (m Map[K, V]) Len() int {
	return len(m)
}

// Values see MapValues
func (m Map[K, V]) Values() []V {
	return MapValues(m)
}

// Keys see MapKeys
func (m Map[K, V]) Keys() []K {
	return MapKeys(m)
}

// HasValue see MapHasValueBy
func (m Map[K, V]) HasValue(cmp func(V) bool) bool {
	return MapHasValueBy(m, cmp)
}

// HasKey see MapHasKey
func (m Map[K, V]) HasKey(k K) bool {
	return MapHasKey(m, k)
}

// KeyBy  see MapKeyBy
func (m Map[K, V]) KeyBy(pred func(V) bool) K {
	return MapKeyBy(m, pred)
}

// KeysBy  see MapKeysBy
func (m Map[K, V]) KeysBy(pred func(V) bool) []K {
	return MapKeysBy(m, pred)
}

// ValuesBy see 	return MapValuesBy(m, pred)
func (m Map[K, V]) ValuesBy(pred func(K) bool) []V {
	return MapValuesBy(m, pred)
}

// Entries see MapEntries
func (m Map[K, V]) Entries() []*Entry[K, V] {
	return MapEntries(m)
}

// Filter see MapFilter
func (m Map[K, V]) Filter(pred func(K, V) bool) []*Entry[K, V] {
	return MapFilter(m, pred)
}

// FilterSitu see MapFilterSitu
func (m Map[K, V]) FilterSitu(pred func(K, V) bool) {
	MapFilterSitu(m, pred)
}

// Pick see MapPick
func (m Map[K, V]) Pick(picker func(K, V) int) []map[K]V {
	return MapPick(m, picker)
}

// Each see MapEach
func (m Map[K, V]) Each(c func(K, V)) {
	MapEach(m, c)
}

// Sort see MapSortByValue
func (m Map[K, V]) Sort(lesser func(V, V) bool) []K {
	return MapSortByValue(m, lesser)
}

// SortStable see MapSortByValueStable
func (m Map[K, V]) SortStable(lesser func(V, V) bool) []K {
	return MapSortByValueStable(m, lesser)
}

// MapEach same as for k, v := range m
func MapEach[K key, V any, M ~map[K]V](m M, c func(K, V)) {
	for k, v := range m {
		c(k, v)
	}
}

// MapEntries fetch map to an Entry slice
func MapEntries[K key, V any, M ~map[K]V](m M) (r []*Entry[K, V]) {
	for k, v := range m {
		r = append(r, &Entry[K, V]{Key: k, Val: v})
	}
	return
}

// MapFilter pick subset of entry
func MapFilter[K key, V any, M ~map[K]V](m M, cmp func(K, V) bool) (r []*Entry[K, V]) {
	for k, v := range m {
		if cmp(k, v) {
			r = append(r, &Entry[K, V]{Key: k, Val: v})
		}
	}
	return
}

// MapKeys key slice of map
func MapKeys[K key, V any, M ~map[K]V](m M) (r []K) {
	for k := range m {
		r = append(r, k)
	}
	return
}

// MapHasKey check key is exists, just _, ok = m[k]
func MapHasKey[K key, V any, M ~map[K]V](m M, k K) (ok bool) {
	_, ok = m[k]
	return
}

// MapKeyOf found key of a value otherwise the K's default value returned
func MapKeyOf[K key, V key, M ~map[K]V](m M, v V) (k K) {
	for k1, val := range m {
		if val == v {
			return k1
		}
	}
	return
}

// MapKeysOf found keys of a value
func MapKeysOf[K key, V key, M ~map[K]V](m M, v V) (k []K) {
	for k1, val := range m {
		if val == v {
			k = append(k, k1)
		}
	}
	return
}

// MapKeyBy found key by predicate
func MapKeyBy[K key, V any, M ~map[K]V](m M, p func(V) bool) (k K) {
	for k1, val := range m {
		if p(val) {
			return k1
		}
	}
	return
}

// MapKeysBy pick keys by value predicate
func MapKeysBy[K key, V any, M ~map[K]V](m M, p func(V) bool) (k []K) {
	for k1, val := range m {
		if p(val) {
			k = append(k, k1)
		}
	}
	return
}

// MapKeyMapping mapping key
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

// MapValues value slice of map
func MapValues[K key, V any, M ~map[K]V](m M) (r []V) {
	for _, v := range m {
		r = append(r, v)
	}
	return
}

// MapValueMapping mapping value
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

// MapValueTo make map value to v
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

// MapValuesBy picks values
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

// MapHasValue check value exists
func MapHasValue[K key, V key, M ~map[K]V](m M, v V) (ok bool) {
	for _, val := range m {
		if val == v {
			return true
		}
	}
	return
}

// MapHasValueBy check value exists by cmp
func MapHasValueBy[K key, V any, M ~map[K]V](m M, cmp func(V) bool) (ok bool) {
	for _, val := range m {
		if cmp(val) {
			return true
		}
	}
	return
}

// MapFilterSitu filter and remove entries (not create new ones)
func MapFilterSitu[K key, V any, M ~map[K]V](m M, pred func(K, V) bool) {
	for k, v := range m {
		if !pred(k, v) {
			delete(m, k)
		}
	}
	return
}

// MapPick in to new maps, pick returns: negative for discard, otherwise for put map N
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

// MapSortByValue sort by value and returns the sorted key slice, see sort.Slice
func MapSortByValue[K key, V any, M ~map[K]V](m M, cmp func(V, V) bool) (r []K) {
	r = append(r, MapKeys(m)...)
	sort.Slice(r, func(i, j int) bool {
		return cmp(m[r[i]], m[r[j]])
	})
	return
}

// MapSortByValueStable sort by value and returns the sorted key slice, see sort.SliceStable
func MapSortByValueStable[K key, V any, M ~map[K]V](m M, cmp func(V, V) bool) (r []K) {
	r = append(r, MapKeys(m)...)
	sort.SliceStable(r, func(i, j int) bool {
		return cmp(m[r[i]], m[r[j]])
	})
	return
}

// MapClear by loop on each key
func MapClear[K key, V any, M ~map[K]V](m M) {
	for k := range m {
		delete(m, k)
	}
}

// Entry hold a key value pair (key must comparable)
type Entry[K key, V any] struct {
	Key K
	Val V
}
