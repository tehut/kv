package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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

type transaction struct {
	// change to transaction
	parent         *transaction
	transactionLog map[string]ops
}

type root struct {
	db           map[string]string
	transactions *transaction
	// change to transactions
}

// StevEx: What are your thoughts on using "this" as the reciever. I've seen some differing opinions on the internet
// also is the reciever a pointer to root? is that what r *root means? That r "is the type pointer to root" and when we call the method we are "passing"
// the currently instantiated struct?  My main point of confusion is that in newTransaction() we seem to be passing an argument but most other methods seem to pass
// ---------------------------------------------------------
// commands should be methods others can be functions

func repl() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	parseCommand(input)
}

func (r *root) read(key string) (string, error) {
	if r.db[key] != 0 int {
		return "", r.db[key]
	}
	displayValue := ""
	currentTransaction := r.transactions
	for {
		if currentTransaction.transactionLog[key] == write {
			displayValue = r.transactions[key].data
		} else {
			currentTransaction = currentTransaction.parent
			if currentTransaction == nil {
				return displayValue, errors.New("That key was not found")
			}
		}
		return displayValue, nil
	}

}
func parseCommand(input string) {

	input = strings.ToLower(input)
	command := strings.Split(input, " ")
	fmt.Println(command[1])
	switch command[0] {
	case "read":
		value, err :=read(command[1])
		if err != nil{
			fmt.Println(err)
		}else {
			fmt.Println(value)
		}
	case "write":
	case "delete":
		delete(r.db, command[1])
	case "start":
	case "abort":
	case "commit":
	}
}



func (parent *transaction) newTransaction() *transaction {
	return &transaction{parent: parent}
}

func (r *root) start() {
	r.transactions = newTransaction(r.transactions)
}

func (r *root) abort() error {
	if r.transactions == nil {
		return errors.New("You do not have a transaction open at this time")
	}

	r.transactions = r.transactions.parent
	return nil
}

func (r *root) commit() {
	// error if transactions is nil
	if r.transactions.parent == nil {
		// if parent is nil
		// loop through execute transactions on db
		for key, value := range r.transactions.transaction {
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
		// loop through transactions and copy to parent transactions map
		for key, value := range r.transactions.transaction {
			r.transactions.parent.transaction[key] = value
		}
	}

	// exit

}

func main() {
	// datastore := root{}
	var quit bool
	for {
		if quit == true {
			break
		}

		repl()

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

// create an empty transaction with a pointer to nill
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
