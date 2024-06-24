package maps

import "fmt"

func maps() {
	websites := map[string]string{
		"baidu":  "www.baidu.com",
		"google": "www.google.com",
		"taobao": "www.taobao.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["mathias"])

	delete(websites, "google")
	fmt.Println(websites)

	for index, value := range websites {
		fmt.Println(index, value)
	}
}
