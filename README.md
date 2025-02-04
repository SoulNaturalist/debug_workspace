# Debugging Environment for Rapid Development

This code snippet represents a debugging environment that has been created for personal use. The purpose of this environment is to streamline the development process by providing quick access to various functionalities without the need for extensive implementation each time. This allows for faster testing and iteration, ultimately saving time in the development cycle.

## Example Code

The following code demonstrates how different components can be integrated into a single workflow, showcasing the capabilities of the debugging environment:

```go
package main

import (
	"debug_workspace/ai"
	"debug_workspace/bt"
	"debug_workspace/ngrok"
	"debug_workspace/redis"
)

func main() {
	// Start network tunnel using ngrok
	ngrok.Tunnel("NGROK_TOKEN", "3000")
	
	// Generate tests by prompting the AI system with the given prompt and AI ID
	ai.GenereateTestsByPrompt("hi", 2)
	
	// Open new screen for background tasks
	bt.CreateNewScreen()
	
	// Kill all screens only on Linux and OS X
	bt.KillAllScreen()
	
	// Get value by key from Redis database
	redis.GetValue("super_key")
	
	// Set value and key in Redis database
	redis.SetValue("k", "v")
}
```