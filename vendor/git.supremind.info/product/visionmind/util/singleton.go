// +build !windows

package util

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	log "qiniupkg.com/x/log.v7"
)

// Flock holds the lock data
type Flock struct {
	lockHolder *os.File
	lockFile   string
	lockTime   time.Time
}

// NewFlock create a Flock
func NewFlock(file string) (f *Flock, err error) {
	if file == "" {
		err = errors.New("cannot create flock on empty path")
		return
	}
	var (
		finfo  os.FileInfo
		holder *os.File
	)

	holder, err = os.Create(file)
	if err != nil {
		err = errors.New("lockFile create failed")
		return
	}
	finfo, _ = os.Stat(file)

	return &Flock{
		lockFile:   file,
		lockTime:   finfo.ModTime(),
		lockHolder: holder,
	}, nil
}

func (f *Flock) lock() (err error) {
	return syscall.Flock(int(f.lockHolder.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

func (f *Flock) unlock(clear bool) {
	syscall.Flock(int(f.lockHolder.Fd()), syscall.LOCK_UN)
	if f.lockHolder != nil {
		f.lockHolder.Close()
		if clear {
			os.Remove(f.lockFile)
		}
	}
}

// Lock lock until the lockfile gets removed or modified
func (f *Flock) Lock(callback func(error)) (err error) {
	var finfo os.FileInfo

	if callback != nil {
		defer func() {
			callback(err)
		}()
	}

	if e := f.lock(); e != nil {
		err = fmt.Errorf("try lock file %s failed", f.lockFile)
		return
	}
	for {
		time.Sleep(time.Second * 5)
		if finfo, err = os.Stat(f.lockFile); err != nil {
			f.unlock(true)
			return fmt.Errorf("lock file %s check failed", f.lockFile)
		}
		log.Debug(finfo.ModTime(), f.lockTime)
		if finfo.ModTime() != f.lockTime {
			// 锁文件被修改了, 则退出
			f.unlock(false) // NOTE 文件同时被修改了, 说明有其他进程在hold文件, 不删除文件
			return fmt.Errorf("lock file %s invalid modified", f.lockFile)
		}
	}
}
