<script setup lang="ts">
import { formatUptime } from "~~/utils/function";
import CarbonChevronUp from "~icons/carbon/chevron-up";
import CarbonChevronDown from "~icons/carbon/chevron-down";

const props = defineProps({
  slug: String,
  sites: Array<any>,
});
const status = ref();
</script>

<template>
  <div class="my-8 md:my-10 w-full">
    <Card :slug="slug"/>

    <div class="flex items-center justify-center pt-2">
      <button class="btn">
        <CarbonChevronDown></CarbonChevronDown>
      </button>
    </div>

    <div v-if="true" class="flex flex-col items-center m-4 border rounded-xl p-4">
      <Card v-for="site in sites" :site="site"/>
    </div>
  </div>
</template>

<style>
.btn {
  @apply px-4 py-1 text-sm text-purple-500 font-semibold
    rounded-full border border-purple-200 dark:border-purple-800
    hover:text-white hover:bg-purple-500 hover:border-transparent
    focus:outline-none focus:ring-2 ring-purple-500 ring-opacity-40;
}
</style>
