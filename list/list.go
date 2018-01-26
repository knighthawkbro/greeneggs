package list

import "fmt"

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	data interface{} // Value of Node
	next *node       // Pointer to the next Node
	list *List       // Pointer to the list it is attached to
}

// nextNode (Private) - Returns the next node in the list
func (n node) nextNode() *node {
	// returns nil if there is not list AND if the pointer to the next
	// node is the same as the head's next node there for there is next node
	if next := n.next; n.list != nil && next != &n.list.head {
		return next
	}
	return nil
}

// List (Public) - The container for all the linked nodes in a set
type List struct {
	head node // the begining node
	size int  // size of the list
}

// init (Private) - Generates a linked list with Size=0 and head pointing to itself
func (l *List) init() *List {
	l.head.next = &l.head
	l.size = 0
	return l
}

// New (Public) - Returns an initialized list.
func New() *List { return new(List).init() }

// Size (Public) - Returns the length variable for the list as an integer
func (l *List) Size() int { return l.size }

// Add (Public) - Returns the node in a singly linked list, just adds to the front of the list
func (l *List) Add(v interface{}) error {
	if v == nil {
		return fmt.Errorf("Cannot add a nil value to set")
	}
	if l.Contains(v) {
		return fmt.Errorf("%v already exists in the set", v)
	}
	new := &node{data: v, list: l}
	prev := l.head.next
	l.head.next = new
	new.next = prev
	l.size++
	return nil
}

// Remove (Public) - Removes the first item on a list and returns it
func (l *List) Remove() interface{} {
	if l.size == 0 {
		return nil
	}
	removed := l.head.data
	l.head = *l.head.next
	l.size--
	return removed
}

// Get (Public) - Returns the first item list
func (l *List) Get() interface{} {
	if l.size == 0 {
		return nil
	}
	return l.head.data
}

// Contains (Public) - Returns true or false whether an item was contained in the list
func (l *List) Contains(i interface{}) bool {
	for current := l.head.next; current != nil; current = current.nextNode() {
		if current.data == i {
			return true
		}
	}
	return false
}

// String (Public) - Allows for the fmt.Print* functions to print the list struct as a string.
func (l *List) String() string {
	if l.size == 0 {
		return "[ ]"
	}
	result := "[ "
	for current := l.head.next; current != nil; current = current.nextNode() {
		result += fmt.Sprintf("%v ", current.data)
	}
	return result + "]"
}
