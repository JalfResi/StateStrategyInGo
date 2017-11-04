# State Strategy Pattern in Go
This is an example of the State Strategy pattern implemented in Go. This is based off the PHP example by Sebastian Bergmann [sebastianbergmann/state](https://github.com/sebastianbergmann/state).

## State Machine
The example consists of a Door object, which can be in one of three states; open, closed or locked. There four transitions between these states; open, close, lock and unlock. 

![Door State Machine](https://github.com/JalfResi/StateStrategyInGo/blob/master/DoorStateDiagram.png)

## Implementation Benefits
The Strategy Pattern allows us to encapsulate the actions of each state transition within the State objects themselves. The State Pattern allows us to modify the behaviour of the encapsulating object. Th combination of the two allow us to keep the data of an object separate from the actions that modify that data using a State Machine to ensure that actions can only be executed when the object is in certain states.

For strict State controlled actions, this combination of patterns can simplify things, especially if implementing commands via net/rpc:

```golang
func OpenDoorCommand() error {
    door, err := repo.Load()
        if err != nil {
        return err
    }

    if door.Can(OpenDoorState{}) {
        err := door.Open()
        if err != nil {
            return err
        }

        if door.IsOpen() {
            err := repo.Save(door)
            if err != nil {
                return err
            }
        }
    }
}
```

## References
For more information on the [Strategy Pattern](https://en.wikipedia.org/wiki/Strategy_pattern), see Strategy (315) in the Gang of Four authored [Design Patterns](https://en.wikipedia.org/wiki/Design_Patterns) book

For more information on the [State Pattern](https://en.wikipedia.org/wiki/State_pattern), see State (305) in the Gang of Four authored [Design Patterns](https://en.wikipedia.org/wiki/Design_Patterns) book

For the original PHP example that inspired this, see [sebastianbergmann/state](https://github.com/sebastianbergmann/state)