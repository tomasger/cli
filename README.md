# cli

**cli** is a simple application written in Go programming language.
Its purpose is to retreive data from an API, store and display it. For now it implements the most basic functionality, however it is expected to be expanded into a enterprise-scale program, capable or retrieving data from various APIs, storing it in various persistent data storages and displaying the data in various methods.
# Usage
Available parameters are as follows:
```
./cli -l, --logging=[debug|warn]
```
- `-l debug` prints all debugging data to the **info.log** file along with warnings and errors
- `-l warn` prints only warnings and errors to the **info.log** file

Available commands are **login** and **servers**.
### Login
Usage:
```
./cli --username "username" --password "password"
```
This command saves user's credentials in a persistent data storage for further use.
### Servers
Usage:
```
./cli servers [--local, -s, --sort=[best|alphabetical]]
```
**Servers** Displays the server list that was fetched from the API.
Additional parameters are for:
  - `--local` shows the server list from persistent data storage rather than the API
  - `--sort best` sorts the servers in the order from closest to furthest
  - `--sort alphabetical` sorts the servers in the alphabetical order

# External packages
**cli** uses these external packages:
```
github.com/jessevdk/go-flags
github.com/google/go-cmp/cmp
github.com/sirupsen/logrus
github.com/pkg/errors
```
