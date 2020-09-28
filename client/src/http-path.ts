import axios from "axios";

const PORT = 7979

export default axios.create({
  baseURL: `http://localhost:${PORT}/query`,
  headers: {
    "Content-type": "application/json",
    "Accept": "application/json"
  }
});