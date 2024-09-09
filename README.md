

// Approach 1

// load nodes to a queue

// start executing each node as goroutine

// Block goroutine execution until dependencies are met

// notify engine about completion of a node execution

// Approach 2

// send a node execution event to a redis queue

// worker picks up the message and execute that node.

// once the worker execution is complete, it publishes event mentioning the next node/nodes to be executed

// the entire execution will be asynchronous and will be marked as complete when the last node is executed
