package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Figures struct {
	Name []string `json:"figuresName"`
}

var figuresData = []Figures{
	{Name: []string{"Cone", "Cube", "Cylinder"}},
}

func getFiguresData(c *gin.Context) {
	c.JSON(http.StatusOK, figuresData)
}

func postFiguresVolumes(c *gin.Context) {
	// var input string

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	result := measure("...")
	c.JSON(http.StatusOK, map[string]string{"result": result})
}

type geometry interface {
	volume() float64
}

type cube struct {
	side float64
}

type cone struct {
	radius, height float64
}

type cylinder struct {
	radius, height float64
}

func (cu cube) volume() float64 {
	exponent := 3.0
	return math.Pow(cu.side, exponent)
}

func (cy cylinder) volume() float64 {
	exponent := 2.0
	return math.Pi * math.Pow(cy.radius, exponent) * cy.height
}

func (co cone) volume() float64 {
	exponent := 2.0
	return math.Pi * math.Pow(co.radius, exponent) * co.height
}

func measure(figure string) string {
	switch figure {
	case cone:

	}
}

func main() {
	r := gin.Default()
	r.GET("/get", getFiguresData)
	r.POST("/post", postFiguresVolumes)
	r.Run() // listen and serve on 0.0.0.0:8080
	//var shape geometry
	fmt.Println("Введите название фигуры:")
	var shapeName string
	fmt.Scan(&shapeName)

	/*switch shapeName {

	case "куб":
		cu := cube{}
		fmt.Println("Введите сторону куба:")
		fmt.Scan(&cu.side)
		fmt.Println("Объем куба равен:")
		shape = cu
	case "конус":
		co := cone{}
		fmt.Println("Введите радиус и высоту конуса:")
		fmt.Scan(&co.radius, &co.height)
		fmt.Println("Объем конуса равен:")
		shape = co
	case "цилиндр":
		cy := cylinder{}
		fmt.Println("Введите радиус и высоту цилиндра:")
		fmt.Scan(&cy.radius, &cy.height)
		fmt.Println("Объем цилиндра равен:")
		shape = cy
	default:
		fmt.Println("Неизвестная фигура")
	*/

	//fmt.Println(measure(shape))
}
