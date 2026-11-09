(() => {
  const API_GATEWAY = "http://localhost:8081/api";
  const $ = (id) => document.getElementById(id);

  function mostrarMensaje(mensaje, esError = false) {
    const estado = $("estado-estadisticas");
    estado.textContent = mensaje;
    estado.classList.toggle("error-texto", esError);
  }

  function crearCelda(valor) {
    const celda = document.createElement("td");
    celda.textContent = valor == null || valor === "" ? "—" : String(valor);
    return celda;
  }

  function mostrarHistorial(historial) {
    const cuerpo = $("tabla-historial");
    cuerpo.replaceChildren();

    if (!Array.isArray(historial) || historial.length === 0) {
      mostrarMensaje("No hay reproducciones registradas para este usuario.");
      return;
    }

    historial.forEach((audio) => {
      const fila = document.createElement("tr");
      fila.append(
        crearCelda(audio.titulo),
        crearCelda(audio.tipo),
        crearCelda(audio.reproducciones),
        crearCelda(audio.ultimaReproduccion)
      );
      cuerpo.append(fila);
    });

    mostrarMensaje(`${historial.length} audio(s) en el historial.`);
  }

  async function consultarEstadisticas(evento) {
    evento?.preventDefault();
    const idUsuario = $("id-usuario").value.trim();
    if (!idUsuario) {
      mostrarMensaje("No hay un ID de usuario. Reproduce un audio primero o ingresa el ID usado por el cliente.", true);
      return;
    }

    mostrarMensaje("Consultando historial…");
    $("btn-consultar-estadisticas").disabled = true;
    try {
      const respuesta = await fetch(`${API_GATEWAY}/estadisticas/historial/${encodeURIComponent(idUsuario)}`);
      if (!respuesta.ok) {
        throw new Error(`El API Gateway respondió HTTP ${respuesta.status}.`);
      }
      const historial = await respuesta.json();
      mostrarHistorial(historial);
    } catch (error) {
      $("tabla-historial").replaceChildren();
      mostrarMensaje(`No fue posible consultar el historial: ${error.message}`, true);
    } finally {
      $("btn-consultar-estadisticas").disabled = false;
    }
  }

  document.addEventListener("DOMContentLoaded", () => {
    try {
      $("id-usuario").value = sessionStorage.getItem("spotify-client-id") || "";
    } catch (_) {
      // El ID puede ingresarse manualmente si el navegador bloquea localStorage.
    }
    $("form-estadistica").addEventListener("submit", consultarEstadisticas);
  });
})();
