# Internal Layer

This directory contains the core components of the application, structured according to the principles of Domain-Driven Design (DDD) and potentially the Command Query Responsibility Segregation (CQRS) pattern.

## Overview

The `internal` layer is divided into three main subdirectories:

* **`app`:** This layer contains the application-specific logic, including command and query handlers. It acts as a mediator between the external API layer and the internal domain and infrastructure layers.
* **`domain`:** This layer encapsulates the core business logic of the application, including aggregates, entities, value objects, and domain events. It defines the language and rules of the business domain.
* **`infra`:** This layer provides the technical implementation details for interacting with external systems, such as databases, message queues, and external services.

## Subfolders

### `application`

This subfolder contains two main subfolders:

* **`commands`:** Contains the command handlers that process commands (actions that change the application state) and interact with the domain layer to enforce business rules.
* **`queries`:** Contains the query handlers that process queries (requests for data) and retrieve the necessary information from the domain layer.

The separation of commands and queries is a key aspect of CQRS, promoting better organization, maintainability, and scalability of the application logic.

### `domain`

This subfolder contains the following subfolders:

* **`aggregates`:** Contains the aggregate root entities, which are the primary building blocks of the domain model. Aggregates encapsulate related entities and value objects and ensure the consistency of the domain data.
* **`entities`:** Contains the individual entities that make up the aggregates. Entities have their own identity and lifecycle, and they represent the core concepts of the domain.
* **àbstracts`:** Contains abstract interfaces that define common behavior or contracts for domain objects.
* **`valueobjects`:** Contains value objects, which are immutable objects that represent a specific value or concept in the domain (e.g., an address, a money amount, or a date range).
* **`events`:** (If using CQRS) Contains the domain events that are emitted by aggregates when their state changes. These events are used to communicate changes between the command and query sides of the application.

### `infra`

This subfolder contains the following subfolders:

* **`messaging`:** (If using CQRS) Contains the messaging infrastructure for publishing and consuming domain events.
* **`repositories`:** Contains the repositories that provide access to the underlying data storage (e.g., a database) for the domain objects.
* **`db`:** Contains the database-specific implementation details, such as database schema definitions, migrations, and data access code.

## CQRS Considerations

If the application uses CQRS, the `internal` layer plays a crucial role in implementing the separation of concerns between commands and queries. The command handlers interact with the domain model to process commands and emit events, while the query handlers use a separate data model (potentially a different database) to efficiently serve queries.

## Further Information

For more detailed information on the specific implementation of the application, domain, and infrastructure layers, refer to the README files and source code within the respective subdirectories.
