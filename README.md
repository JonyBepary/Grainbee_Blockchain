
# GrainBee Project: Chaincode Development

![Hyperledger](https://img.shields.io/badge/hyperledger-2F3134?style=for-the-badge&logo=hyperledger&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Codepen](https://img.shields.io/badge/Codepen-000000?style=for-the-badge&logo=codepen&logoColor=white)
![Markdown](https://img.shields.io/badge/Markdown-000000?style=for-the-badge&logo=markdown&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=Swagger&logoColor=white)
![Flutter](https://img.shields.io/badge/Flutter-02569B?style=for-the-badge&logo=flutter&logoColor=white)
![Latex](https://img.shields.io/badge/Overleaf-47A141?style=for-the-badge&logo=Overleaf&logoColor=white)
![IOT](https://img.shields.io/badge/Arduino_IDE-00979D?style=for-the-badge&logo=arduino&logoColor=white)

## Project Overview

GrainBee is a comprehensive blockchain-based solution designed to manage and track the distribution of rations. The project leverages Hyperledger Fabric to ensure secure, transparent, and efficient ration management. Key components include a robust backend implemented in Go, a user-friendly frontend client app, and a set of APIs for interaction.

## Technologies and Expertise

- **Hyperledger Fabric**: The core blockchain platform used for building the chaincode.
- **Chaincode (CC) Tools**: Utilities and libraries for developing and managing chaincode.
- **Docker**: Containerization technology for consistent development and deployment environments.
- **Go Programming Language**: The primary language used for backend development.
- **Industry Best Practices and Optimizations**: Implementation of best practices for performance and security.
- **Version Control System (VCS) Usage**: Git for version control and collaboration.
- **Unit Testing Implementation**: Comprehensive unit tests to ensure code reliability.
- **Concurrent Programming with Go Routines and Channels**: Efficient handling of concurrent operations.
- **API Management System**: Tools and practices for managing and documenting APIs.

## Architecture

The GrainBee system is built on a modular architecture that includes:

- **Backend**: Implemented in Go, the backend handles all blockchain interactions and business logic.
- **Frontend**: A client application that interacts with the backend via APIs to provide a user-friendly interface.
- **API**: A set of RESTful APIs that facilitate communication between the frontend and backend.

The components interact as follows:
- The frontend sends requests to the backend via the API.
- The backend processes these requests, interacts with the Hyperledger Fabric network, and returns responses to the frontend.

## Key Features

- **Ration Management**: Create, update, and delete rations.
- **Member Management**: Issue and update ration cards for members.
- **Inventory Management**: Replenish and track inventory levels.
- **Distribution Point Management**: Create and manage distribution points.
- **Pickup Schedule Management**: Set and retrieve pickup schedules.
- **Concurrent Updates**: Efficient handling of concurrent updates using Go routines and channels.

## Installation Instructions

### Prerequisites

- Git
- cURL
- Go (version 1.16 or later)
- Docker
- Docker Compose
- Hyperledger Fabric (version 2.x)


### Steps and Usage Instructions

1. **Clone the Repository**
   ```sh
   git clone https://github.com/JonyBepary/Grainbee_Blockchain.git
   cd grainbee
   ```

2. **Install Dependencies**
   ```sh
   go mod tidy
   ```


2. **Download Fabric Binaries**

   If you don't have the Fabric binaries, download them:
   ```sh
   curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
   ./install-fabric.sh --fabric-version 2.5.3 binary
   rm install-fabric.sh
   ```

3. **Create Docker Network**

   Create a Docker network for the Fabric components:
   ```sh
   docker network create cc-tools-demo-net
   ```

4. **Bring Up the Network**

   Start the Fabric network with the desired number of organizations (e.g., 3 organizations):
   ```sh
   ./network.sh up createChannel -n 3
   ```

5. ** Package and Deploy Chaincode**

   Deploy the chaincode to the network by following this [tutorial](https://hyperledger-fabric.readthedocs.io/en/release-1.3/chaincode4noah.html#packaging).
6. **To deploy this chaincode in Amazon Managed Blockchain**

   Follow the documentation in this [link](https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/what-is-managed-blockchain.html) to create a network in Amazon Managed Blockchain and deploy chain code.

### Running the Application Locally

1. **Start the Hyperledger Fabric Network**
   ```sh
   ./startDev.sh -n 1
   ```

2. **Run the Backend Server**
   ```sh
   go run main.go
   ```

3. **Access the Frontend**
   - Open the frontend client app (access through http://localhost:3000).

### Common Operations

- **Create a New Ration**
  ```sh
  curl -X POST http://localhost:8080/api/createRation -d '{"id": "ration1", "category": "rice", "description": "10kg rice", "package": "bag", "distributedBy": "dist1", "quantity": 10, "expiryDate": "2023-12-31", "mfgDate": "2023-01-01", "batchNumber": 1}'
  ```

- **Update Member Information**
  ```sh
  curl -X PUT http://localhost:8080/api/updateMemberInfo -d '{"nid": "12345", "name": "John Doe", "dateOfBirth": "1990-01-01", "height": 170, "address": "123 Main St", "contactInformation": "john@example.com", "familySize": 4, "income": 50000, "disabilityStatus": false, "rationCardNumber": "RC1234", "rationCardStatus": "active", "rationCardIssuedDate": "2023-01-01", "rationCardExpiryDate": "2023-12-31", "rationCardCategory": "general"}'
  ```

## API Documentation

### Overview

The GrainBee API provides endpoints for managing rations, members, inventory, distribution points, and pickup schedules.

### Authentication

- **API Key**: Include the API key in the `Authorization` header of your requests.

### Endpoints

- **Create Ration**: `POST /api/createRation`
- **Update Member Info**: `PUT /api/updateMemberInfo`
- **Replenish Inventory**: `POST /api/replenishInventory`
- **Set Pickup Schedule**: `POST /api/setPickupSchedule`
- **Get Pickup Schedule**: `GET /api/getPickupSchedule`

## Testing

### Running Unit Tests

1. **Navigate to the Project Directory**
   ```sh
   cd grainbee
   ```

2. **Run the Tests**
   ```sh
   go test ./...
   ```

### Additional Testing Procedures

- **Mock Tests**: Utilize mock stubs to simulate blockchain interactions.
- **Integration Tests**: Ensure that the backend and frontend components interact correctly.

## Blockchain Visualization

### Overview of Blockchain Structure

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/10.jpeg?raw=true" width="1280" alt="Overview of Blockchain Structure">

### Full view of the entry form

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/4.png?raw=true" width="1280" alt="Full view of the entry form">

### Manage Members

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/5.png?raw=true" width="1280" alt="Manage Members">


### Detailed View of a Member Assets

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/1.png?raw=true" width="1280" alt="Detailed View of a Member Assets">


### Detailed View of a Ration Card

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/2.png?raw=true" width="1280" alt="Detailed View of Detailed View of a Ration Card">

### Curl Commands for Blockchain Operations

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/3.png?raw=true" width="1280" alt="Curl Commands for Blockchain Operationin Blockchain">

### Create Distributor

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/6.png?raw=true" width="1280" alt="Create Distributor">

### Ration Management

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/7.png?raw=true" width="1280" alt="Ration Management">

### Get Member QR Code

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/8.png?raw=true" width="1280" alt="Get Member QR Code">

### Read Member Information

<img src="https://github.com/JonyBepary/Grainbee_Blockchain/blob/main/images/9.png?raw=true" width="1280" alt="Read Member Information">


## Contributing

### Guidelines

- Follow the existing code style and conventions.
- Write clear and concise commit messages.
- Include unit tests for new features and bug fixes.

### Submitting Pull Requests

1. **Fork the Repository**
   A fork is a new repository that shares code and visibility settings with the original “upstream” repository. follow this [link](https://docs.github.com/en/get-started/quickstart/fork-a-repo) to fork the repository.

2. **Create a New Branch**
   ```sh
   git checkout -b feature/new-feature
   ```

3. **Make Your Changes**
   - Implement the new feature or bug fix.
   - Write unit tests for the new code.
   - Update the documentation as needed.
4. **Commit Your Changes**
   ```sh
   git commit -m "Add new feature"
   ```

5. **Push Your Changes**
   ```sh
   git push origin feature/new-feature
   ```

6. **Open a Pull Request**

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact Information

For support or inquiries, please contact the maintainers at [sohelahmedjony@gmail.com](mailto:sohelahmedjony@gmail.com).
