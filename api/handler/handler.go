package handler

import (
	"database/sql"
)

type HandlerV1 struct {
	db *sql.DB 
}

func NewHandlerV1(db *sql.DB) *HandlerV1 {
	return &HandlerV1 {
		db: db,
	}
}
