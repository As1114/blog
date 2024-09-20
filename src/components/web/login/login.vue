<template>
  <a-modal :footer="false" :visible="props.visible" class="login"
           @cancel="close">
    <template #title>Login</template>
    <div class="login_way">
      <a-tabs default-active-key="1">
        <a-tab-pane key="1" title="账号登录">
          <account_login @ok="close"></account_login>
        </a-tab-pane>
        <a-tab-pane key="2" title="其他登录">
        </a-tab-pane>
      </a-tabs>
    </div>
  </a-modal>
</template>

<script lang="ts" setup>
import Account_login from "@/components/web/login/account_login.vue";

interface Props {
  visible: boolean
}

const props = defineProps<Props>()

const emits = defineEmits(["update:visible"])


function close() {
  emits("update:visible", false)
}
</script>

<style>

</style>
