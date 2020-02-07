package main

import "mattmeinzer.com/plants/pkg/models"

type templateData struct {
	Plant  *models.Plant
	Plants []*models.Plant
}
