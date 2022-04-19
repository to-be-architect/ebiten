// Copyright 2022 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package glfwwin

import (
	"fmt"
	"runtime"
	"unsafe"

	"golang.org/x/sys/windows"
)

// For the definitions, see https://github.com/wine-mirror/wine
const (
	_BI_BITFIELDS                         = 3
	_CCHDEVICENAME                        = 32
	_CCHFORMNAME                          = 32
	_CDS_TEST                             = 0x00000002
	_CDS_FULLSCREEN                       = 0x00000004
	_DBT_DEVTYP_DEVICEINTERFACE           = 0x00000005
	_DEVICE_NOTIFY_WINDOW_HANDLE          = 0x00000000
	_DIB_RGB_COLORS                       = 0
	_DISP_CHANGE_SUCCESSFUL               = 0
	_DISP_CHANGE_RESTART                  = 1
	_DISP_CHANGE_FAILED                   = -1
	_DISP_CHANGE_BADMODE                  = -2
	_DISP_CHANGE_NOTUPDATED               = -3
	_DISP_CHANGE_BADFLAGS                 = -4
	_DISP_CHANGE_BADPARAM                 = -5
	_DISP_CHANGE_BADDUALVIEW              = -6
	_DISPLAY_DEVICE_ACTIVE                = 0x00000001
	_DISPLAY_DEVICE_MODESPRUNED           = 0x08000000
	_DISPLAY_DEVICE_PRIMARY_DEVICE        = 0x00000004
	_DM_BITSPERPEL                        = 0x00040000
	_DM_PELSWIDTH                         = 0x00080000
	_DM_PELSHEIGHT                        = 0x00100000
	_DM_DISPLAYFREQUENCY                  = 0x00400000
	_EDS_ROTATEDMODE                      = 0x00000004
	_ENUM_CURRENT_SETTINGS         uint32 = 0xffffffff
	_HORZSIZE                             = 4
	_IDC_ARROW                            = 32512
	_LOGPIXELSX                           = 88
	_LOGPIXELSY                           = 90
	_MAPVK_VSC_TO_VK                      = 1
	_PM_REMOVE                            = 0x0001
	_SPI_GETFOREGROUNDLOCKTIMEOUT         = 0x2000
	_SPI_SETFOREGROUNDLOCKTIMEOUT         = 0x2001
	_SPIF_SENDCHANGE                      = _SPIF_SENDWININICHANGE
	_SPIF_SENDWININICHANGE                = 2
	_SW_HIDE                              = 0
	_TLS_OUT_OF_INDEXES            uint32 = 0xffffffff
	_USER_DEFAULT_SCREEN_DPI              = 96
	_VERTSIZE                             = 6
	_VK_ADD                               = 0x6B
	_VK_DECIMAL                           = 0x6E
	_VK_DIVIDE                            = 0x6F
	_VK_MULTIPLY                          = 0x6A
	_VK_NUMPAD0                           = 0x60
	_VK_NUMPAD1                           = 0x61
	_VK_NUMPAD2                           = 0x62
	_VK_NUMPAD3                           = 0x63
	_VK_NUMPAD4                           = 0x64
	_VK_NUMPAD5                           = 0x65
	_VK_NUMPAD6                           = 0x66
	_VK_NUMPAD7                           = 0x67
	_VK_NUMPAD8                           = 0x68
	_VK_NUMPAD9                           = 0x69
	_VK_SUBTRACT                          = 0x6D
	_WMSZ_BOTTOM                          = 6
	_WMSZ_BOTTOMLEFT                      = 7
	_WMSZ_BOTTOMRIGHT                     = 8
	_WMSZ_LEFT                            = 1
	_WMSZ_RIGHT                           = 2
	_WMSZ_TOP                             = 3
	_WMSZ_TOPLEFT                         = 4
	_WMSZ_TOPRIGHT                        = 5
	_WS_BORDER                            = 0x00800000
	_WS_CAPTION                           = _WS_BORDER | _WS_DLGFRAME
	_WS_CLIPSIBLINGS                      = 0x04000000
	_WS_CLIPCHILDREN                      = 0x02000000
	_WS_DLGFRAME                          = 0x00400000
	_WS_EX_APPWINDOW                      = 0x00040000
	_WS_EX_CLIENTEDGE                     = 0x00000200
	_WS_EX_OVERLAPPEDWINDOW               = _WS_EX_WINDOWEDGE | _WS_EX_CLIENTEDGE
	_WS_EX_TOPMOST                        = 0x00000008
	_WS_EX_WINDOWEDGE                     = 0x00000100
	_WS_MAXIMIZEBOX                       = 0x00010000
	_WS_MINIMIZEBOX                       = 0x00020000
	_WS_POPUP                             = 0x80000000
	_WS_SYSMENU                           = 0x00080000
	_WS_THICKFRAME                        = 0x00040000
)

