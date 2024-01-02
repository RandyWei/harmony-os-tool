<script setup lang="ts">

import { Ref, onMounted, ref } from 'vue'
import { ListApps } from "../../wailsjs/go/entry/Application";
import { models } from '../../wailsjs/go/models';

const apps:Ref<models.App[]> = ref([])

onMounted(async () => {
  apps.value = await ListApps()
  console.log(apps)
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
                            <el-button type="primary" size="small">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="智慧增值服务" shadow="never" class="min-width"  :body-style="{padding: '0px'}">
                <el-table :data="apps" :flexible="true">
                    <el-table-column prop="name" label="应用名称"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-card>

            <el-card header="系统功能" shadow="never" class="min-width"  :body-style="{padding: '0px'}">
                <el-table :data="apps" :flexible="true">
                    <el-table-column prop="name" label="应用名称"></el-table-column>
                    <el-table-column label="操作" width="100px">
                        <template #default="scope">
                            <el-button type="primary" size="small">{{ scope.row.installed ? "卸载" : "装回" }}</el-button>
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