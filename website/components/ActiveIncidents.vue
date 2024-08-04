<script setup lang="ts">
import {ref} from "vue";

const runtimeConfig = useRuntimeConfig();
const incidents = ref([]);

const fetchComputedData = async () => {
  try {
    incidents.value = await $fetch(
        `https://api.github.com/repos/${runtimeConfig.public.OWNER}/${runtimeConfig.public.REPO}/issues`
    );
  } catch (error) {
    console.error(error);
  }
};

watchEffect(() => {
  fetchComputedData();
});

</script>

<template>
  <div class="mt-12 md:mt-20 font-semibold inline-flex items-end w-full">
    <h2 class="text-xl md:text-3xl">Active Incidents</h2>
  </div>

  <div class="mt-4 md:mt-8">
    <div v-if="!incidents || incidents.length == 0">
      <div class="bg-gray-50 rounded-xl p-8 h-30 flex items-center justify-center">No incident reported</div>
    </div>
    <div v-else>
      <div class="bg-red-100 text-red-700 p-4 rounded-xl" role="alert">
        <p class="font-bold">Be Warned</p>
        <p>Something not ideal might be happening.</p>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>