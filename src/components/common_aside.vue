<template>
  <div :class="{collapsed:store.collapsed}" class="common_aside">
    <a-layout-sider
      :collapse="store.collapsed"
      :theme="store.themeString"
      collapsible
      @collapse="collapse"

    >
      <div class="avatar">
        <a-avatar v-if="store.collapsed" :size="40"><img alt="" loading="lazy" src="/public/images/classroom.jpg">
        </a-avatar>
        <a-avatar v-else :size="180"><img alt="" loading="lazy" src="/public/images/classroom.jpg"></a-avatar>
      </div>
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

function collapse(collapsed: boolean) {
  store.setCollapsed(collapsed)
}
</script>
<style>
.common_aside {
  width: 200px;

  .arco-layout-sider-has-trigger {
    padding-bottom: 0;

    .avatar {
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
