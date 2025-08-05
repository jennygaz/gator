package main

import (
	"github.com/jennygaz/gator/internal/config"
	"github.com/jennygaz/gator/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}
