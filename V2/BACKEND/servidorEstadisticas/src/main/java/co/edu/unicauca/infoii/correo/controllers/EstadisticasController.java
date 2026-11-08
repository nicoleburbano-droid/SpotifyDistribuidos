package co.edu.unicauca.infoii.correo.controllers;

import co.edu.unicauca.infoii.correo.DTOs.AudioReproduccionDTO;
import co.edu.unicauca.infoii.correo.services.EstadisticasService;
import java.util.List;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/estadisticas")
public class EstadisticasController {
    private final EstadisticasService service;

    public EstadisticasController(EstadisticasService service) {
        this.service = service;
    }

    @GetMapping("/cliente/{clienteId}/audios")
    public List<AudioReproduccionDTO> obtenerAudiosReproducidos(@PathVariable String clienteId) {
        return service.obtenerAudiosReproducidos(clienteId);
    }
}

