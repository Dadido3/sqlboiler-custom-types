package shopspring

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/shopspring/decimal"
)

var (
	nullBytes = []byte("null")
)

var (
	_ driver.Valuer = Decimal{}
	_ driver.Valuer = NullDecimal{}
	_ sql.Scanner   = &Decimal{}
	_ sql.Scanner   = &NullDecimal{}
)

// Decimal is a DECIMAL in sql. Its zero value is valid for use with both
// Value and Scan.
//
// Although decimal can represent NaN and Infinity it will return an error
// if an attempt to store these values in the database is made.
//
// Because it cannot be nil, when Big is nil Value() will return "0"
// It will error if an attempt to Scan() a "null" value into it.
type Decimal struct {
	decimal.Decimal
}

// NullDecimal is the same as Decimal, but allows the Big pointer to be nil.
// See documentation for Decimal for more details.
//
// When going into a database, if Big is nil it's value will be "null".
type NullDecimal struct {
	*decimal.Decimal
}

// NewDecimal creates a new decimal from a decimal
func NewDecimal(d decimal.Decimal) Decimal {
	return Decimal{Decimal: d}
}

// NewNullDecimal creates a new null decimal from a decimal
func NewNullDecimal(d *decimal.Decimal) NullDecimal {
	return NullDecimal{Decimal: d}
}

// Randomize implements sqlboiler's randomize interface
func (d *Decimal) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	d.Decimal = *randomDecimal(nextInt, fieldType, false)
}

// Value implements driver.Valuer.
func (n NullDecimal) Value() (driver.Value, error) {
	if n.Decimal == nil {
		return nil, nil
	}
	return n.Decimal.Value()
}

// Scan implements sql.Scanner.
func (n *NullDecimal) Scan(val interface{}) error {
	if n == nil {
		return fmt.Errorf("%T is nil", n)
	}
	if val == nil {
		n.Decimal = nil
	} else {
		dec := new(decimal.Decimal)
		if err := dec.Scan(val); err != nil {
			return err
		}
		n.Decimal = dec
	}
	return nil
}

// UnmarshalJSON allows marshalling JSON into a null pointer
func (n *NullDecimal) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		if n != nil {
			n.Decimal = nil
		}
		return nil
	}

	n.Decimal = new(decimal.Decimal)
	return n.Decimal.UnmarshalJSON(data)
}

// MarshalText marshals a decimal value
func (n NullDecimal) MarshalText() ([]byte, error) {
	if n.Decimal == nil {
		return nullBytes, nil
	}

	return n.Decimal.MarshalText()
}

// UnmarshalText allows marshalling text into a null pointer
func (n *NullDecimal) UnmarshalText(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		if n != nil {
			n.Decimal = nil
		}
		return nil
	}

	n.Decimal = new(decimal.Decimal)
	return n.Decimal.UnmarshalText(data)
}

// String impl
func (n NullDecimal) String() string {
	if n.Decimal == nil {
		return "nil"
	}
	return n.Decimal.String()
}

// MarshalJSON marshals a decimal value
func (n NullDecimal) MarshalJSON() ([]byte, error) {
	if n.Decimal == nil {
		return nullBytes, nil
	}

	return n.Decimal.MarshalText()
}

// IsZero implements qmhelper.Nullable
func (n NullDecimal) IsZero() bool {
	return n.Decimal == nil
}

// Randomize implements sqlboiler's randomize interface
func (n *NullDecimal) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	n.Decimal = randomDecimal(nextInt, fieldType, shouldBeNull)
}

func randomDecimal(nextInt func() int64, fieldType string, shouldBeNull bool) *decimal.Decimal {
	if shouldBeNull {
		return nil
	}

	randVal := fmt.Sprintf("%d.%d", nextInt()%10, nextInt()%10)
	random, err := decimal.NewFromString(randVal)
	if err != nil {
		panic("randVal could not be turned into a decimal")
	}

	return &random
}
