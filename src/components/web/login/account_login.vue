<template>
  <a-form ref="formRef" :model="loginForm" :size="'large'" class="account_form" layout="vertical">
    <a-form-item field="account"
                 hide-asterisk label="account">
      <a-input v-model="loginForm.account" placeholder="please enter your email..."></a-input>
    </a-form-item>
    <a-form-item field="password"
                 hide-asterisk label="password">
      <a-input-password v-model="loginForm.password" autocomplete="current-password"
                        placeholder="please enter your password..."></a-input-password>
    </a-form-item>
    <a-form-item class="captcha" field="captcha" hide-asterisk
                 label="captcha">
      <a-input v-model="loginForm.captcha"></a-input>
      <img
        v-if="picPath"
        :src="picPath"
        alt="请输入验证码"
        @click="loginVerify()"
      >
    </a-form-item>
    <a-button type="primary" @click="login">登录</a-button>
  </a-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";
import {userLogin, type userLoginType} from "@/api/user";
import {captchaGen} from "@/api/system";

const store = useStore()
const formRef = ref()
const emits = defineEmits(["ok"])
const picPath = ref('')

const loginForm = reactive<userLoginType>({
  account: "",
  password: "",
  captcha: '',
  captcha_id: '',
})

async function login() {
  let res = await userLogin(loginForm)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  await store.setToken(res.data)
  Message.success(res.msg)
  emits("ok")
}

function formReset() {
  formRef.value.resetFields(Object.keys(loginForm))
  formRef.value.clearValidate(Object.keys(loginForm))
}

defineExpose({
  formReset
})


async function loginVerify() {
  let res = await captchaGen()
  picPath.value = res.data.pic_path;
  loginForm.captcha_id = res.data.captcha_id;
}

loginVerify()

</script>

<style>
.account_form {
  .arco-form-item {
    margin-bottom: 10px;

    img {
      height: 35px;
    }
  }
}
</style>
