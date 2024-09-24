package main

import (
	//"fmt"
	//"bufio"
	//"fmt"
	"fmt"
	"net/http"
	//"strings"

	"github.com/labstack/echo/v4"
)

type foodRecipe struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Time        string   `json:"time"`
	Flavour     string   `json:"flavour"`
	Category    string   `json:"category"`
	Dietary     string   `json:"dietary"`
}
/*
func home(c echo.Context) error {

	return c.String(http.StatusOK, "What's Cookin?\nfind yummy foods\nshare your art")
}*/
func getRecipes(c echo.Context) error {

	Recipe := foodRecipe{
		Name:        "milk tea",
		Ingredients: []string{"tea leaves: 2tbsp", "sugar :2tbsp", "milk: 100ml"},
		Time:        "5 minutes",
		Flavour:     "sweet",
		Category:    "drinks",
	}
	Recipe2 := foodRecipe{
		Name:        "buff momo",
		Ingredients: []string{"onion", "buffkeema", "flour/ dough"},
		Time:        "1 hour",
		Flavour:     "spicy/meatful",
		Category:    "food",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Newest Food Recipies",
		"recipe":  Recipe,
		"recipe2": Recipe2,
	})
}


func submitHandler(c echo.Context) error {
	recipie := c.FormValue("foodSearch")
	response:= "Dish: "+ recipie
	return c.String(http.StatusOK,response)
}

func food(c echo.Context) error {
	var recipe = []foodRecipe {
		 {
			Name : "buff momo",
			Ingredients : []string{"keema","flour","oil"},
			Time: "1 hour",
			Category : "dish",
		},

	}
	response := fmt.Sprintf("name: %s, ingredients: %v, time: %s, category: %s",
recipe[0].Name, recipe[0].Ingredients, recipe[0].Time, recipe[0].Category)

	return c.String(http.StatusOK,response)
}

func main() {
	e := echo.New()

	e.Static("/", "index.html")
	e.GET("/recipies", getRecipes)
	e.POST("/submit",submitHandler)
	e.Logger.Fatal(e.Start(":6969"))

}
