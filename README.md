
## Installation Steps

#### Install Go

If you haven't already, you need to install Go. You mentioned that you are using Go version 1.21.1.

You can download and install Go from the official website: 

```
https://go.dev/dl/
```

Make sure Go is properly installed and added to your system's PATH.


### Set up your Go module

You mentioned that your project is hosted on GitHub at `https://github.com/ankitdmon/Image-Compressor` 
To work with your project, set up your Go module as follows:

```
https://github.com/ankitdmon/myproject
 ``` 

This will create a go.mod file in your project directory.

### Install project dependencies

Your `go.mod` file lists the project dependencies. To install these dependencies, run:

```
go mod tidy
```
This command will download and install the necessary dependencies for your project.

### Verify Installation

To verify that all dependencies are installed successfully, you can run:
```
go list -m all
```
This will display a list of all the installed modules and their versions.

### Building and Running Your Project

Now, you should be able to build and run your project. You may have a `main.go` file or an entry point for your project. To run your application, use the `go run` command:

```
go run main.go
```
Replace main.go with the actual entry point of your project.

# Project Structure

This project follows a `Microservices architecture` with two main components: `Consumer` and `Producer`

### Consumer Component

The `Consumer` component includes the following folders and files:

- **`consumer/`** (Consumer component root folder)
  - `main.go` - The entry point for the consumer component.
  - `.env` - Configuration file for environment variables.
  - **`messaging/`** (Messaging related functionality)
    - `consumer.go` - Consumer logic for handling messages.
  - **`utils/`** (Utility functions for the consumer component)
    - `imageUtils.go` - Utility functions for image manipulation.
    - `imageUtils_test.go` - Unit tests for imageUtils.go.


### Producer Component

The `Producer` component includes the following folders and files:

- **`producer/`** (Producer component root folder)
  - **`controllers/`** (Controllers for handling HTTP requests)
    - `controller.go` - Controller logic.
  - **`initializers/`** (Initialization functions for the producer component)
    - `DB.go` - Database initialization.
    - `loadENV.go` - Environment variable loading.
  - **`messaging/`** (Messaging related functionality for the producer)
    - `producer.go` - Logic for producing messages.
  - **`migrate/`** (Database migration scripts)
    - `migrate.go` - Database migration logic.
  - **`models/`** (Data models for the producer component)
    - `product.go` - Product model.
    - `product_test.go` - Unit tests for the product model.
    - `user.go` - User model.
    - `user_test.go` - Unit tests for the user model.
  - **`routes/`** (HTTP route definitions)
    - `route.go` - Route configuration.
  - **`test/`** (Test files for the producer component)
    - `productAPI_test.go` - Integration tests for the product API.
  - **`utils/`** (Utility functions for the producer component)
    - `userData.go` - Utility functions for user data.


In this structure, the `Consumer` and `Producer` components are organized following the `MVC` pattern, with separate folders for models, controllers, views, and utility functions specific to each component. This approach promotes code separation and maintainability.

# Interacting with the System

You can interact with the system using Postman, a popular API testing tool. Follow these steps to use Postman with your project:

## Using Postman

1. Open Postman.

2. Set the endpoint to `http://localhost:3000/product` with `POST` method.

3. Choose one of the following request payloads and enter it into the request body.

### Request Payloads:

#### Product 1
```json
{
  "user_id": 1,
  "product_name": "",
  "product_description": "",
  "product_images": "",
  "product_price": 12000.2
}
```

#### Product 2
```json
{
  "user_id": 1,
  "product_name": "",
  "product_description": "",
  "product_images": "",
  "product_price": 2000.20
}
```
