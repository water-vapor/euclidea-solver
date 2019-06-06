package hashset

// Serializable is an interface represents Serializable objects
type Serializable interface {
	Serialize() interface{}
}

// HashSet is a hash set holds Serializable objects
type HashSet struct {
	m map[interface{}]interface{}
}

// NewHashSet return an empty hash set
func NewHashSet() *HashSet {
	m := make(map[interface{}]interface{})
	return &HashSet{m: m}
}

// Clone deep copies a hash set
func (hs *HashSet) Clone() *HashSet {
	ret := NewHashSet()
	for key, value := range hs.m {
		ret.m[key] = value
	}
	return ret
}

// Add adds an element
func (hs *HashSet) Add(key Serializable) {
	hs.m[key.Serialize()] = key
}

// Remove removes an element
func (hs *HashSet) Remove(key Serializable) {
	delete(hs.m, key.Serialize())
}

// Contains checks the existence of an element
func (hs *HashSet) Contains(key Serializable) bool {
	_, exists := hs.m[key.Serialize()]
	if exists {
		return true
	}
	return false
}

// Dict returns the dict for iteration
func (hs *HashSet) Dict() map[interface{}]interface{} {
	return hs.m
}

// Values returns a slice of all objects
func (hs *HashSet) Values() []interface{} {
	ret := make([]interface{}, len(hs.m))
	index := 0
	for _, elem := range hs.m {
		ret[index] = elem
		index++
	}
	return ret
}

// Empty checks whether the hash set is empty
func (hs *HashSet) Empty() bool {
	return len(hs.m) == 0
}
