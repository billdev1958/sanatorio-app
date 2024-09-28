import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import App from "./App"; // Este es el layout (antes llamado Layout)
import Home from "./pages/Homepage"; // Página de ejemplo
import Citas from "./pages/Citas";

const root = document.getElementById("root");

render(
  () => (
    <Router>
      {/* Rutas */}
      <Route path="/" component={App}>
        <Route path="/" component={Home} /> {/* Ruta para HomePage */}
        <Route path="/hello-world" component={() => <div>Hello World!</div>} /> {/* Ruta adicional */}
        <Route path="/citas" component={Citas} /> {/* Ruta para HomePage */}
      </Route>
    </Router>
  ),
  root!
);
