import { createPinia } from 'pinia'
const pinia = createPinia()


import { defineStore } from 'pinia'
export const useStore = defineStore('main', {
    // other options...
    state() {
        return {
            counter: 0
        }
    },
    actions: {
        increament() {
            this.counter++
        }
    }
})

export default pinia