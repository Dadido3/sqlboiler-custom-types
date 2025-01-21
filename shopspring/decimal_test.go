package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
)

func TestDecimal_Value(t *testing.T) {
	t.Parallel()

	tests := []string{
		"3.14",
		"0",
		"43.4292",
	}

	for i, test := range tests {
		d := NewDecimal(decimal.RequireFromString(test))

		val, err := d.Value()
		if err != nil {
			t.Errorf("%d) %+v", i, err)
		}

		s, ok := val.(string)
		if !ok {
			t.Errorf("%d) wrong type returned", i)
		}

		if s != test {
			t.Errorf("%d) want: %s, got: %s", i, test, s)
		}
	}

	zero := Decimal{}
	if _, err := zero.Value(); err != nil {
		t.Error("zero value should not error")
	}
}

func TestDecimal_Scan(t *testing.T) {
	t.Parallel()

	tests := []string{
		"3.14",
		"0",
		"43.4292",
	}

	for i, test := range tests {
		var d Decimal
		if err := d.Scan(test); err != nil {
			t.Error(err)
		}

		if got := d.String(); got != test {
			t.Errorf("%d) want: %s, got: %s", i, test, got)
		}
	}

	var d Decimal
	if err := d.Scan(nil); err == nil {
		t.Error("it should disallow scanning from a null value")
	}
}

func TestNullDecimal_Value(t *testing.T) {
	t.Parallel()

	tests := []string{
		"3.14",
		"0",
		"43.4292",
	}

	for i, test := range tests {
		d := NewDecimal(decimal.RequireFromString(test))

		val, err := d.Value()
		if err != nil {
			t.Errorf("%d) %+v", i, err)
		}

		s, ok := val.(string)
		if !ok {
			t.Errorf("%d) wrong type returned", i)
		}

		if s != test {
			t.Errorf("%d) want: %s, got: %s", i, test, s)
		}
	}

	zero := NullDecimal{}
	if _, err := zero.Value(); err != nil {
		t.Error("zero value should not error")
	}
}

func TestNullDecimal_Scan(t *testing.T) {
	t.Parallel()

	tests := []string{
		"3.14",
		"0",
		"43.4292",
	}

	for i, test := range tests {
		var d NullDecimal
		if err := d.Scan(test); err != nil {
			t.Error(err)
		}

		if got := d.String(); got != test {
			t.Errorf("%d) want: %s, got: %s", i, test, got)
		}
	}

	var d NullDecimal
	if err := d.Scan(nil); err != nil {
		t.Error(err)
	}
	if d.Decimal != nil {
		t.Error("it should have been nil")
	}
}

func TestDecimal_JSON(t *testing.T) {
	t.Parallel()

	s := struct {
		D Decimal `json:"d"`
	}{}

	err := json.Unmarshal([]byte(`{"d":"54.45"}`), &s)
	if err != nil {
		t.Error(err)
	}

	want := decimal.RequireFromString("54.45")
	if s.D.Cmp(want) != 0 {
		t.Error("D was wrong:", s.D)
	}
}

func TestDecimal_Text(t *testing.T) {
	t.Parallel()

	d := new(Decimal)

	err := d.UnmarshalText([]byte(`54.45`))
	if err != nil {
		t.Error(err)
	}

	want := decimal.RequireFromString("54.45")
	if d.Cmp(want) != 0 {
		t.Error("D was wrong:", d)
	}
}

func TestNullDecimal_JSON(t *testing.T) {
	t.Parallel()

	s := struct {
		N NullDecimal `json:"n"`
	}{}

	err := json.Unmarshal([]byte(`{"n":"54.45"}`), &s)
	if err != nil {
		t.Error(err)
	}

	want := decimal.RequireFromString("54.45")
	if s.N.Cmp(want) != 0 {
		fmt.Println(want, s.N)
		t.Error("N was wrong:", s.N)
	}
}

func TestNullDecimal_Text(t *testing.T) {
	t.Parallel()

	n := new(NullDecimal)

	err := n.UnmarshalText([]byte(`54.45`))
	if err != nil {
		t.Error(err)
	}

	want := decimal.RequireFromString("54.45")
	if n.Cmp(want) != 0 {
		fmt.Println(want, n)
		t.Error("N was wrong:", n)
	}
}

func TestNullDecimal_JSONNil(t *testing.T) {
	t.Parallel()

	var n NullDecimal
	b, _ := json.Marshal(n)
	if string(b) != `null` {
		t.Errorf("want: null, got: %s", b)
	}

	n2 := new(NullDecimal)
	b, _ = json.Marshal(n2)
	if string(b) != `null` {
		t.Errorf("want: null, got: %s", b)
	}
}

func TestNullDecimal_TextNil(t *testing.T) {
	t.Parallel()

	var n NullDecimal
	b, _ := n.MarshalText()
	if string(b) != `null` {
		t.Errorf("want: null, got: %s", b)
	}

	n2 := new(NullDecimal)
	b, _ = n2.MarshalText()
	if string(b) != `null` {
		t.Errorf("want: null, got: %s", b)
	}
}

func TestNullDecimal_IsZero(t *testing.T) {
	t.Parallel()

	var nullable qmhelper.Nullable = NullDecimal{}

	if !nullable.IsZero() {
		t.Error("it should be zero")
	}

	nullable = NullDecimal{Decimal: &decimal.Decimal{}}
	if nullable.IsZero() {
		t.Error("it should not be zero")
	}
}
