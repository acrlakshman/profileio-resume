package profileio

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
)

func generateTemplateForBasicTemplate(profile *Profile, templateFile string) {
	// Populate sorted section list
	sortedSectionList := GetSortedSectionArray(profile)

	// preProcessProfile data based on themes
	preProcessProfile(profile)

	// Generate template file
	GenerateTemplate(profile, &sortedSectionList, templateFile)
}

func compare(f1, f2 string) bool {
	chunkSize := 1000

	fContents1, err := os.Open(f1)
	defer fContents1.Close()
	if err != nil {
		panic(err)
	}

	fContents2, err := os.Open(f2)
	defer fContents2.Close()
	if err != nil {
		panic(err)
	}

	for {
		byteS1 := make([]byte, chunkSize)
		_, err1 := fContents1.Read(byteS1)

		byteS2 := make([]byte, chunkSize)
		_, err2 := fContents2.Read(byteS2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				panic("Error")
			}
		}

		if !bytes.Equal(byteS1, byteS2) {
			return false
		}
	}
}

func TestProfileIO(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("./samples/profile_resume.json")

	profile := Profile{}

	// Map json data to profile object
	PopulateProfile(&jsonData, &profile)

	// Test basic template
	profile.Config.Theme.Value = "basic"
	templateFile := "./basic.tmpl"
	generateTemplateForBasicTemplate(&profile, templateFile)
	defer os.Remove(templateFile)

	// diff basic.tmpl with ./samples/basic.tmpl
	if !compare("./basic.tmpl", "./samples/basic.tmpl") {
		t.Errorf("generated template (\"basic\") differs from the reference one, either make sure to update the reference template or check the code changes.")
	}

	// Test panther template
	profile.Config.Theme.Value = "panther"
	templateFile = "./panther.tmpl"
	generateTemplateForBasicTemplate(&profile, templateFile)
   defer os.Remove(templateFile)

	// diff panther.tmpl with ./samples/panther(_os).tmpl
   pantherFileRef := "./samples/panther.tmpl"

	if runtime.GOOS == "darwin" {
	   pantherFileRef = "./samples/panther_mac.tmpl"
   }

	if !compare("./panther.tmpl", pantherFileRef) {
		t.Errorf("generated template (\"panther\") differs from the reference one, either make sure to update the reference template or check the code changes.")
	}
}
