<script lang="ts" setup>
import {Ref, onMounted, reactive,ref} from 'vue'
import {WaitForDevice,InstallAdb,GetAppDir} from '../../wailsjs/go/entry/App'
import  {ConnectState}  from '../models/ConnectState'
import DeviceView from "./DeviceView.vue";
import {models} from "../../wailsjs/go/models";

const connection = reactive({
  deviceConnectState: ConnectState.CONNECTING,
  errorTip: ""
})

const adb = reactive({
  installLoading: false,
  installText:"下载安装ADB",
  installError:"",
})

const device :Ref<models.Device> = ref(new models.Device) 

const data = reactive({
  appDir : "",
})

/**
 * 检查环境
 * 主要是检查是否已经安装了adb
 */
function checkEnv() {
  connection.deviceConnectState = ConnectState.CONNECTING
  WaitForDevice().then(result => {
    connection.deviceConnectState = ConnectState.CONNECTED
    //如果设备列表为空，则需要一直检测
    if (result.length === 0) {
      device.value = new models.Device()
      setTimeout(() => {
        checkEnv()
      }, 1000)
      return
    }else {
      device.value = result[0]
    }
  }).catch(err => {
    connection.deviceConnectState = ConnectState.ERROR
    connection.errorTip = err
  })
  
}
/**
 * 安装adb
 */
function installAdb() {
  adb.installLoading = true
  adb.installText = "下载安装中..."
  InstallAdb().then(result => {
    checkEnv()
    adb.installText = "下载安装ADB"
    adb.installLoading = false
  }).catch(err => {
    console.error(err)
    adb.installError = err
    adb.installText = "下载安装ADB"
    adb.installLoading = false
  })
}

function getAppDir() {
  GetAppDir().then(result => {
    data.appDir = result
  })
}

onMounted(() => {
  checkEnv()
  getAppDir() 
})

</script>

<template>
  <main>
    <div v-loading="connection.deviceConnectState === ConnectState.CONNECTING" element-loading-background="#00000000" id="result" class="w-full h-screen">
      
      <el-result
      v-if="connection.deviceConnectState === ConnectState.ERROR"
      class="h-full">
        <template #icon>
          <img src="../assets/images/error.png" width="300"  />
        </template>
        <template #title><h2 class="text-red">{{ connection.errorTip }}</h2></template>
        <template #sub-title>
          <div class="flex-col w-full " >
            <ul class="w-full">
              <li class="text-start">
                <div><h3>手动安装</h3></div>
                <div>如果比较熟悉adb命令的话，可以尝试手动安装并配置好环境变量</div>
                <div>官方下载链接：https://developer.android.com/tools/releases/platform-tools</div>
                <div>配套下载链接：https://gitee.com/RandyWei/harmony-os-tool/tree/main/attachment</div>
                <div>完成上述步骤后重启程序即可</div>
              </li>
              <li class="text-start">
                <h3>自动安装</h3>
                <div>如果不想手动安装，可以尝试点击下面的按钮自动下载安装</div>
                <div>安装目录为：{{ data.appDir }}</div>
                <div>自动安装的不会配置到环境变量，如果想卸载可以在设置中进行卸载</div>
                <div class="flex justify-center pt pb"><el-button type="primary" class="w-full" @click="installAdb()" :loading="adb.installLoading">下载安装ADB</el-button></div>
                <div><span class="text-red">{{ adb.installError }}</span></div>
              </li>
            </ul>

          </div>
        </template>
      </el-result>

      <DeviceView :device="device"></DeviceView>

    </div>
    
  </main>
</template>

<style scoped>

</style>
