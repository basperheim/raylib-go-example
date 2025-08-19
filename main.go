package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func main() {
	rl.InitWindow(800, 450, "Raylib-Go Orbital Camera")
	defer rl.CloseWindow()

	camera := rl.Camera3D{
		Position:   rl.NewVector3(10.0, 10.0, 10.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Make camera orbit around origin
		angle := float64(rl.GetTime()) * 0.5 // radians/sec
		radius := 10.0

		camera.Position.X = float32(radius * math.Cos(angle))
		camera.Position.Z = float32(radius * math.Sin(angle))
		camera.Position.Y = 5.0

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)
		rl.DrawCube(rl.NewVector3(0, 1, 0), 2, 2, 2, rl.Red)
		rl.DrawCubeWires(rl.NewVector3(0, 1, 0), 2, 2, 2, rl.Maroon)
		rl.DrawGrid(10, 1.0)
		rl.EndMode3D()

		rl.DrawText("Orbital camera around cube", 10, 40, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}
