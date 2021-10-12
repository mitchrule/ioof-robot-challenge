# IOOF Robot Challenge - Mitch Rule (October 2021)

## Compiling
The Go compiler is required to build this project. It can be downloaded [here](https://golang.org/dl/).

After installing the Go compiler, clone this project and run the following from the root directory.

```
go build
```

This will produce the executable file `ioof-robot-challenge`.

## Running
This program takes input from stdin. This can be through user input or piped in from the command line. A summary of the available commands are as follows.

* `PLACE x,y,f`
  * Place the robot on the table at coordinates `(x,y)` facing in direction `f`.
  * When the program starts it will only accept place commands until it receives valid input.
* `MOVE`
  * Move the robot one unit in the direction it is facing.
* `LEFT`
  * Turn the robot 90 degrees to the left.
* `RIGHT`
  * Turn the robot 90 degrees to the right.
* `REPORT`
  * Report the current location and facing of the robot.

## Examples
A number of example input files are provided in the `test-data` directory.