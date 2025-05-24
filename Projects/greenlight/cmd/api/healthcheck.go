package main

import (
    "fmt"
    "net/http"
)

// 该程序写入一个纯文本响应，包含有关应用程序状态、环境和版本信息。
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "status: available")
    fmt.Fprintf(w, "environment: %s\n", app.config.env)
    fmt.Fprintf(w, "version: %s\n", version)
}


