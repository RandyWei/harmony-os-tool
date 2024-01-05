<script lang="ts" setup>
import Home from "./components/Home.vue";
import Settings from "./components/Settings.vue";
import HuaWei from "./components/HuaWei.vue";
import About from "./components/About.vue";
import Help from "./components/Help.vue";
import {WaitForDevice} from '../wailsjs/go/entry/Application'
import { onMounted, ref,Ref } from "vue";


//需要监测设备连接状态来决定是否激活其他的Tab
const enabled = ref(false)
const activeName = ref('1')
const connectState = ref(-1) 
/**
 * 检查环境
 * 主要是检查是否已经安装了adb
 */
 function checkEnv() {
  WaitForDevice().then(result => {
    //如果设备列表为空，则需要一直检测
    if (result.length > 0) {
      enabled.value = true
      connectState.value++
    } else {
      enabled.value = false
      activeName.value = '1'
      connectState.value = -1
    }
  }).catch(err => {
    enabled.value = false
    activeName.value = '1'
    connectState.value = -1
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
        <el-tab-pane name="1"><template #label><span class="font-size-4">首页</span></template><Home :connectState="connectState"></Home></el-tab-pane>
        <el-tab-pane name="2" label="华为" lazy ><template #label><span class="font-size-4">华为</span></template><HuaWei :connectState="connectState"></HuaWei></el-tab-pane>
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
