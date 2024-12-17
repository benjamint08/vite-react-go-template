#
<img src="go_react_vite_tailwind.png">

# Go, React, Tailwind and Vite template (mongohelper branch)

By [benjamint08](https://github.com/benjamint08)

## Description

This is a template for a fullstack application using Go, MongoDB, React and Vite. This template application is a simple todo list.

## Features

- Go backend
- MongoDB database
- React frontend
- Tailwind CSS
- Vite for frontend development

## Installation

1. Clone the repository
2. Run `npm install`
3. Run `go run main.go -dev` for development
4. Run `go run main.go -build` to build the application
5. Run `go run main.go` to run the application in production mode

## Adding API endpoints

To add an API endpoint, create a new file in the `handlers` directory. Check the `handlers/todo_handlers.go` file for an example.

You can then add the handler to the router in the `main.go` file.

## Production (Docker)

1. Run `docker build -t <your_image_name> .`
2. Run `docker run -d -p 8080:8080 -e MONGO_URI="<MONGO_URI" <your_image_name>`
3. Visit `http://localhost:8080` in your browser
4. Success