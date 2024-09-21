<template>
  <RouterView/>
</template>
<script lang="ts" setup>
import {useStore} from "@/stores";

const store = useStore()
store.loadToken()
store.loadTheme()
</script>
<style>

</style>
