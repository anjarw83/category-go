package main

import (
	"encoding/json"
	"fmt"
)

// 2. List Lengkap Berupa Json Pada File category.json
// 3. List Category
type Category struct {
	Name   string `json:"name"`
	Childs []struct {
		Name   string `json:"name"`
		Childs []struct {
			Name string `json:"name"`
		} `json:"childs"`
	} `json:"childs"`
}

// Sampling; Copast Value from category.json for complete Data
const CATEGORY_DATA = `[
	{
		"name": "Buku",
        "childs":
        [
            {
                "name": "Arsitektur & Desain",
                "childs":
                [
                    {
                        "name": "Buku Bangunan"
                    },
                    {
                        "name": "Buku Codes & Standars"
                    }
				]
			},
			{
                "name": "Buku Hukum",
                "childs":
                [
                    {
                        "name": "Buku Gender & Hukum"
                    },
                    {
                        "name": "Buku Hukum Dagang"
                    }
				]
			}
		]
	},
	{
        "name": "Rumah Tangga",
        "childs":
        [
            {
                "name": "Dekorasi",
                "childs":
                [
                    {
                        "name": "Cover Kipas Angin"
                    },
                    {
                        "name": "Cover Kursi"
                    },
                    {
                        "name": "Hiasan Dinding"
                    }
				]
			},
			{
                "name": "Furniture",
                "childs":
                [
                    {
                        "name": "Bedside Table"
                    },
                    {
                        "name": "Cermin Badan"
                    },
                    {
                        "name": "Kasur"
                    }
				]
			}
		]
	}
]`

func main() {
	var category []Category

	Data := []byte(CATEGORY_DATA)
	err := json.Unmarshal(Data, &category)

	if err != nil {
		// if error is not nil
		fmt.Println(err)
	}

	initIndex := [3]int{0, 0, 0}
	showCategories(category, initIndex, 0)
}

func showCategories(data []Category, childIndex [3]int, activeIndex int) {
	total_category := len(data)

	// Exit On Completion
	if childIndex[0] >= total_category {
		return
	}

	if activeIndex == 0 {
		// onLastIndex || Exit On Completeion
		if childIndex[0] == len(data[childIndex[0]].Childs) {
			return
		}
		fmt.Println(data[childIndex[0]].Name)

		// ShiftDown
		if len(data[childIndex[0]].Childs) > 0 {
			activeIndex += 1
		}

		showCategories(data, childIndex, activeIndex)
	} else if activeIndex == 1 {
		// onLastIndex || ShiftUp
		if len(data[childIndex[0]].Childs) == childIndex[activeIndex] {
			childIndex[activeIndex] = 0
			childIndex[activeIndex-1] += 1
			activeIndex -= 1
		} else {
			fmt.Println("   ", data[childIndex[0]].Childs[childIndex[1]].Name)

			// ShiftDown
			if len(data[childIndex[0]].Childs[childIndex[1]].Childs) > 0 {
				activeIndex += 1
			}
		}

		showCategories(data, childIndex, activeIndex)
	} else if activeIndex == 2 {
		fmt.Println("      ", data[childIndex[0]].Childs[childIndex[1]].Childs[childIndex[2]].Name)
		childIndex[activeIndex] += 1

		// onLastIndex || ShiftUp
		if len(data[childIndex[0]].Childs[childIndex[1]].Childs) == childIndex[activeIndex] {
			childIndex[activeIndex] = 0
			childIndex[activeIndex-1] += 1
			activeIndex -= 1
		}
		showCategories(data, childIndex, activeIndex)
	}

}
