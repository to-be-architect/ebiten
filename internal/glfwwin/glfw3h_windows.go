// Copyright 2002-2006 Marcus Geelnard
// Copyright 2006-2019 Camilla Löwy
// Copyright 2022 The Ebiten Authors
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would
//    be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such, and must not
//    be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source
//    distribution.

package glfwwin

const (
	CursorNormal   = 0x00034001
	CursorHidden   = 0x00034002
	CursorDisabled = 0x00034003
)

const (
	KeyUnknown = -1

	// Printable keys
	KeySpace        = 32
	KeyApostrophe   = 39 // '
	KeyComma        = 44 // ,
	KeyMinus        = 45 // -
	KeyPeriod       = 46 // .
	KeySlash        = 47 // /
	Key0            = 48
	Key1            = 49
	Key2            = 50
	Key3            = 51
	Key4            = 52
	Key5            = 53
	Key6            = 54
	Key7            = 55
	Key8            = 56
	Key9            = 57
	KeySemicolon    = 59 // ;
	KeyEqual        = 61 // =
	KeyA            = 65
	KeyB            = 66
	KeyC            = 67
	KeyD            = 68
	KeyE            = 69
	KeyF            = 70
	KeyG            = 71
	KeyH            = 72
	KeyI            = 73
	KeyJ            = 74
	KeyK            = 75
	KeyL            = 76
	KeyM            = 77
	KeyN            = 78
	KeyO            = 79
	KeyP            = 80
	KeyQ            = 81
	KeyR            = 82
	KeyS            = 83
	KeyT            = 84
	KeyU            = 85
	KeyV            = 86
	KeyW            = 87
	KeyX            = 88
	KeyY            = 89
	KeyZ            = 90
	KeyLeftBracket  = 91  // [
	KeyBackslash    = 92  // \
	KeyRightBracket = 93  // ]
	KeyGraveAccent  = 96  // `
	KeyWorld1       = 161 // non-US #1
	KeyWorld2       = 162 // non-US #2

	// Function keys
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyF13          = 302
	KeyF14          = 303
	KeyF15          = 304
	KeyF16          = 305
	KeyF17          = 306
	KeyF18          = 307
	KeyF19          = 308
	KeyF20          = 309
	KeyF21          = 310
	KeyF22          = 311
	KeyF23          = 312
	KeyF24          = 313
	KeyF25          = 314
	KeyKP0          = 320
	KeyKP1          = 321
	KeyKP2          = 322
	KeyKP3          = 323
	KeyKP4          = 324
	KeyKP5          = 325
	KeyKP6          = 326
	KeyKP7          = 327
	KeyKP8          = 328
	KeyKP9          = 329
	KeyKPDecimal    = 330
	KeyKPDivide     = 331
	KeyKPMultiply   = 332
	KeyKPSubtract   = 333
	KeyKPAdd        = 334
	KeyKPEnter      = 335
	KeyKPEqual      = 336
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyMenu         = 348

	KeyLast = KeyMenu
)

const (
	MouseButton1      = 0
	MouseButton2      = 1
	MouseButton3      = 2
	MouseButton4      = 3
	MouseButton5      = 4
	MouseButton6      = 5
	MouseButton7      = 6
	MouseButton8      = 7
	MouseButtonLast   = MouseButton8
	MouseButtonLeft   = MouseButton1
	MouseButtonRight  = MouseButton2
	MouseButtonMiddle = MouseButton3
)

var (
	NotInitialized     Error = 0x00010001
	NoCurrentContext   Error = 0x00010002
	InvalidEnum        Error = 0x00010003
	InvalidValue       Error = 0x00010004
	OutOfMemory        Error = 0x00010005
	ApiUnavailable     Error = 0x00010006
	VersionUnavailable Error = 0x00010007
	PlatformError      Error = 0x00010008
	FormatUnavailable  Error = 0x00010009
	NoWindowContext    Error = 0x0001000A
)

type PeripheralEvent int

const (
	Connected    PeripheralEvent = 0x00040001
	Disconnected PeripheralEvent = 0x00040002
)

const (
	DONT_CARE = -1
)

type Error int

func (e Error) Error() string {
	switch e {
	case NotInitialized:
		return "the GLFW library is not initialized"
	case NoCurrentContext:
		return "there is no current context"
	case InvalidEnum:
		return "invalid argument for enum parameter"
	case InvalidValue:
		return "invalid value for parameter"
	case OutOfMemory:
		return "out of memory"
	case ApiUnavailable:
		return "the requested API is unavailable"
	case VersionUnavailable:
		return "the requested API version is unavailable"
	case PlatformError:
		return "a platform-specific error occurred"
	case FormatUnavailable:
		return "the requested format is unavailable"
	case NoWindowContext:
		return "the specified window has no context"
	default:
		return "ERROR: UNKNOWN GLFW ERROR"
	}
}

type VidMode struct {
	Width       int
	Height      int
	RedBits     int
	GreenBits   int
	BlueBits    int
	RefreshRate int
}

type GammaRamp struct {
	Red   []int
	Green []int
	Blue  []int
}

type Image struct {
	Width  int
	Height int
	Pixels []byte
}
