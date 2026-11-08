package co.edu.unicauca.servidorchat.capaEventos;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import org.springframework.context.event.EventListener;
import org.springframework.messaging.simp.stomp.StompHeaderAccessor;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.messaging.SessionConnectedEvent;
import org.springframework.web.socket.messaging.SessionDisconnectEvent;
import org.springframework.web.socket.messaging.SessionSubscribeEvent;
import org.springframework.web.socket.messaging.SessionUnsubscribeEvent;

/**
 * Ecos del servidor: muestra cuando un cliente se registra (suscribe) o se
 * desregistra (desuscribe) del canal de un audio.
 */
@Component
public class RegistroCanalesListener {

    /** clave sesion:suscripcion -> canal */
    private final Map<String, String> suscripciones = new ConcurrentHashMap<>();

    @EventListener
    public void alConectar(SessionConnectedEvent evento) {
        StompHeaderAccessor h = StompHeaderAccessor.wrap(evento.getMessage());
        System.out.println("[ECO] Cliente conectado, sesión " + h.getSessionId());
    }

    @EventListener
    public void alSuscribir(SessionSubscribeEvent evento) {
        StompHeaderAccessor h = StompHeaderAccessor.wrap(evento.getMessage());
        String canal = h.getDestination();
        suscripciones.put(h.getSessionId() + ":" + h.getSubscriptionId(), canal);
        System.out.println("[ECO] Sesión " + h.getSessionId() + " registrada en " + canal
                + " (oyentes: " + oyentes(canal) + ")");
    }

    @EventListener
    public void alDesuscribir(SessionUnsubscribeEvent evento) {
        StompHeaderAccessor h = StompHeaderAccessor.wrap(evento.getMessage());
        String canal = suscripciones.remove(h.getSessionId() + ":" + h.getSubscriptionId());
        if (canal != null) {
            System.out.println("[ECO] Sesión " + h.getSessionId() + " desregistrada de " + canal
                    + " (oyentes: " + oyentes(canal) + ")");
        }
    }

    @EventListener
    public void alDesconectar(SessionDisconnectEvent evento) {
        String sesion = evento.getSessionId();
        suscripciones.keySet().removeIf(clave -> clave.startsWith(sesion + ":"));
        System.out.println("[ECO] Cliente desconectado, sesión " + sesion);
    }

    private long oyentes(String canal) {
        return suscripciones.values().stream().filter(canal::equals).count();
    }
}
