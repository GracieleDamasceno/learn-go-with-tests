package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance().String()
		want := Bitcoin(10).String()

		assertStrings(got, want, t)
	})

	t.Run("withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance().String()
		want := Bitcoin(10).String()

		assertStrings(got, want, t)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		error := wallet.Withdraw(Bitcoin(100))

		assertError(error, "insuficient funds", t)
	})

}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(got error, want string, t testing.TB) {
	if got == nil {
		t.Fatal("Expected an error, got none")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
