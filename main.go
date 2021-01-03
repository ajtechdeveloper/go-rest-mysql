package main
import (
 "fmt"
 "github.com/gin-gonic/gin"
 _ "github.com/go-sql-driver/mysql" 
 "github.com/jinzhu/gorm"
)
var db *gorm.DB
var err error

type Employee struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Department string `json:"department"`
	Salary int `json:"salary"`
   }

func main() {
 db, _ = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/softwaredevelopercentral?charset=utf8&parseTime=True&loc=Local")
 if err != nil {
    fmt.Println(err)
 }
 defer db.Close()
 db.AutoMigrate(&Employee{})
 r := gin.Default()
 r.GET("/employee/", GetAllEmployees)
 r.GET("/employee/:id", GetEmployee)
 r.POST("/employee", CreateEmployee)
 r.PUT("/employee/:id", UpdateEmployee)
 r.DELETE("/employee/:id", DeleteEmployee)
 r.Run(":8080")
}

func GetAllEmployees(c *gin.Context) {
   var employee []Employee
   if err := db.Find(&employee).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   } else {
      c.JSON(200, employee)
   }
  }
func GetEmployee(c *gin.Context) {
   id := c.Params.ByName("id")
   var employee Employee
   if err := db.Where("id = ?", id).First(&employee).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   } else {
      c.JSON(200, employee)
   }
  }
func CreateEmployee(c *gin.Context) {
   var employee Employee
   c.BindJSON(&employee)
   db.Create(&employee)
   c.JSON(200, employee)
  }
func UpdateEmployee(c *gin.Context) {
 var employee Employee
 id := c.Params.ByName("id")
 if err := db.Where("id = ?", id).First(&employee).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 }
 c.BindJSON(&employee)
 db.Save(&employee)
 c.JSON(200, employee)
}
func DeleteEmployee(c *gin.Context) {
   id := c.Params.ByName("id")
   var employee Employee
   d := db.Where("id = ?", id).Delete(&employee)
   fmt.Println(d)
   c.JSON(200, gin.H{"Employee with ID# " + id: "is deleted"})
  }


