package maps

// In this wmi package you create the type that matches the class you want to
// collect and then you add the members that you need.
// see as an example:
// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-process
type Win32_Process struct {
	Name            string
	ParentProcessId uint32
}
