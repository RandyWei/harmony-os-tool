<script setup lang="ts">

import { Ref, onMounted, ref } from 'vue'
import { ListApps,UninstallApp,InstallExistingApp } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';

const features:Ref<models.App[][]> = ref([])

// 卸载/装回
async function OpreationApp (app:models.App)  {
        console.log("OpreationApp")
    if(app.installed){
        console.log("卸载")
        const result = await UninstallApp(app.id,app.related_ids)
        console.log("卸载结果:",result)
    }else{
        console.log("装回")
        const result = await InstallExistingApp(app.id,app.related_ids)
        console.log("装回结果:",result)
    }
    features.value = await ListApps()
}

onMounted(async () => {
    features.value = await ListApps()
  console.log("apps:",features.value)
})

</script>

<template>

    <el-scrollbar class="h-screen">
        <div class="pb-10 h-full w-full">
            <el-card header="性能" shadow="never" class="w-full"  :body-style="{padding: '10px'}">
            
            </el-card>

            <el-card header="华为全家桶" shadow="never" class="w-full"  :body-style="{padding: '0px'}">
                <el-table :data="features[1]" :flexible="true">
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
                <el-table :data="features[2]" :flexible="true">
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
                <el-table :data="features[3]" :flexible="true">
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
                <el-table :data="features[4]" :flexible="true">
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