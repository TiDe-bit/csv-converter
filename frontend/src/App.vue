<script lang="ts" setup>
  import options from './components/options.vue'
  import { GetSavedOptions } from '../wailsjs/go/app/App'
import { VueElement } from 'vue';

  export interface FromToLike {
    From: number;
    To: number;
  }

  const allOptionSetters: Array<VueElement> = []

  const addNewOption = () => {
    allOptionSetters.push(<options />)
  }


  const allOptions = new Map<string,FromToLike>()
  GetSavedOptions().then(savedPairs => {
    savedPairs.forEach(pair => {
      allOptions.set(pair.To.toString(), pair)
    } 
    )    
  })



</script>



<template>
  <div class="container">
    <button class="btn" @click="addNewOption">New Option</button>

    <options/>
  </div>
</template>

<style>
.container {
  height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
  align-content: center;
}
</style>
