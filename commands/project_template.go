// Code generated by go-bindata.
// sources:
// project_template/contracts/token.sol
// project_template/contracts/token_test.go
// project_template/libraries/token.sol
// DO NOT EDIT!

package commands

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

var _contractsTokenSol = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\xcb\x4e\x2b\x31\x0c\x86\xf7\x7d\x0a\x2f\xa7\x3a\x73\xba\x40\x82\xcd\x50\x1e\x81\x0d\xec\x50\x55\xa5\x19\xa7\xb2\xc8\x38\x23\xc7\xa9\x8a\x50\xdf\x9d\xa4\x73\xa1\x53\xb6\xce\xe7\xff\xe2\xd8\xc0\x2a\xc6\x2a\xbc\x87\x4f\x64\xf8\x5e\x01\x44\x15\xe2\x23\xf4\xe9\xe0\xc9\x02\x9b\x0e\x9b\x3c\xc5\x13\x72\xa6\xc4\x70\x74\x28\x95\x69\x5b\xc1\x18\x81\xb8\xc5\x33\xb6\xb0\x77\x12\xba\x1a\xfe\x8c\x35\xd4\x90\x88\xf5\xe1\xf1\x09\xf6\x27\xe3\x13\xae\x8b\x5a\x67\xfa\xbe\x98\xcc\x3a\xdb\x97\x2b\xb6\x9e\x6c\x0f\xc6\x1b\xb6\x18\x9b\x55\xa6\x5d\x62\xab\x14\x78\x08\x59\x4d\x7a\xc4\xa4\x64\xfc\x5b\xea\x7b\xff\x55\x4f\xb9\xb5\x30\xaf\x39\xf5\xfa\xda\x06\x66\xa9\x0f\x3d\x6f\x82\xd0\x91\x78\x07\xdb\xe5\x72\x73\x05\x4b\xd5\xfc\x32\x0b\x94\xe9\x65\xe1\xaf\xf7\xf5\x05\x2d\xd2\x09\x65\x28\x09\xa6\x0b\xa9\x94\x10\xd4\x24\x1c\xab\x43\x08\x1e\x62\x72\x8e\x2c\x61\x79\x18\x12\x91\x83\x6a\x4e\xd5\xc5\xe3\x26\x62\xbe\x97\xec\xe0\xf9\x4e\x01\x9c\xf1\x11\x9b\x65\x8d\xdb\x85\xff\xdb\x71\xe3\x8e\x99\x72\xed\xe0\xdf\x92\x98\x3f\xf0\x57\xa5\xbe\x69\x31\xda\x0f\xec\x98\x41\x25\x8d\xa7\xb8\xac\x7e\x02\x00\x00\xff\xff\xb5\x22\x41\x04\x2f\x02\x00\x00")

func contractsTokenSolBytes() ([]byte, error) {
	return bindataRead(
		_contractsTokenSol,
		"contracts/token.sol",
	)
}

