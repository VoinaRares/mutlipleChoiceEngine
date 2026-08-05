package pathing

import (
	"encoding/json"
	"errors"
	"os"
)

type nodes struct {
	HeadNodeId string `json:"head_node_id"`
	Nodes      []Node `json:"nodes"`
}

func BuildOptionTree(path string) (OptionTree, error) {
	tree := OptionTree{
		Options: make(map[string]*Node),
	}

	nodeArray, err := parseJsonToNodes(path)
	if err != nil {
		return OptionTree{}, err
	}
	tree.HeadId = nodeArray.HeadNodeId

	for _, node := range nodeArray.Nodes {
		// Declare nodeCopy so that we allocate new memory for the variables inside the tree
		nodeCopy := node
		tree.Options[node.OptionId] = &nodeCopy
	}

	err = validateOptionTree(&tree)
	if err != nil {
		return OptionTree{}, err
	}

	return tree, nil
}

func parseJsonToNodes(path string) (nodes, error) {

	jsonFile, err := os.Open(path)
	if err != nil {
		return nodes{}, err
	}
	defer jsonFile.Close()

	var nodeArray nodes
	return nodeArray, json.NewDecoder(jsonFile).Decode(&nodeArray)
}

// we use pointers for efficiency here
func validateOptionTree(tree *OptionTree) error {
	if tree.HeadId == "" {
		return errors.New("option tree: head_node_id is missing from story file")
	}
	if _, ok := tree.Options[tree.HeadId]; !ok {
		return errors.New("invalid OptionTree! HeadId was not part of the tree")
	}
	for _, node := range tree.Options {
		for _, childrenId := range node.ChildrenIds {
			if _, ok := tree.Options[childrenId]; !ok {
				return errors.New("invalid OptionTree! ChildrenId was not part of the tree")
			}

		}
	}
	return nil
}
