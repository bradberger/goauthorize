package transaction

import (
    "testing"
)

// TestNewCreditCardStripsWhitespace should remove all whitespace from card.
func TestNewCreditCardStripsWhitespace(t *testing.T) {
    cardNum := "4222 2222 2222 2222"
    card := NewCreditCard(cardNum, "", "")
    if card.CardNumber != "4222222222222222" {
        t.Fail()
    }
}

// TestNewCreditCardStripsWhitespace should remove all non-numeric characters from card number.
func TestNewCreditCardStripsDashes(t *testing.T) {
    cardNum := "4222-2222-2222-2222"
    card := NewCreditCard(cardNum, "", "")
    if card.CardNumber != "4222222222222222" {
        t.Fail()
    }
}

func TestNewCreditCardStripsInvalid(t *testing.T) {
    cardNum := "acbdefgABCDEFG!@#$%^&*()_+"
    card := NewCreditCard(cardNum, "", "")
    if card.CardNumber != "" {
        t.Fail()
    }
}
