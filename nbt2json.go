package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"

	"github.com/landru27/nbt"
)

func main() {
	var err error
	var bufDotDat []byte
	var bufTemp bytes.Buffer
	var rdrTemp *bytes.Reader

	fileDotDat := flag.String("filein", "UNDEFINED", "a .dat file from a Minecraft world")
	flag.Parse()

	// read in the Minecraft .dat file
	bufDotDat, err = ioutil.ReadFile(*fileDotDat)
	panicOnErr(err)
	rdrDotDat := bytes.NewReader(bufDotDat)
	panicOnErr(err)

	// uncompress the data so that we can work with it
	rdrDotDatGZip, err := gzip.NewReader(rdrDotDat)
	panicOnErr(err)
	io.Copy(&bufTemp, rdrDotDatGZip)

	// parse the data out of Minecraft's NBT format into data structures we interact with
	rdrTemp = bytes.NewReader(bufTemp.Bytes())
	dataDotDat, err := nbt.ReadNBTData(rdrTemp, nbt.TAG_NULL, "")
	panicOnErr(err)

	// output the .dat data to JSON
	var bufJSON []byte
	bufJSON, err = json.MarshalIndent(&dataDotDat, "", "  ")
	panicOnErr(err)
	os.Stdout.Write(bufJSON)
}

func panicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}
