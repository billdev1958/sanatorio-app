import { Component } from 'solid-js';

interface FormInputProps {
  type: string;
  name: string;
  placeholder: string;
  required?: boolean;
}

const FormInput: Component<FormInputProps> = (props) => {
  return (
    <div class="input-group">
      <input
        type={props.type}
        name={props.name}
        placeholder={props.placeholder}
        required={props.required}
      />
    </div>
  );
};

export default FormInput;
