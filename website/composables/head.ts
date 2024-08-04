export const useCustomHead = (
  title = process.env.NAME,
  description = process.env.DESCRIPTION,
  image = process.env.LOGO_URL,
  cname = process.env.CNAME
) => {
  useHead({
    title,
    htmlAttrs: { lang: "en" },
    //viewport: "width=device-width, initial-scale=1",
    //charset: "utf-8",
    link: [{ rel: "icon", href: "/favicon.png" }],
    meta: [
      {
        name: "description",
        content: description,
      },
      { name: "twitter:card", content: "summary_large_image" },
      { name: "twitter:site", content: "@zernonia" },
      { name: "twitter:title", content: title },
      {
        name: "twitter:description",
        content: description,
      },
      { name: "twitter:image", content: image },
      { property: "og:type", content: "website" },
      { property: "og:title", content: title },
      { property: "og:url", content: "https://"+cname },
      { property: "og:image", content: image ?? image },
      {
        property: "og:description",
        content: description,
      },
    ],
  })
}
