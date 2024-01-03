<script setup lang="ts">

import { Util } from '../utils/util'
</script>

<template>
    <img src="../assets/images/raw_code.jpg">
    <div>这是原作者的二维码，请支持原作者</div>

    <div style="margin-top: 1rem;">本项目为开源项目</div>
    <div>GitHub: <a href="#" target="_blank" @click="Util.OpenUrl('https://github.com/RandyWei/harmony-os-tool')">https://github.com/RandyWei/harmony-os-tool</a></div>
    <div>Gitee: <a href="#" target="_blank" @click="Util.OpenUrl('https://gitee.com/RandyWei/harmony-os-tool')">https://gitee.com/RandyWei/harmony-os-tool</a></div>
</template>