func contractsTokenSol() (*asset, error) {
	bytes, err := contractsTokenSolBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/token.sol", size: 559, mode: os.FileMode(420), modTime: time.Unix(1472428910, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _contractsToken_testGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\xdd\x6e\xdb\x36\x14\xbe\xb6\x9e\x82\x33\xb0\x41\x2c\x64\xca\x76\x96\xa5\x33\xd0\x8b\x24\x4b\xba\x60\x4b\x1a\xcc\x06\x76\x19\x30\x12\x2d\x13\x91\x48\x8d\xa4\xd2\x78\x43\xde\x7d\xe7\x90\xb4\x6a\x27\x99\xeb\x4e\x17\x16\x75\x78\xfe\xbe\xef\xfc\xb8\xe5\xc5\x03\xaf\x04\x69\xb8\x54\x49\x92\xe7\x95\x9e\x55\x42\x09\xc3\x9d\x20\xfc\x5e\xc2\x99\x8c\x46\x56\xd7\xc4\xe9\x07\xa1\x18\x9e\x46\xa3\xf6\xa1\xf2\x06\x70\xd4\x9d\x8b\x57\x95\x4e\x12\xd9\xb4\xda\x38\x92\x26\x83\x61\x61\xd6\xad\xd3\xb9\x28\x4a\xcb\x87\xf0\xbd\xac\x79\x85\xef\x5a\xfb\x57\xc3\xdd\x2a\x07\xff\x78\xd6\x16\x7f\x9d\xb0\x4e\x2a\x10\xc0\xb9\x92\x6e\xd5\xdd\xb3\x42\x37\xb9\x70\x2b\x61\x44\xd7\x40\x66\xa3\xfe\xcc\x8b\x42\x77\xca\xd9\x1c\x52\x04\x2f\xaa\x1c\xfe\x1f\xa3\xfc\x1e\xc0\x0b\x55\xda\x43\xac\xe1\xa2\xd1\xea\x30\x4d\x23\x0e\xd5\xcb\xdd\xba\x15\x87\xc5\xf7\x7c\xbe\xd0\xb4\xce\x08\x57\xac\x4c\xee\xc9\x5b\xae\x73\x6e\xad\x30\x6e\x98\xd0\x24\x79\xe4\x06\x0b\xf1\x20\xd6\x63\xe2\x9f\x77\xbe\x16\xec\xd6\xc8\x47\xa8\xee\x6f\x62\xed\x2f\x27\xff\x79\xc9\xcb\xd2\x04\xd3\x80\x9d\x9d\x82\x40\x58\x1b\x6e\x26\x6f\xdf\x74\x6e\x15\x1d\x22\xc3\x6c\x61\xb8\xb2\xbc\x70\x9f\x5a\x07\xb7\x91\x6f\xb2\xd1\x88\xf4\xb3\xb9\x6c\xba\x1a\xc2\x96\x67\x41\x92\x0c\x3e\x9e\xce\xef\x7e\xbf\xba\xbe\x5a\x90\x0f\x04\xba\x84\xdd\x88\xcf\x57\xca\xa5\xc7\x63\x7c\x68\x32\xb8\xf9\x74\x73\x7e\xe1\xbd\xec\xdc\x4f\x29\x22\x5f\x76\xaa\x20\x0b\x60\xe4\x1a\x7a\x34\x6d\xc8\xbb\xd8\x5a\xec\x9a\x92\x7f\x02\x23\x19\xb9\x03\xcb\xc0\x29\xfb\x18\x1b\x1e\x50\xa7\x34\x90\xb2\xef\x3e\xf0\xd2\xdf\xde\x76\xf7\x60\xb1\xd0\x91\x83\x14\xdd\xa3\xb0\x96\x05\x18\xd0\x0d\x5b\xfb\xf4\x27\xbb\xfa\xc8\x21\xc2\x02\xfe\x00\x17\x08\x45\xb9\xe1\x51\x1b\xef\x9f\x7e\xe1\x12\x14\x37\x34\x82\xf2\x4b\x26\xa1\x01\x06\xd8\x67\x1e\x83\x95\xf6\x34\xcc\x00\xb0\x30\x18\xc4\x04\x66\xc4\x03\xca\x50\x74\xc6\x6b\xae\x0a\x31\xdb\xe6\xf4\xe7\xe9\xf4\xe8\xe8\x64\x3a\x3e\xfa\xe9\xfd\xf1\x8f\x27\x27\xc7\xef\xc7\x27\x14\x95\x9f\xb3\x43\x7d\x4f\xbe\xdd\x37\x00\xc4\x85\xc1\x6e\xb9\xb1\x02\x49\xd7\x96\x5d\x3c\x49\x97\x36\xec\x8f\x4e\xa5\x94\x26\xcf\xb1\xce\x73\x59\x41\x75\x90\x95\xd7\x3d\x4c\x03\x89\x41\xe5\x52\x61\xf1\x61\x5e\x3a\xa3\x08\x9a\xa6\x3c\x24\xf9\xa2\x89\x33\xe2\x9e\xa0\x65\x70\x2e\xfb\xf6\x95\x5a\x51\x92\xbe\x16\x66\x44\x18\xa3\x8d\x6f\xab\x81\x85\x30\x1c\xbc\x0b\x2f\x25\xb3\xbe\xe2\x18\x3f\x75\x4f\xf8\xfe\x95\xdb\x55\x4a\xd9\xd9\x1a\x5a\x32\xa5\x19\x89\xc5\x1c\xc8\xa5\xb7\xf9\xee\x03\x51\xb2\xf6\xde\x36\x99\xc2\xb7\xf7\x87\xac\x24\xbd\x14\xbc\xfd\x09\x6b\x60\xbe\x09\x99\xf6\xc1\xc1\xdb\x73\x4f\x4e\x29\xda\x5a\xaf\x53\xa9\xa4\x93\xbc\x9e\x77\x6d\x5b\xaf\x71\x32\x2b\x06\xec\x67\x44\xf1\x46\x10\xd8\x21\x30\x1b\x94\xbc\x5b\xe0\x1a\x9f\x03\x03\x00\x0c\x53\xb8\x83\x21\xc8\xc2\x72\xef\x11\xfd\xe2\x1d\x7a\x4d\x6c\xad\xbd\x3d\x8a\x85\x8c\xbd\x89\xc7\x9d\x24\x50\x80\xd1\xe1\x0d\x03\xfb\x06\x7c\xf8\x9b\x60\x97\xdc\xf1\x7a\x99\x0e\x2f\xb9\xac\x45\x09\x99\x44\x3c\x44\x89\xcf\x21\x2f\x28\x9d\x72\x06\x62\xce\xc8\xf7\x8f\x43\x9f\xa5\xc7\xdf\x8f\x07\x3b\x87\xda\x42\xdb\x60\x90\x48\xdd\x0f\xdb\x38\x31\xd4\x79\xef\x03\x9f\x80\x17\xc4\xdb\xab\x6b\x16\xd7\xd9\x1e\xa8\x10\xf4\x79\x6b\xf5\x5c\x05\xb4\xf2\x6f\xe8\x4d\xf7\x65\x03\x2d\x7c\xab\x84\xdc\x81\xce\x58\x9f\xad\x89\x98\xc0\x7a\xcb\xc8\xd0\xe7\x38\xc4\xac\x3d\x4b\xb0\x8d\x40\x3b\xfc\xcf\xde\x80\xc0\xaf\x21\xbf\xea\xd9\xc5\x5f\x1d\xaf\x53\xd7\xdb\x84\xaa\xfa\x05\xe1\xe7\x6d\xc7\x36\xce\xa0\xf5\xad\x3f\x7e\xc3\xc9\xab\x4c\xa2\x17\xba\x03\xce\xc3\x5f\x7e\x23\xb2\xa3\x5d\x5c\x77\x7d\x4f\x85\xd4\x7a\x9f\x61\x63\xec\x64\x42\xdf\xee\x91\x96\x2b\x59\xa4\x7b\x6b\x6e\xe1\x5b\x98\xb3\x43\xa8\x30\xa2\x90\xad\x14\xca\x7d\x55\x7b\xb2\x9f\xb8\x29\x00\xdd\x89\xfb\x15\x9e\x41\xfd\x65\x6c\xa4\xfb\xdf\x00\x00\x00\xff\xff\x41\x67\x7f\x33\x9a\x09\x00\x00")

