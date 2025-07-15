package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4" // go mod tidy เพื่อดึง dependencies ที่จำเป็น
)

type Movies struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Rating string `json:"rating"`
}

var movies = []Movies{
	{Title: "Inception", Year: 2010, Rating: "PG-13", Id: 1},
	{Title: "The Matrix", Year: 1999, Rating: "R", Id: 2},
	{Title: "Interstellar", Year: 2014, Rating: "PG-13", Id: 3},
	{Title: "The Dark Knight", Year: 2008, Rating: "PG-13", Id: 4},
	{Title: "Pulp Fiction", Year: 1994, Rating: "R", Id: 5},
	{Title: "The Shawshank Redemption", Year: 1994, Rating: "R", Id: 6},
	{Title: "Forrest Gump", Year: 1994, Rating: "PG-13", Id: 7},
	{Title: "The Godfather", Year: 1972, Rating: "R", Id: 8},
	{Title: "The Lord of the Rings: The Return of the King", Year: 2003, Rating: "PG-13", Id: 9},
	{Title: "Fight Club", Year: 1999, Rating: "R", Id: 10},
	{Title: "The Social Network", Year: 2010, Rating: "PG-13", Id: 11},
	{Title: "The Avengers", Year: 2012, Rating: "PG-13", Id: 12},
	{Title: "Gladiator", Year: 2000, Rating: "R", Id: 13},
	{Title: "Titanic", Year: 1997, Rating: "PG-13", Id: 14},
	{Title: "Jurassic Park", Year: 1993, Rating: "PG-13", Id: 15},
}

func getHello(c echo.Context) error {
	// This function handles the GET request to the /hello endpoint.
	fmt.Println("Received a request at /hello")
	return c.String(http.StatusOK, "Hello, World!")
}

func handlerFunc(c echo.Context) error {
	// This endpoint demonstrates the use of middleware.
	fmt.Println("Handling request at /middleware")
	response := map[string]any{
		"message": "This is a response from the /middleware endpoint",
		"slice":   []string{"item1", "item2", "item3"},
	}
	// TODO: return map ออกไปเป็น json ได้เลย
	return c.JSON(200, response)
}

func middlewareExample(next echo.HandlerFunc) echo.HandlerFunc {
	// This is a middleware function that can be used to log requests.
	fmt.Println("Middleware: Request received")
	return func(c echo.Context) error {
		fmt.Println("Middleware: Before handling the request")
		err := next(c)
		fmt.Println("Middleware: After handling the request")
		return err
	}
}

func getById(c echo.Context) error {
	// This function handles the GET request to the /:id endpoint.
	id := c.Param("id")
	fmt.Printf("Received a request for ID: %s\n", id)
	response := map[string]string{
		"id":      id,
		"message": "This is a response for the requested ID",
	}
	// TODO: return map ออกไปเป็น json ได้เลย
	return c.JSON(200, response)
}

func GetHelloPrefixApi(c echo.Context) error {
	// This function handles the GET request to the /api/hello endpoint.
	fmt.Println("Received a request at /api/hello")
	return c.String(200, "Hello from prefix API!")
}

func GetByIdPrefixApi(c echo.Context) error {
	// This function handles the GET request to the /api/:id endpoint.
	id := c.Param("id")
	fmt.Printf("Received a request for ID in API: %s\n", id)
	response := map[string]string{
		"id":      id,
		"message": "This is a response for the requested ID in API",
	}
	// TODO: return map ออกไปเป็น json ได้เลย
	return c.JSON(200, response)
}

// http://localhost:8080/movies?year=2010
func getMovies(c echo.Context) error {
	// This function handles the GET request to the /movies endpoint.
	fmt.Println("Received a request at /movies")
	// TODO: return struct ออกไปเป็น json ได้เลย

	y := c.QueryParam("year")

	if y == "" {
		return c.JSON(http.StatusOK, movies)
	}

	// TODO: strconv.Atoi ใช้แปลง string เป็น int
	year, err := strconv.Atoi(y)
	if err != nil {
		fmt.Printf("Error converting year to integer: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid year format"})
	}

	fmt.Printf("Filtering movies by year: %d\n", year)
	var filteredMovies []Movies                                         // TODO:ถ้าประกาศตัวแปรแบบนี้ จะเป็น nil
	fmt.Printf("filteredMovies == nil => %#v\n", filteredMovies == nil) // true

	result := []Movies{}                           // TODO: ถ้าประกาศตัวแปรแบบนี้ จะเป็น empty slice จะไม่เป็น nil
	fmt.Printf("a => %#v\n", result)               // invalid: slices cannot be compared to each other
	fmt.Printf("a == nil => %#v\n", result == nil) // false

	for _, movie := range movies {
		if movie.Year == year {
			result = append(result, movie)
		}
	}

	return c.JSON(http.StatusOK, result)

}

func createMovie(c echo.Context) error {
	// This function handles the POST request to the /movies endpoint.
	var movie Movies
	// TODO: bind the request body to the movie struct
	if err := c.Bind(&movie); err != nil {
		// TODO: return map ออกไปเป็น json ได้เลย
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	movies = append(movies, movie)
	fmt.Printf("Added movie: %+v\n", movie)
	// TODO: return struct ออกไปเป็น json ได้เลย
	return c.JSON(http.StatusCreated, movies)
}

// http://localhost:8080/movies/1?year=2010
func getMoviesById(c echo.Context) error {
	// This function handles the GET request to the /movies/:id endpoint.
	// TODO: get parameter id from the request
	id := c.Param("id")
	fmt.Printf("Received a request for movie ID: %s\n", id)
	year := c.QueryParam("year")
	fmt.Printf("Query parameter year: %s\n", year)

	for _, movie := range movies {
		if fmt.Sprintf("%d", movie.Id) == id && fmt.Sprintf("%d", movie.Year) == year {
			return c.JSON(http.StatusOK, movie)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Movie not found"})
}

func main() {
	// This is the entry point of the application.
	// You can add your application logic here.
	e := echo.New()

	e.GET("/hello", getHello)

	// TODO: parameter ตัวแรกเป็นน path ของ endpoint
	// TODO: ตัวที่สองเป็น handler function ที่จะถูกเรียกเมื่อมี request มาที่ path
	// TODO: ตัวที่สามเป็นต้นไปจะเป็น middleware ที่จะถูกเรียกก่อน handler function
	e.GET("/middleware", handlerFunc, middlewareExample)

	e.GET("/:id", getById)

	// TODO: grouping endpoints with a prefix
	api := e.Group("/api", middlewareExample)
	api.GET("/hello", GetHelloPrefixApi)
	api.GET("/:id", GetByIdPrefixApi)

	moviesGroup := e.Group("/movies", middlewareExample)
	moviesGroup.GET("", getMovies)
	moviesGroup.POST("", createMovie)
	moviesGroup.GET("/:id", getMoviesById)

	port := "8080"

	fmt.Printf("Starting server on port %s...\n", port)

	e.Logger.Fatal(e.Start(":" + port))
}
