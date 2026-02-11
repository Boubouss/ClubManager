package middlewares

import (
	"ClubManager/users/internal/service"
	"ClubManager/shared/models"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type loggingService struct {
  next service.UserService
} 

func NewLoggingService(next service.UserService) service.UserService {
  return &loggingService{
    next: next,
  }
}

func (s *loggingService) CreateUser(ctx context.Context, data *models.UserForm) (userId uuid.UUID, err error) {
  defer func(begin time.Time){
    logrus.WithFields(logrus.Fields{
      "type": "CreateUser",
      "took": time.Since(begin),
      "userId": userId,
      "err": err,
    })
  }(time.Now())

  return s.next.CreateUser(ctx, data)
}
