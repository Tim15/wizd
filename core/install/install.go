package install

import (
	"fmt"
	"os"

	"github.com/wizpkg/wizd/core/config"
)

// Install a package to the specified store
func Install(packageName string, packageVersion string) string {
	store := config.GetConfig("storeLocation")

	fmt.Println("Installing to store: ", store)
	fmt.Println(packageName, packageVersion)

	path := store + "/packages/" + packageName + "/" + packageVersion
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Println(path, "does not exist")
		// os.MkdirAll(path)
	} else {
		fmt.Println("path exists")
	}
	// if (os.e)
	return packageName + packageVersion
}
