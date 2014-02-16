package gocoa

/*
#cgo CFLAGS: -I/System/Library/Frameworks/CoreFoundation.framework/Versions/A/Headers/
#cgo LDFLAGS: -lobjc

#include <CoreFoundation.h>
*/
import "C"

import (
	"unsafe"
	"reflect"
)

func NSDictionary(key string, value Object) Object {
	return ClassForName("NSDictionary").Instance("alloc").Call("initWithObject:forKey:", value, NSString(key))
}

func NSString(inString string) Object {
	l := C.CFIndex(len(inString))
	ret := C.CFStringCreateWithBytes(nil, *(**C.UInt8)(unsafe.Pointer(&inString)),
		l, C.kCFStringEncodingUTF8, 0)
	return Object(unsafe.Pointer(ret))
}

func NSStringToString(inString Object) string {
	cr := C.CFStringRef(unsafe.Pointer(inString))

	var usedBufLen C.CFIndex
	rng := C.CFRange { C.CFIndex(0), C.CFStringGetLength(cr) }
	n := int(C.CFStringGetBytes(cr, rng, C.kCFStringEncodingUTF8, 0, 0, nil, 0, &usedBufLen))
	if n <= 0 { return "" }

	buf := make([]byte, int(usedBufLen))
	C.CFStringGetBytes(cr, rng, C.kCFStringEncodingUTF8, 0, 0, (*C.UInt8)(unsafe.Pointer(&buf[0])), C.CFIndex(len(buf)), &usedBufLen)

	sh := &reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&buf[0])),
		Len: int(usedBufLen),
	}
	return *(*string)(unsafe.Pointer(sh))
}

const (
	BlackColor     = "blackColor"
	BlueColor      = "blueColor"
	BrownColor     = "brownColor"
	ClearColor     = "clearColor"
	CyanColor      = "cyanColor"
	DarkGrayColor  = "darkGrayColor"
	GrayColor      = "grayColor"
	GreenColor     = "greenColor"
	LightGrayColor = "lightGrayColor"
	MagentaColor   = "magentaColor"
	OrangeColor    = "orangeColor"
	PurpleColor    = "purpleColor"
	RedColor       = "redColor"
	WhiteColor     = "whiteColor"
	YellowColor    = "yellowColor"
)

func NSColor(color string) Object {
	return ClassForName("NSColor").Instance(color)
}


const (
	NSBorderlessWindowMask		NSUInteger = 0
    NSTitledWindowMask			NSUInteger = 1 << 0
    NSClosableWindowMask		NSUInteger = 1 << 1
    NSMiniaturizableWindowMask	NSUInteger = 1 << 2
    NSResizableWindowMask		NSUInteger = 1 << 3
)

const (
    NSBackingStoreRetained	 	NSUInteger = 0
    NSBackingStoreNonretained	NSUInteger = 1
    NSBackingStoreBuffered		NSUInteger = 2
)
