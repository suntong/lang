package main

import (
	"fmt"
	// For Go 1.21+
	"maps"
	"slices"
	"sync"
)

// Relation represents a parsed foreign key relationship.
type Relation struct {
	FkTable             string
	FkConstraintName    string
	FkColumns           []string
	ReferencedTable     string
	ReferencedPkColumns []string // Columns in the referenced table targeted by the FK
}

// LabelSet represents the set of labels for a node (table).
// It's a map where keys are label strings and values are empty structs for set semantics.
type LabelSet map[string]struct{}

// LPA encapsulates the state and logic for the Label Propagation Algorithm.
type LPA struct {
	// graph stores the directed graph where a key node is influenced by nodes in its value slice.
	// Specifically, FkTable -> ReferencedTable means FkTable is influenced by ReferencedTable.
	// So, graph[FkTable] will contain ReferencedTable as an influencer.
	graph map[string][]string

	// nodes is a sorted list of all unique table names. This primarily aids in
	// deterministic initialization of goroutines if needed, though the parallel
	// nature makes strict order less critical for computation itself.
	nodes []string

	// currentLabels holds the set of labels for each node.
	// Access to currentLabels is protected by mu.
	currentLabels map[string]LabelSet
	mu            sync.RWMutex
}

// NewLPA creates and initializes a new LPA instance from a list of relations.
// It builds the graph structure and sets initial labels (each node gets its own name as a label).
func NewLPA(relations []Relation) (*LPA, error) {
	adj := make(map[string][]string)
	nodeSet := make(map[string]struct{}) // To collect unique node names
	initialLabels := make(map[string]LabelSet)

	if relations == nil {
		// Handle nil input gracefully by returning an empty LPA state.
		// Depending on strictness, an error could be returned instead:
		// return nil, fmt.Errorf("input relations cannot be nil")
		return &LPA{
			graph:         adj,
			nodes:         []string{},
			currentLabels: initialLabels,
		}, nil
	}

	for _, rel := range relations {
		if rel.FkTable == "" || rel.ReferencedTable == "" {
			// validation for empty table names
			fmt.Printf("Warning: Relation contains empty FkTable ('%s') or ReferencedTable ('%s')\n", rel.FkTable, rel.ReferencedTable)
		}

		// Add both tables to the set of known nodes
		nodeSet[rel.FkTable] = struct{}{}
		nodeSet[rel.ReferencedTable] = struct{}{}

		// The FkTable is influenced by the ReferencedTable.
		// Add ReferencedTable to the list of influencers for FkTable.
		adj[rel.FkTable] = append(adj[rel.FkTable], rel.ReferencedTable)
	}

	// Create a sorted list of nodes. Sorting helps in producing deterministic
	// outputs if the order of processing matters in any part (e.g., logging, testing).
	nodeList := make([]string, 0, len(nodeSet))
	for nodeName := range nodeSet {
		nodeList = append(nodeList, nodeName)
	}
	slices.Sort(nodeList) // Sorts in-place. Requires Go 1.21+.

	// Initialize labels: each node starts with its own name as its only label.
	// Also, ensure all nodes (even those only appearing as ReferencedTable)
	// exist in the adjacency list, potentially with an empty list of influencers.
	for _, nodeName := range nodeList {
		initialLabels[nodeName] = LabelSet{nodeName: {}}
		if _, ok := adj[nodeName]; !ok {
			// This node might only be a ReferencedTable and not an FkTable for any relation,
			// or it's an isolated table if it was mentioned but not part of any processed relation.
			// It has no influencers based on the FK definitions.
			adj[nodeName] = []string{}
		}
	}

	return &LPA{
		graph:         adj,
		nodes:         nodeList,
		currentLabels: initialLabels,
	}, nil
}

