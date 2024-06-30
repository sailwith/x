package pay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Pay struct {
	endpoint string
	bundleID string
}

func New(endpoint string, bundleID string) *Pay {
	return &Pay{
		endpoint: endpoint,
		bundleID: bundleID,
	}
}

// IAP
func (p *Pay) GetRecentOrder(receipt string) (*InApp, error) {
	data := make(map[string]interface{})
	data["receipt-data"] = receipt
	bytesData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(p.endpoint, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload ReceiptPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return nil, err
	}

	return payload.RecentOrder(p.bundleID)
}
