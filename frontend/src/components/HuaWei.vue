<script setup lang="ts">

import { Ref, onMounted, reactive, watch,ref } from 'vue'
import { UninstallApp, InstallExistingApp,EnableApp,DisableApp,ListModuleApps,ListModules } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';
import { EventName } from '../models/event'
import { Util } from '../utils/util'
import * as wailsRuntime from "../../wailsjs/runtime/runtime";

const props = defineProps({
    device: {
        type: models.Device,
        default: () => ({} as models.Device)
    }
})

const pageData = reactive({
    modules: [] as models.Module[],
    loadings:[] as boolean[]
})

// 卸载/装回
async function OpreationApp(index: number,moduleId:string, app: models.App) {
    pageData.loadings[index] = true
    if (app.installed) {
        const result = await UninstallApp(app.id, app.related_ids)
        console.log("卸载结果:", result)
    } else {
        const result = await InstallExistingApp(app.id, app.related_ids)
        console.log("装回结果:", result)
    }
    ListModuleApps(props.device.brand,moduleId)
}

//EnableApp
async function EnabledApp(index: number,moduleId:string,app: models.App) {
    pageData.loadings[index] = true
    if (app.installed) {
        const result = await DisableApp(app.id)
        console.log("禁用结果:", result)
        //因为禁用进程有延迟，所以需要延迟刷新列表
        setTimeout(() => {
            ListModuleApps(props.device.brand,moduleId)
        }, 1000);
    } else {
        const result = await EnableApp(app.id)
        console.log("启用结果:", result)
        //因为启用进程有延迟，所以需要延迟刷新列表
        ListModuleApps(props.device.brand,moduleId)
    }
}

//使用浏览器打开链接
function OpenUrl(url: string) {
    Util.OpenUrl(url)
}


async function loadApp(){

    //监听APP刷新事件
    wailsRuntime.EventsOn(EventName.Event_Refresh_App_List, (data: models.EventData) => {
        for (let i = 0; i < pageData.modules.length; i++) {
            if (pageData.modules[i].id == data.type) {
                pageData.loadings[i] = false
                pageData.modules[i].apps = data.data
                break
            }
        }
    })

    ListModules(props.device.brand).then(result => {
        pageData.modules = result
        for (let i = 0; i < pageData.modules.length; i++) {
            pageData.loadings[i] = true
            ListModuleApps(props.device.brand,pageData.modules[i].id)
        }
    }).catch(err => {
    })
    
}

onMounted(async() => {
    loadApp()
})

</script>

<template>
    <el-scrollbar class="h-screen">
        <div class="pb-10 h-full w-full">
            <el-card v-for="(module,index) in pageData.modules" :key="module.id" :header="module.name" shadow="never" class="w-full mb" :body-style="{ padding: '0px' }"   >
                <div class="text-start p-5" v-if="module.name.search('性能')!=-1">
                    <p>
                        性能模式是华为推出的一种特殊的模式，可以让华为手机在游戏中获得更好的性能，
                        但是开启后会影响到系统的稳定性，请谨慎使用，如有问题请联系作者
                    </p>
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
                <el-table :data="module.apps" :flexible="true" v-loading="pageData.loadings[index]">
                    <el-table-column prop="name" label="功能名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="EnabledApp(index,module.id, scope.row)" v-if="module.type==='disable'">
                                {{ scope.row.installed ? "禁用" : "启用" }}
                            </el-button>
                            <el-button type="primary" size="small" @click="OpreationApp(index,module.id, scope.row)" v-else>{{ scope.row.installed
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