<script lang="ts" setup>
import Home from "./components/Home.vue";
import Settings from "./components/Settings.vue";

//TODO 需要监测设备连接状态来决定是否激活其他的Tab



</script>
<template>
  <div class="common-layout h-screen">
    <el-tabs type="border-card" tab-position="left" stretch class="w-full h-full">
        <el-tab-pane><template #label><span class="font-size-4">首页</span></template><Home></Home></el-tab-pane>
        <el-tab-pane label="华为"><template #label><span class="font-size-4">华为</span></template>User</el-tab-pane>
        <el-tab-pane label="Config"><template #label><span class="font-size-4">Config</span></template>Config</el-tab-pane>
        <el-tab-pane label="Role"><template #label><span class="font-size-4">Role</span></template>Role</el-tab-pane>
        <el-tab-pane><template #label><span class="font-size-4">设置</span></template><Settings></Settings></el-tab-pane>
      </el-tabs>
  </div>
</template>

<style>
.el-tabs--border-card{
  border: 0;
}
.el-tabs--left.el-tabs--border-card .el-tabs__header.is-left{
  border: 0;
}

.el-tabs__item.is-left.is-active{
  border: 0  !important;
}

.el-tabs__item.is-left{
  min-width: 200px;
  justify-content: center !important;
  text-align: center !important;
  padding:  2rem 0;
  border: 0 !important;
}
</style>
