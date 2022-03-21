package model

import (
	"time"
)

type User struct {
	ID         int32     `json:"ID" sql:"ID"`
	Name       string    `json:"name" sql:"name"`
	Email      string    `json:"email" sql:"email"`
	UserName   string    `json:"uname" sql:"uname"`
	Password   string    `json:"password" sql:"password"`
	TokenHash  string    `json:"tokenhash" sql:"tokenhash"`
	IsVerified bool      `json:"isverfiied" sql:"isverified"`
	CreatedAt  time.Time `json:"createdat" sql:"createdat"`
}

type VerificationDataType int

const (
	MailConfirmation VerificationDataType = iota + 1
	PassReset
)

type VerificationData struct {
	Email     string               `json:"email" validate:"required" sql:"email"`
	Code      string               `json:"code" validate:"required" sql:"code"`
	ExpiresAt time.Time            `json:"expiresat" sql:"expiresat"`
	Type      VerificationDataType `json:"type" sql:"type"`
}
