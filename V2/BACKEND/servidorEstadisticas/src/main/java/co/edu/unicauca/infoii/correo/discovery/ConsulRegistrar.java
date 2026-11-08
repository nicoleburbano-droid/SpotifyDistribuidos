package co.edu.unicauca.infoii.correo.discovery;

import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Component
@EnableScheduling
public class ConsulRegistrar {
    private static final Logger log = LoggerFactory.getLogger(ConsulRegistrar.class);
    private final HttpClient http = HttpClient.newBuilder().connectTimeout(Duration.ofSeconds(3)).build();
    private final ObjectMapper json = new ObjectMapper();

    @Value("${consul.address:http://127.0.0.1:8500}") private String consulAddress;
    @Value("${spring.application.name:servidorEstadisticas}") private String serviceName;
    @Value("${server.port:6000}") private int servicePort;
    @Value("${service.host:127.0.0.1}") private String serviceHost;
    private String serviceId;
    private String checkId;

    @PostConstruct
    void registrar() {
        serviceId = serviceName + "-" + servicePort;
        checkId = "check-" + serviceId;
        Map<String, Object> check = Map.of("CheckID", checkId, "TTL", "15s", "DeregisterCriticalServiceAfter", "1m");
        Map<String, Object> registration = Map.of("ID", serviceId, "Name", serviceName, "Address", serviceHost,
                "Port", servicePort, "Tags", List.of("spring", "rest", "rabbitmq-consumer"), "Check", check);
        request("PUT", "/v1/agent/service/register", encode(registration));
        heartbeat();
    }

    @Scheduled(fixedDelay = 5000)
    void heartbeat() {
        if (checkId != null) {
            try {
                request("PUT", "/v1/agent/check/pass/" + checkId, "");
            } catch (IllegalStateException e) {
                log.warn("No se pudo renovar el health check de {} en Consul", serviceName, e);
            }
        }
    }

    @PreDestroy
    void desregistrar() {
        if (serviceId != null) request("PUT", "/v1/agent/service/deregister/" + serviceId, "");
    }

    private String encode(Object value) {
        try { return json.writeValueAsString(value); }
        catch (Exception e) { throw new IllegalStateException("No se pudo serializar el registro de Consul", e); }
    }

    private void request(String method, String path, String body) {
        try {
            HttpRequest.Builder builder = HttpRequest.newBuilder(URI.create(consulAddress.replaceAll("/$", "") + path))
                    .timeout(Duration.ofSeconds(3));
            if (body.isEmpty()) builder.method(method, HttpRequest.BodyPublishers.noBody());
            else builder.header("Content-Type", "application/json").method(method, HttpRequest.BodyPublishers.ofString(body));
            HttpResponse<Void> response = http.send(builder.build(), HttpResponse.BodyHandlers.discarding());
            if (response.statusCode() < 200 || response.statusCode() >= 300)
                throw new IllegalStateException("Consul respondió HTTP " + response.statusCode());
        } catch (Exception e) {
            throw new IllegalStateException("No fue posible comunicarse con Consul en " + consulAddress, e);
        }
    }
}
