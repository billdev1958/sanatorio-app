import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import { BrowserRouter, Route, Routes, Navigate, Outlet } from 'react-router-dom'
import Homepage from './pages/Homepage.tsx'
import InvestigacionPage from './pages/InvestigacionPage.tsx'
import BlogPage from './pages/UsersMainPage.tsx'
import BlogArticle from './pages/BlogArticle.tsx'
import Login from './pages/Login.tsx'
import RegisterDoctor from './pages/RegisterDoctors.tsx'
import UsersMainPage from './pages/UsersMainPage.tsx'
import RegisterUser from './pages/RegisterUsers.tsx'
import ControlMenu from './pages/ControlMenu.tsx'
import UpdateUser from './pages/UpdateUser.tsx'


// Layout protegido
const ProtectedRoute = () => {
  const isAuthenticated = !!localStorage.getItem('token'); // Verifica si el usuario está autenticado

  return isAuthenticated ? <Outlet /> : <Navigate to="/login" />;
};

export const AppRoutes = () => {
  return (
        <Routes>
          {/* Ruta de Login fuera del layout protegido */}
          <Route path="/login" element={<Login />} />

          {/* Rutas protegidas */}
          <Route path="/" element={<ProtectedRoute />}>
            <Route element={<App />}>
              <Route index element={<Homepage />} /> {/* Esta es la homepage */}
              <Route path="control" element={<ControlMenu />} />
              <Route path="user" element={<UsersMainPage />} />
              <Route path="register/user" element={<RegisterUser />} />
              <Route path="register/doctor" element={<RegisterDoctor />} />
              <Route path="/user/update/:userId" element={<UpdateUser />} />
              <Route path="investigacion" element={<InvestigacionPage />} />
              <Route path="blog" element={<BlogPage />} />
              <Route path="blog/:id" element={<BlogArticle />} />
            </Route>
          </Route>
        </Routes>
  );
};

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <AppRoutes />
    </BrowserRouter>
  </React.StrictMode>
)
