package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/caiguanhao/gotogether"
	"github.com/gosuri/uiprogress"
)

type Reader struct {
	Reader   io.Reader
	Size     int64
	Update   func(percent int)
	progress int64
}

func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.progress += int64(n)
	r.Update(int(r.progress * 100 / r.Size))
	return n, err
}

var hasError bool

func writeError(info string, err error) bool {
	if err == nil {
		return false
	}
	hasError = true
	f, ferr := os.OpenFile("errors.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if ferr != nil {
		return true
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("[%s] %s - %s\n", time.Now().Format(time.RFC3339), info, err.Error()))
	return true
}

func main() {
	println("DLFILES by CGH")
	p := uiprogress.New()
	p.Width = 30
	p.RefreshInterval = time.Millisecond * 150
	p.Start()
	gotogether.Enumerable(files).Queue(func(item interface{}) {
		file := item.([]string)
		remote, local, name := file[0], file[1], file[2]
		err := os.MkdirAll(filepath.Dir(local), 0755)
		if writeError(local, err) {
			return
		}
		resp, err := http.Get(remote)
		if writeError(remote, err) {
			return
		}
		defer resp.Body.Close()
		out, err := os.Create(local)
		if writeError(local, err) {
			return
		}
		defer out.Close()
		bar := p.AddBar(100).PrependCompleted()
		bar.AppendFunc(func(b *uiprogress.Bar) string {
			return name
		})
		io.Copy(out, &Reader{Reader: resp.Body, Size: resp.ContentLength, Update: func(percent int) {
			bar.Set(percent)
		}})
		time.Sleep(150 * time.Millisecond)
	}).WithConcurrency(4).Run()
	p.Stop()
	p.Bars = nil
	if hasError {
		println("下载时有错误发生，请查看 errors.txt ！")
	}
	if runtime.GOOS == "windows" {
		println("下载已完成，你可以退出！")
		time.Sleep(10 * time.Second)
	} else {
		println("下载已完成！")
	}
}
