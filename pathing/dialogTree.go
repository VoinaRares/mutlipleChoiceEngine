package pathing

type OptionTree struct {
	HeadId  string
	Options map[string]*Node
}

type Node struct {
	// The id might not be useful anymore... I will keep it for now. Refactoring should be easy.
	OptionId     string         `json:"option_id"`
	OptionText   string         `json:"option_text"`
	ChildrenIds  []string       `json:"children_ids"`
	PreviewText  string         `json:"preview_text"`
	Requirements map[string]int `json:"requirements"`
	StateChanges map[string]int `json:"state_changes"`
}
