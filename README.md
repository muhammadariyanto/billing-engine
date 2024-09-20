# Billing Engine

## Overview
This service is designed to handle the core functionality of a billing engine. It processes loan and billing schedule for each loan.

For a detailed explanation of the process, refer to the [Sequence Diagram Process](./docs/sequence_diagram.md).

## How to Run

1. **Configuration Setup**:
    - Copy the `.config.example.yaml` file into a new file named `.config.yaml`.
      ```bash
      cp .config.example.yaml .config.yaml
      ```

2. **Run the Service**:
    - Use the following command to run the HTTP server.
      ```bash
      make run-http
      ```

## Run Unit Tests

To run the unit tests for the service, use the following command:

```bash
make test
```

## Manual Testing
I have prepared a postman collection to carry out testing on this application, please import the file `Amartha - Biling Engine.postman_collection.json` to the postman application.
