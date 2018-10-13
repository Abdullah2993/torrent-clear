package main

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

func Install(path string) error {
	k, _, err := registry.CreateKey(syscall.HKEY_CLASSES_ROOT, `SystemFileAssociations\.torrent\shell\clear`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetStringValue("", "Clear Comments")
	if err != nil {
		return err
	}

	k, _, err = registry.CreateKey(syscall.HKEY_CLASSES_ROOT, `SystemFileAssociations\.torrent\shell\clear\command`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = k.SetStringValue("", fmt.Sprintf(`"%s" "%%1"`, path))
	if err != nil {
		return err
	}
	return nil
}
