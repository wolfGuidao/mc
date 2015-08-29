// +build ignore

/*
 * Minio Client (C) 2014, 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Version struct {
	Date string
}

func writeVersion(version Version) error {
	var versionTemplate = `// --------  DO NOT EDIT --------
// This file is autogenerated by genversion.go during the release process.

package main

// Version autogenerated
const Version = {{if .Date}}"{{.Date}}"{{else}}""{{end}}
`
	t := template.Must(template.New("version").Parse(versionTemplate))
	versionFile, err := os.OpenFile("version.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer versionFile.Close()
	err = t.Execute(versionFile, version)
	if err != nil {
		return err
	}
	return nil
}

func genVersion() {
	t := time.Now().UTC()
	date := t.Format(time.RFC3339Nano)
	// Tag is of following format
	//
	//   RELEASE.[WeekDay]-[Month]-[Day]-[Hour]-[Min]-[Sec]-GMT-[Year]
	//
	tag := fmt.Sprintf(
		"RELEASE.%s-%s-%02d-%02d-%02d-%02d-GMT-%d",
		t.Weekday().String()[0:3],
		t.Month().String()[0:3],
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Year())
	fmt.Println("Release-Tag: " + tag)
	fmt.Println("Release-Version: " + t.Format(http.TimeFormat))
	version := Version{Date: date}
	err := writeVersion(version)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println("Successfully generated ‘version.go’")
}

func main() {
	genVersion()
}
