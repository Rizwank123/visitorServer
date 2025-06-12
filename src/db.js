const { Pool } = require('pg');

const pool = new Pool({
  connectionString: 'postgres://postgres:root@localhost:5433/visitor_db?sslmode=disable'
});

const initDB = async () => {
  try {
    const client = await pool.connect();
    
    const schema = `
      CREATE TABLE IF NOT EXISTS visitors (
        id SERIAL PRIMARY KEY,
        ip TEXT,
        network TEXT,
        version TEXT,
        city TEXT,
        region TEXT,
        region_code TEXT,
        country TEXT,
        country_name TEXT,
        country_code TEXT,
        country_code_iso3 TEXT,
        country_capital TEXT,
        country_tld TEXT,
        continent_code TEXT,
        in_eu BOOLEAN,
        postal TEXT,
        latitude DOUBLE PRECISION,
        longitude DOUBLE PRECISION,
        timezone TEXT,
        utc_offset TEXT,
        country_calling_code TEXT,
        currency TEXT,
        currency_name TEXT,
        languages TEXT,
        country_area INTEGER,
        country_population BIGINT,
        asn TEXT,
        org TEXT,
        visited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      );
    `;

    await client.query(schema);
    client.release();
    console.log('Database initialized successfully');
  } catch (err) {
    console.error('Database initialization error:', err);
    process.exit(1);
  }
};

module.exports = {
  pool,
  initDB
}; 