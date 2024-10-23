// src/components/Section.tsx
import { Component, createSignal, JSX } from 'solid-js';

interface SectionProps {
  title: string;
  children: JSX.Element;
}

const Section: Component<SectionProps> = (props) => {
  const [isOpen, setIsOpen] = createSignal(false);

  return (
    <div class="section">
      <div class="section-header" onClick={() => setIsOpen(!isOpen())}>
        <h2>{props.title}</h2>
        <button class="toggle-btn">{isOpen() ? '-' : '+'}</button>
      </div>
      {isOpen() && <div class="section-content">{props.children}</div>}
    </div>
  );
};

export default Section;
