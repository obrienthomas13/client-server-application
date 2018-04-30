# Client Server Application

The Client-Server Application is an exercise in TCP communications via UNIX programming. Essentially, this software application has a server and a client are developed from first principles and are capable of communicating to each other through a simple TCP implementation. In order to run the server, it must be able to connect to an free unix address on the local machine and have access to the local directory to store incoming files. Provided the client is given the correct unix address and gain ownership over a free unix address as well, it can create a TCP connection with the server, break up the file into packets, send the file to the server, have the server recompile the packets, and then store the file into the server's directory.

## Prerequisites
- The local machine must be some Unix-based machine like the following:
  - macOS, Linux, etc
- Ensure that [*Go*](https://golang.org/) is installed on the local machine
- Clone 2 instances of this repository with the following command

`git clone https://github.com/obrienthomas13/client-server-application.git`

## Usage
- Change the enter working directory to be inside the repository
- Command to run the server:

`go run unix-server.go <server-unix-address>`
  - **server-unix-address**: location where the server will attempt to listen to if not occupied

- Change the enter working directory to be inside the repository of the other instance
- Command to run the client:

`go run unix-server.go <client-unix-address> <server-unix-address>`
  - **client-unix-address**: location where the client will attempt to live in
  - **server-unix-address**: address of the server
- With the client running, the user will be prompted with `Enter a file name: ` to enter in the name of existing files
- Type in a file name, and the file will be sent over to the server!