type (
	_BOOL       int32
	_HBITMAP    windows.Handle
	_HCURSOR    windows.Handle
	_HDC        windows.Handle
	_HDEVNOTIFY windows.Handle
	_HGDIOBJ    windows.Handle
	_HGLRC      windows.Handle
	_HICON      windows.Handle
	_HINSTANCE  windows.Handle
	_HMENU      windows.Handle
	_HMODULE    windows.Handle
	_HMONITOR   windows.Handle
	_LPARAM     uintptr
	_LRESULT    uintptr
	_WPARAM     uintptr
)

type _DPI_AWARENESS_CONTEXT windows.Handle

const (
	intSize = 32 << (^uint(0) >> 63)

	_DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 _DPI_AWARENESS_CONTEXT = (1 << intSize) - 4
)

type _MONITOR_DPI_TYPE int32

const (
	_MDT_EFFECTIVE_DPI _MONITOR_DPI_TYPE = 0
	_MDT_ANGULAR_DPI   _MONITOR_DPI_TYPE = 1
	_MDT_RAW_DPI       _MONITOR_DPI_TYPE = 2
	_MDT_DEFAULT       _MONITOR_DPI_TYPE = _MDT_EFFECTIVE_DPI
)

type _PROCESS_DPI_AWARENESS int32

const (
	_PROCESS_DPI_UNAWARE           _PROCESS_DPI_AWARENESS = 0
	_PROCESS_SYSTEM_DPI_AWARE      _PROCESS_DPI_AWARENESS = 1
	_PROCESS_PER_MONITOR_DPI_AWARE _PROCESS_DPI_AWARENESS = 2
)

type _BITMAPV5HEADER struct {
	bV5Size          uint32
	bV5Width         int32
	bV5Height        int32
	bV5Planes        uint16
	bV5BitCount      uint16
	bV5Compression   uint32
	bV5SizeImage     uint32
	bV5XPelsPerMeter int32
	bV5YPelsPerMeter int32
	bV5ClrUsed       uint32
	bV5ClrImportant  uint32
	bV5RedMask       uint32
	bV5GreenMask     uint32
	bV5BlueMask      uint32
	bV5AlphaMask     uint32
	bV5CSType        uint32
	bV5Endpoints     _CIEXYZTRIPLE
	bV5GammaRed      uint32
	bV5GammaGreen    uint32
	bV5GammaBlue     uint32
	bV5Intent        uint32
	bV5ProfileData   uint32
	bV5ProfileSize   uint32
	bV5Reserved      uint32
}

type _CIEXYZ struct {
	ciexyzX _FXPT2DOT30
	ciexyzY _FXPT2DOT30
	ciexyzZ _FXPT2DOT30
}

type _CIEXYZTRIPLE struct {
	ciexyzRed   _CIEXYZ
	ciexyzGreen _CIEXYZ
	ciexyzBlue  _CIEXYZ
}

type _DEV_BROADCAST_DEVICEINTERFACE_W struct {
	dbcc_size       uint32
	dbcc_devicetype uint32
	dbcc_reserved   uint32
	dbcc_classguid  windows.GUID
	dbcc_name       [1]uint16
}

