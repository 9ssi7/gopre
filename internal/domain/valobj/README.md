# Value Objects

This directory contains the **value objects** that represent descriptive aspects of your application's domain model. Value objects are immutable objects that are defined by their attributes and have no conceptual identity.

## Overview

In Domain-Driven Design (DDD), value objects play a crucial role in modeling concepts that are important in the domain but don't have a distinct identity. They are used to represent attributes of entities or other value objects and are often used to encapsulate complex data structures or business rules.

## Characteristics of Value Objects

* **Immutability:** Value objects are immutable, meaning their state cannot be changed once they are created. This ensures that they are safe to share and pass around without the risk of unexpected side effects.
* **Value Equality:** Value objects are compared by value, not by reference. Two value objects are considered equal if all their attributes have the same values.
* **No Identity:** Value objects do not have a unique identifier. They are defined solely by their attributes.
* **Behavior:** Value objects can have methods that encapsulate domain logic related to their specific value or concept.

## Go Implementation

In Go, value objects are typically represented as structs. Immutability is enforced by making the struct fields unexported (lowercase) and providing constructor functions or methods to create new instances.

## Examples

Here are some examples of value objects you might find in a typical e-commerce application, implemented in Go:

```go
// Address value object
type Address struct {
    street     string
    city       string
    state      string
    postalCode string
    country    string
}

func NewAddress(street, city, state, postalCode, country string) Address {
    return Address{
        street:     street,
        city:       city,
        state:      state,
        postalCode: postalCode,
        country:    country,
    }
}

// Quantity value object
type Quantity int

func NewQuantity(value int) Quantity {
    if value < 0 {
        // Handle negative quantity error
    }
    return Quantity(value)
}
```

## Guidelines for Defining Value Objects

* **Identify domain concepts without identity:** Look for concepts in your domain that are important but don't have a unique identity (e.g., addresses, money amounts, dates).
* **Make value objects immutable:** Ensure that value objects cannot be modified after they are created.
* **Implement value equality:** Define how value objects are compared for equality.
* **Encapsulate domain logic:** If a value object has associated business logic, define methods on the struct to encapsulate that logic.

## Usage in Entities and Aggregates

Value objects are often used as attributes of entities or other value objects. For example, a `Customer` entity might have an `Address` value object as an attribute.

## Further Information

For more detailed information on value objects and their role in DDD, refer to the DDD literature and resources. You can also find examples and best practices for implementing value objects in Go online.
