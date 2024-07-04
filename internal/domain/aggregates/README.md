# Aggregates

This directory contains the **aggregates** that represent the primary building blocks of your application's domain model. In Domain-Driven Design (DDD), aggregates are clusters of associated objects that are treated as a single unit for data changes.

## Overview

Aggregates help maintain data consistency and encapsulate domain logic. They consist of one or more entities and value objects that are closely related and have a consistent lifecycle. The aggregate root is the main entry point for interacting with the aggregate and is responsible for enforcing business rules and maintaining the aggregate's internal consistency.

## Characteristics of Aggregates

* **Consistency Boundaries:** Aggregates define transaction boundaries. Changes to multiple objects within an aggregate are treated as a single atomic operation, ensuring that the aggregate remains in a consistent state.
* **Root Entity:** Each aggregate has a single root entity, also known as the aggregate root. The root entity is the only object within the aggregate that can be referenced directly from outside the aggregate.
* **Internal Structure:** Aggregates can contain other entities and value objects. These objects are only accessible through the aggregate root and are considered part of the aggregate's internal implementation.
* **Domain Logic:** Aggregates encapsulate the business logic and rules related to the aggregate's domain. This logic is implemented as methods on the aggregate root and ensures that the aggregate's state remains valid.

## Go Implementation

In Go, aggregates are typically represented as structs. The aggregate root is the main struct, and it can contain fields that reference other entities or value objects within the aggregate.

## Example

```go
// domain/aggregates/order.go
package aggregates

import (
    "time"

    "your-project/domain/entities"
    "your-project/domain/valueobjects"
)

type Order struct {
    ID            string              `gorm:"primaryKey"`
    CustomerID    string              `gorm:"not null"`
    Status        valueobjects.Status `gorm:"not null"`
    OrderItems    []entities.OrderItem
    ShippingInfo  entities.ShippingInfo `gorm:"embedded"`
    BillingInfo   entities.BillingInfo  `gorm:"embedded"`
    TotalAmount   float64             `gorm:"not null"`
    CreatedAt     time.Time           `gorm:"autoCreateTime"`
    UpdatedAt     time.Time           `gorm:"autoUpdateTime"`
}

// (Methods for adding/removing items, updating shipping/billing info, etc.)
```

In this example, `Order` is the aggregate root. It contains references to `OrderItem` entities and `ShippingInfo` and `BillingInfo` value objects. The `Order` struct also has methods for interacting with the aggregate, such as adding or removing items, updating shipping or billing information, and changing the order status.

## CQRS and Aggregates

In a CQRS architecture, aggregates are primarily used on the command side of the application. They are responsible for processing commands and ensuring that the domain rules are enforced. When the state of an aggregate changes, it emits domain events that are used to update the query model.

## Event Synchronization in CQRS

In a CQRS setup with separate databases for commands and queries, domain events play a crucial role in synchronizing data between the two models. When an aggregate's state changes, it publishes domain events to a message broker (e.g., RabbitMQ, Kafka). The query side subscribes to these events and updates its own data model accordingly.

## Further Information

For more detailed information on aggregates and their role in DDD, refer to the DDD literature and resources. You can also find examples and best practices for implementing aggregates in Go online.
