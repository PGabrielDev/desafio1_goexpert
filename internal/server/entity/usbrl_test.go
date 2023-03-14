package entity

import "testing"

func Test_create_usbrl_when_bid_is_blank(t *testing.T) {
	usbrl := USDBRL{}

	err := usbrl.Validate()

	if err == nil {
		t.Errorf("Error expected")
	}

}

func Test_create_usbrl_when_bid_is_valid_value(t *testing.T) {
	usbrl := USDBRL{Bid: "50.00"}

	err := usbrl.Validate()

	if err != nil {
		t.Errorf("Error expected")
	}

}
