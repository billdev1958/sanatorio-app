import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import App from "./App"; // Este es el layout (antes llamado Layout)
import Citas from "./pages/Citas";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import HistorialMedicoPage from "./pages/MedicalHistoryP";


const root = document.getElementById("root");

render(
  () => (
    <Router>
      {/* Rutas */}
      <Route path="/" component={App}>
        <Route path="/" component={Dashboard} /> {/* Ruta para HomePage */}
        <Route path="/citas" component={Citas} /> {/* Ruta para HomePage */}
        <Route path="/login" component={Login} /> {/* Ruta para Login */}
        <Route path="/register" component={Register} /> {/* Ruta para Login */}
        <Route path="/medicalhistory" component={HistorialMedicoPage} /> {/* Ruta para Login */}



      </Route>
    </Router>
  ),
  root!
);
