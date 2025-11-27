# Go-MySQL RESTful API

A RESTful CRUD API for category management using Go and MySQL


## API Documentation

Read the OpenAPI Documentation [here](apispec.json)

## Environment Variables

- DB_USERNAME
- DB_PASSWORD
- DB_HOST
- DB_PORT
- DB_NAME
- SERVER_PORT
- API_KEY

## Initial SQL Query

Do the [initial SQL query](initial_queries.sql) to create the category table.

## API Request Authentication

The value of `X-API-Key` in API requests must match the `API_KEY` from the environment variables to pass the authentication.
