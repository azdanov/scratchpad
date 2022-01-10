package main

import "github.com/azdanov/scratchpad/pkg/models"

type templateData struct {
	Scratch   *models.Scratch
	Scratches []*models.Scratch
}
