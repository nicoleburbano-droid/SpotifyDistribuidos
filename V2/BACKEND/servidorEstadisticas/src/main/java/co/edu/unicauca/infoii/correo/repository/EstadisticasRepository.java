package co.edu.unicauca.infoii.correo.repository;

import co.edu.unicauca.infoii.correo.DTOs.AudioReproduccionDTO;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.io.IOException;
import java.nio.file.AtomicMoveNotSupportedException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardCopyOption;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Repository;

@Repository
public class EstadisticasRepository {
    private final ObjectMapper json = new ObjectMapper();
    private final Path archivo;
    private final Map<String, Map<String, Long>> reproducciones = new HashMap<>();

    public EstadisticasRepository(@Value("${estadisticas.archivo:./datos/estadisticas.json}") String rutaArchivo) {
        archivo = Path.of(rutaArchivo);
        cargar();
    }

    public synchronized void registrarReproduccion(String clienteId, String titulo) {
        Map<String, Long> audiosCliente = reproducciones.computeIfAbsent(clienteId, clave -> new HashMap<>());
        Long anterior = audiosCliente.put(titulo, audiosCliente.getOrDefault(titulo, 0L) + 1);
        try {
            guardar();
        } catch (RuntimeException e) {
            if (anterior == null) audiosCliente.remove(titulo);
            else audiosCliente.put(titulo, anterior);
            if (audiosCliente.isEmpty()) reproducciones.remove(clienteId);
            throw e;
        }
    }

    public synchronized List<AudioReproduccionDTO> obtenerPorCliente(String clienteId) {
        return reproducciones.getOrDefault(clienteId, Map.of()).entrySet().stream()
                .map(entrada -> new AudioReproduccionDTO(entrada.getKey(), entrada.getValue()))
                .sorted(Comparator.comparing(AudioReproduccionDTO::titulo, String.CASE_INSENSITIVE_ORDER))
                .toList();
    }

    private void cargar() {
        if (!Files.exists(archivo)) return;
        try {
            Map<String, Map<String, Long>> almacenadas = json.readValue(
                    archivo.toFile(), new TypeReference<>() {});
            if (almacenadas != null) almacenadas.forEach((cliente, audios) ->
                    reproducciones.put(cliente, new HashMap<>(audios)));
        } catch (IOException e) {
            throw new IllegalStateException("No se pudo cargar el historial de reproducciones desde " + archivo, e);
        }
    }

    private void guardar() {
        Path temporal = archivo.resolveSibling(archivo.getFileName() + ".tmp");
        try {
            Files.createDirectories(archivo.toAbsolutePath().getParent());
            json.writeValue(temporal.toFile(), reproducciones);
            try {
                Files.move(temporal, archivo, StandardCopyOption.ATOMIC_MOVE, StandardCopyOption.REPLACE_EXISTING);
            } catch (AtomicMoveNotSupportedException e) {
                Files.move(temporal, archivo, StandardCopyOption.REPLACE_EXISTING);
            }
        } catch (IOException e) {
            try { Files.deleteIfExists(temporal); } catch (IOException ignored) { }
            throw new IllegalStateException("No se pudo guardar el historial de reproducciones en " + archivo, e);
        }
    }
}
