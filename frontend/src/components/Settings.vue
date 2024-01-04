<script lang="ts" setup>
import {useDark} from '@vueuse/core'
import { watch,onMounted,reactive } from 'vue';
import {models} from "../../wailsjs/go/models";
import * as wailsRuntime from "../../wailsjs/runtime/runtime";

import {GetAdbDir,GetLogDir,RemoveLog,RemoveAdb} from '../../wailsjs/go/entry/Application'

const isDark = useDark()

watch(isDark, (val) => {
  if(val){
    wailsRuntime.WindowSetDarkTheme()
  }else{
    wailsRuntime.WindowSetLightTheme()
  }
})

const settings = reactive({
  adbDir: {} as models.DirModel,
  logDir: {} as models.DirModel,
})

async function clearLog(){
  await RemoveLog()
  GetLogDir().then(result => {
    settings.logDir = result
  })
}

async function uninstallAdb(){
    await RemoveAdb()
    GetAdbDir().then(result => {
      settings.adbDir = result
    })
}

onMounted(() => {
  GetAdbDir().then(result => {
    settings.adbDir = result
  })

  GetLogDir().then(result => {
    settings.logDir = result
  })
})

</script>
<template>
    <el-space direction="vertical" class="w-full margin-top-10" fill>
        <el-card class="w-full" shadow="hover" :body-style="{padding: '10px'}">
            <el-row class="w-full">
                <el-col :span="12" class="text-start flex items-center">
                    暗夜模式
                </el-col>
                <el-col :span="12" class="text-end">
                    <el-switch v-model="isDark">
                    </el-switch>
                </el-col>
            </el-row>
        </el-card>

        <el-card class="w-full" shadow="hover" :body-style="{padding: '10px'}" v-if="settings.adbDir">
            <el-row class="w-full">
                <el-col :span="20" class="text-start ">
                    <div>ADB</div>
                    <div class="text-size-3 w-full">目录:{{ settings.adbDir.path }}</div>
                    <div class="text-size-3 w-full">大小:{{ settings.adbDir.size }}</div>
                </el-col>
                <el-col :span="4" class="text-end flex items-center">
                    <el-button type="primary" @click="uninstallAdb()">卸载ADB</el-button>
                </el-col>
            </el-row>
        </el-card>

        <el-card class="w-full" shadow="hover" :body-style="{padding: '10px'}">
            <el-row class="w-full">
                <el-col :span="20" class="text-start ">
                    <div>日志</div>
                    <div class="text-size-3 w-full">目录:{{ settings.logDir.path }}</div>
                    <div class="text-size-3 w-full">大小:{{ settings.logDir.size }}</div>
                </el-col>
                <el-col :span="4" class="text-end flex items-center  justify-content-end">
                    <el-button type="primary" @click="clearLog()">删除日志</el-button>
                </el-col>
            </el-row>
        </el-card>
    </el-space>
</template>