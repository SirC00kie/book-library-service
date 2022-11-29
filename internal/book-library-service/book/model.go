package book

type Book struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Type        string `json:"type" bson:"type"`
	Content     string `json:"content" bson:"content"`
	Author      string `json:"author" bson:"author"`
	Year        int    `json:"year" bson:"year"`
	Description string `json:"description" bson:"description"`
}

type CreateUserDTO struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
