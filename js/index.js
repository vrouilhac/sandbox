const fs = require('fs');
const path = require('path');

const config = JSON.parse(fs.readFileSync(path.resolve(__dirname, "config.json")));

const regex1 = new RegExp("@[ar]_[0-9a-z]*", "g");
const regex2 = /[ar]_[0-9a-z]*/g;


const test = `
a_34b3 is good here and represent very well what is 
stated in\n @r_3b4. And there is also a_9b32s
`;

console.log("Test1 : " + test.match(regex1));
console.log("Test2 : " + test.match(regex2));
