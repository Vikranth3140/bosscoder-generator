# Bosscoder Generator

Bosscoder Generator is a backend project written in Go that leverages the OpenAI API to generate linkedin post alerts.

## Prerequisites

- Go 1.16 or higher
- OpenAI API Key

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/kundu-ramit/bosscoder-generator.git
    cd bosscoder-generator
    ```

2. **Set your OpenAI API key:**
    ```sh
    apiKey := os.Getenv("OPENAI_API_KEY")
    ```

3. **Navigate to the backend directory:**
    ```sh
    cd backend
    ```

4. **Run the Go application:**
    ```sh
    go run bosscoder_generate.go
    ```

## Usage

Once the server is running, you can interact with the backend to generate sample linkedin post alerts using the OpenAI API.

## License

This project is licensed under the MIT License.