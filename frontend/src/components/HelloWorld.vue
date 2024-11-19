<script setup>
import {reactive} from 'vue'
import {Run, Bind, ShowMessageDialog} from '../../wailsjs/go/main/App'

const data = reactive({
  port: "",
  resultText: "輸入串口 👇",
  bind: false,
  active: false,
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
    message = "串口綁定失敗"
    data.bind = false
    ShowMessageDialog("Error", message)
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{data.resultText}}</div>
    <div id="input" class="input-box">
      <template v-if="data.active && data.bind">
        <div>登錄成功</div>
      </template>
      <template v-else>
        <input id="port" v-model="data.port" autocomplete="off" class="input" type="text"/>
        <button class="btn" @click="bind">綁定串口</button> <br/>
        <button class="btn" @click="run">啟動</button>
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
</style>
