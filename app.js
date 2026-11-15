const http = require('http');
const os = require('os');
http.createServer((req, res) => {
  res.end('Respondo desde el pod: ' + os.hostname() + '\n');
}).listen(8080);
