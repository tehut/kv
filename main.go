
package main

import (
	"fmt"
	"strings"
)

type Ops interface{
}

type Write struct{
    data string
}

type Read struct{
}

type Delete struct{
}


type  Node struct {
    parent *Node 
	ops map[string]Ops
   }


type Root struct{
	dictionary map[string]string
	operations *Node

}

// StevEx: What are your thoughts on using "this" as the reciever. I've seen some differing opinions on the internet
        // also is the reciever a pointer to root? is that what r *root means? That r "is the type pointer to root" and when we call the method we are "passing"
        // the currently instantiated struct?  My main point of confusion is that in newNode() we seem to be passing an argument but most other methods seem to pass 
// ---------------------------------------------------------

func (r *root) newNode(parent) *Node {
        return &Node{parent:parent}
}


func (r *root) start() *Node {
    var parent *Node
        if operations == nil {
            parent = *r
        }else {parent = operations}

        newOperations := r.newNode(operations)
        r.operations = newOperations
}

func (r *root) fetchAncestor() *Node{
     ancestorNode := &operations.parent
     return ancestorNode
}

func (r *root) abort() *Node {
    if operations == nil{
        fmt.Println("You do not have a transaction open at this time")
    }else {
        r.operations = fetchAncestor()
    }
}

func (r *root) commit() map[string]string{
    currentNodeMap = r.operations
    for key,value :range currentNodeMap {
        // realized I could have just done an if statement after I wrote it as a switch
        // kept it as a switch mostly to play with the syntax in go--which is almost identical to js
        switch value {
            case Write struct:
            // have not checked if go allows for overwriting maps or if we should just implem
            dictionary[key] = value.data
            case Delete struct:
            delete(dictionary,key) 
        }
    }
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
		
}