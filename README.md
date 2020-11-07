# ProfileIO Resume

LaTeX based resume generator for [ProfileIO]

## How to use?

While [profileio-resume] is primarily designed to work with [ProfileIO], it can be used standalone or in a different project. Basic usage looks as follows:

```sh
go get github.com/acrlakshman/profileio-resume
```

```go
// main.go
package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/acrlakshman/profileio-resume/profileio"
)

func main() {
	// For sample profile_resume.json, please refer to
	// "profileio-resume -> profileio -> samples -> profile_resume.json
	jsonData, _ := ioutil.ReadFile("./profile_resume.json")
	profileio.ProfileIO(jsonData)

	app := "xelatex"
	if !commandExists(app) {
		app = "pdflatex"
		if !commandExists(app) {
			log.Fatal("Cannot compile resume.tex")
		}
	}

	// profileio-resume, writes tex and cls files to a
	// folder "resume" in the current directory.
	os.Chdir("./resume")
	cmd := exec.Command(app, "resume.tex")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}
```

[ProfileIO]: https://github.com/acrlakshman/profileio
[profileio-resume]: https://github.com/acrlakshman/profileio-resume
