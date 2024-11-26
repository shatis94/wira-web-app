<template>
  <div>
    <h1>Player Rankings</h1>
    <table>
      <thead>
        <tr>
          <th>Rank</th>
          <th>Username</th>
          <th>Score</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(player, index) in players" :key="player.char_id">
          <td>{{ index + 1 }}</td>
          <td>{{ player.username }}</td>
          <td>{{ player.reward_score }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'; // Import Axios for API calls

export default {
  name: 'App',
  data() {
    return {
      players: [] // Initialize players as an empty array
    };
  },
  created() {
    this.fetchPlayerRankings(); // Fetch data when the component is created
  },
  methods: {
    async fetchPlayerRankings() {
      try {
        // Call your backend API to get player rankings
        const response = await axios.get('http://localhost:8080/api/rankings'); // Ensure your backend runs on this endpoint
        this.players = response.data; // Assign the fetched data to players
      } catch (error) {
        console.error('Error fetching player rankings:', error); // Log any errors
      }
    }
  }
};
</script>


<style>
/* Add your styles here */
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  padding: 8px 12px;
  border: 1px solid #ddd;
}
</style>
