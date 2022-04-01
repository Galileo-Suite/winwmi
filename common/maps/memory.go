package maps

// Win32_PhysicalMemory describes the physical memory available
// to the Windows OS.
// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-physicalmemory
type Win32_PhysicalMemory struct {
	Capacity    uint64
	Description string
}
