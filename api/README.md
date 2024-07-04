# API Layer

This layer contains the various APIs through which the application interacts with the external world.  It enables clients to access the application's functionality using different protocols and data formats.

## Subfolders

Each subfolder represents a different API type and contains the implementation details specific to that API:

* **`cli`:** (If applicable) Contains the command-line interface (CLI) for interacting with the application. See the [`cli/README.md`](./cli/README.md) file for details.
* **`http`:** Contains handlers and route definitions for RESTful APIs using the HTTP protocol. See the [`http/README.md`](./http/README.md) file for details.
* **`graphql`:** Contains GraphQL schema definitions and resolvers for the GraphQL API. See the [`graphql/README.md`](./graphql/README.md) file for details.
* **`openapi`:** (If applicable) Contains OpenAPI (Swagger) specifications for documenting and generating client code for the RESTful APIs. See the [`openapi/README.md`](./openapi/README.md) file for details.
* **`rpc`:** (If applicable) Contains gRPC service definitions (`*.proto` files) for the gRPC API. See the [`rpc/README.md`](./rpc/README.md) file for details.

## CQRS and APIs

If the application adopts the Command Query Responsibility Segregation (CQRS) architectural pattern, this separation should be reflected in the API layer as well. This means that commands (actions that change the application state) and queries (actions that retrieve data) should be handled by different APIs or endpoints.

In a RESTful API context, this could mean using different HTTP methods for commands (e.g., `POST`, `PUT`, `DELETE`) and queries (`GET`). Alternatively, you could use different URL paths (e.g., `/api/commands/...` for commands and `/api/queries/...` for queries).

In a GraphQL API context, this separation is typically achieved by using `Mutation` types for commands and `Query` types for queries.

Please refer to the README files within each API subfolder for more specific implementation details and guidance on how CQRS is applied within that API type.
