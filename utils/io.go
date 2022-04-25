package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func GetJsonStream(path string) interface{} {
	var json_data interface{}
	fp, err := os.Open(path)
	Panic(err)
	defer fp.Close()
	string_data, err := ioutil.ReadAll(fp)
	Panic(err)
	err = json.Unmarshal(string_data, &json_data)
	Panic(err)
	return json_data
}

func getFile(data, path, module string) {
	var buf bytes.Buffer
	tmpl := template.Must(template.New("").Parse(string(data)))
	fp, err := os.Create(strings.Replace(path, ".tmpl", ".go", -1))
	Panic(err)
	err = tmpl.Execute(&buf, map[string]string{
		"module": module,
	})
	fp.Write(bytesHandler(buf))
	Panic(err)
	fp.Close()
}

func parseIndex(name, path string) *frame {
	rootPath := os.Getenv("GOPATH") + "\\src\\gos\\"
	path = strings.Replace(path+"/packages/index.yml", "tree", "raw", -1)
	res, err := http.Get(path)
	Panic(err)
	MkDirs(rootPath + "packages")
	fp, err := os.Create(rootPath + "packages/index.yml")
	defer fp.Close()
	Panic(err)
	data, err := ioutil.ReadAll(res.Body)
	Panic(err)
	fp.Write(data)
	conf := GetConf("index", "/packages/")
	subpath := conf.GetString(name)
	path = strings.Replace(path, "index.yml", subpath+"/"+name+".json", -1)
	res, err = http.Get(path)
	Panic(err)
	MkDirs(rootPath + "packages/" + subpath)
	fp, err = os.Create(rootPath + "packages/" + subpath + "/" + name + ".json")
	Panic(err)
	data, err = ioutil.ReadAll(res.Body)
	Panic(err)
	fp.Write(data)
	json_data := GetJsonStream(rootPath + "/packages/" + subpath + "/" + name + ".json")
	f := &frame{}
	for k, v := range json_data.(map[string]interface{}) {
		if k == "path" {
			for _, _v := range v.([]interface{}) {
				f.path = append(f.path, _v.(string))
			}
		}
		if k == "version" {
			f.version = v.(string)
		}
	}
	return f
}
