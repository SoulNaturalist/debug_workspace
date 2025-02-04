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
	// Kill all screens only linux and os x
	bt.KillAllScreen()
	// Get value by key redis db
	redis.GetValue("super_key")
	// Set value and key redis db
	redis.SetValue("k", "v")

}
