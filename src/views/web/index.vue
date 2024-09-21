<template>
  <div class="web">
    <div class="web_index">
      <div class="web_index_header">
        <web_nav></web_nav>
      </div>
      <main>
        <common_aside v-if="store.isLogin" :class="{collapsed:store.collapsed}" :show-avatar="true"
                      :show-stats="false"></common_aside>
        <div :class="{isLogin:!store.isLogin}" class="web_index_content">
          <router-view v-slot="{Component}">
            <transition mode="out-in" name="fade">
              <component :is="Component"></component>
            </transition>
          </router-view>
        </div>
      </main>
      <div class="web_index_footer">
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Web_nav from "@/components/web/web_nav.vue";
import Common_aside from "@/components/common_aside.vue";
import {useStore} from "@/stores";

const store = useStore()
</script>

<style lang="scss">
.web {
  .web_index {
    display: flex;
    flex-direction: column;
    align-items: center;

    .web_index_header {
      height: 80px;
      width: 100%;
    }

    main {
      width: var(--main_width);
      display: flex;
      height: calc(100vh - 80px);
      transition: all .3s;

      .common_aside {
        width: 200px;

        .arco-layout-sider {
          height: calc(100vh - 128px);
          position: fixed;
        }
      }

      .common_aside.collapsed {
        width: 48px;

        & ~ .web_index_content {
          width: calc(100% - 48px);
        }
      }

      .web_index_content {
        width: calc(100% - 200px);
      }

      .web_index_content.isLogin {
        width: 100%;
      }
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
