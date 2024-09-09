package engine

import (
	"bitbucket.org/shisheodev/on-demand-automation/engine/node"
	"context"
	"log"
	"sync"
)

func (i impl) ExecuteWorkflow(ctx context.Context, workflowID string) {
	nodes := getNodes(workflowID)

	doneTasks := make(map[string]bool)

	var wg sync.WaitGroup
	var mu sync.Mutex
	//stopChan := make(chan struct{})
	errChan := make(chan error, len(nodes)) // Channel to collect errors

	for _, n := range nodes {
		mu.Lock()
		doneTasks[n.ID] = false // default
		mu.Unlock()

		wg.Add(1)

		go func(n *node.Node) {
			defer wg.Done()

			//select {
			//case <-stopChan:
			//	return
			//default:
			//}

			// Execute the node's task and send any errors to the error channel
			if err := n.ExecuteTask(&node.TaskParams{
				DoneTasks:         doneTasks,
				Mu:                &mu,
				TerminateWorkflow: false,
			}); err != nil {
				errChan <- err

				if n.TerminateWorkflowOnFailure {
					log.Println("terminating workflow")
					//close(stopChan)
				}
			}
		}(&n)
	}

	go func() {
		wg.Wait()
		close(errChan) // Close the error channel after all tasks are done
	}()

	// Collect and handle errors
	for err := range errChan {
		if err != nil {
			// Handle the error appropriately (log, return, etc.)
			log.Printf("workflow execution failed: %v", err)
		}
	}

	log.Println("Execution complete!")
}

func getNodes(workflowID string) []node.Node {
	// TODO: retrieve workflow nodes from DB

	return []node.Node{
		{
			ID:           "node1",
			Stage:        "execution",
			Type:         "text",
			Dependencies: nil,
			NextNode: []node.NextNode{
				{
					NodeID: "node2",
				},
			},
			Config: node.Config{
				Text: &node.TextNodeConfig{EndpointURL: ""},
			},
		},
		{
			ID:    "node2",
			Stage: "execution",
			Type:  "plugin",
			Dependencies: []node.Dependency{
				{
					NodeID: "node1",
				},
			},
			NextNode: []node.NextNode{
				{
					NodeID: "node3",
				},
			},
			Config: node.Config{
				Plugin: &node.PluginNodeConfig{
					EndpointURL: "",
				},
			},
			TerminateWorkflowOnFailure: true,
		},
		{
			ID:    "node3",
			Stage: "execution",
			Type:  "model",
			Dependencies: []node.Dependency{
				{
					NodeID: "node2",
				},
			},
			NextNode: []node.NextNode{
				{
					NodeID: "",
				},
			},
			Config: node.Config{
				Model: &node.ModelNodeConfig{
					ModelEndpoint: "",
					Temperature:   0.7,
				},
			},
		},
	}
}
