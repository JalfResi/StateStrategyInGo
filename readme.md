# State Strategy Pattern in Go
This is an example of the State Strategy pattern implemented in Go. This is based off the PHP example by Sebastian Bergmann [sebastianbergmann/state](https://github.com/sebastianbergmann/state).

## State Machine
The example consists of a Door object, which can be in one of three states; open, closed or locked. There four transitions between these states; open, close, lock and unlock. 

![Door State Machine](https://github.com/JalfResi/)