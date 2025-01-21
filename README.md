# Custom types for SQLBoiler

This module contains some (one for now) custom type(s) to be used with SQLBoiler.

## Types

- `Shopspring/Decimal`: Replacement for SQLBoilers decimal type using [shopspring/decimal](github.com/shopspring/decimal) instead of [ericlagergren/decimal](github.com/ericlagergren/decimal).

## Usage

Add the following to your `sqlboiler.toml`:

```toml
[[types]]
  [types.match]
    type = "types.Decimal"
    nullable = false

  [types.replace]
    type = "shopspring.Decimal"

  [types.imports]
    third_party = ['"github.com/Dadido3/sqlboiler-custom-types/shopspring"']

[[types]]
  [types.match]
    type = "types.NullDecimal"
    nullable = true

  [types.replace]
    type = "shopspring.NullDecimal"

  [types.imports]
    third_party = ['"github.com/Dadido3/sqlboiler-custom-types/shopspring"']
```
