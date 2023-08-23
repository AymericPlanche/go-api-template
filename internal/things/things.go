package things

// Thing is just an example of entity
// Note that for simplicity we use here the same struct for app and persistance. In some cases it's preferred to use 2 different types
type Thing struct {
	ID   ID     `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// ID represents the ID of at thing
type ID string
