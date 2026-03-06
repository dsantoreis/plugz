import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

export default defineConfig({
  integrations: [
    starlight({
      title: 'Plugz Docs',
      description: 'Go skills runtime and marketplace docs.',
      defaultLocale: 'en',
      social: { github: 'https://github.com/dsantoreis/plugz' },
      sidebar: [{ label: 'Docs', autogenerate: { directory: '.' } }]
    })
  ]
});
