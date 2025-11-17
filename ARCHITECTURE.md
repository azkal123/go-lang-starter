# Architecture Documentation

## Hexagonal Architecture Overview

Aplikasi ini menggunakan **Hexagonal Architecture** (Ports & Adapters) yang memisahkan business logic dari infrastruktur eksternal.

## Layer Structure

### 1. Domain Layer (`internal/domain/`)

**Tujuan**: Berisi business logic murni, tidak bergantung pada framework atau database.

#### Entities (`entity/`)
- `User` - Entity user dengan password hashing
- `Role` - Entity role dengan permission management
- `Permission` - Entity permission (user:read, dorm:update, dll)
- `Dormitory` - Entity dormitory
- `UserRole`, `RolePermission`, `UserDormitory` - Junction tables

#### Repository Interfaces (`repository/`)
- `UserRepository` - Interface untuk operasi user
- `RoleRepository` - Interface untuk operasi role
- `PermissionRepository` - Interface untuk operasi permission
- `DormitoryRepository` - Interface untuk operasi dormitory

#### Service Interfaces (`service/`)
- `TokenService` - Interface untuk JWT token operations
- `AuthService` - Interface untuk authentication operations

#### Errors (`errors/`)
- Domain-specific errors yang digunakan di seluruh aplikasi

### 2. Application Layer (`internal/application/`)

**Tujuan**: Mengimplementasikan use cases dan business workflows.

#### Use Cases (`usecase/`)
- `AuthUseCase` - Register, Login, RefreshToken
- `UserUseCase` - CRUD operations untuk user
- `DormitoryUseCase` - CRUD operations untuk dormitory

#### DTOs (`dto/`)
- Request/Response models untuk HTTP layer
- `AuthDTO`, `UserDTO`, `DormitoryDTO`

### 3. Infrastructure Layer (`internal/infrastructure/`)

**Tujuan**: Implementasi konkret dari interfaces yang didefinisikan di domain layer.

#### Database (`database/`)
- `postgres.go` - Database connection dan migration menggunakan GORM

#### Repositories (`repository/`)
- Implementasi konkret dari repository interfaces
- Menggunakan GORM untuk database operations

#### Services (`service/`)
- `jwt_service.go` - Implementasi JWT token service

### 4. Interface Layer (`internal/interfaces/`)

**Tujuan**: Entry point untuk aplikasi, menangani HTTP requests.

#### Handlers (`http/handler/`)
- `auth_handler.go` - HTTP handlers untuk authentication
- `user_handler.go` - HTTP handlers untuk user management
- `dormitory_handler.go` - HTTP handlers untuk dormitory management

#### Middleware (`http/middleware/`)
- `auth_middleware.go` - JWT authentication, permission checking, dormitory guard

#### Router (`http/router/`)
- `router.go` - Route configuration menggunakan Gin

## Data Flow

```
HTTP Request
    ↓
Router (Gin)
    ↓
Middleware (JWT Auth, Permission Check, Guard)
    ↓
Handler
    ↓
Use Case (Business Logic)
    ↓
Repository Interface (Port)
    ↓
Repository Implementation (Adapter)
    ↓
Database (PostgreSQL)
```

## Dependency Rule

**Aturan Penting**: Dependencies hanya boleh mengarah ke dalam (inward).

- Domain layer: **TIDAK** bergantung pada layer lain
- Application layer: Bergantung pada Domain layer
- Infrastructure layer: Bergantung pada Domain layer
- Interface layer: Bergantung pada Application dan Domain layer

```
Interface → Application → Domain
     ↓           ↓
Infrastructure → Domain
```

## Key Design Patterns

### 1. Repository Pattern
- Abstraksi data access layer
- Memudahkan testing dengan mock repository
- Memungkinkan pertukaran database tanpa mengubah business logic

### 2. Dependency Injection
- Use cases menerima dependencies melalui constructor
- Memudahkan testing dan maintainability

### 3. Middleware Pattern
- JWT authentication
- Permission checking
- Dormitory guard

## Testing Strategy

### Unit Tests
- Test use cases dengan mock repositories
- Test domain entities dan business logic

### Integration Tests
- Test repository implementations dengan test database
- Test HTTP handlers dengan test server

### Example Test Structure
```
internal/
  application/
    usecase/
      auth_usecase_test.go
  infrastructure/
    repository/
      user_repository_test.go
  interfaces/
    http/
      handler/
        auth_handler_test.go
```

## Scalability Considerations

1. **Horizontal Scaling**: Stateless design memungkinkan multiple instances
2. **Database**: Repository pattern memudahkan migrasi ke database lain
3. **Caching**: Dapat ditambahkan di repository layer tanpa mengubah use cases
4. **Message Queue**: Dapat ditambahkan sebagai adapter di infrastructure layer

## Security

1. **Password Hashing**: Menggunakan bcrypt di domain entity
2. **JWT Tokens**: Access token (15 min) dan Refresh token (7 days)
3. **Permission System**: Role-based dan permission-based authorization
4. **Guard System**: Kontrol akses ke dormitory tertentu

## Future Enhancements

1. **Caching Layer**: Redis untuk session dan data caching
2. **Event System**: Domain events untuk async processing
3. **API Versioning**: Support multiple API versions
4. **GraphQL**: Alternative to REST API
5. **Microservices**: Split into smaller services jika diperlukan
