# comps-winners-ui

Provides a simple but fun User Interface to determine a competition winner. It connects to the `competition-database` and pulls finished competitions and their entries. From there, it uses a random selection alogirthm to pick a winner from the entries. The project does not store any data. It's a "dummy" front end that displays data from a database.

## Running the project

* Copy the provided `.env.example` to a new `.env` file and replace the variables with the correct values
* You can run the project with Go

```sh
go run ./cmd/api/*.go
```

## License
For full transparency and visibiltiy, we've put this repo under the MIT license. You're free to view and use the code as you see fit. For more information, please read the [LICENSE](LICENSE) file.
