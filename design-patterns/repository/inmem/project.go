package inmem

import (
	"fmt"

	"github.com/ryan0x44/go-resources/design-patterns/repository/project"
)

// ProjectRepository implements the Project repository
type ProjectRepository struct {
	projects  map[int64]string
	highestID int64
}

// FindAll returns all entries in the projects map as a slice of Projects
func (r *ProjectRepository) FindAll() ([]project.Project, error) {
	projects := []project.Project{}
	for projectID, projectName := range r.projects {
		projects = append(projects, project.Project{
			ID:   projectID,
			Name: projectName,
		})
	}
	return projects, nil
}

// Store will add/update the projects map with the given Project
func (r *ProjectRepository) Store(p project.Project) (project.Project, error) {
	if r.projects == nil {
		r.projects = map[int64]string{}
	}
	if p.ID <= 0 {
		// Auto-increment ID for new projects
		r.highestID++
		p.ID = r.highestID
	} else {
		// Check ID exists when updating projects
		if _, exists := r.projects[p.ID]; !exists {
			return p, fmt.Errorf("Could not update project with ID: %d", p.ID)
		}
	}
	r.projects[p.ID] = p.Name
	return p, nil
}

// Delete will remove the Project from the projects map
func (r *ProjectRepository) Delete(p project.Project) error {
	if r.projects == nil {
		return nil
	}
	if _, exists := r.projects[p.ID]; !exists {
		return fmt.Errorf("Could not find project with ID: %d", p.ID)
	}
	delete(r.projects, p.ID)
	return nil
}
