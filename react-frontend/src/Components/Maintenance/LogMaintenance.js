import React from 'react';
import { Select, Input, Button, message } from 'antd';
import { useLocation } from 'react-router-dom';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const CreateaMaintenanceLog = async (maintenance, mileage, vin, navigate) => {
  const token = localStorage.getItem('token');
  const username = localStorage.getItem('username');

  try {
    const response = await fetch('http://localhost:8080/maintenances', {
      method: 'POST',
      body: JSON.stringify({
        car_vin: vin,
        maintenance_type: maintenance,
        mileage: mileage,
      }),
      headers: {
        'Content-type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (response.ok) {
      const successMessage = `Successfully Added Maintenance Record ${maintenance} at ${mileage}km`;
      message.success(successMessage);
      navigate(`/home/${username}`);
    } else {
      console.error('Failed to add maintenance Record:', response.statusText);
    }
  } catch (error) {
    console.error('Error during POST request:', error);
  }
};

function LogMaintenance() {
  const location = useLocation();
  const { car } = location.state || {};

  const [selectedMaintenance, setSelectedMaintenance] = useState([]);
  const [mileage, setMileage] = useState('');

  const options = [
    { value: 'Oil-Change', label: 'Oil Change' },
    { value: 'Tire-Rotation', label: 'Tire Rotation' },
    { value: 'General-Checkup', label: 'General Checkup' },
    { value: 'Engine-Swap', label: 'Engine Swap' },
    { value: 'Transmission', label: 'Transmission' },
    { value: 'Paint', label: 'Paint' },
  ];

  const handleChange = (value) => {
    setSelectedMaintenance(value);
  };

  const navigate = useNavigate();

  const handleCreateMaintenance = (vin) => {
    for (let i = 0; i < selectedMaintenance.length; i++) {
      let maintenanceType = selectedMaintenance[i];
      let mileageInt = parseInt(mileage);
      if (isNaN(mileageInt)) {
        message.error('Mileage needs to be a valid number');
        return;
      }
      CreateaMaintenanceLog(maintenanceType, mileageInt, vin, navigate);
    }
  };

  return (
    <div>
      <h1>Log your Maintenance</h1>
      <h2>
        {car.make} {car.model} | {car.year}
      </h2>
      <h3>Enter type of maintenances done</h3>
      <Select
        mode='tags'
        style={{
          width: '100%',
        }}
        placeholder='Tags Mode'
        onChange={handleChange}
        options={options}
      />
      <h4>Enter Mileage done at</h4>
      <Input
        type='number'
        value={mileage}
        onChange={(e) => setMileage(e.target.value)}
      />
      <Button type='primary' onClick={() => handleCreateMaintenance(car.vin)}>
        Primary Button
      </Button>
    </div>
  );
}

export default LogMaintenance;
