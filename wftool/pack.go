package wftool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// error handling, in general, needs attention

// PackCWL serializes a single cwl byte to json
func PackCWL(cwl []byte, id string, path string, graph *[]string) ([]byte, error) {
	cwlObj := new(interface{})
	yaml.Unmarshal(cwl, cwlObj)
	*cwlObj = nuConvert(*cwlObj, primaryRoutine, id, false, path, graph)
	// printJSON(cwlObj)
	j, err := json.MarshalIndent(cwlObj, "", "    ")
	if err != nil {
		return nil, err
	}
	return j, nil
}

// PackCWLFile ..
// 'path' is relative to prevPath
// except in the case where prevPath is "", and path is absolute
// which is the first call to packCWLFile
//
// at first call
// first try absolute path
// if err, try relative path - path relative to working dir
// if err, fail out
//
// always only handle absolute paths - keep things simple
// assume prevPath is absolute
// and path is relative to prevPath
// construct absolute path of `path`
//
// so:
// 'path' is relative to 'prevPath'
// 'prevPath' is absolute
// 1. construct abs(path)
// 2. ..
func PackCWLFile(path string, prevPath string, graph *[]string) (err error) {

	fmt.Println("path: ", path)
	fmt.Println("prev path: ", prevPath)

	///// here get absolute path of 'path' before reading file ////
	if prevPath != "" {
		var wd string
		if !strings.ContainsAny(prevPath, "/") {
			prevPath = fmt.Sprintf("./%v", prevPath)
		}
		if err = os.Chdir(filepath.Dir(prevPath)); err != nil {
			fmt.Println("err 1: ", err)
			return err
		}
		if err = os.Chdir(filepath.Dir(path)); err != nil {
			fmt.Println("err 2: ", err)
			return err
		}
		if wd, err = os.Getwd(); err != nil {
			fmt.Println("err 3: ", err)
			return err
		}
		path = fmt.Sprintf("%v/%v", wd, filepath.Base(path))
	}

	//////////////////
	cwl, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("err 4: ", err)
		return err
	}
	// copying cwltool's pack id scheme
	// not sure if it's actually good or not
	// but for now, doing this
	id := fmt.Sprintf("#%v", filepath.Base(path))
	// 'path' here is absolute - implies prevPath is absolute
	j, err := PackCWL(cwl, id, path, graph)
	if err != nil {
		fmt.Println("err 5: ", err)
	}
	printJSON(j)
	*graph = append(*graph, string(j))
	return nil
}

// PrintJSON pretty prints a struct as JSON
func printJSON(i interface{}) {
	var see []byte
	var err error
	see, err = json.MarshalIndent(i, "", "   ")
	if err != nil {
		fmt.Printf("error printing JSON: %v", err)
	}
	fmt.Println(string(see))
}
