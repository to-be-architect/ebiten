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

import (
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	_GLFW_INSERT_FIRST = 0
	_GLFW_INSERT_LAST  = 1
)

var _glfw library

type initconfig struct {
	hasButtons bool
}

type wndconfig struct {
	width          int
	height         int
	title          string
	resizable      bool
	visible        bool
	decorated      bool
	focused        bool
	autoIconify    bool
	floating       bool
	maximized      bool
	centerCursor   bool
	focusOnShow    bool
	scaleToMonitor bool
}

type ctxconfig struct {
	client     int
	source     int
	major      int
	minor      int
	forward    bool
	debug      bool
	noerror    bool
	profile    int
	robustness int
	release    int
	share      *Window
}

type fbconfig struct {
	redBits        int
	greenBits      int
	blueBits       int
	alphaBits      int
	depthBits      int
	stencilBits    int
	accumRedBits   int
	accumGreenBits int
	accumBlueBits  int
	accumAlphaBits int
	auxBuffers     int
	stereo         bool
	samples        int
	sRGB           bool
	doublebuffer   bool
	transparent    bool
	handle         uintptr
}

type context struct {
	client     int
	source     int
	major      int
	minor      int
	revision   int
	forward    bool
	debug      bool
	noerror    bool
	profile    int
	robustness int
	release    int

	wgl struct {
		dc       _HDC
		handl    _HGLRC
		interval int
	}
}

// TODO: Give correct function types
type (
	PosCallback             uintptr
	SizeCallback            uintptr
	CloseCallback           uintptr
	RefreshCallback         uintptr
	FocusCallback           uintptr
	IconifyCallback         uintptr
	MaximizeCallback        uintptr
	FramebufferSizeCallback uintptr
	ScaleCallback           uintptr
	MouseButtonCallback     uintptr
	CursorPosCallback       uintptr
	CursorEnterCallback     uintptr
	ScrollCallback          uintptr
	KeyCallback             uintptr
	CharacterCallback       uintptr
	CharModsCallback        uintptr
	DropCallback            uintptr
	MonitorCallback         func(monitor *Monitor, event PeripheralEvent)
)

type Window struct {
	resizable    bool
	decorated    bool
	autoIconify  bool
	floating     bool
	focusOnShow  bool
	shouldClose  bool
	userPointer  unsafe.Pointer
	doublebuffer bool
	videoMode    VidMode
	monitor      *Monitor
	cursor       *Cursor

	minwidth  int
	minheight int
	maxwidth  int
	maxheight int
	numer     int
	denom     int

	stickyKeys         bool
	stickyMouseButtons bool
	lockKeyMods        bool
	cursorMode         int
	mouseButtons       [MouseButtonLast + 1]byte
	keys               [KeyLast + 1]byte
	// Virtual cursor position when cursor is disabled
	virtualCursorPosX float64
	virtualCursorPosY float64
	rawMouseMotion    bool

	context context

	callbacks struct {
		pos         PosCallback
		size        SizeCallback
		close       CloseCallback
		refresh     RefreshCallback
		focus       FocusCallback
		iconify     IconifyCallback
		maximize    MaximizeCallback
		fbsize      FramebufferSizeCallback
		scale       ScaleCallback
		mouseButton MouseButtonCallback
		cusorPos    CursorPosCallback
		cursorEnter CursorEnterCallback
		scroll      ScrollCallback
		key         KeyCallback
		character   CharacterCallback
		charmods    CharModsCallback
		drop        DropCallback
	}

	win32 struct {
		handle    windows.HWND
		bigIcon   _HICON
		smallIcon _HICON

		cursorTracked  bool
		frameAction    bool
		iconified      bool
		maximized      bool
		transparent    bool // Whether to enable framebuffer transparency on DWM
		scaleToMonitor bool

		// Cached size used to filter out duplicate events
		width  int
		height int

		// The last received cursor position, regardless of source
		lastCursorPosX int
		lastCursorPosY int

		// The last recevied high surrogate when decoding pairs of UTF-16 messages
		highSurrogate uint16
	}
}

type Monitor struct {
	name string

	widthMM  int
	heightMM int

	window *Window

	modes       []*VidMode
	currentMode *VidMode

	originalRamp GammaRamp
	currentRamp  GammaRamp

	win32 struct {
		handle _HMONITOR

		// This size matches the static size of DISPLAY_DEVICE.DeviceName
		adapterName string
		displayName string
		modesPruned bool
		modeChanged bool
	}
}

type Cursor struct {
	win32 struct {
		handle _HCURSOR
	}
}

type tls struct {
	win32 struct {
		allocated bool
		index     uint32
	}
}

type library struct {
	initialized bool

	hints struct {
		init        initconfig
		framebuffer fbconfig
		window      wndconfig
		context     ctxconfig
		refreshRate int
	}

	errors  []error
	cursors []*Cursor
	windows []*Window

	monitors []*Monitor

	errorSlot   tls
	contextSlot tls
	errorLock   sync.Mutex

	callbacks struct {
		monitor MonitorCallback
	}

	win32 struct {
		helperWindowHandle       windows.HWND
		deviceNotificationHandle _HDEVNOTIFY
		foregroundLockTimeout    uint32
		acquiredMonitorCount     int
		clipboardString          string
		keycodes                 [512]int16
		scancodes                [KeyLast + 1]int16
		keynames                 [KeyLast + 1]string

		// Where to place the cursor when re-enabled
		restoreCursorPosX float64
		restoreCursorPosY float64

		// The window whose disabled cursor mode is active
		disabledCursorWindow *Window
		rawInput             []_RAWINPUT
		mouseTrailSize       uint32
	}
}
