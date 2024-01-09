<script setup lang="ts">
import { defineProps, onMounted, ref, watch } from "vue";
import {models} from "../../wailsjs/go/models";
import {Free,Top} from "../../wailsjs/go/entry/Application";

const props = defineProps({
    device: {
        type: Object,
        default: () => ({})
    }
})

const memInfo = ref("")
const topInfo = ref(new models.TopInfo())
const scrollBarHeight = ref(0)

async function FreeMem() {
    memInfo.value = await Free()
}

async function FetchTopInfo() {
    topInfo.value = await Top()
    onResize()
}

function onResize(){
    console.log("onResize")
    //重新计算表格高度
    let bodyElement = document.querySelector("body") as HTMLElement
    let tableElement = document.querySelector("#table") as HTMLElement
    if(!tableElement||tableElement==null||!tableElement.offsetParent){
        setTimeout(() => {
            onResize()
        }, 500)
        return
    }
    let bodyHeight = document.body.clientHeight | bodyElement.clientHeight
    let tableTopContainerHeight = (tableElement.offsetParent as HTMLElement).offsetLeft - tableElement.offsetLeft 
    scrollBarHeight.value = bodyHeight - tableTopContainerHeight
}

onMounted(() => {
    // FreeMem()
    FetchTopInfo()
    window.addEventListener("resize", onResize)
})

</script>

<template>
        <div class="w-full pb-10" id="table-container">
            <div class="sticky top-0 z-9">
                <div class="text-start w-full">当前连接的设备： {{ device.brand }} {{ device.model }}</div>
                <div class="text-start w-full text-size-3">设备序号： {{ device.id }}</div>
                <hr>
                <div class="text-start">
                    设备运行情况：
                    <pre class="color-green">{{ topInfo.tasks }}</pre>
                    <pre class="color-green">{{ topInfo.mem }}</pre>
                </div>
            </div>
            <el-table :data="topInfo.processes" id="table" :height="scrollBarHeight">
                <el-table-column prop="pid" label="进程ID"></el-table-column>
                <el-table-column prop="virt" label="虚拟内存"></el-table-column>
                <el-table-column prop="res" label="物理内存"></el-table-column>
                <el-table-column prop="shr" label="共享内存"></el-table-column>
                <el-table-column prop="cpu" label="CPU使用率(%)"></el-table-column>
                <el-table-column prop="mem" label="内存占比"></el-table-column>
                <el-table-column prop="time" label="占用CPU时间"></el-table-column>
                <el-table-column prop="args" label="进程包名"></el-table-column>
                <el-table-column prop="desc" label="备注"></el-table-column>
            </el-table>
        </div>
</template>