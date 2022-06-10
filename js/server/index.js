const axios = require("axios");

async function main() {
  try {
    const { data } = await axios.get("http://localhost:3000");
    console.log({ data: JSON.parse('{"name": "Vincent"}'), d: data });
  } catch (err) {
    console.log({ err });
  }
}

main();
