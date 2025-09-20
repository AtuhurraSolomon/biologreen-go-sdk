Bio-Logreen Official Go SDK
The official Go SDK for the Bio-Logreen Facial Authentication API.

This SDK provides a simple and idiomatic way for Go developers to integrate passwordless, face-based authentication into their applications.

Features
Simple client initialization

Register new users with SignupFace

Authenticate existing users with LoginFace

Helper utility to convert image files to Base64

Clear error handling for API responses

Installation
To install the Bio-Logreen Go SDK, run the following command in your terminal:

go get [github.com/AtuhurraSolomon/biologreen-go-sdk](https://github.com/AtuhurraSolomon/biologreen-go-sdk)

Usage
Here is a basic example of how to use the client to sign up a new user.

1. Initialize the Client
First, import the package and create a new client using your secret API key obtained from the Bio-Logreen developer dashboard.

package main

import (
    "fmt"
    "log"

    "[github.com/AtuhurraSolomon/biologreen-go-sdk/biologreen](https://github.com/AtuhurraSolomon/biologreen-go-sdk/biologreen)"
)

func main() {
    // It is recommended to load your API key from an environment variable
    apiKey := "YOUR_SECRET_API_KEY"

    client, err := biologreen.NewClient(apiKey)
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    // ... use the client
}

2. Sign Up a New User
To register a new user, provide the path to their image and any optional custom data you want to associate with them.

    // Path to the user's image file
    imagePath := "path/to/user-image.jpg"

    // Use the built-in helper to convert the image to a Base64 string
    imageBase64, err := biologreen.ImageFileToBase64(imagePath)
    if err != nil {
        log.Fatalf("Error converting image: %v", err)
    }

    // Prepare the request
    signupReq := biologreen.SignupRequest{
        ImageBase64: imageBase64,
        CustomFields: map[string]interface{}{
            "plan":    "premium",
            "user_id": "your-internal-user-id-123",
        },
    }

    // Call the API
    signupResp, err := client.SignupFace(signupReq)
    if err != nil {
        log.Fatalf("API call failed: %v", err)
    }

    // Print the successful response
    fmt.Println("--- Signup Success! ---")
    fmt.Printf("Bio-Logreen User ID: %d\n", signupResp.UserID)
    fmt.Printf("Is New User: %t\n", signupResp.IsNewUser)
    fmt.Printf("Custom Fields: %v\n", signupResp.CustomFields)

3. Log In an Existing User
Logging in an existing user is just as simple.

    // Path to the image for login
    loginImagePath := "path/to/login-attempt.jpg"
    loginImageBase64, err := biologreen.ImageFileToBase64(loginImagePath)
    if err != nil {
        log.Fatalf("Error converting image: %v", err)
    }

    // Prepare and send the request
    loginReq := biologreen.LoginRequest{ImageBase64: loginImageBase64}
    loginResp, err := client.LoginFace(loginReq)
    if err != nil {
        log.Fatalf("Login API call failed: %v", err)
    }

    // On success, the response contains the matched user's data
    fmt.Println("--- Login Success! ---")
    fmt.Printf("Authenticated User ID: %d\n", loginResp.UserID)

Local Development
If you are running the Bio-Logreen API on your local machine (e.g., via Docker), you can point the client to your local server by providing the URL during initialization:

    localApiUrl := "http://localhost:8000/v1"
    client, err := biologreen.NewClient(apiKey, localApiUrl)

For issues or feature requests, please open an issue on GitHub.
