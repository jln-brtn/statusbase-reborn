<script setup lang="ts">
//@ts-ignore
import Banner from "assets/banner.png"
import { useWindowSize } from "@vueuse/core"

const runtimeConfig = useRuntimeConfig();
const gridCount = useGridCount()
const { width: windowWidth } = useWindowSize()

watchEffect(() => {
  const n = windowWidth.value
  gridCount.value = n > 768 ? 45 : n > 500 ? 30 : 20
})
</script>
<template>
  <div class="w-full flex flex-col items-center relative">
    <img alt="background banner" class="w-full h-54 object-cover" :src="Banner" />
    <div class="w-full max-w-screen-md absolute top-0 px-6 xl:px-0">
      <div class="mt-8 sm:mt-20 mb-10 flex flex-col sm:flex-row items-center justify-between">
        <NuxtLink
          to="/"
          class="text-white text-3xl md:text-4xl font-semibold inline-flex items-center text-shadow-md shadow-green-500"
        >
          <Logo class="mr-4"></Logo> {{ runtimeConfig.public.NAME }}
        </NuxtLink>
      </div>
      <NuxtPage />

      <Footer></Footer>
    </div>
  </div>
</template>
