<script setup lang="ts">
import CarbonChevronUp from "~icons/carbon/chevron-up";
import CarbonChevronDown from "~icons/carbon/chevron-down";

const props = defineProps({
  name: String,
  slug: String,
  sites: Array<any>,
});
// status = ref();
const showCards = ref(false);

const toggleCards = () => {
  showCards.value = !showCards.value;
};
</script>

<template>
  <div class="my-8 md:my-10 w-full">
    <Card :name="name" :slug="slug"/>

    <div class="flex items-center justify-center">
      <button class="btn" @click="toggleCards">
        <CarbonChevronDown v-if="!showCards"></CarbonChevronDown>
        <CarbonChevronUp v-if="showCards"></CarbonChevronUp>
      </button>
    </div>

    <div v-if="showCards" class="flex flex-col items-center m-4 border rounded-xl p-4">
      <Card v-for="site in sites" :name="site.name" :desc="site.desc" :slug="site.slug"/>
    </div>
  </div>
</template>

<style>
.btn {
  @apply px-4 py-1 text-sm text-green-500 font-semibold
    rounded-full border border-green-200 dark:border-green-800
    hover:text-white hover:bg-green-500 hover:border-transparent
    focus:outline-none focus:ring-2 ring-green-500 ring-opacity-40;
}
</style>
