// UNIT TEST

package addresses

import "testing"

func TestAddressType(t *testing.T) {
	addressForTest := "Avenue"

	kindOfAddressExpected := "Avenue"

	kindOfAddressReceived := AddressType(addressForTest)
	if kindOfAddressReceived != kindOfAddressExpected {
		t.Errorf("Received kind if different than expected! Expected %s and got %s", kindOfAddressExpected, kindOfAddressReceived)
	}
}