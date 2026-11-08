package co.edu.unicauca.servidorchat.capaControladores;

import co.edu.unicauca.servidorchat.DTOs.ReaccionDTO;
import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.Set;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.stereotype.Controller;
import org.springframework.util.StringUtils;

@Controller
public class ReaccionesController {
  private static final String DESTINO_REACCIONES = "/brokerDeReacciones/audios/";
  private static final Set<String> REACCIONES_VALIDAS = Set.of("like", "heart", "sad", "fun");
  private final SimpMessagingTemplate messagingTemplate;

  public ReaccionesController(SimpMessagingTemplate messagingTemplate) {
    this.messagingTemplate = messagingTemplate;
  }

  @MessageMapping("/enviarReaccion")
  public void processMessage(@Payload ReaccionDTO reaccion) {
    if (reaccion == null) {
      return;
    }

    String audioId = reaccion.audioId() == null ? "" : reaccion.audioId().trim();
    String tipoReaccion = reaccion.reaccion() == null ? "" : reaccion.reaccion().trim();
    if (!StringUtils.hasText(audioId) || audioId.length() > 256 || !REACCIONES_VALIDAS.contains(tipoReaccion)) {
      return;
    }
    String canalAudio = Base64.getUrlEncoder().withoutPadding()
        .encodeToString(audioId.getBytes(StandardCharsets.UTF_8));
    messagingTemplate.convertAndSend(DESTINO_REACCIONES + canalAudio,
        new ReaccionDTO(audioId, tipoReaccion));
  }
}

