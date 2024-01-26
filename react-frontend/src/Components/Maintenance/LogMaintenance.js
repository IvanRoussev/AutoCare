import React from 'react';
import { Select, Input, Button, message } from 'antd';
import { useLocation } from 'react-router-dom';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../../styles/LogMaintenance.css';

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

  const handleBackToHomeClick = () => {
    const username = localStorage.getItem('username');
    navigate(`/home/${username}`);
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
    <div className='logMaintenance-page'>
      <div className='titles-maintenance'>
        <h1>Log your Maintenance</h1>
        <h3>Enter type of maintenances done</h3>
      </div>
      <div className='maintenance-container'>
        <h2>
          {car.make} {car.model} | {car.year}
        </h2>
        <div className='maintenance-form'>
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
          <button
            className='back-to-cars-from-logMaintenance'
            onClick={handleBackToHomeClick}
          >
            Back to Cars
          </button>
          <Button
            type='primary'
            className='button-create-maintenance'
            onClick={() => handleCreateMaintenance(car.vin)}
          >
            Log Maintenance
          </Button>
        </div>
      </div>
    </div>
  );
}

export default LogMaintenance;
