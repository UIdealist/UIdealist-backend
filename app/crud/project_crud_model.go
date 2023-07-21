package crud

// Project creation struct
type ProjectCreate struct {
	Name        string `json:"name" validate:"required;lte=45"`
	Description string `json:"description" validate:"required;lte=255"`
	Icon        string `json:"icon" validate:"required;lte=45"`
}

// Project update struct
type ProjectUpdate struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required;lte=45"`
	Description string `json:"description" validate:"required;lte=255"`
	Icon        string `json:"icon" validate:"required;lte=45"`
}

// Project delete struct
type ProjectDelete struct {
	ID string `json:"id"`
}