type _DEVMODEW struct {
	dmDeviceName       [_CCHDEVICENAME]uint16
	dmSpecVersion      uint16
	dmDriverVersion    uint16
	dmSize             uint16
	dmDriverExtra      uint16
	dmFields           uint32
	dmPosition         _POINTL
	_                  [8]byte // the rest of union
	dmColor            int16
	dmDuplex           int16
	dmYResolution      int16
	dmTTOption         int16
	dmCollate          int16
	dmFormName         [_CCHFORMNAME]uint32
	dmLogPixels        uint16
	dmBitsPerPel       uint32
	dmPelsWidth        uint32
	dmPelsHeight       uint32
	dmDisplayFlags     uint32 // union with DWORD dmNup
	dmDisplayFrequency uint32
	dmICMMethod        uint32
	dmICMIntent        uint32
	dmMediaType        uint32
	dmDitherType       uint32
	dmReserved1        uint32
	dmReserved2        uint32
	dmPanningWidth     uint32
	dmPanningHeight    uint32
}

type _DISPLAY_DEVICEW struct {
	cb           uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	StateFlags   uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

type _FXPT2DOT30 int32

type _ICONINFO struct {
	fIcon    int32
	xHotspot uint32
	yHotspot uint32
	hbmMask  _HBITMAP
	hbmColor _HBITMAP
}

type _MONITORINFO struct {
	cbSize    uint32
	rcMonitor _RECT
	rcWork    _RECT
	dwFlags   uint32
}

type _MONITORINFOEXW struct {
	cbSize    uint32
	rcMonitor _RECT
	rcWork    _RECT
	dwFlags   uint32
	szDevice  [_CCHDEVICENAME]uint16
}

type _MSG struct {
	hwnd     windows.HWND
	message  uint32
	wParam   _WPARAM
	lParam   _LPARAM
	time     uint32
	pt       _POINT
	lPrivate uint32
}

type _POINT struct {
	x int32
	y int32
}

type _POINTL struct {
	x int32
	y int32
}

type _RAWINPUT struct {
	header _RAWINPUTHEADER
	mouse  _RAWMOUSE

	// RAWMOUSE is the biggest among RAWHID, RAWKEYBOARD, and RAWMOUSE.
	// Then, padding is not needed here.
}

type _RAWINPUTHEADER struct {
	dwType  uint32
	dwSize  uint32
	hDevice windows.Handle
	wParam  uintptr
}

type _RAWMOUSE struct {
	usFlags            uint16
	ulButtons          uint32 // TODO: Check alignments
	ulRawButtons       uint32
	lLastX             int32
	lLastY             int32
	ulExtraInformation uint32
}

type _RECT struct {
	left   int32
	top    int32
	right  int32
	bottom int32
}

var (
	gdi32    = windows.NewLazySystemDLL("gdi32.dll")
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")
	shcore32 = windows.NewLazySystemDLL("shcore32.dll")
	user32   = windows.NewLazySystemDLL("user32.dll")

	procCreateBitmap     = gdi32.NewProc("CreateBitmap")
	procCreateDCW        = gdi32.NewProc("CreateDCW")
	procCreateDIBSection = gdi32.NewProc("CreateDIBSection")
	procDeleteDC         = gdi32.NewProc("DeleteDC")
	procDeleteObject     = gdi32.NewProc("DeleteObject")
	procGetDeviceCaps    = gdi32.NewProc("GetDeviceCaps")

	procGetModuleHandleW          = kernel32.NewProc("GetModuleHandleW")
	procIsWindows8Point1OrGreater = kernel32.NewProc("IsWindows8Point1OrGreater")
	procIsWindowsVistaOrGreater   = kernel32.NewProc("IsWindowsVistaOrGreater")
	procTlsAlloc                  = kernel32.NewProc("TlsAlloc")
	procTlsFree                   = kernel32.NewProc("TlsFree")
	procTlsGetValue               = kernel32.NewProc("TlsGetValue")
	procTlsSetValue               = kernel32.NewProc("TlsSetValue")

	procGetDpiForMonitor       = shcore32.NewProc("GetDpiForMonitor")
	procSetProcessDpiAwareness = shcore32.NewProc("SetProcessDpiAwareness")

	procAdjustWindowRectEx            = user32.NewProc("AdjustWindowRectEx")
	procAdjustWindowRectExForDpi      = user32.NewProc("AdjustWindowRectExForDpi")
	procChangeDisplaySettingsExW      = user32.NewProc("ChangeDisplaySettingsExW")
	procCreateIconIndirect            = user32.NewProc("CreateIconIndirect")
	procCreateWindowExW               = user32.NewProc("CreateWindowExW")
	procDispatchMessageW              = user32.NewProc("DispatchMessageW")
	procEnumDisplayDevicesW           = user32.NewProc("EnumDisplayDevicesW")
	procEnumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	procEnumDisplaySettingsW          = user32.NewProc("EnumDisplaySettingsW")
	procEnumDisplaySettingsExW        = user32.NewProc("EnumDisplaySettingsExW")
	procGetDC                         = user32.NewProc("GetDC")
	procGetDpiForWindow               = user32.NewProc("GetDpiForWindow")
	procGetMonitorInfoW               = user32.NewProc("GetMonitorInfoW")
	procLoadCursorW                   = user32.NewProc("LoadCursorW")
	procMapVirtualKeyW                = user32.NewProc("MapVirtualKeyW")
	procPeekMessageW                  = user32.NewProc("PeekMessageW")
	procRegisterDeviceNotificationW   = user32.NewProc("RegisterDeviceNotificationW")
	procReleaseDC                     = user32.NewProc("ReleaseDC")
	procSetCursor                     = user32.NewProc("SetCursor")
	procSetProcessDPIAware            = user32.NewProc("SetProcessDPIAware")
	procSetProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")
	procShowWindow                    = user32.NewProc("ShowWindow")
	procSystemParametersInfoW         = user32.NewProc("SystemParametersInfoW")
	procToUnicode                     = user32.NewProc("ToUnicode")
	procTranslateMessage              = user32.NewProc("TranslateMessage")
)

func _AdjustWindowRectEx(lpRect *_RECT, dwStyle uint32, menu bool, dwExStyle uint32) error {
	var bMenu uintptr
	if menu {
		bMenu = 1
	}
	r, _, e := procAdjustWindowRectEx.Call(uintptr(unsafe.Pointer(lpRect)), uintptr(dwStyle), bMenu, uintptr(dwExStyle))
	if r == 0 {
		return fmt.Errorf("glfwwin: AdjustWindowRectEx failed: %w", e)
	}
	return nil
}

func _AdjustWindowRectExForDpi(lpRect *_RECT, dwStyle uint32, menu bool, dwExStyle uint32, dpi uint32) error {
	var bMenu uintptr
	if menu {
		bMenu = 1
	}
	r, _, e := procAdjustWindowRectExForDpi.Call(uintptr(unsafe.Pointer(lpRect)), uintptr(dwStyle), bMenu, uintptr(dwExStyle), uintptr(dpi))
	if r == 0 {
		return fmt.Errorf("glfwwin: AdjustWindowRectExForDpi failed: %w", e)
	}
	return nil
}

func _AdjustWindowRectExForDpi_Available() bool {
	return procAdjustWindowRectExForDpi.Find() == nil
}

func _ChangeDisplaySettingsExW(deviceName string, lpDevMode *_DEVMODEW, hwnd windows.HWND, dwflags uint32, lParam unsafe.Pointer) int32 {
	var lpszDeviceName *uint16
	if deviceName != "" {
		var err error
		lpszDeviceName, err = windows.UTF16PtrFromString(deviceName)
		if err != nil {
			panic("glfwwin: device name must not include a NUL character")
		}
	}

	r, _, _ := procChangeDisplaySettingsExW.Call(uintptr(unsafe.Pointer(lpszDeviceName)), uintptr(unsafe.Pointer(lpDevMode)), uintptr(hwnd), uintptr(dwflags), uintptr(lParam))
	runtime.KeepAlive(lpszDeviceName)
	runtime.KeepAlive(lpDevMode)

	return int32(r)
}

func _CreateBitmap(nWidth int32, nHeight int32, nPlanes uint32, nBitCount uint32, lpBits unsafe.Pointer) (_HBITMAP, error) {
	r, _, e := procCreateBitmap.Call(uintptr(nWidth), uintptr(nHeight), uintptr(nPlanes), uintptr(nBitCount), uintptr(lpBits))
	if r == 0 {
		return 0, fmt.Errorf("glfwwin: CreateBitmap failed: %w", e)
	}
	return _HBITMAP(r), nil
}

func _CreateDCW(driver string, device string, port string, pdm *_DEVMODEW) (_HDC, error) {
	var lpszDriver *uint16
	if driver != "" {
		var err error
		lpszDriver, err = windows.UTF16PtrFromString(driver)
		if err != nil {
			panic("glfwwin: driver must not include a NUL character")
		}
	}

	var lpszDevice *uint16
	if device != "" {
		var err error
		lpszDevice, err = windows.UTF16PtrFromString(device)
		if err != nil {
			panic("glfwwin: device must not include a NUL character")
		}
	}

	var lpszPort *uint16
	if port != "" {
		var err error
		lpszPort, err = windows.UTF16PtrFromString(port)
		if err != nil {
			panic("glfwwin: port must not include a NUL character")
		}
	}

	r, _, e := procCreateDCW.Call(uintptr(unsafe.Pointer(lpszDriver)), uintptr(unsafe.Pointer(lpszDevice)), uintptr(unsafe.Pointer(lpszPort)))
	runtime.KeepAlive(lpszDriver)
	runtime.KeepAlive(lpszDevice)
	runtime.KeepAlive(lpszPort)
	if r == 0 {
		return 0, fmt.Errorf("glfwwin: CreateDCW failed: %w", e)
	}
	return _HDC(r), nil
}

func _CreateDIBSection(hdc _HDC, pbmi *_BITMAPV5HEADER, usage uint32, hSection windows.Handle, offset uint32) (_HBITMAP, *byte, error) {
	// pbmi is originally *BITMAPINFO.
	var bits *byte
	r, _, e := procCreateDIBSection.Call(uintptr(hdc), uintptr(unsafe.Pointer(pbmi)), uintptr(usage), uintptr(unsafe.Pointer(&bits)), uintptr(hSection), uintptr(offset))
	if r == 0 {
		return 0, nil, fmt.Errorf("glfwwin: CreateDIBSection failed: %w", e)
	}
	return _HBITMAP(r), bits, nil
}

func _CreateIconIndirect(piconinfo *_ICONINFO) (_HICON, error) {
	r, _, e := procCreateIconIndirect.Call(uintptr(unsafe.Pointer(piconinfo)))
	if r == 0 {
		return 0, fmt.Errorf("glfwwin: CreateIconIndirect failed: %w", e)
	}
	return _HICON(r), nil
}

func _CreateWindowExW(dwExStyle uint32, className string, windowName string, dwStyle uint32, x, y, nWidth, nHeight int, hWndParent windows.HWND, hMenu _HMENU, hInstance _HINSTANCE, lpParam unsafe.Pointer) (windows.HWND, error) {
	var lpClassName *uint16
	if className != "" {
		var err error
		lpClassName, err = windows.UTF16PtrFromString(className)
		if err != nil {
			panic("glfwwin: class name msut not include a NUL character")
		}
	}

	var lpWindowName *uint16
	if windowName != "" {
		var err error
		lpWindowName, err = windows.UTF16PtrFromString(windowName)
		if err != nil {
			panic("glfwwin: window name msut not include a NUL character")
		}
	}

	r, _, e := procCreateWindowExW.Call(
		uintptr(dwExStyle), uintptr(unsafe.Pointer(lpClassName)), uintptr(unsafe.Pointer(lpWindowName)), uintptr(dwStyle),
		uintptr(x), uintptr(y), uintptr(nWidth), uintptr(nHeight),
		uintptr(hWndParent), uintptr(hMenu), uintptr(hInstance), uintptr(lpParam))
	runtime.KeepAlive(lpClassName)
	runtime.KeepAlive(lpWindowName)

	if r == 0 {
		return 0, fmt.Errorf("glfwwin: CreateWindowExW failed: %w", e)
	}
	return windows.HWND(r), nil
}

func _DeleteDC(hdc _HDC) error {
	r, _, e := procDeleteDC.Call(uintptr(hdc))
	if r == 0 {
		return fmt.Errorf("glfwwin: DeleteDC failed: %w", e)
	}
	return nil
}

func _DeleteObject(ho _HGDIOBJ) error {
	r, _, e := procDeleteObject.Call(uintptr(ho))
	if r == 0 {
		return fmt.Errorf("glfwwin: DeleteObject failed: %w", e)
	}
	return nil
}

func _DispatchMessageW(lpMsg *_MSG) _LRESULT {
	r, _, _ := procDispatchMessageW.Call(uintptr(unsafe.Pointer(lpMsg)))
	return _LRESULT(r)
}

func _EnumDisplayDevicesW(device string, iDevNum uint32, dwFlags uint32) (_DISPLAY_DEVICEW, bool) {
	var lpDevice *uint16
	if device != "" {
		var err error
		lpDevice, err = windows.UTF16PtrFromString(device)
		if err != nil {
			panic("glfwwin: device name must not include a NUL character")
		}
	}

	var displayDevice _DISPLAY_DEVICEW
	displayDevice.cb = uint32(unsafe.Sizeof(displayDevice))
	r, _, _ := procEnumDisplayDevicesW.Call(uintptr(unsafe.Pointer(lpDevice)), uintptr(iDevNum), uintptr(unsafe.Pointer(&displayDevice)), uintptr(dwFlags))
	runtime.KeepAlive(lpDevice)

	if r == 0 {
		return _DISPLAY_DEVICEW{}, false
	}
	return displayDevice, true
}

func _EnumDisplayMonitors(hdc _HDC, lprcClip *_RECT, lpfnEnum uintptr, dwData _LPARAM) error {
	r, _, e := procEnumDisplayMonitors.Call(uintptr(hdc), uintptr(unsafe.Pointer(lprcClip)), uintptr(lpfnEnum), uintptr(dwData))
	if r == 0 {
		return fmt.Errorf("glfwwin: EnumDisplayMonitors failed: %w", e)
	}
	return nil
}

func _EnumDisplaySettingsExW(deviceName string, iModeNum uint32, dwFlags uint32) (_DEVMODEW, bool) {
	var lpszDeviceName *uint16
	if deviceName != "" {
		var err error
		lpszDeviceName, err = windows.UTF16PtrFromString(deviceName)
		if err != nil {
			panic("glfwwin: device name must not include a NUL character")
		}
	}

	var dm _DEVMODEW
	dm.dmSize = uint16(unsafe.Sizeof(dm))

	r, _, _ := procEnumDisplaySettingsExW.Call(uintptr(unsafe.Pointer(lpszDeviceName)), uintptr(iModeNum), uintptr(unsafe.Pointer(&dm)), uintptr(dwFlags))
	runtime.KeepAlive(lpszDeviceName)

	if r == 0 {
		return _DEVMODEW{}, false
	}
	return dm, true
}

func _EnumDisplaySettingsW(deviceName string, iModeNum uint32) (_DEVMODEW, bool) {
	var lpszDeviceName *uint16
	if deviceName != "" {
		var err error
		lpszDeviceName, err = windows.UTF16PtrFromString(deviceName)
		if err != nil {
			panic("glfwwin: device name must not include a NUL character")
		}
	}

	var dm _DEVMODEW
	dm.dmSize = uint16(unsafe.Sizeof(dm))

	r, _, _ := procEnumDisplaySettingsW.Call(uintptr(unsafe.Pointer(lpszDeviceName)), uintptr(iModeNum), uintptr(unsafe.Pointer(&dm)))
	runtime.KeepAlive(lpszDeviceName)

	if r == 0 {
		return _DEVMODEW{}, false
	}
	return dm, true
}

func _GetDC(hWnd windows.HWND) (_HDC, error) {
	r, _, e := procGetDC.Call(uintptr(hWnd))
	if r == 0 {
		return 0, fmt.Errorf("glfwwin: GetDC failed: %w", e)
	}
	return _HDC(r), nil
}

func _GetDeviceCaps(hdc _HDC, index int32) int32 {
	r, _, _ := procGetDeviceCaps.Call(uintptr(hdc), uintptr(index))
	return int32(r)
}

func _GetDpiForWindow(hwnd windows.HWND) uint32 {
	r, _, _ := procGetDpiForWindow.Call(uintptr(hwnd))
	return uint32(r)
}

func _GetDpiForWindow_Available() bool {
	return procGetDpiForWindow.Find() == nil
}

func _GetModuleHandleW(moduleName string) (_HMODULE, error) {
	var lpModuleName *uint16
	if moduleName != "" {
		var err error
		lpModuleName, err = windows.UTF16PtrFromString(moduleName)
		if err != nil {
			panic("glfwwin: module name must not include a NUL character")
		}
	}

	r, _, e := procGetModuleHandleW.Call(uintptr(unsafe.Pointer(lpModuleName)))
	runtime.KeepAlive(lpModuleName)

	if r == 0 {
		return 0, fmt.Errorf("glfwwin: GetModuleHandleW failed: %w", e)
	}
	return _HMODULE(r), nil
}

func _GetMonitorInfoW(hMonitor _HMONITOR) (_MONITORINFO, bool) {
	var mi _MONITORINFO
	mi.cbSize = uint32(unsafe.Sizeof(mi))
	r, _, _ := procGetMonitorInfoW.Call(uintptr(hMonitor), uintptr(unsafe.Pointer(&mi)))
	if r == 0 {
		return _MONITORINFO{}, false
	}
	return mi, true
}

func _GetMonitorInfoW_Ex(hMonitor _HMONITOR) (_MONITORINFOEXW, bool) {
	var mi _MONITORINFOEXW
	mi.cbSize = uint32(unsafe.Sizeof(mi))
	r, _, _ := procGetMonitorInfoW.Call(uintptr(hMonitor), uintptr(unsafe.Pointer(&mi)))
	if r == 0 {
		return _MONITORINFOEXW{}, false
	}
	return mi, true
}

func _GetDpiForMonitor(hmonitor _HMONITOR, dpiType _MONITOR_DPI_TYPE) (dpiX, dpiY uint32, err error) {
	r, _, e := procGetDpiForMonitor.Call(uintptr(hmonitor), uintptr(dpiType), uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))
	if r != 0 {
		return 0, 0, fmt.Errorf("glfwwin: GetDpiForMonitor failed: %w", e)
	}
	return dpiX, dpiY, nil
}

