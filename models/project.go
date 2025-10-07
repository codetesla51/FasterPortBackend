package main
import (
        "context"
        "encoding/json"
        "errors"
        "fmt"
        "github.com/codetesla51/portBackend/config"
        "github.com/jackc/pgx/v5"
)

type Project struct {
        ID            int      `json:"id"`
        Name          string   `json:"name"`
        Slug          string   `json:"slug"`
        TechStack     []string `json:"tech_stack"`
        DisplayStatus bool     `json:"display_status"`
        Image         string   `json:"image"`
        Description   string   `json:"description"`
}

func GetVisibleProjects(limit, offset int) ([]Project, error) {
        query := `
        SELECT id, name, slug, tech_stack, display_status, image, description
        FROM projects
        WHERE display_status = TRUE
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `

        rows, err := config.DB.Query(context.Background(), query, limit, offset)
        if err != nil {
                return nil, err
        }
        defer rows.Close()

        projects := []Project{}
        for rows.Next() {
                var p Project
                var techStackJSON string
                
                if err := rows.Scan(
                        &p.ID,
                        &p.Name,
                        &p.Slug,
                        &techStackJSON,
                        &p.DisplayStatus,
                        &p.Image,
                        &p.Description,
                ); err != nil {
                        return nil, err
                }
                
                // Parse tech_stack JSON string into array
                if err := json.Unmarshal([]byte(techStackJSON), &p.TechStack); err != nil {
                        // If unmarshal fails, set empty array
                        p.TechStack = []string{}
                }
                
                projects = append(projects, p)
        }

        return projects, nil
}

func GetProjectBySlug(slug string) (*Project, error) {
        var p Project
        var techStackJSON string
        
        query := `
        SELECT id, name, slug, tech_stack, display_status, image, description
        FROM projects
        WHERE slug = $1 AND display_status = true
    `
    
        err := config.DB.QueryRow(context.Background(), query, slug).Scan(
                &p.ID,
                &p.Name,
                &p.Slug,
                &techStackJSON,
                &p.DisplayStatus,
                &p.Image,
                &p.Description,
        )
        if err != nil {
                if errors.Is(err, pgx.ErrNoRows) {
                        return nil, errors.New("project not found")
                }
                return nil, fmt.Errorf("db error: %w", err)
        }

        // Parse tech_stack JSON string into array
        if err := json.Unmarshal([]byte(techStackJSON), &p.TechStack); err != nil {
                // If unmarshal fails, set empty array
                p.TechStack = []string{}
        }

        return &p, nil
}