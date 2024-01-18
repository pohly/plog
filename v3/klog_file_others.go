//go:build !windows
// +build !windows

package plog

import (
	"os/user"
)

func getUserName() string {
	userNameOnce.Do(func() {
		current, err := user.Current()
		if err == nil {
			userName = current.Username
		}
	})

	return userName
}
