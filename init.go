package main

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func configGlfw() *glfw.Window {
	err := glfw.Init()
	if err != nil {
		log.Fatalln(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	window.MakeContextCurrent()
	return window
}

func configGl() {
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}
