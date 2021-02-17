package new

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/go-kratos/kratos/cmd/kratos/internal/base"
)

const (
	serviceLayoutURL = "https://github.com/go-kratos/kratos-layout.git"
)

// Project is a project template.
type Project struct {
	Name string
}

// Generate generate template project.
func (p *Project) Generate(ctx context.Context, dir string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", p.Name)
	}
	fmt.Printf("Creating service %s\n", p.Name)
	repo := base.NewRepo(serviceLayoutURL)

	if err := repo.CopyTo(ctx, to, p.Name, []string{".git", ".github"}); err != nil {
		return err
	}
	os.Rename(
		path.Join(to, "cmd", "server"),
		path.Join(to, "cmd", p.Name),
	)
	return nil
}
