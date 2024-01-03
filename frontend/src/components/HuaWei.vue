<script setup lang="ts">

import { Ref, onMounted, reactive, watch,ref } from 'vue'
import { ListApps, UninstallApp, InstallExistingApp,EnableApp,DisableApp } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';
import { EventName } from '../models/event'
import { Util } from '../utils/util'
import * as wailsRuntime from "../../wailsjs/runtime/runtime";

const props = defineProps({
    device: {
        type: models.Device,
        default: () => ({} as models.Device)
    },
    connectState: {
        type: Number,
        default: () => (-1)
    }
})

const first = ref(true)

const features = reactive({
    apps0: [] as models.App[],
    apps1: [] as models.App[],
    apps2: [] as models.App[],
    apps3: [] as models.App[],
    apps4: [] as models.App[],
    apps5: [] as models.App[]
})

const loading = reactive({
    apps0: false,
    apps1: false,
    apps2: false,
    apps3: false,
    apps4: false,
    apps5: false
})

// 卸载/装回
async function OpreationApp(type: number, app: models.App) {
    if (app.installed) {
        const result = await UninstallApp(app.id, app.related_ids)
        console.log("卸载结果:", result)
    } else {
        const result = await InstallExistingApp(app.id, app.related_ids)
        console.log("装回结果:", result)
    }
    switch (type) {
        case 0:
            loading.apps0 = true
            break;
        case 1:
            loading.apps1 = true
            break;
        case 2:
            loading.apps2 = true
            break;
        case 3:
            loading.apps3 = true
            break;
        case 4:
            loading.apps4 = true
            break;
        case 5:
            loading.apps5 = true
            break;
        default:
            break;
    }
    ListApps(type)
}

//EnableApp
async function EnabledApp(type: number,app: models.App) {
    loading.apps0 = true
    if (app.installed) {
        const result = await DisableApp(app.id)
        console.log("禁用结果:", result)
        //因为禁用进程有延迟，所以需要延迟刷新列表
        setTimeout(() => {
            ListApps(type)
        }, 1000);
    } else {
        const result = await EnableApp(app.id)
        console.log("启用结果:", result)
        //因为启用进程有延迟，所以需要延迟刷新列表
        ListApps(type)
    }
}

//使用浏览器打开链接
function OpenUrl(url: string) {
    Util.OpenUrl(url)
}

//监听设备连接
watch(() => props.connectState, (newVal, oldVal) => {
    console.log("监听设备连接", newVal, oldVal)
    if (newVal == -1) {
        first.value = true
        features.apps0 = []
        features.apps1 = []
        features.apps2 = []
        features.apps3 = []
        features.apps4 = []
        features.apps5 = []
    } else if (newVal > 0 && first.value){
        first.value = false
        loadApp()
    }
})

function loadApp(){
    loading.apps0 = true
    loading.apps1 = true
    loading.apps2 = true
    loading.apps3 = true
    loading.apps4 = true
    loading.apps5 = true
    //监听APP刷新事件
    wailsRuntime.EventsOn(EventName.Event_Refresh_App_List, (data: models.EventData) => {
        switch (data.type) {
            case 0:
                features.apps0 = data.data
                loading.apps0 = false
                break;
            case 1:
                features.apps1 = data.data
                loading.apps1 = false
                break;
            case 2:
                features.apps2 = data.data
                loading.apps2 = false
                break;
            case 3:
                features.apps3 = data.data
                loading.apps3 = false
                break;
            case 4:
                features.apps4 = data.data
                loading.apps4 = false
                break;
            case 5:
                features.apps5 = data.data
                loading.apps5 = false
                break;
            default:
                break;
        }
    })

    ListApps(0)
    ListApps(1)
    ListApps(2)
    ListApps(3)
    ListApps(4)
    ListApps(5)
}

onMounted(async() => {
    // loadApp()
})

</script>

<template>
    <el-scrollbar class="h-screen">
        <div class="pb-10 h-full w-full">
            <el-card shadow="never" class="w-full mb" :body-style="{ padding: '0px' }">
                <el-table :data="features.apps0" :flexible="true" v-loading="loading.apps0">
                    <el-table-column prop="name" label="功能名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="EnabledApp(0, scope.row)">{{ scope.row.installed
                                ? "禁用" : "启用" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="华为全家桶" shadow="never" class="w-full mb" :body-style="{ padding: '0px' }">
                <el-table :data="features.apps1" :flexible="true" v-loading="loading.apps1">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(1, scope.row)">{{ scope.row.installed
                                ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="智慧增值服务" shadow="never" class="w-full mb" :body-style="{ padding: '0px' }">
                <el-table :data="features.apps2" :flexible="true" v-loading="loading.apps2">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(2, scope.row)">{{ scope.row.installed
                                ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="系统功能" shadow="never" class="w-full mb" :body-style="{ padding: '0px' }">
                <el-table :data="features.apps3" :flexible="true" v-loading="loading.apps3">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(3, scope.row)">{{ scope.row.installed
                                ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="冗余服务" shadow="never" class="w-full mb" :body-style="{ padding: '0px' }">
                <el-table :data="features.apps4" :flexible="true"  v-loading="loading.apps4">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(4, scope.row)">{{ scope.row.installed
                                ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="性能" shadow="never" class="w-full mb" :body-style="{ padding: '0px' }">
                <div class="text-start p-5">
                    <p>据说能够用巅峰性能模式卡出全局120高刷，具体方法自行去酷安搜一下
                    ，早先我是为了打游戏才弄的这个模式的，处理器越差提升越明显，但是记得做好散热
                    ，游戏前务必打开【游戏加速】和手机的【性能模式】，原神帧率和稳定性提升明显
                    ，如您第一次开启，不想调试效果，直接全部解除就行</p>
                    <p>
                        此链接是荣耀9启用后的效果视频【B站链接】
                    </p>
                    <p> 
                        <a href="#" target="_blank" @click="OpenUrl('https://www.bilibili.com/video/BV1aS4y1d7c8?spm_id_from=333.999.0.0')">https://www.bilibili.com/video/BV1aS4y1d7c8?spm_id_from=333.999.0.0</a>
                    </p>
                    <p> 选择后按Ctrl/Command + C复制，粘贴至浏览器地址即可观看
                    </p>
                </div>
                <el-table :data="features.apps5" :flexible="true"  v-loading="loading.apps5">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(5, scope.row)">{{ scope.row.installed
                                ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>
        </div>
    </el-scrollbar>
</template>

<style scoped>
.min-width {
    width: 100vw;
}
</style>