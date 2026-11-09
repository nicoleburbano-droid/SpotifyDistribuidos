// Administrador web: registra un audio con sus metadatos a través del API Gateway.
// POST /api/audios/almacenamiento  (multipart: archivo, tipo, metadatos en JSON)
const GATEWAY = 'http://localhost:8081';
const $ = (id) => document.getElementById(id);

// Campos de cada tipo, iguales a los DTO del servidor de audios.
// El primero es el título del audio (nombre del archivo en el streaming).
const TIPOS = {
  musica: {
    nombre: 'Música', lista: '/musica/all', vector: 'vectorMusica',
    campos: [
      { clave: 'tituloCancion', etiqueta: 'Título de la canción', obligatorio: true },
      { clave: 'artistaPrincipal', etiqueta: 'Artista principal', ayuda: 'El intérprete o banda líder del proyecto', obligatorio: true },
      { clave: 'album', etiqueta: 'Álbum', ayuda: 'El disco, EP o sencillo que agrupa la obra', obligatorio: true },
      { clave: 'genero', etiqueta: 'Género musical', ayuda: 'Pop, Rock, Jazz...', obligatorio: true },
      { clave: 'selloDiscografico', etiqueta: 'Sello discográfico' },
      { clave: 'anioLanzamiento', etiqueta: 'Año de lanzamiento' },
    ],
  },
  audioLibros: {
    nombre: 'Audiolibros', lista: '/audiolibros/all', vector: 'vectorAudioLibros',
    campos: [
      { clave: 'tituloLibro', etiqueta: 'Título del libro', ayuda: 'El nombre de la obra literaria completa', obligatorio: true },
      { clave: 'autor', etiqueta: 'Autor', ayuda: 'El escritor original del texto', obligatorio: true },
      { clave: 'narrador', etiqueta: 'Narrador', ayuda: 'El actor de voz que realiza la lectura', obligatorio: true },
      { clave: 'editorial', etiqueta: 'Editorial' },
      { clave: 'isbn', etiqueta: 'ISBN', numero: true },
      { clave: 'capitulo', etiqueta: 'Capítulo', numero: true },
    ],
  },
  podcast: {
    nombre: 'Podcasts', lista: '/podcasts/all', vector: 'vectorPodcast',
    campos: [
      { clave: 'tituloEpisodio', etiqueta: 'Título del episodio', ayuda: 'El tema específico de ese audio', obligatorio: true },
      { clave: 'nombrePodcast', etiqueta: 'Nombre del podcast', ayuda: 'El título general del programa o serie', obligatorio: true },
      { clave: 'anfitrion', etiqueta: 'Anfitrión (host)', ayuda: 'La persona que dirige o presenta', obligatorio: true },
      { clave: 'numeroTemporada', etiqueta: 'Temporada', numero: true },
      { clave: 'notasShow', etiqueta: 'Notas del episodio' },
      { clave: 'clasificacionContenido', etiqueta: 'Clasificación' },
    ],
  },
  ruidoBlanco: {
    nombre: 'Ruido blanco', lista: '/ruido-blanco/all', vector: 'vectorRuidoBlanco',
    campos: [
      { clave: 'titulo', etiqueta: 'Título', obligatorio: true },
      { clave: 'tipoSonido', etiqueta: 'Tipo de sonido', ayuda: 'La variante de frecuencia', opciones: ['Ruido Blanco', 'Marrón', 'Rosa'], obligatorio: true },
      { clave: 'fuenteAudio', etiqueta: 'Fuente del audio', ayuda: 'Lluvia, Ventilador, Bosque...', obligatorio: true },
      { clave: 'usoSugerido', etiqueta: 'Uso sugerido', opciones: ['Dormir', 'Concentración', 'Meditación'], obligatorio: true },
      { clave: 'proveedorContenido', etiqueta: 'Proveedor del contenido' },
      { clave: 'duracionBucle', etiqueta: 'Duración del bucle (s)', numero: true },
      { clave: 'frecuenciaDominante', etiqueta: 'Frecuencia dominante' },
    ],
  },
};

