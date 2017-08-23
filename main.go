
package main

import (
	"fmt"
	"strings"
)

type Ops interface{
}

type  Node struct {
    parent *Node 
	ops map[string]Ops
	// including type in name is redundent with strong typing
	// go likes CamelCase
   }
// func (a *Node) append(input string) {
// 	// should check for/strip punctuation if using Title
	
// 	to_append := strings.Title(input)
// 	// switch statement nothing fancy
// } 

type Root struct{
	dictionary map[string]string
	operations *Node

	//func newNode(parent *Node) *Node {(that's what it returns; a pointer to anode) 
		// return &Node{parent:parent}
}


type Write struct{
	op string
}

type Read struct{
	op string
}

type Delete struct{
	op string
}
// make start a method on root
	// 

// func recievers should always be pointers unless you have a compelling reason to use pass by value instead of pass by ref
// includes a function that 
	// 
// START ABORT and COMMIT are methods on ROOT
// READ/WRITE/DELETE are methods on ROOT

    func (r *root) read(k){
		display_value := ""
		element_value :=""
		for current_map[k] != nil{

		}
        until current_map[k] != nil {
            
                func(*element)Next{
					element_value =current_map[k]
				
            }
        }
	   
		display_value = element_value
		// include 

    func write (k,v){
        current_map[k]=v
    }

    func delete(k){
        delete(current_map,k)
    }

    func abort{
        // reset the head of the linked list to previous
    }

    func commit {
        for range current_map(k,v){
        &parent[k]=v
        }
    }

    func quit {
        quit = true
    }
}


func main (){
quit = false

while quit == false 
	
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
		
}