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

func IsExecutableCurrentUser(info os.FileInfo) bool {
	mode := info.Mode()

	if IsExecutableAll(mode) {
		return true
	} else {
		user, _ := user.Current()

		if uid, _ := strconv.ParseUint(user.Uid, 10, 32); IsExecutableOwner(mode) && info.Sys().(*syscall.Stat_t).Uid == uint32(uid) {
			return true
		}
	}

	return false
}