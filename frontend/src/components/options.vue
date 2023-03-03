<script lang="ts" setup>
import { booleanLiteral, numberLiteralTypeAnnotation, STATEMENT_OR_BLOCK_KEYS } from '@babel/types'
import { def } from '@vue/shared'
import {reactive} from 'vue'
import { FromToLike } from '../App.vue';
  
const {initialFrom = 0, initialTo = 0, setFunction} = defineProps<{
  initialFrom?: number,
  initialTo?: number,
  setFunction: (pair?: FromToLike) => void
}>()

const data = reactive({
  from: initialFrom,
  to: initialTo,
  visible: true
})

const saveOption = () => {
  setFunction({From: data.from, To: data.to})
}

const deleteOption = () => {
  setFunction(undefined)
  data.visible = false
}
</script>

<template name="Option">
  <div v-if="data.visible" class="container">
    <div id="result" class="result">{{ data.to }}</div>
    <div id="input" class="input-box">
      <input id="from" v-model="data.from" :on-change="() => initialFrom = Math.abs(data.from)" autocomplete="off" class="input" type="number"/>
      <input id="to" v-model="data.to" :on-change="() => initialTo = Math.abs(data.to)" autocomplete="off" class="input" type="number"/>
      <button class="btn" @click="saveOption">Save Option</button>
      <button class="btn" @click="deleteOption">X</button>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: block;
}

.result {
  line-height: 20px;
}

.input-box {
  display: flex;
  align-items: start;
  justify-content: start;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  margin-right: 1em;
  margin-left: 1em;
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
