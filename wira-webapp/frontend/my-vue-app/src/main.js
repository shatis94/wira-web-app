import Vue from 'vue';
import App from './App.vue';
import { fetchData } from './api.js';  // Importing fetchData

Vue.config.productionTip = false;

// Single Vue instance to render the app
new Vue({
  render: h => h(App),  // Rendering the App component
  data() {
    return {
      players: [],  // Holds the fetched player data
    };
  },
  created() {
    this.loadPlayerRankings();  // Fetch data when the component is created
  },
  methods: {
    async loadPlayerRankings() {
      try {
        this.players = await fetchData();  // Fetch data using fetchData() from api.js
      } catch (error) {
        console.error('Error loading player rankings:', error);
      }
    },
  },
}).$mount('#app');
