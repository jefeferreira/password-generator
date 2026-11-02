# Password Generator

A simple and secure password generator written in Go, using `crypto/rand` for cryptographically secure random password generation.


## Usage

Run the application:

```bash
go run ./cmd/cli
```

By default, it generates one **16-character** password containing uppercase letters, lowercase letters, numbers, and symbols.

### Available Flags

| Flag | Default | Description |
| --- | --- | --- |
| `-length` | `16` | Password length |
| `-uppercase` | `true` | Include uppercase letters |
| `-lowercase` | `true` | Include lowercase letters |
| `-numbers` | `true` | Include numbers |
| `-symbols` | `true` | Include special characters |
| `-count` | `1` | Number of passwords to generate |

### Examples

Generate a password with default settings:

```bash
go run ./cmd/cli
```

Generate a **24-character** password:

```bash
go run ./cmd/cli -length 24
```

Generate a password without symbols:


```bash
go run ./cmd/cli -symbols=false
```

Generate a password containing only lowercase letters and numbers:


```bash
go run ./cmd/cli -uppercase=false -symbols=false
```

Generate 5 passwords:

```bash
go run ./cmd/cli -count 5
```

Generate 3 passwords with 20 characters each:


```bash
go run ./cmd/cli -length 20 -count 3
```

## Security

Passwords are generated using Go's `crypto/rand` package to ensure cryptographically secure randomness.

