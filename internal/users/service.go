package users

import (
	"context"
	"fmt"
	"log"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, givenPassword string) error
}

type TokenGenerator interface {
	GenerateUserTokens(userID string) (string, error)
	ValidateToken(signedToken string) (string, error)
	GenerateRandomToken() (string, error)
}

type Mailer interface {
	SendVerificationEmail(toEmail, verifyToken string) error
}

type Service struct {
	repo          Repository
	hasher        PasswordHasher
	generateToken TokenGenerator
	mailer        Mailer
}

func NewService(r Repository, h PasswordHasher, g TokenGenerator, m Mailer) *Service {
	return &Service{
		repo:          r,
		hasher:        h,
		generateToken: g,
		mailer:        m,
	}
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest) error {
	exists, err := s.repo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if exists {
		return ErrEmailAlreadyExists
	}

	hash, err := s.hasher.HashPassword(req.Password)
	if err != nil {
		return err
	}

	VerificationToken, err := s.generateToken.GenerateRandomToken()
	if err != nil {
		return err
	}

	user := &User{
		Email:             req.Email,
		Password:          hash,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		VerificationToken: VerificationToken,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	go func() {
		err := s.mailer.SendVerificationEmail(
			user.Email,
			VerificationToken,
		)
		if err != nil {
			// log error, DO NOT panic
			log.Printf("failed to send verification email: %v", err)
		}
	}()

	return nil
}

func (s *Service) VerifyEmail(ctx context.Context, token string) error {
	err := s.repo.VerifyEmail(ctx, token)
	if err != nil {
		return ErrInvalidVerificationToken
	}

	return nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (userRes *UserResponse, token string, error error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		fmt.Printf("1")
		return nil, "", err
	}
	err = s.hasher.VerifyPassword(user.Password, req.Password)
	if err != nil {
		fmt.Printf("2")
		return nil, "", err
	}

	token, err = s.generateToken.GenerateUserTokens(user.ID.String())

	if err != nil {
		fmt.Printf("3")

		return nil, "", err
	}

	return ToUserResponse(user), token, nil

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
