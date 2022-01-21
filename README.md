# recon

Short for reactive config. A library for managing and validating configuration files for Go. Supports loading configuration from files (JSON and YAML), SQL database, and REST API endpoints. Supports caching with redis.

## Loaders

## Validation

Available rules

### Boundary

#### Numeric boundaries

Available for all numeric types: `rune`, `byte`, `float32`, `float64`, and sizes of `uint` and `int`

`minimum: uint` asserts the value to be higher or equal 

`maximum: uint` asserts the value to be lower or equal

`exclusiveMinimum: uint` asserts the value to be higher

`exclusiveMaximum: uint` asserts the value to be lower

#### String validation

Available for `string` type

`minLength: uint` asserts the length of the value to be longer or equal

`maxLength: uint` asserts the length of the value to be shorter or equal

`pattern: string (regex)` asserts the value to pass the given regular expression rule

`format: email | hostname | ip | uri` asserts the format

#### Array validation

`minItems: uint` asserts the number of items to be greater than or equal to this number

`maxItems: uint` asserts the number of items to be less than or equal to this number

`uniqueItems` asserts that all of the items in this array are unique

#### Map validation

`minProperties: uint` asserts the number of records in this map to be greater than or equal to this number 

`maxProperties: uint` asserts the number of records in this map to be less than or equal to this number 

`required: []string` asserts that this map contains the given set of values as keys