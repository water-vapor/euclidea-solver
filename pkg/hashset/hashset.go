package hashset

// An interface represents Serializable objects
type Serializable interface {
	Serialize() interface{}
}

// A hash set holds Serializable objects
type HashSet struct {
	m map[interface{}]interface{}
}

func NewHashSet() *HashSet {
	m := make(map[interface{}]interface{})
	return &HashSet{m: m}
}

// Deep copy a hash set
func (hs *HashSet) Clone() *HashSet {
	ret := NewHashSet()
	for key, value := range hs.m {
		ret.m[key] = value
	}
	return ret
}

func (hs *HashSet) Add(key Serializable) {
	hs.m[key.Serialize()] = key
}

func (hs *HashSet) Remove(key Serializable) {
	delete(hs.m, key.Serialize())
}

func (hs *HashSet) Contains(key Serializable) bool {
	_, exists := hs.m[key.Serialize()]
	if exists {
		return true
	}
	return false
}

// Returns the dict for iteration
func (hs *HashSet) Dict() map[interface{}]interface{} {
	return hs.m
}

// Returns a slice of all objects
func (hs *HashSet) Values() []interface{} {
	ret := make([]interface{}, len(hs.m))
	index := 0
	for _, elem := range hs.m {
		ret[index] = elem
		index++
	}
	return ret
}

func (hs *HashSet) Empty() bool {
	return len(hs.m) == 0
}
