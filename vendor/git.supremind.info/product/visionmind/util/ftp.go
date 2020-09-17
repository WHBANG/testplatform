package util

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jlaffaye/ftp"
	log "qiniupkg.com/x/log.v7"
)

type FtpConn struct {
	*ftp.ServerConn
	isUtf8Available bool
	host            string
}

func NewFtp(host string, verbose bool, remoteBasePath string, isUtf8Available bool) (*FtpConn, error) {
	var (
		user     = "anonymous"
		password = "anonymous"
		address  string
	)
	userPassIndex := strings.Index(host, "@")
	if userPassIndex == -1 {
		address = host
	} else {
		userPass := host[:userPassIndex]
		userIndex := strings.Index(userPass, ":")
		if userIndex == -1 {
			user = userPass
			password = ""
		} else {
			user = userPass[:userIndex]
			password = userPass[userIndex+1:]
		}
		address = host[userPassIndex+1:]
	}

	options := []ftp.DialOption{}
	if verbose {
		options = append(options, ftp.DialWithDebugOutput(os.Stderr))
	}

	ftpConn, err := ftp.Dial(address, options...)
	if err != nil {
		log.Errorf("ftp.Dial(%s): %s", address, err.Error())
		return nil, err
	}
	ftpConnection := &FtpConn{ServerConn: ftpConn, isUtf8Available: isUtf8Available, host: host}

	err = ftpConnection.Login(user, password)
	if err != nil {
		log.Errorf("ftp.Login(%s): %s", user, err.Error())
		return nil, err
	}

	remoteBaseEncoded, err := ftpConnection.handleEncoding(remoteBasePath)
	if err != nil {
		return nil, err
	}
	err = ftpConnection.ChangeDir(remoteBaseEncoded)
	if err != nil {
		log.Errorf("ftp.ChangeDir(%s): %s", remoteBasePath, err.Error())
		return nil, err
	}

	log.Infof("connected to the ftp server %s with user %s at %s\n", address, user, remoteBasePath)
	return ftpConnection, nil
}

func (f *FtpConn) FtpUploadLocal(filePath, newFileName, remotePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("os.Open(%s): %s", filePath, err.Error())
		return "", err
	}
	defer file.Close()
	return f.ftpUpload(file, newFileName, remotePath)
}

func (f *FtpConn) FtpUploadRemote(remote, newFileName, remotePath string) (string, error) {
	resp, err := http.DefaultClient.Get(remote)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return f.ftpUpload(resp.Body, newFileName, remotePath)
}

func (f *FtpConn) ftpUpload(reader io.Reader, newFileName, remotePath string) (string, error) {
	backupDir, err := f.CurrentDir()
	if err != nil {
		log.Errorf("ftp.CurrentDir(): %s", err.Error())
		return "", err
	}

	err = f.changeDir(remotePath)
	if err != nil {
		return "", err
	}

	fileNameEncoded, err := f.handleEncoding(newFileName)
	if err != nil {
		return "", err
	}
	err = f.Stor(fileNameEncoded, reader)
	if err != nil {
		log.Errorf("ftp.Stor(%s): %s", newFileName, err.Error())
		return "", err
	}
	filePath, err := f.CurrentDir()
	if err != nil {
		return "", err
	}
	filePath = "ftp://" + path.Join(f.host, filePath, newFileName)
	err = f.ChangeDir(backupDir)
	return filePath, err
}

func (f *FtpConn) changeDir(path string) error {
	pathList := strings.Split(path, "/")
	for _, subPath := range pathList {
		pathDecoded, err := f.handleEncoding(subPath)
		if err != nil {
			return err
		}
		err = f.ChangeDir(pathDecoded)
		if err != nil {
			err = f.MakeDir(pathDecoded)
			if err != nil {
				log.Errorf("ftp.MakeDir(%s): %s", subPath, err.Error())
				return err
			}
			err = f.ChangeDir(pathDecoded)
			if err != nil {
				log.Errorf("ftp.ChangeDir(%s): %s", subPath, err.Error())
				return err
			}
		}
	}
	return nil
}

func (f *FtpConn) handleEncoding(raw string) (string, error) {
	if f.isUtf8Available {
		return raw, nil
	}

	GBKEncoded, err := EncodeGBK(raw)
	if err != nil {
		log.Errorf("util.EncodeGBK(%s): %s", raw, err.Error())
		return "", err
	}
	return GBKEncoded, nil
}
