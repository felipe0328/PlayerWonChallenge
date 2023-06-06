# PlayerWon Take Home Challenge

## Description
WebService created to deliver Ads based on **used ID**, **User Country** and player **language** based on the requirements described in [this document](https://gist.github.com/victorhurdugaci/c168d5e8e5befaac1dbf334470166220).

This project uses an in memorySQLDB to avoid the need of implementing a DB.

## Used Libraries
- [Gin](https://gin-gonic.com)
- [RamSQL](https://github.com/proullon/ramsql)
- [Swaggo](https://github.com/swaggo/swag)
- [Mockery](https://github.com/vektra/mockery)
- [Testify](https://github.com/stretchr/testify)
- [fmt](https://pkg.go.dev/fmt)
- [GolangCi-Lint](https://golangci-lint.run/)


## Paths
| Request Type  | Endpoint              | Description                   |
|--             |--                     |--                             |
| GET           | `/docs/index.html`    | Project swagger documentation |
| POST          | `/ads`                | Receive a new video           |


## Project Commands
| Command               | Description           |
| --                    |--                     |
| `go run main.go`      | Run debug server      |
| `go test ./...`       | Run project Unit tests|
| `go fmt ./...`        | Format code           |
| `golangci-lint run`   | Run golang linter     |