package fn

// Set interface
// FIXME: current golang (include 1.20) not treat interface as comparable
type Set[K comparable] interface {
	Clear()                      //Clear remove all values
	Values() []K                 //Values as a slice
	Each(func(K) bool)           //Each loop for value
	EachIndex(func(int, K) bool) //EachIndex loop with index and value
	Add(K) bool                  //Add with exists check,return true if success.
	Remove(K) bool               //Remove delete from underlying map with exists check,return true if success.
	Put(K)                       //Put Add without existence check
	Delete(K)                    //Delete without existence check.
	Len() int                    //Len the len.
	Exists(K) bool               //Exists check exists
}

// HashSet Set impl with map[K]struct{}
// FIXME: current golang (include 1.20) not treat interface as comparable
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

// Exists a value in set
func (m HashSet[K]) Exists(k K) (ok bool) {
	_, ok = m[k]
	return
}

// Clear all values
func (m HashSet[K]) Clear() {
	for k := range m {
		delete(m, k)
	}
}

// Values as slice
func (m HashSet[K]) Values() []K {
	return MapKeys(m)
}

// Each action on each value
func (m HashSet[K]) Each(fn func(K) bool) {
	for k := range m {
		if fn(k) {
			break
		}
	}
}

// EachIndex action on each index and value
func (m HashSet[K]) EachIndex(fn func(int, K) bool) {
	i := 0
	for k := range m {
		if fn(i, k) {
			break
		}
		i++
	}
}

// Add a value when it not exists
func (m HashSet[K]) Add(k K) bool {
	if m.Exists(k) {
		return false
	}
	m[k] = Nothing
	return true
}

// Remove a value
func (m HashSet[K]) Remove(k K) bool {
	if m.Exists(k) {
		delete(m, k)
		return true
	}
	return false
}

// Put a value, override old one if present
func (m HashSet[K]) Put(k K) {
	m[k] = Nothing
}

// Delete one value
func (m HashSet[K]) Delete(k K) {
	if m.Exists(k) {
		delete(m, k)
	}
}

// Len of set
func (m HashSet[K]) Len() int {
	return len(m)
}
