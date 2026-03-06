import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

export default defineConfig({
  site: 'https://dsantoreis.github.io/plugz',
  base: '/plugz',
  integrations: [
    starlight({
      title: 'Plugz Docs',
      description: 'Go skills runtime and marketplace docs.',
      defaultLocale: 'en',
      social: { github: 'https://github.com/dsantoreis/plugz' },
      sidebar: [{
        label: 'Documentation',
        items: [
          { label: 'Getting Started', link: '/getting-started/' },
          { label: 'Architecture', link: '/architecture/' },
          { label: 'API Reference', link: '/api-reference/' },
          { label: 'Deployment', link: '/deployment/' }
        ]
      }]
    })
  ]
});
