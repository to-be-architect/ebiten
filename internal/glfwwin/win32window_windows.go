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
	"math"
	"reflect"
	"unsafe"
)

func (w *Window) getWindowStyle() uint32 {
	var style uint32 = _WS_CLIPSIBLINGS | _WS_CLIPCHILDREN

	if w.monitor != nil {
		style |= _WS_POPUP
	} else {
		style |= _WS_SYSMENU | _WS_MINIMIZEBOX
		if w.decorated {
			style |= _WS_CAPTION
			if w.resizable {
				style |= _WS_MAXIMIZEBOX | _WS_THICKFRAME
			}
		} else {
			style |= _WS_POPUP
		}
	}

	return style
}

func (w *Window) getWindowExStyle() uint32 {
	var style uint32 = _WS_EX_APPWINDOW

	if w.floating {
		style |= _WS_EX_TOPMOST
	}

	return style
}

func chooseImage(images []*Image, width, height int) *Image {
	var leastDiff uint = math.MaxUint32
	var closest *Image
	for _, image := range images {
		currDiff := abs(image.Width*image.Height - width*height)
		if currDiff < leastDiff {
			closest = image
			leastDiff = currDiff
		}
	}
	return closest
}

func createIcon(image *Image, xhot, yhot int, icon bool) (_HICON, error) {
	var bi _BITMAPV5HEADER
	bi.bV5Size = uint32(unsafe.Sizeof(bi))
	bi.bV5Width = int32(image.Width)
	bi.bV5Height = int32(-image.Height)
	bi.bV5Planes = 1
	bi.bV5BitCount = 32
	bi.bV5Compression = _BI_BITFIELDS
	bi.bV5RedMask = 0x00ff0000
	bi.bV5GreenMask = 0x0000ff00
	bi.bV5BlueMask = 0x000000ff
	bi.bV5AlphaMask = 0xff000000

	dc, err := _GetDC(0)
	if err != nil {
		return 0, err
	}
	defer _ReleaseDC(0, dc)

	color, targetPtr, err := _CreateDIBSection(dc, &bi, _DIB_RGB_COLORS, 0, 0)
	if err != nil {
		return 0, err
	}
	defer _DeleteObject(_HGDIOBJ(color))

	mask, err := _CreateBitmap(int32(image.Width), int32(image.Height), 1, 1, nil)
	if err != nil {
		return 0, err
	}
	defer _DeleteObject(_HGDIOBJ(mask))

	source := image.Pixels
	var target []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&target))
	h.Data = uintptr(unsafe.Pointer(targetPtr))
	h.Len = image.Width * image.Height
	h.Cap = image.Width * image.Height
	for i := 0; i < len(source)/4; i++ {
		target[4*i] = source[4*i+2]
		target[4*i+1] = source[4*i+1]
		target[4*i+2] = source[4*i+0]
		target[4*i+3] = source[4*i+3]
	}

	var iconInt32 int32
	if icon {
		iconInt32 = 1
	}
	ii := _ICONINFO{
		fIcon:    iconInt32,
		xHotspot: uint32(xhot),
		yHotspot: uint32(yhot),
		hbmMask:  mask,
		hbmColor: color,
	}
	handle, err := _CreateIconIndirect(&ii)
	if err != nil {
		return 0, err
	}

	return handle, nil
}

func getFullWindowSize(style uint32, exStyle uint32, contentWidth, contentHeight int, dpi uint32) (fullWidth, fullHeight int, err error) {
	rect := _RECT{
		left:   0,
		top:    0,
		right:  int32(contentWidth),
		bottom: int32(contentHeight),
	}
	if _AdjustWindowRectExForDpi_Available() {
		if err := _AdjustWindowRectExForDpi(&rect, style, false, exStyle, dpi); err != nil {
			return 0, 0, err
		}
	} else {
		if err := _AdjustWindowRectEx(&rect, style, false, exStyle); err != nil {
			return 0, 0, err
		}
	}
	return int(rect.right - rect.left), int(rect.bottom - rect.top), nil
}

func (w *Window) applyAspectRatio(edge int, area *_RECT) error {
	ratio := float32(w.numer) / float32(w.denom)

	var dpi uint32 = _USER_DEFAULT_SCREEN_DPI
	if _GetDpiForWindow_Available() {
		dpi = _GetDpiForWindow(w.win32.handle)
	}

	xoff, yoff, err := getFullWindowSize(w.getWindowStyle(), w.getWindowExStyle(), 0, 0, dpi)
	if err != nil {
		return err
	}

	if edge == _WMSZ_LEFT || edge == _WMSZ_BOTTOMLEFT || edge == _WMSZ_RIGHT || edge == _WMSZ_BOTTOMRIGHT {
		area.bottom = area.top + int32(yoff) + int32(float32(area.right-area.left-int32(xoff))/ratio)
	} else if edge == _WMSZ_TOPLEFT || edge == _WMSZ_TOPRIGHT {
		area.top = area.bottom - int32(yoff) - int32(float32(area.right-area.left-int32(xoff))/ratio)
	} else if edge == _WMSZ_TOP || edge == _WMSZ_BOTTOM {
		area.right = area.left + int32(xoff) + int32(float32(area.bottom-area.top-int32(yoff))*ratio)
	}

	return nil
}

func (w *Window) updateCursorImage() error {
	if w.cursorMode == CursorNormal {
		if w.cursor != nil {
			_SetCursor(w.cursor.win32.handle)
		} else {
			cursor, err := _LoadCursorW(0, _IDC_ARROW)
			if err != nil {
				return err
			}
			_SetCursor(cursor)
		}
	} else {
		_SetCursor(0)
	}
	return nil
}

// TODO: add more functions

func registerWindowClassWin32() error {
	// TODO: Implement this
	return nil
}

func (w *Window) platformSetWindowPos(xpos, ypos int) {
	panic("platformSetWindowPos is not implemented")
}

func (w *Window) platformGetWindowSize() (width, height int) {
	panic("platformGetWindowSize is not implemented")
}

func (w *Window) platformGetWindowFrameSize() (left, top, right, bottom int) {
	panic("platformGetWindowFrameSize is not implemented")
}

func (w *Window) platformSetWindowMonitor(monitor *Monitor, xpos, ypos, width, height, refreshRate int) {
	panic("platformSetWindowMonitor is not implemented")
}
