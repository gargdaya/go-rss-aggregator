# RSS Feed Aggregator

This project is an RSS feed aggregator that collects and displays the latest news and updates from various RSS feeds.

## Table of Contents

-   [Installation](#installation)
-   [Usage](#usage)

## Installation

1. Install the required dependencies:

    ```bash
    go get ./...
    ```

    ```bash
    go mod vendor
    ```

2. Configure env File:
   Set the following environment variables in your `.env` file:
    - `PORT`: Specify the port for the server.
    - `DB_URL`: Specify the database URL.

## Usage

1. Start the aggregator:

    ```bash
    go run main
    ```

    or

    ```bash
    go build && ./rssagg
    ```

    and check for api endpoints in `main.go`
