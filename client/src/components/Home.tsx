import React, { useState } from 'react';
import '../App.css';
const LOGO = new URL('../assets/LOGO.png', import.meta.url).href;

function Home() {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState({ text: '', type: ''});

  const API_URL = import.meta.env.VITE_API_URL
  
  const handleSubmit = async (e: React.FormEvent)=>{
    e.preventDefault()
    const response = await fetch(`${API_URL}/api/subscribers`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json'},
      body: JSON.stringify({email: email})
    });
    console.log("URL da API carregada:", API_URL);
    const data = await response.json();

    if (response.ok){
      setMessage({text: 'Inscrição realizada com sucesso!', type: 'success'})
    }else if(response.status===409){
      setMessage({text: data.error, type: 'error'})
    }else{
      setMessage({text: 'Erro de conexão com o servidor.', type: 'error'})
    }
  }

  return (
    <div className='container'>
      <img src={LOGO} alt='Logo'/>
      <form onSubmit={handleSubmit}>
        <div className='groupemail'>
          <input type='email' placeholder='YOUR EMAIL' value={email} onChange={(e) => setEmail(e.target.value)} />
          <hr></hr>
        </div>
        <button type='submit' className='btn-submit'>SEND</button>
      </form>
      {message.text &&(
        <p className={`message ${message.type}`}>
          {message.text}
        </p>
      )}
    </div>
  );
}

export default Home;