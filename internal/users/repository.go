package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&User{}).
		Where("email = ?", email).
		Count(&count).Error

	return count > 0, err

}

func (r *Repository) Create(ctx context.Context, req *User) error {
	return r.db.Create(req).Error
}

func (r *Repository) VerifyEmail(ctx context.Context, token string) error {
	err := r.db.Model(&User{}).
		Where("verification_token = ?", token).
		Updates(map[string]interface{}{"verification_at": time.Now(), "verification_token": ""}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Need to change
func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
