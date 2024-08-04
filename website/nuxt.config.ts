// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    ssr: false,
    devtools: {
        enabled: true,

        timeline: {
            enabled: true,
        },
    },
    runtimeConfig: {
        public: {
            OWNER: process.env.OWNER,
            REPO: process.env.REPO,
            CNAME: process.env.CNAME,
            LOGO_URL: process.env.LOGO_URL,
            NAME: process.env.NAME,
            DESCRIPTION: process.env.DESCRIPTION,
        },
    },
    modules: ["nuxt-windicss", "unplugin-icons/nuxt"],
})