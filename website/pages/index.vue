<script setup lang="ts">
import YAML from "yaml";

const { data: groups } = await useAsyncData("groups", async () => {
  const response: any = await $fetch(
    "https://jsonplaceholder.typicode.com/posts"
  );
  return YAML.parse(response);
});

const gridCount = useGridCount();
useCustomHead("StatusBase Status Page");
</script>

<template>
  <div class="">
    <OverallStatus></OverallStatus>

    <div class="mt-12 md:mt-20 font-semibold inline-flex items-end">
      <h2 class="text-xl md:text-3xl">Uptime</h2>
      <h3 class="text-gray-400 md:text-xl ml-3 md:ml-6">
        Last {{ gridCount }} days
      </h3>
    </div>

    <Collapse v-for="group in groups"></Collapse>

    <!-- <IncidentReport :incidents="incidents"></IncidentReport> -->
  </div>
</template>
