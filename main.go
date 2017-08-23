package main

import (
	"errors"
)

type ops interface {
}

type write struct {
	data string
}

// type read struct {
// }

type deleteCommand struct {
}

type node struct {
	// change to transaction
	parent *node
	ops    map[string]ops
}

type root struct {
	db         map[string]string
	operations *node
	// change to transactions
}

// StevEx: What are your thoughts on using "this" as the reciever. I've seen some differing opinions on the internet
// also is the reciever a pointer to root? is that what r *root means? That r "is the type pointer to root" and when we call the method we are "passing"
// the currently instantiated struct?  My main point of confusion is that in newNode() we seem to be passing an argument but most other methods seem to pass
// ---------------------------------------------------------
// commands should be methods others can be functions

func newNode(parent *node) *node {
	return &node{parent: parent}
}

func (r *root) start() {
	r.operations = newNode(r.operations)
}

func (r *root) abort() error {
	if r.operations == nil {
		return errors.New("You do not have a transaction open at this time")
	}

	r.operations = r.operations.parent
	return nil
}

func (r *root) commit() {
	// error if operations is nil
	if r.operations.parent == nil {
		// if parent is nil
		// loop through execute operations on db
		for key, value := range r.operations.ops {
			// realized I could have just done an if statement after I wrote it as a switch
			// kept it as a switch mostly to play with the syntax in go--which is almost identical to js
			switch value.(type) {
			case write:
				r.db[key] = value.(write).data
			case deleteCommand:
				delete(r.db, key)
			}
		}
	} else {
	// loop through operations and copy to parent operations map
		for key,value :range r.operations.ops {
			r.operations.parent.operations[key] = value
		}
	}

	// exit

}

// Psuedocode implementation
//     func (r *root) read(k){
// 		display_value := ""
// 		element_value :=""
// 		for current_map[k] != nil{

// 		}
//         until current_map[k] != nil {

//                 func(*element)Next{
// 					element_value =current_map[k]

//             }
//         }

// 		display_value = element_value
// 		// include

//     func write (k,v){
//         current_map[k]=v
//     }

//     func delete(k){
//         delete(current_map,k)
//     }

//     func abort{
//         // reset the head of the linked list to previous
//     }

//     func commit {
//         for range current_map(k,v){
//         &parent[k]=v
//         }
//     }

//     func quit {
//         quit = true
//     }
// }

// func main (){
// quit = false

// while quit == false

// create an empty node with a pointer to nill
// run command prompt
// downcase result
// switch statement
//     read
//     write
//     delete
//     start
//     commit
//     abort
// 	quit

// }
