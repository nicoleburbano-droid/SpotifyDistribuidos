package co.edu.unicauca.infoii.correo.capaServicios;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import org.springframework.stereotype.Service;

import co.edu.unicauca.infoii.correo.DTOs.EntradaHistorialDTO;

/** Lleva, por usuario, los audios que ha reproducido y cuántas veces. */
@Service
public class ServicioEstadisticas {

    /** usuario -> (tipo|titulo -> entrada) */
    private final Map<String, Map<String, EntradaHistorialDTO>> historial = new ConcurrentHashMap<>();

    public int registrar(String idUsuario, String titulo, String tipo, String fecha) {
        String usuario = (idUsuario == null || idUsuario.isBlank()) ? "anonimo" : idUsuario;
        Map<String, EntradaHistorialDTO> audios = historial.computeIfAbsent(usuario, u -> new ConcurrentHashMap<>());
        EntradaHistorialDTO entrada = audios.computeIfAbsent(tipo + "|" + titulo,
                k -> new EntradaHistorialDTO(titulo, tipo, 0, null));
        synchronized (entrada) {
            entrada.setReproducciones(entrada.getReproducciones() + 1);
            entrada.setUltimaReproduccion(fecha);
            return entrada.getReproducciones();
        }
    }

    public List<EntradaHistorialDTO> historialDe(String idUsuario) {
        List<EntradaHistorialDTO> lista = new ArrayList<>(historial.getOrDefault(idUsuario, Map.of()).values());
        lista.sort(Comparator.comparingInt(EntradaHistorialDTO::getReproducciones).reversed());
        return lista;
    }
}
