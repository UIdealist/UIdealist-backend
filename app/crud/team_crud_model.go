package crud

// Team creation struct
type TeamCreate struct {
	Name string `json:"name" validate:"required;lte=45"`
}

// Team update struct
type TeamUpdate struct {
	Name string `json:"name" validate:"required;lte=45"`
}

// Team member add struct
type TeamMemberAdd struct {
	UserID string `json:"userId" validate:"required;lte=45"`
}

// Team member remove struct
type TeamMemberRemove struct {
	UserID string `json:"userId" validate:"required;lte=45"`
}
