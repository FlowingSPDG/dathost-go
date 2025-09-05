package dathost

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetCurrentAccount implements DatHostClientv01.
func (dc *dathostClientv01) GetCurrentAccount(ctx context.Context) (*Account, error) {
	ep := "https://dathost.net/api/0.1/account"

	req, err := http.NewRequestWithContext(ctx, "GET", ep, nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var account Account
	if err := json.NewDecoder(res.Body).Decode(&account); err != nil {
		return nil, err
	}
	return &account, nil
}

// ListInvoices implements DatHostClientv01.
func (dc *dathostClientv01) ListInvoices(ctx context.Context) ([]Invoice, error) {
	ep := "https://dathost.net/api/0.1/account/invoices"

	req, err := http.NewRequestWithContext(ctx, "GET", ep, nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var invoices []Invoice
	if err := json.NewDecoder(res.Body).Decode(&invoices); err != nil {
		return nil, err
	}
	return invoices, nil
}

// GetInvoiceAsHTML implements DatHostClientv01.
func (dc *dathostClientv01) GetInvoiceAsHTML(ctx context.Context, id string) (string, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/account/invoices/%s", id)

	req, err := http.NewRequestWithContext(ctx, "GET", ep, nil)
	if err != nil {
		return "", err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(html), nil
}
