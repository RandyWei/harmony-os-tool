<script lang="ts" setup>
import {useDark,useToggle} from '@vueuse/core'
import { watch } from 'vue';

import * as wailsRuntime from "../../wailsjs/runtime/runtime";
const isDark = useDark()

watch(isDark, (val) => {
  console.log(val)
  if(val){
    wailsRuntime.WindowSetDarkTheme()
  }else{
    wailsRuntime.WindowSetLightTheme()
  }
})

</script>
<template>
    <el-container class="margin-top-10">
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
    </el-container>
</template>