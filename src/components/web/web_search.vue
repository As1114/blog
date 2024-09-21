<template>
  <div id="search" class="search">
    <a-auto-complete
      v-model:model-value="key"
      :data="options"
      @change="fetchData"
      @select="enter"
    >
    </a-auto-complete>
    <a-button>
      <template #icon>
        <icon-search/>
      </template>
    </a-button>
  </div>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {articleSearchByTitle, type searchDataType} from "@/api/article";
import {useRouter} from "vue-router";

const router = useRouter()
const key = ref<string>('');
const options = ref<{ label: string; value: string }[]>([]);

const fetchData = async () => {
  if (key.value) {
    const res = await articleSearchByTitle(key.value);
    options.value = transformData(res.data.list);
  } else {
    options.value = []
  }
};

const transformData = (data: searchDataType[]) => {
  return data.map(item => ({
    label: item.title,
    value: item.id,
  }));
};

function enter(id: string) {
  router.push(`/article/${id}`)
}
</script>

<style>
.search {
  display: flex;
}
</style>
