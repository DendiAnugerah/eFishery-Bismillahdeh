# Documentation

Simple API documentation about this project

Base URL: localhost:9000

### Category

To create a new category, use the following request\
POST -> /category/create

Example:

```
Parameter JSON
{
    name: "Ikan"
}
```

### Product

To create a new product use:\
POST -> /product/create

Example:

```
{
    "name": "Ikan hiu",
    "description": "Ikan hiu berbahaya",
    "price": 2000,
    "id_category": 1
}
```
