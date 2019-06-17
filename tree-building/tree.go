package tree

import (
	"errors"
	"sort"
)

// Record struct
type Record struct {
	ID     int
	Parent int
}

// Node struct
type Node struct {
	ID       int
	Children []*Node
}

// Build reconstructs a tree structure from records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Ensure records are sorted by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].Parent != 0 {
		return nil, errors.New("root node has parent")
	}

	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}

	if len(records) == 1 {
		return &Node{}, nil
	}

	// Build nodes
	nodes := make(map[int]*Node, len(records))
	for _, r := range records {
		nodes[r.ID] = &Node{ID: r.ID}
	}

	// Link nodes
	for i, r := range records[1:] {
		if r.ID == 0 {
			return nil, errors.New("duplicate root")
		}

		if r.ID-records[i].ID > 1 {
			return nil, errors.New("non-continuous")
		}

		if records[i].ID == r.ID {
			return nil, errors.New("duplicate node")
		}

		if r.ID == r.Parent {
			return nil, errors.New("cycle directly")
		}

		if r.ID < r.Parent {
			return nil, errors.New("cycle indirectly")
		}

		n := nodes[r.Parent]
		n.Children = append(n.Children, nodes[r.ID])
	}

	return nodes[0], nil
}
