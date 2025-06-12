const express = require('express');
const cors = require('cors');
const { initDB } = require('./db');
const { addVisitor, getVisitorCount, getAllVisitors } = require('./handlers');

const app = express();
const port = process.env.PORT || 8080;

// Middleware
app.use(cors());
app.use(express.json());

// Routes
app.post('/visitor', addVisitor);
app.get('/visitor/count', getVisitorCount);
app.get('/visitor', getAllVisitors);

// Initialize database and start server
const startServer = async () => {
  try {
    await initDB();
    app.listen(port, () => {
      console.log(`Server starting on :${port}`);
    });
  } catch (err) {
    console.error('Failed to start server:', err);
    process.exit(1);
  }
};

startServer(); 