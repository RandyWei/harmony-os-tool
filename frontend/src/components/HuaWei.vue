<script setup lang="ts">

import { Ref, onMounted, ref } from 'vue'
import { ListApps,UninstallApp,InstallExistingApp } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';

const apps:Ref<models.App[]> = ref([])

// 卸载/装回
async function OpreationApp (app:models.App)  {
        console.log("OpreationApp")
    if(app.installed){
        console.log("卸载")
        const result = await UninstallApp(app.id)
        console.log("卸载结果:",result)
    }else{
        console.log("装回")
        const result = await InstallExistingApp(app.id)
        console.log("装回结果:",result)
    }
    apps.value = await ListApps()
}

onMounted(async () => {
  apps.value = await ListApps()
})

</script>

<template>

    <el-container>
        <el-space wrap direction="horizontal" alignment="start" class="w-full">
            <el-card header="性能" shadow="never" class="min-width"  :body-style="{padding: '10px'}">
            
            </el-card>

            <el-card header="华为全家桶" shadow="never" class="min-width"  :body-style="{padding: '0px'}">
                <el-table :data="apps" :flexible="true">
                    <el-table-column prop="name" label="应用名称"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="智慧增值服务" shadow="never" class="min-width"  :body-style="{padding: '0px'}">
                <el-table :data="apps" :flexible="true">
                    <el-table-column prop="name" label="应用名称"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="系统功能" shadow="never" class="min-width"  :body-style="{padding: '0px'}">
                <el-table :data="apps" :flexible="true">
                    <el-table-column prop="name" label="应用名称"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small" @click="OpreationApp(scope.row)">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>
        </el-space>
    </el-container>
</template>

<style scoped>
.min-width{
    min-width: 35vw;
}

</style>