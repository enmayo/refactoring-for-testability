package main

import (
	"io/ioutil"
	"os"
)

// Ideally, these would be moved to a more central place where they can be utilized by various other packages that depend on them
// ----------------------------------------------------
type WriteFileFunc func(filename string, data []byte, perm os.FileMode) error
type ReadDirFunc func(dirname string) ([]os.FileInfo, error)

func NewIoutilPkg() *ioutilPkgImpl{
	return &ioutilPkgImpl{ioutil.WriteFile, ioutil.ReadDir}
}

type IoutilPkg interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
	ReadDir(dirname string) ([]os.FileInfo, error)
}

type ioutilPkgImpl struct {
	// Having these members to allow for testing this shim
	writeFile WriteFileFunc
	readDir ReadDirFunc
}

func (ioutilpkg *ioutilPkgImpl) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutilpkg.writeFile(filename, data, perm)
}

func (ioutilpkg *ioutilPkgImpl) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutilpkg.readDir(dirname)
}
// ----------------------------------------------------
