<script setup lang="ts">
import { formatUptime } from "~~/utils/function";

const props = defineProps({
  site: Object
});
const status = ref();
</script>

<template>
  <div class="my-8 md:my-10 w-full">
    <div class="mb-4 flex items-center justify-between">
      <div class="mx-0.5 truncate">
        <div class="flex items-center">
          <StatusIcon :uptime="status?.overallUptime"></StatusIcon>
          <span class="ml-2 md:text-xl">{{ props.site?.slug }}</span>
        </div>
        <h4 v-if="props.site?.desc" class="text-sm text-gray-400">{{ props.site?.desc }}</h4>
      </div>
      <div>
        <p v-if="status?.overallUptime" class="ml-3 md:text-xl text-gray-400">
          {{ formatUptime(status?.overallUptime) }}
        </p>
        <div v-else class="bg-gray-100 rounded-lg w-24 h-7 animate-pulse"></div>
      </div>
    </div>
    <StatusGrid ref="status" />
  </div>
</template>
