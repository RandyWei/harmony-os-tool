<script lang="ts" setup>
import {onMounted, reactive} from 'vue'
import {WaitForDevice,InstallAdb,GetAppDir} from '../../wailsjs/go/entry/App'
import  {ConnectState}  from '../models/ConnectState'


const data = reactive({
  deviceName: "",
  appDir : "",
  installLoading: false,
  installText:"下载安装ADB",
  installError:"",
  deviceConnectState: ConnectState.CONNECTING,
  errorTip: ""
})

/**
 * 检查环境
 * 主要是检查是否已经安装了adb
 */
function checkEnv() {
  data.deviceConnectState = ConnectState.CONNECTING
  WaitForDevice().then(result => {
    // data.deviceConnectState = result
  }).catch(err => {
    data.deviceConnectState = ConnectState.ERROR
    data.errorTip = err
  })
  
}
/**
 * 安装adb
 */
function installAdb() {
  data.installLoading = true
  data.installText = "下载安装中..."
  InstallAdb().then(result => {
    checkEnv()
    data.installText = "下载安装ADB"
    data.installLoading = false
  }).catch(err => {
    console.error(err)
    data.installError = err
    data.installText = "下载安装ADB"
    data.installLoading = false
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
    <div v-loading="data.deviceConnectState === ConnectState.CONNECTING" element-loading-background="#00000000" id="result" class="w-full h-screen">
      
      <el-result
      v-if="data.deviceConnectState === ConnectState.ERROR"
      class="h-full">
        <template #icon>
          <img src="../assets/images/error.png" width="300"  />
        </template>
        <template #title><h2 class="text-red">{{ data.errorTip }}</h2></template>
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
                <div class="flex justify-center pt pb"><el-button type="primary" class="w-full" @click="installAdb()" :loading="data.installLoading">下载安装ADB</el-button></div>
                <div><span class="text-red">{{ data.installError }}</span></div>
              </li>
            </ul>

          </div>
        </template>
      </el-result>

    </div>
    
  </main>
</template>

<style scoped>

</style>
