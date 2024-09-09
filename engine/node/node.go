package node

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type TaskParams struct {
	DoneTasks         map[string]bool
	Mu                *sync.Mutex
	TerminateWorkflow bool
}

func (n *Node) ExecuteTask(params *TaskParams) error {

	for {
		//log.Println("waiting for dependency...")
		done := n.canExecute(params)
		if done {
			log.Println("dependencies satisfied for node: ", n.ID)
			break
		}
	}

	var err error
	switch n.Type {
	case TypeText:
		err = n.executeTextNode()
	case TypeModel:
		err = n.executeModelNode()
	case TypePlugin:
		err = n.executePluginNode()
	default:
		err = fmt.Errorf("unknown node type: %v", n.Type)
	}

	if err != nil {
		log.Println("encountered error in node: ", n.ID)
		if n.TerminateWorkflowOnFailure {
			params.Mu.Lock()
			params.TerminateWorkflow = true
			params.Mu.Unlock()
		}
	}

	params.Mu.Lock()
	params.DoneTasks[n.ID] = true
	params.Mu.Unlock()

	return err
}

func (n *Node) canExecute(params *TaskParams) bool {

	time.Sleep(1 * time.Second)
	params.Mu.Lock()
	if params.TerminateWorkflow {
		log.Println("cancelling node execution: ", n.ID)
		params.Mu.Unlock()
		return false
	}
	params.Mu.Unlock()

	if n.Dependencies == nil {
		return true
	}

	dependencies := n.Dependencies

	dependencyCount := len(dependencies)
	doneCount := 0

	for _, dep := range dependencies {

		params.Mu.Lock()
		done := params.DoneTasks[dep.NodeID]
		params.Mu.Unlock()

		if done {
			doneCount++
		}
	}

	if doneCount == dependencyCount {
		return true
	}

	return false
}

func (n *Node) executeTextNode() error {
	log.Println("executing text node task: ", n.ID)

	return nil
}

func (n *Node) executeModelNode() error {
	log.Println("executing model node task: ", n.ID)

	return nil
}

func (n *Node) executePluginNode() error {
	log.Println("executing plugin node task: ", n.ID)

	return errors.New("failed")
}
