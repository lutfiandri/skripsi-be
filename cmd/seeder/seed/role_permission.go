package seed

import (
	"context"
	"time"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type PermissionSeed struct {
	Permission domain.Permission
	RoleIds    []uuid.UUID
}

func getMainSeeds() []any {
	roleAdminId, _ := uuid.Parse(constant.RoleAdminId)
	roleCustomerId, _ := uuid.Parse(constant.RoleCustomerId)
	now := time.Now()

	mainSeeds := []any{
		// PERMISSION ==========================================
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionCreatePermission,
				Group:       constant.PermissionGroupPermission,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionReadPermission,
				Group:       constant.PermissionGroupPermission,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdatePermission,
				Group:       constant.PermissionGroupPermission,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionDeletePermission,
				Group:       constant.PermissionGroupPermission,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},

		// ROLE =================================================
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionCreateRole,
				Group:       constant.PermissionGroupRole,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionReadRole,
				Group:       constant.PermissionGroupRole,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdateRole,
				Group:       constant.PermissionGroupRole,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionDeleteRole,
				Group:       constant.PermissionGroupRole,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},

		// DEVICE TYPE ==========================================
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionCreateDeviceType,
				Group:       constant.PermissionGroupDeviceType,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionReadDeviceType,
				Group:       constant.PermissionGroupDeviceType,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId, roleCustomerId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdateDeviceType,
				Group:       constant.PermissionGroupDeviceType,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionDeleteDeviceType,
				Group:       constant.PermissionGroupDeviceType,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},

		// DEVICE ==============================================
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionCreateDevice,
				Group:       constant.PermissionGroupDevice,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionReadDevice,
				Group:       constant.PermissionGroupDevice,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId, roleCustomerId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdateDevice,
				Group:       constant.PermissionGroupDevice,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleCustomerId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionDeleteDevice,
				Group:       constant.PermissionGroupDevice,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId, roleCustomerId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdateVersionDevice,
				Group:       constant.PermissionGroupDevice,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionAcquireDevice,
				Group:       constant.PermissionGroupDevice,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleCustomerId},
		},

		// OAUTH CLIENT ==============================================
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionCreateOAuthClient,
				Group:       constant.PermissionGroupOAuthClient,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionReadOAuthClient,
				Group:       constant.PermissionGroupOAuthClient,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdateOAuthClient,
				Group:       constant.PermissionGroupOAuthClient,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionDeleteOAuthClient,
				Group:       constant.PermissionGroupOAuthClient,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},

		// OAUTH SCOPE ==============================================
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionCreateOAuthScope,
				Group:       constant.PermissionGroupOAuthScope,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionReadOAuthScope,
				Group:       constant.PermissionGroupOAuthScope,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionUpdateOAuthScope,
				Group:       constant.PermissionGroupOAuthScope,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
		PermissionSeed{
			Permission: domain.Permission{
				Id:          uuid.New(),
				Code:        constant.PermissionDeleteOAuthScope,
				Group:       constant.PermissionGroupOAuthScope,
				Description: "",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			RoleIds: []uuid.UUID{roleAdminId},
		},
	}

	return mainSeeds
}

func SeedRolesAndPermissions(mongo *mongo.Database) {
	mainSeeds := getMainSeeds()

	seedPermissions(mongo, mainSeeds)
	seedRoles(mongo, mainSeeds)
}

func seedPermissions(mongo *mongo.Database, mainSeeds []any) {
	permissions := []any{}
	for _, s := range mainSeeds {
		seed := s.(PermissionSeed)
		permissions = append(permissions, seed.Permission)
	}

	mongo.Collection(domain.PermissionCollection).InsertMany(context.Background(), permissions)
}

func seedRoles(mongo *mongo.Database, mainSeeds []any) {
	permissionMap := make(map[uuid.UUID][]uuid.UUID)
	for _, s := range mainSeeds {
		seed := s.(PermissionSeed)

		for _, roleId := range seed.RoleIds {
			permissionMap[roleId] = append(permissionMap[roleId], seed.Permission.Id)
		}
	}

	roleAdminId, _ := uuid.Parse(constant.RoleAdminId)
	roleCustomerId, _ := uuid.Parse(constant.RoleCustomerId)
	now := time.Now()

	roles := []any{
		domain.Role{
			Id:            roleAdminId,
			Name:          "admin",
			PermissionIds: permissionMap[roleAdminId],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
		domain.Role{
			Id:            roleCustomerId,
			Name:          "customer",
			PermissionIds: permissionMap[roleCustomerId],
			CreatedAt:     now,
			UpdatedAt:     now,
		},
	}

	mongo.Collection(domain.RoleCollection).InsertMany(context.Background(), roles)
}
