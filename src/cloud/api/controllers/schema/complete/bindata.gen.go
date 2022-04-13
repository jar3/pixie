// Code generated for package complete by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package complete

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xdc\x56\x5f\xe1\x36\x1d\x84\x0e\x0a\xa2\x9b\x38\x84\x9a\xd6\x60\x2f\x29\x77\xb1\x58\xc4\x77\x17\x0b\xa5\x15\xb7\x9f\xe4\xfb\xbf\x9f\x9e\x46\x72\x05\x1b\x3a\x82\xc3\x83\xd2\x00\x2f\x07\x80\xc9\xb8\xc6\xca\x72\x31\xa5\x3d\x46\x0a\x70\xb4\xc4\xd2\xf8\x55\x80\xf5\x44\x94\x52\xab\x77\x00\x3d\x25\xae\x87\x52\x7a\x36\x3a\xe9\x9d\xa4\xe0\x39\x2f\x9b\x1b\xd5\x96\x50\xbc\x7b\x3b\x37\xce\xfe\xa8\xc6\x79\x36\x8a\x39\xc0\x79\xfa\xf8\xcb\x3f\x3d\x82\x3d\xa5\xcc\x3a\xeb\x1d\x40\x75\x43\x69\xa8\xd5\x66\x79\x34\x8e\x94\x0d\x63\xb7\xcb\x01\xb6\xad\xa2\x7d\x85\x9f\x00\x00\x00\xff\xff\xef\x77\x02\x34\xfc\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 252, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x1a\xdb\x6e\xe3\xb8\xf5\xdd\x5f\x71\x66\xe7\x61\x13\x20\x0d\x16\x45\x77\x51\xf8\xa9\x5a\x5b\xb3\xa3\x26\x71\x5c\xdb\x99\xed\x62\x31\x18\xd0\xd2\xb1\x45\x44\x26\xb5\x24\xe5\xc4\x2d\xe6\xdf\x8b\x43\x52\x12\x29\x2b\x73\x6b\xd1\x37\x8b\x97\x73\xbf\xd3\xf8\x6c\x50\x14\x60\x4e\x35\xc2\x3f\x1a\x54\x27\xf8\xf7\x04\xa0\xd1\xa8\xa6\xf0\xa0\x51\x65\x62\x27\x5f\x4d\x00\xa4\xda\x4f\xe1\x5e\xed\xdb\x6f\x3a\xb1\x46\x63\xb8\xd8\x6b\x77\xb2\xfd\x6a\x77\x13\x63\x14\xdf\x36\x06\xfd\x7e\xff\xed\xe1\xd1\xa2\x9e\xc2\xef\x1d\x9a\xf7\xb4\x91\x57\x8d\x36\xa8\x2e\x78\x31\x85\x6c\xfe\xea\x72\x0a\x33\xb7\xd2\x62\xf6\x07\x7e\x3e\x2d\xd8\x01\x2f\x04\x3b\xe0\x14\xd6\x46\x71\xb1\x7f\xf9\x30\xa1\x09\x77\x42\x4c\x33\x29\x04\xe6\x86\x4b\x71\x8e\xb3\xdf\xeb\x01\xf2\x44\x19\xbe\x63\xb9\xb9\x60\xfe\xc7\xe6\x54\xe3\x14\x92\xe0\xcb\x82\xb8\xcd\xda\x25\xba\xc8\x1a\x23\x73\x79\xa8\x2b\x34\x78\xc1\x45\xdd\x98\x96\xec\x2b\xc8\x1b\xa5\xa5\x5a\x4a\x3d\x85\x4c\x98\x2b\x60\x16\xe5\x14\x92\xe0\x4e\x62\xd7\x08\xf8\x55\x4b\xf9\x43\x36\x6f\x61\x5c\xc6\x87\x57\xa8\x9b\xea\x0c\xed\x1b\x8e\x55\x31\xc4\xbd\xa3\x45\xcf\x41\x70\x36\x15\x86\x9b\xd3\x0d\x17\xc5\xd5\x04\x00\x40\xe1\x1f\x0d\x57\x58\x24\x6a\x4f\x87\x49\xa0\xe3\xc7\xdf\x7f\x01\x79\x96\x90\x96\xc6\x09\xc0\x6b\x58\xe7\x8a\xd7\xe6\xb0\x57\x80\xa2\xa8\x25\x17\x46\x5f\x81\xc2\x1d\x2a\x30\x12\x0a\x99\x6b\xe0\x02\xf2\x4a\x36\x05\xab\xf9\x75\xad\xa4\x91\x13\x80\x8a\x1f\xf1\x1d\xc7\x27\x22\xe7\xd6\xff\xbe\x43\xc3\x0a\x66\x98\x53\x72\x7b\x62\x26\x85\x41\x61\x74\xa0\xe3\xdb\xc1\x16\x1d\xd7\x96\x0e\x02\xe7\x28\x8a\x81\xb9\xdd\x11\x50\xeb\x68\xc3\xf3\x34\xc7\xba\x92\x27\x78\xc4\x93\x9e\x00\x14\xf6\xeb\x80\xc2\xdc\xe0\x89\x10\xcc\xc3\x85\x18\x4f\x74\x36\x40\x13\x5d\xf1\x58\x92\x65\xd6\xa2\x60\x35\xf7\xb0\x93\x65\x76\x06\xd4\xed\x06\xd0\xdc\x21\x0f\x66\x59\x35\x7b\x2e\x26\x00\xb5\xfd\xa1\x2f\x1e\xb9\x28\xa6\x7e\x99\xf4\x7a\x39\x85\xdf\xdd\x97\x03\xa7\x90\x78\xe5\x52\xb8\x45\xf2\x0f\x0b\xdb\xbb\xe2\x95\x07\xf4\x0e\x95\xb6\xb6\xdc\xbb\x68\x7f\xe1\x95\x45\xbd\x29\xb9\x86\x27\x5e\x55\xb0\x45\x50\x58\x57\x2c\xc7\x02\xb6\xa7\x21\x8a\x99\x14\x3b\xbe\x07\x29\x72\x04\x56\x55\x90\x4b\xa1\x9b\x03\x2a\x0d\x25\x3b\x22\x6c\x11\x05\x34\x75\xc1\x0c\x16\xd7\x2e\xc2\xac\xc6\x00\x84\x54\xf6\x4c\xb9\xbd\x51\xd6\x46\xaf\x8d\x82\x8e\x2e\xaf\x3b\x53\x5a\xc5\x4b\x03\x24\x6e\x71\x00\x7e\x8e\x86\xf1\x0a\x8b\xe1\xd5\xc9\xc7\xc9\x24\x0c\xd9\x77\x8d\x61\xb4\x6d\xa3\xf6\x4c\x21\x33\xe8\x43\x57\x14\x0a\xe1\x6f\x05\xd6\x0a\x73\x92\xcd\x85\x42\xa6\x49\x23\xdf\xf9\x03\x1a\x98\x42\x10\xf2\x09\x72\x0b\xa0\x80\x23\x67\x50\x3f\x7b\x33\xfc\xee\x72\x02\xf0\x60\xe5\xfa\x8e\xff\x8b\xdb\xa0\x48\xf2\xf0\x5e\x4e\x4e\x9e\xcd\x5f\x5d\xc1\x31\xd8\x9c\x42\x5a\x70\xc3\xb6\x55\x74\x65\x24\x3e\x3b\x92\x23\xbb\x3e\x33\x73\x80\x39\x52\xd4\x98\xbf\xe0\x15\x3f\x4b\x59\x21\x13\x3d\x38\x67\xd8\xbd\x81\xb7\x00\xdc\xf7\xf8\x4d\xc7\x60\x98\xc7\x2e\x74\x97\xde\x5a\x66\xa2\x34\x77\x79\x9e\xf6\xd6\x68\xe2\x4c\x77\xc1\x82\x24\x18\x42\x09\x92\xe1\xe5\x58\x7a\xcc\xc4\x91\x3b\x72\x2e\xf0\xc0\x78\x15\xf8\xd5\x8e\x2b\x6d\x16\x61\xda\xbb\x82\x8a\x0d\x96\x2e\xdb\xec\x4d\x60\x62\xfe\x96\xa8\x0e\x5c\x93\x53\xea\x0b\xca\xd3\x9d\x02\x9b\x78\x33\x26\x38\xd8\xe8\x81\x87\x3a\xbc\x57\xfb\x0b\xa9\xf6\x43\x2a\xb2\x79\x8f\xfd\x5e\xed\x3b\xe1\x4a\xb5\xef\x10\xcb\x7e\xbd\x47\x1a\x1c\x26\x38\x41\xe9\xe1\xf0\x39\xd6\x36\xf2\x11\x45\x00\xec\xb2\xc3\x3d\x01\x58\xe1\x51\x3e\x62\x52\x55\xc1\x59\x1d\x1f\x0e\x2c\x60\x85\x07\x79\xb4\xbc\xbe\x51\xf2\x40\xec\x04\xd2\x09\x8f\xc6\xf1\xd2\xb1\xf6\xd9\x48\x73\x05\x28\x88\xad\xa2\x03\xd4\xad\x0c\x42\xe4\x15\x05\xb6\x1d\x0f\x65\x11\x02\xd5\xa3\x86\xbb\x1a\x89\x26\x56\xb6\x2e\x65\xf5\xa0\x06\x07\xc7\xdc\x67\x08\xeb\xf3\x20\x9c\x92\x9d\x93\xbd\x40\x49\x88\xe8\xe3\x64\x62\x83\x57\x6b\x44\x36\x78\xf9\x73\x13\x80\xa8\xa4\x9b\x00\xc4\x0e\x40\x29\x8a\xe7\xa6\x51\xd1\x99\xa1\xe5\xb9\xa5\xbe\x00\xa1\x05\xae\x93\xba\x56\xf2\x18\xe8\xa0\xa7\x25\x9b\xa7\x4b\x66\x4a\x4b\x4a\x36\x4f\x87\xc0\x6a\x66\xca\xfe\xbb\xbd\xe4\x8d\xf2\x33\xf4\x17\xf2\xc0\xb8\x18\x42\x74\xca\x77\x14\xb1\x4a\x47\x7a\xe0\x05\x12\x31\x94\x3b\x3c\x5d\x94\x33\x42\xb1\xb5\xae\x61\x51\x33\xc1\xaa\x93\xe1\xb9\xbe\xaf\x8d\xa4\xda\x2e\x02\xe5\xc8\x0a\x2f\xf7\xa1\xc6\x5e\x37\xb2\x51\x6b\x44\xf1\xd2\x3d\x5b\x30\xbe\x10\xbd\xc6\x01\x8c\xdf\xfa\x22\x9a\x3b\x42\xe3\x12\x66\x20\x62\x9f\xa8\x12\x73\xa7\xa7\xf0\xa6\x92\xcc\xb8\xb2\x49\xe7\xe7\x4a\x72\x80\x06\x00\x1e\x29\x3f\xf4\xca\xf8\x1a\x78\xa3\x75\xdb\x7f\x41\x5f\x04\xef\x7f\x42\x26\x8a\xe6\x30\x52\xcc\xaf\x0d\x33\x68\x11\x24\xe9\xfa\xc3\xc3\xe2\x66\x71\xff\xeb\xc2\x7f\x2d\xd3\xc5\x3c\x5b\xfc\xe2\xbf\x56\x0f\x8b\x45\xff\xf5\x26\xc9\x6e\xd3\xb9\xff\xd8\xa4\xab\xbb\x6c\x91\x6c\xd2\xf9\x28\xa6\xbe\x4b\x71\x88\x92\x4d\x80\xe8\x35\x24\x02\xb0\xe0\xc6\x37\x38\x20\x73\xea\x7c\x80\xef\x80\xd9\xec\x03\x25\xd3\x70\x90\x05\xdf\x71\x2c\xc0\x94\x08\xce\x8a\x0c\x3e\x1b\x2a\x08\xb9\xd0\xa8\xc8\x86\x40\x2a\x28\x28\xdc\xd0\xef\xbc\x64\x8a\xe5\x54\xc8\x5c\xf7\xf5\x24\x17\x79\xd5\x14\xa8\xa9\x4c\xb2\x17\x84\x85\xf7\x88\xa7\xad\x64\xaa\x00\x26\x0a\xa8\x99\x76\x00\xe4\xe1\xc0\x44\x61\xaf\x13\xc5\xe9\x3c\xdb\x38\x72\x41\x63\x85\x79\x4f\xaf\xa8\x4e\xe3\x44\xe7\xa5\xd4\x28\x80\x89\xa8\xe1\x02\xdd\xec\xf7\xa8\xe9\xee\x75\x4b\x56\xc1\xa9\x0a\xd3\x60\xfb\x97\xd7\x96\xa8\xe8\x8a\x29\x99\x01\x6e\x40\x97\xb2\xa9\x0a\xa0\x9c\x64\x0f\x11\xaa\xef\xb5\x6f\x15\xa9\x29\xa2\x45\x41\x82\x61\x14\x43\x6a\xc5\x49\xbb\x86\x6d\x5b\x2e\xd6\xe9\x6d\x3a\xdb\x7c\xc2\x1e\xa8\xaa\xf7\xe6\x70\x13\x99\xc3\xcd\x87\xe5\xfd\xdc\xff\x5a\xbf\x9b\xb5\xbf\x66\xab\x6c\xb9\xf1\x1f\x8b\xe4\x2e\x5d\x2f\x93\x59\xda\x7e\xdf\xcf\xd3\xde\xe3\x02\x54\xeb\x4e\x02\x16\x95\xeb\x2a\xc6\x69\x19\x84\x4e\x6f\xd9\x94\x44\x82\xec\x38\x01\x38\x30\x93\x97\x58\x64\xa2\xc0\x67\xdb\x88\x66\xc2\xbc\xa7\xee\x8c\xec\x7b\x0c\xb8\x35\xfc\x8e\xba\x0d\xdb\x0e\x88\x22\x93\x21\x53\x2b\xf0\x19\xe4\xce\x0a\xd6\xb0\xad\xd3\x84\x29\x51\x87\x7a\x74\xc5\xf2\x4e\x2a\x12\xb3\x61\x5b\x4b\x85\x6d\xdb\x2d\xa0\x5f\x4b\x34\x25\x2a\x6f\x37\x64\x5c\x2c\xb8\x4c\xf7\xc0\x90\x1d\x10\x7c\x87\xd0\x36\x3e\x07\xf6\xe8\xb4\xec\x4d\x11\xf0\x19\xf3\xc6\x46\x4e\xc2\xd3\x7f\x25\x3b\x43\x81\x94\x80\xf7\x21\x13\x42\xfa\x06\x8d\x79\xcf\xea\xfb\x51\xfd\xb8\x2e\x3c\x10\xc3\x4e\xaa\x03\x33\xd4\x05\x38\xdf\x23\x62\x3b\x47\xd4\xbe\x42\x79\x2a\x79\x5e\x5a\xc3\xb7\x5d\x57\xcd\x94\x76\x2d\xdb\xb9\x39\xcb\xce\xe6\x9d\xbd\xb3\xed\xda\xc8\x1a\x6a\xa9\xb9\xa5\x97\xf8\xeb\x70\x66\xe1\x6c\x22\x12\xe8\x90\x06\xa2\x8b\xc1\x91\x55\xbc\xb8\x0a\xe4\xd3\x0a\xf0\xda\xe6\xfb\xb4\x5b\x0f\x85\xf5\x1a\x92\xaa\x8a\x54\x4a\x6a\x41\x96\x97\x81\xf6\x89\x48\xed\x75\xbc\x8e\xa4\x1b\xd9\xcf\xb8\x50\x83\xf9\x46\x20\xd9\x17\x22\x83\xf6\x56\xd1\xf2\x47\x05\x01\x2f\xb0\xf8\x52\xb5\xbe\x8a\xe4\x24\x15\x08\x69\xcd\x96\x3a\xcc\x46\x09\x2c\x40\x59\x4a\x9c\xe5\xd6\x4c\x19\xce\x2a\xb8\x30\xaa\xc1\x4b\x3a\xde\x91\x74\xb1\x63\x95\x46\xea\xf6\x4a\xa6\x93\xa2\xb0\xfa\x61\xd5\x9d\x75\x37\x3d\x52\x33\xcd\xa4\x30\x8c\x0b\x54\xe4\x60\x8d\xcb\xeb\xc3\xe2\x67\x3c\x65\x79\x57\xed\x8f\x1d\x50\x6b\xb6\x8f\x96\xda\x36\x35\x5c\xd1\x86\x29\x33\x93\x8d\x30\xd6\xe5\x7a\x52\x6e\xfe\xaa\xd3\x23\x0a\x27\xee\x11\x60\xb6\x69\xda\xf0\x03\x46\x64\x50\xdb\x34\x58\x6c\x01\x2e\x65\xf1\x4d\x5c\x35\xfa\xab\xd9\xca\x5b\x31\xda\x29\x65\x2c\x53\x37\x2b\x40\x62\x8d\x76\x5b\x36\xdb\x11\xc2\x98\x3c\x6c\xb4\xf7\x6d\x76\xc0\x82\xb3\xc1\x02\x77\x8c\xac\xd2\x2a\x80\x72\x98\x90\xa6\xf4\xee\xf4\x28\xe4\x93\x20\x93\x9f\xad\xa3\xa4\x4d\xf7\xfc\x79\x0d\x25\xb2\xca\x94\x27\xba\x5a\x22\x53\x66\x8b\xcc\x5b\x96\xc2\x1c\xf9\x11\x0b\x4a\xb5\x0a\xf7\x4d\xc5\x14\x70\x61\x50\x51\x79\x6b\xf3\xad\x29\x5d\x0c\xf0\x93\x03\x02\xa7\x50\xd7\x52\x14\x44\x81\x91\x76\xc8\x88\xda\x68\x4f\xc4\xdb\x34\xb9\xdd\xbc\xfd\xed\x9c\x88\x46\x04\x64\xd8\xb0\xd9\x43\xcc\xdd\xc8\x96\xea\x07\x09\x4b\xfe\xcc\x11\x66\x95\x6c\x5c\xc6\xe7\xda\xbb\x57\x1b\x5e\x7a\x1e\xae\x60\x6b\xa3\x9d\xf8\xde\xc0\x1f\x0d\xaa\x93\x0d\x27\xe4\x9a\x5a\x1e\xd0\xab\xcd\x67\x71\x85\x1a\x0f\xdb\x0a\x35\xbc\xdd\x6c\x96\xdf\x6b\xf8\xf1\x87\x1f\xbc\xf6\x3b\xf9\x8d\x13\x6f\xa3\xfd\x5e\xda\xa1\x26\xd7\x3d\xad\x9e\x8f\x5f\x56\xcb\x59\xcb\x01\xe5\x8b\xad\x42\xf6\xa8\xaf\x2d\x80\x52\xd6\xe8\xa2\x31\x33\x5d\xe9\xd0\x32\x6e\xe1\xe6\x44\xe8\x96\xe5\x8f\x54\xa8\x70\x81\x96\x65\x72\xfe\x03\xc5\x16\xf0\x14\x39\x4a\x3c\x9d\xf3\x6c\x3d\xbb\x5f\x2c\xd2\xd9\xc6\x56\x78\x43\x39\x53\x6f\x49\xba\x79\x2a\x51\x0c\x05\xcd\xdd\x4a\xad\x64\x8e\x5a\x53\xe8\x6c\x8f\xb7\x32\x58\xce\x93\x8d\x2b\x23\x1d\x5c\x37\x21\x72\xf5\x52\xcb\xb9\x13\x3b\x2d\x51\xd8\xd2\xe4\xc2\x4c\x9c\x40\xda\x60\xb6\x6b\x94\xcb\xa6\xce\x8c\xdd\x70\x4f\x03\xdb\xca\xc6\x89\xe0\xc9\x47\x3d\x6e\x42\xdb\x94\x6a\x48\xca\x39\x8f\x9e\x96\x27\xa6\xc1\xa8\x93\xb7\x3f\x87\xc0\x91\xb4\xb3\xf3\xb7\xd6\x6a\x84\x7c\x22\x86\x19\x6c\x59\x11\x09\xd0\x32\x99\xf6\x35\xf2\x40\x82\x05\xee\x15\x2b\x7a\x05\x07\xf2\xab\xf8\x23\x56\x27\x42\xbb\xc5\xc0\xe2\x08\xf7\x81\xef\x4b\x43\xcb\x76\xe4\xe2\x4d\x95\xda\x8c\x56\x6b\xe9\x2f\xab\x64\xee\x4a\x70\x1b\xad\xc2\x29\x9b\xf5\xf6\x9a\x69\x6d\x4a\x25\x9b\x7d\x99\x0e\xe6\x0e\x41\xfc\x0e\x06\x84\x71\xdf\xd1\x46\xb2\x28\x8c\xb4\x11\xf3\x6d\xeb\x33\x51\xf0\x8b\xc7\x7f\xd1\xd8\xaf\xdb\x1d\x0e\x83\xfb\x47\x98\x97\x77\xce\x3a\x71\x85\xc6\x9c\x66\xe3\x9b\xe7\x2f\x0f\x6d\x80\x55\xb2\x5a\x56\x4c\x60\x17\xd7\x6d\xe5\xd8\x7d\xb9\x80\xda\xc5\x95\x39\x33\xec\xf3\xc7\x45\x73\x58\xc8\x02\xb5\x8f\xbd\x76\x21\x13\xda\xa8\x86\xba\x39\x2c\xe2\x4d\x27\xd3\xbb\xf3\x8c\x50\x2b\x3c\x72\xd9\xe8\xf5\x98\xd0\xcf\xf6\xa3\x7c\x35\x54\x65\xfc\x5e\xe5\x94\x5a\x27\x45\xa1\x50\x47\x79\xc9\xc8\x47\x14\xe7\xad\x68\x3f\x32\xb4\x57\xcf\x06\x2f\xdc\xee\xdd\x72\xf1\x18\xdd\x7d\x0d\xab\xcf\xbc\xd4\x58\xe8\xc3\x07\x9a\xcf\x8d\x4d\x86\x6d\xed\x57\xa2\x69\x5f\x63\x7c\x49\xe0\x70\x4e\xcf\xa8\xb0\x1a\x78\xae\xda\xd3\x21\x05\x47\xae\xff\xbe\xbe\x5f\x7c\x0b\x11\xf1\xeb\xd1\x57\x71\x6a\xcb\xaf\x96\xca\xd8\x6b\xbf\x0a\xf9\x0b\xfc\x0f\xde\xb5\xbc\x7b\xc4\xac\x77\x3d\x63\xf0\xa4\x69\xc1\x00\x44\x0d\xbd\xfd\xbc\xcd\x16\x0f\xff\xfc\x90\xdc\xcd\x7f\xfa\x4b\xbb\x34\x4f\x56\xbf\x66\x8b\x78\x6d\x76\xbf\xd8\x24\xd9\x22\x5d\x7d\x58\xa7\x9b\x0f\xbf\x25\x77\xb7\xeb\xf1\xad\x11\x78\xf1\x81\x4d\x7a\xb7\xbc\xa5\xa0\xeb\x80\x74\x2e\xd0\xbf\xb7\xba\x37\x6c\x15\xd9\xae\x2e\xd9\x9f\x7f\xfc\x29\xe2\x31\x1e\x51\x7d\x4d\x0c\x1d\x1f\x70\x05\x33\x72\xa7\xf1\xf3\x81\xe3\xf9\xc5\x60\xce\xed\x9c\xee\x85\xb9\xa0\xd3\xbf\x9b\x04\xff\x49\x61\x65\x1f\x6a\x88\x71\x7d\xdd\x96\xac\x76\x6f\xb4\x5e\x0d\x06\xd1\xe3\x6d\xb5\x8d\xed\x72\x2f\x83\xde\x8b\x30\x68\x33\x12\x98\x75\x53\xd7\x52\x19\xdd\x4d\x7a\xa3\xc1\x61\xf7\xae\x75\x9e\x77\xe0\x85\x89\x77\x67\x6f\xfd\x6b\xa3\xe5\x62\x19\x8e\x25\x96\x37\x1f\x56\xe9\x26\x5d\x6c\xb2\xfb\x45\x5f\xa5\x87\x0f\x83\x63\x8c\x1f\x59\xd5\xe0\x79\xa0\xeb\x9f\x20\xed\xad\x6e\xe0\x1e\x3d\x06\xae\xf3\x12\x0f\xed\xe3\x69\x55\xc9\xa7\x59\xa3\x8d\x3c\xa4\xcf\xc4\xfd\xc3\xea\x36\xe2\xcc\x1e\xc8\x84\xc6\xbc\x51\xb8\xb9\x5d\x8f\xa4\xdb\x73\xd8\xa3\x24\x8f\x6a\xe7\xcc\x6c\xbe\x81\xf3\x97\x01\xe8\x81\x14\xc6\xce\xf8\xff\x4c\x0c\x45\xd0\x99\x0b\x1f\xe1\xbd\x63\x7d\xf4\x79\xe4\x13\xa2\xff\x56\x64\xaf\x81\x32\x77\x6f\x83\xb1\x83\x0c\x9e\x26\xbe\x20\x2a\x8f\x78\xc9\xce\x36\x29\x22\x3f\xad\xbb\xcc\x3e\x7c\xd9\x19\xfc\xf1\x24\x9b\x3b\x76\xdc\x83\xf8\xf0\x4d\x62\x49\xcd\x84\x19\x31\x97\x17\x9e\x80\xff\xdf\x54\x9f\x67\x87\x71\x4e\x3e\xa1\xac\x51\x1e\x63\x6b\x1c\x63\xf2\x8b\x86\x81\x03\xc6\x46\xf8\x3a\x67\x6b\x84\xab\x11\xa6\x3e\xc1\xd3\xc7\xc9\x7f\x02\x00\x00\xff\xff\xb4\xb8\x24\x41\x37\x25\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 9527, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
