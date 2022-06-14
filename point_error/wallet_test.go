package point_error

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

func TestBitcoin_String(t *testing.T) {
	bitcoin := Bitcoin(10)
	if bitcoin.String() != "10 BTC" {
		t.Errorf("bitcoin format not expected")
	}

}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func assertError(t *testing.T, got error, want string) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got '%s',want '%s'", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
