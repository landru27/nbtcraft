package main

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/landru27/nbt"
)

func main() {
	var err error
	var bufJSON []byte
	var nbtJSON nbt.NBT
	var bufTemp bytes.Buffer
	var bufGZip bytes.Buffer

	fileJSON := flag.String("filein", "UNDEFINED", "a JSON file made from a .dat file from a Minecraft world")
	fileDotDat := flag.String("fileout", "UNDEFINED", "a .dat file to create for use in a Minecraft world")
	flag.Parse()

	// read in the JSON-of-a-.dat file
	bufJSON, err = ioutil.ReadFile(*fileJSON)
	panicOnErr(err)
	// we unmarshal with Decoder and UseNumber so that Minecraft's long ints are handled correctly
	d := json.NewDecoder(bytes.NewReader(bufJSON))
	d.UseNumber()
	err = d.Decode(&nbtJSON)
	panicOnErr(err)

	// parse the data out of our internal datatype format into Minecraft's NBT format
	err = nbt.WriteNBTData(&bufTemp, &nbtJSON)
	panicOnErr(err)

	// compress the data so that Minecraft can work with it
	wtrDotDatGZip := gzip.NewWriter(&bufGZip)
	panicOnErr(err)
	wtrDotDatGZip.Write(bufTemp.Bytes())
	wtrDotDatGZip.Close()

	// output the NBT data to a .dat file
	fh, err := os.Create(*fileDotDat)
	panicOnErr(err)
	defer fh.Close()
	err = binary.Write(fh, binary.BigEndian, bufGZip.Bytes())
	panicOnErr(err)
}

func panicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}
