package lib

import (
	"crypto/sha1"
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
		if len(icon.Free) > 0 {
			freeItems = append(freeItems, icon)
		}
	}
	index := int(hash[0]) % len(freeItems)

	return icons[index].Name
}
