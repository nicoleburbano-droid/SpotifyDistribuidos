package co.edu.unicauca.infoii.correo.services;

import co.edu.unicauca.infoii.correo.DTOs.AudioReproduccionDTO;
import co.edu.unicauca.infoii.correo.repository.EstadisticasRepository;
import java.util.List;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;

@Service
public class EstadisticasService {
    private final EstadisticasRepository repository;

    public EstadisticasService(EstadisticasRepository repository) {
        this.repository = repository;
    }

    public void registrarReproduccion(String clienteId, String titulo) {
        if (!StringUtils.hasText(clienteId) || !StringUtils.hasText(titulo)) {
            throw new IllegalArgumentException("clienteId y titulo son obligatorios");
        }
        repository.registrarReproduccion(clienteId.trim(), titulo.trim());
    }

    public List<AudioReproduccionDTO> obtenerAudiosReproducidos(String clienteId) {
        if (!StringUtils.hasText(clienteId)) {
            throw new IllegalArgumentException("clienteId es obligatorio");
        }
        return repository.obtenerPorCliente(clienteId.trim());
    }
}

