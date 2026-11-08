let clienteChat = null;
let suscripcionAudio = null;
let audioActual = "";

function setConectado(conectado) {
document.getElementById('btnConectar').disabled = conectado;
document.getElementById('btnDesconectar') .disabled = !conectado;
}

function conectar(){
    const socket = new SockJS('http://localhost:8081/api/reacciones/ws');
    clienteChat = Stomp.over(socket);
    clienteChat.connect({}, onConnected);
}

function onConnected(frame) {
    console.log("Conectado: " + frame);
    setConectado(true);
    actualizarSuscripcionAudio();
}

function recibirReaccion(reaccion){
    let datos;
    try {
        datos = JSON.parse(reaccion.body);
    } catch (error) {
        console.error("Formato de reacción inválido:", error);
        return;
    }
    console.log("Reacción reenviada:", datos);
    const contenedor = document.getElementById("mensajes");

    //crear burbuja
    const burbuja = document.createElement("div");
    burbuja.classList.add("reaccion-burbuja");

    let icono = "";
    switch(datos.reaccion){
        case "like":
            icono = '<i class="fa-solid fa-thumbs-up icon-like"></i>';
            break;
        case "heart":
            icono = '<i class="fa-solid fa-heart icon-heart"></i>';
            break;
        case "sad":
            icono = '<i class="fa-solid fa-face-sad-tear icon-sad"></i>';
            break;
        case "fun":
            icono = '<i class="fa-solid fa-face-laugh icon-fun"></i>';
            break;

    }

    burbuja.innerHTML = icono;
    contenedor.appendChild(burbuja);

    //eliminar burbuja despues de la animacion
    setTimeout(
        ()=>{
            burbuja.remove();
        }, 3000);
}

function enviarReaccionServidor(event){
    const icon = event.target;
    handleReactionClick(icon);

    const reaccion = event.target.dataset.reaction;
    if(clienteChat && clienteChat.connected){
        if (!audioActual) {
            alert("Selecciona primero el audio que estás escuchando");
            return;
        }
        clienteChat.send("/app/enviarReaccion", {"content-type": "application/json"}, JSON.stringify({audioId: audioActual, reaccion}));

    }else{
        alert("No estas conectado");
    }
}

function seleccionarAudioActual() {
    const entrada = document.getElementById("audio-actual");
    audioActual = entrada.value.trim();
    if (!audioActual) {
        alert("Ingresa el ID o título del audio");
        return;
    }
    actualizarSuscripcionAudio();
}

function codificarAudioId(audioId) {
    const bytes = new TextEncoder().encode(audioId);
    let binario = "";
    bytes.forEach(byte => binario += String.fromCharCode(byte));
    return btoa(binario).replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/g, "");
}

function actualizarSuscripcionAudio() {
    if (!clienteChat || !clienteChat.connected || !audioActual) return;
    if (suscripcionAudio) suscripcionAudio.unsubscribe();
    const canal = codificarAudioId(audioActual);
    suscripcionAudio = clienteChat.subscribe(`/brokerDeReacciones/audios/${canal}`, recibirReaccion);
}

function desconectar () {
    if (clienteChat !== null) {
        if (suscripcionAudio) suscripcionAudio.unsubscribe();
        suscripcionAudio = null;
        clienteChat.disconnect(() => {
            setConectado(false);
        });
        clienteChat = null;
    }
}

