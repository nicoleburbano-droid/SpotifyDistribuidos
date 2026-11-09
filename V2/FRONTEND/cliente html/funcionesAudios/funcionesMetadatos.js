(() => {
  const API_GATEWAY = "http://localhost:8081/api";
  const RUTAS = {
    musica: { lista: "/musica/all", detalle: "/musica/titulo/", campoLista: "vectorMusica", campoTitulo: "tituloCancion", formato: "mp3" },
    audioLibros: { lista: "/audiolibros/all", detalle: "/audiolibros/titulo/", campoLista: "vectorAudioLibros", campoTitulo: "tituloLibro", formato: "mp3" },
    podcast: { lista: "/podcasts/all", detalle: "/podcasts/titulo/", campoLista: "vectorPodcast", campoTitulo: "tituloEpisodio", formato: "mp3" },
    ruidoBlanco: { lista: "/ruido-blanco/all", detalle: "/ruido-blanco/titulo/", campoLista: "vectorRuidoBlanco", campoTitulo: "titulo", formato: "mp3" },
  };
  const $ = (id) => document.getElementById(id);
  const estado = { tipo: "", audio: null, detalle: null, urlAnterior: "" };

  async function consultar(ruta) {
    const respuesta = await fetch(`${API_GATEWAY}${ruta}`);
    if (!respuesta.ok) throw new Error(`El gateway respondió HTTP ${respuesta.status}.`);
    return respuesta.json();
  }

  function tituloDe(audio, tipo) {
    const ruta = RUTAS[tipo];
    return audio.titulo || audio[ruta?.campoTitulo] || "Audio sin título";
  }

  function mostrarEstado(mensaje, error = false) {
    const nodo = $("estado-audios");
    nodo.textContent = mensaje;
    nodo.classList.toggle("error-texto", error);
  }

  function limpiarSeleccion() {
    estado.audio = null;
    estado.detalle = null;
    $("detalle-audio").replaceChildren(document.createTextNode("Selecciona un audio para consultar sus detalles."));
    $("btn-reproducir").disabled = true;
    $("estado-streaming").textContent = "";
  }

  async function cargarTipos() {
    const selector = $("tipo-audio");
    try {
      const tipos = await consultar("/tipos-audio");
      selector.replaceChildren(new Option("Selecciona un tipo", ""));
      tipos.forEach(({ tipo }) => selector.add(new Option(tipo, tipo)));
      if (tipos.length) {
        selector.value = tipos[0].tipo;
        await cargarAudios(tipos[0].tipo);
      } else mostrarEstado("No hay tipos de audio disponibles.");
    } catch (error) {
      selector.replaceChildren(new Option("No disponible", ""));
      mostrarEstado(`No fue posible consultar los tipos: ${error.message}`, true);
    }
  }

  async function cargarAudios(tipo) {
    estado.tipo = tipo;
    limpiarSeleccion();
    const lista = $("lista-audios");
    lista.replaceChildren(Object.assign(document.createElement("p"), { className: "texto-tenue", textContent: "Cargando catálogo…" }));
    const ruta = RUTAS[tipo];
    if (!ruta) {
      lista.replaceChildren(Object.assign(document.createElement("p"), { className: "texto-tenue", textContent: "Este tipo no tiene un endpoint de catálogo disponible." }));
      mostrarEstado("Tipo de audio no soportado por el catálogo.", true);
      return;
    }
    try {
      const datos = await consultar(ruta.lista);
      const audios = datos[ruta.campoLista] || [];
      if (!audios.length) {
        lista.replaceChildren(Object.assign(document.createElement("p"), { className: "texto-tenue", textContent: datos.Mensaje || "No hay audios en este tipo." }));
        mostrarEstado("Catálogo consultado.");
        return;
      }
      lista.replaceChildren(...audios.map((audio) => {
        const titulo = tituloDe(audio, tipo);
        const boton = document.createElement("button");
        boton.type = "button";
        boton.className = "tarjeta audio-opcion";
        boton.textContent = titulo;
        boton.addEventListener("click", () => seleccionarAudio(audio, titulo));
        return boton;
      }));
      mostrarEstado(`${audios.length} audio(s) disponibles.`);
    } catch (error) {
      lista.replaceChildren(Object.assign(document.createElement("p"), { className: "texto-tenue", textContent: `No fue posible cargar el catálogo: ${error.message}` }));
      mostrarEstado("Error consultando el catálogo.", true);
    }
  }

  function renderizarDetalle(base, especifico) {
    const contenedor = $("detalle-audio");
    const datos = { ...base, ...(Array.isArray(especifico) ? especifico[0] : especifico || {}) };
    estado.detalle = datos;
    const filas = Object.entries(datos).filter(([, valor]) => valor !== null && valor !== "" && typeof valor !== "object");
    const lista = document.createElement("dl");
    lista.className = "detalle-lista";
    filas.forEach(([clave, valor]) => {
      const termino = document.createElement("dt");
      termino.textContent = clave.replace(/([A-Z])/g, " $1").replace(/^./, (letra) => letra.toUpperCase());
      const definicion = document.createElement("dd");
      definicion.textContent = String(valor);
      lista.append(termino, definicion);
    });
    contenedor.replaceChildren(lista);
  }

  async function seleccionarAudio(audio, titulo) {
    estado.audio = { ...audio, titulo };
    try { localStorage.setItem("spotify-audio-actual", titulo); } catch (_) { /* almacenamiento opcional */ }
    $("btn-reproducir").disabled = true;
    $("detalle-audio").textContent = "Consultando detalles…";
    document.querySelectorAll(".audio-opcion").forEach((boton) => boton.classList.toggle("seleccionada", boton.textContent === titulo));
    const ruta = RUTAS[estado.tipo];
    try {
      const respuesta = await consultar(`${ruta.detalle}${encodeURIComponent(titulo)}`);
      renderizarDetalle(audio, respuesta[ruta.campoLista] || respuesta.objAudio || {});
      $("btn-reproducir").disabled = audio.disponible === false;
      mostrarEstado(audio.disponible === false ? "El audio está marcado como no disponible." : "Detalles del audio cargados.", audio.disponible === false);
    } catch (error) {
      $("detalle-audio").textContent = `No fue posible consultar los detalles: ${error.message}`;
      mostrarEstado("Error consultando el detalle del audio.", true);
    }
  }

  function reproducir() {
    const audio = estado.audio;
    if (!audio) return;
    const player = $("audio-player");
    if (estado.urlAnterior) URL.revokeObjectURL(estado.urlAnterior);
    player.hidden = false;
    $("estado-streaming").textContent = "Solicitando el audio al servicio de streaming…";
    if (typeof window.iniciar_streaming_cancion !== "function") {
      $("estado-streaming").textContent = "No se cargó el cliente gRPC-Web.";
      return;
    }
    window.iniciar_streaming_cancion(audio.titulo, RUTAS[estado.tipo]?.formato || "mp3", estado.tipo);
  }

  document.addEventListener("DOMContentLoaded", () => {
    $("tipo-audio").addEventListener("change", (evento) => cargarAudios(evento.target.value));
    $("btn-reproducir").addEventListener("click", reproducir);
    cargarTipos();
  });
})();
