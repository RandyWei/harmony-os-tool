<script setup lang="ts">

import { Util } from '../utils/util'
</script>

<template>
    <el-card header="原作者" shadow="never" class="mb">
        <el-row>
            <el-col :span="12">
                <img src="../assets/images/header.jpg">
                <div>
                    B站：<a href="#" target="_blank" @click="Util.OpenUrl('https://space.bilibili.com/430813939')">一只靓仔琦</a>
                </div>
                <div>
                    微信公众号：靓玩科技
                </div>
            </el-col>
            <el-col :span="12">
                <img src="../assets/images/raw_code.jpg">
                <div>这是原作者的二维码，请支持原作者</div>
            </el-col>
        </el-row>
    </el-card>
    
    <el-card header="本项目" shadow="never">
        <div class="text-start">GitHub: <a href="#" target="_blank" @click="Util.OpenUrl('https://github.com/RandyWei/harmony-os-tool')">https://github.com/RandyWei/harmony-os-tool</a></div>
        <div class="text-start">Gitee: <a href="#" target="_blank" @click="Util.OpenUrl('https://gitee.com/RandyWei/harmony-os-tool')">https://gitee.com/RandyWei/harmony-os-tool</a></div>
    </el-card>
    
</template>
<style scoped>
a{
    color: cadetblue;
}
</style>