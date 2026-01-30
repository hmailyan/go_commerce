package brands

import "context"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) error {

}
