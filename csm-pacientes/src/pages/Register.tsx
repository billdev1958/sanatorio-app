import { createSignal } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import NavBar from '../components/NavBar'; // Importa el NavBar
import FormInput from '../components/FormInput'; // Importa el componente reutilizable
import { registerUser } from '../services/RegisterService'; // Importa el servicio de registro
import { RegisterPatientRequest } from '../models/Login&Registers';
import Loader from '../core/components/Loader';
import { useMessage } from '../core/domain/messageProvider'; // Contexto para mensajes de éxito
import AlertMessage from '../core/components/AlertMessage'; // Nuestro AlertMessage con auto-cierre

const Register = () => {
  const [name, setName] = createSignal<string>("");
  const [lastname1, setLastname1] = createSignal<string>("");
  const [lastname2, setLastname2] = createSignal<string>("");
  const [curp, setCurp] = createSignal<string>("");
  const [sex, setSex] = createSignal<string>("");
  const [phone, setPhone] = createSignal<string>("");
  const [email, setEmail] = createSignal<string>("");
  const [password, setPassword] = createSignal<string>("");
  const [registerError, setRegisterError] = createSignal<string | null>(null);
  const [loading, setLoading] = createSignal(false);

  const navigate = useNavigate();
  const { setSuccessMessage } = useMessage();

  const handleSubmit = async (e: Event) => {
    e.preventDefault();
    setRegisterError(null);
    setLoading(true);

    const user: RegisterPatientRequest = {
      dependency_id: 1,
      name: name(),
      lastname1: lastname1(),
      lastname2: lastname2(),
      curp: curp(),
      sex: sex(),
      phone: phone(),
      email: email(),
      password: password(),
    };

    try {
      await registerUser(user);
      setSuccessMessage("Registro exitoso, por favor inicia sesión.");
      navigate('/login', { replace: true });
    } catch (error: any) {
      setRegisterError(error.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      {loading() && <Loader fullScreen={true} size="60px" />}

      {registerError() && (
        <AlertMessage
          type="error"
          message={registerError()!}
          onClose={() => setRegisterError(null)}
        />
      )}

      <NavBar toggleMenu={() => {}} />

      <div class="form-container">
        <div class="form-card">
          <h2>Registrar Usuario</h2>
          <form class="form" onSubmit={handleSubmit}>
            <div class="input-group select-wrapper">
              <label for="afiliation">Afiliación</label>
              <select name="afiliation" id="afiliation" required onInput={(e: InputEvent) => setSex((e.target as HTMLSelectElement).value)}>
                <option value="">Selecciona una opción</option>
                <option value="1">Confianza</option>
                <option value="2">FAAPA</option>
                <option value="3">SUTES</option>
                <option value="4">Estudiante</option>
                <option value="5">Externo</option>
              </select>
            </div>

            <FormInput
              type="text"
              name="name"
              placeholder="Nombre"
              required={true}
              value={name()}
              onInput={(e: InputEvent) => setName((e.target as HTMLInputElement).value)}
            />

            <FormInput
              type="text"
              name="lastname1"
              placeholder="Apellido Paterno"
              required={true}
              value={lastname1()}
              onInput={(e: InputEvent) => setLastname1((e.target as HTMLInputElement).value)}
            />

            <FormInput
              type="text"
              name="lastname2"
              placeholder="Apellido Materno"
              required={true}
              value={lastname2()}
              onInput={(e: InputEvent) => setLastname2((e.target as HTMLInputElement).value)}
            />

            <FormInput
              type="text"
              name="curp"
              placeholder="CURP"
              required={true}
              value={curp()}
              onInput={(e: InputEvent) => setCurp((e.target as HTMLInputElement).value)}
            />

            <div class="input-group select-wrapper">
              <label for="sex">Sexo</label>
              <select name="sex" id="sex" required onInput={(e: InputEvent) => setSex((e.target as HTMLSelectElement).value)}>
                <option value="">Selecciona tu sexo</option>
                <option value="M">Masculino</option>
                <option value="F">Femenino</option>
                <option value="Otro">Otro</option>
              </select>
            </div>

            <FormInput
              type="tel"
              name="phone"
              placeholder="Teléfono"
              required={true}
              value={phone()}
              onInput={(e: InputEvent) => setPhone((e.target as HTMLInputElement).value)}
            />

            <FormInput
              type="email"
              name="email"
              placeholder="Correo electrónico"
              required={true}
              value={email()}
              onInput={(e: InputEvent) => setEmail((e.target as HTMLInputElement).value)}
            />

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

          {/* Eliminamos el <p class="error-message"> y usamos AlertMessage en su lugar */}

          <div class="form-links">
            <p>¿Ya tienes una cuenta? <a href="#login">Iniciar sesión</a></p>
          </div>
        </div>
      </div>
    </>
  );
};

export default Register;
