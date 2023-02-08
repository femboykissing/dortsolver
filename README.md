# Dort Solver
This project uses the API from Dort, which you can purchase at https://discord.gg/arkoselabs.

I made this in about 5 minutes while stressed.

### Example for Outlook
```go
package main

import (
	"log"
	"github.com/femboykissing/dortsolver"
)

func main() {
	solver := dortsolver.Solver{
		ApiKey:    "your_key",
		PublicKey: "B7D8911C-5CC8-A9A3-35B0-554ACEE604DA",
		SiteUrl:   "https://iframe.arkoselabs.com",
	}

	response, err := solver.Solve()
	if err != nil {
		log.Fatalf("Failed to solve: %s\n", err.Error())
		return
	}

	log.Printf("%s\n", response.GameToken)
}
```