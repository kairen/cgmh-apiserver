// Code generated by go-bindata.
// sources:
// api/swagger-spec/v1.yml
// DO NOT EDIT!

package swagger

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiSwaggerSpecV1Yml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5b\x73\xe2\x3a\xf2\x7f\xe7\x53\x74\x65\xfe\x55\xfc\xb7\x6a\x03\x36\x24\x73\xe1\x69\xb2\x49\x26\x87\xd9\x4c\x92\x9d\x24\xb3\x5b\xb5\xb5\x35\x25\xdb\x0d\x68\x62\x4b\x3e\x92\x0c\x87\x39\xb5\xdf\x7d\x4b\xf2\x1d\x0c\x18\x92\x09\xb9\xc0\x43\x02\x96\xbb\xa5\xee\xfe\x75\xab\xd5\xb2\x2c\x27\x64\x38\x44\xd1\x83\x66\xa7\x65\x35\x1b\x94\x0d\x78\xaf\x01\xe0\xa1\x74\x05\x0d\x15\xe5\xac\x07\xc7\x67\x5f\x7e\x83\x4f\x5c\x04\x12\x8e\xae\xfa\x0d\x80\x31\x0a\x69\x5a\xec\x96\xd5\xb2\x1a\x00\x8a\x2a\x1f\x2b\x6e\x54\x28\x02\x79\x39\xb8\x46\x31\xa6\x2e\xf6\x60\xa4\x54\xd8\x6b\xb7\x7d\xee\x12\x7f\xc4\xa5\x6a\x00\xb8\x9c\x29\xe2\x2a\xdd\x29\x00\x06\x84\xfa\x3d\xb8\x9b\xfa\xd8\x72\x3e\x52\x36\xa1\x4c\x2a\xe2\xde\xb5\x5c\x1e\x34\x00\x7c\xea\x22\x93\x18\xdf\xcb\x48\x80\x3d\x38\x0a\x89\x3b\x42\xe8\x98\x61\x00\x44\xc2\xcf\x7a\x99\x4c\x26\x2d\x62\x9a\x5b\x5c\x0c\xdb\x09\xb1\x6c\x9f\xf7\x8f\x4f\x2f\xae\x4f\xf7\x3b\x2d\xab\x35\x52\x81\xdf\xd0\x23\xe9\x41\x3e\x28\x45\x86\xb2\xd7\xd8\x4f\x7b\x88\xd4\x68\x56\x23\xb7\x12\x85\x69\x40\xa6\xa8\x4b\xf4\x45\x23\x71\x4a\xa3\xdb\xab\x69\x0a\x37\x9d\xe3\x18\xfd\xca\xbb\x4c\x4b\xe9\xde\x2b\x4e\x99\x82\xdf\xa8\x54\x82\x4f\x2b\x69\xf2\x3b\xb8\x98\x96\x68\xb5\x41\x2a\x49\x74\x83\xb9\xf3\x8d\xfe\x0b\xd7\xee\x08\x03\x94\x0d\x19\xff\xd7\x1a\xd0\x9a\x6c\x84\x44\x8d\xa4\xd6\xf9\x5e\x9b\x44\x6a\xd4\xf6\xf9\x90\xb2\xbd\xd8\x08\xa1\x56\x9d\xf9\x06\x10\xab\x2d\xfe\xbe\x9f\xaa\x4d\x7f\x64\x14\x04\x44\x4c\x7b\xd0\x34\xdd\x1a\xfa\x66\xd2\xe6\x72\x26\x23\xd3\x5b\x4a\x48\xc2\xd0\x4f\x74\xda\xfe\x21\x39\x4b\x1a\x42\xc1\xbd\xc8\xad\x73\x23\x11\x24\x40\x85\xa2\x70\x2b\x65\x3d\x70\xb8\x37\x4d\x2e\xa4\xe0\x29\x5d\x12\xf8\x7b\x44\x05\x7a\x3d\x50\x22\xc2\xec\xb2\xd1\x06\xe9\x65\xbf\x01\xf6\xfe\x4f\xe0\x60\xaf\x07\x7b\x6f\xda\x1e\x0e\x28\xa3\x7a\x04\x32\x51\x4b\x23\x65\x26\x43\xae\xd1\x96\x13\x36\x3b\x96\xd5\x2c\xf2\x29\x19\xa4\x79\x1d\xb9\x2e\x4a\xd9\x2c\xdc\x30\xdf\x35\x80\x9a\x86\xd8\x03\xee\xfc\x40\x57\x95\x1a\x42\xc1\x43\x14\x8a\x16\xbb\x8c\x3f\x01\x4a\x49\x86\x38\x7b\x39\xe5\x25\x95\xa0\x6c\x38\xd7\x88\x7f\x90\x20\xd4\x4e\xdd\xbc\xbc\x6b\xce\xb4\xba\xdc\x5b\xc8\x8e\x32\x85\x43\x83\xfe\x05\xfc\x3a\x96\x35\xd3\xe8\x11\x45\x16\xb1\xab\x90\x74\xb9\xb4\x86\x94\xdf\x21\xab\x6a\x58\x21\x73\x59\x6e\x9c\x7e\x1e\x39\x67\x2e\xbd\xa4\x9f\xfb\xb7\x3f\xfb\xf6\x05\xed\xcb\x3e\xfb\x7a\xe8\x1e\xf7\xdf\xf6\xef\xc2\x7f\x7d\x3b\xfe\xfc\xa1\x85\xd3\xcf\xbe\xf3\xcf\x4f\xa1\x93\xeb\xa8\x79\xb0\xd4\xd0\x7d\x36\x26\x3e\xf5\xe0\x54\xc7\x3a\xe0\x02\xae\x88\x94\x13\x2e\xbc\x12\x87\xee\x12\x0e\xc6\x8b\x18\x57\x40\x5c\x45\xc7\x44\x61\x91\xf4\x70\x45\xe7\x0a\x05\x23\x3e\xe8\x90\x8c\x02\x4e\x85\xe0\xa2\x99\xfb\xb6\xc0\x21\x95\x0a\xc5\x66\xee\xfd\x35\xa1\x06\x86\x13\x88\x24\x8a\x97\xe7\xe5\x99\x82\x76\x8e\xbe\x8e\xa3\xd7\xf4\x89\x2b\x32\xf5\x39\x79\x48\x38\x4b\x54\x29\x96\xa3\xf5\xa0\x2c\x51\x81\x1a\xa1\xc1\x31\x84\x33\x3e\xfa\x92\x00\xad\x55\xf4\xf2\xd0\xcc\xef\xb6\x8c\x65\x4c\xe3\x3b\xf7\xbd\x39\xfc\xdc\x1b\xd9\x03\x2e\x5c\x5c\x1b\xde\xe8\x46\x82\xaa\x69\xde\xed\x3e\xfc\x0d\x89\x40\xa1\x6f\xea\xc1\xbf\x8f\xbc\x80\xb2\xff\xcc\x39\xc3\x27\xdd\x19\x88\xd7\xe2\x12\x05\xdd\xee\xfc\x62\xeb\xc9\x5c\x8a\xb3\x87\xc8\xe7\x36\x4c\xd5\x7e\x9d\x2b\x87\xb4\x3d\xb6\xdb\xda\x9f\x12\x37\x1e\xe2\x02\x37\x4e\x96\x94\x50\x72\xcc\x73\x2a\x15\x10\xdf\x37\x1e\x99\x81\x6d\x3d\x37\x5f\x1b\xdb\x09\xb4\x77\xc8\x2e\xb0\x23\x42\x90\xe9\x5c\x1b\x55\x18\x54\x62\x7a\x41\xe0\x31\x38\x28\x82\xd3\x5e\xb6\x0a\x60\x7a\x22\xe0\x82\xfe\x7c\x88\xfc\x7f\xc9\x04\x52\x89\xbc\xdb\xd0\x23\x0a\x4b\x49\x7e\x55\xb8\x8e\xa3\xb3\xd6\x5c\x36\xa4\xb9\x08\x7e\xdf\x70\xad\xc7\xf0\x5d\xaf\x54\x36\x8e\xd6\x97\x7f\x7f\x72\x60\x8e\x17\x79\x23\x22\xc1\x41\x64\x10\x19\x75\x7b\x5b\x59\x88\xd7\x05\xeb\xaa\x48\x4a\xfd\xaa\x04\xff\x71\x41\xee\xa1\x8f\x0a\xd7\xc0\xf9\x89\x21\x28\xe1\x7c\x65\x78\x25\xc5\xf0\xfa\xcb\xbd\x62\x23\x74\x46\x11\xad\x98\x50\xeb\x41\x33\xb2\x2c\xcb\xb2\x5f\x92\xaf\xf1\xad\x2e\x7c\xb7\xef\x17\xa5\x44\xa4\xfd\xa7\x06\xc7\x7f\x37\xc9\x47\xce\x50\xc5\x6b\x03\x67\x6a\x10\xb6\xcb\x30\x1e\x33\xfe\x6e\x09\x2f\x52\x11\x15\xc9\x55\xab\xd0\xca\xe0\x7a\x9d\xe2\x25\xe6\xf1\xff\xa6\x9e\x89\x6d\xc7\xe7\xee\xdd\x5f\x1e\x2e\xdc\x3e\xc6\x9a\xd1\x24\x21\x89\x2a\x1e\x20\x34\xd6\x5c\x98\x9c\x10\x45\x9e\x42\xcc\x10\xdc\xc7\x8d\x10\x50\x48\x23\x41\x33\x79\x86\x46\x37\xb2\xbf\x3e\x93\xfb\x38\x46\xff\xde\x36\x37\x5c\x9e\xa1\xd1\x63\xe9\x5f\x9f\xd5\x43\x4e\xd9\xca\x8a\xe3\x4a\xab\x1b\x2e\xcf\xd0\xea\xb1\xf4\xaf\xcb\xea\x45\x3f\x5f\x98\x0c\xa6\x8f\x32\x40\x75\x75\xca\xf0\xa8\x5f\x9e\x22\xbb\xf2\xd4\x93\x2c\x4f\x95\x82\xde\xe3\xd7\xa7\x16\xee\x45\x57\xa3\xef\x58\xa0\x0e\x38\x0c\x27\x31\xfe\x8c\x82\x9e\x57\xd0\x79\x45\xb3\xcc\x92\xe9\xa4\xda\xbc\xc9\x7c\xb2\x33\xed\x33\x30\xed\xb2\xaa\x5b\xb5\x75\x93\xb2\xdb\xce\xba\x8f\x67\xdd\x03\xeb\x60\x09\xd1\x05\x57\xf0\x89\x47\xec\xc1\x53\x8b\xf6\x9f\x35\x8a\x4d\x95\x20\x39\x43\x95\x00\xc4\x99\x42\xff\x64\x97\x2f\xfc\xc2\x62\xd3\x36\xa7\xfe\x32\x5e\x3c\x1c\x90\xc8\x5f\xb9\x06\xa9\x33\x69\xc4\x9c\x5e\x52\x59\x7f\xf3\xa2\xbe\x3f\xb6\x2c\xcb\xea\x3c\xaf\x90\x75\x5f\x40\x99\xf5\xdc\x28\x7e\x64\x7a\x55\x08\x9a\x7d\x02\x1b\xaa\x97\x3a\x86\x25\x24\x3c\x97\xed\x8c\x6a\x64\xfc\x1e\xa1\x98\x9d\x87\xf4\x42\xf3\xf6\x36\x0b\x67\x0b\xec\x57\x96\xf4\x1f\x9a\x8f\x29\xba\xeb\xd5\x75\x14\xd1\x4c\x49\xbb\x68\x58\xc9\x6d\x1b\xab\x27\x83\x8c\xef\x29\xda\xb6\x8c\xfc\x01\x17\xc1\x2a\xc4\x27\xe7\x06\xa0\x1a\xe8\x9a\x83\xdc\x01\x3c\xff\xbc\x7a\x80\x1b\x4c\x6d\x6b\x89\xb1\xb8\x3a\x50\x09\xe3\xa4\x38\x40\x0c\x8c\xb7\xfd\x00\x4b\x51\x71\x2f\x61\x3f\xfd\x49\x65\xae\x73\xa8\x7c\xfa\xdb\xf1\x4b\xd2\xda\x4a\x30\x27\x59\xed\x0e\xcc\x3b\x30\x3f\x45\x30\x2f\xab\xfe\x54\xe2\x39\x29\xfe\x94\xf1\xfc\xba\x97\x67\x2f\xee\x99\xab\x27\xf4\x18\xfa\xb3\x70\xa3\x52\xe6\x5e\xab\x66\x56\xe5\x5a\x67\xa8\x8c\x53\xed\x2a\x66\x8f\x1f\xaa\xb7\x84\x95\x7a\x8f\x67\x2d\xcb\x2b\x0c\x60\x62\x36\xcf\xab\x0a\xaf\x07\xfe\x8a\x9e\xc9\x7a\x03\x5f\xb8\x87\x3e\x9c\xe4\x2a\x68\x14\xd4\xa1\xf9\x99\x33\xea\x31\xe3\x39\xf7\xcc\xb4\xdd\x88\x2d\x64\x0e\xbd\x24\xdf\xd3\x33\x2f\x31\x88\xe6\xfc\x37\x7e\x6d\x43\x36\xde\x4a\x6f\xcd\xa7\x32\x89\xe2\xa3\x42\xa9\x92\x97\x3a\x40\xe5\x19\x9f\xe5\x3c\xc2\x8f\x52\x4e\x2c\x33\x1e\x73\x4e\x6b\x03\x91\xb8\xef\x5d\x15\xa5\xda\x07\x86\x93\xab\x5f\x2c\x67\xa1\xd3\x0d\x44\xd5\x9f\xc2\x20\xeb\x72\xb0\x42\x64\x52\x7d\x74\x75\x74\xcc\x8f\xb6\xad\xad\xb2\x07\x56\x47\x7a\x90\xfa\xbe\x68\x8c\xf3\xba\xe4\x2b\x19\x22\x73\xa7\xc9\x8f\x88\x51\x95\x52\x70\x69\x9c\x20\xfd\x39\xe2\x0c\x1f\x42\x28\x2d\xcf\x03\x62\x19\xd2\x38\x58\x93\xf4\x33\x1f\xa5\x67\x25\x63\xc9\xeb\x12\x1e\x9f\x7d\xf9\x2d\xb9\xaa\xb5\x94\x91\x55\x51\x65\x44\x47\x7d\x50\x48\x52\x31\x7f\x70\xe7\xc6\xbc\xf5\xa5\x16\x2d\xb2\x21\x65\x2c\x9b\x6e\x8d\x01\x6a\xe3\xf7\x83\xdd\xe9\x1e\x1c\xbe\x7d\xf7\x5e\x47\x30\x1c\xa3\x5f\x13\x33\x05\x60\x0c\xc3\xe8\x4a\x50\x37\xfd\xe9\x91\x69\xfe\x73\x1e\x03\xb4\xb6\xf1\xe2\x6d\x93\x6e\xd6\x50\x0a\xd4\x47\x91\xe2\x30\x44\x86\x44\xe8\x79\xd4\x99\x82\x34\x01\x7b\x03\x4b\x37\xcf\x4f\xbf\x9d\x9e\x43\x37\x9d\x23\x52\x79\x66\xe9\x67\xb3\x9a\x8a\x6c\x26\x95\xbd\x36\xa9\x9d\x93\x16\xa4\xab\x3b\xf0\x0b\x2e\x02\xe2\x97\x0e\xd0\x24\x1b\x6f\xb3\x2c\x1c\xce\xfd\x79\x06\x03\xe2\x4b\x6d\x27\xcd\x60\x81\xdd\xe7\x0d\x58\x3e\xde\x52\x39\x42\x9d\x1c\x10\xd5\x2b\x9e\x53\x58\xb4\xde\x5a\xcf\xac\xf7\x0e\x20\xbb\x08\x30\x4b\x5b\x8a\x00\xfa\x53\x88\x02\xab\xa9\xf7\xe2\xd7\x29\xd9\x7b\x45\xe2\xfe\xc9\x7a\x2e\x9e\x62\xc1\xec\xe2\xd4\x76\x9d\x43\x2b\xf3\x1d\xc1\xfd\xda\xf2\x9a\xf3\xb9\xa9\x59\xcd\x99\x88\x7a\xbe\x52\x48\x93\xcd\x19\x8a\xf5\x3d\xec\x7b\x5e\x47\xdf\xb9\xd9\xab\x76\xb3\x18\x0d\xa5\x5c\xa6\x0e\x22\x92\xcd\xbd\x6d\xa1\x42\xa2\x2b\x70\xce\x3f\x17\x09\xec\x3a\x83\x03\xbb\x7b\xf0\xce\x71\xec\x0f\xef\xde\x0f\xde\x0e\xba\x1d\xeb\xfd\x3b\xa7\xe3\x0e\x2c\x1b\xbb\x87\x76\xaa\x86\x78\x0d\xf9\x6c\x94\xf0\xc8\x21\x23\x8f\x6c\xcf\x40\x37\x9b\x84\xe1\xfc\xc4\xc7\x43\x8b\xf9\x10\x12\x6d\x3c\x9d\xe5\x47\x1a\x6a\xe6\xd3\xa5\x9d\xfb\x7d\x30\x35\x9e\xc2\xef\x31\xf1\xa3\x45\x19\xf5\x46\xfa\x48\x9f\x46\xca\x3a\xda\x48\x9b\x66\x58\x6b\x26\xbb\xa5\x67\x35\x6a\x1b\x7d\xe5\x64\xb8\xc2\xac\x9e\xf3\x3c\xb5\x65\x6e\xa4\x6b\x2c\x66\x3a\x96\xfd\x7e\xdf\xfa\xb0\xdf\x39\xbc\xb1\xed\x5e\xb7\xdb\x3b\x78\xdf\xb2\x2c\xab\x59\xad\xa8\xfe\xf5\xe5\x09\x51\x18\x57\x2f\x82\xcd\xb0\x9a\xbc\x65\xf4\x22\x5f\x11\x26\x57\x4e\x0b\x75\x85\xe4\xd2\x55\x56\x1d\xd8\xd7\x46\xd6\x5d\x14\xe8\xa8\x70\x32\xae\x83\x88\x79\x94\x0d\xaf\x79\x24\xb2\xa5\xa5\x12\x84\x32\xca\x86\xa6\x16\xb7\xb8\xe5\x84\xca\xbb\x6b\xfa\x73\xb6\xed\x94\x8d\xa9\xe0\x2c\x40\x96\x96\x2e\x58\x14\x38\x28\x2e\x07\x67\x57\xb7\xc9\x15\x9f\xb0\x61\x44\x86\x33\x43\xbc\x76\x47\xe8\x45\x7e\x7a\x15\xff\x08\xd1\x55\x37\x34\x1b\xb7\x40\x22\xb3\x02\x88\x3b\x22\x62\x78\xff\xd5\xaf\x4b\xec\x4e\xd7\x3e\xb0\x3d\xe2\xd6\x8e\x5d\xeb\x82\xbc\xc6\x4c\x90\x22\xbf\x60\xe4\xda\x50\xd4\x99\x5d\xb3\x4c\x7f\xba\x6e\x59\xcb\x9e\x4d\x2d\x8b\x40\xda\x78\x8d\x53\x80\x5e\x6d\x69\xcc\x0b\x78\xaf\x62\xc2\x54\x2a\x03\xd8\xba\x1c\xa8\x70\x0a\xc1\xa0\x84\xef\x95\x2c\x82\x90\x64\x83\x48\x3b\x0f\x34\xd6\x8f\x84\x3b\xa2\x0a\x5d\x15\x89\x35\x44\xb9\xb8\x28\x31\x39\x27\xd3\xb4\x00\x00\xab\x03\x53\x5a\x91\xe1\xb2\x76\x87\xb7\x4e\xc4\x54\x04\xf6\xdb\x96\x75\x50\x28\xb3\x18\x6f\xad\x5f\xed\xf8\xd6\x3f\xe9\x1f\xc1\x0d\x4a\x9f\xc0\x37\x3b\x8f\x69\x18\x72\x77\x74\xb3\x4e\x90\xb4\xc1\x23\x53\xe8\xc2\x88\x47\xa2\xc4\x45\x87\x8d\xd5\x8a\x08\x48\xbc\x02\x4d\x93\x3b\xa2\xea\x52\xce\xa8\xb0\x22\xa0\xd5\x96\xe1\x66\x84\xc0\x07\x03\xea\x12\x1f\x42\x32\x44\xf0\xf8\x84\x15\x77\x51\x2b\x63\xe2\x26\xd3\x4f\x89\x4f\x69\x9f\x6e\x15\x8f\x8e\x35\xc7\xa5\x10\x85\xeb\x8b\x8a\x4c\x72\x31\xf0\xf9\x64\x56\xb6\x1b\x4d\x58\x97\xcd\x75\x14\xa2\x18\x53\x99\x6f\x07\x79\x44\x91\xb5\x58\xf4\x03\x32\xcc\x0e\xd2\x17\xa6\x8f\x1a\x98\x89\x3d\x38\x8d\xa7\xe9\x44\x53\xbb\xe7\xab\xa9\x1a\xf1\x2c\x9c\xe6\x53\xd0\x2c\x83\x99\x4d\xe1\xea\x0d\x61\xa9\x88\x50\x7a\xe2\x2f\x5e\x5c\xb2\x15\x9c\x4e\x14\x5e\x9c\x2b\xe4\x9f\xf9\xbc\x23\x77\x70\xd3\xce\xbc\x5f\xd3\x8d\xfd\xb6\x59\x8e\xe5\xe9\x1c\xfd\xe4\xf4\x61\x3f\x8a\x3e\xba\x59\x34\x8c\xf3\x8f\x8d\xd4\x10\xa3\xaa\x6a\x70\x55\x3b\xfa\xb9\x9b\x1f\x96\xb6\xf3\x89\xab\x22\xe2\xaf\xcf\xa5\x50\xcf\x33\xf9\x54\x6d\xd7\xf8\x1a\x31\x08\xa6\x59\x4c\x88\xe7\xb4\x6c\x3b\x5d\x95\xf4\xbd\x9c\xd5\x57\x2c\x69\xa9\x94\x6e\xed\xc5\x8d\x7f\x85\x3e\x8b\x8b\x00\x40\x98\x07\x47\xe6\x6b\x56\xfe\x24\x52\xc5\x3b\xfb\x6b\x4d\x47\x9b\xe6\xec\xc6\xdc\x02\xcd\xcb\x3d\x1f\xa5\xc3\xc2\xae\x7f\xcd\xb5\x82\xa6\xc8\x72\x7a\x63\x8b\x05\x99\x71\x7c\xe3\x3d\xb2\xe3\x47\x32\xf4\x1b\x38\xe6\x41\xc8\x19\x32\x25\x1b\x6e\xf6\x55\x77\x9c\x3e\xb9\x71\x9d\xbe\xe4\x1f\x8a\x0f\x6e\x34\x8a\xe3\x32\x2f\xff\x4f\xc6\x6d\xee\xee\x81\x63\x6e\x6d\x34\xfe\x17\x00\x00\xff\xff\x68\x85\x27\xb7\x36\x62\x00\x00")

func apiSwaggerSpecV1YmlBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerSpecV1Yml,
		"api/swagger-spec/v1.yml",
	)
}

func apiSwaggerSpecV1Yml() (*asset, error) {
	bytes, err := apiSwaggerSpecV1YmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/swagger-spec/v1.yml", size: 25142, mode: os.FileMode(420), modTime: time.Unix(1544173894, 0)}
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
	"api/swagger-spec/v1.yml": apiSwaggerSpecV1Yml,
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
	"api": &bintree{nil, map[string]*bintree{
		"swagger-spec": &bintree{nil, map[string]*bintree{
			"v1.yml": &bintree{apiSwaggerSpecV1Yml, map[string]*bintree{}},
		}},
	}},
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