// Run executes the Overlapping Label Propagation Algorithm.
// It iterates up to maxIterations or until labels converge.
// Goroutines are used to compute label updates for nodes in parallel within each iteration.
// This implementation uses a synchronous update model for labels: all nodes' labels
// for iteration `k` are based on labels from iteration `k-1`.
// The overlapping nature (union of labels) means labels are only added,
// which makes the algorithm inherently stable and less prone to oscillations
// compared to majority-voting LPAs.
func (l *LPA) Run(maxIterations int) (map[string]LabelSet, int) {
	if len(l.nodes) == 0 {
		return make(map[string]LabelSet), 0
	}

	for iter := 0; iter < maxIterations; iter++ {
		changedInIteration := false

		// proposedLabelUpdates will store the newly computed LabelSets for each node in this iteration.
		// Access to this map from goroutines must be synchronized.
		proposedLabelUpdates := make(map[string]LabelSet, len(l.nodes))
		var proposedMu sync.Mutex

		var wg sync.WaitGroup

		// Create a snapshot of currentLabels for all goroutines to read from.
		// This ensures that all nodes in this iteration base their updates on the
		// state from the end of the previous iteration (or initial state for iter 0).
		// This is key for the "synchronous parallel" update style.
		l.mu.RLock()
		labelsSnapshot := make(map[string]LabelSet, len(l.currentLabels))
		for node, ls := range l.currentLabels {
			labelsSnapshot[node] = cloneLabelSet(ls) // Deep copy each LabelSet
		}
		l.mu.RUnlock()

		for _, nodeName := range l.nodes { // Iterate over a consistent list of nodes
			wg.Add(1)
			go func(currentNodeName string) {
				defer wg.Done()

				// Start with the node's own inherent label. This ensures it's always part of its set.
				newNodeLabels := LabelSet{currentNodeName: {}}

				// Get the list of nodes that influence the currentNode.
				// l.graph is read-only after LPA initialization, so no lock needed here.
				influencers := l.graph[currentNodeName]

				for _, influencerName := range influencers {
					// Get the labels of the influencer from the snapshot of the previous iteration's state.
					if influencerLabels, ok := labelsSnapshot[influencerName]; ok {
						for label := range influencerLabels {
							newNodeLabels[label] = struct{}{} // Add label to the set (union operation)
						}
					}
				}

				// Safely store the computed labels for this node.
				proposedMu.Lock()
				proposedLabelUpdates[currentNodeName] = newNodeLabels
				proposedMu.Unlock()
			}(nodeName)
		}
		wg.Wait() // Wait for all goroutines to compute their proposed labels.

		// Apply the proposed updates to currentLabels and check for convergence.
		l.mu.Lock()
		for nodeName, newLabels := range proposedLabelUpdates {
			// Compare newLabels with the actual l.currentLabels[nodeName] (not the snapshot)
			// to determine if a change occurred in this iteration.
			if !areLabelSetsEqual(l.currentLabels[nodeName], newLabels) {
				l.currentLabels[nodeName] = newLabels // Update the main label set for the next iteration
				changedInIteration = true
			}
		}
		l.mu.Unlock()

		if !changedInIteration {
			// fmt.Printf("LPA converged after %d iterations.\n", iter+1)
			return l.GetLabelsCopy(), iter + 1 // Converged. iter+1 because iter is 0-indexed.
		}
	}

	fmt.Printf("LPA reached max iterations (%d).\n", maxIterations)
	return l.GetLabelsCopy(), maxIterations // Max iterations reached
}

// GetLabelsCopy returns a deep copy of the current labels.
// This is safe for callers to read or modify without affecting the LPA's internal state.
func (l *LPA) GetLabelsCopy() map[string]LabelSet {
	l.mu.RLock()
	defer l.mu.RUnlock()

	copiedLabels := make(map[string]LabelSet, len(l.currentLabels))
	for node, ls := range l.currentLabels {
		copiedLabels[node] = cloneLabelSet(ls)
	}
	return copiedLabels
}

// cloneLabelSet creates a deep copy of a LabelSet.
// Uses maps.Clone (Go 1.21+). For older Go, implement manually.
func cloneLabelSet(ls LabelSet) LabelSet {
	if ls == nil {
		return make(LabelSet) // Return empty, non-nil LabelSet
	}
	// For Go 1.21+
	return maps.Clone(ls)
	/*
	   // Manual clone for older Go versions (e.g., < Go 1.21):
	   cloned := make(LabelSet, len(ls))
	   for k, v := range ls {
	   	cloned[k] = v
	   }
	   return cloned
	*/
}

