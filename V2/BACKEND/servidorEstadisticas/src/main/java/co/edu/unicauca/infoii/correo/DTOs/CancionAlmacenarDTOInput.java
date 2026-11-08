package co.edu.unicauca.infoii.correo.DTOs;

/** Notificación que publica el servidor de streaming cuando termina de enviar un audio. */
public class CancionAlmacenarDTOInput {
	private String titulo;
	private String mensaje;
	private String idUsuario;
	private String tipo;
	private String fecha;

	public CancionAlmacenarDTOInput() {
	}

	public String getTitulo() { return titulo; }
	public void setTitulo(String titulo) { this.titulo = titulo; }

	public String getMensaje() { return mensaje; }
	public void setMensaje(String mensaje) { this.mensaje = mensaje; }

	public String getIdUsuario() { return idUsuario; }
	public void setIdUsuario(String idUsuario) { this.idUsuario = idUsuario; }

	public String getTipo() { return tipo; }
	public void setTipo(String tipo) { this.tipo = tipo; }

	public String getFecha() { return fecha; }
	public void setFecha(String fecha) { this.fecha = fecha; }
}