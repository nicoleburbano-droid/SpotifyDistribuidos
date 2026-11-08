package co.edu.unicauca.infoii.correo.capaControladores;

import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import co.edu.unicauca.infoii.correo.DTOs.EntradaHistorialDTO;
import co.edu.unicauca.infoii.correo.capaServicios.ServicioEstadisticas;

@RestController
public class EstadisticasController {

    private final ServicioEstadisticas servicio;

    public EstadisticasController(ServicioEstadisticas servicio) {
        this.servicio = servicio;
    }

    /** GET /historial/{idUsuario}: audios reproducidos por el usuario y cuántas veces. */
    @GetMapping("/historial/{idUsuario}")
    public List<EntradaHistorialDTO> historial(@PathVariable String idUsuario) {
        System.out.println("[REST] GET /historial/" + idUsuario);
        return servicio.historialDe(idUsuario);
    }
}
