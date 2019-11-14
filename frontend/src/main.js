import 'babel-polyfill';
import Vue from 'vue';

// Setup Vuetify
import Vuetify from 'vuetify';
const opts = {
	theme: {
		dark: false
	},
}
Vue.use(Vuetify, opts);
import 'vuetify/dist/vuetify.min.css';
import '@mdi/font/css/materialdesignicons.css';

import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		vuetify: new Vuetify(opts),
		render: h => h(App)
	}).$mount('#app');
});
