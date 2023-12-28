<script lang="ts" setup>


</script>
<template>
  <div class="common-layout h-screen">
    <el-tabs type="border-card" tab-position="left" stretch class="w-full h-full">
        <el-tab-pane><template #label><span class="font-size-4">首页</span></template>首页</el-tab-pane>
        <el-tab-pane label="华为"><template #label><span class="font-size-4">华为</span></template>User</el-tab-pane>
        <el-tab-pane label="Config"><template #label><span class="font-size-4">Config</span></template>Config</el-tab-pane>
        <el-tab-pane label="Role"><template #label><span class="font-size-4">Role</span></template>Role</el-tab-pane>
        <el-tab-pane label="Task"><template #label><span class="font-size-4">Task</span></template>Task</el-tab-pane>
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
.el-tabs__item{
  min-width: 200px;
  justify-content: center;
  text-align: center;
  padding:  2rem 0;
  border: 0;
}
</style>
