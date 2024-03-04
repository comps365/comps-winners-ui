# comps-winners-ui

Provides a simple but fun User Interface to determine a competition winner. It connects to the `competition-database` and pulls finished competitions and their entries. From there, it uses a random selection alogirthm to pick a winner from the entries. The project does not store any data. It's a "dummy" front end that displays data from a database.

## Picking the winner

To pick the winner we're following these steps:
1. The project pulls the entries from the database and stores them in the browser's memory
2. Then, the entries are shuffled. This is triggered manually by the presenter and can be done as many times as they want. We are using the [Fisher-Yates shuffle](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle)
3. When the winner selection is triggered, we're randomly selecting an integer `n` in this range `0 >= n <= len(entries) - 1)` for 10s, constantly
4. The last randomly-generated `n`, after the 10s are up, is used as an index for the `entries`. The entry on that `index`, is declared the winner

> [!IMPORTANT]
> All the above are happening on the client-side, away from any servers. So it's impossible for us, or anyone else, to tamper with the winner selection. We've built this to be as random, tamper-proof and transparent as we could.

## Running the project
* Copy the provided `.env.example` to a new `.env` file and replace the variables with the correct values
* You can run the project with Go

```sh
go run ./cmd/api/*.go
```

## License
For full transparency and visibiltiy, we've put this repo under the MIT license. You're free to view and use the code as you see fit. For more information, please read the [LICENSE](LICENSE) file.
