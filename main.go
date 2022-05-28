package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	template.New("tt")
	test := &Template{
		Sections:         map[string]string{},
		Params:           []string{},
		TemplateBody:     "",
		TemplateFileName: "root.xml",
		RootSection:      "root",
		ChildTemplates:   map[string]*Template{},
	}
	addfuncchild("001_PT_011", test.ChildTemplates)
	addfuncchild("001_PT_012", test.ChildTemplates)
	addfuncchild("001_PT_013", test.ChildTemplates)
	addfuncchild("001_TT_015", test.ChildTemplates)
	addfuncchild("001_PT_021", test.ChildTemplates)
	addfuncchild("001_PT_022", test.ChildTemplates)
	addfuncchild("001_PT_023", test.ChildTemplates)
	addfuncchild("001_TT_025", test.ChildTemplates)
	addfuncchild("001_PT_031", test.ChildTemplates)
	addfuncchild("001_PT_032", test.ChildTemplates)
	addfuncchild("001_PT_033", test.ChildTemplates)
	addfuncchild("001_TT_035", test.ChildTemplates)
	addfuncchild("001_PT_041", test.ChildTemplates)
	addfuncchild("001_PT_042", test.ChildTemplates)
	addfuncchild("001_PT_043", test.ChildTemplates)
	addfuncchild("001_TT_045", test.ChildTemplates)
	addfuncchild("001_PT_051", test.ChildTemplates)
	addfuncchild("001_PT_052", test.ChildTemplates)
	addfuncchild("001_PT_053", test.ChildTemplates)
	addfuncchild("001_TT_055", test.ChildTemplates)
	addfuncchild("001_PT_061", test.ChildTemplates)
	addfuncchild("001_PT_062", test.ChildTemplates)
	addfuncchild("001_PT_063", test.ChildTemplates)
	addfuncchild("001_TT_065", test.ChildTemplates)
	addfuncchild("001_AT_014", test.ChildTemplates)
	addfuncchild("001_AT_024", test.ChildTemplates)
	addfuncchild("001_AT_034", test.ChildTemplates)
	addfuncchild("001_AT_044", test.ChildTemplates)
	addfuncchild("001_AT_054", test.ChildTemplates)
	addfuncchild("001_AT_064", test.ChildTemplates)
	addfuncchild("001_AT_013", test.ChildTemplates)
	test.ReadBody("root")
	os.WriteFile("result.tpj", []byte(test.Sections["root"]), 0666)
	fmt.Print(NEXTMAP)
}
func addfuncchild(name string, children map[string]*Template) {
	child := &Template{
		Sections: map[string]string{},
		Params: []string{
			"NEXT NW_ID 1 1 NW_ID_1",
			"NEXT NW_ID 1 1 NW_ID_2",
			"NEXT NW_ID 1 1 NW_ID_3",
			"NEXT NW_ID 1 1 NW_ID_4",
			"NEXT NW_ID 1 1 NW_ID_5",
			"NEXT NW_ID 1 1 NW_ID_6",
			"NEXT NW_ID 1 1 NW_ID_7",
			"NEXT NW_ID 1 1 NW_ID_8",
			"NEXT NW_ID 1 1 NW_ID_9",
			"NEXT NW_ID 1 1 NW_ID_10",
			"NEXT NW_ID 1 1 NW_ID_11",
			"NEXT NW_ID 1 1 NW_ID_12",
			"NEXT NW_ID 1 1 NW_ID_13",
			"NEXT NW_ID 1 1 NW_ID_14",
			"NEXT CONST 0 1 CONST_1",
			"NEXT CONST 0 1 CONST_2",
			"NEXT CONST 0 1 CONST_3",
			"NEXT CONST 0 1 CONST_4",
			"NEXT FL_H3 10511 100 FL_H3",
			"NEXT FL_H2 10513 100 FL_H2",
			"NEXT FL_H1 10515 100 FL_H1",
			"NEXT FL_L3 10517 100 FL_L3",
			"NEXT FL_L2 10519 100 FL_L2",
			"NEXT FL_L1 10521 100 FL_L1",
			"NEXT AI_H3 10523 100 AI_H3",
			"NEXT AI_L3 10525 100 AI_L3",
			"NEXT FL_IMV 10527 100 FL_IMV",
			"NEXT FL_SET 10529 100 FL_SET",
			"NEXT RC 10531 100 RC",
			"NEXT AI 10533 100 AI",
			"NEXT RS 10535 100 RS",
			"NEXT FL 10537 100 FL",
			"NEXT SCAL 18610 9 SCAL",
			"NEXT SCAL_RESULT 18510 4 SCAL_RESULT",
		},
		TemplateFileName: "ai.xml",
		RootSection:      "Network",
		ChildTemplates: map[string]*Template{
			"Const": {
				Sections:         map[string]string{},
				Params:           []string{},
				TemplateFileName: "aiconst.xml",
				RootSection:      "Const",
			},
			"ai_conf": {
				Sections: map[string]string{},
				Params: []string{
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_1",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_2",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_3",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_4",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_5",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_6",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_7",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_8",
					"NEXT SCAL_ADDR 48100 1 SCAL_ADDR_9",
					"NEXT SCAL_AI 48000 4 SCAL_AI",
				},
				TemplateFileName: "aielemconf.xml",
				RootSection:      "Elem_config",
			},
		},
	}
	children[name] = child
}

