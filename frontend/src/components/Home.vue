<script lang="ts" setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/entry/App'
import  {ConnectState}  from '../models/ConnectState'


const data = reactive({
  deviceName: "",
  deviceConnectState: ConnectState.CONNECTING,
})

function greet() {
 
  Greet("data.name").then(result => {
    data.deviceName = result
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.deviceName }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.deviceName" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>

<style scoped>

</style>
