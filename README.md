# HNG12 Stage 0 - Public API

## Overview

This project is a simple public API that returns JSON-formatted information, including:

Your registered email address on the HNG12 Slack workspace.

The current datetime in ISO 8601 format (UTC).

The GitHub repository URL for this project.


## Technology Stack

Programming Language: Go

Framework: Standard Go net/http package

Deployment: Hosted on a publicly accessible endpoint

CORS Handling: Implemented using standard HTTP headers


## API Documentation

Endpoint

URL: https://hngstage0api.onrender.com/

Method: GET


Response Format (200 OK)

{
  "email": "your-email@example.com",
    "current_datetime": "2025-01-30T09:30:00Z",
      "github_url": "https://github.com/Ayobami0/HNGStage0"
      }

How to Run Locally

Prerequisites: Install Go

1. Clone this repository:

```
git clone https://github.com/Ayobami0/HNGStage0
cd your-repo
```


2. Run the server:

```go run main.go```


3. Open your browser or use Postman/cURL to test the endpoint:

```curl http://localhost:8080```
