<template>
  <div class="web_index">
    <div class="web_index_header">
      </div>
    <main>
        <div class="web_index_aside"></div>
        <div class="web_index_content">
          <router-view v-slot="{Component}">
            <transition mode="out-in" name="fade">
              <component :is="Component"></component>
            </transition>
          </router-view>
        </div>
      </main>
    <div class="web_index_footer"></div>
  </div>
</template>

<script lang="ts" setup>


</script>

<style lang="scss">

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
