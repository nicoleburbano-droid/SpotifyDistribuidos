package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.amqp.rabbit.annotation.Queue;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Service;

import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.capaServicios.ServicioEstadisticas;

/** Comunicación asíncrona: consume las reproducciones que publica el servidor de streaming. */
@Service
public class MessageConsumer {

    private final ServicioEstadisticas servicio;

    public MessageConsumer(ServicioEstadisticas servicio) {
        this.servicio = servicio;
    }

    @RabbitListener(queuesToDeclare = @Queue(value = "notificaciones_audios", durable = "true"))
    public void notificacionesCanciones(CancionAlmacenarDTOInput notificacion) {
        int veces = servicio.registrar(notificacion.getIdUsuario(), notificacion.getTitulo(),
                notificacion.getTipo(), notificacion.getFecha());
        System.out.println("[ECO] Reproducción recibida: usuario=" + notificacion.getIdUsuario()
                + ", audio=" + notificacion.getTitulo() + " (" + notificacion.getTipo() + ")"
                + ", veces=" + veces);
        System.out.println("Mensaje: " + notificacion.getMensaje());
    }
}
