<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();
const href = `https://github.com/${runtimeConfig.public.OWNER}/${runtimeConfig.public.REPO}`
</script>

<template>
  <div class="flex flex-col md:flex-row items-center my-20 space-y-2 md:space-x-2 md:space-y-0">
    <p class="flex items-center w-max"> 
      This page is <a class="ml-1 underline" :href="href"> open source</a>, powered by <a class="ml-1 underline" href="https://github.com/jln-brtn/statusbase-reborn">Statusbase Reborn</a>
    </p>
  </div>
</template>
