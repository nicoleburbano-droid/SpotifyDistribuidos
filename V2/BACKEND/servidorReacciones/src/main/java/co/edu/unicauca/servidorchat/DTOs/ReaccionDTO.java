package co.edu.unicauca.servidorchat.DTOs;

/** Reacción que un cliente envía mientras reproduce un audio. */
public class ReaccionDTO {
    private String usuario;
    private String reaccion;     // like, heart, sad, fun
    private String tituloAudio;
    private String hora;         // la pone el servidor al reenviarla

    public ReaccionDTO() {
    }

    public String getUsuario() { return usuario; }
    public void setUsuario(String usuario) { this.usuario = usuario; }

    public String getReaccion() { return reaccion; }
    public void setReaccion(String reaccion) { this.reaccion = reaccion; }

    public String getTituloAudio() { return tituloAudio; }
    public void setTituloAudio(String tituloAudio) { this.tituloAudio = tituloAudio; }

    public String getHora() { return hora; }
    public void setHora(String hora) { this.hora = hora; }
}
