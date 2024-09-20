<template>
  <div class="search">
    <a-auto-complete v-model:model-value="key" :data="options" @select="goTo">
    </a-auto-complete>
    <a-button>
      <template #icon>
        <icon-search/>
      </template>
    </a-button>
  </div>
</template>
<script lang="ts" setup>
import {ref} from "vue";

const key = ref<string>('')


const options = [
  {label: "Option 1", value: "1"},
  {label: "Option 2", value: "2"},
  {label: "Option 3", value: "3"},
];

function goTo() {
  console.log(key.value)
}
</script>
<style>
.search {
  display: flex;
}
</style>
