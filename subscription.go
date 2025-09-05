package dathost

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// UpdateSubscription implements DatHostClientv01.
func (dc *dathostClientv01) UpdateSubscription(ctx context.Context, id string, req UpdateSubscriptionRequest) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/subscription", id)

	data := url.Values{}
	data.Set("action", req.Action)
	if req.Months > 0 {
		data.Set("months", fmt.Sprintf("%d", req.Months))
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", ep, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	dc.addHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
