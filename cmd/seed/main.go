package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
	"github.com/your-org/go-backend-starter/internal/infrastructure/database"
	infraRepo "github.com/your-org/go-backend-starter/internal/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	ctx := context.Background()

	// Initialize repositories
	permissionRepo := infraRepo.NewPermissionRepository()
	roleRepo := infraRepo.NewRoleRepository()
	userRepo := infraRepo.NewUserRepository()
	dormitoryRepo := infraRepo.NewDormitoryRepository()

	// Create permissions
	permissions := []*entity.Permission{
		{ID: uuid.New(), Name: "user:read", Slug: "user-read", Resource: "user", Action: "read", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "user:create", Slug: "user-create", Resource: "user", Action: "create", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "user:update", Slug: "user-update", Resource: "user", Action: "update", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "user:delete", Slug: "user-delete", Resource: "user", Action: "delete", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "dorm:read", Slug: "dorm-read", Resource: "dorm", Action: "read", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "dorm:create", Slug: "dorm-create", Resource: "dorm", Action: "create", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "dorm:update", Slug: "dorm-update", Resource: "dorm", Action: "update", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "dorm:delete", Slug: "dorm-delete", Resource: "dorm", Action: "delete", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	log.Println("Creating permissions...")
	for _, perm := range permissions {
		if err := permissionRepo.Create(ctx, perm); err != nil {
			log.Printf("Failed to create permission %s: %v", perm.Name, err)
		} else {
			log.Printf("Created permission: %s", perm.Name)
		}
	}

	// Create roles
	adminRole := &entity.Role{
		ID:        uuid.New(),
		Name:      "admin",
		Slug:      "admin",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Permissions: []entity.Permission{
			*permissions[0], *permissions[1], *permissions[2], *permissions[3],
			*permissions[4], *permissions[5], *permissions[6], *permissions[7],
		},
	}

	staffRole := &entity.Role{
		ID:        uuid.New(),
		Name:      "staff",
		Slug:      "staff",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Permissions: []entity.Permission{
			*permissions[0], *permissions[4], *permissions[6], // user:read, dorm:read, dorm:update
		},
	}

	userRole := &entity.Role{
		ID:        uuid.New(),
		Name:      "user",
		Slug:      "user",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Permissions: []entity.Permission{
			*permissions[4], // dorm:read
		},
	}

	log.Println("Creating roles...")
	if err := roleRepo.Create(ctx, adminRole); err != nil {
		log.Printf("Failed to create admin role: %v", err)
	} else {
		log.Println("Created admin role")
	}

	if err := roleRepo.Create(ctx, staffRole); err != nil {
		log.Printf("Failed to create staff role: %v", err)
	} else {
		log.Println("Created staff role")
	}

	if err := roleRepo.Create(ctx, userRole); err != nil {
		log.Printf("Failed to create user role: %v", err)
	} else {
		log.Println("Created user role")
	}

	// Create admin user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	adminUser := &entity.User{
		ID:        uuid.New(),
		Email:     "admin@example.com",
		Password:  string(hashedPassword),
		Name:      "Admin User",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Roles:     []entity.Role{*adminRole},
	}

	log.Println("Creating admin user...")
	if err := userRepo.Create(ctx, adminUser); err != nil {
		log.Printf("Failed to create admin user: %v", err)
	} else {
		log.Println("Created admin user: admin@example.com / admin123")
	}

	// Create sample dormitories
	dormitories := []*entity.Dormitory{
		{
			ID:          uuid.New(),
			Name:        "Dormitory A",
			Address:     "123 Main Street",
			Description: "Main dormitory building",
			Capacity:    100,
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "Dormitory B",
			Address:     "456 Oak Avenue",
			Description: "Secondary dormitory building",
			Capacity:    80,
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	log.Println("Creating dormitories...")
	for _, dorm := range dormitories {
		if err := dormitoryRepo.Create(ctx, dorm); err != nil {
			log.Printf("Failed to create dormitory %s: %v", dorm.Name, err)
		} else {
			log.Printf("Created dormitory: %s", dorm.Name)
		}
	}

	log.Println("Seed data created successfully!")
}
