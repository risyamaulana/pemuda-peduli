package main

import (
	"log"

	"pemuda-peduli/src/common/infrastructure"
	"pemuda-peduli/src/common/infrastructure/web"
	"pemuda-peduli/src/common/interfaces"

	achievementApp "pemuda-peduli/src/achievement/applications"
	adminUserApp "pemuda-peduli/src/admin_user/applications"
	albumApp "pemuda-peduli/src/album/applications"
	authAdminApp "pemuda-peduli/src/auth/admin/applications"
	authUserApp "pemuda-peduli/src/auth/user/applications"
	bannerApp "pemuda-peduli/src/banner/applications"
	beneficariesApp "pemuda-peduli/src/beneficaries/applications"
	beritaApp "pemuda-peduli/src/berita/applications"
	kontakKamiApp "pemuda-peduli/src/kontak_kami/applications"
	partnerKamiApp "pemuda-peduli/src/partner_kami/applications"
	programDonasiApp "pemuda-peduli/src/program_donasi/applications"
	programKamiApp "pemuda-peduli/src/program_kami/applications"
	roleApp "pemuda-peduli/src/role/applications"
	teamApp "pemuda-peduli/src/team/applications"
	tentangKamiApp "pemuda-peduli/src/tentang_kami/applications"
	testimoniApp "pemuda-peduli/src/testimoni/applications"
	tokenApp "pemuda-peduli/src/token/applications"
	tujuanKamiApp "pemuda-peduli/src/tujuan_kami/applications"
	userApp "pemuda-peduli/src/user/applications"

	"github.com/fasthttp/router"
	"github.com/joho/godotenv"
)

// App entry point
func main() {
	// App and Routing Initialization
	var apps = map[string]interfaces.IApplication{}
	router := web.NewRouter()
	initialize(apps, router)

	// Turn on Web API Server
	ws, _ := web.NewWebServer(router, 6969)
	go ws.Listen()

	ctx := infrastructure.CaptureSignal()
	<-ctx.Done()

	ctx = infrastructure.GracefulShutdown()
	// Stop serving API
	ws.Shutdown(ctx)
	<-ctx.Done()
	// Clean up each app resource
	destroy(apps)

	log.Println("🟢 PEMUDA_PEDULI app has been shut down successfully. Asta lavista!")
}

// Setup application modules
func initialize(apps map[string]interfaces.IApplication, r *router.Router) {

	// env Load
	err := godotenv.Load(".sample.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// Register applications to run
	apps["token"] = tokenApp.NewTokenApp()
	apps["auth-admin"] = authAdminApp.NewAuthAdminApp()
	apps["auth-user"] = authUserApp.NewAuthUserApp()
	apps["role"] = roleApp.NewRoleApp()
	apps["admin-user"] = adminUserApp.NewAdminUserApp()
	apps["user"] = userApp.NewUserApp()

	apps["achievement"] = achievementApp.NewAchievementApp()
	apps["album"] = albumApp.NewAlbumApp()
	apps["banner"] = bannerApp.NewBannerApp()
	apps["beneficaries"] = beneficariesApp.NewBeneficariesApp()
	apps["berita"] = beritaApp.NewBeritaApp()
	apps["kontak-kami"] = kontakKamiApp.NewKontakKamiApp()
	apps["partner-kami"] = partnerKamiApp.NewPartnerKamiApp()
	apps["program-donasi"] = programDonasiApp.NewProgramDonasiApp()
	apps["program-kami"] = programKamiApp.NewProgramKamiApp()
	apps["team"] = teamApp.NewTeamApp()
	apps["tentang-kami"] = tentangKamiApp.NewTentangKamiApp()
	apps["testimoni"] = testimoniApp.NewTestimoniApp()
	apps["tujuan-kami"] = tujuanKamiApp.NewTujuanKamiApp()

	for _, v := range apps {
		// log.Printf("Initializing app %s", k)
		v.Initialize(r)
	}
}

func destroy(apps map[string]interfaces.IApplication) {
	for _, v := range apps {
		v.Destroy()
	}
}
