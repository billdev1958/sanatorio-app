import { Component } from 'solid-js';

interface FormInputProps {
  type: string;
  name: string;
  placeholder: string;
  required?: boolean;
  value: string; // Agrega la propiedad value
  onInput: (e: InputEvent) => void; // Agrega el manejador onInput
}

const FormInput: Component<FormInputProps> = (props) => {
  return (
    <div class="input-group">
      <input
        type={props.type}
        name={props.name}
        placeholder={props.placeholder}
        required={props.required}
        value={props.value} // Conecta el valor del input
        onInput={props.onInput} // Maneja el evento onInput
      />
    </div>
  );
};

export default FormInput;
