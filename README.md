# Description

Cosmart software engineer assignment

## Prerequisite

- Go (Golang)  
  See [Golang Installation](https://golang.org/doc/install)

- MySQL  
  See [MySQL Installation](https://www.mysql.com/downloads/)

## Installation

1. Clone this repository and install the prerequisites above
2. Copy `.env` from `env.sample` and modify the configuration value appropriately
3. install dependencies with `make dep`
4. test with `make test`
4. build and run with `make run`

## Migration

1. Make sure env is configured correctly
2. migrate the schema with `make migrate`
3. seed the database with data from openlibrary with `make seed`

## API LIST

### **Get Books**

Get a list of books based on the query genre

#### **Details**

URL : `/books`

Method : `GET`

Query :
| Key | Example Value | Description |
| ----------- | ----------- | ----------- |
| offset | 0 | offset of the books
| limit | 10 | limit of the books
| genre | love | genre of the books

#### **Example Curl** :
```curl
curl --location --request GET 'localhost:3000/books?limit=10&offset=0&genre=love'
```
#### **Success Response**
```json
{
    "success": true,
    "message": "Success",
    "data": [
        {
            "ID": 19,
            "CreatedAt": "2023-01-29T02:44:16.053+07:00",
            "UpdatedAt": "2023-01-29T02:44:16.053+07:00",
            "DeletedAt": null,
            "Title": "Mémoires d'Hadrien",
            "Author": "Marguerite Yourcenar",
            "Edition": 124,
            "Genre": "Love"
        },
        {
            "ID": 30,
            "CreatedAt": "2023-01-29T02:44:16.057+07:00",
            "UpdatedAt": "2023-01-29T02:44:16.057+07:00",
            "DeletedAt": null,
            "Title": "A Wrinkle in Time",
            "Author": "Madeleine L'Engle",
            "Edition": 82,
            "Genre": "Love"
        }
    ]
}
```

### **Get a Single Book**

Get a single book by their ID

#### **Details**

URL : `/books/:id`

Method : `GET`

Params : 
| Key | Example Value | Description |
| ----------- | ----------- | ----------- |
| id | 1 | id of the book

#### **Example Curl** :
```curl
curl --location --request GET 'localhost:3000/books?limit=10&offset=0&genre=love'
```
#### **Success Response**
```json
{
    "success": true,
    "message": "Success",
    "data": {
        "ID": 1,
        "CreatedAt": "2023-01-29T02:44:16.045+07:00",
        "UpdatedAt": "2023-01-29T02:44:16.045+07:00",
        "DeletedAt": null,
        "Title": "Wuthering Heights",
        "Author": "Emily Brontë",
        "Edition": 1746,
        "Genre": "British and irish fiction (fictional works by one author)"
    }
}
```

### **Create a Single Book**

Create a single book

#### **Details**

URL : `/books`

Method : `POST`

Json Parameters : 
| Key | Example Value | Description |
| ----------- | ----------- | ----------- |
| title | "book title" | title of the book
| author | "lorem ipsum" | author of the book
| genre | "gaming" | genre of the book
| edition | 1 | book edition

#### **Example Curl** :
```curl
curl --location --request POST 'localhost:3000/books' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "how to please your dota carry",
    "author": "lorem ipsum",
    "genre": "gaming",
    "edition": 1
}'
```

#### **Success Response**
```json
{
    "success": true,
    "message": "Book created successfully",
    "data": {
        "ID": 1003,
        "CreatedAt": "2023-01-29T17:53:26.98+07:00",
        "UpdatedAt": "2023-01-29T17:53:26.98+07:00",
        "DeletedAt": null,
        "Title": "how to please your carry",
        "Author": "lorem ipsum",
        "Edition": 1,
        "Genre": "gaming"
    }
}
```

### **Create a Pickup Schedule**

Create a pickup schedule

#### **Details**

URL : `/schedule-borrow`

Method : `POST`

Json Parameters : 
| Key | Example Value | Description |
| ----------- | ----------- | ----------- |
| name | "bung messi" | name of the borrower
| book_id | 3 | id of the borrowed book
| pickup_time | "2023-01-31T22:00:00.123Z" | pickup time of the book

#### **Example Curl** :
```curl
curl --location --request POST 'localhost:3000/schedule-borrow' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Farras",
    "book_id": 3,
    "pickup_time": "2023-01-31T12:00:00.123Z"
}'
```

#### **Success Response**
```json
{
    "success": true,
    "message": "Borrow Schedule Created Successfully",
    "data": {
        "ID": 1,
        "CreatedAt": "2023-01-29T18:03:02.689+07:00",
        "UpdatedAt": "2023-01-29T18:03:02.689+07:00",
        "DeletedAt": null,
        "Name": "Farras",
        "BookId": 3,
        "PickupTime": "2023-01-31T12:00:00.123Z",
        "DueTime": "2023-02-07T12:00:00.123Z",
        "ReturnTime": null
    }
}
```

### **Return Book**

update the schedule for returning the book

#### **Details**

URL : `/schedule-return`

Method : `POST`

Json Parameters : 
| Key | Example Value | Description |
| ----------- | ----------- | ----------- |
| schedule_id | 1 | id of the schedule
| return_time | "2023-01-31T22:00:00.123Z" | return time of the book

#### **Example Curl** :
```curl
curl --location --request POST 'localhost:3000/schedule-return' \
--header 'Content-Type: application/json' \
--data-raw '{
    "schedule_id": 1,
    "return_time": "2023-02-02T22:00:00.123Z"
}'
```

#### **Success Response**
```json
{
    "success": true,
    "message": "Success",
    "data": {
        "ID": 1,
        "CreatedAt": "2023-01-29T18:03:02.689+07:00",
        "UpdatedAt": "2023-01-29T18:04:57.874+07:00",
        "DeletedAt": null,
        "Name": "Farras",
        "BookId": 3,
        "PickupTime": "2023-01-31T19:00:00.123+07:00",
        "DueTime": "2023-02-07T19:00:00.123+07:00",
        "ReturnTime": "2023-02-02T22:00:00.123Z"
    }
}
```