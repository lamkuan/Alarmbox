package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"wails/execute"
	"wails/lib/adwan"

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

func (a *App) shutdown(_ context.Context) {
	if !serial.IsPortNil() {
		_ = serial.Write(serial.GlobalOff)
	}
}

func (a *App) Login(username, password string) (err error) {

	if username == "" || password == "" {
		err = errors.New("用戶名或密碼不能為空")
		return
	}

	success, err := adwan.Login(username, password, config.IP)

	if err != nil {
		log.Println(err)
		return
	}

	if success == false {
		err = errors.New("登錄失敗： 帳號或密碼錯誤")
		return
	}

	adwan.Account.Username = username
	adwan.Account.Password = password

	return
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

func (a *App) ShowMessageDialog(title, message string) {
	_, _ = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog, // 彈窗類型：Info, Warning, Error 等
		Title:   title,              // 彈窗標題
		Message: message,            // 彈窗內容
	})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
