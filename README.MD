GENERATION:
gqlgen generate --config codegen/gqlgen-merchant.yml
gqlgen generate --config codegen/gqlgen-retailer.yml


src/
├── schemas/          → GraphQL schema + generated code
├── internal/         → Resolvers (business logic)
├── merchant/         → Merchant server (main.go)
├── retailer/         → Retailer server (main.go)
├── common/           → Shared structs/models


RUNNING SERVER:
go run src/merchant/main.go
go run src/retailer/main.go