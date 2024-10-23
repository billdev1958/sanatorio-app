import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import App from "./App"; // Este es el layout (antes llamado Layout)
import Home from "./pages/Homepage"; // Página de ejemplo
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
        <Route path="/" component={Home} /> {/* Ruta para HomePage */}
        <Route path="/hello-world" component={() => <div>Hello World!</div>} /> {/* Ruta adicional */}
        <Route path="/citas" component={Citas} /> {/* Ruta para HomePage */}
        <Route path="/login" component={Login} /> {/* Ruta para Login */}
        <Route path="/register" component={Register} /> {/* Ruta para Login */}
        <Route path="/dashboard" component={Dashboard} /> {/* Ruta para Login */}
        <Route path="/medicalhistory" component={HistorialMedicoPage} /> {/* Ruta para Login */}



      </Route>
    </Router>
  ),
  root!
);
