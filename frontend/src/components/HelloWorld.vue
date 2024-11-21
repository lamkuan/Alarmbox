<script setup>
import {reactive} from 'vue'
import {Run, Bind, ShowMessageDialog, Login} from '../../wailsjs/go/main/App'

const account = reactive({
  username: '',
  password: '',
  logged: false,
})

const userStatus = reactive({
  loginSuccessful: false,
})

function login() {
  let message = ""

  data.loading = true

  if (account.logged) {
      message = "Logged in"
      ShowMessageDialog("Success", message)
      return
  }

  Login(account.username, account.password)
      .then(() => {
          message = "登錄成功"
          userStatus.loginSuccessful = true
          account.logged = true
          ShowMessageDialog("Success", message)
          data.loading = false
      })
      .catch((err) => {
          message = "登錄失敗: " +  err.toString()
          ShowMessageDialog("Error", message)
          data.loading = false
      })
}

const data = reactive({
  port: "",
  resultText: "輸入串口 👇",
  bind: false,
  active: false,
  loading: false,
})

function run() {

  let message = ""

  if (data.active) {
    message = "Alarmbox 已經啟動"
    ShowMessageDialog("Error", message)
    return
  }

  Run()
  .then(() => {
    message = "啟動成功"
    data.active = true
    ShowMessageDialog("Success", message)
  })
  .catch((err) => {
    message = "啟動失敗: " + err.toString()
    data.active = false
    ShowMessageDialog("Error", message)
  })
}

function bind() {

  let message = ""

  if (data.bind) {
    message = "串口已綁定成功"
    ShowMessageDialog("Error", message)
    return
  }

  Bind(data.port)
  .then(() => {
    message = "串口綁定成功"
    data.bind = true
    ShowMessageDialog("Success", message)

  })
  .catch((err) => {
    message = "串口綁定失敗: " + err.toString()
    data.bind = false
    ShowMessageDialog("Error", message)
  })
}

</script>

<template>
  <main>
<!--    <div id="result" class="result">{{data.resultText}}</div>-->
    <div id="input" class="input-box">
      <template v-if="userStatus.loginSuccessful">
        <div>登錄成功</div>
        串口： <input id="port" v-model="data.port" autocomplete="off" class="input" type="text"/>
        <button class="btn" @click="bind">綁定串口</button> <br/>
        <button class="btn" @click="run">啟動</button>
      </template>
      <template v-else>
        <div id="loginText" class="loginText">輸入Adwan的帳號和密碼</div>
        用戶名： <input id="username" v-model="account.username" type="text" /> <br/>
        密碼： <input id="password" v-model="account.password" type="password" /> <br/>
        <template v-if="data.loading">
          <div class="loading-overlay">
            <div class="spinner"></div>
          </div>
        </template>
        <template v-else>
          <button class="btn" @click="login">登錄成功</button>
        </template>
      </template>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.start_btn {
  width: 100px;
  height: 30px;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 5px solid #ccc;
  border-top-color: #007bff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
