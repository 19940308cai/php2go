// php2go functions

package php2go

import (
	"syscall"
)

// umask()
func Umask(mask int) int {
	return syscall.Umask(mask)
}

// disk_free_space()
func DiskFreeSpace(directory string) (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(directory, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Bfree * uint64(fs.Bsize), nil
}

// disk_total_space()
func DiskTotalSpace(directory string) (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(directory, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Blocks * uint64(fs.Bsize), nil
}
