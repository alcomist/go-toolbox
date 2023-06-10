package config

import (
	"fmt"
	"github.com/alcomist/go-framework/file"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

func LoadIni(f string) (*ini.File, error) {
	return ini.ShadowLoad(f)
}

func MustGet(s string) []*ini.Section {

	sections := make([]*ini.Section, 0)

	f, err := LoadIni(file.Home() + "/config.ini")
	if err != nil {
		log.Println(err)
		return sections
	}

	if len(s) == 0 {
		return f.Sections()
	}

	if !f.HasSection(s) {
		log.Fatal(fmt.Errorf("ini file has no section : %v", s))
	}

	for _, section := range f.Sections() {
		name := section.Name()
		if strings.Index(name, s) != -1 {
			sections = append(sections, section)
		}
	}

	return sections
}
