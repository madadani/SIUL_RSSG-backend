package routes

import (
	"net/http"

	"siul-pbj-api/controllers"
	"siul-pbj-api/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine) {
	// Base API Group
	v1 := r.Group("/api/v1")

	// ==========================================
	// 1. PUBLIC ROUTES (No Auth Required)
	// ==========================================
	public := v1.Group("/")
	{
		public.POST("/usulan", controllers.SubmitUsulan)
		public.GET("/usulan/:nomor_tiket", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Detail usulan", "data": gin.H{"nomor_tiket": c.Param("nomor_tiket")}})
		})
		public.GET("/kategori-belanja", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "List kategori", "data": []interface{}{}})
		})
		public.GET("/master-barang", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "List master barang", "data": []interface{}{}})
		})
	}


	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		// Using real AuthGuard
		auth.POST("/logout", middlewares.AuthGuard(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logged out successfully"})
		})
		auth.GET("/me", middlewares.AuthGuard(), func(c *gin.Context) {
			// Now grabbing from JWT claims set by middleware
			userId, _ := c.Get("user_id")
			username, _ := c.Get("username")
			role, _ := c.Get("role")
			c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": userId, "username": username, "role": role}})
		})
	}

	// ==========================================
	// 3. PEP ROUTES (Role: PEP)
	// ==========================================
	pep := v1.Group("/pep")
	pep.Use(middlewares.AuthGuard(), middlewares.RoleGuard("pep"))
	{
		pep.GET("/usulan", controllers.GetUsulanMasuk)
		pep.GET("/usulan/:id", dummyHandler("Detail usulan PEP"))
		pep.POST("/usulan/:id/disposisi", controllers.DisposisiKePPTK)
		pep.PUT("/usulan/:id/anggaran", dummyHandler("Geser Anggaran oleh PEP"))

		// Master Data
		master := pep.Group("/master")
		master.GET("/kategori-belanja", dummyHandler("List Master Kategori"))
		master.GET("/users/pptk", controllers.GetPPTKUsers)
		master.PUT("/users/:id/kewenangan", dummyHandler("Edit Kewenangan User"))
	}

	// ==========================================
	// 4. PPTK ROUTES (Role: PPTK)
	// ==========================================
	pptk := v1.Group("/pptk")
	pptk.Use(middlewares.AuthGuard(), middlewares.RoleGuard("pptk"))
	{
		pptk.GET("/usulan", controllers.GetUsulanPPTK)
		pptk.POST("/usulan/:id/disposisi", controllers.DisposisiKePPKOM)
		pptk.POST("/usulan/:id/return", controllers.ReturnKePEP)
	}

	// ==========================================
	// 5. PPKOM ROUTES (Role: PPKOM)
	// ==========================================
	ppkom := v1.Group("/ppkom")
	ppkom.Use(middlewares.AuthGuard(), middlewares.RoleGuard("ppkom"))
	{
		ppkom.GET("/usulan", controllers.GetUsulanPPKOM)
		ppkom.POST("/usulan/:id/setujui", controllers.SetujuiDanDisposisiKePP)
		ppkom.POST("/usulan/:id/tolak", controllers.TolakDanKembalikanKePPTK)
	}

	// ==========================================
	// 6. PP ROUTES (Role: PP)
	// ==========================================
	pp := v1.Group("/pp")
	pp.Use(middlewares.AuthGuard(), middlewares.RoleGuard("pp"))
	{
		pp.GET("/usulan", controllers.GetUsulanPP)
		pp.POST("/usulan/:id/realisasi", controllers.RealisasiUsulan)
	}
}

// Dummy handlers
func dummyHandler(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": msg, "data": gin.H{"id": c.Param("id")}})
	}
}
