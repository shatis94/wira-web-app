const express = require('express');
const { Client } = require('pg');
const cors = require('cors');  // Import CORS to enable cross-origin requests

const app = express();
const port = 8080;

// Setup PostgreSQL client
const client = new Client({
  user: 'postgres',      // replace with your PostgreSQL username
  host: 'localhost',
  database: 'wira',      // replace with your database name
  password: '55429298',  // replace with your PostgreSQL password
  port: 5433,            // use the port you set up for PostgreSQL
});

// Connect to PostgreSQL
client.connect()
  .then(() => console.log('Connected to PostgreSQL'))
  .catch(err => console.error('Connection error', err.stack));

// Middleware for parsing JSON requests
app.use(express.json());

// Enable CORS for all incoming requests
app.use(cors());  // This will allow all incoming requests, useful for frontend to access the API

// Example route to get player rankings
app.get('/api/rankings', async (req, res) => {
  try {
    const result = await client.query('SELECT * FROM scores ORDER BY reward_score DESC LIMIT 10');
    res.json(result.rows);  // Return the player rankings
  } catch (err) {
    console.error(err.message);
    res.status(500).send('Server error');
  }
});

// Example route to add a new score
app.post('/api/score', async (req, res) => {
  const { char_id, reward_score } = req.body;

  try {
    const result = await client.query(
      'INSERT INTO scores (char_id, reward_score) VALUES ($1, $2) RETURNING *',
      [char_id, reward_score]
    );
    res.json(result.rows[0]);  // Return the newly inserted score
  } catch (err) {
    console.error(err.message);
    res.status(500).send('Server error');
  }
});

// Pagination example for rankings
app.get('/api/rankings', async (req, res) => {
  try {
    const result = await client.query('SELECT * FROM scores ORDER BY reward_score DESC LIMIT 10');
    res.json(result.rows); // Ensure this sends the rows to the frontend
  } catch (err) {
    console.error('Database query failed:', err);
    res.status(500).send('Server error');
  }
});

// Start the server
app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});

// Middleware to disable caching during development
app.use((req, res, next) => {
  res.setHeader('Cache-Control', 'no-store');
  next();
});
