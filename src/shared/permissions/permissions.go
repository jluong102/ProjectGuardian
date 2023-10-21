package permissions 

import "os"
import "os/user"
import "syscall"
import "strconv"

func IsExecutableAll(mode os.FileMode) bool {
	return mode & 0001 != 0
}

func IsExecutableOwner(mode os.FileMode) bool {
	return mode & 0100 != 0
}

func IsExecutableGroup(mode os.FileMode) bool {
	return mode & 0010 != 0
}

// Probably a better way to check the uid/gid then this 
func IsExecutableCurrentUser(info os.FileInfo) bool {
	mode := info.Mode()

	if IsExecutableAll(mode) {
		return true
	} else {
		user, _ := user.Current()

		if uid, _ := strconv.ParseUint(user.Uid, 10, 32); IsExecutableOwner(mode) && info.Sys().(*syscall.Stat_t).Uid == uint32(uid) {
			return true
		} else if gid, _ := strconv.ParseUint(user.Gid, 10, 32); IsExecutableGroup(mode) && info.Sys().(*syscall.Stat_t).Gid == uint32(gid) {
			return true
		}
	}

	return false
}