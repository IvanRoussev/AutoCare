import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { message, Timeline } from 'antd';
import '../../styles/timeline.css';

function ShowTimeline() {
  const location = useLocation();
  const { car } = location.state || {};
  console.log(car);

  const token = localStorage.getItem('token');

  const [maintenanceData, setMaintenanceData] = useState([]);

  useEffect(() => {
    const createTimeline = async () => {
      const url = `http://localhost:8080/maintenances/${car.vin}?page_id=1&page_size=5`;
      try {
        const response = await fetch(url, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();

          let maintenances = [];
          for (let i = 0; i < data.length; i++) {
            let maintenanceType = data[i]['maintenance_type'];
            let date = data[i]['created_at'];
            let mileage = data[i]['mileage'];

            let sentence = `${maintenanceType} done on ${date} at ${mileage}km`;
            maintenances.push({ children: sentence }); // Push new items to the array
          }

          setMaintenanceData(maintenances); // Set the array to the state
        } else {
          console.error('Error Creating Timeline:', response.statusText);
          const errorMessage = `Error Creating Timeline: ${car.year} ${car.make} ${car.model}`;
          message.error(errorMessage);
        }
      } catch (error) {
        console.error('Error Creating Timeline::', error);
        const errorMessage = `Error Creating Timeline: ${car.year} ${car.make} ${car.model}`;
        message.error(errorMessage);
      }
    };
    createTimeline();
  }, [car, token]);
  console.log(maintenanceData);

  return (
    <div className='timeline-page'>
      <div>
        <h1 className='car-title-maintenanceLog'>
          {car.make} {car.model} Maintenance Log
        </h1>
      </div>
      <div className='timeline'>
        <Timeline items={maintenanceData} />
      </div>
    </div>
  );
}

export default ShowTimeline;
