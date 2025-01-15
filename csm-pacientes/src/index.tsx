import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import App from "./App"; // Layout principal
import Citas from "./Citas/Pages/Citas";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import HistorialMedicoPage from "./pages/MedicalHistoryP";
import { AuthProvider } from "./services/AuthContext";
import ProtectedRoute from "./services/ProtectedRoute"; // Importamos el componente ProtectedRoute
import RegisterBeneficiary from "./pages/RegisterBeneficiary";
import DashboardAdmin from "./Admin/Pages/Menu";
import RegisterOfficeForm from "./Admin/Pages/RegisterOffice";
import RegisterOfficeScheduleForm from "./Admin/Pages/RegistrarHorarios";
import RegisterDoctorForm from "./Admin/Pages/RegistrarDoctores";
import ConsultasHome from "./Citas/Pages/ConsultasHome";

const root = document.getElementById("root");

render(
  () => (
    <AuthProvider>
      <Router>
        {/* Rutas públicas */}
        <Route path="/login" component={Login} />
        <Route path="/register" component={Register} />

        {/* Rutas protegidas */}
        <Route path="/" component={() => (
          <ProtectedRoute>
            <App>
              <Dashboard />
            </App>
          </ProtectedRoute>
        )} />
        
          <Route path="/admin" component={() => (
          <ProtectedRoute>
            <App>
              <DashboardAdmin />
            </App>
          </ProtectedRoute>
        )} />
        
        <Route path="/admin/office" component={() => (
          <ProtectedRoute>
            <App>
              <RegisterOfficeForm />
            </App>
          </ProtectedRoute>
        )} />

        <Route path="/admin/doctor" component={() => (
          <ProtectedRoute>
            <App>
              <RegisterDoctorForm />
            </App>
          </ProtectedRoute>
        )} />

        <Route path="/admin/schedule" component={() => (
          <ProtectedRoute>
            <App>
              <RegisterOfficeScheduleForm />
            </App>
          </ProtectedRoute>
        )} />

        <Route path="/citas" component={() => (
          <ProtectedRoute>
            <App>
              <Citas />
            </App>
          </ProtectedRoute>
        )} />
          <Route path="/consultas" component={() => (
          <ProtectedRoute>
            <App>
              <ConsultasHome />
            </App>
          </ProtectedRoute>
        )} />
        <Route path="/medicalhistory" component={() => (
          <ProtectedRoute>
            <App>
              <HistorialMedicoPage />
            </App>
          </ProtectedRoute>
        )} />
        <Route path="/beneficiary" component={() => (
          <ProtectedRoute>
            <App>
              <RegisterBeneficiary />
            </App>
          </ProtectedRoute>
        )} />
      </Router>
    </AuthProvider>
  ),
  root!
);
