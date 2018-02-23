package binding

/*

#include <stdio.h>
#include <stdarg.h>

// The gateway function
int debugCallback_cgo(int level, const char *fmt, va_list args, void *data)
{
	void debugCallback(int, const char *, va_list, void *);
	debugCallback(level, fmt, args, data);
}
*/
import "C"
