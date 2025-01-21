# Custom types for SQLBoiler

This module contains some (one for now) custom type(s) to be used with SQLBoiler.

## Types

- `shopspring/Decimal`: Replacement for SQLBoilers decimal type using [shopspring/decimal](https://github.com/shopspring/decimal) instead of [ericlagergren/decimal](https://github.com/ericlagergren/decimal).

## Usage

Add the following to your `sqlboiler.toml`:

```toml
# Use shopspring/decimal instead of ericlagergren/decimal.
[[types]]
  match.type = "types.Decimal"
  match.nullable = false
  replace.type = "shopspring.Decimal"
  imports.third_party = ['"github.com/Dadido3/sqlboiler-custom-types/shopspring"']

[[types]]
  match.type = "types.NullDecimal"
  match.nullable = true
  replace.type = "shopspring.NullDecimal"
  imports.third_party = ['"github.com/Dadido3/sqlboiler-custom-types/shopspring"']
```
