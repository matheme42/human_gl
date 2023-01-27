package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// makeVao initializes and returns a vertex array from the points provided.
func initLimbs() [6]uint32 {

	var vaoList [6]uint32

	vaoList[0] = createLimb(0.5, 0.5, 0.5)
	return vaoList
}

func createLimb(x float32, y float32, z float32) uint32 {

	var (
		triangle = []float32{
			0, 0, 0,
			0, y, 0,
			x, 0, 0,

			x, y, 0,
			0, y, 0,
			x, 0, 0,

			x, 0, z,
			0, 0, z,
			x, 0, 0,

			0, 0, 0,
			0, 0, z,
			x, 0, z,

			0, y, z,
			0, 0, z,
			0, y, 0,

			0, 0, 0,
			0, 0, z,
			0, y, z,
		}
	)

	mgl32.Perspective(80, 16/9, 0.4, 100)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(triangle), gl.Ptr(triangle), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	return vao
}
