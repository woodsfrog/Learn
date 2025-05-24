package main

import (
    "flag"
    "fmt"
    "log/slog"
    "net/http"
    "os"
    "time"
)

// 声明一个包含应用程序版本号的字符串
const version = "1.0.0"

// 定义了一个config结构来保存我们的所有配置设置。
type config struct {
    port int
    env string
}

// 定义了一个应用程序结构体保存HTTP处理程序和中间的依赖库
type application struct {
    config config
    logger *slog.Logger
}

func main(){
    // 声明config结构实例
    var cfg config

    // 将port env命令行的值读取到config中，使用默认端口4000和development环境。
    flag.IntVar(&cfg.port,"port",4000,"API server port")
    flag.StringVar(&cfg.env,"env","development","Environment (development|staging|production)")
    flag.Parse()

    // 初始化一个记录器，将日志写入标准的out系统
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // 声明应用程序的实体，其中包含config结构和记录器
    app := &application{
        config: cfg,
        logger: logger,
    }

    // 声明一个新的servemux并添加一个请求，该请求分发给/v1/healthcheck，后续章节创建
    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

    // 声明一个HTTP服务器，该服务器侦听配置文件给的端口，进行一些合理的超时设置，并且写入错误记录器
    srv := &http.Server{
        Addr:            fmt.Sprintf(":%d", cfg.port),
        Handler:         mux,
        IdleTimeout:     time.Minute,
        ReadTimeout:     5 * time.Second,
        WriteTimeout:    10 * time.Second,
        ErrorLog:        slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }

    // 开启HTTP服务
    logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

    err := srv.ListenAndServe()
    logger.Error(err.Error())
    os.Exit(1)
}
