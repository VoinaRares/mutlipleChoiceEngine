package traversel

import (
	"awesomeProject/pathing"
	"errors"
	"slices"
)

// The main story of the game. This tree will be used to traverse the story
var optionTree pathing.OptionTree
var current *pathing.Node

type Option struct {
	Id   string
	Text string
}

// Initialize this initializes the path and the options inside the tree
func Initialize(path string) ([]Option, error) {
	var err error
	optionTree, err = pathing.BuildOptionTree(path)
	if err != nil {
		return nil, err
	}

	current = optionTree.Options[optionTree.HeadId]
	options := getOptions(optionTree.HeadId)
	return options, nil
}

// ChooseOption uses the id of an option Node to return the next possible strings
func ChooseOption(optionId string) ([]Option, error) {
	if !slices.Contains(current.ChildrenIds, optionId) {
		return nil, errors.New("option does not exist")
	}

	/*
		I might think about moving this into a function because it is basically the traverse operation.
		For now, it would be a line, so it doesn't really make sense...
		In the future it could hold extra logic
	*/
	current = optionTree.Options[optionId]

	return getOptions(current.OptionId), nil
}

func getOptions(nodeId string) []Option {
	headNode := optionTree.Options[nodeId]
	var options []Option

	for _, optionId := range headNode.ChildrenIds {
		newOption := Option{
			Id:   optionId,
			Text: optionTree.Options[optionId].OptionText,
		}
		options = append(options, newOption)
	}

	return options
}
