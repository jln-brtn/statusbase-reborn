<script setup lang="ts">
import type { Report } from "~~/utils/interface";
import type { Dayjs } from "dayjs";
import dayjs from "dayjs/esm";

const gridCount = useGridCount()
const props = defineProps({
  slug: String,
})

const getDateArray = function (start: Date, days: number) {
  var arr: Dayjs[] = []
  for (let i = days - 1; i >= 0; i--) {
    let dt = dayjs.utc(start).subtract(i, "day")
    arr.push(dt)
  }
  return arr
}

const { data: computedData } = await useAsyncData("computedData", async () => {
  const response: any = await $fetch(
    `https://raw.githubusercontent.com/jln-brtn/statusbase-reborn/master/ci/logs/${props.slug}.csv`
  );

  let dates = getDateArray(new Date(), gridCount.value)

  console.log(props.slug);
  

  const datesPostProcess = dates.map((i) => {
    let dataGroupByDates: number[] = response
      ?.filter((j: { time: string | number | dayjs.Dayjs | Date | null | undefined; }) => i.isSame(dayjs.utc(j.time), "day"))
      .map((i: { status: string; }) => (i.status === "success" ? 1 : 0))

    let uptime = dataGroupByDates?.length ? dataGroupByDates.reduce((a, v) => a + v, 0) / dataGroupByDates.length : -1

    return {
      date: i,
      uptime,
      count: dataGroupByDates?.length,
    }
  })
  console.log(datesPostProcess);

  return datesPostProcess;
});

const overallUptime = computed(() => {
  /*let dateWithUptimeData = computedData.filter((i) => i.uptime >= 0)
  return (
    dateWithUptimeData.reduce((a, v) => a + v.uptime * v.count, 0) / dateWithUptimeData.reduce((a, v) => a + v.count, 0)
  )*/
})

defineExpose({ overallUptime })
</script>

<template>
  <div v-if="computedData" class="flex justify-between">
    <template v-for="data in computedData" :key="data.date">
      <Status v-bind="data"></Status>
    </template>
  </div>
  <div v-else class="h-8.5 w-full bg-gray-100 rounded-lg animate-pulse"></div>
</template>
