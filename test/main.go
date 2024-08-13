package main

import (
	"fmt"
	"net/http"
	"strings"

	"moul.io/http2curl"
)

func main2() {
    mermaidCode := `
    graph TD
        A[Start] --> B{Is it working?}
        B -->|Yes| C[Continue]
        B -->|No| D[Fix it]
        D --> E[End]
    `

    // Replace this URL with the actual API endpoint if available
    url := "https://app.diagrams.net/"
    payload := strings.NewReader(fmt.Sprintf(`{"code":"%s"}`, mermaidCode))

    req, err := http.NewRequest("POST", url, payload)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")

	 // Convert the request to a cURL command
	 command, err := http2curl.GetCurlCommand(req)
	 if err != nil {
		 fmt.Println("Error converting to cURL:", err)
		 return
	 }
	   // Print the cURL command
	   fmt.Println(command)
	

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
	}



    defer resp.Body.Close()

    fmt.Println("Flowchart created successfully.")
}
