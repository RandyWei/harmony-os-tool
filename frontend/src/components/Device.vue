<script setup lang="ts">


const props = defineProps({
    device: {
        type: String,
        default: () => ("")
    }
})
</script>

<template>
    <el-container>
        当前连接的设备：{{ device }}
    </el-container>
</template>