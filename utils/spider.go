package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type frame struct {
	version string
	path    []string
	key     string
	branch  string
}

func InitProject(framework_name, module_name string) {
	var urls []string
	reg := regexp.MustCompile(`\w*\.tmpl`)
	dst := GetConf("gos", "").GetString("repository") + "tree/main/"
	f := parseIndex(framework_name, dst)
	dst += framework_name + "/" + f.version + "/"
	wg.Add(len(f.path))
	for cnt := 0; cnt < len(f.path); cnt++ {
		if strings.IndexByte(f.path[cnt], '.') > 0 {
			urls = append(urls, dst+f.path[cnt])
			wg.Done()
			continue
		}
		go func(index int) {
			res, err := http.Get(dst + f.path[index])
			Panic(err)
			data, err := ioutil.ReadAll(res.Body)
			Panic(err)
			tmp := reg.FindAll(data, -1)
			for i := 0; i < len(tmp); i++ {
				urls = append(urls, dst+f.path[index]+"/"+string(tmp[i]))
			}
			wg.Done()
		}(cnt)
	}
	wg.Wait()
	urlsf := RemoveDuplicate(urls)
	wg.Add(len(urlsf))
	for _, url := range urlsf {
		go func(url string) {
			res, err := http.Get(url)
			Panic(err)
			data, err := ioutil.ReadAll(res.Body)
			Panic(err)
			path := url[len(dst)-1:]
			file := reg.Find([]byte(url))
			if len(path) == len(file) {
				getFile(string(data), path, module_name)
			} else {
				MkDirs(path[:len(path)-len(file)])
				getFile(string(data), path, module_name)
			}
			wg.Done()
		}(strings.Replace(url.(string), "tree", "raw", -1))
	}
	wg.Wait()
	Gomod("init", module_name)
	Gomod("tidy")
}

func GetLibary(libary_name string) {
	conf := newConf("packages", "yaml", ".")
	if conf.GetString("scaffold."+libary_name) != "" {
		fmt.Println(libary_name + " exist!")
		return
	}
	var urls []string
	repository_addr := GetConf("gos", "").GetString("repository")
	dst := repository_addr + "tree/main/"
	f := parseIndex(libary_name, dst)
	if f == nil {
		return
	}
	dir := ""
	dir_slice := strings.Split(libary_name, "-")
	for i := 0; i < len(dir_slice)-1; i++ {
		dir += dir_slice[i] + "/"
	}
	dst = strings.Replace(dst+dir+dir_slice[len(dir_slice)-1]+"/"+f.version+"/", "main", f.branch, -1)
	wg.Add(len(f.path))
	for cnt := 0; cnt < len(f.path); cnt++ {
		if strings.IndexByte(f.path[cnt], '.') > 0 {
			urls = append(urls, dst+f.path[cnt])
			wg.Done()
			continue
		} else {
			panic("throw exception")
		}
	}
	wg.Wait()
	urlsf := RemoveDuplicate(urls)
	wg.Add(len(urlsf))
	MkDirs(f.branch + "/" + dir)
	reg := regexp.MustCompile(`\w*\.\w*`)
	for _, url := range urlsf {
		go func(url string) {
			res, err := http.Get(url)
			Panic(err)
			data, err := ioutil.ReadAll(res.Body)
			Panic(err)
			file := reg.Find([]byte(url[len(repository_addr):]))
			fp, err := os.Create(f.branch + "/" + dir + string(file))
			defer fp.Close()
			Panic(err)
			fp.Write(data)
			wg.Done()
		}(strings.Replace(url.(string), "tree", "raw", -1))
	}
	wg.Wait()
	conf.Set("scaffold."+libary_name, f.version)
	conf.WriteConfigAs("./packages.yml")
	Gomod("tidy")
}
