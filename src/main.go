package main

import (
	"./binding"
	"fmt"
)

func main() {
	lib := binding.Init()
	defer binding.Done(lib)
	fmt.Printf("lib: %v \n", *lib)

	binding.ASS_set_message_cb(lib)

	renderer := binding.ASS_renderer_init(lib)
	defer binding.ASS_renderer_done(renderer)
	fmt.Printf("renderer: %v \n", *renderer)

	track := binding.ASS_read_file(lib, "sample.ass")
	fmt.Printf("track: %v \n", *track)

	binding.ASS_set_frame_size(renderer, 1280, 720)

	frame := binding.ASS_render_frame(renderer, track, 1000)
	fmt.Printf("frame: %v \n", frame)

}
