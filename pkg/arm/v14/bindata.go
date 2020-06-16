// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package arm generated by go-bindata.// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5d\x73\x1b\xb7\xce\xbe\x2e\x7f\x05\xba\x72\x9b\x26\x35\xb5\x76\xda\xbe\x6f\x47\xad\x33\xe3\x38\x4e\xc7\xa7\x6e\xec\x63\x27\xa7\x17\x69\x26\x43\x2d\xb1\x12\x23\x2e\xb9\xe5\x87\x64\x55\xd1\x7f\x3f\x03\xee\x4a\x96\x2c\xd9\x49\x1a\xf7\xe4\xc2\x91\x48\x10\x20\xc1\x07\x78\x00\xaa\xf3\x65\xde\x57\x26\xef\x0b\x3f\x04\x8e\x57\x8c\x75\xe0\xb9\x75\x10\xd0\x07\x65\x06\x3d\xd0\x76\x00\xc2\x48\x90\xce\xd6\x20\xb4\x86\xe0\x44\x59\xaa\x02\xc2\x50\x04\x98\xd8\xa8\x25\x38\x1b\x03\xc2\x58\x09\x08\x43\x84\x4a\xf8\x80\x0e\x8e\x4f\x9f\xb2\x0e\x5c\x1c\x5f\x9e\xbd\xba\x38\x3a\xfe\xe5\xe2\xec\xd5\xf9\x41\x36\xb5\xd1\x71\x87\xde\x46\x57\x20\x1f\x38\x1b\xeb\x8c\x75\xe0\xec\xf2\xed\xf3\x7f\x3f\x7b\x71\x90\xd9\x1a\x8d\x1f\xaa\x32\x74\x77\xd6\x56\x76\xad\x17\x12\xc7\xdd\x42\xdb\x28\x33\xd6\x61\x1d\x50\x75\x10\x7d\x8d\x1e\xf8\x09\x9c\xbc\x38\x7f\xf5\x12\xb8\x87\x9d\x6f\xa4\x1a\xc0\xb7\x7e\x68\x5d\x80\x6c\xa7\xd5\x9b\xc1\x7b\x08\x42\x69\xe0\xfb\x0f\x81\xbf\x83\xd3\xb3\x5f\x80\x73\x6d\x07\xbc\x76\x58\xaa\x2b\xc8\x7e\x7d\xf5\xf4\x18\x48\x14\x9e\x5d\x9c\x9d\xf7\xb2\xcf\xd3\x4f\x3a\x18\x9b\xcd\x40\x95\xd0\x3d\xb2\xa6\x54\x83\xee\x25\x16\xd1\xa9\x30\x3d\x17\xa1\x18\x9e\x8b\x62\x24\x06\xe8\x61\x3e\x67\xda\x0e\x06\xe8\x80\x87\xd6\x71\xdc\x07\xe1\x42\xac\xbb\x7e\x08\x99\x32\x3e\x08\xad\x95\x19\x80\x43\x09\xe4\xf2\x42\x1a\x28\x92\xce\xe8\x44\x50\xd6\x80\x35\xb0\xf3\xcd\xd0\xfa\x60\x44\x85\x0f\x33\x56\x88\x00\x4f\xf2\xb1\x70\xb9\x56\xfd\x7c\x1a\xab\xbc\xd0\x0a\x4d\xe0\x05\xba\xd0\xad\xb1\x82\x9f\x7f\x7e\x70\x7c\xf6\xfc\x01\x6d\xf1\x08\x5d\x38\xf4\x4f\xa7\x01\xfd\x72\xaf\x34\xa6\x4a\x55\x88\x80\xbe\xdb\xee\xf5\x02\x6b\xeb\x55\xb0\x6e\x9a\xa6\xe1\x3d\x5c\x06\x47\xfb\x9a\xcf\xd9\xf1\xd9\xf3\xdb\x8d\x8e\x70\x7a\xd3\xe6\xb9\x53\x63\x11\xf0\x57\x9c\x7e\xa2\xe5\x5f\x71\xba\x61\xf8\xa3\x1d\x78\x78\x71\x06\xbe\xbd\x05\x88\xb5\x24\x1b\xf0\x7a\x36\x6b\xf5\xf9\x7f\x59\x65\x3e\x70\x5d\xd9\x2e\x64\x30\x9f\xbf\xd9\x70\x79\x69\x1d\x88\x10\xb0\xaa\x03\x28\x03\xb3\xfd\x6e\xf7\x87\xf9\x4f\x20\x2d\x03\x98\xc6\x0a\xda\x6d\x00\x9f\x02\xff\x13\x3e\xcd\x66\x32\x09\x5f\x7f\x0d\x7d\x87\x62\xc4\x00\xee\x3c\xf0\xeb\xc5\x36\x76\x66\xed\xa7\xf9\x9b\xed\x47\x6f\xf7\xd4\x60\xa8\x14\x4a\xa3\xcc\x18\x10\x66\x5f\xbf\x5e\x59\x0d\x5c\x07\xf8\x01\xde\xbc\xf9\x89\xa2\xdb\x80\xd7\x88\x35\xec\xff\x04\xa8\x3d\x02\x5e\xa9\x40\x5f\x4a\xc5\xa4\x35\xf8\x81\xdb\x70\x58\xd9\xf1\xa7\x81\x99\xbc\x57\x68\x14\x86\x92\x0f\x73\x15\x70\x57\xc2\x9d\xe0\xbe\x03\x84\x6c\x36\x43\x23\xe7\x73\xca\x72\x85\x43\x11\x90\xac\x07\xa1\x0c\x3a\xa8\xa3\xd6\xe4\x25\x87\x81\x55\x23\xa9\x1c\xf0\xfa\x5a\x99\x75\x6a\xa0\x4c\xde\x95\xb6\x18\xa1\xbb\x01\xf7\xf5\xc9\xbc\x39\x51\xf7\x9d\xb7\x66\x15\xf6\xdd\x67\xe8\xd4\x18\x65\xf7\xc8\x56\x7d\x65\x50\x9e\x54\x62\x80\xe7\x51\xeb\xcb\x64\x75\x01\x84\x0d\x88\x6b\x43\xb9\xe7\x16\x6b\x90\x3b\x6b\x43\x4e\x47\x7a\x79\xf6\xec\xac\x07\x12\x35\x06\x4c\xa9\xb8\xb4\x5a\xdb\x09\x69\x4a\xa9\xb6\x39\x33\x79\x59\x94\x94\xa2\x55\x00\xe5\xa1\x2f\x46\x28\x41\x99\x60\xc1\x46\x07\xff\xf9\x0d\x14\xed\xcb\xb3\xb4\x46\x48\x09\xbc\x84\xf6\xd8\x4c\x95\xf0\x25\x0c\x1c\xae\x78\x66\xb1\x0d\x0c\x45\x5e\xfa\x20\xfa\x0d\x50\x18\x80\x9f\xfa\x80\x55\x11\x34\xf8\x60\xeb\x56\x07\x4f\xb7\x19\xeb\x6e\x50\x15\xba\x0f\x4a\x79\x74\x63\x55\xe0\x6d\x72\x2b\xf3\xd5\xa8\xf4\xdd\xab\xd2\xd3\x76\x73\x89\xe3\x5c\x2a\x3f\xca\xc5\x5f\xd1\x61\xbe\xa4\x9c\x5a\xb8\xb0\xcf\x00\xb0\x18\x5a\x78\x70\xb7\x18\x6c\x9c\x11\x48\x3d\x0c\x5c\xfd\x67\xb4\x41\x00\xec\xc1\xde\x03\x78\xf2\xe4\xfa\xe8\xb4\x0d\x1b\x4d\xb8\xb9\x92\x01\x38\xf4\xc1\x3a\x2c\xac\x01\x7e\xb1\x65\xbe\x41\x14\x69\x6a\x51\x24\x05\x56\xd6\xdc\x40\x11\x03\xc8\x88\xb8\x24\x21\xc9\x65\x3d\xc8\xde\xd9\xe8\x8c\xd0\x32\xdb\xa5\x39\xa9\x3c\xb1\x16\xd7\x38\x10\xc5\x94\x3b\x1c\x28\x1f\xdc\x34\xeb\x41\x70\x11\x59\x83\xa7\x75\x5f\x0a\x17\x36\x9d\xb9\x5d\xe0\xc6\xdd\x95\x8a\xb1\xd6\x33\x29\x78\x08\xe3\x6d\x2e\x4b\xd0\xf6\xdd\x17\x56\x62\xca\x5e\x4f\x92\xab\x0d\x49\x7d\xbd\x15\x45\x18\x0a\xb9\x0d\x43\xcb\x5b\xbd\x79\x57\xbe\xf0\x6a\x3f\xd7\xd1\xec\xc1\xfb\xf7\xcd\xe9\x6e\xbb\xd6\x15\xd1\x1b\x06\x9b\x0b\x95\x58\x8a\xa8\x83\xff\xa8\x0b\xa5\x75\xb7\x5f\x67\x9a\x25\xbf\x74\x40\x14\x05\xd6\x54\x44\xc1\x8f\xdf\x7f\xff\x1d\x10\x45\x50\x4c\x0a\x59\x29\xef\x29\x08\x29\xf5\x38\xab\x35\x79\xd2\x3a\x90\x3e\x71\x47\x28\xea\xdd\xb4\xa0\xfd\xf0\x7d\xa2\x91\x2f\x6a\x67\x83\x3d\xd8\x99\x49\x1f\xbe\xfa\x6a\xf7\xd1\x9c\x7d\x51\x5b\x17\x9a\x81\x4e\xe7\xd1\xee\x9c\x7d\x71\x5d\xb1\x1c\xa6\x8a\xea\xe4\xe2\xf8\xf7\xc3\xd3\xd3\xb7\x87\xa7\xa7\x67\xbf\x53\x32\xdb\x49\x4a\x80\x57\x74\xa9\x01\x81\xf3\xe6\xff\x17\xc7\xbf\xd3\xe0\x62\x9a\x4b\x52\x0d\x3b\xe9\x2f\x7f\x07\x87\x47\x47\xc7\xe7\x2f\x81\x4f\xda\x14\xbf\xb0\xc3\xbd\x18\x63\x8b\x59\x3f\xf5\x4d\xd6\xcb\x17\xb3\xe4\x82\x49\x22\x0c\x02\x10\xb9\xc1\x10\x18\x26\x42\x0c\xd0\x84\x54\x53\x1a\x0c\x13\xeb\x46\x10\x83\xd2\x2a\x28\xf4\x30\xb0\x89\x98\x82\x05\x27\x8a\x94\x9c\xa5\xa2\x84\xd5\xa5\x82\xac\x5c\x2e\x76\xd1\x78\xe8\x63\x69\x1d\x82\x34\x9e\xb2\xd8\xc8\xd8\x89\x81\x60\x93\x8f\x1b\x4b\x08\x68\x24\xc4\x1a\x26\x2a\x0c\x81\xc8\x6c\x0a\x3e\x25\x56\x36\x19\x2a\x8d\x89\xe7\x96\x5c\x03\x5c\x3e\x84\x83\x03\xc8\xb2\xc4\x75\xd2\x5e\x33\xdd\x47\x30\x1b\xe1\x9f\xce\xb8\x19\x02\x97\x8d\x14\xcc\xe7\x77\x97\x09\x77\x07\xd2\xb5\x96\xcf\xab\x04\x3e\xda\xca\xa7\x16\x04\xff\xb7\x77\x5b\x45\x40\xc5\xfa\x8b\xb3\x97\xc7\x3d\x38\x31\x50\xc6\x10\x1d\xee\x42\x65\xc7\xd8\xb4\x10\xca\x94\xd6\x55\x2d\xf9\xc7\xe0\x95\x44\xb0\x25\xa0\x19\x2b\x67\x4d\x45\xd7\x3d\x16\x4e\x35\x98\xea\x30\x8f\x01\xbe\xbd\x62\x78\x95\xd0\x79\x79\x78\xf9\xea\xe2\xe4\xe0\xc1\xca\x51\x7e\x4b\x9e\x68\x4f\xd2\xcc\xc3\x7c\xfe\x20\x2d\xe4\x57\x8b\x7c\xe5\xa2\x01\xce\x6b\xa7\xc6\x4a\xe3\x00\x25\x70\x4e\xb5\x05\x5f\x40\x92\x50\x01\x7c\x0c\x79\x2f\xa7\x8f\xbd\xbf\x80\x63\x6b\xed\x6e\xbf\xb5\x37\xc0\xa2\x21\x83\xcd\x0a\xc6\x9a\x9a\x8b\x17\x82\x07\x17\x7d\xa0\xd8\xf0\x18\x52\x54\xc4\x1a\x06\x68\x70\x2c\xd2\x6d\xd2\x88\x0f\xa2\x18\x81\xf0\xe0\x2d\x51\xb5\x4f\x90\x5e\xaf\x92\x94\x07\x2d\x94\x24\x87\x41\x7f\xca\x3a\x49\xa4\x35\x7d\x5d\xd2\xec\x36\x2b\xb5\xf5\x48\xb9\x47\xa5\x40\x69\x43\xe4\x16\xe1\xca\x3a\x64\x1d\xda\x8a\x87\xd2\xd9\x6a\x4d\xb6\x76\xb6\x40\xef\x29\xb2\x26\x8a\x8a\xa5\xa1\xaa\x49\x5f\xb3\x7f\xd6\x6c\xc3\x23\xf8\x61\xd3\x16\x46\x2a\xe7\x0a\x04\x01\x52\x4c\xc1\x1a\x3d\xa5\xd3\xd4\xd8\x24\x42\x69\x0b\xcf\xf2\xe8\x5d\xae\x6d\x21\x74\x6a\x43\xc5\x5f\x1e\x0b\xd9\x1e\x96\x8a\x9e\xbe\xf0\xa8\x95\xa1\xe8\x84\xf3\xfd\x67\x1f\x94\xf7\xb6\x0c\x13\xe1\x3e\x5a\xbe\xd0\xa2\x12\xe3\x85\x34\xeb\x00\x1a\x42\x5a\x4a\x4f\x0d\xf3\xad\xdf\x4a\xcb\x90\x9e\x5d\x13\x64\x34\x95\xf0\x23\xa8\xa4\x97\x0b\x02\x85\xc6\xce\xfa\xd7\xca\x9a\xeb\x91\x52\x47\x34\x61\xf9\x7d\x45\x5d\xbb\x81\xfb\x52\xd7\x1c\xe2\xf3\xb4\xb1\x0e\x9c\x2b\x03\xa3\xd8\xc7\xc6\x73\x09\x45\xd1\x23\x24\xcf\x82\xa8\x15\x27\x59\x74\xcc\x53\x28\x29\xe0\x0e\x21\xf3\x9d\x6f\xe0\x51\x33\xde\x83\x87\xdd\x47\x9d\x3f\xf6\x87\x21\xd4\xbe\x97\xe7\x2b\x25\x7e\x27\x6b\x68\xbf\xad\x6a\x9b\x44\x96\x13\x4b\x9a\xee\xb5\xc5\x7b\x53\xbc\x7c\x61\xe0\xcd\xc0\xbd\xda\x20\xde\x49\x7f\xee\x5f\xab\x97\xf7\xe0\x8e\xd4\x29\x24\x35\x6d\x8f\xc2\xd8\x6c\xc6\x29\xc3\x1b\x84\x9d\xee\x53\x51\x8c\x62\xfd\x54\xdb\xfe\x0b\x22\xc4\x2c\xfb\xe0\xfb\xc4\x92\xdb\xa9\x12\x1a\xa3\x9b\x6e\xf4\x6f\x94\xe9\x02\xd1\x28\x0c\x30\xa4\xb8\xef\x27\x2b\xa9\x95\xbb\x28\xd7\x2b\xa7\xfc\x11\x23\x8e\xa1\x7d\x3c\x53\xee\x60\x7d\xae\x5d\xd7\xb4\x66\x3b\x2b\x72\x7f\x9b\x9a\x8f\x43\x21\x9b\x33\x7f\x26\x3b\xaf\x29\xfa\x27\x09\x7a\xdd\xd0\xfd\x71\xf4\x9d\xfb\x94\x76\x62\xb4\x15\x92\x9c\xd8\x5c\x42\xb6\x4e\xa3\x9b\xcc\xf9\x07\x83\xc4\x9e\x1b\xf1\xd7\xdb\x1c\xda\x26\x9c\xde\xf9\x6a\x67\xc7\x4a\xa2\xcb\x7b\xf9\x5b\x29\x82\xc8\xdf\x12\xdd\xb5\xd2\xab\x00\xe8\xe5\x36\x12\x45\xd3\xd4\x87\x7c\x46\x50\x6a\x0e\xd1\x68\xe2\xfd\x16\xee\x07\xb4\xf2\x46\x04\xcc\xe7\xad\x90\x4c\xcf\xa1\x89\x7b\x0f\xc8\x58\x0b\xc6\xae\xec\xb7\x02\xa2\x48\x73\x0b\x57\xdd\xed\xd0\xd6\xfe\x42\x98\xae\x70\x11\x26\x8f\x17\x8d\xc5\xdf\xc5\x74\x53\xfe\xd0\x99\x3f\x13\xd3\x6b\x8a\xfe\x49\x4c\xaf\x1b\xfa\x1f\x61\xba\xf1\x72\xe2\x75\x23\x6a\x3f\xb4\xe1\x93\x30\x4d\x28\xea\x2d\x3f\x2d\xa7\x56\xf3\x55\x6f\xfd\x5b\x83\x4e\x8e\x70\xfc\xf2\xe8\xd9\xd1\xcb\xd3\xb7\x87\xe7\x27\x07\xd9\x77\xd9\x2d\xa0\x5d\x77\x0a\xc9\x90\x96\x44\xe8\xed\x7e\x17\x40\x59\x8b\x84\x0d\x5c\x52\xdc\x70\x4a\x98\xeb\xb9\xd4\xe0\xa4\x15\x48\xad\xcf\x4a\xc6\x6e\x87\x95\x51\x41\x09\xcd\x0b\x1d\x53\x8c\x66\xad\x0f\xf7\xd2\xbf\x83\x05\xbf\xac\x8d\xf6\x1e\x7f\xf7\xe3\xde\xee\xea\xd0\xfe\x56\xc1\xfd\x4d\xc1\xc7\x5b\x05\x1f\x27\xc1\x6c\xfb\x96\x78\xb0\x23\x34\xc9\x2d\xbc\xb4\x8e\xa7\x56\xff\x86\xa8\x90\x63\x74\x41\x79\xe4\x35\xa2\xe3\xd1\x69\x0f\x5b\xa8\x31\x99\x61\xac\x1a\x6f\x7a\x29\x7f\x74\x63\x6c\xe3\x09\x72\xe9\xcf\x35\x4a\x5a\x7b\x1e\xb8\xa1\xf7\x63\x90\x89\xa9\xe9\xcc\x12\x3d\x53\x17\x3b\x9f\x33\x16\xa2\x41\xc9\x85\xac\xa8\x10\x2f\xa9\x81\xbd\x2e\x66\xda\xb7\x04\x5e\x6b\x91\x7a\x2e\xaa\xd1\x85\xf6\x16\x0c\xa2\xbc\x96\xeb\xa6\x82\xad\x3b\xb6\x3a\x56\xe8\x81\x80\xd1\xbc\x83\xca\x45\x3b\x7d\x55\x7a\x68\x5e\xb7\x0a\x6a\xa2\xa9\xd3\x5e\xbc\x86\x56\xb0\xf7\xff\x3f\xec\x6d\x7b\x15\xbd\x45\x3f\xed\xa3\x79\x90\x4a\x25\x82\x9f\x7a\x6d\x07\xe0\x15\xf5\x04\x13\x84\x4a\x18\x31\x40\x40\xaa\x1b\xc2\x90\x44\xc2\xd0\xd9\x38\x18\xc2\xe2\x4d\x6b\xa5\x8e\x6d\x1f\xb6\x16\x5a\xb6\x56\xba\xb6\xde\x98\x66\x1d\x30\x36\x60\x0f\x44\xb0\x95\x2a\xf8\xb5\xc7\xd2\x1b\x41\xe1\x84\x1f\x82\xb6\xb6\xf6\x10\x4d\x50\x7a\xf1\xeb\x95\xf2\x10\xeb\xcd\xaa\x7c\xab\x96\xa5\xb1\xfb\xf8\xc5\xc7\x17\x43\x94\x31\x39\x6c\x35\x2a\x1d\xf6\xad\x0d\x54\x76\x17\xb6\xaa\xd3\x03\xef\xb6\x47\xfd\x8c\xf9\x61\x0c\x44\x2c\x94\xc2\x9a\x35\xdf\x3e\x66\xb3\x19\xa5\xc8\xf9\x7c\xa3\x2f\xb8\xf3\x3c\xcb\x77\xb5\xc5\xab\xf9\x7f\x03\x00\x00\xff\xff\x0e\x17\xed\x4b\x2c\x1c\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x5d\x6f\xdb\xc8\x15\x7d\x9f\x5f\x71\x56\x5a\xc4\x5d\x6c\x28\x25\x05\xd2\x02\xce\x26\x40\x9a\x64\x81\x74\xd1\xb5\x61\x6f\xdb\x87\x20\x0f\x23\xce\x25\x35\xab\xe1\x5c\x66\x3e\x24\x2b\x8a\xfe\x7b\x71\x49\x2a\x96\x2d\x5b\x4e\x90\xbe\x49\xe4\x9d\xfb\x79\xee\x39\xc3\xf1\x0f\xd3\x99\xf5\xd3\x99\x8e\x73\x14\x74\xa5\xd4\x66\x03\x5b\x61\xf2\x9a\x7d\x65\xeb\xc9\x25\x95\x39\xd8\xb4\x3e\xd7\xa9\x9c\x9f\xeb\x72\xa1\x6b\x8a\xd8\x6e\x95\xe3\xba\xa6\x80\x22\xc1\xb3\xa1\x22\x26\x1d\x52\x6e\x27\x71\x8e\x91\xf5\x31\x69\xe7\xac\xaf\x11\xc8\x60\xae\x13\x4a\xe3\x51\x76\x1e\x73\xd0\xc9\xb2\x07\x7b\xfc\xf8\x97\x39\xc7\xe4\x75\x43\x3f\x8d\x54\xa9\x13\x5e\x4e\x97\x3a\x4c\x9d\x9d\x4d\xd7\xb9\x99\x96\xce\x92\x4f\x45\x49\x21\x4d\x5a\x6a\xf0\xcb\x2f\x27\x6f\xcf\x7e\x3d\x91\x04\x5f\x53\x48\xaf\xe2\x3f\xd6\x89\xe2\x97\x4c\xe5\x99\xad\x6c\xa9\x13\xc5\xc9\x90\xe9\x05\xb5\x1c\x6d\xe2\xb0\xee\x5e\xe3\x33\x2e\x53\x90\xbc\xb6\x5b\xf5\xf6\xec\xd7\xfb\x83\x2e\x68\x7d\x3b\xe6\x79\xb0\x4b\x9d\xe8\x37\x5a\x7f\x63\xe4\xdf\x68\x7d\x10\xf8\x2b\xdb\xf7\xea\xe2\x0c\x71\x98\x00\x72\x6b\x24\x02\xde\x6f\x36\x83\xb7\xf8\x4f\xb6\xfe\x81\x51\x8d\x1e\x63\x84\xed\xf6\xc3\x41\xc3\x2b\x0e\xd0\x29\x51\xd3\x26\x58\x8f\xcd\xd3\xc9\xe4\xd9\xf6\x39\x0c\x2b\x60\x9d\x1b\x0c\x69\xa0\x58\xa3\xf8\x88\x6f\x8b\xd9\x85\xc4\xa3\x47\x98\x05\xd2\x0b\x05\x1c\x29\xf7\xfd\x2e\x89\x1f\x37\xc3\xaf\xed\x87\xbb\x0b\x1f\x32\xea\xf1\x53\x69\xeb\xc8\x8c\x14\x04\xad\xef\xdf\xef\x9d\x46\xe1\x12\x9e\xe1\xc3\x87\xe7\x48\x73\xf2\x88\x8e\xa8\xc5\xd3\xe7\x20\x17\x09\x74\x65\x93\xfc\xa9\xac\x32\xec\xe9\xe8\x24\x02\x35\xbc\xfc\x36\x18\x4b\xe7\x4a\x47\xda\x43\x3b\xa7\x42\x83\x22\x54\x38\x0a\xeb\x23\xf0\x53\x9b\x0d\x79\xb3\xdd\x2a\x35\x46\x19\x48\x27\x92\xe8\x49\x5b\x4f\x01\x6d\x76\x4e\x7a\x14\x28\xa9\x66\x61\x6c\x40\xd1\x5e\x3b\xe3\x60\x6b\xeb\xa7\x13\xc3\xe5\x82\xc2\x2d\xa0\xdf\x7c\x39\xed\x2b\x9a\xfc\x19\xd9\xef\x03\x7e\xf2\x86\x82\x5d\x92\x99\xbc\xe6\x66\x66\x3d\x99\x77\x8d\xae\xe9\x3c\x3b\x77\xd9\x45\xdd\x81\xe0\x00\xdc\xce\xa3\x88\xf7\xa5\x82\x69\x60\x4e\x53\x29\xe9\x8f\xb3\x37\x67\xa7\x30\xe4\x28\x91\x8c\x0a\x15\x3b\xc7\x2b\xf1\x54\x07\xce\x6d\x5f\xb3\x74\x59\x57\x89\x02\x6c\x82\x8d\x98\xe9\x05\x19\x58\x9f\x18\x9c\x03\xfe\xf3\x2f\x58\xc9\x2b\xaa\xee\x8c\x36\x06\x45\x85\xa1\x6c\x65\x2b\xfc\x80\x3a\xd0\x5e\x67\x76\x69\x50\x2a\xa7\x55\x4c\x7a\xd6\xc3\x44\x01\x71\x1d\x13\x35\x65\x72\x88\x89\xdb\xc1\x47\xd1\x4d\x33\xb7\x93\x64\x1b\x0a\x0f\x5a\x45\x0a\x4b\x5b\xd2\x7d\x76\x7b\xef\x9b\x45\x15\x27\x57\x55\x94\x74\xa7\x86\x96\x53\x63\xe3\x62\xaa\x3f\xe5\x40\xd3\x40\x91\x73\x28\xa9\x68\x75\x48\x4f\x15\x40\xe5\x9c\x71\x72\xdc\x0c\x07\x35\x42\xdc\xa3\x0e\xed\xc7\xcc\x49\x03\x4f\xf0\xe4\x04\x2f\x5f\x5e\x97\x2e\x69\x70\xf6\xe9\xf6\x49\x05\x04\x8a\x89\x03\x95\xec\x51\x5c\x1c\xbc\xdf\x6c\x0a\xd9\x3b\xfa\x88\xc9\x05\x3b\x12\xd2\xaa\x82\x96\xad\x57\x40\x0f\x36\x09\x32\x00\xcc\x68\x6a\xd8\xdf\x02\x98\x02\x46\x8e\xeb\xc2\x08\xc8\xc2\xe8\x14\xa3\x3f\x39\x07\xaf\x9d\x19\x3d\x96\x77\xc6\x46\x3d\x73\x54\x38\xaa\x75\xb9\x2e\x02\xd5\x36\xa6\xb0\x1e\x9d\x22\x85\x4c\xaa\x87\x9a\xe4\x41\xde\xf4\x71\xf7\x3b\xae\x43\x3a\x6c\xf9\xdd\x06\xb7\x26\x5c\x59\xa5\x86\xfe\x75\x2b\x26\x9b\x30\xb0\x5d\xb7\x00\x71\xf2\x3b\x1b\xea\xf8\xed\x65\x37\x10\x2f\x56\x8f\x04\xd0\xab\x8e\x6e\xc4\xb5\x60\x58\xe8\x04\x2b\xad\x6b\xf2\x09\xda\x1b\x78\x4a\x2b\x0e\x0b\xe4\x64\x9d\x4d\x96\x22\x6a\xee\x68\x2d\x31\x82\x2e\xbb\xe5\x36\x56\x00\x3f\x51\x63\x69\xef\xee\x70\xc8\x3e\x62\x46\x15\x07\x82\xf1\x51\xb6\x60\xe1\x79\xe5\x91\xb8\xdb\x9b\x3e\x12\x75\x9d\xc8\x2d\x56\x36\xcd\x21\x54\xb8\x46\xec\x16\x53\xad\xe6\xd6\x51\xc7\x92\x5f\xb8\x0a\x85\xf9\x09\x2f\x5e\x60\x34\xea\x98\xd2\xf0\x35\x4f\x3e\xc8\x8b\xd2\x17\xa9\xf0\xb0\x35\x97\xbd\x15\xb6\xdb\xe3\x02\x73\xbc\xc1\xd7\x5e\xbe\x47\x43\xbe\x3a\xc6\xb7\x4a\xc9\xdf\x9e\xdc\xa7\x25\x63\x35\xc6\xef\x67\x7f\xbc\x3d\xc5\x3b\x8f\x2a\xa7\x1c\xe8\x31\x1a\x5e\x0a\xbd\x69\xe9\x42\xc5\xa1\x19\x84\x23\xa7\x68\x0d\x81\x2b\x90\x5f\xda\xc0\xbe\x91\x51\x2f\x75\xb0\x02\xfb\xa8\xc6\x2a\x52\xc2\xcf\x57\x8a\xae\x5a\x0e\x09\x97\xaf\x2e\xff\x7d\xf1\xee\xc5\xc9\x5e\x29\xff\xe5\xb0\xa0\x30\x54\xd2\xbf\xc7\x76\x7b\xd2\x1d\x2c\xae\x76\x28\x0e\xd9\xa3\x28\xda\x60\x97\xd6\x51\x4d\x06\x45\x21\xba\x54\xec\xe0\x28\x88\x40\xb1\xc4\xf4\x74\x2a\x3f\x4f\x3f\xa1\xa0\x21\xda\xf1\xbe\x0d\x13\x50\xd9\x4b\xc0\xfe\x84\x52\xbd\x5a\x17\xa5\x2e\x52\xc8\x31\xc9\x5e\x44\x4a\xdd\x46\xe4\x16\x35\x79\x5a\xea\x6e\x96\xf2\x24\x26\x5d\x2e\xa0\x23\x22\x0b\xcd\xc7\x0e\xce\x37\x15\xd6\x46\x38\x6d\x8d\x34\x0c\xb3\xb5\x1a\x77\x26\x43\xe8\x6b\x39\x7c\xdc\x9f\x74\x1c\x29\x20\xcd\x6d\xb7\x24\xc3\x7a\xdc\x63\xdc\x70\x20\x35\x96\x54\x22\xaa\xc0\xcd\x0d\xdb\x36\x70\x49\x31\xca\x56\xad\xac\x08\xed\xdc\xb6\xe2\xaf\xcf\x5f\xf5\x69\x44\x42\x9c\x73\x76\xa6\xeb\x31\xfb\x92\xa0\x61\xf4\x1a\xec\xdd\x5a\xaa\x69\xbb\x64\x48\xa0\x18\xd5\x34\xc7\x30\x75\x5c\x6a\xd7\xdd\xb8\xf5\xa7\x48\xa5\x19\x8a\x15\xc1\x9c\xe9\x48\xce\x7a\xd9\x4c\x9c\x3f\x7d\xf3\xa0\x7d\xe4\x2a\xad\x74\xf8\x6a\xfb\xd2\xe9\x46\x2f\x77\xd6\x6a\x0c\xf2\x82\xb4\x8e\x9a\x7a\x3e\xbc\x39\x95\x81\x37\xa3\xba\xa6\xcd\xec\x1b\x1d\x17\x68\x4c\x34\x3b\x5a\x45\x1f\xe7\xe6\xdf\x86\xfd\xf5\x93\xca\x65\xf2\xe9\xcb\xff\x3d\x77\x43\x02\xff\x2f\x77\x7d\x11\xdf\xe7\x4d\x1d\x93\xb6\x94\x3d\x99\x42\x9b\x46\xe0\x51\x09\xa5\x72\x4b\x3e\xce\x6d\x95\x0a\x01\x57\x60\x57\xb4\x4e\x7b\xea\x75\x49\x28\xe2\x81\x53\xc2\x64\xfb\x22\x26\x22\x42\xd0\x2e\x32\x3c\x91\xb9\xb6\x9c\x74\x83\x9d\x2c\xd9\xe5\x86\x22\xe4\xba\xd7\xdf\x08\xcd\x4e\x18\x44\xeb\x7b\x9d\x2f\x45\x0e\x44\x33\x76\xf7\xc2\x06\x4f\xfe\xfe\xec\xc9\x5d\xf7\xc3\x7b\xfc\x4b\x1e\xbd\xfe\x76\xb7\xdf\xb8\x8e\x8e\x6b\x44\x2b\x08\x5f\x11\x1a\xed\x75\x4d\xa0\x25\x85\x75\x9a\x8b\x49\x9a\x07\xce\xf5\x1c\x3b\x09\xdf\x9b\xca\xa0\xe3\x3b\x2f\x77\xce\x8d\xdb\x83\xd7\x6a\x0c\xcf\x89\x4e\xa1\x13\x37\xb6\x2c\x6e\xf6\x0c\x65\x90\x0f\x56\xc7\xdc\x46\x64\x9f\xac\x43\xa3\x63\x77\x49\x8c\xc8\xed\x21\xc6\xee\xf4\xf2\x25\xd8\xf7\x7f\xf3\xc6\x72\x4e\x26\x77\xed\xda\xfb\x1e\x40\xa0\x19\x73\x12\xe2\x28\xb9\x69\xbb\x8b\xee\x5d\x9f\x36\x23\x15\xe7\x39\x19\x91\xf4\xa2\x18\xce\xfc\xfc\x57\xb9\xfe\xbb\x48\xdb\xed\x01\xc6\x8f\x56\x83\xcf\x9f\xfb\x2b\xd2\xee\xeb\xe1\x7f\x01\x00\x00\xff\xff\x38\xb6\xac\x3c\xe3\x0f\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
