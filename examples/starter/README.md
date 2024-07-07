# Starter Template

This is a starter template for a simple api using Go, Postgres, JWT and Docker.

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install)
- [Postman](https://www.postman.com/downloads/)
- [Git](https://git-scm.com/downloads)

### Installation

1. Clone the repo:

  ```sh
  git clone https://github.com/9ssi7/gopre.git
  ```

2. Change directory:

  ```sh
  cd gopre/examples/starter
  ```

3. Run the following command to make the environment file:

  ```sh
  make env
  ```

4. Run the following command to make the jwt private and public keys:

  ```sh
  make jwt
  ```

5. Run the following command to start the application:

  ```sh
  make dev
  ```

6. Open Postman public workspace and import the collection:

  [![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/ssi97/workspace/9ssi7-s-public-apis)

## License

Distributed under the Apache License 2.0. See `LICENSE` for more information.
