import { defineConfig } from 'astro/config';
import tailwind from '@astrojs/tailwind';
import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';

export default defineConfig({
  site: 'https://vftalk.my.id',
  integrations: [
    tailwind(), svelte(),
    sitemap({
      filter: (page) =>
        page !== 'https://vftalk.my.id/chats/' &&
        page !== 'https://vftalk.my.id/explore/' &&
        page !== 'https://vftalk.my.id/profile/' &&
        page !== 'https://vftalk.my.id/landingpage/' &&
        page !== 'https://vftalk.my.id/setting/',
    }),
  ],
});
