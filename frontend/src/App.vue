<script lang="ts" setup>
import Option from './components/options.vue'
import { GetSavedOptions, Convert, SaveOptions } from '../wailsjs/go/app/App'
import { reactive, onMounted, Fragment  } from 'vue';

  export interface FromToLike {
    From: number;
    To: number;
  }

  class Pair {
    private fromTo?: FromToLike
    public readonly index: number

    constructor(index: number) {
      this.index = index
    }

    public SetFromTo(value?: FromToLike) {
      this.fromTo = value
    }

    public GetFromTo(): FromToLike | undefined {
      return this.fromTo
    }
  }

  const state = reactive<{optionsList: Array<Pair>}>({
    optionsList: []
  })

  const joinOptionsListItems = (): FromToLike[] => {
    const actualSettings: FromToLike[] = []
    state.optionsList.forEach((value) => {
      const pair = value.GetFromTo()
      if (pair) {
        actualSettings.push(pair)
      }
    })
    return actualSettings
  }

  const addNewOption = () => {
    const pair = new Pair(state.optionsList.length)
    pair.SetFromTo({From: 0, To: 0})
    state.optionsList.push(pair)
    console.log(state.optionsList)
  }

  onMounted(async () => {
    const savedPairs = await GetSavedOptions()
    savedPairs.forEach((pair, index) => {
      const newPair = new Pair(index)
      newPair.SetFromTo(pair)
      state.optionsList.push(newPair)
    })
  })
  
  const convert = async () => {
    await SaveOptions(joinOptionsListItems())
    await Convert()
  }

</script>

<template>
  <div class="container">
    <button class="btn" @click="addNewOption()">New Option</button>
    <button class="btn" @click="convert()">Apply Options</button>
    <div v-for="item in state.optionsList" class="list" >
      <div v-if="item.GetFromTo()">
        <Option :initial-from="item.GetFromTo()?.From" :initial-to="item.GetFromTo()?.To" :set-function="(value?: FromToLike) => item.SetFromTo(value)" />
      </div>
    </div>
  </div>
</template>

<style>
.list {
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  vertical-align: center;
  max-width: max-content;
}

.container {
  height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
  align-content: center;
}
</style>
