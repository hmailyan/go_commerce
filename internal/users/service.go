package users

import (
	"context"
)

type PasswordHasher interface {
	GenerateUserTokens(password string) (string, error)
}

type TokenGenerator interface {
	Generate(userID string) (string, error)
	ValidateToken(signedToken string) (string, error)
}

type Service struct {
	repo          Repository
	hasher        PasswordHasher
	generateToken TokenGenerator
}

func NewService(r Repository, h PasswordHasher, g TokenGenerator) *Service {
	return &Service{
		repo:          r,
		hasher:        h,
		generateToken: g,
	}
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest) (*SignUpOutput, error) {
	exists, err := s.repo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	hash, err := s.hasher.GenerateUserTokens(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:     req.Email,
		Password:  hash,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken.Generate(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &SignUpOutput{
		ID:    user.ID.String(),
		Token: token,
	}, nil
}
func (s *Service) Me(ctx context.Context, token string) (user *User, error error) {
	// For simplicity, let's assume we have a method to get user by ID
	uid, err := s.generateToken.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	user, err = s.repo.FindByID(ctx, uid)
	if err != nil {
		return nil, ErrCantFindUser
	}
	return user, nil
}
