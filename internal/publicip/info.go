package publicip

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/publicip/models"
)

var (
	ErrTooManyRequests = errors.New("too many requests sent for this month")
	ErrBadHTTPStatus   = errors.New("bad HTTP status received")
)

func Info(ctx context.Context, client *http.Client, ip net.IP) ( //nolint:interfacer
	result models.IPInfoData, err error) {
	const baseURL = "https://api.myip.com/"
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return result, err
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusTooManyRequests:
		return result, fmt.Errorf("%w: %s", ErrTooManyRequests, baseURL)
	default:
		return result, fmt.Errorf("%w: %d %s", ErrBadHTTPStatus,
			response.StatusCode, response.Status)
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	countryCode := strings.ToLower(result.CountryCode)
	country, ok := constants.CountryCodes()[countryCode]
	if ok {
		result.Country = country
	}
	return result, nil
}
