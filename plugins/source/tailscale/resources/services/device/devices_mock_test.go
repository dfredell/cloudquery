package device

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func createDevices(mux *httprouter.Router) error {
	var device tailscale.Device
	if err := faker.FakeObject(&device); err != nil {
		return err
	}

	mux.GET("/api/v2/tailnet/:tailnet/devices", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := map[string][]*tailscale.Device{
			"devices": {&device},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var route tailscale.DeviceRoutes
	if err := faker.FakeObject(&route); err != nil {
		return err
	}

	mux.GET("/api/v2/device/:device/routes", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(route)
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestDevices(t *testing.T) {
	client.MockTestHelper(t, Devices(), createDevices)
}
