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
func repl(root root) {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	root.parseCommand(input)
}

func newRoot() root {
	return root{transactions: nil, db: make(map[string]string)}
}

func (r *root) write(key string, value string) {
	if r.transactions != nil {
		r.transactions.transactionLog[key] = write{data: value}
	} else {
		r.db[key] = value
	}
}

func (r *root) read(key string) (string, error) {
	displayValue := ""

	currentTransaction := r.transactions

	for currentTransaction != nil {
		if _, ok := currentTransaction.transactionLog[key]; ok {
			switch currentTransaction.transactionLog[key].(type) {
			case write:
				displayValue = currentTransaction.transactionLog[key].(write).data
				fmt.Println(displayValue, "  this is the data in the write struct")
				return displayValue, nil
			}

		} else {
			if currentTransaction.parent != nil {
				currentTransaction = currentTransaction.parent
			}
		}

	}
	// I'm still using the pointer that's a member of the root struct here because I'm not sure how scope works
	// in Go.  Will currentTransaction be the original value when we drop out of the loop or will it still be the
	// last assigned value? I'd expect that it would be the last assigned value.

	if r.transactions == nil {
		if value, ok := r.db[key]; ok {
			return value, nil
		}
		return displayValue, errors.New("That key was not found")
	}
	return displayValue, nil
}

// }
func (r *root) parseCommand(input string) {
	// only to string command[0]
	input = strings.ToLower(input)
	command := strings.Fields(input)
	// fmt.Println(command[1])
	switch command[0] {
	case "read":
		value, err := r.read(command[1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	case "write":
		fmt.Println(len(command))
		if len(command) == 3 {
			r.write(command[1], command[2])
		} else {
			err := errors.New("You have not supplied enough arguments for the WRITE command")
			fmt.Println(err)
		}
	case "delete":
		delete(r.db, command[1])
	case "start":
		r.start()
		fmt.Println(command[0])
	case "abort":
		r.abort()
	case "commit":
		r.commit()
	}
}

func (r *root) newTransaction(passedtransaction *transaction) transaction {
	var t transaction
	if passedtransaction != nil {
		t = transaction{parent: passedtransaction}
	} else {
		t = transaction{parent: nil}
		return t
	}
	return t
}

func (r *root) start() {
	var newestTransaction transaction
	if r.transactions != nil {
		newestTransaction = r.newTransaction(r.transactions)

	} else {
		newestTransaction = r.newTransaction(nil)
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
	if r.transactions.parent == nil {
		for key, value := range r.transactions.transactionLog {
			switch value.(type) {
			case write:
				r.db[key] = value.(write).data
			case deleteCommand:
				delete(r.db, key)
			}
		}
	} else {
		for key, value := range r.transactions.transactionLog {
			r.transactions.parent.transactionLog[key] = value
		}
		// reset the pointer to the parent or nil.
		// if there are no open transactions set transaction pointer to nil
		// if transactions are open set pointer to parent
		// check for open transactions on root.transactions
	}
	// exit
}

func main() {
	datastore := newRoot()

	var quit bool
	for {
		if quit == true {
			break
		}
		repl(datastore)
	}
}
