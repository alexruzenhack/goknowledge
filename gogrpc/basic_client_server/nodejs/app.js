'use strict'

const path = require('path');
const PROTO_PATH = path.join('/home', 'alex', 'go', 'src', 'goknowledge', 'gogrpc', 'basic_client_server', 'pb', 'messages.proto');
const SERVER_ADDR = 'localhost:50000';

const grpc = require('grpc');
const HelloService = grpc.load(PROTO_PATH).HelloService;
const client = new HelloService(SERVER_ADDR, grpc.credentials.createInsecure());

function main() {
  client.sayHello({Name: 'Alex'}, function (error, response){
    if (error) {
      console.log(error);
      return;
    }
    console.log(response.Message);
  });
}

main();