func _IsWindows8Point1OrGreater() bool {
	r, _, _ := procIsWindows8Point1OrGreater.Call()
	return r != 0
}

func _IsWindowsVistaOrGreater() bool {
	r, _, _ := procIsWindowsVistaOrGreater.Call()
	return r != 0
}

func _LoadCursorW(hInstance _HINSTANCE, lpCursorName uintptr) (_HCURSOR, error) {
	r, _, e := procLoadCursorW.Call(uintptr(hInstance), lpCursorName)
	if r == 0 {
		return 0, fmt.Errorf("glfwwin: LoadCursorW: %w", e)
	}
	return _HCURSOR(r), nil
}

func _MapVirtualKeyW(uCode uint32, uMapType uint32) uint32 {
	r, _, _ := procMapVirtualKeyW.Call(uintptr(uCode), uintptr(uMapType))
	return uint32(r)
}

func _PeekMessageW(lpMsg *_MSG, hWnd windows.HWND, wMsgFilterMin uint32, wMsgFilterMax uint32, wRemoveMsg uint32) bool {
	r, _, _ := procPeekMessageW.Call(uintptr(unsafe.Pointer(lpMsg)), uintptr(hWnd), uintptr(wMsgFilterMin), uintptr(wMsgFilterMax), uintptr(wRemoveMsg))
	return r != 0
}

