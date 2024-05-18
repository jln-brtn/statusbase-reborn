<script setup lang="ts">
import YAML from "yaml";

const { data: config } = await useAsyncData("config", async () => {
  const response: any = await $fetch(
    "https://raw.githubusercontent.com/jln-brtn/statusbase-reborn/master/ci/config.yml"
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

    <Collapse v-for="group in config.groups" :name="group.name" :sites="group.sites"/>

    <!-- <IncidentReport :incidents="incidents"></IncidentReport> -->
  </div>
</template>