function mensaje(texto, clase) {
  $('mensaje').textContent = texto;
  $('mensaje').className = clase || '';
}

// El formulario cambia según el tipo elegido
function pintarCampos() {
  const tipo = TIPOS[$('tipo').value];
  $('leyenda').textContent = 'Metadatos de ' + tipo.nombre.toLowerCase();
  const cont = $('campos');
  cont.innerHTML = '';
  tipo.campos.forEach((c) => {
    const label = document.createElement('label');
    label.className = 'campo';
    const texto = document.createElement('span');
    texto.textContent = c.etiqueta + (c.ayuda ? ' (' + c.ayuda + ')' : '');
    if (c.obligatorio) texto.className = 'obligatorio';
    let control;
    if (c.opciones) {
      control = document.createElement('select');
      c.opciones.forEach((o) => control.add(new Option(o, o)));
    } else {
      control = document.createElement('input');
      control.type = c.numero ? 'number' : 'text';
    }
    control.name = c.clave;
    control.required = !!c.obligatorio;
    label.append(texto, control);
    cont.appendChild(label);
  });
  cargarLista();
}

async function cargarLista() {
  const tipo = TIPOS[$('tipo').value];
  const lista = $('lista');
  lista.innerHTML = '';
  try {
    const r = await fetch(GATEWAY + '/api' + tipo.lista);
    const datos = await r.json();
    (datos[tipo.vector] || []).forEach((a) => {
      const li = document.createElement('li');
      const b = document.createElement('button');
      b.type = 'button';
      b.textContent = a[tipo.campos[0].clave];
      li.appendChild(b);
      lista.appendChild(li);
    });
  } catch (e) {
    mensaje('No hay conexión con el API Gateway o con el servidor de audios.', 'error');
  }
}

$('formulario').addEventListener('submit', (e) => {
  e.preventDefault();
  const idTipo = $('tipo').value;
  const metadatos = {};
  TIPOS[idTipo].campos.forEach((c) => {
    const valor = $('formulario').elements[c.clave].value.trim();
    metadatos[c.clave] = c.numero ? (parseInt(valor, 10) || 0) : valor;
  });
  const datos = new FormData();
  datos.append('tipo', idTipo);
  datos.append('metadatos', JSON.stringify(metadatos));
  datos.append('archivo', $('archivo').files[0]);

  const boton = $('btnRegistrar');
  const progreso = $('progreso');
  boton.disabled = true;
  progreso.hidden = false;
  progreso.value = 0;
  mensaje('Subiendo audio...');

  // XMLHttpRequest para mostrar el progreso de subida
  const xhr = new XMLHttpRequest();
  xhr.open('POST', GATEWAY + '/api/audios/almacenamiento');
  xhr.upload.onprogress = (ev) => { if (ev.lengthComputable) progreso.value = (ev.loaded / ev.total) * 100; };
  xhr.onload = () => {
    boton.disabled = false;
    progreso.hidden = true;
    let r = {};
    try { r = JSON.parse(xhr.responseText); } catch (_) { /* sin cuerpo */ }
    if (xhr.status === 201) {
      mensaje(r.mensaje + '. Ya aparece en el cliente web.', 'success');
      $('formulario').reset();
      $('tipo').value = idTipo; // el reset vuelve al primer tipo
      pintarCampos();
    } else {
      mensaje('No se registró: ' + (r.mensaje || 'error ' + xhr.status), 'error');
    }
  };
  xhr.onerror = () => {
    boton.disabled = false;
    progreso.hidden = true;
    mensaje('No hay conexión con el API Gateway.', 'error');
  };
  xhr.send(datos);
});

Object.entries(TIPOS).forEach(([id, t]) => $('tipo').add(new Option(t.nombre, id)));
$('tipo').addEventListener('change', pintarCampos);
pintarCampos();
