import { Component, For } from 'solid-js';

interface InfoPanelProps {
  title: string;
  fields: { label: string; value: string }[];
}

const InfoPanel: Component<InfoPanelProps> = (props) => {
  return (
    <div class="info-panel">
      <h2><i class="fas fa-notes-medical"></i> {props.title}</h2>
      <For each={props.fields}>
        {(field) => (
          <p><strong>{field.label}:</strong> {field.value}</p>
        )}
      </For>
    </div>
  );
};

export default InfoPanel;
