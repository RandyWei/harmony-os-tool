<script setup lang="ts">

import { Ref, onMounted, reactive, ref } from 'vue'
import { ListApps, UninstallApp, InstallExistingApp } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';
import { EventName } from '../models/event'
import * as wailsRuntime from "../../wailsjs/runtime/runtime";

const features = reactive({
    apps1: [] as models.App[],
    apps2: [] as models.App[],
    apps3: [] as models.App[],
    apps4: [] as models.App[],
    apps5: [] as models.App[]
})

const loading = reactive({
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

onMounted(() => {
    loading.apps1 = true
    loading.apps2 = true
    loading.apps3 = true
    loading.apps4 = true
    loading.apps5 = true
    wailsRuntime.EventsOn(EventName.Event_Refresh_App_List, (data: models.EventData) => {
        switch (data.type) {
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

    ListApps(1)
    ListApps(2)
    ListApps(3)
    ListApps(4)
    ListApps(5)

})

</script>

<template>
    <el-scrollbar class="h-screen">
        <div class="pb-10 h-full w-full">
            <el-card header="性能" shadow="never" class="w-full" :body-style="{ padding: '10px' }">

            </el-card>

            <el-card header="华为全家桶" shadow="never" class="w-full" :body-style="{ padding: '0px' }">
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

            <el-card header="智慧增值服务" shadow="never" class="w-full" :body-style="{ padding: '0px' }">
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

            <el-card header="系统功能" shadow="never" class="w-full" :body-style="{ padding: '0px' }">
                <el-table :data="features.apps4" :flexible="true" v-loading="loading.apps4">
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

            <el-card header="冗余服务" shadow="never" class="w-full" :body-style="{ padding: '0px' }">
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