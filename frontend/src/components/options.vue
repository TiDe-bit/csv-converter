<script lang="ts" setup>
export default {

  data(from: number, to: number) {
    import {reactive} from 'vue'
    import {Convert} from '../../wailsjs/go/app/App'

    return {
      from: 0,
      to: 0
    }
    data.from = from;
    data.to = to;
  
    const data = reactive({
      from: 0,
      to: 0,
    })

    const saveOption = () => {
      
    }

    const deleteOption = () => {
      
    }
  
  }
}

</script>

<template>
  <div id="result" class="result">{{ data.to }}</div>
  <div id="input" class="input-box">
    <input id="from" v-model="data.from" autocomplete="off" class="input" type="number"/>
    <input id="to" v-model="data.to" autocomplete="off" class="input" type="number"/>
    <button class="btn" @click="saveOption">AddOption</button>
    <button class="btn" @click="deleteOption">X</button>
  </div>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
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
