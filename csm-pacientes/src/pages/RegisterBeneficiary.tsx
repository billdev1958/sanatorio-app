// RegisterBeneficiary.tsx
import { createSignal } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import FormInput from '../components/FormInput';
import { registerBeneficiary } from '../services/RegisterService';
import { RegisterBeneficiaryRequest } from '../models/Login&Registers';
import { useAuth } from '../services/AuthContext';

const RegisterBeneficiary = () => {
  const [name, setName] = createSignal<string>("");
  const [lastname1, setLastname1] = createSignal<string>("");
  const [lastname2, setLastname2] = createSignal<string>("");
  const [curp, setCurp] = createSignal<string>("");
  const [sex, setSex] = createSignal<string>("");
  const [registerError, setRegisterError] = createSignal<string | null>(null);
  const navigate = useNavigate();
  const { token } = useAuth();

  const handleSubmit = async (e: Event) => {
    e.preventDefault();
    setRegisterError(null);

    const beneficiary: RegisterBeneficiaryRequest = {
      name: name(),
      lastname1: lastname1(),
      lastname2: lastname2(),
      curp: curp(),
      sex: sex(),
    };

    try {
      await registerBeneficiary(beneficiary, token() ?? undefined); // Convierte null a undefined
      navigate('/login', { replace: true });
    } catch (error: any) {
      setRegisterError(error.message);
    }
  };

  return (
    <div class="form-container">
      <div class="form-card">
        <h2>Registrar Usuario</h2>
        <form class="form" onSubmit={handleSubmit}>
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
          <button type="submit" class="form-button">Registrar</button>
        </form>

        {registerError() && <p class="error-message">Error: {registerError()}</p>}
      </div>
    </div>
  );
};

export default RegisterBeneficiary;
