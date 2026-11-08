package co.edu.unicauca.infoii.correo.DTOs;

/** Un audio del historial de un usuario y cuántas veces lo reprodujo. */
public class EntradaHistorialDTO {
	private String titulo;
	private String tipo;
	private int reproducciones;
	private String ultimaReproduccion;

	public EntradaHistorialDTO() {
	}

	public EntradaHistorialDTO(String titulo, String tipo, int reproducciones, String ultimaReproduccion) {
		this.titulo = titulo;
		this.tipo = tipo;
		this.reproducciones = reproducciones;
		this.ultimaReproduccion = ultimaReproduccion;
	}

	public String getTitulo() { return titulo; }
	public void setTitulo(String titulo) { this.titulo = titulo; }

	public String getTipo() { return tipo; }
	public void setTipo(String tipo) { this.tipo = tipo; }

	public int getReproducciones() { return reproducciones; }
	public void setReproducciones(int reproducciones) { this.reproducciones = reproducciones; }

	public String getUltimaReproduccion() { return ultimaReproduccion; }
	public void setUltimaReproduccion(String ultimaReproduccion) { this.ultimaReproduccion = ultimaReproduccion; }
}