// areLabelSetsEqual checks if two LabelSets are identical.
// Uses maps.Equal (Go 1.21+ for generic version). struct{} is comparable.
func areLabelSetsEqual(s1, s2 LabelSet) bool {
	// For Go 1.21+ (specifically for generic maps.Equal)
	return maps.Equal(s1, s2)
	/*
	   // Manual check for older Go versions:
	   if len(s1) != len(s2) {
	   	return false
	   }
	   // Handle cases where one or both might be nil but effectively empty
	   if (s1 == nil && len(s2) == 0) || (s2 == nil && len(s1) == 0) {
	       return true
	   }
	   if s1 == nil || s2 == nil { // one is nil, the other is not and non-empty
	       return false
	   }
	   for k := range s1 {
	   	if _, ok := s2[k]; !ok {
	   		return false
	   	}
	   }
	   return true
	*/
}

// Helper for printing labels in a consistent order (useful for debugging/testing).
func PrintLabels(labels map[string]LabelSet) {
	tableNames := make([]string, 0, len(labels))
	for name := range labels {
		tableNames = append(tableNames, name)
	}
	slices.Sort(tableNames) // Requires Go 1.21+

	for _, tableName := range tableNames {
		labelList := make([]string, 0, len(labels[tableName]))
		for label := range labels[tableName] {
			labelList = append(labelList, label)
		}
		slices.Sort(labelList) // Requires Go 1.21+
		fmt.Printf("Table %s: Labels %v\n", tableName, labelList)
	}
}

// Example Usage (typically in a main function or a test file)
func main() {
	relations := []Relation{
		{FkTable: "orders", ReferencedTable: "customers"},
		{FkTable: "order_items", ReferencedTable: "orders"},
		{FkTable: "order_items", ReferencedTable: "products"},
		{FkTable: "reviews", ReferencedTable: "products"},
		{FkTable: "reviews", ReferencedTable: "customers"},
		// A more complex scenario:
		{FkTable: "employee_projects", ReferencedTable: "employees"},
		{FkTable: "employee_projects", ReferencedTable: "projects"},
		{FkTable: "project_tasks", ReferencedTable: "projects"},
		{FkTable: "employees", ReferencedTable: "departments"}, // Employees in depts
	}

	lpaAlg, err := NewLPA(relations)
	if err != nil {
		fmt.Printf("Error creating LPA: %v\n", err)
		return
	}

	fmt.Println("Initial Labels:")
	PrintLabels(lpaAlg.GetLabelsCopy())
	fmt.Println("---")

	// Run LPA with a maximum of 10 iterations
	finalLabels, iterations := lpaAlg.Run(10)

	fmt.Printf("\n--- Running LPA ---\n")
	fmt.Printf("LPA finished after %d iterations.\n", iterations)
	fmt.Println("\nFinal Labels:")
	PrintLabels(finalLabels)

	// Expected propagation:
	// - "departments" label should reach "employees", then "employee_projects".
	// - "customers" label should reach "orders", then "order_items". It also reaches "reviews".
	// - "products" label should reach "order_items" and "reviews".
	// - "projects" label should reach "employee_projects" and "project_tasks".
}

// Expected Output from example above:
// Initial Labels:
// Table customers: Labels [customers]
// Table departments: Labels [departments]
// Table employee_projects: Labels [employee_projects]
// Table employees: Labels [employees]
// Table order_items: Labels [order_items]
// Table orders: Labels [orders]
// Table products: Labels [products]
// Table project_tasks: Labels [project_tasks]
// Table projects: Labels [projects]
// Table reviews: Labels [reviews]
// ---
//
// --- Running LPA ---
// LPA finished after 3 iterations.
//
// Final Labels:
// Table customers: Labels [customers]
// Table departments: Labels [departments]
// Table employee_projects: Labels [departments, employee_projects, employees, projects]
// Table employees: Labels [departments, employees]
// Table order_items: Labels [customers, order_items, orders, products]
// Table orders: Labels [customers, orders]
// Table products: Labels [products]
// Table project_tasks: Labels [project_tasks, projects]
// Table projects: Labels [projects]
// Table reviews: Labels [customers, products, reviews]
