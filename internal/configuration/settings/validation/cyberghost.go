package validation

import (
	"github.com/qdm12/gluetun/internal/models"
)

func CyberghostCountryChoices(servers []models.CyberghostServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Country
	}
	return makeUnique(choices)
}

func CyberghostHostnameChoices(servers []models.CyberghostServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}