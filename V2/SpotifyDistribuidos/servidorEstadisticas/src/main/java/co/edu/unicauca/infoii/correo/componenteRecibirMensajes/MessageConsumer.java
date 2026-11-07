package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.stereotype.Service;

import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.commons.Simulacion;

import org.springframework.amqp.rabbit.annotation.RabbitListener;

@Service
public class MessageConsumer {
    @RabbitListener(queues = "notificaciones_audios")
    public void notificacionesCanciones(CancionAlmacenarDTOInput objClienteCreado){
        Simulacion.simular(10000);
        System.out.println("Titulo: "+objClienteCreado.getTitulo());
        System.out.println("Mensaje: "+objClienteCreado.getMensaje());
    }
    
}
    
