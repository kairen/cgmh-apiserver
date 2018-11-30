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

var _apiSwaggerSpecV1Yml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5b\x73\x1a\x3b\xf2\x7f\xe7\x53\x74\x39\xff\x2a\xfe\x5b\xb5\x86\x19\xb0\x73\xe1\x29\x5e\xdb\xf1\x21\xeb\xd8\xde\xd8\xce\x6e\xd5\xd6\x56\x4a\x33\xd3\x80\xe2\x19\x69\x8e\xa4\x81\x43\x4e\xed\x77\xdf\x92\xe6\x0e\x03\x0c\xd8\x31\xbe\xe0\x87\x04\x46\xea\x96\xba\xfb\xd7\xad\x6e\x49\x83\x9c\x90\xe1\x10\x45\x0f\x9a\x9d\x96\xd5\x6c\x50\x36\xe0\xbd\x06\x80\x87\xd2\x15\x34\x54\x94\xb3\x1e\x1c\x9f\x7d\xf9\x0d\x3e\x71\x11\x48\x38\xba\xea\x37\x00\xc6\x28\xa4\x69\xb1\x5b\x56\xcb\x6a\x00\x28\xaa\x7c\xac\xe8\xa8\x50\x04\xf2\x72\x70\x8d\x62\x4c\x5d\xec\xc1\x48\xa9\xb0\xd7\x6e\xfb\xdc\x25\xfe\x88\x4b\xd5\x00\x70\x39\x53\xc4\x55\x7a\x50\x00\x0c\x08\xf5\x7b\x70\x37\xf5\xb1\xe5\x7c\xa4\x6c\x42\x99\x54\xc4\xbd\x6b\xb9\x3c\x68\x00\xf8\xd4\x45\x26\x31\xee\xcb\x48\x80\x3d\x38\x0a\x89\x3b\x42\xe8\x98\x69\x00\x44\xc2\xcf\x46\x99\x4c\x26\x2d\x62\x9a\x5b\x5c\x0c\xdb\x09\xb1\x6c\x9f\xf7\x8f\x4f\x2f\xae\x4f\xf7\x3b\x2d\xab\x35\x52\x81\xdf\xd0\x33\xe9\x41\x3e\x29\x45\x86\xb2\xd7\xd8\x4f\x47\x88\xd4\x68\x56\x23\xb7\x12\x85\x69\x40\xa6\xa8\x4b\xf4\x43\x23\x71\x4a\xa3\xdb\xab\x69\x0a\x9d\xce\x71\x8c\x7e\x65\x2f\xd3\x52\xea\x7b\xc5\x29\x53\xf0\x1b\x95\x4a\xf0\x69\x25\x4d\xde\x83\x8b\x69\x89\x56\x1b\xa4\x92\x44\x37\x98\x9e\x6f\xf4\xbf\x70\xed\x8e\x30\x40\xd9\x90\xf1\xff\x5a\x03\x5a\x93\x8d\x90\xa8\x91\xd4\x3a\xdf\x6b\x93\x48\x8d\xda\x3e\x1f\x52\xb6\x17\x1b\x21\xd4\xaa\x33\x9f\x00\x62\xb5\xc5\x9f\xf7\x53\xb5\xe9\x3f\x19\x05\x01\x11\xd3\x1e\x34\xcd\xb0\x86\xbe\x99\xb4\xb9\x9c\xc9\xc8\x8c\x96\x12\x92\x30\xf4\x13\x9d\xb6\x7f\x48\xce\x92\x86\x50\x70\x2f\x72\xeb\x74\x24\x82\x04\xa8\x50\x14\xba\x52\xd6\x03\x87\x7b\xd3\xe4\x41\x0a\x9e\xd2\x23\x81\xbf\x47\x54\xa0\xd7\x03\x25\x22\xcc\x1e\x1b\x6d\x90\x5e\xf6\x1d\x60\xef\xff\x04\x0e\xf6\x7a\xb0\xf7\xa6\xed\xe1\x80\x32\xaa\x67\x20\x13\xb5\x34\x52\x66\x32\xe4\x1a\x6d\x39\x61\xb3\x63\x59\xcd\x22\x9f\x92\x41\x9a\xd7\x91\xeb\xa2\x94\xcd\x42\x87\xf9\xa1\x01\xd4\x34\xc4\x1e\x70\xe7\x07\xba\xaa\xd4\x10\x0a\x1e\xa2\x50\xb4\x38\x64\xfc\x17\xa0\x94\x64\x88\xb3\x8f\x53\x5e\x52\x09\xca\x86\x73\x8d\xf8\x07\x09\x42\xed\xd4\xcd\xcb\xbb\xe6\x4c\xab\xcb\xbd\x85\xec\x28\x53\x38\x34\xe8\x5f\xc0\xaf\x63\x59\x33\x8d\x1e\x51\x64\x11\xbb\x0a\x49\x97\x4b\x6b\x48\xf9\x1d\xb2\xaa\x86\x15\x32\x97\xe5\xc6\xe9\xe7\x91\x73\xe6\xd2\x4b\xfa\xb9\x7f\xfb\xb3\x6f\x5f\xd0\xbe\xec\xb3\xaf\x87\xee\x71\xff\x6d\xff\x2e\xfc\xd7\xb7\xe3\xcf\x1f\x5a\x38\xfd\xec\x3b\xff\xfc\x14\x3a\xb9\x8e\x9a\x07\x4b\x0d\xdd\x67\x63\xe2\x53\x0f\x4e\x75\xac\x03\x2e\xe0\x8a\x48\x39\xe1\xc2\x2b\x71\xe8\x2e\xe1\x60\xbc\x88\x71\x05\xc4\x55\x74\x4c\x14\x16\x49\x0f\x57\x0c\xae\x50\x30\xe2\x83\x0e\xc9\x28\xe0\x54\x08\x2e\x9a\xb9\x6f\x0b\x1c\x52\xa9\x50\x6c\xe6\xde\x5f\x13\x6a\x60\x38\x81\x48\xa2\x78\x79\x5e\x9e\x29\x68\xe7\xe8\xeb\x38\x7a\x4d\x9f\xb8\x22\x53\x9f\x93\x87\x84\xb3\x44\x95\x62\x39\x5a\x0f\xca\x12\x15\xa8\x11\x1a\x1c\x43\x38\xe3\xa3\x2f\x09\xd0\x5a\x45\x2f\x0f\xcd\xfc\x6e\xcb\x58\xc6\x34\xbe\x73\xdf\x9b\xc3\xcf\xbd\x91\x3d\xe0\xc2\xc5\xb5\xe1\x8d\x6e\x24\xa8\x9a\xe6\xc3\xee\xc3\xdf\x90\x08\x14\xba\x53\x0f\xfe\x7d\xe4\x05\x94\xfd\x67\xce\x19\x3e\xe9\xc1\x40\xbc\x16\x97\x28\xe8\x76\xe7\x17\x5b\x4f\xe6\x52\x9c\x3d\x44\x3e\xb7\x61\xaa\xf6\xeb\x5c\x39\xa4\xed\xb1\xdd\xd6\xfe\x94\xb8\xf1\x10\x17\xb8\x71\x52\x52\x42\xc9\x31\xcf\xa9\x54\x40\x7c\xdf\x78\x64\x06\xb6\xf5\xdc\x7c\x6d\x6c\x27\xd0\xde\x21\xbb\xc0\x8e\x08\x41\xa6\x73\x6d\x54\x61\x50\x89\xe9\x05\x81\xc7\xe0\xa0\x08\x4e\x7b\x59\x15\xc0\xf4\x42\xc0\x05\xfd\xf9\x10\xf9\xff\x92\x05\xa4\x12\x79\xb7\xa1\x47\x14\x96\x92\xfc\xaa\x70\x1d\x47\x67\xad\xb9\x6c\x4a\x73\x11\xfc\xbe\xe1\x5a\xcf\xe1\xbb\xae\x54\x36\x8e\xd6\x97\x7f\x7f\x72\x60\x8e\x8b\xbc\x11\x91\xe0\x20\x32\x88\x8c\xba\xbd\xad\x14\xe2\x75\xc1\xba\x2a\x92\x52\xbf\x2a\xc1\x7f\x5c\x90\x7b\xe8\xa3\xc2\x35\x70\x7e\x62\x08\x4a\x38\x5f\x19\x5e\x49\x31\xbc\xfe\x72\xaf\xd8\x08\x9d\x51\x44\x2b\x16\xd4\x7a\xd0\x8c\x2c\xcb\xb2\xec\x97\xe4\x6b\x7c\xab\x85\xef\xf6\xfd\xa2\x94\x88\xb4\xff\xd4\xe0\xf8\xef\x26\xf9\xc8\x19\xaa\xb8\x36\x70\xa6\x06\x61\xbb\x0c\xe3\x31\xe3\xef\x96\xf0\x22\x15\x51\x91\x5c\x55\x85\x56\x06\xd7\xeb\x14\x2f\x31\x8f\xff\x37\xfb\x99\xd8\x76\x7c\xee\xde\xfd\xe5\xe1\xc2\xed\x63\xd4\x8c\x26\x09\x49\x54\xf1\x00\xa1\xb1\x66\x61\x72\x42\x14\x79\x0a\x31\x43\x70\x1f\x37\x42\x40\x21\x8d\x04\xcd\xe4\x19\x1a\xdd\xc8\xfe\xfa\x4c\xee\xe3\x18\xfd\x7b\xdb\xdc\x70\x79\x86\x46\x8f\xa5\x7f\x7d\x56\x0f\x39\x65\x2b\x77\x1c\x57\x5a\xdd\x70\x79\x86\x56\x8f\xa5\x7f\x5d\x56\x2f\xfa\xf9\xc2\x64\x30\xbd\xca\x00\xd5\xbb\x53\x86\x47\xfd\xed\x29\xb2\xdb\x9e\x7a\x92\xdb\x53\xa5\xa0\xf7\xf8\xfb\x53\x0b\xcf\xa2\xab\xd1\x77\x2c\x50\x07\x1c\x86\x93\x18\x7f\x46\x41\xcf\x2b\xe8\xbc\xa2\x55\x66\xc9\x72\x52\x6d\xde\x64\x3d\xd9\x99\xf6\x19\x98\x76\xd9\xae\x5b\xb5\x75\x93\x6d\xb7\x9d\x75\x1f\xcf\xba\x07\xd6\xc1\x12\xa2\x0b\xae\xe0\x13\x8f\xd8\x83\xa7\x16\xed\x3f\x6b\x6c\x36\x55\x82\xe4\x0c\x55\x02\x10\x67\x0a\xfd\x93\x5d\xbe\xf0\x0b\x37\x9b\xb6\xb9\xf4\x67\x78\x31\xe9\xf7\x28\xbe\xe1\xba\x0a\x31\xb3\x17\x66\xa1\x3a\x33\x35\x2c\x21\xe1\xb9\xec\x20\x4b\x47\x8d\xdf\x23\x14\xb3\x61\x43\xd7\x05\xb7\xb7\x19\xfa\x16\x18\xb7\x2c\xe9\x3f\x34\x1f\xb3\x47\xaa\x8b\xa1\x28\xa2\x99\x92\x76\xe0\xad\xe4\xb6\x8d\x64\xd7\x20\xe3\x7b\x8a\xb6\x2d\x23\x7f\xc0\x45\xb0\x0a\xf1\xc9\x35\x6f\xa8\x06\xba\xe6\x20\x77\x00\xcf\xff\x5e\x3d\xc0\x0d\xa6\xb6\x96\xec\x2f\x2c\xe6\x2a\x61\x9c\xd4\x72\xc4\xc0\x78\xdb\xf7\x0d\x8a\x8a\x7b\x09\xc7\x9f\x4f\x2a\xd1\x98\x43\xe5\xd3\x3f\x3d\x5d\x52\xba\x56\x82\x39\xa9\x5c\x77\x60\xde\x81\xf9\x29\x82\x79\x59\xb1\x5e\x89\xe7\xa4\x56\x2f\xe3\xf9\x25\x5c\x92\xd9\x5d\x91\x79\x12\xe9\xce\x8c\x9c\xcf\xc2\x8d\x4a\x99\x7b\xad\x2d\x8e\x2a\xd7\x3a\x43\x65\x9c\x6a\xb7\xc1\xf1\xf8\xa1\x7a\x4b\x58\xa9\x77\x9b\x66\x59\x5e\x61\x00\x13\xb3\x79\x5e\x9b\xa6\x7a\xe2\xaf\xe8\x0a\xcd\x1b\xf8\xc2\x3d\xf4\xe1\x24\x57\x41\xa3\xa0\x0e\xcd\xcf\xbc\x52\x1c\x33\x9e\x73\xcf\x4c\xdb\x8d\xd8\x42\xe6\x1d\x85\xe4\x73\xfa\x8a\x42\x0c\xa2\x39\xff\x8d\xdf\xb2\xcf\xe6\x5b\xe9\xad\xf9\x52\x26\x51\x7c\x54\x28\x55\xf2\x0e\x3e\x54\xbe\x92\xb1\x9c\x47\xf8\x51\xca\x89\x65\xe6\x63\x5e\xab\xd9\x40\x24\xee\x7b\x57\x45\xa9\xf6\x81\xe1\xe4\xea\x17\xcb\x59\x18\x74\x03\x51\xf5\x5f\x61\x92\x75\x39\x58\x21\x32\xa9\x3e\xba\x3a\x3a\xe6\x6f\x22\xad\xad\xb2\x07\x56\x47\xfa\xde\xeb\x7d\xd1\x18\xe7\x75\xc9\x47\x32\x44\xe6\x4e\x93\x2f\x11\xa3\x2a\xa5\xe0\xd2\x38\x41\xfa\x75\xc4\x19\x3e\x84\x50\x5a\x9e\x07\xc4\x32\xa4\x71\xb0\x26\xe9\x67\x3e\x4a\x5f\x6d\x8b\x25\xaf\x4b\x78\x7c\xf6\xe5\xb7\xe4\xa9\xd6\x52\x46\x56\x45\x95\x11\x1d\xf5\x41\x21\x49\xc5\xfc\xc1\x9d\x1b\xf3\x23\x1d\xb5\x68\x91\x0d\x29\x63\xd9\x72\x6b\x0c\x50\x1b\xbf\x1f\xec\x4e\xf7\xe0\xf0\xed\xbb\xf7\x3a\x82\xe1\x18\xfd\x9a\x98\x29\x00\x63\x18\x46\x57\x82\xba\xe9\x57\x8f\x4c\xf3\xaf\xf3\x18\xa0\xb5\x8d\xe7\x8f\x75\x62\xde\xcd\x1a\x4a\x81\xfa\x28\x52\x1c\x86\xc8\x90\x08\xbd\x8e\x3a\x53\x90\x26\x60\x6f\x60\xe9\xe6\xf9\xe9\xb7\xd3\x73\xe8\xa6\x6b\x44\x2a\xcf\x2c\xfd\x6c\x56\x53\x91\xcd\xa4\xb2\xd7\x26\xb5\x73\xd2\x82\x74\x75\x27\x7e\xc1\x45\x40\xfc\xec\x7d\x07\xfd\xff\x02\xf3\xcd\xdb\xa1\xfc\x52\x41\xe5\x40\x7a\x8d\x27\xaa\x57\xbc\x1d\xbe\xa8\x6c\x5a\xcf\x3a\xf7\x8e\x03\x3b\x47\x9e\xa5\x2d\x39\xb2\xfe\x33\xce\xdc\x3f\x59\xcf\xd9\x52\x73\x9a\xf3\x94\xda\x20\x3e\xb4\x32\x14\x0b\xee\xd7\x9e\xb2\x79\xb1\x31\xb5\x8c\xb9\x4c\x3e\x4b\xe9\x70\xee\x57\x00\x22\x4f\x58\xcd\xe5\xf3\x7a\x54\x03\xe2\x4b\x4c\x9c\xe4\x7b\xbe\xa3\xbd\xf3\x94\x57\xed\x29\x31\x1a\x4a\x59\x45\x1d\x44\x24\xc7\x6c\xdb\x42\x85\x44\x57\xe0\x9c\x7f\x2e\x12\xd8\x75\x06\x07\x76\xf7\xe0\x9d\xe3\xd8\x1f\xde\xbd\x1f\xbc\x1d\x74\x3b\xd6\xfb\x77\x4e\xc7\x1d\x58\x36\x76\x0f\xed\x54\x0d\x71\x35\xf7\x6c\x94\xf0\xc8\x21\x23\x8f\x6c\xcf\x40\x37\x9b\x84\xe1\xfc\xaa\xfc\x43\x8b\xf9\x10\x12\x6d\xbc\x9c\xe5\x77\xc1\x6b\x66\xb6\xa5\x33\xf4\x7d\x30\xbb\x2d\x85\xef\x63\xe2\x47\x8b\x72\xdb\x8d\xf4\xd1\x49\xf1\x9c\x0e\xb4\x91\x36\xcd\xb4\xd6\x4c\x3b\x4b\xb7\x26\x6a\x1b\x7d\xe5\x62\xb8\xc2\xac\x9e\xf3\x3c\xb5\x65\x3a\xd2\x35\xca\x8a\x8e\x65\xbf\xdf\xb7\x3e\xec\x77\x0e\x6f\x6c\xbb\xd7\xed\xf6\x0e\xde\xb7\x2c\xcb\x6a\x56\x2b\xaa\x7f\x7d\x79\x42\x14\xc6\xfb\x08\xc1\x66\x58\x4d\x7e\x9e\xf1\x22\xaf\xcd\x92\x27\xa7\x85\x0a\x3f\x79\x74\x95\xd5\xe9\xfb\xda\xc8\x7a\x88\x02\x1d\x15\x4e\xc6\x75\x10\x31\x8f\xb2\xe1\x35\x8f\x44\x56\xe4\x29\x41\x28\xa3\x6c\x68\x76\xc5\x16\xb7\x9c\x50\x79\x77\x4d\x7f\xce\xb6\x9d\xb2\x31\x15\x9c\x05\xc8\xd2\x4d\x04\x16\x05\x0e\x8a\xcb\xc1\xd9\xd5\x6d\xf2\xc4\x27\x6c\x18\x91\xe1\xcc\x14\xaf\xdd\x11\x7a\x91\x9f\x3e\xc5\x3f\x42\x74\xd5\x0d\xcd\xe6\x2d\x90\xc8\x6c\x2b\xc2\x1d\x11\x31\xbc\x7f\x1d\xea\x12\xbb\xd3\xb5\x0f\x6c\x8f\xb8\xb5\x63\xd7\xba\x20\xaf\xb1\x12\xa4\xc8\x2f\x18\xb9\x36\x14\x75\x66\xd7\x2c\xd3\x9f\xae\xbb\xc1\x64\xcf\xa6\x96\x45\x20\x6d\x5c\xa6\x14\xa0\x57\x5b\x1a\xf3\xcb\xa5\x57\x31\x61\x2a\x95\x01\x6c\x5d\x0e\x54\x38\x85\x60\x50\xc2\xf7\x4a\x16\x41\x48\xb2\x49\xa4\x83\x07\x1a\xeb\x47\xc2\x1d\x51\x85\xae\x8a\xc4\x1a\xa2\x5c\x5c\x94\x98\x9c\x93\x69\x5a\xc3\xc3\xea\xc0\x94\xee\x8d\x70\x59\x7b\xc0\x5b\x27\x62\x2a\x02\xfb\x6d\xcb\x3a\x28\x6c\x78\x18\x6f\xad\xbf\xef\xf0\xad\x7f\xd2\x3f\x82\x1b\x94\x3e\x81\x6f\x76\x1e\xd3\x30\xe4\xee\xe8\x66\x9d\x20\x69\x83\x47\xa6\xd0\x85\x11\x8f\x44\x89\x8b\x0e\x1b\xab\x15\x11\x90\xb8\x02\x4d\x93\x3b\xa2\xea\x52\xce\xa8\xb0\x22\xa0\xd5\x96\xe1\x66\x84\xc0\x07\x03\xea\x12\x1f\x42\x32\x44\xf0\xf8\x84\x15\xcf\x33\x2b\x63\xe2\x26\xcb\x4f\x89\x4f\xe9\xc4\x6c\x15\x8f\x8e\x35\xc7\xa5\x10\x85\xeb\x8b\x8a\x4c\x72\x31\xf0\xf9\x64\x56\xb6\x1b\x4d\x58\x97\xcd\x75\x14\xa2\x18\x53\x99\x1f\xcc\x78\x44\x91\xb5\x58\xf4\x03\x32\xcc\xde\x40\x2e\x2c\x1f\x35\x30\x13\x7b\x70\x1a\x4f\xd3\x85\xa6\xf6\xc8\x57\x53\x35\xe2\x59\x38\xcd\x97\xa0\x59\x06\x33\xc7\xb3\xd5\x47\xb3\x52\x11\xa1\xf4\xc2\x5f\x7c\xb8\xe4\x50\x36\x5d\x28\xbc\x38\x57\xc8\xff\xe6\xf3\x8e\xdc\xc1\x4d\x3b\xf3\x7e\xcd\x30\xf6\xdb\x66\x39\x96\xa7\x6b\xf4\x93\xd3\x87\xfd\x28\xfa\xe8\x66\xd1\x30\xce\x3f\x36\x52\x43\x8c\xaa\xaa\xc9\x55\x9d\xad\xe7\x6e\x7e\x58\x3a\x58\x27\xae\x8a\x88\xbf\x3e\x97\xc2\x7e\x9e\xc9\xa7\x6a\xbb\xc6\xd7\x88\x41\x30\xcd\x62\x42\xbc\xa6\x65\x07\xdb\xaa\xa4\xef\xe5\xac\xbe\x62\x49\x4b\xa5\x74\x6b\x2f\x6e\xfc\x2b\xf4\x59\xbc\x09\x00\x84\x79\x70\x64\x3e\xee\x65\x5e\x2d\x55\x7c\xc6\xbe\xd6\x72\xb4\x69\xce\x6e\xcc\x2d\xd0\xfc\x2a\xe2\xa3\x0c\x58\x38\x7f\xaf\x59\x2b\x68\x8a\x2c\xa7\x37\xb6\x58\x90\x19\xc7\x1d\xef\x91\x1d\x3f\x92\xa1\xdf\xc0\x31\x0f\x42\xce\x90\x29\xd9\x70\xb3\x8f\x7a\xe0\xf4\x0e\xc5\x75\xfa\xeb\xe8\x50\xbc\x42\xd1\x28\xce\xcb\xfc\x6a\x7a\x32\x6f\xd3\xbb\x07\x8e\xe9\xda\x68\xfc\x2f\x00\x00\xff\xff\x79\x7a\xc2\xae\x6f\x5f\x00\x00")

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

	info := bindataFileInfo{name: "api/swagger-spec/v1.yml", size: 24431, mode: os.FileMode(420), modTime: time.Unix(1543561338, 0)}
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

