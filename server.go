package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "github.com/ssor/quickshare/file_tools"
    "github.com/ssor/quickshare/server/libs/cfg"
    "fmt"
    "github.com/mkideal/cli"
)

const (
//assetsRoot = "./assets"
)

type Arg struct {
    cli.Helper
    Root string `cli:"root" usage:"assets root" dft:"./assets"`
}

func main() {

    os.Exit(cli.Run(new(Arg), func(ctx *cli.Context) error {
        argv := ctx.Argv().(*Arg)
        if e := os.MkdirAll(argv.Root, os.ModePerm); e != nil {
            return e
        }
        fmt.Println("mk dir: ", argv.Root)

        router := gin.Default()
        router.GET("/list", fileList(argv.Root))
        router.Static("/assets", argv.Root)

        hostName, err := cfg.GetLocalAddr()
        if err != nil {
            panic(err)
        }

        return router.Run(fmt.Sprintf("%s:8888", hostName.String()))
    }))

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

func fileList(assetsRoot string) gin.HandlerFunc {
    f := func(c *gin.Context) {
        list := file_tools.GetFileList(assetsRoot)
        c.JSON(200, list)
    }
    return f
}
