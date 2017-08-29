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

type deleteCommand struct {
}

type transaction struct {
	parent         *transaction
	transactionLog map[string]ops
}

type root struct {
	db           map[string]string
	transactions *transaction
}

// ---------------------------------------------------------------------------
func repl(root *root) bool {
	var quit bool
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	root.parseCommand(input)
	return quit
}

func newRoot() root {
	return root{transactions: nil, db: make(map[string]string)}
}

func newTransaction(passedtransaction *transaction) transaction {
	var t transaction
	if passedtransaction != nil {
		t = transaction{parent: passedtransaction, transactionLog: make(map[string]ops)}
	} else {
		t = transaction{parent: nil, transactionLog: make(map[string]ops)}
		return t
	}
	return t
}

// ---------------------------------------------------------------------------

func (r *root) write(key string, value string) {
	if r.transactions != nil {
		r.transactions.transactionLog[key] = write{data: value}
	} else {
		r.db[key] = value
	}
}

func (r *root) deleteFunction(key string) {
	if r.transactions != nil {
		r.transactions.transactionLog[key] = deleteCommand{}
	} else {
		delete(r.db, key)
	}
}

func (r *root) read(key string) (string, error) {
	currentTransaction := r.transactions
	for currentTransaction != nil {
		if _, ok := currentTransaction.transactionLog[key]; ok {
			switch currentTransaction.transactionLog[key].(type) {
			default:
				return "", errors.New("Do not recognize command, delete and write are recognized transactions")

			case deleteCommand:
				return "", errors.New("We have deleted that key")
			case write:
				return currentTransaction.transactionLog[key].(write).data, nil
				// fmt.Println(displayValue, "  this is the data in the write struct")

			}

		} else {
			currentTransaction = currentTransaction.parent
		}

	}
	// I'm still using the pointer that's a member of the root struct here because I'm not sure how scope works
	// in Go.  Will currentTransaction be the original value when we drop out of the loop or will it still be the
	// last assigned value? I'd expect that it would be the last assigned value.

	if value, ok := r.db[key]; ok {
		return value, nil
	}
	return "", errors.New("That key was not found")

}

func (r *root) start() {
	var newestTransaction transaction
	if r.transactions != nil {
		newestTransaction = newTransaction(r.transactions)

	} else {
		newestTransaction = newTransaction(nil)
	}
	r.transactions = &newestTransaction
}

func (r *root) abort() error {
	if r.transactions == nil {
		return errors.New("You do not have a transaction open at this time")
	}

	r.transactions = r.transactions.parent
	return nil
}

func (r *root) commit() {
	if r.transactions == nil {
		return
	}
	if r.transactions.parent == nil {
		for key, value := range r.transactions.transactionLog {
			switch value.(type) {
			case write:
				r.db[key] = value.(write).data
			case deleteCommand:
				delete(r.db, key)
			}
			// I've already confirmed that there is no open parent transaction so at this point I reset the head to nil
			r.transactions = nil
		}
	} else {
		for key, value := range r.transactions.transactionLog {
			r.transactions.parent.transactionLog[key] = value
		}
		// I've already checked that r.transactions.parent exists and now im resetting the head
		r.transactions = r.transactions.parent

	}
	// exit
	// sometimes I commit and it doesn't exit!
}

func (r *root) parseCommand(input string) {
	// only to string command[0]
	input = strings.ToLower(input)
	command := strings.Fields(input)
	// fmt.Println(command[1])
	switch command[0] {
	default:
		fmt.Println("That is not a valid command")
	case "read":
		if len(command) == 2 {
			value, err := r.read(command[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		} else {
			fmt.Println("That is not a valid command")
		}
	case "write":
		if len(command) == 3 {
			r.write(command[1], command[2])
		} else {
			err := errors.New("You have not supplied enough arguments for the WRITE command")
			fmt.Println(err)
		}
	case "delete":
		if len(command) == 2 {
			r.deleteFunction(command[1])
		} else {
			fmt.Println("That is not a valid command")
		}
	case "start":
		if len(command) != 1 {
			fmt.Println("that is not a valid command")
		} else {
			r.start()
		}
	case "abort":
		if len(command) != 1 {
			fmt.Println("that is not a valid command")
		} else {
			r.abort()
		}
	case "commit":
		if len(command) != 1 {
			fmt.Println("that is not a valid command")
		} else {
			r.commit()
		}
	case "quit":
		if len(command) != 1 {
			fmt.Println("that is not a valid command")
		} else {
			fmt.Println("Exiting")
			os.Exit(0)
		}
	}

}

func main() {
	datastore := newRoot()

	for {
		repl(&datastore)

	}
}
