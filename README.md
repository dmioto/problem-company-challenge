# Problem Company Challenge

 A server that handles requests and performs operations with a Postgres database, using Docker to setup.

## Customer Model
- First Name
- Last Name
- Email
- Password (hash)

## Running...

### Using Ubuntu 18.04

Runing the following command, docker will automatically initiate, define the API port, environment variables, build and run Postgres backend container.
```sh
sudo start.sh
```

### Using other Operational Systems

Create a Docker Network to make the connection between Postgres and backend. We'll consider the name of this network to be "problem-network".
```sh
sudo docker network create problem-network
```

Start the Postgres container and connect it to the problem-network.
```sh
sudo docker run -d -e POSTGRES_PASSWORD=postgres --name=problem-database --network=problem-network postgres:14.5-alpine
```

Run the application using the following commands:

Create a docker network to make the communication between the two containers (database and backend), which then will run in the same host.
```sh
sudo docker network create problem-network
```

Run the database. We'll consider the name of this database to be "problem-database".
```sh
sudo docker run -e POSTGRES_PASSWORD=postgres --name=problem-database --network=problem-network postgres:14.5-alpine
```

Get the IP of the database container inside of the problem-network.
```sh
IP=sudo docker network inspect problem-network | grep -o -P "(?<=\"IPv4Address\": \").*(?=/)"
echo $IP
```

Run the backend.
```sh
sudo docker run --network=problem-network teste
```


