package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func deleteLines(x int) {
	clearLine := "\033[2K"
	moveCursorUp := "\033[A"
	// Clear the last x lines
	for i := 0; i < x; i++ {
		fmt.Print(moveCursorUp) // Move cursor up
		fmt.Print(clearLine)    // Clear the current line
	}

	// Move cursor up for the remaining lines
	//for i := 0; i < x; i++ {
	//	fmt.Print(moveCursorUp) // Move cursor up
	//}
}

const (
	SPEED_MODIFIER = 10
	TARGET_FPS     = 120
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Moving Circle Example")
	defer rl.CloseWindow()

	xPos := float32(screenWidth) / 2
	yPos := float32(screenHeight) / 2
	radius := float32(50)
	speed := float32(20)

	rl.SetTargetFPS(TARGET_FPS)

	// collect info only once every X frames
	cycle := 0
	step := (float32(TARGET_FPS) / 10)

	var lastFrameDrawTime time.Duration
	for !rl.WindowShouldClose() {
		startRenderTime := time.Now()
		speedModifier := rl.GetFrameTime() * SPEED_MODIFIER
		// collect info
		cycle++
		if float32(cycle) > step {
			cycle = 0
			deleteLines(4)
			fmt.Printf("\n")
			fmt.Println("FPS: ", rl.GetFPS())
			fmt.Println("Time since last frame: ", rl.GetFrameTime()*1000, "ms")
			fmt.Println("Time to process last frame: ", float32(lastFrameDrawTime.Microseconds())/1000, "ms")
		}

		// Update position
		xPos += speed * speedModifier
		if xPos > float32(screenWidth)-radius || xPos < radius {
			speed = -speed
		}

		// Drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawCircle(int32(xPos), int32(yPos), radius, rl.Maroon)
		rl.DrawFPS(20, 20)
		lastFrameDrawTime = time.Since(startRenderTime)
		rl.EndDrawing()
	}
}
