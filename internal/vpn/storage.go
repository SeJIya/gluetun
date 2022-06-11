package vpn

import (
	"github.com/qdm12/gluetun/internal/models"
)

func (l *Loop) GetServerList() (servers []models.Server, err error) {
	vpnSettings := l.state.GetSettings()

	provider := vpnSettings.Provider
	provider.ServerSelection.Countries = []string{}
	provider.ServerSelection.Regions = []string{}
	provider.ServerSelection.Cities = []string{}
	provider.ServerSelection.Hostnames = []string{}

	return l.storage.FilterServers(*provider.Name, provider.ServerSelection)
}
