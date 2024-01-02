<script setup lang="ts">
import { defineProps } from "vue";
import {models} from "../../wailsjs/go/models";

const props = defineProps({
    device: {
        type: models.Device,
        default: () => ({} as models.Device)
    }
})
</script>

<template>
    <el-container>
        当前连接的设备：{{ device.id }} {{ device.product }} {{ device.model }}
    </el-container>
</template>