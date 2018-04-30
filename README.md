# Client Server Application

The Client-Server Application is an exercise in TCP communications via UNIX programming. Essentially, this software application has a server and a client are developed from first principles and are capable of communicating to each other through a simulated TCP implementation. In order to run the server, it must be able to connect to a free unix address on the local machine and have access to the local directory to store incoming files. Provided the client is given the correct unix address and gains ownership over a free unix address as well, it can create a TCP connection with the server, deconstruct files into packets, send the packet to the server, have the server recompile the packets, and then store the file into the server's directory.

## Prerequisites
- The local machine must be some Unix-based machine like the following:
  - macOS, Linux, etc
- Ensure that [*Go*](https://golang.org/) is installed on the local machine
- Clone an instance of this repository with the following command

`git clone https://github.com/obrienthomas13/client-server-application.git`

## Usage
- Change the enter working directory to be inside the application's repository
- Command to run the server:

`go run unix-server.go <server-unix-address>`
  - **server-unix-address**: location where the server will attempt to listen to if not occupied

- In another Bash session, change the enter working directory to be inside the applications repository
- Command to run the client:

`go run unix-client.go <client-unix-address> <server-unix-address>`
  - **client-unix-address**: location where the client will attempt to live in
  - **server-unix-address**: address of the server
- With the client running, the user will be prompted with `Enter a file name: ` to enter in the name of existing files
- Type in a file name, and the file will be sent over to the server!
- To end either the *server* or *client* simply force the program to end with `CTRL-C`

- The server is capable of listening to multiple clients. To do this simply open another local session of Bash and follow the previous instructions for the client.
