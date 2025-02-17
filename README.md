# Go Web Toolkit

## Overview

Go Web Toolkit is a reusable Go module designed to simplify common tasks in web application development. Instead of rewriting the same code for every project, you can use this toolkit to handle JSON operations, file uploads, error handling, and more with ease.

## Features

- Read and write JSON
- Produce well-formatted error messages
- Upload files with restrictions on MIME type and size
- Create directories on the server
- Generate random strings
- Download static files
- Post JSON to a remote service
- Create URL-safe slugs from strings
- Fully tested and built using Go's standard library

## Installation

To install Go Web Toolkit, use:

```sh
go get github.com/yourusername/go-web-toolkit
```

## Usage

Import the package in your Go project:

```go
import "github.com/yourusername/go-web-toolkit"
```

Example usage:

```go
package main

import (
    "fmt"
    "github.com/yourusername/go-web-toolkit"
)

func main() {
    slug := toolkit.CreateSlug("Hello World!")
    fmt.Println(slug) // Output: hello-world
}
```

## Get the latest tagged version

To get the latest tagged version of the module, run:

```sh
go get -u github.com/qaiswardag/goweb-toolkit@v0.0.0
```

## License

This project is licensed under the MIT License.
