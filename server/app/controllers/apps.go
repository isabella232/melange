package controllers

import (
	"fmt"
	"net/http"

	"getmelange.com/app/framework"
	"getmelange.com/app/packaging"
)

type appRequest struct {
	Repository string
	Id         string
}

type InstallPluginController struct {
	Packager *packaging.Packager
}

func (i *InstallPluginController) Handle(req *http.Request) framework.View {
	r := &appRequest{}
	err := DecodeJSONBody(req, r)
	if err != nil {
		fmt.Println("Couldn't deocde JSON Body", err)
		return framework.Error500
	}

	err = i.Packager.InstallPlugin(r.Repository)
	if err != nil {
		fmt.Println("Error install application", err)
		return framework.Error500
	}

	return &framework.HTTPError{
		ErrorCode: 200,
		Message:   "OK",
	}
}

type UninstallPluginController struct {
	Packager *packaging.Packager
}

func (i *UninstallPluginController) Handle(req *http.Request) framework.View {
	r := &appRequest{}
	err := DecodeJSONBody(req, r)
	if err != nil {
		fmt.Println("Couldn't deocde JSON Body", err)
		return framework.Error500
	}

	err = i.Packager.UninstallPlugin(r.Id)
	if err != nil {
		fmt.Println("Error uninstalling application", err)
		return framework.Error500
	}

	return &framework.HTTPError{
		ErrorCode: 200,
		Message:   "OK",
	}
}

// ServerLists will return the list of Trackers or Servers from
// getmelange.com.
type PluginStoreController struct {
	Packager *packaging.Packager
}

// Handle will decodeProviders from getmelange.com then return them in JSON.
func (s *PluginStoreController) Handle(req *http.Request) framework.View {
	packages, err := s.Packager.DecodeStore()
	if err != nil {
		fmt.Println("Getting store", err)
		return framework.Error500
	}

	return &framework.JSONView{
		Content: packages,
	}
}