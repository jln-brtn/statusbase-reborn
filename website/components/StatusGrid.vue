<script setup lang="ts">
import dayjs from "dayjs/esm";
import Papa from "papaparse";

const gridCount = useGridCount();
const props = defineProps({
  slug: String,
});

const getDateArray = function (start: Date, days: number) {
  var arr: string[] = []; // Change the type to string
  for (let i = days - 1; i >= 0; i--) {
    let dt = dayjs.utc(start).subtract(i, "day");
    arr.push(dt.format("YYYY-MM-DD")); // Convert Dayjs object to string
  }
  return arr;
};

const { data: computedData } = await useAsyncData("computedData", async () => {
  try {
    const response: any = await $fetch(
      `https://raw.githubusercontent.com/jln-brtn/statusbase-reborn/master/ci/logs/${props.slug}.csv`
    );

    const records = Papa.parse(response, {
      header: true,
      skipEmptyLines: true,
    });

    let dates: string[] = getDateArray(new Date(), gridCount.value); // Dates array is now an array of strings

    let globalResult = [];
    for (const index in dates) {
      const date = dayjs(dates[index]); // Convert string back to Dayjs object
      let dateResult = {
        date: dates[index], // Store date as string
        uptime: 1,
      };
      let dataGroupByDates: any = records.data.filter((j: any) =>
        date.isSame(dayjs.utc(j.time), "day")
      );

      if (dates.length > 0 && parseInt(index) + 1 === dates.length) {
        dataGroupByDates.push({
          time: dayjs().format("YYYY-MM-DD HH:mm").toString(),
          status: "success",
        });
      }

      for (let index = 1; index < dataGroupByDates.length; index++) {
        const currentMeasure = dataGroupByDates[index];
        const previousMeasure = dataGroupByDates[index - 1];
        if (previousMeasure.status === "error") {
          const currentDate = dayjs(currentMeasure.time);
          const previousDate = dayjs(previousMeasure.time);
          dateResult.uptime -= currentDate.diff(previousDate, "minute") / 1440;
        }
      }
      globalResult.push(dateResult);
    }
    return globalResult;
  } catch (error) {
    console.error(error);
  }
});

const overallUptime = computed(() => {
  /*let dateWithUptimeData = computedData.filter((i) => i.uptime >= 0)
  return (
    dateWithUptimeData.reduce((a, v) => a + v.uptime * v.count, 0) / dateWithUptimeData.reduce((a, v) => a + v.count, 0)
  )*/
  return 0.5;
});

defineExpose({ overallUptime });
</script>

<template>
  <div v-if="computedData" class="flex justify-between">
    <template v-for="data in computedData" :key="data.date">
      <Status v-bind="data"></Status>
    </template>
  </div>
  <div v-else class="h-8.5 w-full bg-gray-100 rounded-lg animate-pulse"></div>
</template>
