package commands

// AddBookCommand - модель добавления книги в каталог
type AddBookCommand struct {
	Name   string  `json:"name"   validate:"required"`
	Author string  `json:"author" validate:"required"`
	Price  float64 `json:"price"  validate:"required"`
}
