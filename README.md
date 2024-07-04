# gopre: A Go Project Template for CQRS and DDD

**gopre** is a comprehensive Go project template designed to facilitate the development of robust and scalable applications using the principles of Domain-Driven Design (DDD) and the Command Query Responsibility Segregation (CQRS) architectural pattern. It provides a well-structured foundation for building complex applications with a focus on maintainability, testability, and adaptability.

## Overview

This template embodies a layered architecture that strictly adheres to DDD principles. It encourages a clear separation of concerns between the domain, application, and infrastructure layers, ensuring that the core business logic remains isolated from technical implementation details.

While the template promotes the use of CQRS, it is not mandatory. You can choose to implement CQRS fully, partially, or not at all, depending on your project's specific requirements.

## Key Features

* **Domain-Driven Design (DDD):** The project structure is aligned with DDD principles, emphasizing the importance of the domain model and its ubiquitous language.
* **Command Query Responsibility Segregation (CQRS):** The template provides a clear separation of commands (actions that change the application state) and queries (actions that retrieve data), facilitating independent scaling and optimization of each aspect.
* **Layered Architecture:** The codebase is organized into distinct layers (domain, application, infrastructure), promoting modularity and maintainability.
* **Clean Architecture:** The template follows the principles of Clean Architecture, ensuring that the domain layer remains independent of external frameworks and technologies.
* **API Flexibility:** Supports multiple API styles, including RESTful APIs (using Gin), GraphQL APIs (using gqlgen), and gRPC APIs (using Protocol Buffers).
* **Testability:** The project includes a comprehensive suite of unit, integration, and end-to-end tests to ensure code quality and prevent regressions.
* **Docker Integration:** Dockerfiles and a `docker-compose.yml` file are provided for easy containerization and deployment.

## Project Structure

The project follows a standard Go project layout with the following key directories:

* **`api`:** Contains the implementation of different API types (HTTP, GraphQL, gRPC, OpenAPI).
* **`cmd`:** Contains the main entry points for the application's services (cmdsrv, querysrv, or myapp).
* **`internal`:** Contains the core application layers (application, domain, infrastructure).
* **`pkg`:** Contains shared packages and modules used throughout the application.
* **`test`:** Contains unit, integration, and end-to-end tests.

## Suitable Projects

gopre is well-suited for:

* **Complex business domains:** Where a clear and well-structured domain model is essential.
* **Scalable applications:** Where independent scaling of reads and writes is beneficial.
* **Evolving requirements:** Where the ability to adapt to changing business needs is crucial.
* **Collaborative teams:** Where a well-defined architecture can facilitate communication and collaboration among developers.

## Not Suitable For

* **Simple CRUD applications:** Where the overhead of CQRS and DDD might not be justified.
* **Projects with tight deadlines:** Where the initial setup and learning curve of CQRS and DDD might be a constraint.

## Getting Started

1. Clone the repository.
2. Open a terminal and navigate to the examples directory.
3. Run the following command to start the application:

```bash
make run
```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
