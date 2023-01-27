package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func draw(vao uint32, window *glfw.Window, program uint32, count int32) {
	gl.UseProgram(program)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, count)
}

func main() {
	window := configGlfw()
	defer glfw.Terminate()
	configGl()
	var program = CreateProgramShader()
	var vao = initLimbs()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		draw(vao[0], window, program, 36)
		glfw.PollEvents()
		window.SwapBuffers()
	}
}
