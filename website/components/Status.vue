<script setup lang="ts">
import { formatUptime, statusColor, statusString } from "~~/utils/function"
import { Tippy } from "vue-tippy";
import { DateTime } from "luxon";

const props = defineProps({
  date: String,
  uptime: Number,
})

const uptimeClass = computed(() => statusColor(props.uptime))
</script>

<template>
  <ClientOnly>
    <tippy interactive trigger="click mouseenter" placement="bottom">
      <template #default="{ state }">
        <div
          class="w-2 h-7 rounded cursor-pointer transform transition hover:scale-125"
          :class="[uptimeClass, { 'scale-125': state.isMounted }]"
        ></div>
      </template>

      <template #content>
        <div class="px-4 py-2 flex flex-col items-center">
          <p>
            {{ DateTime.fromISO(props.date).toFormat("dd MMM yyyy") }}
          </p>
          <p class="text-white text-center mt-2 px-4 py-2 rounded-lg" :class="uptimeClass">
            {{ statusString(uptime, "message") }}
          </p>
          <p class="mt-2 text-sm text-gray-400">{{ formatUptime(uptime) }}</p>
        </div>
      </template>
    </tippy>
  </ClientOnly>
</template>
