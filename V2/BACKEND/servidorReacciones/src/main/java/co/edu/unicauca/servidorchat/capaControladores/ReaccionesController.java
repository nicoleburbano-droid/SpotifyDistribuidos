package co.edu.unicauca.servidorchat.capaControladores;

import java.time.LocalTime;
import java.time.format.DateTimeFormatter;

import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.messaging.handler.annotation.DestinationVariable;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.stereotype.Controller;

import co.edu.unicauca.servidorchat.DTOs.ReaccionDTO;

/**
 * Cada audio tiene su propio canal: /brokerDeReacciones/audios/{canal}.
 * El cliente publica y se suscribe usando el mismo canal codificado.
 * La reacción se reenvía (callback) solo a los clientes suscritos a ese canal,
 * usando la conexión WebSocket que ya tienen abierta.
 */
@Controller
public class ReaccionesController {

  private static final DateTimeFormatter FORMATO = DateTimeFormatter.ofPattern("HH:mm:ss");

  @MessageMapping("/enviarReaccion/{canal}")
  @SendTo("/brokerDeReacciones/audios/{canal}")
  public ReaccionDTO processMessage(@DestinationVariable String canal, ReaccionDTO reaccion) {
      reaccion.setHora(LocalTime.now().format(FORMATO));
      System.out.println("[ECO] Reacción '" + reaccion.getReaccion() + "' de " + reaccion.getUsuario()
              + " reenviada al canal " + canal);
      return reaccion;
  }
}
