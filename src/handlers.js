const { pool } = require('./db');

const addVisitor = async (req, res) => {
  try {
    const {
      ip, network, version, city, region, region_code, country, country_name,
      country_code, country_code_iso3, country_capital, country_tld,
      continent_code, in_eu, postal, latitude, longitude, timezone,
      utc_offset, country_calling_code, currency, currency_name,
      languages, country_area, country_population, asn, org
    } = req.body;

    const query = `
      INSERT INTO visitors (
        ip, network, version, city, region, region_code, country, country_name,
        country_code, country_code_iso3, country_capital, country_tld,
        continent_code, in_eu, postal, latitude, longitude, timezone,
        utc_offset, country_calling_code, currency, currency_name,
        languages, country_area, country_population, asn, org
      ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,
        $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)
    `;

    const values = [
      ip, network, version, city, region, region_code, country, country_name,
      country_code, country_code_iso3, country_capital, country_tld,
      continent_code, in_eu, postal, latitude, longitude, timezone,
      utc_offset, country_calling_code, currency, currency_name,
      languages, country_area, country_population, asn, org
    ];

    await pool.query(query, values);
    res.status(201).send();
  } catch (err) {
    console.error('Error adding visitor:', err);
    res.status(500).json({ error: 'Internal server error' });
  }
};

const getVisitorCount = async (req, res) => {
  try {
    const result = await pool.query('SELECT COUNT(*) FROM visitors');
    res.send(result.rows[0].count);
  } catch (err) {
    console.error('Error getting visitor count:', err);
    res.status(500).json({ error: 'Internal server error' });
  }
};

const getAllVisitors = async (req, res) => {
  try {
    const result = await pool.query(`
      SELECT ip, network, version, city, region, region_code, country,
        country_name, country_code, country_code_iso3, country_capital,
        country_tld, continent_code, in_eu, postal, latitude, longitude,
        timezone, utc_offset, country_calling_code, currency, currency_name,
        languages, country_area, country_population, asn, org
      FROM visitors
    `);
    res.json(result.rows);
  } catch (err) {
    console.error('Error getting all visitors:', err);
    res.status(500).json({ error: 'Internal server error' });
  }
};

module.exports = {
  addVisitor,
  getVisitorCount,
  getAllVisitors
}; 