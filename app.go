package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"wails/execute"

	"wails/config"
	"wails/lib/serial"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Run() (err error) {
	if !serial.IsPortNil() {
		err = execute.Run()
	} else {
		err = errors.New("沒有串口綁定")
	}

	return
}

func (a *App) Bind(name string) (err error) {
	config.Name = name

	err = serial.InitPort()

	return
}

// 顯示彈窗
func (a *App) ShowMessageDialog(title, message string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog, // 彈窗類型：Info, Warning, Error 等
		Title:   title,              // 彈窗標題
		Message: message,            // 彈窗內容
	})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
