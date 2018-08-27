package main

import (
	"fmt"
	"log"

	"github.com/ryan0x44/go-resources/design-patterns/repository/inmem"
	"github.com/ryan0x44/go-resources/design-patterns/repository/project"
)

func main() {
	// Instantiate in-memory repository implementation
	projectsRepo := inmem.ProjectRepository{}

	// Create some projects
	_, err := projectsRepo.Store(
		project.Project{
			Name: "First project",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	second, err := projectsRepo.Store(
		project.Project{
			Name: "Second project",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Update a project
	second.Name = "My 2nd project"
	_, err = projectsRepo.Store(second)
	if err != nil {
		log.Fatal(err)
	}

	// Get all projects
	all, err := projectsRepo.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Projects:")
	for _, p := range all {
		fmt.Printf("%d) %s\n", p.ID, p.Name)
	}
}
