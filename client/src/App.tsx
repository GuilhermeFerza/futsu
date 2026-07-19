import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './components/Home';
import AdminLancamento from './components/AdminLancamento'; 

function App() {
  const ADMIN_PATH = import.meta.env.VITE_PATH_ADM
  return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="*" element={<Home />} />
          <Route path={`${ADMIN_PATH}`} element={<AdminLancamento />} />
        </Routes>
      </BrowserRouter>
  );
}

export default App;