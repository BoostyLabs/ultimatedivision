package userauth

import (
	"context"
	"crypto/subtle"
	"github.com/google/uuid"
	"time"

	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"

	"ultimatedivision/internal/auth"
	"ultimatedivision/users"
)

const (
	// TokenExpirationTime after passing this time token expires.
	TokenExpirationTime = 24 * time.Hour
)

var (
	// ErrUnauthenticated should be returned when user performs unauthenticated action.
	ErrUnauthenticated = errs.Class("user unauthenticated error")

	// Error is a error class for internal auth errors.
	Error = errs.Class("user auth internal error")
)

// Service is handling all user authentication logic.
//
// architecture: Service
type Service struct {
	users  users.DB
	signer auth.TokenSigner
}

// NewService is a constructor for user auth service.
func NewService(users users.DB, signer auth.TokenSigner) *Service {
	return &Service{
		users:  users,
		signer: signer,
	}
}

// Token authenticates User by credentials and returns auth token.
func (service *Service) Token(ctx context.Context, email string, password string) (token string, err error) {
	user, err := service.users.GetByEmail(ctx, email)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return "", ErrUnauthenticated.Wrap(err)
	}

	claims := auth.Claims{
		Email:     user.Email,
		ExpiresAt: time.Now().Add(TokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return token, nil
}

// Authorize validates token from context and returns authorized Authorization.
func (service *Service) Authorize(ctx context.Context) (_ auth.Claims, err error) {
	tokenS, ok := auth.GetToken(ctx)
	if !ok {
		return auth.Claims{}, ErrUnauthenticated.Wrap(err)
	}

	token, err := auth.FromBase64URLString(string(tokenS))
	if err != nil {
		return auth.Claims{}, Error.Wrap(err)
	}

	claims, err := service.authenticate(token)
	if err != nil {
		return auth.Claims{}, ErrUnauthenticated.Wrap(err)
	}

	err = service.authorize(ctx, claims)
	if err != nil {
		return auth.Claims{}, ErrUnauthenticated.Wrap(err)
	}

	return *claims, nil
}

// authenticate validates token signature and returns authenticated *satelliteauth.Authorization.
func (service *Service) authenticate(token auth.Token) (_ *auth.Claims, err error) {
	signature := token.Signature

	err = service.signer.SignToken(&token)
	if err != nil {
		return nil, err
	}

	if subtle.ConstantTimeCompare(signature, token.Signature) != 1 {
		return nil, errs.New("incorrect signature")
	}

	claims, err := auth.FromJSON(token.Payload)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// authorize checks claims and returns authorized User.
func (service *Service) authorize(ctx context.Context, claims *auth.Claims) (err error) {
	if !claims.ExpiresAt.IsZero() && claims.ExpiresAt.Before(time.Now()) {
		return ErrUnauthenticated.Wrap(err)
	}

	_, err = service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return errs.New("authorization failed. no user with email: %s", claims.Email)
	}

	return nil
}

// GetUserByEmail returns user by email from DB.
func (service *Service) GetUserByEmail(ctx context.Context, email string) (users.User, error) {
	return service.users.GetByEmail(ctx, email)
}

// RegisterUser - register a new user.
func (service *Service) RegisterUser(ctx context.Context, email, password, nickName, firstName, lastName string) error {
	user := users.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: []byte(password),
		NickName:     nickName,
		FirstName:    firstName,
		LastName:     lastName,
		LastLogin:    time.Time{},
		Status:       users.StatusActive,
		CreatedAt:    time.Now(),
	}

	err := user.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}

	return service.users.Create(ctx, user)
}
