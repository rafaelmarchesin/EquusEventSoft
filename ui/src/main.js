import g from 'guark'
import Vue from 'vue'
import App from './App.vue'
import store from './store'
import vuetify from './plugins/vuetify'
import router from './router'

Vue.config.productionTip = false

new Vue({
    store,
    render: h => h(App),
    created: () => g.hook("created"),
    vuetify,
    router,
    mounted: () => g.hook("mounted")
}).$mount('#app')