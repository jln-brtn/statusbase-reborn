<script setup lang="ts">
import { formatUptime } from "~~/utils/function";
import CarbonChevronUp from "~icons/carbon/chevron-up";

const props = defineProps({
  meta_data: Object,
  report_data: Object,
});
const status = ref();
</script>

<template>
  <div class="my-8 md:my-10 w-full">
    <div class="mb-4 flex items-center justify-between">
      <div class="mx-0.5 truncate">
        <div class="flex items-center">
          <StatusIcon :uptime="status?.overallUptime"></StatusIcon>
          <span class="ml-2 md:text-xl">{{ meta_data?.title }}</span>
        </div>
      </div>
      <div>
        <p v-if="status?.overallUptime" class="ml-3 md:text-xl text-gray-400">
          {{ formatUptime(status?.overallUptime) }}
        </p>
        <div v-else class="bg-gray-100 rounded-lg w-24 h-7 animate-pulse"></div>
      </div>
    </div>
    <StatusGrid ref="status" :data="report_data?.body" />
    <div class="flex items-center justify-center pt-2">
      <button class="btn">
        <CarbonChevronUp class=""></CarbonChevronUp>
      </button>
    </div>

    <div class="flex flex-col items-center m-4 border rounded-xl p-4">
      <Card></Card>
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
