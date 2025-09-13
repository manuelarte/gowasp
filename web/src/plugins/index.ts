/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Types
import type { App } from 'vue'
import router from '../router'
import pinia from '../stores'

// Plugins
// eslint-disable-next-line perfectionist/sort-imports
import { DataLoaderPlugin } from 'unplugin-vue-router/data-loaders'
import vuetify from './vuetify'

export function registerPlugins (app: App) {
  app
    .use(DataLoaderPlugin, { router })
    .use(vuetify)
    .use(router)
    .use(pinia)
}
