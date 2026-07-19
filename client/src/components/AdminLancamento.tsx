import React, { useState } from "react";
import '../App.css'


export default function AdminLancamento(){
    const [nomeMusica, setNomeMusica] = useState('');
    const[link, setLink] = useState('');
    const [senhaAdmin, setSenhaAdmin] = useState('');
    const [status, setStatus] = useState('');
    const [loading, setLoading] = useState(false);

    const dispararEmails = async (e: React.FormEvent)=>{
        e.preventDefault();
        setLoading(true);
        setStatus("Processando...");

        const API_URL = import.meta.env.VITE_API_URL

        try{
            const response = await fetch(`${API_URL}/api/notify-release`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    nome_musica: nomeMusica,
                    link: link,
                    senha_admin: senhaAdmin,
                }),
            });

            const data = await response.json()

            if(response.ok){
                setStatus(`Sucesso ${data.mensagem}`)
                setNomeMusica('');
                setLink('');
            }else{
                setStatus(`Erro: ${data.erro || data.error}`);
            }
        }catch(error){
            console.error(error);
            setStatus("Erro de conexão com a API")
        }finally{
            setLoading(false)
        }
    };

    return(
        <div className="admin-container">
        <h2>Painel de Lançamento</h2>
        <p>Dispare a notificação para todos os inscritos.</p>

        <form className="admin-form" onSubmit={dispararEmails}>
            
            <div className="input-group">
            <label>Nome da Música:</label>
            <input 
                type="text" 
                value={nomeMusica}
                onChange={(e) => setNomeMusica(e.target.value)}
                required 
            />
            <hr />
            </div>

            <div className="input-group">
            <label>Link para Ouvir:</label>
            <input 
                type="url" 
                value={link}
                onChange={(e) => setLink(e.target.value)}
                required 
            />
            <hr />
            </div>

            <div className="input-group">
            <label>Senha de Administrador:</label>
            <input 
                type="password" 
                value={senhaAdmin}
                onChange={(e) => setSenhaAdmin(e.target.value)}
                required 
            />
            <hr />
            </div>

            <button 
            type="submit" 
            className="btn-submit"
            disabled={loading}
            >
            {loading ? 'Disparando...' : 'Notificar Inscritos'}
            </button>

        </form>

        {status && (
            <div className={`message ${status.includes('❌') ? 'error' : 'success'}`}>
            {status}
            </div>
        )}
        </div>
    );
}