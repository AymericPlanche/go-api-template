package things

// Thing is just an example of entity
type Thing struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

// ID represents the ID of at thing
type ID string
