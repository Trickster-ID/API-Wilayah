package main

import (
	"encoding/json"
	"net/http"
	"wilayah/config"
	"wilayah/controller"
	"wilayah/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var(
	db				*gorm.DB							= config.SetUpDatabaseConnection()
	provinceRepo	repository.ProvincesRepo			= repository.NewProvincesRepo(db)
	cityRepo		repository.CitiesRepo				= repository.NewCitiesRepo(db)
	districtRepo	repository.DistrictsRepo			= repository.NewDistrictsRepo(db)
	subdistrictRepo	repository.SubdistrictsRepo			= repository.NewSubdistrictsRepo(db)

	provinceCtrl	controller.ProvController			= controller.NewProvController(provinceRepo)
	cityCtrl		controller.CitiesController			= controller.NewCitiesController(cityRepo)
	districtCtrl	controller.DistrictsController		= controller.NewDistrictsController(districtRepo)
	subdistrictCtrl controller.SubdistrictsController	= controller.NewSubdistrictsController(subdistrictRepo)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	apiGroup := r.Group("/api/")
	{ 
		apiGroup.GET("provinsi", provinceCtrl.ProvGetAll)
		apiGroup.GET("provinsi/:id", provinceCtrl.ProvGetById)
		apiGroup.GET("provinsi/c/:contain", provinceCtrl.ProvGetByContain)

		apiGroup.GET("kota", cityCtrl.CitGetAll)
		apiGroup.GET("kota/:id", cityCtrl.CitGetById)
		apiGroup.GET("kota/p/:parent", cityCtrl.CitGetByParent)
		apiGroup.GET("kota/c/:contain", cityCtrl.CitGetByContain)

		apiGroup.GET("kecamatan", districtCtrl.DisGetAll)
		apiGroup.GET("kecamatan/:id", districtCtrl.DisGetById)
		apiGroup.GET("kecamatan/p/:parent", districtCtrl.DisGetByParent)
		apiGroup.GET("kecamatan/c/:contain", districtCtrl.DisGetByContain)

		apiGroup.GET("desa", subdistrictCtrl.SubdisGetAll)
		apiGroup.GET("desa/:id", subdistrictCtrl.SubdisGetById)
		apiGroup.GET("desa/p/:parent", subdistrictCtrl.SubdisGetByParent)
		apiGroup.GET("desa/c/:contain", subdistrictCtrl.SubdisGetByContain)
	}
	r.GET("", mainPage)
	r.Run(":8888")
}

func mainPage(cx *gin.Context) {
	dd := `
    {
        "Provinsi":{
			"GetAll":{
				"link":"/api/provinsi",
				"method":"GET",
				"desc":"get all provinsi",
				"e.g":"localhost:8888/api/provinsi/"
			},
			"GetById":{
				"link":"/api/provinsi/p/:id",
				"method":"GET",
				"desc":"get provinsi by selected id",
				"e.g":"localhost:8888/api/provinsi/p/5"
			},
			"GetByContain":{
				"link":"/api/provinsi/c/:contain",
				"method":"GET",
				"desc":"get provinsi by string contain province name",
				"e.g":"localhost:8888/api/provinsi/c/bant"
			}
		},
		"Kota":{
			"GetAll":{
				"link":"/api/kota",
				"method":"GET",
				"desc":"get all kota",
				"e.g":"localhost:8888/api/kota/"
			},
			"GetById":{
				"link":"/api/kota/:id",
				"method":"GET",
				"desc":"get kota by selected id",
				"e.g":"localhost:8888/api/kota/12"
			},
			"GetByParent":{
				"link":"/api/kota/p/:parent",
				"method":"GET",
				"desc":"get kota by selected provinsi/parent",
				"e.g":"localhost:8888/api/kota/p/3"
			},
			"GetByContain":{
				"link":"/api/kota/c/:contain",
				"method":"GET",
				"desc":"get kota by string contain kota name",
				"e.g":"localhost:8888/api/kota/c/seran"
			}
		},
		"Kecamatan":{
			"GetAll":{
				"link":"/api/kecamatan",
				"method":"GET",
				"desc":"get all kecamatan",
				"e.g":"localhost:8888/api/kecamatan/"
			},
			"GetById":{
				"link":"/api/kecamatan/:id",
				"method":"GET",
				"desc":"get kecamatan by selected id",
				"e.g":"localhost:8888/api/kecamatan/12"
			},
			"GetByParent":{
				"link":"/api/kecamatan/p/:parent",
				"method":"GET",
				"desc":"get kecamatan by selected kota/parent",
				"e.g":"localhost:8888/api/kecamatan/p/11"
			},
			"GetByParent":{
				"link":"/api/kecamatan/c/:contain",
				"method":"GET",
				"desc":"get kecamatan by string contain kecamatan name",
				"e.g":"localhost:8888/api/kecamatan/c/kramatwa"
			}
		},
		"Desa":{
			"GetAll":{
				"link":"/api/desa",
				"method":"GET",
				"desc":"get all desa",
				"e.g":"localhost:8888/api/desa/"
			},
			"GetById":{
				"link":"/api/desa/:id",
				"method":"GET",
				"desc":"get desa by selected id",
				"e.g":"localhost:8888/api/desa/12"
			},
			"GetByParent":{
				"link":"/api/desa/p/:parent",
				"method":"GET",
				"desc":"get desa by selected kecamatan/parent",
				"e.g":"localhost:8888/api/desa/p/12"
			},
			"GetByParent":{
				"link":"/api/desa/c/:contain",
				"method":"GET",
				"desc":"get desa by selected kecamatan/parent",
				"e.g":"localhost:8888/api/desa/c/kramatwat"
			}
		}
    }`
	var obj interface{}
    _ = json.Unmarshal([]byte(dd), &obj)
	cx.JSON(http.StatusOK, obj)
}