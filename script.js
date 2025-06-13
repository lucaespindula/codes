let seuVotoPara = document.querySelector('.d-1-1 span'); 
let cargo = document.querySelector('.d-1-2 span');
let descricao = document.querySelector('.d-1-4');
let aviso = document.querySelector('.d-2');
let lateral = document.querySelector('.d-1-right');
let numeros = document.querySelector('.d-1-3');

let etapaAtual = 0;
let numero = '';
let votoBranco = false;
let votos = [];

const contagemVotos = {
    '38111': { nome: 'Fulano de Tal 38111', votos: 0 },
    '77222': { nome: 'Beltrano da Silva 77222', votos: 0 },
    '99': { nome: 'Ciclano 99', votos: 0 },
    '84': { nome: 'Zulano 84', votos: 0 },
    'branco': { nome: 'Voto em Branco', votos: 0 },
    'nulo': { nome: 'Voto Nulo', votos: 0 }
};

function comecarEtapa() {
    let etapa = etapas[etapaAtual];
    votoBranco = false;
    numero = '';
    let numeroHtml = '';

    for (let i = 0; i < etapa.numeros; i++) {
        if (i === 0) {
            numeroHtml += '<div class="numero pisca"></div>';
        } else {
            numeroHtml += '<div class="numero"></div>';
        }
    }

    seuVotoPara.style.display = 'none';
    cargo.innerHTML = etapa.titulo;
    descricao.innerHTML = '';
    aviso.style.display = 'none';
    lateral.innerHTML = '';
    numeros.innerHTML = numeroHtml;
}

function atualizaInterface() {
    let etapa = etapas[etapaAtual];
    let candidato = etapa.candidatos.filter((item) => item.numero === numero);

    if (candidato.length > 0) {
        candidato = candidato[0];
        seuVotoPara.style.display = 'block';
        aviso.style.display = 'block';
        descricao.innerHTML = `Nome: ${candidato.nome}<br/>Partido: ${candidato.partido}`;
        let fotosHtml = '';
        for (let i in candidato.fotos) {
            if (candidato.fotos[i].small) {
                fotosHtml += `<div class="d-1-image small"><img src="${candidato.fotos[i].url}" alt="" />${candidato.fotos[i].legenda}</div>`;
            } else {
                fotosHtml += `<div class="d-1-image"><img src="${candidato.fotos[i].url}" alt="" />${candidato.fotos[i].legenda}</div>`;
            }
        }
        lateral.innerHTML = fotosHtml;
    } else {
        seuVotoPara.style.display = 'block';
        aviso.style.display = 'block';
        descricao.innerHTML = '<div class="aviso--grande pisca">VOTO NULO</div>';
    }
}

function clicou(n) {
    let elNumero = document.querySelector('.numero.pisca');
    if (elNumero !== null) {
        elNumero.innerHTML = n; 
        numero += n;

        elNumero.classList.remove('pisca');
        if (elNumero.nextElementSibling !== null) {
            elNumero.nextElementSibling.classList.add('pisca');
        } else {
            atualizaInterface();
        }
    }
}

function branco() {
    numero = '';
    votoBranco = true;

    seuVotoPara.style.display = 'block';
    aviso.style.display = 'block';
    numeros.innerHTML = '';
    descricao.innerHTML = '<div class="aviso--grande pisca">VOTO EM BRANCO</div>';
    lateral.innerHTML = '';
}

function corrige() {
    comecarEtapa();
}

function exibirBoletim() {
    let boletimHtml = `
    <div class="boletim-container">
        <h2>BOLETIM DE URNA</h2>
        <ul class="boletim-lista">
`;

    for (let candidato in contagemVotos) {
        let votosCandidato = contagemVotos[candidato].votos;
        let votoTermo = votosCandidato === 1 ? 'voto' : 'votos';
        boletimHtml += `
        <li class="boletim-item">
            <span class="boletim-nome">${contagemVotos[candidato].nome}</span>
            <span class="boletim-votos">${votosCandidato} ${votoTermo}</span>
        </li>
        `;
    }

    boletimHtml += `
        </ul>
    </div>
    `;

    document.querySelector('.tela').innerHTML = boletimHtml;
    document.querySelector('.teclado').style.display = 'none';
}

function confirma() {
    let etapa = etapas[etapaAtual];
    let votoConfirmado = false;

    if (votoBranco === true) {
        votoConfirmado = true;
        contagemVotos['branco'].votos++;
        votos.push({
            etapa: etapas[etapaAtual].titulo,
            voto: 'branco'
        });
    } else if (numero.length === etapa.numeros && contagemVotos[numero]) {
        votoConfirmado = true;
        contagemVotos[numero].votos++;
        votos.push({
            etapa: etapas[etapaAtual].titulo,
            voto: numero
        });
    } else {
        contagemVotos['nulo'].votos++;
        votos.push({
            etapa: etapas[etapaAtual].titulo,
            voto: 'nulo'
        });
    }

    etapaAtual++;
    if (etapas[etapaAtual] !== undefined) {
        comecarEtapa();
    } else {
        alert('Você votou para todos os cargos. Selecione "Nova Votação" para outra pessoa votar, ou "Exibir Boletim", para encerrar a votação!');
    }
}

document.querySelector('#reiniciar').addEventListener('click', () => {
    etapaAtual = 0;
    votos = [];
    comecarEtapa();
});

document.querySelector('#nova-votacao').addEventListener('click', () => {
    etapaAtual = 0;
    comecarEtapa();
});

document.querySelector('#boletim').addEventListener('click', () => {
    exibirBoletim();
});

comecarEtapa();
