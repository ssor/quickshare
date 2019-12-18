package main

import (
    "os"
    "github.com/mkideal/cli"
    "fmt"
    "strings"
    "github.com/ssor/req"
    "github.com/ssor/quickshare/file_tools"
    "path/filepath"
)

const (
    assetsRoot = "./assets"
)

type Arg struct {
    cli.Helper
    Host string `cli:"host" usage:"host of file server" dft:"http://127.0.0.1:8888"`
}

func main() {
    os.MkdirAll(assetsRoot, os.ModePerm)

    os.Exit(cli.Run(new(Arg), func(ctx *cli.Context) error {
        argv := ctx.Argv().(*Arg)
        host := argv.Host
        if len(host) <= 0 {
            return fmt.Errorf("host needed")
        }
        if strings.HasPrefix(host, "http") == false {
            host = "http://" + host
        }
        if strings.HasSuffix(host, "/") {
            host = strings.Replace(host, "/", "", 1)
        }

        return startSyncFile(host)
    }))
}

func startSyncFile(host string) error {
    url := host + "/list"
    r, err := req.Get(url)
    if err != nil {
        return err
    }
    var serverFiles file_tools.FileInfos
    err = r.ToJSON(&serverFiles)
    if err != nil {
        return err
    }

    fmt.Printf("%d files get from server\n", len(serverFiles))

    //for _, fi := range serverFiles {
    //    fmt.Println(fi)
    //}
    localFiles := file_tools.GetFileList(assetsRoot)
    fmt.Printf("%d local files \n", len(localFiles))
    for _, fi := range localFiles {
        fmt.Println(fi)
    }

    newestUpdated := localFiles.Diff(serverFiles)

    if len(newestUpdated) <= 0 {
        fmt.Println("no files changed")
    } else {
        fmt.Printf("%d files changed: \n", len(newestUpdated))
        for _, fi := range newestUpdated {
            if e := downloadFile(fi, host); e != nil {
                fmt.Printf("download file (%s) failed: %s \n", fi, e)
            } else {
                fmt.Printf("download file (%s) OK \n", fi, )
            }
        }
    }
    return nil
}

func downloadFile(fi file_tools.FileInfo, host string) error {
    fileDir := filepath.Join(assetsRoot, filepath.Dir(fi.FullPath))
    if e := os.MkdirAll(fileDir, os.ModePerm); e != nil {
        return e
    }
    url := host + "/assets/" + fi.FullPath
    fmt.Println(url)

    r, err := req.Get(url)
    if err != nil {
        return fmt.Errorf("url: %s err: %s", url, err)
    }
    if e := r.ToFile(filepath.Join(fileDir, filepath.Base(fi.FullPath))); e != nil {
        return e
    }
    return nil
}
