import { createSignal } from 'solid-js';
import NavBar from '../components/NavBar'; // Importa el NavBar
import FormInput from '../components/FormInput'; // Importa el componente reutilizable

const Register = () => {
  // Creamos señales para manejar el estado de cada campo de formulario
  const [name, setName] = createSignal<string>("");
  const [lastname1, setLastname1] = createSignal<string>("");
  const [lastname2, setLastname2] = createSignal<string>("");
  const [curp, setCurp] = createSignal<string>("");
  const [phone, setPhone] = createSignal<string>("");
  const [email, setEmail] = createSignal<string>("");
  const [password, setPassword] = createSignal<string>("");

  return (
    <>
      {/* Incluimos el NavBar siempre */}
      <NavBar toggleMenu={() => {}} /> {/* No necesitamos el menú hamburguesa en el registro, por eso la función está vacía */}

      <div class="form-container">
        <div class="form-card">
          <h2>Registrar Usuario</h2>
          <form class="form">
            {/* Afiliación select */}
            <div class="input-group select-wrapper">
              <label for="afiliation">Afiliación</label>
              <select name="afiliation" id="afiliation" required>
                <option value="">Selecciona una opción</option>
                <option value="Administrativo">Administrativo</option>
                <option value="FAAPA">FAAPA</option>
                <option value="SUTES">SUTES</option>
                <option value="Estudiante">Estudiante</option>
              </select>
            </div>

            {/* Nombre */}
            <FormInput
              type="text"
              name="name"
              placeholder="Nombre"
              required={true}
              value={name()}
              onInput={(e: InputEvent) => setName((e.target as HTMLInputElement).value)}
            />

            {/* Apellido Paterno */}
            <FormInput
              type="text"
              name="lastname1"
              placeholder="Apellido Paterno"
              required={true}
              value={lastname1()}
              onInput={(e: InputEvent) => setLastname1((e.target as HTMLInputElement).value)}
            />

            {/* Apellido Materno */}
            <FormInput
              type="text"
              name="lastname2"
              placeholder="Apellido Materno"
              required={true}
              value={lastname2()}
              onInput={(e: InputEvent) => setLastname2((e.target as HTMLInputElement).value)}
            />

            {/* CURP */}
            <FormInput
              type="text"
              name="curp"
              placeholder="CURP"
              required={true}
              value={curp()}
              onInput={(e: InputEvent) => setCurp((e.target as HTMLInputElement).value)}
            />

            {/* Sexo */}
            <div class="input-group select-wrapper">
              <label for="sex">Sexo</label>
              <select name="sex" id="sex" required>
                <option value="">Selecciona tu sexo</option>
                <option value="Masculino">Masculino</option>
                <option value="Femenino">Femenino</option>
                <option value="Otro">Otro</option>
              </select>
            </div>

            {/* Teléfono */}
            <FormInput
              type="tel"
              name="phone"
              placeholder="Teléfono"
              required={true}
              value={phone()}
              onInput={(e: InputEvent) => setPhone((e.target as HTMLInputElement).value)}
            />

            {/* Correo Electrónico */}
            <FormInput
              type="email"
              name="email"
              placeholder="Correo electrónico"
              required={true}
              value={email()}
              onInput={(e: InputEvent) => setEmail((e.target as HTMLInputElement).value)}
            />

            {/* Contraseña */}
            <FormInput
              type="password"
              name="password"
              placeholder="Contraseña"
              required={true}
              value={password()}
              onInput={(e: InputEvent) => setPassword((e.target as HTMLInputElement).value)}
            />

            <button type="submit" class="form-button">Registrar</button>
          </form>

          <div class="form-links">
            <p>¿Ya tienes una cuenta? <a href="#login">Iniciar sesión</a></p>
          </div>
        </div>
      </div>
    </>
  );
};

export default Register;
