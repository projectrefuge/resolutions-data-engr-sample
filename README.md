# resolutions-data-engr-sample

This repository contains a sample of Go code that uses Regex matching to break apart text documents into small enough chunks to create vector embeddings.

## Prerequisites

- Go 1.16 or higher. You can download it from [here](https://golang.org/dl/).

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/projectrefuge/resolutions-data-engr-sample.git
    cd resolutions-data-engr-sample
    ```

2. Install dependencies (if any):
    ```bash
    go mod tidy
    ```

## Usage

To run the code, execute the following command in the root directory of the repository:
```bash
go run main.go
```

## Example Output

After running the code, you should see text files created in the `resolutions` directory, each containing chunks of the input text document.

## Next Steps

MongoDB Atlas provides a way to create embeddings from text data using the Atlas Vector Search feature. You can learn more about it [here](https://www.mongodb.com/docs/atlas/atlas-vector-search/create-embeddings/). You can choose from Go, Python, Java, or Node.js to create embeddings.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.