<script setup lang="ts">
import { defineProps, onMounted, ref, watch } from "vue";
import {models} from "../../wailsjs/go/models";
import {Free} from "../../wailsjs/go/entry/Application";

const props = defineProps({
    device: {
        type: models.Device,
        default: () => ({} as models.Device)
    }
})

const memInfo = ref("")

async function FreeMem() {
    memInfo.value = await Free()
}

onMounted(() => {
    FreeMem()
})

</script>

<template>
    <div>
        <div class="text-start w-full">当前连接的设备：{{ device.id }} {{ device.product }} {{ device.model }}</div>
        <hr>
        <div class="text-start">
            设备内存信息：
            <pre class="color-green">{{ memInfo }}</pre>
        </div>
    </div>
</template>