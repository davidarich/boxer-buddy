package user

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.MustLoadDLL("user32.dll")
	procEnumWindows         = user32.MustFindProc("EnumWindows")
	procGetWindowTextW      = user32.MustFindProc("GetWindowTextW")
	procSetForegroundWindow = user32.MustFindProc("SetForegroundWindow")
	procGetWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
	//procSystemParametersInfoW = user32.MustFindProc("SystemParametersInfoW")
)

// adapter for interacting with application windows on a Windows OS
type WindowsAdapter struct{}

/*
Enumerates all top-level windows on the screen by passing the handle to each window, in turn, to an application-defined callback function.
EnumWindows continues until the last top-level window is enumerated or the callback function returns FALSE.

ref: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
*/
func (w *WindowsAdapter) EnumWindows(lpEnumFunc uintptr, lParam uintptr) (err error) {
	r1, _, errString := syscall.SyscallN(procEnumWindows.Addr(), 2, uintptr(lpEnumFunc), uintptr(lParam), 0)
	if r1 == 0 {
		if errString != 0 {
			err = error(errString)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

/*
Copies the text of the specified window's title bar (if it has one) into a buffer.
If the specified window is a control, the text of the control is copied.
However, GetWindowText cannot retrieve the text of a control in another application.

ref: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowtextw
*/
func (w *WindowsAdapter) GetWindowTextW(hWnd syscall.Handle, str *uint16, maxCount int32) (length int32, err error) {
	r1, _, errNum := syscall.SyscallN(procGetWindowTextW.Addr(), 3, uintptr(hWnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	length = int32(r1)
	if length != 0 {
		return
	}
	if errNum != 0 {
		err = error(errNum)
		return
	}
	err = syscall.EINVAL
	return
}

/*
Brings the thread that created the specified window into the foreground and activates the window.
Keyboard input is directed to the window, and various visual cues are changed for the user.
The system assigns a slightly higher priority to the thread that created the foreground window than it does to other threads.

ref: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
*/
func (w *WindowsAdapter) SetForegroundWindow(hWnd syscall.Handle) (ok bool) {
	r1, _, _ := syscall.SyscallN(procSetForegroundWindow.Addr(), 1, uintptr(hWnd))
	ok = r1 != 0
	return
}

// returns handles for all windows
func (w *WindowsAdapter) Find() (hWnds []syscall.Handle, err error) {
	// this callback is executed for every window handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		hWnds = append(hWnds, h)
		return 1 // returning 1 will include the window handle in the result
	})
	w.EnumWindows(cb, 0)
	if len(hWnds) == 0 {
		err = fmt.Errorf("no windows found")
	}
	return
}

// set focus to the window matching a given pid
func (w *WindowsAdapter) FocusWindowByPid(pid int) {
	// get all windows
	hWnds, err := w.Find()
	if err != nil {
		return
	}

	// find the window with the current pid
	for i := range hWnds {
		if pid == w.GetWindowThreadProcessId(hWnds[i]) {
			_ = w.SetForegroundWindow(hWnds[i])
			return
		}
	}
}

/*
Retrieves the identifier of the thread that created the specified window and, optionally,
the identifier of the process that created the window.

ref: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowthreadprocessid
*/
func (w *WindowsAdapter) GetWindowThreadProcessId(hWnd syscall.Handle) (pid int) {
	// todo: implement syscall for FocusGame feature
	_,_,_ = syscall.SyscallN(procGetWindowThreadProcessId.Addr(), 1, uintptr(hWnd))
	return 0
}

func NewWindowsAdapter() *WindowsAdapter {
	return &WindowsAdapter{}
}