func contractsToken_testGoBytes() ([]byte, error) {
	return bindataRead(
		_contractsToken_testGo,
		"contracts/token_test.go",
	)
}

func contractsToken_testGo() (*asset, error) {
	bytes, err := contractsToken_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts/token_test.go", size: 2458, mode: os.FileMode(420), modTime: time.Unix(1472428772, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesTokenSol = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x55\x4d\x6b\xdc\x48\x10\xbd\xcf\xaf\xa8\xa3\xc7\x98\x91\x31\xec\x1e\xec\xcb\x9a\xc5\x0b\x7b\x4a\x48\x06\x72\x9c\x69\xf5\x94\xa4\x26\xad\x6e\xd1\x1f\x56\x4c\xc8\x7f\x4f\xf5\x87\x3e\x2c\x8d\x9d\x60\x92\x43\xcc\x80\x07\x4d\xf5\x7b\xaf\xde\xab\x2e\x15\x05\xdc\x97\xd6\x19\xc6\x1d\x70\xad\xd2\x97\x4a\x1b\x70\x0d\x42\xe5\xa5\x84\x87\x0f\xff\xc2\xcd\x35\xec\xf5\x67\x54\x60\x1d\x53\x27\x66\x4e\x1b\x3a\xd7\x38\xd7\xd9\xdb\xa2\xa8\x85\x6b\x7c\xb9\xe3\xba\x2d\x90\x4e\x19\xf4\x6d\xf1\xf0\xff\x7b\x5b\x08\x6b\x3d\xda\xe2\xe6\x7a\xb3\x19\xa1\x13\xcc\xd7\x0d\xd0\x5f\x71\x09\xfb\x46\x58\xa0\x0f\x03\x2b\x45\xdd\x90\x86\x86\xa9\x1a\xc1\xe9\x28\x80\xb8\x89\xba\x64\x16\x47\xe6\x5d\x3c\x5a\x79\xc5\x9d\xd0\x8a\x0a\x1d\x93\x1f\x7d\xd7\xc9\xa7\x8b\x6d\xe8\x20\xd4\x39\x30\xe8\xbc\x51\x16\x2e\xbc\x50\xee\xe6\xaf\xbf\xc1\xc6\x92\xed\x5d\x3c\x4d\x84\x06\x3b\xc9\x38\x9e\xa0\x27\xf5\xb7\xf1\xe9\x50\xda\xf9\x52\x0a\x3e\x47\x4e\xa7\xa2\x56\xe6\x9d\x6e\x99\x13\x9c\x49\xf9\x04\xdc\x20\x73\x18\xe4\xd7\xe8\x1c\x9a\x49\xd7\x60\xe1\x0c\x65\x37\xa1\xd0\xa7\xd5\x8f\xc4\x9e\xfb\x8c\x1d\x8e\x1e\x59\xa1\x38\x0e\x32\x16\xc0\xc4\x65\x10\x94\x76\x11\x8c\x7b\x63\x50\x39\x52\x62\x90\xeb\x5a\x09\x4b\x98\x8c\x6a\x14\x88\xb6\x93\xd8\xd2\x8f\x2c\xea\xd1\x55\x24\x22\xe9\xbc\x11\xaa\x06\x96\x43\x7f\xee\x66\xf9\x14\xab\x28\xca\x4e\x48\x34\x49\xf1\x65\x91\xd2\xa2\xc8\x63\x37\xc0\x5a\xed\xc9\xe3\x00\x19\xd2\xb4\x3f\x74\x6f\x3c\xff\x4f\xc7\x0c\x6b\xe1\xa0\x7b\x45\x3d\xed\x89\x8a\x9d\x4e\x06\xad\x85\xca\xe8\x16\xfa\x46\xf0\x26\x1b\x22\x59\x30\xa1\x17\x34\x81\x25\x86\x3c\x8d\x40\x72\x6c\x82\x4a\x11\x47\x90\x5c\xbd\xe8\x25\x3d\x7c\x57\x5d\x0c\x1c\x89\xf6\xb5\x29\xc9\x67\xb6\x73\xc9\xe4\xb5\x20\x25\x16\xd5\x09\x8e\x87\x47\x26\x3d\x1e\x53\xe3\x21\xbd\xe3\xc1\xe9\x63\x52\x7f\x6c\x6d\xbd\x0b\x65\x68\x8e\xab\x8e\xa9\x74\xde\x6e\x8e\x83\x52\x13\x9d\xa0\x90\x56\xf5\x91\x27\x1d\x79\xee\x76\x20\x25\x43\x28\x3c\x65\x2b\xa4\xf8\xd7\x96\x7c\x6a\xe2\x3d\x4c\xf3\x97\xeb\xa0\xa7\xb1\xb0\x9e\x73\x62\xa7\x6b\x0d\x34\x9e\xc3\x10\x4d\x77\x29\xd7\x4e\x86\x39\x7d\x35\x06\x9b\x24\x6d\x27\xd3\x4a\xad\xe5\x00\xf9\x46\xc7\x0e\xe1\xdf\x11\x02\x77\x1c\x3b\x75\x12\x51\x89\x70\x71\x29\x74\x9d\x89\xb7\x84\xc6\x32\x97\xae\x7c\x8a\x38\x67\x9c\x4d\x39\xfc\xd1\x31\xfc\x47\xad\x4d\x51\x84\x46\xaf\xe0\xd7\x24\x33\x9f\xd4\xc1\x65\x4b\x16\x07\xf4\x90\x14\xd8\x6e\x9d\x9d\x5d\x99\x13\xab\x16\xf7\x38\x3b\xca\x38\x8f\x76\xb1\x52\xa6\x65\x3e\xf4\x9f\x76\xe2\x59\xb8\x73\x5e\xf7\x28\xb2\xd3\xe3\x30\xc4\xcd\x9a\xe1\x5e\xf5\x3c\x9d\xa0\x7d\xf5\x13\x9e\x67\xf4\xc9\xee\xdc\xdb\x5b\x4c\x7e\x71\xc7\x2d\xbc\xa1\x82\xb0\x89\x7f\xb3\xb9\xd9\x98\xfb\xd1\x52\x83\x2d\x13\x33\x66\xa0\x37\x99\xee\xd3\xbb\x28\x90\x2e\x9d\x09\xbf\x86\xa5\xb8\x58\xa3\xb3\x59\xcc\x52\x5f\x5b\xac\x23\xe9\xe0\x14\xed\x72\xaa\xdb\x2f\x37\x8e\x20\xa0\x2f\xa4\x65\x31\xee\xe3\xe3\x33\x63\x7f\x37\xc3\xbb\xcf\x91\xaf\xf1\x16\x9a\xc7\xe7\x2f\xe5\x7c\xb7\xf9\xb6\xf9\x1e\x00\x00\xff\xff\x57\xd5\x56\xbb\x1c\x09\x00\x00")

func librariesTokenSolBytes() ([]byte, error) {
	return bindataRead(
		_librariesTokenSol,
		"libraries/token.sol",
	)
}

func librariesTokenSol() (*asset, error) {
	bytes, err := librariesTokenSolBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/token.sol", size: 2332, mode: os.FileMode(420), modTime: time.Unix(1472424019, 0)}
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
	"contracts/token.sol": contractsTokenSol,
	"contracts/token_test.go": contractsToken_testGo,
	"libraries/token.sol": librariesTokenSol,
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
	"contracts": &bintree{nil, map[string]*bintree{
		"token.sol": &bintree{contractsTokenSol, map[string]*bintree{}},
		"token_test.go": &bintree{contractsToken_testGo, map[string]*bintree{}},
	}},
	"libraries": &bintree{nil, map[string]*bintree{
		"token.sol": &bintree{librariesTokenSol, map[string]*bintree{}},
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

