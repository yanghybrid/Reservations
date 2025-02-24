import React, { useState, useEffect } from "react";

function App() {
  const [reservations, setReservations] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/reservations")
      .then((res) => {
        if (!res.ok) {
          throw new Error(`HTTP error! Status: ${res.status}`);
        }
        return res.json();
      })
      .then((data) => {
        console.log("Fetched data:", data); // Debugging log
        setReservations(data);
      })
      .catch((error) => {
        console.error("Error fetching reservations:", error);
        setError(error.message);
      });
  }, []);

  return (
    <div>
      <h1>Restaurant Reservations</h1>
      {error && <p style={{ color: "red" }}>Error: {error}</p>}
      <ul>
        {reservations.length > 0 ? (
          reservations.map((res) => (
            <li key={res.id}>
              <strong>{res.name}</strong> - {res.guests} guests at{" "}
              {new Date(res.dateTime).toLocaleString()}
            </li>
          ))
        ) : (
          <p>No reservations available.</p>
        )}
      </ul>
    </div>
  );
}

export default App;
