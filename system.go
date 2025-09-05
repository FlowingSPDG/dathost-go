package dathost

import (
	"encoding/json"
	"net/http"
)

// GetCustomDomains implements DatHostClientv01.
func (dc *dathostClientv01) GetCustomDomains() ([]CustomDomain, error) {
	ep := "https://dathost.net/api/0.1/custom-domains"

	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var domains []CustomDomain
	if err := json.NewDecoder(res.Body).Decode(&domains); err != nil {
		return nil, err
	}
	return domains, nil
}
