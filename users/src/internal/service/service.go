package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"ClubManager/users/internal/types"
)

type UserService interface {
  CreateUser(context.Context, *types.UserForm) (uuid.UUID, error)
}

type userService struct {
  db *pgx.Conn
}

func NewUserService(db *pgx.Conn) *userService {
  return &userService{
    db: db,
  }
}

func (s *userService) CreateUser(ctx context.Context, data *types.UserForm) (uuid.UUID, error) {
  // Extraire la db depuis le context
  // Créer la requete sql à partir des infos dans le UserForm 
  // Jouer la requete et récuperer le uuid généré
  return uuid.UUID{}, nil
}
