package fn

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// StaticSinglePageApp create a http.Handler to serve local SPA sources
// @param root: the root path
// @param index: the index file path relative to root
func StaticSinglePageApp(root, index string) http.Handler {
	if root == "" {
		root = "."
	}
	return spa{root, index}
}

type spa [2]string

func (h spa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h[0] == "" {
		h[0] = "."
	}
	path := filepath.ToSlash(filepath.Join(h[0], r.URL.Path))
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h[0], h[1])) // when not found or is dir,redirect to index
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h[0])).ServeHTTP(w, r)
}

// FsSinglePageApp create a http.Handler to serve fs.FS SPA sources, eg: a zip filesystem
// @param f: the fs.FS
// @param root: the root folder in fs.FS, can be empty
// @param index: the index file path in filesystem, panic when not exists
func FsSinglePageApp(f fs.FS, root, index string) http.Handler {
	if root == "" {
		root = "."
	}
	return &spaFs{f, root, Panic1(f.Open(filepath.Join(root, index)))}
}

type spaFs struct {
	fs.FS
	root  string
	index fs.File
}

func (h *spaFs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.ToSlash(filepath.Join(h.root, r.URL.Path))
	fi, err := h.FS.Open(path)
	if os.IsNotExist(err) || Panic1(fi.Stat()).IsDir() {
		if h.index == nil {
			h.index, _ = h.FS.Open("index.html")

		}
		n, m, c, err := content(h.index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, n, m, c)
		return
	} //when not found or is dir,redirect to index
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.FS(h.FS)).ServeHTTP(w, r)
}
func content(file fs.File) (string, time.Time, io.ReadSeeker, error) {
	stat, err := file.Stat()
	if err != nil {
		return "", time.Time{}, nil, err
	}
	return stat.Name(), stat.ModTime(), pack(file), nil
}

type rs struct {
	fs.File
}

var (
	errSeeker = errors.New("missing seeker")
)

func (r rs) Seek(offset int64, whence int) (int64, error) {
	s, ok := r.File.(io.Seeker)
	if !ok {
		return 0, errSeeker
	}
	return s.Seek(offset, whence)
}

func pack(f fs.File) io.ReadSeeker {
	return &rs{f}
}
