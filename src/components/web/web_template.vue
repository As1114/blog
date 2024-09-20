<template>
  <div class="web_template">
    <div class="web_template_content">
      <slot name="content"></slot>
    </div>
    <div :class="{templateCollapsed:store.templateCollapsed}" class="web_template_aside">
      <a-layout-sider
        :collapse="store.templateCollapsed"
        :theme="store.themeString"
        collapsible
        @collapse="collapse"
      >
        <slot name="aside"></slot>
      </a-layout-sider>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {useStore} from "@/stores";

const store = useStore()

function collapse(templateCollapsed: boolean) {
  store.setTemplateCollapsed(templateCollapsed)
}
</script>
<style>
.web_template {
  display: flex;
  background-color: var(--color-neutral-2);

  .web_template_aside {
    width: 200px;

    .arco-layout-sider {
      height: calc(100vh - 80px);
    }

  }

  .web_template_aside.templateCollapsed {
    width: 48px;
  }

  .web_template_content {
    width: calc(100% - 48px);
  }

}
</style>
