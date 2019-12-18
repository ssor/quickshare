package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "github.com/ssor/quickshare/file_tools"
    "github.com/ssor/quickshare/server/libs/cfg"
    "fmt"
)

const (
    assetsRoot = "./assets"
)

func main() {
    os.MkdirAll(assetsRoot, os.ModePerm)

    router := gin.Default()
    router.GET("/list", fileList)
    router.Static("/assets", assetsRoot)

    hostName, err := cfg.GetLocalAddr()
    if err != nil {
        panic(err)
    }

    router.Run(fmt.Sprintf("%s:8888", hostName.String()))
    //config := cfg.NewConfigFrom("config.json")
    //srvShare := apis.NewSrvShare(config)
    //
    //// TODO: using httprouter instead
    //mux := http.NewServeMux()
    //mux.HandleFunc(config.PathLogin, srvShare.LoginHandler)
    //mux.HandleFunc(config.PathStartUpload, srvShare.StartUploadHandler)
    //mux.HandleFunc(config.PathUpload, srvShare.UploadHandler)
    //mux.HandleFunc(config.PathFinishUpload, srvShare.FinishUploadHandler)
    //mux.HandleFunc(config.PathDownload, srvShare.DownloadHandler)
    //mux.HandleFunc(config.PathFileInfo, srvShare.FileInfoHandler)
    //mux.HandleFunc(config.PathClient, srvShare.ClientHandler)
    //
    //server := &http.Server{
    //    Addr:           fmt.Sprintf("%s:%d", config.HostName, config.Port),
    //    Handler:        mux,
    //    MaxHeaderBytes: config.MaxHeaderBytes,
    //    ReadTimeout:    time.Duration(config.ReadTimeout) * time.Millisecond,
    //    WriteTimeout:   time.Duration(config.WriteTimeout) * time.Millisecond,
    //    IdleTimeout:    time.Duration(config.IdleTimeout) * time.Millisecond,
    //}
    //
    //log.Printf("quickshare starts @ %s:%d", config.HostName, config.Port)
    //err := open.Start(fmt.Sprintf("http://%s:%d", config.HostName, config.Port))
    //if err != nil {
    //    log.Println(err)
    //}
    //log.Fatal(server.ListenAndServe())
}

type FileInfo struct {
    FullPath string `json:"full_path"`
    MD5      string `json:"md5"`
}

func fileList(c *gin.Context) {
    //var list []FileInfo
    //filepath.Walk(assetsRoot, func(p string, info os.FileInfo, err error) error {
    //    if p == assetsRoot {
    //        return nil
    //    }
    //    if info.IsDir() {
    //        return nil
    //    }
    //    name := info.Name()
    //    if strings.HasPrefix(name, ".") {
    //        return nil
    //    }
    //
    //    var md5Check string
    //    {
    //        bs, err := ioutil.ReadFile(p)
    //        if err != nil {
    //            fmt.Println("----read file failed: ", p, " error: ", err)
    //        } else {
    //            md5Check = fmt.Sprintf("%x", md5.Sum(bs))
    //        }
    //    }
    //
    //    fi := FileInfo{
    //        FullPath: p,
    //        MD5:      md5Check,
    //    }
    //    list = append(list, fi)
    //    fmt.Println("name: ", name)
    //    fmt.Println("full path: ", p)
    //
    //    return nil
    //})
    list := file_tools.GetFileList(assetsRoot)
    c.JSON(200, list)

}
