package models

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //for postgres
	"github.com/revel/revel"

	eu "github.com/streetcom/lib/errutil"
)

const (
	MAX_GEO_LEVEL = 6 //max number of levels in geotree
)

var (
	// Dbm is db handle
	Dbm *gorp.DbMap
)

func init() {
	initDB()
	updateModels()
}

func initDB() {
	var dbInfo string

	dbLocal := false
	if dbLocal {
		revel.INFO.Println("RUNNING IN DEV MODE. DATABASE WILL BE LOCAL")
		dbInfo = "port=5432 user=postgres password=password dbname=StreetcomLudhiana sslmode=disable"
	} else { //Test database
		revel.INFO.Println("RUNNING IN PROD MODE. DATABASE WILL BE RDS")
		//dbInfo = "host=sc-dev-new.cqwf1pvghoch.us-west-2.rds.amazonaws.com user=scdevadmin password=Ch!pm0nk18 " +
		//	"dbname=LudhianaMod sslmode=disable"
		dbInfo = "host=chipmonkdb.ckidpyxivgoi.ap-south-1.rds.amazonaws.com user=raghavkm password=ch!pMonk " +
			"dbname=StreetcomLudhiana sslmode=disable"
		//dbInfo = "host=stcomm-ludh-dev-db.ca3dsakdzuj7.ap-south-1.rds.amazonaws.com user=SCH0019_RW password=zNk78g7Vuv " +
		//"dbname=STComm-Ludh-QA-DB sslmode=disable"
	}

	// } else {
	// 	revel.INFO.Println("RUNNING IN PROD MODE. DATABASE WILL BE RDS")
	// 	dbInfo = "host=scdemoall.cqwf1pvghoch.us-west-2.rds.amazonaws.com user=havellsDbAdmin password=testD3m0.All " +
	// 		"dbname=scLudhiana sslmode=disable"
	// }

	Db, err := sql.Open("postgres", dbInfo)
	if Db == nil || err != nil {
		revel.ERROR.Println("could not connect to postgres", dbInfo)
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: Db, Dialect: gorp.PostgresDialect{}}
}

func updateModels() {
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}
	t := Dbm.AddTableWithName(Users{}, "users").SetKeys(true, "Id")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Username": 50,
		"Name":     100,
	})

	t = Dbm.AddTableWithName(Light{}, "light")
	t = Dbm.AddTableWithName(Report{}, "report").SetKeys(true, "Id")

	//Create geo tree tables  @handle dynamically
	t = Dbm.AddTableWithName(LevelRoot{}, "level1")

	//Add child table
	t = Dbm.AddTableDynamic(&Level2, "").SetKeys(true, "Id")
	t = Dbm.AddTableDynamic(&Level3, "").SetKeys(true, "Id")
	t = Dbm.AddTableDynamic(&Level4, "").SetKeys(true, "Id")
	t = Dbm.AddTableDynamic(&Level5, "").SetKeys(true, "Id")
	t = Dbm.AddTableDynamic(&Rphase, "").SetKeys(true, "Id")
	t = Dbm.AddTableDynamic(&Yphase, "").SetKeys(true, "Id")
	t = Dbm.AddTableDynamic(&Bphase, "").SetKeys(true, "Id")
	err1 := Dbm.CreateTablesIfNotExists()
	eu.CheckErr(err1, "Error in creating Dynamic tables")

	//Add GeoLight which is actual light lamp located somewhere
	t = Dbm.AddTableWithName(GeoLight{}, "geo_light")
	//Add Devtech energy misc table
	t = Dbm.AddTableWithName(DTMisc{}, "dtenergymisc")
	//Add light location
	t = Dbm.AddTableWithName(LightLocation{}, "light_location")
	//Add schedule table
	//	t = Dbm.AddTableWithName(WirepassSchedule{}, "wirepas_schedule").SetKeys(true, "ID")
	//	t = Dbm.AddTableWithName(DimSchedule{}, "dim_schedule").SetKeys(false,
	//		"StartTime", "Dim", "Transition")
	//	t = Dbm.AddTableWithName(CCTSchedule{}, "cct_schedule").SetKeys(false,
	//		"StartTime", "Dim", "CCT", "Transition")
	//	t = Dbm.AddTableWithName(RGBSchedule{}, "rgb_schedule").SetKeys(false,
	//		"StartTime", "Red", "Green", "Blue", "Transition")
	//	t = Dbm.AddTableWithName(DevtechSchedule{}, "devtech_schedule").SetKeys(false,
	//		"StartTime", "Dim")
	//	t = Dbm.AddTableWithName(LightSchedule{}, "light_schedule").SetKeys(false, "Level", "ID")

	//t = Dbm.AddTableWithName(LightSchedules{}, "light_schedule")
	t = Dbm.AddTableWithName(LightSchedules{}, "lightSchedules")
	t = Dbm.AddTableWithName(DevtSchedule{}, "devtech_schedule").SetKeys(true, "ID")
	t = Dbm.AddTableWithName(NBSchedule{}, "schedules").SetKeys(true, "Id")
	//Status of different lights
	t = Dbm.AddTableWithName(LightStatus{}, "light_status").SetKeys(false, "ID", "Level")
	t = Dbm.AddTableWithName(LoraStatus{}, "lora_status").SetKeys(true, "ID")

	//Add lorawan lights data table
	t = Dbm.AddTableWithName(UplinkPayload{}, "uplink_payload").SetKeys(false, "Deveui", "ID")
	t = Dbm.AddTableWithName(DownlinkPayloadStatus{}, "downlink_status").SetKeys(false,
		"Deveui", "ID", "TransmissionStatus")
	t = Dbm.AddTableWithName(NodeInfo{}, "node_info").SetKeys(true, "ID")
	t = Dbm.AddTableWithName(JoinInfo{}, "join_info").SetKeys(true, "ID")
	t = Dbm.AddTableWithName(NodeStatus{}, "node_status").SetKeys(false, "ID")
	t = Dbm.AddTableWithName(ElectricMap{}, "electric_map").SetKeys(false, "LightID")
	t = Dbm.AddTableWithName(EnergyNode{}, "energy_receiver").SetKeys(true, "ID")

	Dbm.TraceOn("[gorp]", revel.INFO)
	err := Dbm.CreateTablesIfNotExists()
	eu.CheckErr(err, "Error in creating tables")

	//add some test content
	//bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte("havells"), bcrypt.DefaultCost)
	//user := &User{0, "operator havells", "havells", "havells", bcryptPassword}
}
