# Origin Energy Go SDK

This is an unofficial Go SDK for the Origin Energy API.

## Features

- [x] Authenticate via OAuth2 authentication link
- [x] Get access token
- [x] Cache access token/refresh tokens
- [x] Refresh access token
- [x] Get user details
- [x] Get usage data
- [x] Get usage data for a specific date range
- [x] Get usage data for a specific date range and interval
- [x] Get hourly, daily, weekly, monthly and yearly usage data

## Installation

- Duplicate .example.env file and rename it to .env
- Fill in the required fields in the .env file
- Run the following command to run

```bash
go run cmd/main.go
```