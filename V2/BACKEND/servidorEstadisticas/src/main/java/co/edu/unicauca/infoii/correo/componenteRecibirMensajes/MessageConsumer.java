package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.stereotype.Service;

import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.services.EstadisticasService;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@Service
public class MessageConsumer {
    private static final Logger log = LoggerFactory.getLogger(MessageConsumer.class);
    private final EstadisticasService estadisticasService;

    public MessageConsumer(EstadisticasService estadisticasService) {
        this.estadisticasService = estadisticasService;
    }

    @RabbitListener(queues = "notificaciones_audios")
    public void notificacionesCanciones(CancionAlmacenarDTOInput notificacion) {
        try {
            estadisticasService.registrarReproduccion(notificacion.getClienteId(), notificacion.getTitulo());
            log.info("Reproducción registrada para cliente {}: {}", notificacion.getClienteId(), notificacion.getTitulo());
        } catch (IllegalArgumentException | NullPointerException e) {
            log.warn("Se ignoró una notificación de reproducción incompleta: {}", notificacion, e);
        }
    }
}
    
