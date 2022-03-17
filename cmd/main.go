package main

import (
	"log"
	"pemuda-peduli/src/common/infrastructure"
	"pemuda-peduli/src/common/infrastructure/db"
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
	hubungiKamiApp "pemuda-peduli/src/hubungi_kami/applications"
	kategoriProgramDonasiApp "pemuda-peduli/src/kategori_program_donasi/application"
	kontakKamiApp "pemuda-peduli/src/kontak_kami/applications"
	menuExtrasApp "pemuda-peduli/src/menu_extras/applications"
	partnerKamiApp "pemuda-peduli/src/partner_kami/applications"
	penggalangDanaApp "pemuda-peduli/src/penggalang_dana/applications"
	programDonasiApp "pemuda-peduli/src/program_donasi/applications"
	programDonasiFundraiserApp "pemuda-peduli/src/program_donasi_fundraiser/applications"
	programDonasiKategoriApp "pemuda-peduli/src/program_donasi_kategori/applications"
	programDonasiRutinApp "pemuda-peduli/src/program_donasi_rutin/applications"
	programIncidentalApp "pemuda-peduli/src/program_incidental/applications"
	programKamiApp "pemuda-peduli/src/program_kami/applications"
	qrisApp "pemuda-peduli/src/qris/applications"
	roleApp "pemuda-peduli/src/role/applications"
	teamApp "pemuda-peduli/src/team/applications"
	teamFlagApp "pemuda-peduli/src/team_flag/applications"
	tentangKamiApp "pemuda-peduli/src/tentang_kami/applications"
	testimoniApp "pemuda-peduli/src/testimoni/applications"
	tokenApp "pemuda-peduli/src/token/applications"
	transactionApp "pemuda-peduli/src/transaction/applications"
	tujuanKamiApp "pemuda-peduli/src/tujuan_kami/applications"
	userApp "pemuda-peduli/src/user/applications"

	"github.com/fasthttp/router"
	"github.com/joho/godotenv"
)

var DB *db.ConnectTo

// db init hardcoded temporary for testing
func init() {
	DB = db.NewDBConnectionFactory(0)
}

// App entry point
func main() {
	// App and Routing Initialization
	apps := map[string]interfaces.IApplication{}
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
	apps["token"] = tokenApp.NewTokenApp(DB)
	apps["auth-admin"] = authAdminApp.NewAuthAdminApp(DB)
	apps["auth-user"] = authUserApp.NewAuthUserApp(DB)
	apps["role"] = roleApp.NewRoleApp(DB)
	apps["admin-user"] = adminUserApp.NewAdminUserApp(DB)
	apps["user"] = userApp.NewUserApp(DB)

	apps["achievement"] = achievementApp.NewAchievementApp(DB)
	apps["album"] = albumApp.NewAlbumApp(DB)
	apps["banner"] = bannerApp.NewBannerApp(DB)
	apps["beneficaries"] = beneficariesApp.NewBeneficariesApp(DB)
	apps["berita"] = beritaApp.NewBeritaApp(DB)
	apps["kontak-kami"] = kontakKamiApp.NewKontakKamiApp(DB)
	apps["hubungi-kami"] = hubungiKamiApp.NewHubungiKamiApp(DB)
	apps["partner-kami"] = partnerKamiApp.NewPartnerKamiApp(DB)
	apps["program-donasi"] = programDonasiApp.NewProgramDonasiApp(DB)
	apps["program-donasi-fundraiser"] = programDonasiFundraiserApp.NewProgramDonasiFundraiserApp(DB)
	apps["program-donasi-rutin"] = programDonasiRutinApp.NewProgramDonasiRutinApp(DB)
	apps["program-donasi-kategori"] = programDonasiKategoriApp.NewProgramDonasiKategoriApp(DB)
	apps["program-kami"] = programKamiApp.NewProgramKamiApp(DB)
	apps["program-incidental"] = programIncidentalApp.NewProgramIncidentalApp(DB)
	apps["team"] = teamApp.NewTeamApp(DB)
	apps["team-flag"] = teamFlagApp.NewTeamFlagApp(DB)
	apps["tentang-kami"] = tentangKamiApp.NewTentangKamiApp(DB)
	apps["testimoni"] = testimoniApp.NewTestimoniApp(DB)
	apps["tujuan-kami"] = tujuanKamiApp.NewTujuanKamiApp(DB)
	apps["menu-extras"] = menuExtrasApp.NewMenuExtrasApp(DB)
	apps["qris"] = qrisApp.NewQrisApp(DB)
	apps["transaction"] = transactionApp.NewTransactionApp(DB)

	apps["penggalang_dana"] = penggalangDanaApp.NewPenggalangDanaApp(DB)
	apps["kategori_program_donasi"] = kategoriProgramDonasiApp.NewKategoriProgramDonasiApp(DB)

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
