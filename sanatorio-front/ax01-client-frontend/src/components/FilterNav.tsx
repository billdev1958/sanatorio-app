import React, { useState } from 'react';

type FilterNavProps = {
  onFilterChange: (filters: { role?: number; name?: string; date?: string }) => void;
};

const FilterNav = ({ onFilterChange }: FilterNavProps) => {
  const [role, setRole] = useState<number | undefined>(undefined);
  const [name, setName] = useState('');
  const [date, setDate] = useState('');

  const handleRoleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedRole = e.target.value ? parseInt(e.target.value) : undefined;
    setRole(selectedRole);
    onFilterChange({ role: selectedRole, name, date });
  };

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
    onFilterChange({ role, name: e.target.value, date });
  };

  const handleDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setDate(e.target.value);
    onFilterChange({ role, name, date: e.target.value });
  };

  return (
    <div className="filterNav">
      <div className="filterItem">
        <label htmlFor="role">Rol</label>
        <select id="role" value={role || ''} onChange={handleRoleChange}>
          <option value="">Todos</option>
          <option value="1">Administrador</option>
          <option value="2">Usuario</option>
          <option value="3">Doctor</option>
        </select>
      </div>

      <div className="filterItem">
        <label htmlFor="name">Nombre</label>
        <input
          type="text"
          id="name"
          value={name}
          placeholder="Buscar por nombre"
          onChange={handleNameChange}
        />
      </div>

      <div className="filterItem">
        <label htmlFor="date">Fecha de creación</label>
        <input type="date" id="date" value={date} onChange={handleDateChange} />
      </div>
    </div>
  );
};

export default FilterNav;