func _RegisterDeviceNotificationW(hRecipient windows.Handle, notificationFilter unsafe.Pointer, flags uint32) (_HDEVNOTIFY, error) {
	r, _, e := procRegisterDeviceNotificationW.Call(uintptr(hRecipient), uintptr(notificationFilter), uintptr(flags))
	if r == 0 {
		return 0, fmt.Errorf("glfwwin: RegisterDeviceNotificationW failed: %w", e)
	}
	return _HDEVNOTIFY(r), nil
}

func _ReleaseDC(hWnd windows.HWND, hDC _HDC) int32 {
	r, _, _ := procReleaseDC.Call(uintptr(hWnd), uintptr(hDC))
	return int32(r)
}

func _SetCursor(hCursor _HCURSOR) _HCURSOR {
	r, _, _ := procSetCursor.Call(uintptr(hCursor))
	return _HCURSOR(r)
}

func _SetProcessDPIAware() bool {
	r, _, _ := procSetProcessDPIAware.Call()
	return r != 0
}

func _SetProcessDpiAwareness(value _PROCESS_DPI_AWARENESS) error {
	r, _, e := procSetProcessDpiAwareness.Call(uintptr(value))
	if windows.Handle(r) != windows.S_OK {
		return fmt.Errorf("glfwwin: SetProcessDpiAwareness failed: %w", e)
	}
	return nil
}

