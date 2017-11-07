# K/V Repl
This is a small command line REPL (read-eval-print loop) that drives a simple in-memory key/value storage system.
This system should also allow for nested transactions. A transaction can then be committed or aborted.


To "play" add and remove key value pairs with the commands below:

COMMANDS

READ `<key>` Reads and prints, to stdout, the val associated with key. If the value is not present an error is printed to stderr.<br>
WRITE `<key> <val>` Stores val in key.<br>
DELETE `<key>` Removes all key from store. Future READ commands on that key will return an error.<br>
START Start a transaction.<br>
COMMIT Commit a transaction. All actions in the current transaction are committed to the parent transaction or the root store. If there is no current transaction an error is output to stderr.<br>
ABORT Abort a transaction. All actions in the current transaction are discarded.<br>
QUIT Exit the REPL cleanly. A message to stderr may be output.<br>
