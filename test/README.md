# Tests

This directory contains automated tests for the application, ensuring the correctness, reliability, and maintainability of the codebase. Tests are essential for verifying that the application behaves as expected and for catching regressions (bugs introduced by new code changes).

## Overview

The `test` directory is organized to mirror the structure of the main application code. This makes it easier to locate the tests for a specific package or module. The tests are categorized into different types based on their scope and purpose:

* **Unit Tests:** Unit tests focus on testing individual functions, methods, or components in isolation. They verify that each unit of code works correctly in isolation from the rest of the system.
* **Integration Tests:** Integration tests verify the interactions between different components or modules of the application. They ensure that the components work together as expected.
* **End-to-End (E2E) Tests:** E2E tests simulate real user scenarios and test the entire application flow from start to finish. They ensure that the application behaves correctly when all components are integrated.

## Test Organization

The `test` directory typically has subdirectories that correspond to the main application directories:

* **`domain`:** Contains unit tests for the domain layer (aggregates, entities, value objects).
* **`application`:** Contains unit and integration tests for the application layer (command handlers, query handlers, application services).
* **`infrastructure`:** Contains integration tests for the infrastructure layer (repositories, message queues, external service integrations).
* **`api`:** Contains end-to-end tests for the API layer (HTTP, GraphQL, gRPC).

## Test Frameworks and Tools

The following frameworks and tools are commonly used for writing tests in Go:

* **Testing Package:** Go's built-in `testing` package provides the basic framework for writing and running tests.
* **Testify:** A popular assertion library that provides a more expressive and readable way to write test assertions.
* **Gomock:** A mocking framework for creating mock objects for testing interactions with external dependencies.
* **SQLMock:** A library for mocking database interactions in tests.
* **httptest:** A package for testing HTTP handlers and clients.

## Test Examples

### Unit Test (Domain Layer)

```go
func TestOrder_AddOrderItem(t *testing.T) {
    // ... (Test setup)

    err := order.AddOrderItem("product1", 2, 10.99)
    require.NoError(t, err)
    assert.Equal(t, 20.98, order.TotalAmount)
    assert.Len(t, order.OrderItems, 1)
    assert.Equal(t, "product1", order.OrderItems[0].ProductID)

    // ... (More test assertions)
}
```

### Integration Test (Application Layer)

```go
func TestPlaceOrderHandler(t *testing.T) {
    // ... (Test setup with mock repositories and event bus)

    cmd := commands.PlaceOrder{
        // ... (Command data)
    }
    err := handler.Handle(cmd)
    require.NoError(t, err)

    // ... (Verify that the order was created and events were published)
}
```

### E2E Test (API Layer)

```go
func TestCreateOrderAPI(t *testing.T) {
    // ... (Start the API server)

    resp, err := http.Post("/api/commands/orders", "application/json", strings.NewReader(orderJSON))
    require.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)

    // ... (Verify the response body and database state)
}
```

## Running Tests

To run the tests, use the following command:

```bash
go test ./...
```

This will run all tests in the project, including unit, integration, and E2E tests.

## Test Coverage

Test coverage is a metric that measures the percentage of code that is executed by the tests. It is a useful indicator of the thoroughness of your tests. To generate a test coverage report, use the following command:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This will generate an HTML report that you can open in your browser to see which parts of your code are covered by tests.
