# Device Management API

## Description
RESTful API for device management. Supports operations for adding, searching, listing, updating, deleting, and searching by device brand.

## About the Author
- Name: Fernanda Melo
- Date: 08/02/2024

## Technologies Used
- Go
- PostgreSQL
- Docker
- Docker Compose

## Prerequisites
- Docker
- Docker Compose
- DBeaver (optional, for table visualization)

## Project Setup

### Step 1: Clone the Repository
Clone the repository to your local machine:

```bash
git clone https://gitlab.com/seu-usuario/device-management-api.git
cd device-management-api
```

### Step 2: Run Docker Compose
Navigate to the compose directory and run Docker Compose to build and start the containers:

```bash
cd compose
docker-compose up --build
```
### Step 3: Verify the Application
The application will be accessible at http://localhost:8080.

## API Endpoints
- POST /devices: Adds a new device.
- GET /devices/{id}: Fetches a device by identifier.
- GET /devices: Lists all devices.
- PUT /devices/{id}: Fully updates a device.
- PATCH /devices/{id}: Partially updates a device.
- DELETE /devices/{id}: Deletes a device.
- GET /devices/brand/{brand}: Searches devices by brand.

## Run Tests
To run the tests, use the command:

```bash
go test ./tests
```



