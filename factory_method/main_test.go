package factorymethod

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(1000)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(1000)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodNonExsitent(t *testing.T) {
	payment, err := GetPaymentMethod(20)
	if err != nil && payment == nil {
		t.Error("A payment method with ID '20' must exist")
	}

	t.Log("LOG:", err)
}
