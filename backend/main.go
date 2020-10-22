package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/moominzie/user-record/controllers"
	_ "github.com/moominzie/user-record/docs"
	"github.com/moominzie/user-record/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Facultys struct {
	Faculty []Faculty
}

type Faculty struct {
	Fname string
}

type Branchs struct {
	Branch []Branch
}

type Branch struct {
	Brname string
	id     int
}

type Buildings struct {
	Building []Building
}

type Building struct {
	Buname string
}

type Rooms struct {
	Room []Room
}

type Room struct {
	Rname string
	id    int
}

type Employees struct {
	Employee []Employee
}

type Employee struct {
	Employeename  string
	Employeeemail string
	Password      string
}

type Statusts struct {
	Statust []Statust
}

type Statust struct {
	Statustname string
}

type Devices struct {
	Device []Device
}

type Device struct {
	Dname string
	id    int
}

type Symptoms struct {
	Symptom []Symptom
}

type Symptom struct {
	Syname string
}

type StatusRs struct {
	StatusR []StatusR
}

type StatusR struct {
	Sname string
}

type Billingstatuss struct {
	Billingstatus []Billingstatus
}

type Billingstatus struct {
	Billingstatusname string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:g18_repaircom.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewFacultyController(v1, client)
	controllers.NewBranchController(v1, client)
	controllers.NewBuildingController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewUserController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewStatustController(v1, client)
	controllers.NewReturninvoiceController(v1, client)
	controllers.NewDeviceController(v1, client)
	controllers.NewStatusRController(v1, client)
	controllers.NewSymptomController(v1, client)
	controllers.NewRepairInvoiceController(v1, client)
	controllers.NewBillingstatusController(v1, client)
	controllers.NewBillController(v1, client)

	// Set Facultys Data
	facultys := Facultys{
		Faculty: []Faculty{
			Faculty{"Engineering"},
		},
	}
	for _, f := range facultys.Faculty {
		client.Faculty.
			Create().
			SetFname(f.Fname).
			Save(context.Background())
	}

	// Set Branchs Data
	branchs := Branchs{
		Branch: []Branch{
			Branch{"Computer Engineering", 1},
			Branch{"Civil Engineering", 1},
			Branch{"Electrical Engineering", 1},
			Branch{"Aeronautical Engineering", 2},
			Branch{"Engineering Management", 2},
			Branch{"Chemical Engineering", 2},
		},
	}
	for _, br := range branchs.Branch {
		client.Branch.
			Create().
			SetBrname(br.Brname).
			Save(context.Background())
	}

	// Set Buildings Data
	buildings := Buildings{
		Building: []Building{
			Building{"Studying building 1"},
			Building{"Studying building 2"},
		},
	}
	for _, bu := range buildings.Building {
		client.Building.
			Create().
			SetBuname(bu.Buname).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{"B1120", 2001},
			Room{"B1112", 2001},
			Room{"B1215", 2001},
			Room{"B1217", 2001},
			Room{"B2502", 2002},
			Room{"B2510", 2002},
			Room{"B2512", 2002},
			Room{"B2515", 2002},
		},
	}
	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetRname(r.Rname).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"Steve Rogers", "qwwe@gmail.com", "1223"},
			Employee{"Tony Stark", "sddf@gmail.com", "2334"},
			Employee{"Nick Fury", "zxxc@gmail.com", "3445"},
			Employee{"Bruce Banner", "bbbb@gmail.com", "5556"},
			Employee{"Natasha Romanoff", "nnnn@gmail.com", "6667"},
		},
	}
	for _, em := range employees.Employee {
		client.Employee.
			Create().
			SetEmployeename(em.Employeename).
			SetEmployeeemail(em.Employeeemail).
			SetPassword(em.Password).
			Save(context.Background())
	}

	// Set Statusts Data
	statusts := Statusts{
		Statust: []Statust{
			Statust{"In delivery"},
			Statust{"Completed"},
		},
	}
	for _, s := range statusts.Statust {
		client.Statust.
			Create().
			SetStatustname(s.Statustname).
			Save(context.Background())
	}

	// Set Statusrs Data
	statusrs := StatusRs{
		StatusR: []StatusR{
			StatusR{"In process "},
			StatusR{"Done"},
		},
	}
	for _, s := range statusrs.StatusR {
		client.StatusR.
			Create().
			SetSname(s.Sname).
			Save(context.Background())
	}

	// Set Devices Data
	devices := Devices{
		Device: []Device{
			Device{"Acer Predator", 1},
			Device{"Mi Air", 1},
			Device{"Huawei Y800", 1},
			Device{"Lenovo NotePad", 2},
			Device{"Samsung ok10", 2},
			Device{"MacBook Pro 2020", 2},
		},
	}
	for _, d := range devices.Device {
		client.Device.
			Create().
			SetDname(d.Dname).
			Save(context.Background())
	}

	// Set Symptoms Data
	symptoms := Symptoms{
		Symptom: []Symptom{
			Symptom{"Damaged Hardware"},
			Symptom{"Damaged Software"},
		},
	}
	for _, sy := range symptoms.Symptom {
		client.Symptom.
			Create().
			SetSyname(sy.Syname).
			Save(context.Background())
	}

	// Set Billingstatuss Data
	billingstatuss := Billingstatuss{
		Billingstatus: []Billingstatus{
			Billingstatus{"Paid"},
			Billingstatus{"Not Paid"},
		},
	}
	for _, bs := range billingstatuss.Billingstatus {
		client.Billingstatus.
			Create().
			SetBillingstatusname(bs.Billingstatusname).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
