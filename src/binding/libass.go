package binding

/*
#cgo CFLAGS: -I/usr/include/ass/ass
#cgo LDFLAGS: -L/usr/lib -lass
#include <ass/ass.h>
#include <ass/ass_types.h>


void debugCallback_cgo(int level, const char *fmt, va_list args, void *data);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Init() *C.ASS_Library {
	foo := C.ass_library_init()

	fmt.Printf("%v \n", *foo)

	return foo
}

func ASS_renderer_init(lib *C.ASS_Library) *C.ASS_Renderer {
	return C.ass_renderer_init(lib)
}

/**
 * \brief Read subtitles from file.
 * \param library library handle
 * \param fname file name
 * \param codepage encoding (iconv format)
 * \return newly allocated track
 */
//ASS_Track *ass_read_file(ASS_Library *library, char *fname, char *codepage);
func ASS_read_file(lib *C.ASS_Library, patch string) *C.ASS_Track {
	c_path := C.CString(patch)
	return C.ass_read_file(lib, c_path, C.CString("UTF-8"))
}

func ASS_renderer_done(renderer *C.ASS_Renderer) {
	C.ass_renderer_done(renderer)
}

/**
 * \brief Set the frame size in pixels, including margins.
 * The renderer will never return images that are outside of the frame area.
 * The value set with this function can influence the pixel aspect ratio used
 * for rendering. If the frame size doesn't equal to the video size, you may
 * have to use ass_set_pixel_aspect().
 * @see ass_set_pixel_aspect()
 * @see ass_set_margins()
 * \param priv renderer handle
 * \param w width
 * \param h height
 */
// void ass_set_frame_size(ASS_Renderer *priv, int w, int h);

func ASS_set_frame_size(renderer *C.ASS_Renderer, width int, height int) {
	C.ass_set_frame_size(renderer, C.int(width), C.int(height))
}

func Done(lib *C.ASS_Library) {
	C.ass_library_done(lib)
}

/**
 * \brief Render a frame, producing a list of ASS_Image.
 * \param priv renderer handle
 * \param track subtitle track
 * \param now video timestamp in milliseconds
 * \param detect_change compare to the previous call and set to 1
 * if positions changed, or set to 2 if content changed.
 */
// ASS_Image *ass_render_frame(ASS_Renderer *priv, ASS_Track *track, long long now, int *detect_change);
func ASS_render_frame(renderer *C.ASS_Renderer, track *C.ASS_Track, now uint) *C.ASS_Image {
	marker := C.int(0)
	return C.ass_render_frame(renderer, track, C.longlong(now), &marker)
}

/**
 * \brief Register a callback for debug/info messages.
 * If a callback is registered, it is called for every message emitted by
 * libass.  The callback receives a format string and a list of arguments,
 * to be used for the printf family of functions. Additionally, a log level
 * from 0 (FATAL errors) to 7 (verbose DEBUG) is passed.  Usually, level 5
 * should be used by applications.
 * If no callback is set, all messages level < 5 are printed to stderr,
 * prefixed with [ass].
 *
 * \param priv library handle
 * \param msg_cb pointer to callback function
 * \param data additional data, will be passed to callback
 */
// void ass_set_message_cb(ASS_Library *priv, void (*msg_cb) (int level, const char *fmt, va_list args, void *data), void *data);

func ASS_set_message_cb(lib *C.ASS_Library) {
	//callback := func(lvl C.int, data *C.char, add unsafe.Pointer) {
	//fmt.Printf("lvl[%v] - %v \n", lvl, C.GoString(data))
	//
	//}

	C.ass_set_message_cb(lib, (*[0]byte)(C.debugCallback_cgo), unsafe.Pointer(nil))
}

//export debugCallback
func debugCallback(lvl C.int, data *C.char, add unsafe.Pointer) {
	fmt.Printf("lvl[%v] - %v \n", lvl, C.GoString(data))

}
