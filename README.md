
# HTTP Gateway with gRPC Backend


This project provides an HTTP gateway with a gRPC backend implementation. It allows you to interact with the backend services using HTTP requests.

## Getting Started

  
These instructions will help you get the project up and running on your local machine.

  
### Prerequisites

  
- Docker: Make sure you have Docker installed on your system. You can download Docker from the official website: [https://www.docker.com](https://www.docker.com)

  

### Installation

1. Clone the repository:


```shell

git clone https://github.com/jeferagudeloc/grpc-http-gateway.git

```

Navigate to the project directory:

  

```shell

cd grpc-http-gateway

```

Build and run the containers using Docker Compose:

  

```shell

docker-compose up --build

```

The HTTP gateway service will be accessible at http://localhost:8080.
The React app will be accessible at http://localhost:3000.


## Project Structure

The project follows a typical directory structure:

  

*src/server:* Contains the implementation of the gRPC backend server.

*src/gateway:* Contains the implementation of the HTTP gateway.

*src/ui:* Contains the portal for test the gateway.

## Configuration

The project uses Docker Compose for container orchestration. The services and their configurations are defined in the docker-compose.yml file.


***server:*** Builds and runs the gRPC backend server.

***gateway:*** Builds and runs the HTTP gateway service.

***mysql:*** Runs the MySQL database server.

***web:*** Runs the web server

The environment variables for the services can be configured in the docker-compose.yml file.


### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

  

### License

This project is licensed under the MIT License.
