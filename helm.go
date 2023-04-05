package main

import (
	"fmt"
	helmclient "github.com/mittwald/go-helm-client"
)

func main() {

	options := helmclient.Options{Namespace: "playground"}
	client, err := helmclient.New(&options)
	if err != nil {
		panic(err.Error())
	}

	releases, err := client.ListDeployedReleases()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Listing helm releases (found %d)\n", len(releases))
	for _, release := range releases {
		fmt.Printf("- %s\n", release.Name)
	}

	fmt.Printf("[dry-run] Uninstalling releases (found %d)\n", len(releases))
	for _, release := range releases {

		spec := helmclient.ChartSpec{
			Namespace:   "playground",
			ReleaseName: release.Name,
			DryRun:      true,
		}
		err = client.UninstallRelease(&spec)
		if err != nil {
			panic(err.Error())
		}
	}
}
