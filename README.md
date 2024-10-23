# LiteAPI Supplier Project

This project is an API built using the Gin framework in Golang that integrates with the Hotelbeds API. It processes requests to fetch the cheapest hotel rates and returns a customized response. The project includes unit testing and is set up for automated deployment using GitHub Actions to AWS ECS Fargate and EC2 Auto Scaling.

## Running the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/liteapi-supplier.git
   cd liteapi-supplier
   ```

2. Set up environment variables in a .env file:
    ```
    HOTELBEDS_BASE_URL=https://developer.hotelbeds.com/booking-api
    HOTELBEDS_API_KEY=your_api_key
    HOTELBEDS_SECRET=your_secret
    LITEAPI_SUPPLIER_CONFIG=your_liteapi_supplier_config
    PORT=8080
    ```

3. Install Dependencies
    ```bash
    go mod tidy
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

## Running Unit Tests

    ```bash
    go test ./... -v
    ```
