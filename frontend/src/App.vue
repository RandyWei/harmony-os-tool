<script lang="ts" setup>
import ErrorView from "./components/ErrorView.vue";
import Settings from "./components/Settings.vue";
import HuaWei from "./components/HuaWei.vue";
import About from "./components/About.vue";
import Help from "./components/Help.vue";
import {WaitForDevice} from '../wailsjs/go/entry/Application'
import { onMounted, ref,Ref,reactive } from "vue";
import {models} from "../wailsjs/go/models";
import  {ConnectState}  from './models/ConnectState'
import { Util } from './utils/util'


//当前连接的设备
const device :Ref<models.Device|null> = ref(null) 

const connection = reactive({
  deviceConnectState: ConnectState.CONNECTING,
  errorTip: ""
})

//需要监测设备连接状态来决定是否激活其他的Tab
const activeName = ref('1')
/**
 * 检查环境
 * 主要是检查是否已经安装了adb
 */
 function checkEnv() {
  Util.Log("checkEnv")
  WaitForDevice().then(result => {
    Util.Log("checkEnv:"+JSON.stringify(result))
    
    if (result.length > 0) {
      connection.deviceConnectState = ConnectState.CONNECTED
      device.value = result[0]
    } else {
      connection.deviceConnectState = ConnectState.DISCONNECTED
      activeName.value = '1'
    }
  }).catch(err => {
    Util.LogE("checkEnv:"+JSON.stringify(err))
    connection.deviceConnectState = ConnectState.ERROR
    connection.errorTip = err
    activeName.value = '1'
  })
  
}

onMounted(async () => {
  checkEnv()
  setInterval(() => {
    checkEnv()
  }, 1000)
})

</script>
<template>
  <div class="common-layout h-screen">
    <el-tabs v-model="activeName" type="border-card" tab-position="left" stretch class="w-full h-full">
        <el-tab-pane name="1"><template #label><span class="font-size-4">首页</span></template>
          
          <div v-if="connection.deviceConnectState === ConnectState.CONNECTING || connection.deviceConnectState === ConnectState.DISCONNECTED" class="w-full h-screen flex justify-center items-center">
            设备未连接，正在监测设备连接状态...  <span v-loading="true" ></span>
          </div>

          <ErrorView v-else-if="connection.deviceConnectState === ConnectState.ERROR" :message="connection.errorTip"></ErrorView>

          <DeviceView :device="device" v-else-if="connection.deviceConnectState === ConnectState.CONNECTED"></DeviceView>

        </el-tab-pane>

        <el-tab-pane name="2" lazy v-if="device">
          <template #label><span class="font-size-4">{{ device?.brand  }}</span></template>
          <HuaWei :device="device"></HuaWei>
        </el-tab-pane>

        <el-tab-pane name="3"><template #label><span class="font-size-4">设置</span></template><Settings></Settings></el-tab-pane>
        <el-tab-pane name="4"><template #label><span class="font-size-4">帮助</span></template><Help></Help></el-tab-pane>
        <el-tab-pane name="5" label="Role"><template #label><span class="font-size-4">关于</span></template><About></About></el-tab-pane>
      </el-tabs>
  </div>
</template>

<style>
.el-tabs--border-card{
  border: 0;
}
.el-tabs--left.el-tabs--border-card .el-tabs__header.is-left{
  border: 0;
}

.el-tabs__item.is-left.is-active{
  border: 0  !important;
}

.el-tabs__item.is-left{
  min-width: 200px;
  justify-content: center !important;
  text-align: center !important;
  padding:  2rem 0;
  border: 0 !important;
}
</style>
