# cdmi-client-go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/grycap/cdmi-client-go)

A basic Go library to perform the core container and object operations defined in the Cloud Data Management Interface (CDMI) [specification](https://www.snia.org/sites/default/files/CDMI_Spec_v1.1.1.pdf).

## Installation

 ```bash
 go get github.com/grycap/cdmi-client-go
 ```

## Example

```go
package main

import (
    "io"
    "net/url"
    "os"

    "github.com/grycap/cdmi-client-go"
)

func main() {
    // CDMI Server endpoint
    endpoint, _ := url.Parse("https://my-cdmi-server.example")
    // Bearer auth token (if not required set an empty string)
    token := "my-token"
    // Verify SSL certificates
    verify := true

    // Create a new CDMI client
    client := cdmi.New(endpoint, token, verify)

    // Create a container (directory)
    err := client.CreateContainer("newcontainerName/anotherContainer", true)
    if err != nil {
        // Example: ignore error 400 (folder already exists)
        if err != cdmi.ErrBadRequest {
            // Manage error
        }
    }

    // Upload a file
    file, _ := os.Open("/path/to/file")
    defer file.Close()
    err = client.CreateObject("containerName/objectName", file, true)
    if err != nil {
        // Manage error
    }

    // Download a file
    newFile, _ := os.Create("/path/to/new/file")
    defer newFile.Close()
    content, err := client.GetObject("containerName/objectName")
    if err != nil {
        // Manage error
    }
    defer content.Close()
    io.Copy(newFile, content)
}
```

All available methods can be found at pkg.go.dev [reference](https://pkg.go.dev/github.com/grycap/cdmi-client-go).
