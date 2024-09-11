<template>
  <div class="common_aside">
    <a-layout-sider
      :collapse="store.collapsed"
      :theme="store.themeString"
      collapsible
    >
      <a-menu>
        <a-menu-item>
          <IconHome/>
          Menu 1
        </a-menu-item>
        <a-menu-item>
          <IconCalendar/>
          Menu 2
        </a-menu-item>
        <a-menu-item>
          <IconCalendar/>
          Menu 3
        </a-menu-item>
        <a-menu-item>
          <IconCalendar/>
          Menu 4
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
  </div>
</template>
<script lang="ts" setup>
import {useStore} from "@/stores";

const store = useStore()
</script>
<style>
.common_aside {
  .arco-layout-sider-has-trigger {
    padding-bottom: 0;
  }
}
</style>
