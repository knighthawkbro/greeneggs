package array

import (
	"fmt"
)

// Array (Public) - Structure that defines
type Array struct {
	size       int
	collection []interface{}
}

// Init (Public) - initializes the array with whatever size is provided, This is what can be overrided by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]interface{}, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *Array { return new(Array).Init(10) }

// Add (Public) - Returns an error if adding the item failed, like a bool return, just with more information when it fails
func (a *Array) Add(item interface{}) error {
	if item == nil {
		return fmt.Errorf("Cannot add a nil item to set")
	}
	if a.Contains(item) {
		return fmt.Errorf("%v already exists in the set", item)
	}
	a.ensureSpace()
	a.collection[a.size] = item
	a.size++
	return nil
}

// Remove (Public) - Removes the last item in the array, at constant speed.
func (a *Array) Remove() interface{} {
	if a.size == 0 {
		return nil
	}
	removed := a.collection[a.size-1]
	a.collection[a.size-1] = nil
	a.size--
	return removed
}

// Get (Public) - Gets the last item inserted into the array, Constant speed
func (a *Array) Get() interface{} {
	if a.size == 0 {
		return nil
	}
	return a.collection[a.size-1]
}

// Contains (Public) - traverses the array and returns true if the item is found, linear speed.
func (a *Array) Contains(item interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.collection[i] == item {
			return true
		}
	}
	return false
}

// Size (Public) - returns the size of the array.
func (a *Array) Size() int {
	return a.size
}

// String (Public) - For methods that want to call the string version of
// this set.
func (a *Array) String() string {
	result := "[ "
	for i := 0; i < a.size; i++ {
		result += fmt.Sprintf("%v ", a.collection[i])
	}
	return result + "]"
}

// ensureSpace (Private) - Sees if the size and capacity of the array are the same. If so,
// It creates a new array with double the capacity and overwrites the old array with a new
// array, then clears the new array for the GC.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}
