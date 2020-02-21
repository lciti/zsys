// Code generated by vfsgen; DO NOT EDIT.

package config

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// internalAssets statically implements the virtual filesystem provided to vfsgen.
var internalAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 2, 18, 16, 42, 59, 785886507, time.UTC),
		},
		"/zsys.conf": &vfsgen۰CompressedFileInfo{
			name:             "zsys.conf",
			modTime:          time.Date(2020, 2, 18, 9, 52, 50, 950719451, time.UTC),
			uncompressedSize: 1019,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x41\x8f\xda\x40\x0c\x85\xef\xf9\x15\x4f\xe2\xd2\x1e\x5a\x01\xdb\xdd\x4a\xb9\x55\xda\x5b\xbb\x55\x0f\x95\xaa\x1e\x0d\x79\x90\x11\xc9\x4c\x64\x3b\x48\xf9\xf7\xd5\x84\x50\x10\x64\x0f\xe4\x94\x8c\xfd\xbe\x79\xb6\xe3\x3a\x98\x27\x1d\xca\x02\x58\xe0\x3b\xd9\x41\x1c\x0d\xc5\x1c\x11\x53\x10\x8c\xae\x03\x3a\x2a\xfa\x18\x1c\x69\x07\x0f\x2d\x11\x76\x60\x4c\xfd\xbe\x1e\x4f\x6a\xb6\x10\x25\x3a\xa5\x31\xfa\x08\xfc\x5d\x13\x49\x2b\x2a\xb6\x29\x56\xc1\x43\x8a\x39\x11\x9b\x7e\x7b\xa0\xc3\x5c\xd4\x21\xb1\x02\x63\x85\x4a\x9c\x86\x0f\x3b\x4d\x2d\xda\x64\x0e\xe5\x96\xd1\xe1\x09\xa9\xa9\x68\xfe\xb1\x00\xf6\xdb\x51\x24\x3b\xa7\x96\x58\x15\xc0\x81\xec\x1a\x31\x2f\xb1\x5e\x62\x81\xb7\x10\x43\xdb\xb7\x88\x7d\xbb\xa1\x66\x67\x13\xc6\xa2\x74\x56\x27\xb7\x0c\xcc\xa2\xcf\xa3\x45\x00\x9f\x10\xa5\x65\x89\xeb\xe7\xdb\x26\xb8\x8a\x0e\x63\x68\xaa\x6f\xb2\x7d\x96\x61\xfa\xb6\x2b\xe5\xcf\xff\xb7\x4e\x31\xa4\x23\x75\x14\x87\xe8\xd4\xa3\x34\xb7\xf2\x86\x71\xef\xf5\x89\xf1\x63\x7c\xcf\x72\xca\xb6\x3e\xb7\x29\x44\x54\x32\xd8\x45\x68\xd2\x76\x0d\xad\xa3\x9e\x32\xca\xab\x7b\x2b\x71\x31\x5e\xaa\xcc\xea\x2b\xd8\xd8\x42\xed\x1b\x5a\x1e\xf9\xa5\xf6\x5f\xca\x63\x48\xbd\xbd\xca\x50\xdc\x14\xb7\x2a\xe6\xec\x9e\x4f\xef\xbd\x3c\xcd\x82\xff\x90\x87\x5b\xf2\xf3\x83\xe4\xd5\x2c\xf9\x2d\x45\xaf\x6f\xd1\x5f\x66\xd1\x5f\x1f\x44\xff\xa5\xe8\x5d\x3b\xd6\xb3\xe8\xa7\xe5\x83\xec\x75\x86\xdb\x1d\x7d\x39\x4f\x7f\x79\x7e\x17\xbf\x2e\xf6\x8c\x54\x69\xf2\x44\xf3\x62\xa6\xde\x4b\xbc\x2c\xc7\xf4\x05\x5e\x85\x6d\x5e\xbb\x53\x20\xff\x0e\xc6\xbc\x8d\xf6\x2f\x00\x00\xff\xff\x3d\x02\xed\x74\xfb\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/zsys.conf"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
