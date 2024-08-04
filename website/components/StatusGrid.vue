<script setup lang="ts">
import { DateTime, Settings } from "luxon";
Settings.defaultZone = "utc";

import Papa from "papaparse";
import { ref, watch, computed } from "vue";

const runtimeConfig = useRuntimeConfig();
const gridCount = useGridCount();
const props = defineProps({
  slug: String,
});

const getDateArray = function (start: string, days: number) {
  const arr: string[] = []; // Change the type to string
  for (let i = days - 1; i >= 0; i--) {
    let dt = DateTime.fromISO(start).minus({ days: i });
    arr.push(<string>dt.toISO()); // Convert Dayjs object to string
  }
  return arr;
};

const computedData = ref([]);

const fetchComputedData = async () => {
  try {
    const response: any = await $fetch(
        `https://raw.githubusercontent.com/${runtimeConfig.public.OWNER}/${runtimeConfig.public.REPO}/${runtimeConfig.public.BRANCH}/ci/logs/${props.slug}.csv`
    );

    const records = Papa.parse(response, {
      header: true,
      skipEmptyLines: true,
    });

    let dates: string[] = getDateArray(DateTime.now().toISO(), 45); // Dates array is now an array of strings

    let globalResult = [];

    for (const index in dates) {
      const date = DateTime.fromISO(dates[index]);

      let dateResult = {
        date: dates[index], // Store date as string
        uptime: 1,
      };

      let dataGroupByDates: any = records.data.filter((j: any) =>
          date.hasSame(DateTime.fromISO(j.time), "day")
      );

      if (dataGroupByDates.length > 0) {
        if (parseInt(index) + 1 === dates.length) {
          dataGroupByDates.push({
            time: DateTime.now().toISO(),
            status: "success",
          });
        }

        for (let index2 = 1; index2 < dataGroupByDates.length; index2++) {
          const currentMeasure = dataGroupByDates[index2];
          const previousMeasure = dataGroupByDates[index2 - 1];
          if (previousMeasure.status === "error") {
            const currentDate = DateTime.fromISO(currentMeasure.time);
            const previousDate = DateTime.fromISO(previousMeasure.time);

            dateResult.uptime -=
                currentDate.diff(previousDate, "minutes").toObject().minutes / 1440;
          }
        }
      } else {
        dateResult.uptime = -1;
      }

      globalResult.push(dateResult);
    }

    computedData.value = globalResult;
  } catch (error) {
    console.error(error);
  }
};

watch(() => props.slug, fetchComputedData, { immediate: true });

const overallUptime = computed(() => {
  let dateWithUptimeData = computedData.value.filter((i) => i.uptime >= 0)
  return (
      dateWithUptimeData.reduce((a, v) => a + v.uptime, 0) / dateWithUptimeData.reduce((a) => a + 1, 0)
  )
});

defineExpose({ overallUptime });
</script>

<template>
  <div v-if="computedData.length" class="flex justify-between">
    <template v-for="data in computedData" :key="data.date">
      <Status v-bind="data"></Status>
    </template>
  </div>
  <div v-else class="h-8.5 w-full bg-gray-100 rounded-lg animate-pulse"></div>
</template>
