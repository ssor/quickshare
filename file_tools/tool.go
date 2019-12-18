package file_tools

import (
    "fmt"
    "path/filepath"
    "os"
    "strings"
    "io/ioutil"
    "crypto/md5"
)

type FileInfos []FileInfo

func (fis FileInfos) Find(fp string) (fi FileInfo, ok bool) {
    for _, fi := range fis {
        if fi.FullPath == fp {
            return fi, true
        }
    }
    return
}

func (fis FileInfos) Diff(newest FileInfos) (changed FileInfos) {
    for _, nfi := range newest {
        fi, ok := fis.Find(nfi.FullPath)
        if ok == false {
            changed = append(changed, nfi)
            continue
        }
        if nfi.SameWith(fi) == false {
            changed = append(changed, nfi)
            continue
        }
    }
    return
}

type FileInfo struct {
    FullPath string `json:"full_path"`
    MD5      string `json:"md5"`
}

func (fi FileInfo) SameWith(input FileInfo) bool {
    return fi.MD5 == input.MD5
}

func (fi FileInfo) String() string {
    return fmt.Sprintf("path: %s md5: %s", fi.FullPath, fi.MD5)
}

func GetFileList(root string) (list FileInfos) {
    filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
        if p == root {
            return nil
        }
        if info.IsDir() {
            return nil
        }
        name := info.Name()
        if strings.HasPrefix(name, ".") {
            return nil
        }

        var md5Check string
        {
            bs, err := ioutil.ReadFile(p)
            if err != nil {
                fmt.Println("----read file failed: ", p, " error: ", err)
            } else {
                md5Check = fmt.Sprintf("%x", md5.Sum(bs))
            }
        }

        fi := FileInfo{
            FullPath: p,
            MD5:      md5Check,
        }
        list = append(list, fi)
        //fmt.Println("name: ", name)
        //fmt.Println("full path: ", p)

        return nil
    })
    return
}
