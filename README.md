# gmap-menu-scraper-go

This project uses Selenium WebDriver to crawl menu photos from Google Maps.

## Setup

1. Install Go v1.21.0 from the [official website](https://golang.org/dl/).
2. Install Selenium Server, Chrome, and Chromedriver by executing the `setup.sh` script:
    ```sh
    ./setup.sh
    ```

## Running the Scraper

1. Navigate to the project directory:
    ```sh
    cd gmap-menu-scraper-go
    ```
2. Run the Go script:
    ```sh
    go run main.go
    ```

## Building the Project

1. Ensure you are in the project directory:
    ```sh
    cd gmap-menu-scraper-go
    ```
2. Build the Go binary into the `bin` directory:
    - On Linux and macOS:
        ```sh
        go build -o bin/gmap-menu-scraper
        ```
3. After building, you can run the binary directly:
    - On Linux and macOS:
        ```sh
        ./bin/gmap-menu-scraper
        ```

The script will create a directory named `photos` and download the menu photos of the specified restaurant into it.