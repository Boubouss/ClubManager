package models

import (
	"github.com/google/uuid"
)

type User struct {
  Id uuid.UUID
  Username string
  Email string 
  Phonenumber string
}

type UserForm struct {
  Username string
  Email string
  Phonenumber string
  Password string
}

