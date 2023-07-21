package crud

// Idea creation struct
type IdeaCreate struct {
	Title        string  `json:"title" validate:"required;lte=45"`
	Body         string  `json:"body" validate:"required;lte=255"`
	ProjectID    *string `json:"proId"`
	BrainStormID *string `json:"brainstormId"`
}

// Idea update struct
type IdeaUpdate struct {
	Title string `json:"title" validate:"required;lte=45"`
	Body  string `json:"body" validate:"required;lte=255"`
}

// Brainstorm creation struct
type BrainstormCreateIdea struct {
	Title string `json:"title" validate:"required;lte=45"`
	Body  string `json:"body" validate:"required;lte=255"`
}

type BrainstormCreate struct {
	Name      string                 `json:"name" validate:"required;lte=45"`
	Ideas     []BrainstormCreateIdea `json:"ideas"`
	ProjectID *string                `json:"proId"`
}

// Brainstorm update struct
type BrainstormUpdate struct {
	Name string `json:"name" validate:"required;lte=45"`
}
