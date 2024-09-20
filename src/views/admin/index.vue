<template>
  <div class="admin_index">
    <div class="admin_index_header">
      <admin_nav></admin_nav>
    </div>
    <div class="admin_index_content">
      <common_aside :class="{collapsed:store.collapsed}" :show-avatar="false" :show-stats="false"></common_aside>
      <div class="content">
        <router-view v-slot="{Component}">
          <transition mode="out-in" name="fade">
            <component :is="Component"></component>
          </transition>
        </router-view>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import Admin_nav from "@/components/admin/admin_nav.vue";
import Common_aside from "@/components/common_aside.vue";
import {useStore} from "@/stores";

const store = useStore()
</script>
<style lang="scss">
.admin_index {
  display: flex;
  height: 100vh;
  flex-direction: column;

  .admin_index_content {
    height: calc(100vh - 60px);
    display: flex;
    transition: all .3s;

    .common_aside {
      width: 200px;

      .arco-layout-sider {
        height: calc(100vh - 108px);
      }

    }

    .common_aside.collapsed {
      width: 48px;

      & ~ .content {
        width: calc(100% - 48px);
      }
    }

    .content {
      width: calc(100% - 200px);
      overflow-y: auto;
      overflow-x: hidden;
    }
  }
}

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
