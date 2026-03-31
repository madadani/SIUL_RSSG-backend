package routes

import (
	"net/http"

	"siul-pbj-api/internal/auth"
	"siul-pbj-api/internal/middleware"
	"siul-pbj-api/internal/pep"
	"siul-pbj-api/internal/pp"
	"siul-pbj-api/internal/ppkom"
	"siul-pbj-api/internal/pptk"
	"siul-pbj-api/internal/public"
	"siul-pbj-api/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all routes with dependency-injected handlers
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize handlers
	authHandler := auth.NewHandler(db)
	publicHandler := public.NewHandler(db)
	pepHandler := pep.NewHandler(db)
	pptkHandler := pptk.NewHandler(db)
	ppkomHandler := ppkom.NewHandler(db)
	ppHandler := pp.NewHandler(db)

	// Base API Group
	v1 := r.Group("/api/v1")

	// Serve Static Files for Uploads
	r.Static("/uploads", "./uploads")

	// ==========================================
	// 1. PUBLIC ROUTES (No Auth Required)
	// ==========================================
	pub := v1.Group("/")
	{
		pub.POST("/usulan", publicHandler.SubmitUsulan)
		pub.GET("/usulan", publicHandler.GetUsulanPublikList)
		pub.GET("/usulan/:nomor_tiket", publicHandler.GetDetailUsulanPublik)
		pub.GET("/kategori-belanja", publicHandler.GetMasterKategori)
		pub.GET("/master-barang", publicHandler.GetMasterBarang)
	}

	// ==========================================
	// 2. AUTH ROUTES
	// ==========================================
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", middleware.AuthGuard(), func(c *gin.Context) {
			response.Success(c, "Logged out successfully", nil)
		})
		authGroup.GET("/me", middleware.AuthGuard(), func(c *gin.Context) {
			userId, _ := c.Get("user_id")
			username, _ := c.Get("username")
			role, _ := c.Get("role")
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    gin.H{"id": userId, "username": username, "role": role},
			})
		})
	}

	// ==========================================
	// 3. PEP ROUTES (Role: PEP)
	// ==========================================
	pepGroup := v1.Group("/pep")
	pepGroup.Use(middleware.AuthGuard(), middleware.RoleGuard("pep"))
	{
		pepGroup.GET("/usulan", pepHandler.GetUsulanMasuk)
		pepGroup.GET("/usulan/:id", dummyHandler("Detail usulan PEP"))
		pepGroup.GET("/usulan/detail-anggaran", pepHandler.GetDetailAnggaran)
		pepGroup.POST("/usulan/:id/disposisi", pepHandler.DisposisiKePPTK)
		pepGroup.PUT("/usulan/:id/detail_anggaran", dummyHandler("Kelola Detail Anggaran oleh PEP"))

		// Master Data
		master := pepGroup.Group("/master")
		master.GET("/kategori-belanja", publicHandler.GetMasterKategori)
		master.GET("/nama-barang", publicHandler.GetMasterBarang)
		master.GET("/users/all", pepHandler.GetAllUsers)
		master.GET("/users/pptk", pepHandler.GetPPTKUsers)
		master.PUT("/users/:id/kewenangan", dummyHandler("Edit Kewenangan User"))
	}

	// ==========================================
	// 4. PPTK ROUTES (Role: PPTK)
	// ==========================================
	pptkGroup := v1.Group("/pptk")
	pptkGroup.Use(middleware.AuthGuard(), middleware.RoleGuard("pptk"))
	{
		pptkGroup.GET("/usulan", pptkHandler.GetUsulanPPTK)
		pptkGroup.POST("/usulan/:id/disposisi", pptkHandler.DisposisiKePPKOM)
		pptkGroup.POST("/usulan/:id/return", pptkHandler.ReturnKePEP)
		pptkGroup.GET("/users/ppkom", pptkHandler.GetPPKOMUsers)
	}

	// ==========================================
	// 5. PPKOM ROUTES (Role: PPKOM)
	// ==========================================
	ppkomGroup := v1.Group("/ppkom")
	ppkomGroup.Use(middleware.AuthGuard(), middleware.RoleGuard("ppkom"))
	{
		ppkomGroup.GET("/usulan", ppkomHandler.GetUsulanPPKOM)
		ppkomGroup.POST("/usulan/:id/setujui", ppkomHandler.SetujuiDanDisposisiKePP)
		ppkomGroup.POST("/usulan/:id/tolak", ppkomHandler.TolakDanKembalikanKePPTK)
		ppkomGroup.GET("/users/pp", ppkomHandler.GetPPUsers)
	}

	// ==========================================
	// 6. PP ROUTES (Role: PP)
	// ==========================================
	ppGroup := v1.Group("/pp")
	ppGroup.Use(middleware.AuthGuard(), middleware.RoleGuard("pp"))
	{
		ppGroup.GET("/usulan", ppHandler.GetUsulanPP)
		ppGroup.POST("/usulan/:id/realisasi", ppHandler.RealisasiUsulan)
	}
}

// Dummy handlers for unimplemented endpoints
func dummyHandler(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.Success(c, msg, gin.H{"id": c.Param("id")})
	}
}
