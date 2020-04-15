import axios from 'axios';

const HTTP = axios.create({
  baseURL: 'http://127.0.0.1:8008',
  headers: { 'Content-Type': 'application/json' },
});

export default HTTP;
