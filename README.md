#  Facebook Conversion API Integration with Golang

This repository provides a Golang implementation for integrating Facebook Conversion API. The Facebook Conversion API allows you to send web events from your server directly to Facebook's servers for tracking and measurement purposes.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Facebook Pixel**: You should already have a Facebook Pixel set up for your website.
- **Facebook Business Manager Account**: You need access to a Facebook Business Manager account for managing your events.
- **Access Token**: Obtain a Facebook Access Token with the necessary permissions for sending events.

## Installation

Use your preferred Go package manager to install the library:

```shell
go get github.com/yourusername/facebook-conversion-api
```

## Usage
To use this library, follow these steps:

1. Import the package into your Go application:

```go
import "github.com/yourusername/facebook-conversion-api"
```

Initialize the Facebook Conversion API with your Access Token:
```
c := facebook.NewClient("YOUR_ACCESS_TOKEN")

Create and send events to Facebook:
event := &facebook.Event{
    PixelID: "YOUR_PIXEL_ID",
    Data: facebook.EventData{
        EventName: "Purchase",
        EventTime: time.Now().Unix(),
        UserData: facebook.UserData{
            Email: "user@example.com",
        },
    },
}

err := c.SendEvent(event)
if err != nil {
    log.Fatal(err)
}
```

## Example
Check out the main.go file in this repository for a complete working example of how to send events to Facebook Conversion API using Golang.
![78](https://github.com/TABREZ-96/GoLang_Challenge/assets/114156392/5be94279-3563-44d4-b2d4-b656958c3dd8)
![WhatsApp Image 2023-10-05 at 00 15 06](https://github.com/TABREZ-96/GoLang_Challenge/assets/114156392/e1d33a1b-b2e7-4f0a-b0c3-09265354c6b7)
![66](https://github.com/TABREZ-96/GoLang_Challenge/assets/114156392/be5a772e-8c21-40e8-94dc-fea607473a60)


## Acknowledgments
Facebook Conversion API Documentation

