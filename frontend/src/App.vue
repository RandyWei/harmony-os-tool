<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import {useDark,useToggle} from '@vueuse/core'

const isDark = useDark()
const darkToggle = useToggle(isDark)


</script>
<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <HelloWorld/>

  <el-button type="warning" round @click="darkToggle()">Toggle Dark</el-button>

</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
