# nbtcraft
utilities for editing Minecraft .dat files


### build

```
go build -o GOPATH/bin/nbt2json nbt2json.go
go build -o GOPATH/bin/json2nbt json2nbt.go
```


### example usage

```
./nbt2json -filein level.dat | ./ppfjson.pl > level.dat--Hesperia--json
[edit                                         level.dat--Hesperia--json]
./json2nbt -filein                            level.dat--Hesperia--json  -fileout level.dat--Hesperia--nbt
```

