<script setup lang="ts">

import { Ref, onMounted, reactive, ref } from 'vue'
import { ListApps1,ListApps2,ListApps4,ListApps3,ListApps5,UninstallApp,InstallExistingApp } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';

const features = reactive({
    apps1:[] as models.App[],
    apps2:[] as models.App[],
    apps3:[] as models.App[],
    apps4:[] as models.App[],
    apps5:[] as models.App[]
})

// 卸载/装回
async function OpreationApp (app:models.App)  {
        console.log("OpreationApp")
    if(app.installed){
        const result = await UninstallApp(app.id,app.related_ids)
        console.log("卸载结果:",result)
    }else{
        const result = await InstallExistingApp(app.id,app.related_ids)
        console.log("装回结果:",result)
    }
    features.apps1 = await ListApps1()
    features.apps2 = await ListApps2()
    features.apps3 = await ListApps3()
    features.apps4 = await ListApps4()
    features.apps5 = await ListApps5()
}

onMounted( () => {
    ListApps1().then(res => {
        features.apps1 = res
    })
    ListApps2().then(res => {
        features.apps2 = res
    })
    ListApps3().then(res => {
        features.apps3 = res
    })
    ListApps4().then(res => {
        features.apps4 = res
    })
    ListApps5().then(res => {
        features.apps5 = res
    })
})

</script>

<template>

    <el-scrollbar class="h-screen">
        <div class="pb-10 h-full w-full">
            <el-card header="性能" shadow="never" class="w-full"  :body-style="{padding: '10px'}">
            
            </el-card>

            <el-card header="华为全家桶" shadow="never" class="w-full"  :body-style="{padding: '0px'}">
                <el-table :data="features.apps2" :flexible="true">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="智慧增值服务" shadow="never" class="w-full"  :body-style="{padding: '0px'}">
                <el-table :data="features.apps3" :flexible="true">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="系统功能" shadow="never" class="w-full"  :body-style="{padding: '0px'}">
                <el-table :data="features.apps4" :flexible="true">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="冗余服务" shadow="never" class="w-full"  :body-style="{padding: '0px'}">
                <el-table :data="features.apps5" :flexible="true">
                    <el-table-column prop="name" label="应用名称" width="150px"></el-table-column>
                    <el-table-column prop="description" label="备注"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>
        </div>
    </el-scrollbar>

</template>

<style scoped>
.min-width{
    width: 100vw;
}

</style>