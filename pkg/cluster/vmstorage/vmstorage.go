package vmstorage

import "os/exec"



type storage struct {
	cmd *exec.Cmd
}

func New() *storage {
	return &storage{}
}
