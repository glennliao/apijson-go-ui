import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

import Components from 'unplugin-vue-components/vite';
import {AntDesignVueResolver} from 'unplugin-vue-components/resolvers';
import {visualizer} from "rollup-plugin-visualizer";


// https://vitejs.dev/config/

export default defineConfig({
    base: "./",
    plugins: [
        vue(),
        Components({
            directoryAsNamespace:true,
            resolvers: [AntDesignVueResolver({importStyle: true, resolveIcons: true})],
            include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
        }),
        visualizer({
            open: true
        })
    ],
})
