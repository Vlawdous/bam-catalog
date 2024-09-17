package repository

import "bam-catalog/internal/entity"

type BookRepository struct {
	*Repository
}

func (br *BookRepository) findById(id string) *entity.Book {
	searchableBook := &entity.Book{}

	br.First(searchableBook, id)
	return searchableBook
}

func NewBookRepository(repository *Repository) *BookRepository {
	return &BookRepository{Repository: repository}
}
