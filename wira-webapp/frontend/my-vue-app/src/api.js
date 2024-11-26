// api.js
export async function fetchData() {
  try {
    const response = await fetch('http://localhost:8080/api/rankings');  // Replace with your actual API endpoint
    const data = await response.json();
    return data;  // Return the fetched data
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error;  // Propagate the error to be caught in the main.js
  }
}
