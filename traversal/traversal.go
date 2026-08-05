package traversal

import (
	"errors"
	"learningProject/pathing"
	"slices"
)

// The main story of the game. This tree will be used to traverse the story
var optionTree pathing.OptionTree
var current *pathing.Node

type Option struct {
	Id          string
	PreviewText string
}

// Initialize this initializes the path and the options inside the tree
func Initialize(path string) error {
	var err error
	optionTree, err = pathing.BuildOptionTree(path)
	if err != nil {
		return err
	}

	current = optionTree.Options[optionTree.HeadId]
	return nil
}

func GetCurrent() Option {
	return Option{
		Id:          current.OptionId,
		PreviewText: current.PreviewText,
	}
}

// ChooseOption uses the id of an option Node to return the next possible strings
func ChooseOption(optionId string) error {
	if !slices.Contains(current.ChildrenIds, optionId) {
		return errors.New("option does not exist")
	}

	/*
		I might think about moving this into a function because it is basically the traverse operation.
		For now, it would be a line, so it doesn't really make sense...
		In the future it could hold extra logic
	*/
	current = optionTree.Options[optionId]

	return nil
}

func GetOptions(optionId string) []Option {
	head := optionTree.Options[optionId]
	options := make([]Option, 0, len(head.ChildrenIds))

	for _, optId := range head.ChildrenIds {
		newOption := Option{
			Id:          optId,
			PreviewText: optionTree.Options[optId].PreviewText,
		}
		options = append(options, newOption)
	}

	return options
}
