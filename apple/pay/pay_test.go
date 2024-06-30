package pay

import "testing"

func TestGetRecentOrder(t *testing.T) {
	pay := New("https://sandbox.itunes.apple.com/verifyReceipt", "")
	iap, err := pay.GetRecentOrder("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(iap)
}
