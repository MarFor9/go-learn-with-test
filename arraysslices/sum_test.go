package arraysslices

import (
	"reflect"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		checkSum := func(t testing.TB, got, expected int) {
			if got != expected {
				t.Errorf("expected: '%d' but got: '%d'", expected, got)
			}
		}
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		checkSum(t, got, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		expected := []int{6, 15}
		checkSums(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		expected := []int{5, 9}
		checkSums(t, got, expected)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		checkSums(t, got, expected)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all emelents", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concat, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)
	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})
	t.Run("filter accounts base on balance", func(t *testing.T) {
		accounts := []Account{
			{Name: "a", Balance: 10},
			{Name: "b", Balance: 12},
			{Name: "c", Balance: 13},
			{Name: "d", Balance: 14},
		}
		got := Filter(accounts, func(a Account) bool {
			return a.Balance > 12
		})
		want := []Account{{Name: "c", Balance: 13}, {Name: "d", Balance: 14}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %+v but got %+v", want, got)
		}
	})
	t.Run("All balance are above some value", func(t *testing.T) {
		accounts := []Account{
			{Name: "b", Balance: 15},
			{Name: "c", Balance: 13},
			{Name: "d", Balance: 14},
		}
		got := All(accounts, func(a Account) bool {
			return a.Balance > 12
		})
		AssertTrue(t, got)
	})
	t.Run("map collection of A to collection B", func(t *testing.T) {
		data := []string{"a", "b", "c", "c"}

		got := Map(data, func(st string) string {
			return strings.ToUpper(st)
		})
		want := []string{"A", "B", "C", "C"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v but want %+v", got, want)
		}
	})
}
func checkSums(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v got %v", expected, got)
	}
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
