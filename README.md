
# Telemetry Example Project

## Description

This project demonstrates how to use the `telemetry` package for logging and telemetry data. The example application utilizes different logging drivers based on the configuration specified in `config/config.json`.

## Installation

To set up and run the project, follow these steps:

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/kanojiyadhaval/telemetry_example
   ```

2. **Navigate to the Project Directory:**

   ```bash
   cd telemetry_example
   ```

3. **Install Dependencies:**

   Ensure you have Go installed on your system. Then, run:

   ```bash
   go mod tidy
   ```

## Configuration

1. **Create a Configuration File:**

   Ensure the `config/config.json` file exists with the following content:

   ```json
   {
     "driver_type": "json",
     "file_path": "./logs/log.json"
   }
   ```

2. **Ensure Directory Structure:**

   Create the `logs` directory if it does not already exist:

   ```bash
   mkdir -p ./logs
   ```

   Optionally, create an initial empty `log.json` file in the `logs` directory:

   ```bash
   touch ./logs/log.json
   ```

## Usage

Run the application using the following command:

```bash
go run main.go
```

## Example

This project logs a transaction using the `telemetry` package. The `main.go` file demonstrates how to configure and use different logging drivers based on the settings in `config/config.json`.