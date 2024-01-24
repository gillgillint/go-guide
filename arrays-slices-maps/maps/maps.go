package maps

import "fmt"

func main() {
	website := map[string]string{
		"google":             "https://google.com",
		"Amazon Web Service": "https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Amazon Web Service"])

	website["LikedIn"] = "https://linkedin.com"
	fmt.Println(website)

	delete(website, "Amazon Web Service")
	fmt.Println(website)
}
