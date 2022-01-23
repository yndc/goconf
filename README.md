# recon

Short for reactive config. A library for managing and validating configuration files for Go. Supports loading configuration from files (JSON and YAML), SQL database, and REST API endpoints. Supports caching with redis.

## Loaders

`recon` supports loading configuration from different sources. 

### File loaders

`recon` supports loading data from multiple formats. We will deduce the format by their file extensions:
- JSON (`.json`)
- YAML (`.yaml` or `.yml`)
- TOML (`.toml`)
- environment variables for other file extensions

Please check out environment variables mapping if you have nested fields. 

### File Loader

### SQL Database

### REST API

## Config Mapping

`recon` supports nested objects, however some data sources does not support nested objects. Therefore when loading configuration from sources that does not support them, we will map the key to a nested field.

By default, nested fields can be accessed or defined using the character `.`. So if you have an example config like this: 

```
{
    "connections": {
        "postgres": {
            "host": "localhost"
        }
    }
}
```

You can access the deeply nested `"host"` by using `connections.postgres.host`

Arrays can be accessed using the `[INDEX]` format. For example to get the second item from the `connections` array in this configuration:
```
{
    "connections": [
        {
            "type": "postgres",
            "host": "localhost"
        },
        {
            "type": "mysql",
            "host": "localhost2"
        }
    ]
}
```
You can use `connections[1].host`

### Environment variables mapping

Environment variables are tricky, that's why we can't go with the convention over configuration approach. 

There are two major issues when mapping environment variables to a field in a Golang object. 
1. Potential naming clashes with other variables
2. Nested fields (dot character is not a valid character for a variable name)

To solve the first problem, we are giving you an option to setup a prefix for the environment variables mapping. With this prefix set, `recon` will use this prefix when checking and mapping the variables. For example, if you set the prefix as `MY_PREFIX` and you're trying to load a field named `myValue`, then `recon` will watch the environment variable with the name `MY_PREFIX_myValue`. 

Since we can't use the dot character `.` in environment variables, we are using double underscores (`__`) to access deep fields by default. You can override this character when building the schema. So for example, if you set the prefix as `MYAPP` and you're trying to load a field named `postgres` inside an object from `database`, then `recon` will watch the environment variable with the name `MYAPP_database__postgres`. 

## Cache (Optional)

By default `recon` will keep a state of the config in the memory, however if a redis connection is given, `recon` will store the state of the config in `redis`.

## Defaults and Optionals

Go struct tags can be used to set a default value for a field. The format is `default:"DEFAULT_VALUE"`. This is only available on primitive types.

All fields are required by default. To allow optional fields, it is required to set the field as a pointer type (`*`).

Composite types such as arrays and structs (nested objects) are the exception for this behaviour. Arrays can always be empty (as a zero-length array). However it is also possible to add a validation rule to make sure that the array is not empty. See below for the `minItems` rule. 

Structs however need to be value types, e.g. it's not possible to have a pointer to struct in the configuration struct definition. To have a struct as an optional, then all of its fields need to be an optional.

## Validation

Go struct tags can also be used to add a simple validation over the field values. 

#### Numeric validation

Available for all numeric types: `rune`, `byte`, `float32`, `float64`, and sizes of `uint` and `int`
- `minimum: uint` asserts the value to be higher or equal 
- `maximum: uint` asserts the value to be lower or equal
- `exclusiveMinimum: uint` asserts the value to be higher
- `exclusiveMaximum: uint` asserts the value to be lower

#### String validation

Available for `string` type
- `minLength: uint` asserts the length of the value to be longer or equal
- `maxLength: uint` asserts the length of the value to be shorter or equal
- `pattern: string (regex)` asserts the value to pass the given regular expression rule
- `format: email | hostname | ip | uri` asserts the format

#### Array validation
- `minItems: uint` asserts the number of items to be greater than or equal to this number
- `maxItems: uint` asserts the number of items to be less than or equal to this number
- `uniqueItems` asserts that all of the items in this array are unique

#### Map validation
- `minProperties: uint` asserts the number of records in this map to be greater than or equal to this number 
- `maxProperties: uint` asserts the number of records in this map to be less than or equal to this number 
- `required: []string` asserts that this map contains the given set of values as keys