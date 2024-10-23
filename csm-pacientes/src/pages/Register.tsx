import FormInput from '../components/FormInput'; // Importa el componente reutilizable

const Register = () => {
  return (
    <>
      <div class="form-container">
        <div class="form-card">
          <h2>Registrar Usuario</h2>
          <form class="form">
            {/* AfiliationID select */}
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
            />

            {/* Apellido Paterno */}
            <FormInput
              type="text"
              name="lastname1"
              placeholder="Apellido Paterno"
              required={true}
            />

            {/* Apellido Materno */}
            <FormInput
              type="text"
              name="lastname2"
              placeholder="Apellido Materno"
              required={true}
            />

            {/* CURP */}
            <FormInput
              type="text"
              name="curp"
              placeholder="CURP"
              required={true}
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
            />

            {/* Correo Electrónico */}
            <FormInput
              type="email"
              name="email"
              placeholder="Correo electrónico"
              required={true}
            />

            {/* Contraseña */}
            <FormInput
              type="password"
              name="password"
              placeholder="Contraseña"
              required={true}
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
