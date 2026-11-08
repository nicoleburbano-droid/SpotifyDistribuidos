package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Registro struct {
	cliente           *http.Client
	base, id, checkID string
	parar             chan struct{}
	wg                sync.WaitGroup
}
type serviceRegistration struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Address string   `json:"Address"`
	Port    int      `json:"Port"`
	Tags    []string `json:"Tags"`
	Check   struct {
		CheckID                        string `json:"CheckID"`
		TTL                            string `json:"TTL"`
		DeregisterCriticalServiceAfter string `json:"DeregisterCriticalServiceAfter"`
	} `json:"Check"`
}

func Registrar(consulAddr, nombre, host string, puerto int) (*Registro, error) {
	if !strings.Contains(consulAddr, "://") {
		consulAddr = "http://" + consulAddr
	}
	base, err := url.Parse(strings.TrimRight(consulAddr, "/"))
	if err != nil {
		return nil, fmt.Errorf("dirección de Consul inválida: %w", err)
	}
	id := fmt.Sprintf("%s-%s-%d", nombre, host, puerto)
	r := &Registro{cliente: &http.Client{Timeout: 5 * time.Second}, base: base.String(), id: id, checkID: "check-" + id, parar: make(chan struct{})}
	registration := serviceRegistration{ID: id, Name: nombre, Address: host, Port: puerto, Tags: []string{"go", "rest"}}
	registration.Check.CheckID = r.checkID
	registration.Check.TTL = "15s"
	registration.Check.DeregisterCriticalServiceAfter = "1m"
	body, err := json.Marshal(registration)
	if err != nil {
		return nil, err
	}
	if err := r.request(http.MethodPut, "/v1/agent/service/register", body); err != nil {
		return nil, fmt.Errorf("registrando %s en Consul: %w", nombre, err)
	}
	if err := r.request(http.MethodPut, "/v1/agent/check/pass/"+url.PathEscape(r.checkID), nil); err != nil {
		_ = r.request(http.MethodPut, "/v1/agent/service/deregister/"+url.PathEscape(r.id), nil)
		return nil, fmt.Errorf("iniciando health check de %s: %w", nombre, err)
	}
	r.wg.Add(1)
	go r.latido()
	log.Printf("servicio %q registrado en Consul como %s", nombre, id)
	return r, nil
}
func (r *Registro) request(method, path string, body []byte) error {
	req, err := http.NewRequest(method, r.base+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := r.cliente.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Consul respondió %s", resp.Status)
	}
	return nil
}
func (r *Registro) latido() {
	defer r.wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := r.request(http.MethodPut, "/v1/agent/check/pass/"+url.PathEscape(r.checkID), nil); err != nil {
				log.Printf("error renovando health check en Consul: %v", err)
			}
		case <-r.parar:
			return
		}
	}
}
func (r *Registro) Desregistrar() {
	close(r.parar)
	r.wg.Wait()
	if err := r.request(http.MethodPut, "/v1/agent/service/deregister/"+url.PathEscape(r.id), nil); err != nil {
		log.Printf("error desregistrando %s de Consul: %v", r.id, err)
	}
}
