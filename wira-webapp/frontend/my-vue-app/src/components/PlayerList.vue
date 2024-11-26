<template>
    <div>
      <h2>Player List</h2>
      <ul>
        <li v-for="player in players" :key="player.char_id">
          {{ player.username }} ({{ player.email }}) - Score: {{ player.reward_score }}
        </li>
      </ul>
      <button @click="loadMorePlayers">Load More</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'PlayerList',
    data() {
      return {
        players: [],      // Stores the list of players
        currentPage: 1,   // Current page for pagination
        pageSize: 10,     // Number of players per page
      };
    },
    created() {
      this.fetchPlayers();  // Fetch players when the component is created
    },
    methods: {
      // Fetch players data from the backend API with pagination
      async fetchPlayers() {
        try {
          const response = await fetch(`http://localhost:8080/api/players?page=${this.currentPage}`);
          const data = await response.json();
          this.players = data;  // Update the players array
        } catch (error) {
          console.error('Error fetching players:', error);
        }
      },
      // Load more players for pagination
      async loadMorePlayers() {
        this.currentPage += 1;  // Increment page number
        await this.fetchPlayers();  // Fetch the next page of players
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add your styles here */
  div {
    font-family: Arial, sans-serif;
    margin: 20px;
  }
  
  ul {
    list-style-type: none;
    padding: 0;
  }
  
  li {
    margin: 10px;
    font-size: 18px;
  }
  
  button {
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
    margin-top: 20px;
  }
  </style>
  