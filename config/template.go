package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/options"
	"github.com/28251536/codeforces-tool/util"
	"github.com/fatih/color"
)

//Add a template
func (c *Config) AddTemplate() error {
	color.Cyan("Add a template")

	//Select language
	type kv struct {
		K, V string
	}
	langs := []kv{}
	for k, v := range client.Langs {
		langs = append(langs, kv{k, v})
	}
	sort.Slice(langs, func(i, j int) bool { return langs[i].V < langs[j].V })
	for _, t := range langs {
		fmt.Printf("%5v: %v\n", t.K, t.V)
	}
	lang := ""
	for {
		lang = options.ChooseString("langs")

		if val, ok := client.Langs[lang]; ok {
			color.Green(val)
			break
		}

		color.Red("Invalid index. Please input again")
	}

	//Set the template source file path
	path := ""
	for {
		path = options.ChooseString("templatepath")
		//Exit
		if path == "" {
			os.Exit(1)
		}

		//Check if the file exists
		ok := util.FileExist(path)
		if ok {
			break
		}

		color.Red("%v is invalid. Please input again: ", path)
	}

	//Set suffix
	tmpSuffix := strings.Fields(options.ChooseString("suffix"))
	tmpSuffix = append(tmpSuffix, strings.Replace(filepath.Ext(path), ".", "", 1))
	suffixMap := map[string]bool{}
	suffix := []string{}
	for _, s := range tmpSuffix {
		if _, ok := suffixMap[s]; !ok {
			suffixMap[s] = true
			suffix = append(suffix, s)
		}
	}

	//Set alias
	alias := ""
	for {
		alias = options.ChooseString("alias")

		if len(alias) > 0 {
			break
		}
		color.Red("Alias can not be empty. Please input again: ")
	}

	//Set script
	beforeScript := ""
	script := ""
	afterScript := ""
	//BeforeScript
	beforeScript = options.ChooseString("beforeScript")
	//Script
	for {
		script = options.ChooseString("script")
		if len(script) > 0 {
			break
		}

		color.Red("Script can not empty. Please input again")
	}
	//AfterScript
	afterScript = options.ChooseString("afterScript")
	c.Template = append(c.Template, CodeTemplate{
		alias, lang, path, suffix,
		beforeScript, script, afterScript,
	})
	return c.save()
}

//Remove a template
func (c *Config) RemoveTemplate() error {
	color.Cyan("Remove a template")

	if len(c.Template) == 0 {
		color.Red("There is no template. Please add one")
		return nil
	}

	for i, template := range c.Template {
		star := " "
		if i == c.Default {
			star = color.New(color.FgGreen).Sprint("*")
		}
		fmt.Printf(`%v%2v: "%v" "%v"`, star, i, template.Lang, template.Path)
		fmt.Println()
	}

	idx := 0
	for {
		idx = options.ChooseIndex()
		if 0 <= idx && idx < len(c.Template) {
			break
		}
		color.Red("Tamplate index is not exist. Please input again")
	}

	c.Template = append(c.Template[:idx], c.Template[idx+1:]...)
	if idx == c.Default {
		c.Default = 0
	} else if idx < c.Default {
		c.Default--
	}
	return c.save()
}

// SetDefaultTemplate set default template index
func (c *Config) SetDefaultTemplate() error {
	color.Cyan("Set default template")
	if len(c.Template) == 0 {
		color.Red("There is no template. Please add one")
		return nil
	}
	for i, template := range c.Template {
		star := " "
		if i == c.Default {
			star = color.New(color.FgGreen).Sprint("*")
		}
		fmt.Printf(`%v%2v: "%v" "%v"`, star, i, template.Lang, template.Path)
		fmt.Println()
	}

	idx := 0
	for {
		idx = options.ChooseIndex()
		if 0 <= idx && idx < len(c.Template) {
			break
		}
		color.Red("Tamplate index is not exist. Please enter again")
	}

	c.Default = idx
	return c.save()
}

// TemplateByAlias return all template which alias equals to alias
func (c *Config) TemplateByAlias(alias string) []CodeTemplate {
	ret := []CodeTemplate{}
	for _, template := range c.Template {
		if template.Alias == alias {
			ret = append(ret, template)
		}
	}
	return ret
}
