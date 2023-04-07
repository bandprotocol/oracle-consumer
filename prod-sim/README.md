# How to run a simulation network

The purpose is to make the process of running multiple nodes on the same machine easier.

## Setup configuration

You can run `generate_nodes.sh` script to generate a genesis file with the required files for running nodes.

```bash
./prod-sim/generate_nodes.sh
```

## Running the nodes

After you generate all required files from the above step. Now, you can run the blockchain that has two validators (`alice`, `bob`) by using docker.

```bash
docker-compose -f prod-sim/docker-compose.yaml up 
```
