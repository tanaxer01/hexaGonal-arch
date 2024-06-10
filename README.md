
This is a example proyect in which a hexagonal architecture is used to implement
a FeatureFlag storing solution, this __"feature flags"__ are stored in a mongoDB instance
and can be fetched through graphql.

````sh
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── federation.go
│   ├── generated.go
│   ├── model
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphqls
│   └── schema.resolvers.go
├── internal
│   ├── adapters
│   │   ├── handler
│   │   │   └── featureFlag.handler.go
│   │   └── repository
│   │       └── mongoDB.adapter.go
│   └── core
│       ├── domain                      // All entities needed 4 the business logic.
│       │   └── model.go
│       ├── ports                       // All entities needed 4 the business logic.
│       │   └── ports.go
│       └── services                    // Services are the entry points 4 the core, they implement
│           └── featureFlags.service.go // Ifaces 2 communicate with actors.
├── server.go
└── tools.go
```

(In this case __services__ would be the same as __features__ in some similar proyects)
