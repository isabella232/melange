package app

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	adErrors "airdispat.ch/errors"
	"airdispat.ch/message"
	"airdispat.ch/server"
	"airdispat.ch/wire"

	"getmelange.com/app/controllers"
	"getmelange.com/app/framework"
	"getmelange.com/app/models"
	"getmelange.com/router"
)

func (p *Server) HandleData(res http.ResponseWriter, req *http.Request) {
	// Load tables and Store
	tables, err := models.CreateTables(p.Settings)
	if err != nil {
		fmt.Println("Error creating tables", err)
		framework.WriteView(framework.Error500, res)
		return
	}
	store := p.Settings

	// Get DAP Client
	id, frameErr := controllers.CurrentIdentityOrError(store, tables["identity"])
	if frameErr != nil {
		framework.WriteView(frameErr, res)
		return
	}

	adId, err := id.ToDispatch(store, "")
	if err != nil {
		fmt.Println("Error getting AdId", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	r := &router.Router{
		Origin: adId,
		TrackerList: []string{
			"localhost:2048",
		},
	}

	path := req.URL.Path
	components := strings.Split(path, "/")[1:]

	// Really need at least two
	if len(components) < 2 {
		return
	}

	user := components[0]
	name := strings.Join(components[1:], "/")

	srv, auth, err := models.GetAddresses(r, &models.Address{
		Alias: user,
	})
	if err != nil {
		fmt.Println("Error getting Author Information", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	conn, err := message.ConnectToServer(srv.Location)
	if err != nil {
		fmt.Println("Error connecting to server", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	txMsg := server.CreateTransferMessage(name, adId.Address, srv, auth)
	txMsg.Data = true

	err = message.SignAndSendToConnection(txMsg, adId, srv, conn)
	if err != nil {
		fmt.Println("Error sending to server", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	enc, err := message.ReadMessageFromConnection(conn)
	if err != nil {
		fmt.Println("Error reading from server", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	by, typ, h, err := enc.Reconstruct(adId, false)
	if err != nil {
		fmt.Println("Error reconstructing message", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	if typ == wire.ErrorCode {
		fmt.Println(adErrors.CreateErrorFromBytes(by, h))
		framework.WriteView(framework.Error500, res)
		return
	} else if typ != wire.DataCode {
		fmt.Println("Got wrong type, expecting DAT, got", typ)
		framework.WriteView(framework.Error500, res)
		return
	}

	dataMessage, err := message.CreateDataMessageFromBytes(by, h)
	if err != nil {
		fmt.Println("Error creating message from bytes", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	if dataMessage.Name != name {
		fmt.Println("Name mismatch. Expected", name, "got", dataMessage.Name)
		framework.WriteView(framework.Error500, res)
		return
	}

	if dataMessage.Filename != "" {
		cd := fmt.Sprintf(`attachment; filename="%s"`, dataMessage.Filename)
		res.Header().Add("Content-Disposition", cd)
	}

	res.Header().Add("Content-Length", strconv.Itoa(int(dataMessage.TrueLength())))

	decrypted, err := dataMessage.DecryptReader(conn)
	if err != nil {
		fmt.Println("Error decrypting data", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	_, err = io.Copy(res, decrypted)
	if err != nil {
		fmt.Println("Error writing data", err)
		framework.WriteView(framework.Error500, res)
		return
	}

	// Well, now we're screwed
	if !dataMessage.VerifyPayload() {
		fmt.Println("Wait! That was totally not real!")
	}
}
