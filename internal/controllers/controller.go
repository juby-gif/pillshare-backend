package controllers

import (
	"context"
	"database/sql"
	// "fmt"
	"net/http"

	cache "github.com/go-redis/cache/v8"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/internal/repositories"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

type Controller struct {
	db            *sql.DB
	UserRepo      models.UserRepo
	DashboardRepo models.DashboardRepo
	MedicalRepo models.MedicalRepo
	VitalsRepo models.VitalsRepo
	cache         *cache.Cache
}

func New(db *sql.DB) *Controller {
	userRepo := repositories.NewUserRepo(db)
	dashboardRepo := repositories.NewDashboardRepo(db)
	medicalRepo := repositories.NewMedicalRepo(db)
	vitalsRepo := repositories.NewVitalsRepo(db)
	cache := utils.RedisCache()
	return &Controller{
		db:            db,
		UserRepo:      userRepo,
		DashboardRepo: dashboardRepo,
		MedicalRepo: medicalRepo,
		VitalsRepo: vitalsRepo,
		cache:         cache,
	}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		URL         []string
		n           int
		authStatus  bool
		sessionUUID string
		accessToken string
	)

	// Check for the context has key `url_split` and `length` that
	// have values not equal to nil
	if ctx.Value("url_split") != nil && ctx.Value("length") != nil {

		// Get the values of keys `URL` & `n`(length) from the context
		// Assign the values to `URL` and `n`
		URL = ctx.Value("url_split").([]string)
		n = ctx.Value("length").(int)

		// Check for the context has key `is_authorized` and `session_uuid` that
		// have values not equal to nil
		if ctx.Value("is_authorized") != nil && ctx.Value("session_uuid") != nil && ctx.Value("access_token") != nil {

			// Get the values of keys `is_authorized` & `session_uuid` from the context
			// Assign the values to `authStatus` and `sessionUUID`
			authStatus = ctx.Value("is_authorized").(bool)
			sessionUUID = ctx.Value("session_uuid").(string)
			accessToken = ctx.Value("access_token").(string)
		}
	}

	// Get User modal from cache
	var user models.User
	if err := c.cache.Get(ctx, sessionUUID, &user); err == nil {

		// For debugging purpose only
		// fmt.Println(user)

		// Saving the `user` from cache to context with key `user`
		ctx = context.WithValue(ctx, "user", user)
		ctx = context.WithValue(ctx, "user_id", user.UserId)
		r = r.WithContext(ctx)
	}
	switch {

	//--------------------------------------VERSION--------------------------------------// 
	case n == 3 && URL[2] == "version" && r.Method == "GET":
		c.getVersion(w, r)
	//--------------------------------------VERSION--------------------------------------// 


	//-----------------------------------REFRESH TOKEN-----------------------------------// 
	case n == 3 && URL[2] == "refresh-token" && r.Method == "GET":
		c.postRefreshToken(w, r, accessToken)
	//-----------------------------------REFRESH TOKEN-----------------------------------//  


	//--------------------------------------GATEWAY--------------------------------------// 
	case n == 3 && URL[2] == "login" && r.Method == "POST":
		c.postLogin(w, r)
	case n == 3 && URL[2] == "register" && r.Method == "POST":
		c.postRegister(w, r)
	//--------------------------------------GATEWAY--------------------------------------// 


	//-------------------------------------DASHBOARD-------------------------------------// 
	case n == 3 && URL[2] == "dashboard-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getDashboard(w, r)
		}
	
	case n == 3 && URL[2] == "dashboard" && r.Method == "POST":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.postDashboard(w, r)
		}
	//-------------------------------------DASHBOARD-------------------------------------// 


	//------------------------------------USER PROFILE-----------------------------------// 
	case n == 3 && URL[2] == "user" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getUserProfile(w, r)
		}
	case n == 3 && URL[2] == "update-user" && r.Method == "PATCH":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.patchUserProfile(w, r)
		}
	//------------------------------------USER PROFILE-----------------------------------//


	//-------------------------------------NAVIGATION------------------------------------//
	case n == 3 && URL[2] == "nav-header" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getNavHeader(w, r, user)
		}
	//-------------------------------------NAVIGATION------------------------------------//


	//----------------------------------MEDICAL RECORDS---------------------------------//
	case n == 3 && URL[2] == "medical-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getMedicalDatum(w, r)
		}
	case n == 3 && URL[2] == "medical-data" && r.Method == "POST":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.postMedicalRecord(w, r)
		}
	case n == 3 && URL[2] == "vitals-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getVitalsRecord(w, r)
		}
	case n == 3 && URL[2] == "vitals-data" && r.Method == "POST":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.postVitalsRecord(w, r)
		}
	case n == 3 && URL[2] == "heart-rate-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getHeartRateRecord(w, r)
		}
	case n == 3 && URL[2] == "blood-pressure-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getBloodPressureRecord(w, r)
		}
	case n == 3 && URL[2] == "body-temperature-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getBodyTemperatureRecord(w, r)
		}	
	case n == 3 && URL[2] == "glucose-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getGlucoseRecord(w, r)
		}	
	case n == 3 && URL[2] == "oxygen-saturation-datum" && r.Method == "GET":
		if authStatus != true {
			utils.GetCORSErrResponse(w, "You are not Authorized!", http.StatusUnauthorized)
		} else {
			c.getOxygenSaturationRecord(w, r)
		}	
		
	//----------------------------------MEDICAL RECORDS---------------------------------//

	
	default:
		http.NotFound(w, r)
	}
	}
