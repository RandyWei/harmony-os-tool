<script lang="ts" setup>
import { reactive,onMounted} from 'vue'
import {useDark} from '@vueuse/core'
import {InstallAdb,GetAppDir} from '../../wailsjs/go/entry/Application'
import { Util } from '../utils/util'

const isDark = useDark()

const props = defineProps({
    message: {
        type: String,
        default: () => ("")
    }
})

const adb = reactive({
  installLoading: false,
  installText:"下载安装ADB",
  installError:"",
})

const data = reactive({
  appDir : "",
})

/**
 * 安装adb
 */
function installAdb() {
  adb.installLoading = true
  adb.installText = "下载安装中..."
  InstallAdb().then(result => {
    Util.Log("installAdb:"+JSON.stringify(result))
    adb.installText = "下载安装ADB"
    adb.installLoading = false
  }).catch(err => {
    Util.LogE("installAdb:"+JSON.stringify(err))
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
  getAppDir()
})

</script>

<template>
  <main>
    <div id="result" class="w-full h-screen">
      
      <el-result
      class="h-full">
        <template #icon>
          <img src="../assets/images/error.png" width="300" :class=' isDark ? "" : "darkImg" ' />
        </template>
        <template #title><h2 class="text-red">{{ message }}</h2></template>
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
    </div>
    
  </main>
</template>

<style scoped>
.darkImg{
  filter: invert(1) hue-rotate(.5turn);
}
</style>
