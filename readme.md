# Payment Processor App

This Go web service is designed to process payment for a cardholder.The application is containerized using Docker for easy deployment and management.

## Prerequisites

Before running this application, ensure you have the following dependencies installed:

- Docker
- Docker Compose


## Configuration

This application relies on environment variables for configuration. Please ensure the following environment variables are properly set before running the application:

These environment variables can be set either directly in your system environment or by creating a `.env` file in the root directory of the project.

Example `.env` file:

PORT = 3000

DATABASE_URI = postgres://postgres:password@db:5432/postgres?sslmode=disable

POSTGRES_USER = postgres

POSTGRES_PASSWORD = password

TEST_URL = http://localhost


## Usage

To run the application, follow these steps:

1. Clone this repository to your local machine.
2. Navigate to the root directory of the project.
3. Set the required environment variables as described in the Configuration section.
4. Run the following command to build and start the Docker containers:`docker compose up -d --build`
5. Once the containers are up and running, the application will be accessible at `http://localhost:3000`.


## Tests
To run tests for the application, follow these steps:
1. Navigate to the root directory of the project.
2. Run the following command to run the tests: `docker compose up -d --build`
3. Run the command `docker exec -it payment-processor go test /usr/src/app/tests/`

## Docker Deployment

This application is Dockerized, making it easy to deploy in various environments. To deploy the application using Docker, follow the steps outlined in the Usage section above.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or create a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [MIT](LICENSE) file for details.