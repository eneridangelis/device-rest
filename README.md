# Device Management Service

This project is a REST API built in Go for managing devices, with endpoints for creating, reading, updating, deleting, and searching devices by brand. The application uses PostgreSQL as its database and GORM as the ORM.

## Table of Contents

- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Running the Project](#running-the-project)
- [Makefile Commands](#makefile-commands)
- [API Endpoints](#api-endpoints)

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Postman](https://www.postman.com/downloads/)

### Project Structure

The project follows a clean architecture pattern with layers for:
- `api`: Handles the HTTP requests and routing.
- `usecase`: Contains business logic.
- `repository`: Manages database interactions.
- `model`: Defines the data structures used across the layers.
- `cmd`: Entry point to the application.
- `migration`: Contains migration scripts for database setup.

### Running the Project

#### Step 1: Start the Database with Docker

Run the follow command: `make infra-up`

#### Step 2: Run Database Migration

Run the command: `make migrate-up`

#### Step 3: Run the application

Run the command: `make run`

The API should now be running at http://localhost:8080.

### Makefile Commands
The project includes a Makefile to simplify tasks:

`make infra-up`: Starts the necessary infrastructure (PostgreSQL) using Docker.

`make run`: Runs the application.

`make test`: Runs the tests.

`make install-migrate`: Installs the golang-migrate tool.

`make migrate-up`: Applies the migrations to the database.

`make migrate-down`: Rolls back the last migration.

`make create-migration`: Creates a new migration file.

## API Endpoints

### 1. Add Device

- **Endpoint**: `POST /devices`
- **Description**: Adds a new device to the database.
- **Request Body**:
  ```json
  {
    "name": "Device Name",
    "brand": "Device Brand",
  }
- **Response**:
    - **Status Code**: 201 Created

### 2. Get Device by ID

- **Endpoint**: `GET /devices/{id}`
- **Description**: Fetches a device by its ID.
- **Path Parameters**:
  - **id**: The ID of the device.
- **Example**: `GET http://localhost:8080/devices/1`
- **Response**:
    - **Status Code**: 200 OK
    - **Response Body**: 
    ```json
    {
      "id": 1,
      "name": "Device Name",
      "brand": "Device Brand",
      "created_at": "2024-10-01T00:00:00Z",
    }
    
### 3. List All Devices

- **Endpoint**: `GET /devices`
- **Description**: Retrieves a list of all devices.
- **Example**: `GET http://localhost:8080/devices`
- **Response**:
    - **Status Code**: 200 OK
    - **Response Body**: 
    ```json
    [
        {
            "id": 1,
            "name": "Device Name",
            "brand": "Device Brand",
            "created_at": "2024-10-01T00:00:00Z",
        },
        {
            "id": 2,
            "name": "Device Name 2",
            "brand": "Device Brand 2",
            "created_at": "2024-10-01T00:00:00Z",
        }
    ]

### 4. Update Device

- **Endpoint**: `PATCH /devices`
- **Description**: Updates an existing device based on the provided information, total or partial. Only fields present in the body are updated.
- **Request Body**:
  ```json
  {
    "name": "Device Name",
    "brand": "Device Brand",
  }
- **Response**:
    - **Status Code**: 204 No Content

### 5. Delete Device

- **Endpoint**: `DELETE /devices/{id}`
- **Description**: Deletes a device by its ID.
- **Path Parameters**:
  - **id**: The ID of the device.
- **Example**: `DELETE http://localhost:8080/devices/1`
- **Response**:
    - **Status Code**: 204 No Content

### 6. Search Devices by Brand

- **Endpoint**: `GET /devices/search`
- **Description**: Searches for devices by their brand.
- **Query Parameters**:
  - **brand**: Brand name to filter devices by.
- **Example**: `GET http://localhost:8080/devices/search?brand=BrandName`
- **Response**:
    - **Status Code**: 200 OK
    - **Response Body**: 
    ```json
    {
      "id": 1,
      "name": "Device Name",
      "brand": "BrandName",
      "created_at": "2024-10-01T00:00:00Z",
    }
    