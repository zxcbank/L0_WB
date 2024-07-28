package repositories

import (
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-template-microservice-v2/internal/data/contracts"
	"go-template-microservice-v2/internal/data/entities"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
)

type PgBookRepository struct {
	PgGorm *gormpg.PgGorm
}

func NewPgBookRepository(pgGorm *gormpg.PgGorm) contracts.IBookRepository {
	return &PgBookRepository{PgGorm: pgGorm}
}

func (p PgBookRepository) AddBook(bookEntity entities.BookEntity) error {
	err := p.PgGorm.DB.Create(bookEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting book into the database.")
	}

	return nil
}

func (p PgBookRepository) GetBook(id uuid.UUID) (*entities.BookEntity, error) {
	var book entities.BookEntity

	if err := p.PgGorm.DB.First(&book, id).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the book with id %s into the database.", id))
	}

	return &book, nil
}

func (p PgBookRepository) GetAllBook() ([]*entities.BookEntity, error) {
	var books []*entities.BookEntity

	if err := p.PgGorm.DB.Find(&books).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't find the books into the database."))
	}

	return books, nil
}

func (p PgBookRepository) UpdateBook(bookEntity entities.BookEntity) error {
	err := p.PgGorm.DB.Save(bookEntity).Error
	if err != nil {
		return errors.Wrap(err, "error in the inserting book into the database.")
	}

	return nil
}
