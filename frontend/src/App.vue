<script lang="ts" setup>
import options from './components/options.vue'</script>

const allOptions = new Map()

<template>
  <div class="container">
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
