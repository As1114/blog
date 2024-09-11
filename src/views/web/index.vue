<template>
  <div class="web">
    <div class="web_index">
      <div class="web_index_header">
        <web_nav></web_nav>
      </div>
      <main>
        <a-layout>
          <common_aside></common_aside>
          <a-layout-content class="web_index_content">
            <router-view v-slot="{Component}">
              <transition mode="out-in" name="fade">
                <component :is="Component"></component>
              </transition>
            </router-view>
          </a-layout-content>
        </a-layout>
      </main>
      <div class="web_index_footer">

      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>


import Web_nav from "@/components/web/web_nav.vue";
import Common_aside from "@/components/common_aside.vue";
</script>

<style lang="scss">
.web {
  .web_index {
    display: flex;
    flex-direction: column;
    width: 100%;
    justify-content: center;
    align-items: center;

    .web_index_header {
      width: 100%;
    }

    main {
      width: var(--main_width);
    }
  }
}

//过渡动画
.fade-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

.fade-enter-to {
  transform: translateX(0px);
  opacity: 1;
}

.fade-leave-active, .fade-enter-active {
  transition: all 0.3s ease-out;
}
</style>
