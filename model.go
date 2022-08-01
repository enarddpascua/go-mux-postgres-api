package main

import (
	"database/sql"
	"errors"
)

type todo struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func (p *todo) getTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *todo) updateTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *todo) deleteTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *todo) createTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getTodoz(db *sql.DB, start, count int) ([]todo, error) {
	return nil, errors.New("Not implemented")
}
