package lib

import (
	"crypto/sha1"
	"fmt"
	"github.com/schmiddim/font-awesome-golang/download"
	gen "github.com/schmiddim/font-awesome-golang/generated" // Ensure this is the correct path
)

func GetIconForString(activityName string) string {
	icons := gen.Icons
	hasher := sha1.New()
	hasher.Write([]byte(activityName))
	hash := hasher.Sum(nil)

	var freeItems []download.Icon

	for _, icon := range icons {
		isFree := false
		if len(icon.Free) > 0 {
			for _, free := range icon.Free {
				if free != "brands" {
					isFree = true
					break
				}
			}
			if isFree {
				freeItems = append(freeItems, icon)
			}
		}
	}
	index := int(hash[0]) % len(freeItems)
	icon := freeItems[index]
	// https://fontawesome.com/icons/wpforms?f=brands&s=solid problem gibt es nicht
	/**
	  <i class={fmt.Sprint("fas fa-", fa.GetIconForString(af.Kind))   }></i>
	<i class="fa-brands fa-wpforms"></i>
	{Changes: []string{"5.1.0"}, Ligatures: []string{}, Search: map[string][]string{"terms":[]string{"LOL", "emoticon", "face"}}, Styles: []string{"solid", "regular"}, Unicode: "f588", Label: "Face With Tears of Joy", Voted: false, Free: []string{"solid", "regular"}, Name: "grin-tears"},

	<i class="fas fa-grin-tears"></i>
	*/
	str := fmt.Sprintf("fas fa-%s", icon.Name)
	return str
}
