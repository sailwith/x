package pay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecentOrder(t *testing.T) {
	pay := New("https://sandbox.itunes.apple.com/verifyReceipt", "")
	iap, err := pay.GetRecentOrder("")
	assert.NoError(t, err)
	t.Log(iap)
}
