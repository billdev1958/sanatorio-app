import { createSignal } from "solid-js";

type RegisterOfficeRequest = {
  name: string;
};

function RegisterOfficeForm() {
  const [formData, setFormData] = createSignal<RegisterOfficeRequest>({
    name: "",
  });

  const handleChange = (
    e: Event & { currentTarget: HTMLInputElement }
  ): void => {
    const { name, value } = e.currentTarget;
    setFormData({ ...formData(), [name]: value });
  };

  const handleSubmit = (e: Event): void => {
    e.preventDefault();
    console.log("Office Data:", formData());
  };

  return (
    <div class="form-container">
      <form onSubmit={handleSubmit}>
        <div class="form-group">
          <label for="name">Office Name</label>
          <input
            type="text"
            id="name"
            name="name"
            required
            value={formData().name}
            onInput={handleChange}
          />
        </div>

        <button type="submit">Register Office</button>
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
          input {
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

export default RegisterOfficeForm;