func _SetProcessDpiAwarenessContext(value _DPI_AWARENESS_CONTEXT) error {
	r, _, e := procSetProcessDpiAwarenessContext.Call(uintptr(value))
	if r == 0 {
		return fmt.Errorf("glfwwin: SetProcessDpiAwarenessContext failed: %w", e)
	}
	return nil
}

func _SetProcessDpiAwarenessContext_Available() bool {
	return procSetProcessDpiAwarenessContext.Find() == nil
}

func _ShowWindow(hWnd windows.HWND, nCmdShow int) bool {
	r, _, _ := procShowWindow.Call(uintptr(hWnd), uintptr(nCmdShow))
	return r != 0
}

func _SystemParametersInfoW(uiAction uint32, uiParam uint32, pvParam unsafe.Pointer, fWinIni uint32) error {
	r, _, e := procSystemParametersInfoW.Call(uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni))
	if r == 0 {
		return fmt.Errorf("glfwwin: SystemParametersInfoW failed: %w", e)
	}
	return nil
}

func _TlsAlloc() (uint32, error) {
	r, _, e := procTlsAlloc.Call()
	if uint32(r) == _TLS_OUT_OF_INDEXES {
		return 0, fmt.Errorf("glfwwin: TlsAlloc failed: %w", e)
	}
	return uint32(r), nil
}

