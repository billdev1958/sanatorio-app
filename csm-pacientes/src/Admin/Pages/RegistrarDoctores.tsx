import { createSignal } from "solid-js";

type RegisterDoctorRequest = {
  afiliationID: number | "";
  medicalLicense: string;
  specialtyLicense: string;
  name: string;
  lastname1: string;
  lastname2: string;
  sex: string;
  phoneNumber: string;
  email: string;
  password: string;
};

function RegisterDoctorForm() {
  const [formData, setFormData] = createSignal<RegisterDoctorRequest>({
    afiliationID: "",
    medicalLicense: "",
    specialtyLicense: "",
    name: "",
    lastname1: "",
    lastname2: "",
    sex: "",
    phoneNumber: "",
    email: "",
    password: "",
  });

  const handleChange = (
    e: Event & { currentTarget: HTMLInputElement | HTMLSelectElement }
  ): void => {
    const { name, value } = e.currentTarget;
    setFormData({
      ...formData(),
      [name]: name === "afiliationID" ? +value || "" : value,
    });
  };

  const handleSubmit = (e: Event): void => {
    e.preventDefault();
    console.log("Doctor Data:", formData());
    // Aquí puedes enviar los datos a tu API o procesarlos
  };

  return (
    <div class="form-container">
      <form onSubmit={handleSubmit}>
        <div class="form-group">
          <label for="afiliationID">Afiliation ID</label>
          <input
            type="number"
            id="afiliationID"
            name="afiliationID"
            required
            value={formData().afiliationID}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="medicalLicense">Medical License</label>
          <input
            type="text"
            id="medicalLicense"
            name="medicalLicense"
            required
            value={formData().medicalLicense}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="specialtyLicense">Specialty License</label>
          <input
            type="text"
            id="specialtyLicense"
            name="specialtyLicense"
            required
            value={formData().specialtyLicense}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="name">Name</label>
          <input
            type="text"
            id="name"
            name="name"
            required
            value={formData().name}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="lastname1">First Lastname</label>
          <input
            type="text"
            id="lastname1"
            name="lastname1"
            required
            value={formData().lastname1}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="lastname2">Second Lastname</label>
          <input
            type="text"
            id="lastname2"
            name="lastname2"
            value={formData().lastname2}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="sex">Sex</label>
          <select
            id="sex"
            name="sex"
            required
            value={formData().sex}
            onInput={(e) => handleChange(e as Event & { currentTarget: HTMLSelectElement })}
          >
            <option value="">Select...</option>
            <option value="M">Male</option>
            <option value="F">Female</option>
          </select>
        </div>

        <div class="form-group">
          <label for="phoneNumber">Phone Number</label>
          <input
            type="tel"
            id="phoneNumber"
            name="phoneNumber"
            required
            value={formData().phoneNumber}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="email">Email</label>
          <input
            type="email"
            id="email"
            name="email"
            required
            value={formData().email}
            onInput={handleChange}
          />
        </div>

        <div class="form-group">
          <label for="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            required
            value={formData().password}
            onInput={handleChange}
          />
        </div>

        <button type="submit">Register Doctor</button>
      </form>

      <style>
        {`
          .form-container {
            max-width: 400px;
            margin: 0 auto;
            padding: 1em;
            border: 1px solid #ccc;
            border-radius: 5px;
            background: #f9f9f9;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
          }
          .form-group {
            margin-bottom: 1em;
          }
          label {
            display: block;
            margin-bottom: 0.5em;
            font-weight: bold;
          }
          input, select {
            width: 100%;
            padding: 0.5em;
            border: 1px solid #ccc;
            border-radius: 4px;
          }
          button {
            padding: 0.7em 1.5em;
            background: #007bff;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
          }
          button:hover {
            background: #0056b3;
          }
        `}
      </style>
    </div>
  );
}

export default RegisterDoctorForm;
