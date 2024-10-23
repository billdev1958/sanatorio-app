// src/components/widgets/Widget.tsx
import { Component } from 'solid-js';

interface WidgetProps {
  icon: string;        
  title: string;        
  description: string;  
  link: string;         
  buttonText: string;   
}

const Widget: Component<WidgetProps> = (props) => {
  return (
    <div class="widget">
      <h3><i class={props.icon}></i> {props.title}</h3>
      <p>{props.description}</p>
      <a href={props.link} class="btn">{props.buttonText}</a>
    </div>
  );
};

export default Widget;
