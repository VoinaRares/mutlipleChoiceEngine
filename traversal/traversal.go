package traversal

import (
	"errors"
	"learningProject/pathing"
	"learningProject/playerstates"
	"maps"
	"slices"
)

// The main story of the game. This tree will be used to traverse the story
var optionTree pathing.OptionTree
var current *pathing.Node
var playerState playerstates.PlayerStates

type Option struct {
	Id           string
	Text         string
	PreviewText  string
	Requirements map[string]int
}

// Initialize this initializes the path and the options inside the tree
func Initialize(treePath string, playerPath string) error {
	var err error
	optionTree, err = pathing.BuildOptionTree(treePath)
	if err != nil {
		return err
	}

	playerState, err = playerstates.BuildPlayerStates(playerPath)
	if err != nil {
		return err
	}

	err = validateTreePlayerStates()
	if err != nil {
		return err
	}
	current = optionTree.Options[optionTree.HeadId]
	return nil
}

func GetCurrent() Option {
	return Option{
		Id:           current.OptionId,
		Text:         current.OptionText,
		PreviewText:  current.PreviewText,
		Requirements: current.Requirements,
	}
}

// ChooseOption uses the id of an option Node to return the next possible strings
func ChooseOption(optionId string) error {
	if !slices.Contains(current.ChildrenIds, optionId) {
		return errors.New("option does not exist")
	}
	return traverse(optionId)
}

func traverse(optionId string) error {
	if requirementsMet(optionId) {
		current = optionTree.Options[optionId]
	}
	return nil
}

func requirementsMet(optionId string) bool {
	for key, value := range optionTree.Options[optionId].Requirements {
		if playerState.States[key] < value {
			return false
		}
	}
	return true
}

func GetOptions(optionId string) []Option {
	head := optionTree.Options[optionId]
	options := make([]Option, 0, len(head.ChildrenIds))

	for _, optId := range head.ChildrenIds {
		if requirementsMet(optId) {
			newOption := Option{
				Id:           optId,
				Text:         optionTree.Options[optId].OptionText,
				PreviewText:  optionTree.Options[optId].PreviewText,
				Requirements: optionTree.Options[optId].Requirements,
			}
			options = append(options, newOption)
		}
	}

	return options
}

// validate that the tree does not carry different states
func validateTreePlayerStates() error {
	playerStates := slices.Collect(maps.Keys(playerState.States))
	treeStates := optionTree.States

	/*
		This implementation allows us to have additional states in the player JSON.
		This allows us to have more states inside the player, which we will not use.
		The tree and the requirements are still correct because we checked their matches
		when building the tree
	*/
	for _, state := range treeStates {
		if !slices.Contains(playerStates, state) {
			return errors.New("player state does not exist inside tree")
		}
	}
	return nil
}
