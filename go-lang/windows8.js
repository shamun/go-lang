var net = require('net');
var server = net.createServer(function(socket){
  console.log("connection: " + socket.remoteAddress);
  socket.write("Echo server\r\n");
  socket.end("hello world\n");
});

server.listen(7000, "localhost");
console.log("TCP server");