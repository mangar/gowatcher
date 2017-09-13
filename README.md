# GoWatcher

Watch changes (remove, create and changes) on a mapped directory and fire a command for each.

## Usage

	go run gowatcher config.json


### Configuration


Create a json file with the content bellow.

- dir: directory that will be wathed
- created / change / delete: entry for each action
- created / change / delete . cmd: command that will be fired for each event (create, change or remove)


```
{
  "dir":"/Users/marciog/Projects/__temp/golang_changes_file_dir/watch",
  "create":{
    "cmd": "touch `pwd`/log/`date '+%Y%m%d-%H%M%S'`_create.txt"
  },
  "change":{
    "cmd": "touch `pwd`/log/`date '+%Y%m%d-%H%M%S'`_change.txt"
  },
  "delete":{
    "cmd":"touch `pwd`/log/`date '+%Y%m%d-%H%M%S'`_delete.txt"
  }
}
```
