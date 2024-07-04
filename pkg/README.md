# Shared Packages (pkg)

This directory contains reusable packages and modules that provide specific functionality used across different parts of the application. By centralizing these shared components, we promote code reusability, maintainability, and consistency throughout the project.

## Overview

The `pkg` directory serves as a repository for code that is not specific to any particular domain or application layer. Instead, it houses well-defined building blocks that can be leveraged by multiple components within the system. This approach helps avoid code duplication and ensures that common functionalities are implemented consistently.

## Package Organization

The `pkg` directory is organized into subfolders based on the specific functionality or area of concern of the packages. For example:

* **`errors`:** Contains custom error types and error handling utilities.
* **`http`:** Contains HTTP-related utilities, such as middleware, request/response helpers, or custom HTTP clients.
* **`log`:** Contains logging-related utilities, such as custom loggers or log formatting functions.
* **`database`:** Contains database-related utilities, such as connection pooling, query builders, or transaction management helpers.
* **`validation`:** Contains input validation functions and custom validation rules.
* **`security`:** Contains security-related utilities, such as encryption, decryption, or authentication helpers.
* **`time`:** Contains time-related utilities, such as time formatting or parsing functions.
* **`cqrs`:** (If using CQRS) Contains CQRS-specific components, such as event buses, command dispatchers, or query processors.

## Guidelines for Creating Shared Packages

* **Be specific:** Packages in the `pkg` directory should have a clear and well-defined purpose. Avoid generic names like `utils` or `helpers`.
* **Focus on a single responsibility:** Each package should ideally focus on a single area of concern (e.g., error handling, logging, validation).
* **Test thoroughly:** Shared packages should be thoroughly tested to ensure their correctness and reliability, as they are used by multiple parts of the application.
* **Document well:** Provide clear and concise documentation for each package, including its purpose, usage examples, and any relevant caveats or limitations.
* **Versioning:** Consider using semantic versioning for shared packages to manage their evolution and compatibility with other parts of the application.

## Examples

Here are some examples of packages you might find in the `pkg` directory:

* **`pkg/errors/errors.go`:** Defines custom error types and error handling functions.
* **`pkg/http/middleware.go`:** Contains middleware functions for authentication, logging, or request/response manipulation.
* **`pkg/log/logger.go`:** Implements a custom logger with configurable log levels and output formats.
* **`pkg/database/gorm.go`:** Provides a wrapper around GORM for database connection management and query execution.
* **`pkg/validation/validator.go`:** Implements a validator for input data using custom validation rules.
* **`pkg/security/crypto.go`:** Contains functions for encryption and decryption of sensitive data.
* **`pkg/time/format.go`:** Provides functions for formatting and parsing time values.
* **`pkg/cqrs/eventbus.go`:** (If using CQRS) Defines an interface for an event bus and provides a default implementation.

## Further Information

For more detailed information on the specific packages in your application, refer to the source code and README files within the respective subfolders. You can also consult Go programming resources and best practices for creating reusable packages and modules.