var NEXTMAP map[string]int64 = map[string]int64{}

var NEXT = func(nextID, startval, step string) (result int64) {
	sv, err := strconv.ParseInt(startval, 10, 0)
	if err != nil {
		return 0
	}
	stp, err := strconv.ParseInt(step, 10, 0)
	if err != nil {
		return 0
	}
	_, ok := NEXTMAP[nextID]
	if !ok {
		NEXTMAP[nextID] = sv
	}
	result = NEXTMAP[nextID]
	NEXTMAP[nextID] += stp
	return
}
var ADD = func(val, offset string) (result int64) {
	v, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return 0
	}
	o, err := strconv.ParseInt(offset, 10, 0)
	if err != nil {
		return 0
	}
	result = v + o
	return
}

type Template struct {
	Sections         map[string]string
	Params           []string
	ChildTemplates   map[string]*Template
	TemplateFileName string
	TemplateBody     string
	RootSection      string
	DisplayName      string
}

func (t *Template) CalcParams() {
	for i := range t.Params {
		s := strings.Split(t.Params[i], " ")
		switch s[0] {
		case "NEXT":
			if len(s) == 5 {
				t.Sections[s[4]] = fmt.Sprint(NEXT(s[1], s[2], s[3]))

			}
		}
	}
}
func (t *Template) ReadBody(DisplayName string) error {
	t.DisplayName = DisplayName
	t.CalcParams()
	b, err := os.ReadFile(t.TemplateFileName)
	if err == nil {
		t.TemplateBody = string(b)

		fmt.Printf("File read %s\n", t.RootSection)
	}
	for key := range t.ChildTemplates {
		err := t.ChildTemplates[key].ReadBody(key)
		if err != nil {
			fmt.Printf("ERROR3!!!%T %s\n", err, err)
			return err
		}
		for ikey := range t.ChildTemplates[key].Sections {
			t.Sections[ikey] = t.Sections[ikey] + t.ChildTemplates[key].Sections[ikey]
		}
	}
	tem, err := template.New(t.RootSection).Parse(t.TemplateBody)
	if err != nil {
		fmt.Printf("ERROR1!!!%T %s\n", err, err)
		return err
	}
	fmt.Printf("Genered template %s\n", t.RootSection)
	buff := new(bytes.Buffer)
	err = tem.Execute(buff, t)
	if err != nil {
		fmt.Printf("ERROR2!!!%T %s\n", err, err)
	}
	fmt.Printf("Added to Sections %s\n", t.RootSection)
	t.Sections[t.RootSection] = buff.String()
	return nil
}
