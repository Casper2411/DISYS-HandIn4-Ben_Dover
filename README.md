# Ricart–Agrawala Algorithm Peer-to-peer implementation

This is our project for the 4. assignment in Distributed systems course at IT-University of Copenhagen, made by the group "Ben Dover", consisting of Mathilde Gonzalez-Knudsen, Max Brix Koch, Casper Storm Frøding.

## How to run

The amount of nodes is hardcoded to 3.
To initate the nodes, you need to run three separate processes in different terminals:

`go run . 0`

`go run . 1`

`go run . 2`

The nodes ID will then be portnumber(5000) plus the command line number.

## Log file
In the log file it can be seen what the individual node receives and sends to the other notes, and when it is in critical section, noted by:
`Client x entering critical section.`

## Changing the amount of nodes
Line 64 starts a for loop which runs three times, since there are three nodes in the system. Increasing the number to e.g. 5 will enable you to have 5 nodes in the system, and you would simply need to start two more processes with:

`go run . 3`

``go run . 4``
