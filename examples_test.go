// Copyright 2014 - anova r&d bvba. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noko_test

import (
	"fmt"

	"github.com/stellentus/go-noko"
)

// To get started with the API, use the New function
// and pass your domain and the Noko V2 API token
func Example() {
	f := noko.New("mycompany", "MyNokoAPIV2Token")

	// once you have the Noko object, just start using the API
	// through one of the ...API() functions
	f.EntriesAPI().ListEntries()
}

// Some functions, like create and edit functions, require
// additional input. This input can be provided through an
// anonymous function to add values to the Inputs object.
func ExampleInputSetter(f noko.Noko) {
	f.EntriesAPI().CreateEntry("2014-12-22", 60, func(i noko.Inputs) {
		// here you can add addtional input data to your API call
		i["description"] = "My neat #development issue"
		i["project_name"] = "Customer Project"
	})
}

// Some functions, like the query functions, require
// additional parameters. These parameters can be provided through
// an anonymous function to add values to the Parameters object.
func ExampleParameterSetter(f noko.Noko) {
	f.EntriesAPI().ListEntries(func(p noko.Parameters) {
		// here you can add addtional parameter data to your API call
		p["from"] = "2014-11-01"
		p["to"] = "2014-11-30"
	})
}

// Basic example for working with project pages
func ExampleProjectsPage(f noko.Noko) {
	page, _ := f.ProjectsAPI().ListProjects()

	// the Projects field contains all the projects on the current page
	for _, project := range page.Projects {
		fmt.Println("Project name is " + project.Name)
	}

	// check if there's a next page and then go and fetch it
	if page.HasNext() {
		page, _ = page.Next()
	}
}

// The ProjectPage also offers an AllProjects method for reading all
// projects on the current and subsequent pages through a channel.
// Subsequent pages will automatically get retrieved when reading through
// the end of the previous page
func ExampleProjectsPage_AllProjects(f noko.Noko) {
	page, _ := f.ProjectsAPI().ListProjects()

	// The AllProjects() method returns a convenient channel to read all projects
	// from the current page as well as any subsequent pages.
	for project := range page.AllProjects() {
		fmt.Println("Project name is " + project.Name)
	}
}
