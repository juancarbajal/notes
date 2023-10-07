# Notes 

A notes creator running in console. 

The project use Go Lang and sqlite3. 

## Build 

To create the executable run:

```console
go build 
```


## Usage:

### Create a note 

Create a note using the parameter "-m", you can add title with the parameter "-t".

```console
notes -t "title" -m "my new note"
```

The notes are order by date. 

### Search note 

Search a note by a text "-s".

```console
notes -s "text to search"
```

