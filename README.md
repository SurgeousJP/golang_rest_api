# Golang REST API Example

Welcome to the repository for a simple example of a RESTful API built using the Go programming language. This project demonstrates how to create a basic REST API that performs CRUD (Create, Read, Update, Delete) operations on a resource (In this case, a MongoDB book collection)

## Getting Started

Follow the steps below to set up and run the Golang REST API on your local machine.

### Prerequisites

- Go (1.13 or later)
- Git

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/golang-rest-api.git
   ```
2. **Navigate to the project directory**
   ```
   cd golang-rest-api
   ```
3. **Build the project**
    ```
    ./golang-rest-api
    ```
By default, the API server will run on http://localhost:9090.

### API Endpoints
The API provides the following endpoints to manage a resource (e.g., "books"):

**Insert an item (book) (put the book in the body request)**
```
POST("/create", bc.CreateBook)
```
**Get an item (book) (given its name)**
```
GET("/get/:name", bc.GetBook)
```
**Get all books in DB**
```
GET("/getall", bc.GetAllBooks)
```
**Update an item (book) (put the book in the body request)**
```
PATCH("/update", bc.UpdateBook)
```
**Delete an item (book) (given its name)**
```
DELETE("/delete/:name", bc.DeleteBook)
```