func _TlsFree(dwTlsIndex uint32) error {
	r, _, e := procTlsFree.Call(uintptr(dwTlsIndex))
	if r == 0 {
		return fmt.Errorf("glfwwin: TlsFree failed: %w", e)
	}
	return nil
}

func _TlsGetValue(dwTlsIndex uint32) (uintptr, error) {
	r, _, e := procTlsGetValue.Call(uintptr(dwTlsIndex))
	if r == 0 && e != windows.ERROR_SUCCESS {
		return 0, fmt.Errorf("glfwwin: TlsGetValue failed: %w", e)
	}
	return r, nil
}

func _TlsSetValue(dwTlsIndex uint32, lpTlsValue uintptr) error {
	r, _, e := procTlsSetValue.Call(uintptr(dwTlsIndex), lpTlsValue)
	if r == 0 {
		return fmt.Errorf("glfwwin: TlsSetValue failed: %w", e)
	}
	return nil
}

func _ToUnicode(wVirtKey uint32, wScanCode uint32, lpKeyState *byte, pwszBuff *uint16, cchBuff int32, wFlags uint32) int32 {
	r, _, _ := procToUnicode.Call(uintptr(wVirtKey), uintptr(wScanCode),
		uintptr(unsafe.Pointer(lpKeyState)), uintptr(unsafe.Pointer(pwszBuff)), uintptr(cchBuff), uintptr(wFlags))
	return int32(r)
}

func _TranslateMessage(lpMsg *_MSG) bool {
	r, _, _ := procTranslateMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return r != 0
}
