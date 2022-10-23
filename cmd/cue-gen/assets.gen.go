// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// doc.cue
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _docCue = []byte(`// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file describes the available fields for the genopai configuration file.
// Optional fields are indicated with a question mark.

// module corresponds to the Go or CUE module of a repository, which is the
// directory prefix used by packages within the module to import each other.
// For example, for github.com/istio/api, the module is "istio.io/api".
module: string

openapi?: {
	// selfContained specifies whether all references should be included within
	// the OpenAPI output, recursively.
	selfContained: bool | *false

	// fieldFilter defines a regular expression of all fields to omit from the
	// output. It is only allowed to filter fields that add additional
	// constraints. Fields that indicate basic types cannot be removed. It is
	// an error for such fields to be excluded by this filter.
	// Fields are qualified by their Object type. For instance, the
	// minimum field of the schema object is qualified as Schema/minimum.
	fieldFilter?: string

	// expandReferences replaces references with actual objects when generating
	// OpenAPI Schema. It is an error for an CUE value to refer to itself if
	// this option is used.
	expandReferences: bool | *false
}

// The all section specifies settings for generating an aggregate OpenAPI file
// that contains all schema defined all of the directory entries.
all?: {
	// oapiFilename is the filename for the generated OpenAPI output.
	oapiFilename: =~".json$"

	// title overrides the $description associated with a directory.
	title: string

	version: =~#"^v\d."# | *"v1alpha1"
}

// directories is a map of directories, relative to the root, for which to
// process proto files.
directories: [string]: [...{
	// mode indicates which proto files to include in generated output.
	//   all      include all proto files
	//   perFile  generate a separate OpenAPI file for each proto file
	mode?: "all" | "perFile"

	// protoFiles manually specified the list of proto files to include.
	protoFiles?: [...string]

	// oapiFilename is the filename for the generated OpenAPI output. genoapi
	// will pick a default name if unspecified.
	oapiFilename?: =~".json$"

	// title overrides the $description associated with a directory.
	title?: !=""

	// version specifies the version of an OpenAPI file. The default the version
	// as specified in the directory path.
	version?: =~#"^v\d."#
}]

crd?: {
	// the output directory of the CRD file.
	dir?: string

	// the output filename of the CRDs.
	filename?: string

	// maximum description length of the fields.
	maxDescriptionLength?: int
}
`)

func docCueBytes() ([]byte, error) {
	return _docCue, nil
}

func docCue() (*asset, error) {
	bytes, err := docCueBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.cue", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"doc.cue": docCue,
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
	"doc.cue": &bintree{docCue, map[string]*bintree{}},
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
